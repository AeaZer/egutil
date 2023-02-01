package main

import "fmt"

func main() {
	// TODO: 待简化
	handler, err := newGeneratorHandler(
		typeSql,
		"/Users/tommorrow/Downloads/target.xlsx",
		"/Users/tommorrow/Downloads/updateSql.txt",
		"UPDATE m_sound_comment SET `pool` = 40 WHERE `id` = %s AND `sound_id` = %s;",
		[]string{"C", "A"},
		2,
		0,
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
