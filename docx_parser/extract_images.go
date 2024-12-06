package docx_parser

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 提取并保存图片到指定文件夹
// 提取并保存图片到指定文件夹
func extractImageFromDocx(r *zip.ReadCloser, rid string, relationships Relationships, outputDir string) (string, error) {
	// 图片资源ID通常以 "rId" 开头
	var mediaFileName string
	for _, rel := range relationships.Relationship {
		if rid == rel.Id {
			mediaFileName = rel.Target
			break
		}
	}

	for _, f := range r.File {
		// 查找与资源ID对应的图片文件
		if strings.Contains(f.Name, mediaFileName) && strings.HasPrefix(f.Name, "word/media/") {
			rc, err := f.Open()
			if err != nil {
				return "", err
			}
			defer func(rc io.ReadCloser) {
				err := rc.Close()
				if err != nil {
					// empty
				}
			}(rc)

			// 构造图片的输出路径
			imagePath := filepath.Join(outputDir, filepath.Base(f.Name))

			// 将图片数据保存到文件
			imageData, err := io.ReadAll(rc)
			if err != nil {
				return "", err
			}
			err = os.WriteFile(imagePath, imageData, os.ModePerm)
			if err != nil {
				return "", err
			}

			// 返回图片的路径
			return imagePath, nil
		}
	}
	return "", fmt.Errorf("no image found for resource ID: %s", rid)
}
