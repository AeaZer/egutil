package main

import "fmt"

func main() {
	path := "/Users/tommorrow/Downloads/target.xlsx"
	template := "UPDATE m_sound_comment SET `pool` = 40 WHERE `id` = $C AND `sound_id` = $A;"
	handler, err := NewSimpleSQLGenerateHandler(path, template, &GenerateSQLConfig{
		startLine: 1,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = handler.StartGenerate()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
