package main

import (
	"errors"
	"strings"
)

type wds interface {
	absPath() string
	generateType() int
	generate() error
}

type GenerateHandler struct {
	forFormatColsIndex []int  // for sql generate
	template           string // for sql generate

	wds wds // interface for all
}

func NewWds(wdsType int, path, targetPath, template string, forFormatCols []int, startLine, endLine int) (wds, error) {
	if path == "" || targetPath == "" || template == "" || len(forFormatCols) == 0 || (startLine > endLine && endLine > 0) {
		return nil, errors.New("params input error")
	}
	switch wdsType {
	case TypeSQL:
		return newSQLGenerate(path, targetPath, template, forFormatCols, startLine, endLine)
	}
	return nil, errors.New("error wdsType")
}

// TODO: 指定列标识规则待校验
func validateColsFlag(forFormatColsFlag []string) (int, bool) {
	return 0, true
}

// TODO: 需要支持 A* 列
func (g *GenerateHandler) handleTemplate(template string) {
	templateBytes := []byte(template)
	for i := 0; i < len(templateBytes); i++ {
		if templateBytes[i] == dollarByte {
			templateBytes[i] = percentByte
			i++
			if i == len(templateBytes) {
				panic("error template, $ must have char followed")
			}
			g.forFormatColsIndex = append(g.forFormatColsIndex, int(templateBytes[i]-65))
			templateBytes[i] = sByte
		}
	}
	g.template = string(templateBytes)
}

// StartGenerate is a shortcut for g.wds.generate()
func (g *GenerateHandler) StartGenerate() error {
	return g.wds.generate()
}

func generatorTargetPath(path string) string {
	pathSplit := strings.SplitN(path, ".", 2)
	if len(pathSplit) != 2 {
		panic("unexpected path")
	}
	return pathSplit[0] + "_generate.sql"
}

func NewSimpleSQLGenerateHandler(path, template string, config *GenerateSQLConfig) (*GenerateHandler, error) {
	if config == nil {
		config = &GenerateSQLConfig{}
	}
	g := new(GenerateHandler)
	g.handleTemplate(template)
	w, err := NewWds(TypeSQL, path, generatorTargetPath(path), g.template, g.forFormatColsIndex, config.startLine, config.endLine)
	if err != nil {
		return nil, err
	}
	g.wds = w
	return g, nil
}
