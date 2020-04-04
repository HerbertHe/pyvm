/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

const yamlDefault string = `
Author: !!str HerbertHe
GitHub: !!str https://github.com/HerbertHe/pyvm
Gitee: !!str https://gitee.com/HerbertHe/pyvm
Version: !!str v1.0.0
Path: !!map
  PYVM_HOME: !!str  # 命令行文件环境变量
  PYTHON_SYMLINK: !!str # python的默认path
Source: !!str https://pypi.org # pip软件源
`

// configInitCmd represents the configInit command
var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a config file",
	Long:  `pyvm config init`,
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := homedir.Dir()
		var configPath string
		if runtime.GOOS == "windows" {
			configPath = home + `\.pyvm.yaml`
		} else {
			configPath = home + `/.pyvm.yaml`
		}
		if _, err := os.Stat(configPath); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("初始化配置文件")
				configFile, err := os.Create(configPath)
				if err != nil {
					fmt.Printf("创建配置文件失败...\n%v\n", err)
					return
				}
				defer configFile.Close()
				//	写入文件
				_, err = configFile.Write([]byte(yamlDefault))
				if err != nil {
					fmt.Printf("写入文件失败\n%v\n", err)
					return
				}
			}
		} else {
			fmt.Printf("配置文件存在于home目录下: %v\n", configPath)
		}
	},
}

func init() {
	// 是config命令的子命令
	configCmd.AddCommand(configInitCmd)
}
