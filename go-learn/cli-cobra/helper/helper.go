package helper

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(i, j)
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func UpperCase(s string) string {
	return strings.ToUpper(s)
}

func Modify(s string, opt bool) string {
	if opt {
		return s + "_MODIFIED"
	} else {
		return s + "_modified"
	}
}
