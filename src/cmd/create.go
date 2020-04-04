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
	"strings"

	"github.com/spf13/cobra"
)

var (
	envName string
	pythonVersion string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new python development environment",
	Long: `
pyvm create --env_name=<env_name> --py_version=<py_version>
`,
	Run: func(cmd *cobra.Command, args []string) {
		releases, _ := utils.BackPythonVersions()
		if len(envName) == 0 || len(pythonVersion) == 0 {
			fmt.Println("参数env_name, py_version是需要的")
			fmt.Println("可选python版本:")
			for i := 0; i < len(releases); i++ {
				fmt.Printf("%v\n", releases[i])
			}
			return
		}
		// 提前检查env_name目录是否存在
		if len(envName) != 0 && len(pythonVersion) != 0 {
			// 规范化python版本输入
			pythonVersion = strings.Replace(strings.ToLower(pythonVersion), "p", "P", 1)
			for index, value := range releases {
				if index != len(releases) - 1 && value != pythonVersion {
					continue
				}
				if index == len(releases) - 1 && value != pythonVersion {
					fmt.Println("请求的python版本不存在")
					fmt.Println("可选python版本:")
					for i := 0; i < len(releases); i++ {
						fmt.Printf("%v\n", releases[i])
					}
					return
				}
				if value == pythonVersion {
					// 执行新建命令
					// 下载到PYVM_HOME --> envs
					return
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVar(&envName, "env_name", "", "")
	createCmd.Flags().StringVar(&pythonVersion, "py_version", "","")
}
