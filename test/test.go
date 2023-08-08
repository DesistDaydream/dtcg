package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "\u4f60\\u597d"                     // Unicode 编码的 "你好"
	t, err := strconv.Unquote(`"` + s + `"`) // 加上双引号，然后解码
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t) // 输出 "你好"
	}
}
