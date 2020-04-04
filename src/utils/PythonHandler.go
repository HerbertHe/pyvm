package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

// 下载Python
func DownloadPython(version string, fileName string, fullName string, fb func(length, downLen int64)) error {
	downUrl := `https://npm.taobao.org/mirrors/python/` + BackVersionNum(version) + `/` + fileName
	var (
		fSize int64
		buf = make([]byte, 32*1024)
		written int64
	)
	client := new(http.Client)
	resp, err := client.Get(downUrl)
	if err != nil {
		return err
	}
	//	读取服务器返回资源大小
	fSize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		return err
	}
	//	创建文件
	file, err := os.Create(fullName)
	if err != nil {
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		return errors.New("body is null")
	}
	defer resp.Body.Close()
	for {
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			// 写入byte
			nw, ew := file.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			// 写入出错
			if ew != nil {
				err = ew
				break
			}
			// 读取长度不等于写入长度
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
			if er != nil {
				if er != io.EOF {
					err = er
				}
				break
			}
			//	没有错误了，使用回调
			// 跳出循环
			fb(fSize, written)
		}
	}
	return err
}

// 安装Python
func InstallPythonPassive(pythonShell string, dir string) {
	fmt.Println("安装python环境中...")
	args := "DefaultJustForMeTargetDir=" + dir + " AssociateFiles=0 Shortcuts=0 Include_launcher=0 InstallLauncherAllUsers=0"
	cmd := exec.Command(pythonShell, args)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("执行安装时发生错误...\n%v\n", err)
	}
}
