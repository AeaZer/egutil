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
- 生成的文件为 path 同目录下名称为同名_generator.sql 的文件
- template 不一定是 sql（不校验 template 是否是合法的 sql ），可以时任意的字符串
- generateSQLConfig 可以为 nil 表示整个文件所有行都会生成一条 sql

### 不支持
- 不支持 template 中的 ${excel 列标识} `len(列标识) != 1`