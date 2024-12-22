// -------------------------------------------------
// Package docx_parser
// Author: hanzhi
// Date: 2024/12/22
// -------------------------------------------------

package docx_parser

func stylizedBody(body *Body, styles *Styles) {
	// 根据styles生成一个map方便我查找
	var styleFZMap map[string]int = make(map[string]int)
	for _, style := range styles.StyleList {
		if style.StyleId != "" {
			styleFZMap[style.StyleId] = style.FontSize.Value
		}
	}
	//fmt.Println(styleFZMap)
	// 遍历body，寻找paragraph
	for i, content := range body.Contents {
		if content.Type == "paragraph" {
			paragraph := content.Value.(Paragraph)
			if paragraph.StyleId.Value != "" {
				if fontSize, exists := styleFZMap[paragraph.StyleId.Value]; exists {
					for j := range paragraph.Runs {
						paragraph.Runs[j].FontSize.Value = fontSize
					}
				}
			}
			// 写回到 body.Contents[i]
			body.Contents[i].Value = paragraph
		}
	}

	//fmt.Println(styleFZMap)
}
