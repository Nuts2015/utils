package utils

import (
	"os"
)

func CheckExists(path string) {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return
		}
		os.Mkdir(path, 0777)
		return
	}
	return
}

func MakeDir(path string) {

}
