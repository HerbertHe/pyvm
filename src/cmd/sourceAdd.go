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
	sourceName string
	sourceRemote string
)

// sourceAddCmd represents the sourceAdd command
var sourceAddCmd = &cobra.Command{
	Use:   "add",
	Short: "添加自定义pip软件源",
	Long: `
pyvm source add -n <source_name> --remote <source_remote>
`,
	Run: func(cmd *cobra.Command, args []string) {
		sources := viper.GetStringMapString("Sources")
		if len(sourceName) != 0 && len(sourceRemote) != 0 {
			if _, ok := sources[sourceName]; ok {
				fmt.Println("软件源名称已存在，请更换名称或pyvm source ls查看已有pip软件源")
			} else {
				sources[sourceName] = sourceRemote
				viper.Set("Sources", sources)
				fmt.Printf("设置自定义pip软件源成功:\t%v:\t%v\n", sourceName, sourceRemote)
			}
		} else {
			fmt.Println("-n或--remote的值不能为空")
		}
	},
}

func init() {
	sourceCmd.AddCommand(sourceAddCmd)
	sourceAddCmd.Flags().StringVarP(&sourceName, "name", "n", "", "")
	sourceAddCmd.Flags().StringVar(&sourceRemote, "remote", "", "")
}
