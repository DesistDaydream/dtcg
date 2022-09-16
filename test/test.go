package main

import "strconv"

func main() {
	str := "4000"
	int1, _ := strconv.Atoi(str)
	if int(int1) <= 4000 {
		println(str)
	}
}
