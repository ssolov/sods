package sods

import (
	"testing"
)

func TestParseCell1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	cell := parseCell([]byte(`
								<table:table-cell office:value-type="string">
									<text:p>BlaBlaBla</text:p>
								</table:table-cell>`))
	if cell.Text != "BlaBlaBla" {
		t.Fail()
	}
}

func TestParseCell2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	cell := parseCell([]byte(`<table:table-cell table:style-name="ce15"/>`))
	if len(cell.Text) != 0 {
		t.Fail()
	}
}

func TestParseCells1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	cells := parseCells([]byte(`
			<table:table-cell office:value-type="string">
				<text:p>BlaBlaBla</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>BlaBlaBla</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>BlaBlaBla</text:p>
			</table:table-cell>
			<table:table-cell/>
			<table:table-cell table:number-columns-repeated="2" office:value-type="string">
				<text:p>true</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>BlaBlaBla</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>BlaBlaBla</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>BlaBlaBla</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="float" office:value="2">
				<text:p>2</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="float" office:value="365">
				<text:p>365</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="float" office:value="30">
				<text:p>30</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="float" office:value="480">
				<text:p>480</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="float" office:value="2">
				<text:p>2</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>1,2,3,4,5,6,7,8,9,10,11,12,13,14</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>N/A</text:p>
			</table:table-cell>
			<table:table-cell table:style-name="ce14" table:formula="of:=&quot;2019-03-31&quot;" office:value-type="string" office:string-value="2019-03-31">
				<text:p>2019-03-31</text:p>
			</table:table-cell>
			<table:table-cell table:number-columns-repeated="2" office:value-type="string">
				<text:p>N/A</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>BlaBlaBla</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>N/A</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>BlaBlaBla</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="string">
				<text:p>N/A</text:p>
			</table:table-cell>
			<table:table-cell office:value-type="float" office:value="3">
				<text:p>3</text:p>
			</table:table-cell>
			<table:table-cell table:number-columns-repeated="2" office:value-type="float" office:value="1">
				<text:p>1</text:p>
			</table:table-cell>`))

	if len(cells) != 26 {
		t.Fail()
	}
}

func TestCutElementContent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	beforCut := []byte(`.......<table:table-cell office:value-type="string"><text:p>BlaBlaBla</text:p></table:table-cell>........`)
	afterCut := `<table:table-cell office:value-type="string"><text:p>BlaBlaBla</text:p></table:table-cell>`
	cnt, _ := cutElementContent(beforCut, []byte(cellStart), []byte(cellEnd))
	t.Log(string(cnt))

	if string(cnt) != afterCut {
		t.Fail()
	}
}
