package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "起始牌組特製新手套裝[ST-11]"
	re := regexp.MustCompile(`\[(.*?)\]`)
	match := re.FindStringSubmatch(str)
	if len(match) > 1 {
		fmt.Println(match[1]) // ST-11
	}
}
