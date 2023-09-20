package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CheckExists(path string) {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return
		}
		fmt.Println("mkdir path:", path)
		err2 := os.Mkdir(path, 0777)
		fmt.Println("mkdir err:", err2)
		return
	}
	return
}

func ZipFolder(source, target string) error {
	// 创建目标压缩文件
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 创建压缩器
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历源文件夹中的所有文件和子文件夹
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 获取相对路径
		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}

		// 创建文件头信息
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 设置文件头中的相对路径
		header.Name = relPath

		// 判断是否为文件夹
		if info.IsDir() {
			header.Name += "/"
		} else {
			// 添加文件到压缩包中
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	fmt.Println("压缩完成:", target)
	return nil
}
