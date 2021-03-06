package sods

import (
	"testing"
)

const contentXML = `
<?xml version="1.0" encoding="UTF-8"?>
<office:document-content xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0" xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0" xmlns:table="urn:oasis:names:tc:opendocument:xmlns:table:1.0" xmlns:draw="urn:oasis:names:tc:opendocument:xmlns:drawing:1.0" xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0" xmlns:number="urn:oasis:names:tc:opendocument:xmlns:datastyle:1.0" xmlns:presentation="urn:oasis:names:tc:opendocument:xmlns:presentation:1.0" xmlns:svg="urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0" xmlns:chart="urn:oasis:names:tc:opendocument:xmlns:chart:1.0" xmlns:dr3d="urn:oasis:names:tc:opendocument:xmlns:dr3d:1.0" xmlns:math="http://www.w3.org/1998/Math/MathML" xmlns:form="urn:oasis:names:tc:opendocument:xmlns:form:1.0" xmlns:script="urn:oasis:names:tc:opendocument:xmlns:script:1.0" xmlns:ooo="http://openoffice.org/2004/office" xmlns:ooow="http://openoffice.org/2004/writer" xmlns:oooc="http://openoffice.org/2004/calc" xmlns:dom="http://www.w3.org/2001/xml-events" xmlns:xforms="http://www.w3.org/2002/xforms" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:rpt="http://openoffice.org/2005/report" xmlns:of="urn:oasis:names:tc:opendocument:xmlns:of:1.2" xmlns:xhtml="http://www.w3.org/1999/xhtml" xmlns:grddl="http://www.w3.org/2003/g/data-view#" xmlns:tableooo="http://openoffice.org/2009/table" xmlns:drawooo="http://openoffice.org/2010/draw" xmlns:calcext="urn:org:documentfoundation:names:experimental:calc:xmlns:calcext:1.0" xmlns:loext="urn:org:documentfoundation:names:experimental:office:xmlns:loext:1.0" xmlns:field="urn:openoffice:names:experimental:ooo-ms-interop:xmlns:field:1.0" xmlns:formx="urn:openoffice:names:experimental:ooxml-odf-interop:xmlns:form:1.0" xmlns:css3t="http://www.w3.org/TR/css3-text/" office:version="1.2"><office:scripts/><office:font-face-decls><style:font-face style:name="Liberation Sans" svg:font-family="&apos;Liberation Sans&apos;" style:font-family-generic="swiss" style:font-pitch="variable"/><style:font-face style:name="DejaVu Sans" svg:font-family="&apos;DejaVu Sans&apos;" style:font-family-generic="system" style:font-pitch="variable"/><style:font-face style:name="Lohit Devanagari" svg:font-family="&apos;Lohit Devanagari&apos;" style:font-family-generic="system" style:font-pitch="variable"/><style:font-face style:name="Noto Sans CJK SC Regular" svg:font-family="&apos;Noto Sans CJK SC Regular&apos;" style:font-family-generic="system" style:font-pitch="variable"/></office:font-face-decls><office:automatic-styles><style:style style:name="co1" style:family="table-column"><style:table-column-properties fo:break-before="auto" style:column-width="22.58mm"/></style:style><style:style style:name="ro1" style:family="table-row"><style:table-row-properties style:row-height="4.52mm" fo:break-before="auto" style:use-optimal-row-height="true"/></style:style><style:style style:name="ta1" style:family="table" style:master-page-name="Default"><style:table-properties table:display="true" style:writing-mode="lr-tb"/></style:style></office:automatic-styles><office:body><office:spreadsheet><table:calculation-settings table:automatic-find-labels="false" table:use-regular-expressions="false" table:use-wildcards="true"/><table:table table:name="Tabelle1" table:style-name="ta1"><table:table-column table:style-name="co1" table:number-columns-repeated="21" table:default-cell-style-name="Default"/><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row><table:table-row table:style-name="ro1"><table:table-cell table:number-columns-repeated="21" office:value-type="string" calcext:value-type="string"><text:p>aabbccddee</text:p></table:table-cell></table:table-row></table:table><table:named-expressions/></office:spreadsheet></office:body></office:document-content>
`

func TestParseCell1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	cell := parseCell([]byte(`
								<table:table-cell office:value-type="string">
									<text:p>BlaBlaBla</text:p>
								</table:table-cell>`))
	if cell != "BlaBlaBla" {
		t.Fail()
	}
}

func TestParseCell2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	cell := parseCell([]byte(`<table:table-cell table:style-name="ce15"/>`))
	if len(cell) != 0 {
		t.Fail()
	}
}

func TestParseCellWithSpaces(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	cell := parseCell([]byte(`
								<table:table-cell office:value-type="string">
									<text:p><text:s text:c="3"/>BlaBlaBla</text:p>
								</table:table-cell>`))
	if cell != "   BlaBlaBla" {
		t.Fail()
	}
}
func TestParseCells1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	cells := parseCells([]byte(`<table:table-cell office:value-type="string">
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
		t.Logf("Excepted 26 cells but got %d\n", len(cells))
		// for i, c := range cells {
		// 	t.Logf("%d. %s\n", i, c)
		// }
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

func TestRead(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	f, err := ParseContent([]byte(contentXML))
	if err != nil {
		t.Fatal(err)
	}

	if len(f.Sheets) != 1 {
		t.Fatal("Wrong number of sheets")
	}

	if len(f.Sheets[0].Rows) != 100 {
		t.Fatal("Wrong number of rows")
	}

	if len(f.Sheets[0].Rows[0]) != 21 {
		t.Fatal("Wrong number of cells")
	}
}

func BenchmarkRead(b *testing.B) {
	// run the Read function b.N times
	for n := 0; n < b.N; n++ {
		ParseContent([]byte(contentXML))
	}
}
