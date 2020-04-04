/*
Copyright © 2020 HerbertHe

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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "pyvm",
	Short: "A tool for helping you manage python version!",
	Long: `
Help you to manage python version better!
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// 初始化读取配置文件和环境变量
func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".pyvm")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // 自动读取匹配的环境变量

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file not found, please run `pyvm config init` to generate a new config file")
	}
}
