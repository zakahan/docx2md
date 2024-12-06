package docx_parser

import (
	"bytes"
	"github.com/mattn/go-runewidth"
	"strings"
)

func Table2markdown(table Table) string {
	// 对齐表格宽度
	var tableWidth = 0
	for _, row := range table.Rows {
		if tableWidth < len(row.Cells) {
			tableWidth = len(row.Cells)
		}
	}
	// 对齐单元格宽度
	alignArr := make([]int, tableWidth)
	for _, row := range table.Rows {
		for j, cell := range row.Cells {
			if alignArr[j] < vLen(strings.Join(cell.Texts, "")) {
				alignArr[j] = vLen(strings.Join(cell.Texts, ""))
			}
		}
	}

	var buffer bytes.Buffer
	// 默认第一行是表头
	for i, row := range table.Rows {
		if i == 0 {
			widArr := make([]string, tableWidth) // 为了追加个 |--|表格标记
			// 遍历单元格
			for j, cell := range row.Cells {
				cellStr := strings.Join(cell.Texts, "")
				var text = "| " + cellStr // 添加
				buffer.WriteString(text)
				widDif := alignArr[j] - vLen(cellStr)
				if widDif > 0 {
					buffer.WriteString(strings.Repeat(" ", widDif)) // 加多少个-
				} else {
					buffer.WriteString(" ")
				}

				widArr[j] = strings.Repeat("-", alignArr[j])
			}
			// 补足宽度
			for j := len(row.Cells); j < tableWidth; j++ {
				buffer.WriteString("| ")
				buffer.WriteString(strings.Repeat(" ", alignArr[j]))
				widArr[j] = strings.Repeat("-", alignArr[j])
			}
			buffer.WriteString("|\n")
			// -----------------------------
			// 在第一行后面添加的。
			for j := range widArr {
				var text = "| " + widArr[j] + " "
				buffer.WriteString(text)
			}
			// ------------------------------
			// 总之都要添加这个-----------------------------
			buffer.WriteString("|\n")
		} else { //	 非首行
			// 遍历单元格
			for j, cell := range row.Cells {
				cellStr := strings.Join(cell.Texts, "")
				var text = "| " + cellStr // 添加
				buffer.WriteString(text)
				widDif := alignArr[j] - vLen(cellStr)
				if widDif > 0 {
					buffer.WriteString(strings.Repeat(" ", widDif)) // 加多少个-
				} else {
					buffer.WriteString(" ")
				}
			}
			// 补足宽度
			for j := len(row.Cells); j < tableWidth; j++ {
				buffer.WriteString("| ")
				buffer.WriteString(strings.Repeat(" ", alignArr[j]))
			}
			buffer.WriteString("|\n")
		}
	}
	result := buffer.String()
	return result
}

func vLen(s string) int {
	// visual length
	return runewidth.StringWidth(s)
}
