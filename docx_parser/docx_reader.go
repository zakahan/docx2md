package docx_parser

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 修改后的读取和解析函数，将段落和表格混合存储
func ReadDocx(filePath string, outputFileDir string) (*Document, error) {
	r, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}
	defer func(r *zip.ReadCloser) {
		err := r.Close()
		if err != nil {
			// empty
		}
	}(r)

	// 查找 document.xml.rels文件，也就是多媒体依赖
	var documentFileRels *zip.File
	for _, f := range r.File {
		if f.Name == "word/_rels/document.xml.rels" {
			documentFileRels = f
			break
		}
	}
	if documentFileRels == nil {
		return nil, fmt.Errorf("document.xml not found in %s", filePath)
	}
	// 读取 document.xml 的内容
	rcDFR, err := documentFileRels.Open()
	if err != nil {
		return nil, err
	}
	defer func(rc io.ReadCloser) {
		err := rc.Close()
		if err != nil {
			// empty
		}
	}(rcDFR)
	// 解析相关依赖 得到一个rid和images的map
	var relationships Relationships
	err = xml.NewDecoder(rcDFR).Decode(&relationships)
	if err != nil {
		return nil, err
	}

	/* ----------------------------------------------------------------------------- */

	// 查找 document.xml
	var documentFile *zip.File
	for _, f := range r.File {
		if f.Name == "word/document.xml" {
			documentFile = f
			break
		}
	}

	if documentFile == nil {
		return nil, fmt.Errorf("document.xml not found in %s", filePath)
	}

	// 创建保存图片的文件夹
	outputDir := filepath.Join(outputFileDir, "images")
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	// 读取 document.xml 的内容
	rcDF, err := documentFile.Open()
	if err != nil {
		return nil, err
	}
	defer func(rc io.ReadCloser) {
		err := rc.Close()
		if err != nil {
			// empty
		}
	}(rcDF)

	decoder := xml.NewDecoder(rcDF)
	var body Body
	var contentItem ContentItem

	// 解析 XML 并处理段落、表格和图片
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "p" { // 段落
				var para Paragraph
				err := decoder.DecodeElement(&para, &t)
				if err != nil {
					return nil, err
				}

				// 处理段落中的图片
				for _, run := range para.Runs {
					if run.Drawing != nil {
						imagePath, err := extractImageFromDocx(r, run.Drawing.Blip.Embed, relationships, outputDir)
						if err != nil {
							return nil, err
						}
						contentItem = ContentItem{Type: "image", Value: imagePath}
						body.Contents = append(body.Contents, contentItem)
					}
				}

				contentItem = ContentItem{Type: "paragraph", Value: para}
				body.Contents = append(body.Contents, contentItem)
			} else if t.Name.Local == "tbl" { // 表格
				var tbl Table
				err := decoder.DecodeElement(&tbl, &t)
				if err != nil {
					return nil, err
				}
				contentItem = ContentItem{Type: "table", Value: tbl}
				body.Contents = append(body.Contents, contentItem)
			}
		}
	}

	return &Document{Body: body}, nil
}
