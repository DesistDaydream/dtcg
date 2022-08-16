package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "a源b你你c我ddd"
	// 正则匹配 a 中不止中文的字符
	match := "[!^\u4e00-\u9fa5]"
	reg := regexp.MustCompile(match)
	result := reg.ReplaceAllString(a, "")
	fmt.Println(result)
}
