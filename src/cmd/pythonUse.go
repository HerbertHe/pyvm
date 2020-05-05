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
	version string
)

// pythonUseCmd represents the pythonUse command
var pythonUseCmd = &cobra.Command{
	Use:   "use",
	Short: "设置本地使用的python版本",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("PYVM_HOME") + "/pythons/" + version
		symlink := viper.GetString("PYTHON_SYMLINK")
		if utils.IfDirExists(path) {
			err := utils.SetWindowsSymlink(symlink, path)
			if err != nil {
				fmt.Printf("设置软连接错误:\t%v\n", err)
			} else {
				viper.Set("PythonVersion", version)
			}
		} else {
			fmt.Println("本地无所指定的python版本，请执行pyvm python ls查看本地版本")
		}
	},
}

func init() {
	pythonCmd.AddCommand(pythonUseCmd)
	pythonUseCmd.Flags().StringVarP(&version, "version", "v", "", "pyvm python use -v <version>")
}
