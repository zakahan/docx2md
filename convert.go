package docx2md

import (
	"bytes"
	"fmt"
	"github.com/zakahan/docx2md/docx_parser"
	"path/filepath"
	"strings"
	"unicode"
)

func DocxConvert(docPath string, outputDir string) (string, string, error) { // 分别是 文件名， 内容字符串， error
	mdPath, mdDirPath, err := docx_parser.CreateMdDir(docPath, outputDir, ".docx")
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}
	// --------------
	doc, err := docx_parser.ReadDocx(docPath, mdDirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	var buffer bytes.Buffer

	for _, content := range doc.Body.Contents {
		if content.Type == "paragraph" {
			var bufferPar bytes.Buffer // 对一个段落
			var fontSizePar int = 48   // 找个比较大的值，应该不会比这个更离谱，反正都是h1
			para := content.Value.(docx_parser.Paragraph)
			numPr := para.NumPr
			for _, run := range para.Runs {
				// 统计最小字号是多少，按最小的来
				if run.FontSize.Value < fontSizePar {
					fontSizePar = run.FontSize.Value
				}
				for _, text := range run.Text {
					bufferPar.WriteString(text.Value)
				}
			}
			// runs结束了在一个字符串内
			paragraphStr := bufferPar.String()
			paragraphStr = word2Heading(paragraphStr, fontSizePar, numPr)
			buffer.WriteString(paragraphStr)
		} else if content.Type == "table" {
			table := content.Value.(docx_parser.Table)
			tableStr := docx_parser.Table2markdown(table)
			buffer.WriteString(tableStr)

		} else if content.Type == "image" {
			imagePath := content.Value.(string)
			imageName := filepath.Base(imagePath)
			// 图片选择添加相对路径
			buffer.WriteString("![" + imageName + "](" + filepath.Join("images", imageName) + ")")
		}
		buffer.WriteString("\n")

	}
	markdownStr := buffer.String()
	err = docx_parser.SaveFile(mdPath, markdownStr)
	return mdPath, markdownStr, err

}

func word2Heading(value string, fontSize int, numPr *bool) string {
	value = getTrimedStr(value)
	if len(value) == 0 {
		return "\n"
	}
	if numPr != nil || docx_parser.CheckString(value) {
		var maxHeadingLength = 45
		var h1 = 48
		var h2 = 36
		var h3 = 28
		var h4 = 24

		if h1 <= fontSize {
			return "# " + value
		} else if h2 <= fontSize {
			return "## " + value
		} else if h3 <= fontSize {
			return "### " + value
		} else if h4 <= fontSize && len(value) < maxHeadingLength {
			return "#### " + value
		} else {
			return value
		}

	} else {
		var maxHeadingLength = 15
		var h1 = 48
		var h2 = 36
		var h3 = 28

		if h1 <= fontSize {
			return "# " + value
		} else if h2 <= fontSize {
			return "## " + value
		} else if h3 <= fontSize && len(value) < maxHeadingLength {
			return "### " + value
		} else {
			return value
		}

	}

}

func getTrimedStr(s string) string {
	// 使用 TrimFunc 去掉字符串两端的所有空白字符
	trimmed := strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	return trimmed
}
