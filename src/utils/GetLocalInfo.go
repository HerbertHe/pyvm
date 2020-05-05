package utils

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// 设置windows环境下的软连接
func SetWindowsSymlink (symlink, target string) error {
	cmd := exec.Command("mklink", "/D", symlink, target)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// 遍历文件夹基本方法
func GetLocalDirs(path string) ([]string, error) {
	var dirsList []string
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, fi := range dirs {
		if fi.IsDir() {
			dirsList = append(dirsList, fi.Name())
			return dirsList, nil
		}
	}
	return nil, nil
}

// 判读文件夹是否存在
func IfDirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return false
}