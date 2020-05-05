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
	"github.com/spf13/viper"
)

var (
	listRemotePython bool
)

// pythonLsCmd represents the pythonLs command
var pythonLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "列出python环境",
	Long: `
pyvm python ls 列出本地环境
pyvm python ls --remote 列出可下载版本
`,
	Run: func(cmd *cobra.Command, args []string) {
		if listRemotePython {
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
			fmt.Println("本地python环境为:")
			dirs, err := utils.GetLocalPythonVersions()
			if err != nil {
				fmt.Printf("查询发生错误:\t%v\n", err)
			}

			if len(dirs) == 0 {
				fmt.Println("本地未安装python")
			} else {
				for _, fv := range dirs {
					// 判断正在使用的版本
					if viper.Get("PythonVersion").(string) == fv {
						fmt.Printf("%v\t*\n", fv)
						continue
					}
					fmt.Println(fv)
				}
			}
		}
	},
}

func init() {
	pythonCmd.AddCommand(pythonLsCmd)
	pythonLsCmd.Flags().BoolVar(&listRemotePython, "remote", false, "pyvm python ls --remote")
}
