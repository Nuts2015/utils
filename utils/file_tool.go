package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
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
		err2 := os.MkdirAll(path, 0777)
		fmt.Println("mkdir err:", err2)
		return
	}
	return
}

// IsDirEmpty 检查给定路径的文件夹是否为空
func IsDirEmpty(dirName string) (bool, error) {
	f, err := os.Open(dirName)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // 尝试读取第一个文件名
	if err == io.EOF {
		return true, nil
	}
	return false, err
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

func UnzipFolder(source, target string) error {
	// 打开要解压的压缩文件
	zipFile, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 遍历压缩文件中的所有文件和文件夹
	for _, file := range zipFile.File {
		// 获取解压后的文件路径
		filePath := filepath.Join(target, file.Name)

		// 如果当前项为文件夹，则创建对应的文件夹
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		// 创建解压后的文件
		err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return err
		}

		outputFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer outputFile.Close()

		// 打开压缩文件中的文件
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// 将压缩文件中的内容复制到解压后的文件中
		_, err = io.Copy(outputFile, rc)
		if err != nil {
			return err
		}
	}

	fmt.Println("解压完成:", source)
	return nil
}

func FilterDir(path, name string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
