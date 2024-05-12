package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	filenames := []string{
		"ST17-08.png",
		"ch_BT15-003_01.jpg",
		"ch_BT15-110_02_D的副本.jpeg",
		"ch_P-035_D的副本.png",
		"BT13-001_02.png",
		"1708412303430BT15-001.png",
	}

	for _, filename := range filenames {
		handleFileName(filename)
	}
}

func handleFileName(fileName string) {
	// 去除中文及其他非ASCII字符
	removeNonASCIIRe := regexp.MustCompile(`[\p{Han}]+`)
	// 匹配至少10位的连续数字，作为时间戳；以及开头的 ch_，这两者都要去掉
	timeStampRe := regexp.MustCompile(`^\d{10,}|^ch_`)
	// 解析卡包名称、编号和数字序号
	re := regexp.MustCompile(`([A-Za-z]+-[\d]+)_\D*(\d+)\D*\.\w+`)

	// 去除时间戳
	noTimeStamp := timeStampRe.ReplaceAllString(fileName, "")
	// 去除中文及其他非ASCII字符
	noHanzi := removeNonASCIIRe.ReplaceAllString(noTimeStamp, "")
	// 将 BT15-001_01.png,BT15-001_02.png,BT13-001_02.png,
	// 等等这类字符串中的 _01 替换为 _P1, _02 替换为 _P2, 以此类推，还有 3, 4, 5, etc.
	newFilename := re.ReplaceAllStringFunc(noHanzi, func(m string) string {
		parts := re.FindStringSubmatch(m)
		num, _ := strconv.Atoi(parts[2])                // 转换序号部分
		return fmt.Sprintf("%s_P%d.png", parts[1], num) // 生成标准格式的字符串
	})

	fmt.Println(newFilename)
}
