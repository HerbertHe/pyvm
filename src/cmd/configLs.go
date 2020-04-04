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
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configLsCmd represents the configLs command
var configLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all configuration",
	Long: `pyvm config ls`,
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := homedir.Dir()
		viper.SetConfigName(".pyvm")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)
		if err := viper.ReadInConfig(); err == nil {
			fmt.Printf("Version:\t%v\n", viper.GetString("Version"))
			fmt.Printf("Source:\t%v\n", viper.GetString("Source"))
			fmt.Printf("PYVM_HOME:\t%v\n", viper.GetString("PYVM_HOME"))
			fmt.Printf("PYTHON_SYMLINK:\t%v\n", viper.GetString("PYTHON_SYMLINK"))
		}
	},
}

func init() {
	configCmd.AddCommand(configLsCmd)
}
