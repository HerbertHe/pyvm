package utils

import (
	"fmt"
	"os/exec"
)

func InstallPythonPassive(pythonShell string, dir string) {
	fmt.Println("安装python环境中...")
	args := "DefaultJustForMeTargetDir=" + dir + " AssociateFiles=0 Shortcuts=0 Include_launcher=0 InstallLauncherAllUsers=0"
	cmd := exec.Command(pythonShell, args)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("执行安装时发生错误...\n%v\n", err)
	}
}
