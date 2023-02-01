package main

import "fmt"

func main() {
	handler, err := newGeneratorHandler(
		typeSql,
		"C:\\Users\\34749\\Downloads\\target.xlsx",
		"C:\\Users\\34749\\Downloads\\updateSql.txt",
		"UPDATE m_sound_comment SET police = 1 where id = %s AND dm_id = %s",
		[]string{"A", "C"},
		2,
		5,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = handler.wds.generator()
	if err != nil {
		fmt.Println(err)
		return
	}
}
