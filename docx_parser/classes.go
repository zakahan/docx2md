package docx_parser

import "encoding/xml"

type ContentItem struct {
	Type  string      // "paragraph", "table", "image"
	Value interface{} // Paragraph, Table 或 图片路径 (string)
}

type Document struct {
	Body Body `xml:"body"`
}

type Body struct {
	Contents []ContentItem
}

type Paragraph struct { // 段落类型 w:p
	Runs    []Run  `xml:"r"`          // 段落包含多个 run
	NumPr   *bool  `xml:"pPr>numPr"`  // 检查是否存在编号信息
	StyleId PStyle `xml:"pPr>pStyle"` // 段落样式
}

type PStyle struct { // 段落样式
	Value string `xml:"val,attr"`
}

type Run struct { // 文本运行，可能包含文本或图片
	FontSize FontSize `xml:"rPr>sz"`
	//FontBold    *bool    `xml:"rPr>b"`        // 会出现连续**问题，故去掉
	//FontIncline *bool    `xml:"rPr>i"`
	Text    []Text   `xml:"t"`
	Drawing *Drawing `xml:"drawing,omitempty"` // 可能包含图片
}

type Drawing struct { // 图片嵌套在 <w:drawing> 中
	Blip Blip `xml:"inline>graphic>graphicData>pic>blipFill>blip"`
}

type Blip struct {
	Embed string `xml:"embed,attr"`
}

type Table struct {
	XMLName xml.Name `xml:"tbl"`
	Rows    []Row    `xml:"tr"`
}

type Row struct {
	Cells []Cell `xml:"tc"`
}

type Cell struct {
	Texts []string `xml:"p>r>t"`
}

type Text struct {
	Value string `xml:",chardata"`
}

// FontSize -------------------------------
type FontSize struct {
	Value int `xml:"val,attr"`
}

// Relationships -------------------------------------
// 定义结构体以匹配XML格式
type Relationships struct {
	XMLName      xml.Name       `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationships"`
	Relationship []Relationship `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationship"`
}

type Relationship struct {
	Id     string `xml:"Id,attr"`
	Type   string `xml:"Type,attr"`
	Target string `xml:"Target,attr"`
}

/* -------------------------------------------------------------- */

// Styles 样式表
type Styles struct {
	XMLName   xml.Name
	StyleList []Style `xml:"style"`
}

type Style struct {
	Name     Name     `xml:"name"`
	StyleId  string   `xml:"styleId,attr"`
	FontSize FontSize `xml:"rPr>sz"`
}

type Name struct {
	Value string `xml:"val,attr"`
}
