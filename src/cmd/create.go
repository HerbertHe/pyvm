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
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	envName       string
	pythonVersion string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new python development environment",
	Long:  `pyvm create --env_name=<env_name> --py_version=<py_version>`,
	Run: func(cmd *cobra.Command, args []string) {
		releases, _ := utils.BackPythonVersions()
		if len(envName) == 0 || len(pythonVersion) == 0 {
			fmt.Println("参数env_name, py_version是需要的")
			fmt.Println("可选python版本:")
			for i := 0; i < len(releases); i++ {
				fmt.Printf("%v\t", releases[i])
			}
			return
		}
		// 提前检查env_name目录是否存在
		if len(envName) != 0 && len(pythonVersion) != 0 {
			// 规范化python版本输入
			pythonVersion = strings.Replace(strings.ToLower(pythonVersion), "p", "P", 1)
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
					home, _ := homedir.Dir()
					viper.SetConfigName(".pyvm")
					viper.SetConfigType("yaml")
					viper.AddConfigPath(home)
					if err := viper.ReadInConfig(); err == nil {
						pyvmHome := viper.GetString("PYVM_HOME")
						if len(pyvmHome) != 0 {
							baseName := strings.Replace(pythonVersion, "P", "p", 1)
							baseName = strings.Replace(baseName, "_", "-", 1)
							var fileName string
							var fullName string
							switch runtime.GOOS {
							case "windows":
								{
									fileName = baseName + "-amd64.exe"
									fullName = pyvmHome + `\packages\` + fileName
								}
							case "darwin":
								{
									fileName = baseName + "-macosx10.9.pkg"
									fullName = pyvmHome + `/packages/` + fileName
								}
							}
							if _, err := os.Stat(fullName); err != nil {
								if os.IsNotExist(err) {
									//	文件不存在
									fmt.Printf("正在从淘宝镜像站获取: %v\n", fileName)
									if err := utils.DownloadPython(baseName, fileName, fullName, func(length, downLen int64) {
										// 打印下载进度
										fmt.Printf("文件总长度: %v\t已下载: %v\t已完成: %v\n", length, downLen, float32(downLen)/float32(length))
										if float32(downLen)/float32(length) == float32(1) {
											// 这里跳不出循环
											fmt.Printf("下载完成！文件位于: %v\n", fullName)
											return
										}
									}); err == nil {
									} else {
										fmt.Printf("下载报错:\t%v\n", err)
									}
								} else {
									//	其他错误
									fmt.Printf("发生其他错误: %v\n", err)
								}
							} else {
								//	文件存在
								fmt.Printf("文件已存在于%v\n", fullName)
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
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVar(&envName, "env_name", "", "")
	createCmd.Flags().StringVar(&pythonVersion, "py_version", "", "")
}
