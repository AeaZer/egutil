package main

import (
	"errors"
	"strings"
)

const excelColPattern = "^[A-Z]+$"
const (
	dollarByte  byte = 36
	percentByte byte = 37
	sByte       byte = 115
)

type wds interface {
	absPath() string
	generateType() int
	generator() error
}

type generatorHandler struct {
	forFormatColsIndex []int
	template           string

	wds wds
}

func newWds(wdsType int, path, targetPath, template string, forFormatCols []int, startLine, endLine int) (wds, error) {
	if path == "" || targetPath == "" || template == "" || len(forFormatCols) == 0 || (startLine > endLine && endLine > 0) {
		return nil, errors.New("params input error")
	}
	switch wdsType {
	case typeSql:
		return newSql(path, targetPath, template, forFormatCols, startLine, endLine)
	}
	return nil, errors.New("error wdsType")
}

func validateColsFlag(forFormatColsFlag []string) (int, bool) {
	return 0, true
}

func getColIndex(flag string) int {
	for _, v := range flag {
		return int(v - 65)
	}
	return 0
}

// TODO: 需要支持 A* 列
func (g *generatorHandler) handleTemplate(template string) {
	templateBytes := []byte(template)
	for i := 0; i < len(templateBytes); i++ {
		if templateBytes[i] == dollarByte {
			templateBytes[i] = percentByte
			i++
			g.forFormatColsIndex = append(g.forFormatColsIndex, int(templateBytes[i]-65))
			templateBytes[i] = sByte
		}
	}
	g.template = string(templateBytes)
}

func (g *generatorHandler) startGenerator() error {
	return g.wds.generator()
}

func generatorTargetPath(path string) string {
	pathSplit := strings.SplitN(path, ".", 2)
	if len(pathSplit) != 2 {
		panic("unexpected path")
	}
	return pathSplit[0] + "_generator.sql"
}

func newSimpleSqlGeneratorHandler(path, template string, config *generatorSQLConfig) (*generatorHandler, error) {
	if config == nil {
		config = &generatorSQLConfig{}
	}
	g := new(generatorHandler)
	g.handleTemplate(template)
	w, err := newWds(typeSql, path, generatorTargetPath(path), g.template, g.forFormatColsIndex, config.startLine, config.endLine)
	if err != nil {
		return nil, err
	}
	g.wds = w
	return g, nil
}
