# egutil
generator sql with excel

## quick start
```golang
    path := "/Users/tommorrow/Downloads/target.xlsx"
	template := "UPDATE m_sound_comment SET `pool` = 40 WHERE `id` = $C AND `sound_id` = $A;"
	handler, err := newSimpleSQLGenerateHandler(path, template, &generateSQLConfig{
		startLine: 1,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = handler.startGenerator()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
```