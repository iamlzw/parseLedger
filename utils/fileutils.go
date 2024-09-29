package utils

import (
	"fmt"
	"os"
)

func IsDirExists(path string) bool {
	// 获取文件或文件夹的信息
	info, err := os.Stat(path)

	// 如果发生错误，并且错误是因为文件或目录不存在
	if os.IsNotExist(err) {
		return false
	}

	// 如果没有发生错误，判断它是否是一个目录
	return info.IsDir()
}

func CreateDirIfNotExists(path string) error {
	if !IsDirExists(path) {
		// 如果目录不存在，则创建该目录
		err := os.MkdirAll(path, os.ModePerm) // os.ModePerm 表示 0777 权限
		if err != nil {
			return err
		}
		fmt.Printf("Directory %s created successfully\n", path)
	}
	return nil
}
