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

// sourceLsCmd represents the sourceLs command
var sourceLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "列出所有的pip可用源",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		for k, _ := range viper.GetStringMapString("Source") {
			for key, value := range viper.GetStringMapString("Sources") {
				if k == key {
					fmt.Printf("%v:\t%v\t*\n", key, value)
					continue
				}
				fmt.Printf("%v:\t%v\t\n", key, value)
			}
		}
	},
}

func init() {
	sourceCmd.AddCommand(sourceLsCmd)
}
