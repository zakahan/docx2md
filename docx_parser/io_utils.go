package docx_parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateMdDir(documentPath string, outputDir string, suffix string) (string, string, error) {
	// 获取文件名称 并换后缀
	docxName := filepath.Base(documentPath)
	mdName := strings.TrimSuffix(docxName, suffix)
	mdName = mdName + ".md"

	// ---------------------------------
	// 检查outputDir是否存在，如果不存在就报错
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		fmt.Printf("path does not exist: %s\n", outputDir)
		return "", "", err
	}
	// --------------------------------
	// 创建uuid路径
	//uuidStr := uuid.New().String()
	//mdDirPath := filepath.Join(outputDir, uuidStr)
	mdDirPath := outputDir
	mdPath := filepath.Join(mdDirPath, mdName)
	// 检查路径是否存在，如果不存在则创建
	err := os.MkdirAll(mdDirPath, 0755)
	if err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return "", "", err
	}
	return mdPath, mdDirPath, err
}

func SaveFile(filePath string, mdStr string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			// 啥也不干
		}
	}(file)
	// 写入内容
	_, err = file.Write([]byte(mdStr))
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
		return err
	}
	//
	return err
}
