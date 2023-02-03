package main

const (
	TypeSQL = iota + 1
	TypePD
)

const excelColPattern = "^[A-Z]+$"

const (
	dollarByte  byte = 36  // '$'
	percentByte byte = 37  // '%'
	sByte       byte = 115 // 's'
)
