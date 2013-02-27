package lexer

import (
	"bytes"
	"fmt"
)

type Word string

type Words []Word

func GetWords(b []byte) (w Words) {
	fields := bytes.FieldsFunc(b, func(r rune) bool {
		return false
	})
	fmt.Printf("%v\n", fields)
	return
}
