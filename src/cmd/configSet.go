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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	pyvmHome string
	pythonSymlink string
)

// configSetCmd represents the configSet command
var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "修改pyvm配置",
	Long: `
请谨慎修改PYVM_HOME和PYTHON_SYMLINK的值，否则将会影响使用

pyvm config set --PYVM_HOME=<path>
pyvm config set --PYTHON_SYMLINK=<path>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(pyvmHome) != 0 {
			viper.Set("PYVM_HOME", pyvmHome)
			fmt.Printf("已将pyvm根目录改为:\t%v\n", pyvmHome)
		}

		if len(pythonSymlink) != 0 {
			viper.Set("PYTHON_SYMLINK", pythonSymlink)
			fmt.Printf("已将python软链接路径改为:\t%v\n", pythonSymlink)
		}
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
	configSetCmd.Flags().StringVar(&pyvmHome, "PYVM_HOME", "", "pyvm config set --PYVM_HOME=<path>")
	configSetCmd.Flags().StringVar(&pythonSymlink, "PYTHON_SYMLINK", "", "pyvm config set --PYTHON_SYMLINK=<path>")
}
