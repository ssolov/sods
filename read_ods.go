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

	textStart    = "<text:"
	textStartLen = len(textStart)

	textSpace    = "<text:s/>"
	textSpaceLen = len(textSpace)

	textSpaceCount    = "text:c=\""
	textSpaceCountLen = len(textSpaceCount)

	cellRepeat    = "table:number-columns-repeated=\""
	cellRepeatLen = len(cellRepeat)
)

// Cell is a cell of the ods table
type Cell string

// Row is a row with cells of the ods table
type Row []Cell

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

// Returns the content between first begin and last end pattern inclusiv these patterns
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
	// split rows by the begin pattern
	rawRows := bytes.Split(rawTable, []byte(rowStart))
	if len(rawRows) > 0 {
		rows = make([]Row, 0, len(rows))

		for _, rawRow := range rawRows {
			if len(rawRow) < 1 {
				continue
			}

			var err error
			// get cells and ignore all empty cells after last non empty cell
			if rawRow, err = cutElementContent(rawRow, []byte(cellStart), []byte(cellEnd)); err != nil {
				continue
			}

			row := parseCells(rawRow)
			if len(row) > 0 {
				rows = append(rows, row)
			}
		}
	}

	return rows
}

func parseCells(rawRow []byte) Row {
	rawCells := bytes.Split(rawRow, []byte(cellStart))
	row := make(Row, 0, len(rawCells))

	for _, rawCell := range rawCells {
		if len(rawCell) < 1 {
			continue
		}

		cell := Cell(parseCell(rawCell))
		row = append(row, cell)

		// search and copy repeted cells
		rpCells, err := repeatCell(rawCell, cell)
		if err == nil && len(rpCells) > 0 {
			row = append(row, rpCells...)
		}
	}

	return row
}

func parseCell(cell []byte) string {
	bg := bytes.Index(cell, []byte(textStart))
	if bg != -1 {
		return html.EscapeString(addNSpaces(parseElementContent(cell, bg+textStartLen)))
	}

	return ""
}

func addNSpaces(txt string) string {
	txt = strings.Replace(txt, textSpace, " ", -1)

	esbIdx := strings.Index(txt, textSpaceCount)
	if esbIdx > -1 {
		txt = txt[esbIdx+textSpaceCountLen:]
		eseIdx := strings.Index(txt, "/>")
		if eseIdx > -1 {
			spc := parseAttrValue([]byte(txt[:eseIdx]), 0)
			txt = txt[eseIdx+2:]
			if spc != "" {
				n, err := strconv.Atoi(spc)
				if err == nil {
					return strings.Repeat(" ", n) + txt
				}
			}
		}
	}

	return txt
}

func repeatCell(cont []byte, c Cell) (Row, error) {
	var row Row
	bg := bytes.Index(cont, []byte(cellRepeat))
	if bg != -1 {
		rpCount, err := strconv.Atoi(parseAttrValue(cont, bg+cellRepeatLen))
		if err != nil {
			return row, err
		}

		row = make(Row, 0, rpCount-1)
		for i := 1; i < rpCount; i++ {
			// copy the cell append n time
			row = append(row, c)
		}
	}

	return row, nil
}

// Read function opens and reads the content of ods file
func Read(odsFileName string) (*File, error) {
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

	return ParseContent(content)
}

// ParseContent function will parse the content of content.xml
func ParseContent(content []byte) (*File, error) {
	var odsFile *File

	// found first begin and last end of table
	content, _ = cutElementContent(content, []byte(tableStart), []byte(tableEnd))
	// split tables by begin pattern
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
		// read the name of table
		sh.Name = parseAttrValue(tbl, bytes.Index(tbl, []byte(tableName))+tableNameLen)

		var err error
		// only get rows from this table
		if tbl, err = cutElementContent(tbl, []byte(rowStart), []byte(rowEnd)); err != nil {
			continue
		}

		sh.Rows = parseRows(tbl)
		odsFile.Sheets = append(odsFile.Sheets, sh)
	}

	return odsFile, nil

}
