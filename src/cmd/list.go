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
	"cn.jieec.pyvm/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	listPython bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List python versions available",
	Long: `
pyvm list 列出所有本地的python环境
pyvm list -p 列出所有可下载的python版本
`,
	Run: func(cmd *cobra.Command, args []string) {
		if listPython {
			fmt.Println("正在请求最新的可用的python版本更新信息...")
			releaseNum, releaseDate := utils.BackPythonVersions()
			if releaseNum == nil || releaseDate == nil {
				fmt.Println("内部错误！")
				return
			}
			for i := 0; i < len(releaseNum); i++ {
				fmt.Printf("\n%v\t\t%v", releaseNum[i], releaseDate[i])
			}
		} else {
			fmt.Println("本地可用的python环境为:")
			//	这个地方有文件统计envs下的环境名称和python版本
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&listPython, "python", "p", false, "pyvm list -p")
}
