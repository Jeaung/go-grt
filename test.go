package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abc"
	s := strings.Split(str, "c")
	fmt.Println(s[0], s[1])
}
