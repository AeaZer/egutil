package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type GenerateSQLConfig struct {
	startLine, endLine int
}

type sqlHandler struct {
	path               string
	targetPath         string
	template           string
	forFormatCols      []int
	startLine, endLine int
}

func (s *sqlHandler) absPath() string {
	return s.path
}

func (s *sqlHandler) generateType() int {
	return TypeSql
}

func newSQLHandler(path, targetPath, template string, forFormatCol []int, startLine, endLine int) (*sqlHandler, error) {
	return &sqlHandler{
		path:          path,
		targetPath:    targetPath,
		template:      template,
		forFormatCols: forFormatCol,
		startLine:     startLine,
		endLine:       endLine,
	}, nil
}

func (s *sqlHandler) generate() error {
	excel, err := excelize.OpenFile(s.path)
	if err != nil {
		fmt.Println("【源文件】打开失败", err)
		return err
	}
	rows := excel.GetRows("Sheet1")
	totalLength := len(rows)
	// forbidden panic of out index
	if totalLength < s.endLine {
		s.endLine = totalLength
	}
	if s.endLine == 0 {
		s.endLine = totalLength
	}
	generatorSqls := make([]string, 0, s.endLine-s.startLine+1)
	for row := s.startLine - 1; row < s.endLine; row++ {
		values := make([]interface{}, 0, len(s.forFormatCols))
		for _, col := range s.forFormatCols {
			values = append(values, rows[row][col])
		}
		generatorSqls = append(generatorSqls, fmt.Sprintf(s.template, values...))
	}

	file, err := os.OpenFile(s.targetPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("【写入文件】打开失败", err)
		return err
	}
	// 及时关闭 file 句柄
	defer file.Close()
	// 写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for _, generatorSql := range generatorSqls {
		err = writeRowString(write, generatorSql)
		if err != nil {
			fmt.Println("【写入】时发生错误", err)
			return err
		}
	}
	fmt.Println("row count", len(generatorSqls))
	// Flush 将缓存的文件真正写入到文件中
	write.Flush()
	return nil
}

func writeRowString(writer *bufio.Writer, writeTarget string) error {
	_, err := writer.WriteString(writeTarget + "\n")
	if err != nil {
		return err
	}
	return nil
}
