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
	"strings"
)

var (
	sourceUse string
)

// sourceUseCmd represents the sourceUse command
var sourceUseCmd = &cobra.Command{
	Use:   "use",
	Short: "设置pip源",
	Long: `
pyvm source use -s <source_name>
`,
	Run: func(cmd *cobra.Command, args []string) {
		sources := viper.GetStringMapString("Sources")
		sourceUse = strings.ToLower(sourceUse)
		if _, ok := sources[sourceUse]; ok {
			source := map[string]string{sourceUse: sources[sourceUse]}
			viper.Set("source", source)
			fmt.Printf("已将pip源设置为: %v\n", sourceUse)
		} else {
			fmt.Println("没有找到设置的源，请使用pyvm source ls查看可用源")
		}
	},
}

func init() {
	sourceCmd.AddCommand(sourceUseCmd)
	sourceUseCmd.Flags().StringVarP(&sourceUse, "source", "s", "pypi", "pyvm source use -s <source name>")
}
