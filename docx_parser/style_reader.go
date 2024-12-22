// -------------------------------------------------
// Package docx_parser
// Author: hanzhi
// Date: 2024/12/22
// -------------------------------------------------

package docx_parser

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
)

func ReadStyle(r *zip.ReadCloser, filePath string) (*Styles, error) {
	// 查找 document.xml.rels文件，也就是多媒体依赖
	var styleFileRels *zip.File
	for _, f := range r.File {
		if f.Name == "word/styles.xml" {
			styleFileRels = f
			break
		}
	}

	if styleFileRels == nil {
		return nil, fmt.Errorf("styles.xml not found in %s", filePath)
	}
	// 读取style.xml的内容
	rcDFR, err := styleFileRels.Open()
	if err != nil {
		return nil, err
	}
	defer func(rc io.ReadCloser) {
		err := rc.Close()
		if err != nil {
			// empty
		}
	}(rcDFR)
	// 解析
	var stylesList Styles
	err = xml.NewDecoder(rcDFR).Decode(&stylesList)
	if err != nil {
		return nil, err
	}

	return &stylesList, nil
}
