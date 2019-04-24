package main

import (
	"fmt"
)

func join(delimiter string, vals ...string) string {
	tmp := ""
	for _, val := range vals {
		if tmp != "" {
			tmp += delimiter
		}
		tmp += val
	}
	return tmp
}

func main() {
	fmt.Println(join(",", []string{}...))
	fmt.Println(join(",", []string{"あああ"}...))
	fmt.Println(join(",", []string{"あいうえお", "かきくけこ", "さしすせそ"}...))
	fmt.Println(join("__", []string{"あいうえお", "かきくけこ", "さしすせそ"}...))
}
