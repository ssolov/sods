package sods

import (
	"archive/zip"
	"bytes"
	"fmt"
	"html"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	tableStart    = "<table:table "
	tableStartLen = len(tableStart)

	tableEnd    = "</table:table>"
	tableEndLen = len(tableEnd)

	tableName    = "table:name=\""
	tableNameLen = len(tableName)

	rowStart    = "<table:table-row "
	rowStartLen = len(rowStart)

	rowEnd    = "</table:table-row>"
	rowEndLen = len(rowEnd)

	cellStart    = "<table:table-cell"
	cellStartLen = len(cellStart)

	cellEnd    = "</table:table-cell>"
	cellEndLen = len(cellEnd)

	cellTypePara    = "office:value-type=\""
	cellTypeParaLen = len(cellTypePara)

	cellValuePara    = "office:value=\""
	cellValueParaLen = len(cellValuePara)

	textStart    = "<text:"
	textStartLen = len(textStart)

	cellRepeat    = "table:number-columns-repeated=\""
	cellRepeatLen = len(cellRepeat)

	cellFormula    = "table:formula=\""
	cellFormulaLen = len(cellFormula)
)

// Cell is a cell of the ods table
type Cell struct {
	Type    string // office:value-type="string"
	Value   string // office:value="2"
	Formula string // table:formula="of:=&quot;2019-03-31&quot;"
	Text    string // <text:p>2019-03-31</text:p>

}

// Row is a row with cells of the ods table
type Row struct {
	Cells []Cell
}

// Sheet is the ods sheet. Wich contains the rows.
type Sheet struct {
	Name string
	Rows []Row
}

// File is a holder for the ods sheets
type File struct {
	Sheets []Sheet
}

func getContentXML(r *zip.ReadCloser) ([]byte, error) {
	for _, f := range r.File {
		if f.Name == "content.xml" {
			rc, err := f.Open()
			if err != nil {
				return []byte{}, err
			}

			return ioutil.ReadAll(rc)
		}
	}

	return []byte{}, fmt.Errorf("content.xml not found")
}

func parseAttrValue(attr []byte, start int) string {
	var v string
	if len(attr) <= start {
		return v
	}

	attr = attr[start:]
	end := bytes.Index(attr, []byte("\""))
	if end != -1 {
		v = string(attr[:end])
	}

	return v
}

func parseElementContent(elm []byte, start int) string {
	var c string
	if len(elm) <= start {
		return c
	}

	elm = elm[start:]
	start = bytes.Index(elm, []byte(">"))
	if start != -1 {
		elm = elm[start+1:]
		end := bytes.Index(elm, []byte("</"))
		if end != -1 {
			c = string(elm[:end])
		}
	}

	return c
}

func cutElementContent(elm, begin, end []byte) ([]byte, error) {
	i := bytes.Index(elm, begin)
	if i < 0 {
		return elm, fmt.Errorf("'%s' in elment not found", string(begin))
	}

	elm = elm[i:]
	i = bytes.LastIndex(elm, end)
	if i < 0 {
		return elm, fmt.Errorf("'%s' in elment not found", string(begin))
	}

	return elm[:i+len(end)], nil

}

func parseRows(rawTable []byte) []Row {
	var rows []Row
	rawRows := bytes.Split(rawTable, []byte(rowStart))
	if len(rawRows) > 0 {
		rows = make([]Row, 0, len(rows))

		for _, rawRow := range rawRows {
			if len(rawRow) < 1 {
				continue
			}

			row := Row{}
			var err error
			if rawRow, err = cutElementContent(rawRow, []byte(cellStart), []byte(cellEnd)); err != nil {
				continue
			}

			row.Cells = parseCells(rawRow)
			if len(row.Cells) > 0 {
				rows = append(rows, row)
			}
		}
	}

	return rows
}

func parseCells(rawRow []byte) []Cell {
	rawCells := bytes.Split(rawRow, []byte(cellStart))
	cells := make([]Cell, 0, len(rawCells))

	for _, rawCell := range rawCells {
		if len(rawCell) < 1 {
			continue
		}

		cell := parseCell(rawCell)
		cells = append(cells, cell)

		rpCells, err := repeatCell(rawCell, cell)
		if err == nil && len(rpCells) > 0 {
			cells = append(cells, rpCells...)
		}
	}

	return cells
}

// parseCell ...
func parseCell(cell []byte) Cell {
	var c = Cell{}

	bg := bytes.Index(cell, []byte(cellTypePara))
	if bg != -1 {
		c.Type = parseAttrValue(cell, bg+cellTypeParaLen)
	}

	bg = bytes.Index(cell, []byte(cellValuePara))
	if bg != -1 {
		c.Value = parseAttrValue(cell, bg+cellValueParaLen)
	}

	bg = bytes.Index(cell, []byte(cellFormula))
	if bg != -1 {
		c.Formula = html.EscapeString(parseAttrValue(cell, bg+cellFormulaLen))
	}

	bg = bytes.Index(cell, []byte(textStart))
	if bg != -1 {
		c.Text = parseElementContent(cell, bg+textStartLen)
		c.Text = strings.Replace(c.Text, "<text:s/>", " ", -1)
		c.Text = html.EscapeString(c.Text)
	}

	return c
}

func repeatCell(cont []byte, c Cell) ([]Cell, error) {
	var cells []Cell
	bg := bytes.Index(cont, []byte(cellRepeat))
	if bg != -1 {
		rpCount, err := strconv.Atoi(parseAttrValue(cont, bg+cellRepeatLen))
		if err != nil {
			return cells, err
		}

		cells = make([]Cell, 0, rpCount-1)
		for i := 1; i < rpCount; i++ {
			cells = append(cells, c)
		}
	}

	return cells, nil
}

// Open function opens and reads the content of given ods file
func Open(odsFileName string) (*File, error) {
	var odsFile *File

	r, err := zip.OpenReader(odsFileName)
	if err != nil {
		return odsFile, err
	}
	defer r.Close()

	content, err := getContentXML(r)
	if err != nil {
		return odsFile, err
	}

	content, _ = cutElementContent(content, []byte(tableStart), []byte(tableEnd))
	tables := bytes.Split(content, []byte(tableStart))
	if len(tables) < 1 {
		return odsFile, nil
	}

	odsFile = new(File)
	odsFile.Sheets = make([]Sheet, 0, len(tables))

	for _, tbl := range tables {
		if len(tbl) < 1 {
			continue
		}

		sh := Sheet{}
		sh.Name = parseAttrValue(tbl, bytes.Index(tbl, []byte(tableName))+tableNameLen)

		var err error
		if tbl, err = cutElementContent(tbl, []byte(rowStart), []byte(rowEnd)); err != nil {
			continue
		}

		sh.Rows = parseRows(tbl)
		odsFile.Sheets = append(odsFile.Sheets, sh)
	}

	return odsFile, nil
}
