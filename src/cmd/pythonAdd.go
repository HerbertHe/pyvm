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
	"os"
	"runtime"
	"strings"
)

var (
	pythonVersion string
)

// pythonAddCmd represents the pythonAdd command
var pythonAddCmd = &cobra.Command{
	Use:   "add",
	Short: "添加一个新的python版本",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		releases, _ := utils.BackPythonVersions()
		if len(pythonVersion) == 0 {
			fmt.Println("参数py_version是需要的")
			fmt.Println("可选python版本:")
			for i := 0; i < len(releases); i++ {
				fmt.Printf("%v\t", releases[i])
				if i % 7 == 0 {
					fmt.Printf("\n")
				}
			}
			return
		}
		if len(pythonVersion) != 0 {
			pythonVersion = utils.FormatPythonInput(pythonVersion)
			for index, value := range releases {
				if index != len(releases)-1 && value != pythonVersion {
					continue
				}
				if index == len(releases)-1 && value != pythonVersion {
					fmt.Println("请求的python版本不存在")
					fmt.Println("可选python版本:")
					for i := 0; i < len(releases); i++ {
						fmt.Printf("%v\t", releases[i])
						if i % 7 == 0 {
							fmt.Printf("\n")
						}
					}
					return
				}
				if value == pythonVersion {
					if err := viper.ReadInConfig(); err == nil {
						pyvmHome := viper.GetString("PYVM_HOME")
						if len(pyvmHome) != 0 {
							baseName := strings.Replace(pythonVersion, "P", "p", 1)
							baseName = strings.Replace(baseName, "_", "-", 1)
							var fileName string
							var fullName string
							var pythonLocal string
							switch runtime.GOOS {
							case "windows":
								{
									fileName = baseName + "-amd64.exe"
									fullName = pyvmHome + `\packages\` + fileName
									pythonLocal = pyvmHome + `\pythons\` + pythonVersion
								}
							case "darwin":
								{
									fileName = baseName + "-macosx10.9.pkg"
									fullName = pyvmHome + `/packages/` + fileName
									pythonLocal = pyvmHome + `/pythons/` + pythonVersion
								}
							}
							_, err := os.Stat(fullName)
							_, err1 := os.Stat(pythonLocal)
							// 文件都不存在
							if err1 != nil && err != nil {
								if os.IsNotExist(err) && os.IsNotExist(err1) {
									//	文件不存在
									fmt.Printf("正在从淘宝镜像站获取: %v\n", fileName)
									if err := utils.DownloadPython(baseName, fileName, fullName, func(length, downLen int64) {
										// 打印下载进度
										fmt.Printf("文件总长度: %v\t已下载: %v\t已完成: %.2f%%\n", length, downLen, float32(downLen)/float32(length)*100)
									}); err == nil {
									} else {
										fmt.Printf("下载报错:\t%v\n", err)
									}
									// passive安装
									if err = utils.InstallPythonPassive(fullName, pythonLocal); err != nil {
										fmt.Printf("执行安装时发生错误...\n%v\n", err)
									} else {
										fmt.Printf("python环境安装成功, 位于%v\n", pythonLocal)
									}
								} else {
									//	其他错误
									fmt.Printf("发生其他错误: %v\n%v\n", err, err1)
								}
							}
							// 安装包存在，没有安装
							if err == nil && err1 != nil {
								if os.IsNotExist(err1) {
									fmt.Printf("文件已存在 %v\n", fullName)
									fmt.Println("正在执行安装...")
									// passive安装
									if err = utils.InstallPythonPassive(fullName, pythonLocal); err != nil {
										fmt.Printf("执行安装时发生错误...\n%v\n", err)
									} else {
										fmt.Printf("python环境安装成功, 位于%v\n", pythonLocal)
									}
								} else {
									fmt.Printf("发生其他错误: %v\n", err1)
								}
							}
						} else {
							fmt.Printf("没有配置PYVM_HOME参数, 请在%v文件中配置，为pyvm的安装目录\n", viper.ConfigFileUsed())
						}
					}
					break
				}
			}
		}
	},
}

func init() {
	pythonCmd.AddCommand(pythonAddCmd)
	pythonAddCmd.Flags().StringVar(&pythonVersion, "py_version", "", "")
}
