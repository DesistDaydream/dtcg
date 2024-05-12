package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	aDir = "D:\\Projects\\dtcg\\images\\cn"
	bDir = "E:\\Game\\DCGO_Application\\Textures\\Card"
)

func main() {
	err := filepath.Walk(aDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是文件则拷贝
		if !info.IsDir() {
			err := copyFile(path, filepath.Join(bDir, info.Name()))
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
