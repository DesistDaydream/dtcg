package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	cnDir         = "D:\\Projects\\dtcg\\images\\cn"
	gameCardDir   = "E:\\Game\\DCGO_Standalone\\Textures\\Card"
	gameENCardDir = "E:\\Game\\DCGO_Standalone\\Textures\\Card.en"
)

func copyCNCardPictureToCardDir() {
	err := filepath.Walk(cnDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是文件则拷贝
		if !info.IsDir() {
			err := copyFile(path, filepath.Join(gameCardDir, info.Name()))
			if err != nil {
				fmt.Printf("无法拷贝文件 %s: %v\n", path, err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
	}
}

func copyNoexistPicture() {
	nameList := "BT13-097_P1, BT12-001_P1, BT12-007_P1, BT12-010_P1, BT12-011_P1, BT12-011_P2, BT12-011_P3, BT12-016_P1, BT12-018_P2, BT10-026_P1, BT10-055, BT10-082_P1, BT9-058_P1, BT9-068_P1, BT8-008_P3, BT8-013_P3, BT8-032, BT8-042_P1, BT8-042_P2, BT8-042_P3, BT8-053_P3, BT8-112_P4, BT7-021_P1, BT7-021_P2, BT7-023_P1, BT7-030_P1, BT7-035_P1, BT7-035_P2, BT7-036_P1, BT7-036_P2, BT7-042_P1, BT7-004_P1, BT7-044_P2, BT7-046_P1, BT7-046_P2, BT7-047_P1, BT7-054_P1, BT7-062_P2, BT7-062_P3, BT7-071_P1, BT7-071_P2, BT7-073_P1, BT7-078_P1, BT7-080_P3, BT7-081_P1, BT7-081_P2, BT7-081_P3, BT7-081_P4, BT7-110_P1, BT7-110_P2, BT7-110_P3, BT7-110_P4, BT6-001_P2, BT6-009_P3, BT6-011_P1, BT6-016_P2, BT6-017_P1, BT6-021_P1, BT6-038_P1, BT6-038_P2, BT6-038_P3, BT6-048_P1, BT6-048_P2, BT6-048_P3, BT6-005_P2, BT6-056_P1, BT6-064_P2, BT6-065_P1, BT6-077_P1, BT6-083_P1, BT6-085_P1, BT5-007_P3, BT5-007_P4, BT5-010_P2, BT5-010_P3, BT5-011_P1, BT5-012_P2, BT5-020_P3, BT5-024_P1, BT5-026_P2, BT5-032_P1, BT5-032_P2, BT5-003_P1, BT5-034_P1, BT5-035_P1, BT5-042_P2, BT5-042_P3, BT5-045_P2, BT5-062_P1, BT5-071_P1, BT5-071_P2, BT5-071_P3, BT5-083_P1, BT5-109_P1, BT4-011_P3, BT4-011_P4, BT4-011_P5, BT4-030_P1, BT4-038_P1, BT4-048_P2, BT4-005_P2, BT4-005_P3, BT4-063_P1, BT4-072_P1, BT4-074_P1, BT4-075_P1, BT4-006_P1, BT4-078_P1, BT4-084_P1, BT4-084_P2, BT4-084_P3, BT4-087_P1, BT4-088_P3, BT4-090_P3, BT4-090_P4, BT3-002_P2, BT3-024_P2, BT3-031_P2, BT3-003_P1, BT3-003_P2, BT3-054_P1, BT3-103_P1, BT3-111_P2, BT3-075_P2, BT3-006_P2, BT3-077_P1, BT3-091_P2, BT3-091_P3, BT3-091_P4, BT3-112_P2, BT2-091, BT2-092, BT2-024_P1, BT2-028_P2, BT2-029, BT2-035_P1, BT2-035_P2, BT2-035_P3, BT2-004_P2, BT2-004_P3, BT2-044_P1, BT2-046_P1, BT2-051_P1, BT2-051_P2, BT2-055, BT2-055_P1, BT2-055_P2, BT2-056, BT2-056_P2, BT2-060, BT2-065_P1, BT2-065_P2, BT2-066_P3, BT2-066_P4, BT1-010_P3, BT1-025_P3, BT1-038_P3, BT1-115_P1, BT1-060_P2, BT1-063_P1, BT1-107_P2, BT1-077_P2, BT1-108_P2, BT1-110_P4, BT1-084_P2, BT1-084_P4, BT1-084_P5, EX4-060_P2, EX3-010_P1, EX3-020_P1, EX3-041_P1, EX3-049_P1, EX3-061_P1, EX2-009_P1, EX2-009_P2, EX2-009_P3, EX2-056_P1, EX2-073_P2, EX2-045_P2, EX1-007_P1, EX1-007_P2, EX1-007_P3, EX1-008_P2, EX1-024_P2, EX1-048_P1, EX1-048_P2, EX1-048_P3, EX1-049_P1, ST9-04_P2, ST9-04_P3, ST9-05_P2, ST9-05_P3, ST9-05_P4, ST8-03_P1, ST8-04_P2, ST8-04_P3, ST8-04_P4, ST8-06_P3, ST7-03_P4, ST7-06_P1, ST7-06_P2, ST7-06_P3, ST7-06_P4, ST5-12_P1, ST4-01, ST4-08_P3, ST4-08_P4, ST4-08_P5, ST4-13_P1, ST3-05_P2, ST3-07_P2, ST3-11, ST3-13_P2, ST3-14_P2, ST2-02_P1, ST2-11_P2, ST2-13_P3, ST2-16_P2, ST1-03_P3, ST1-03_P5, ST1-03_P6, ST1-13_P2, ST1-16_P3, ST1-16_P4, P-001_P2, P-002_P1, P-009_P1, P-009_P2, P-010_P1, P-010_P2, P-010_P3, P-029_P1, P-041_P2, P-059, P-062, P-079_P1, P-088, P-003_P1, P-004_P2, P-004_P3, P-008_P3, P-008_P4, P-030_P1, P-061, P-064, P-089, P-005_P2, P-006_P1, P-028_P1, P-028_P2, P-043_P1, P-081_P1, P-087, P-032_P1, P-060, P-063, P-090, P-013_P1, P-014_P1, P-016_P2, P-016_P3, P-016_P4, P-033_P2, P-078, P-018_P1, P-046_P1, P-077_P1, P-080_P1"
	names := strings.Split(nameList, ", ")
	fmt.Println(len(names))
	for _, name := range names {
		name = fmt.Sprintf("%s.png", name)
		// fmt.Println(filepath.Join(gameENCardDir, name), filepath.Join(gameCardDir, name))
		copyFile(filepath.Join(gameENCardDir, name), filepath.Join(gameCardDir, name))
	}
}

// 拷贝文件
func copyFile(src, dst string) error {
	fmt.Println(src, dst)
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

func main() {
	// copyCNCardPictureToCardDir()
	copyNoexistPicture()
}
