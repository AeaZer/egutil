package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
)

type PDGenerate struct {
	path               string
	folderPath         string
	startLine, endLine int
}

func NewPDGenerate(path, targetPath string, startLine, endLine int) *PDGenerate {
	return &PDGenerate{
		path:       path,
		folderPath: targetPath,
		startLine:  startLine,
		endLine:    endLine,
	}
}

func getPictureName(pictureURI string) (string, error) {
	uris := strings.SplitN(pictureURI, ".", 2)
	if len(uris) != 2 {
		return "", fmt.Errorf("无法时别的 picture uri: %s", pictureURI)
	}
	return uris[1], nil
}

func (p *PDGenerate) absPath() string {
	return p.path
}

func (p *PDGenerate) generateType() int {
	return TypePD
}

func (p *PDGenerate) generate() error {
	excel, err := excelize.OpenFile(p.path)
	if err != nil {
		fmt.Println("【源文件】打开失败", err)
		return err
	}
	rows := excel.GetRows("Sheet1")
	totalLength := len(rows)
	// forbidden panic of out index
	if totalLength < p.endLine {
		p.endLine = totalLength
	}
	if p.endLine == 0 {
		p.endLine = totalLength
	}

	return nil
}
