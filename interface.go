package main

import (
	"errors"
	"fmt"
)

const excelColRegexp = "^[A-Z]+$"

type wds interface {
	absPath() string
	generateType() int
	generator() error
}

type generatorHandler struct {
	forFormatColsFlag []string // need trans
	forFormatCols     []int
	wds               wds
}

func newWds(wdsType int, path, targetPath, template string, forFormatCols []int, startLine, endLine int) (wds, error) {
	if path == "" || targetPath == "" || template == "" || len(forFormatCols) == 0 || startLine > endLine {
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

func (g *generatorHandler) transFlagToIndex() error {
	index, ok := validateColsFlag(g.forFormatColsFlag)
	if !ok {
		return fmt.Errorf("error col flag %s", g.forFormatColsFlag[index])
	}
	for _, colFlag := range g.forFormatColsFlag {
		g.forFormatCols = append(g.forFormatCols, getColIndex(colFlag))
	}
	return nil
}

func newGeneratorHandler(wdsType int, path, targetPath,
	template string, forFormatCols []string, startLine, endLine int) (*generatorHandler, error) {
	g := &generatorHandler{
		forFormatColsFlag: forFormatCols,
	}
	err := g.transFlagToIndex()
	if err != nil {
		return nil, err
	}
	w, err := newWds(wdsType, path, targetPath, template, g.forFormatCols, startLine, endLine)
	if err != nil {
		return nil, err
	}
	g.wds = w
	return g, nil
}

func newSimpleSqlGeneratorHandler(path, template string, forFormatCols []string, config *generatorSQLConfig) {
	if config == nil {

	}
}
