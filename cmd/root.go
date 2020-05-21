/*
Copyright © 2020 Arnoud Kleinloog <arnoud@kleinloog.ch>

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
	"os"
	"strings"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "httplog",
	Short: "HTTP Logger",
	Long: `Simple HTTP Logger
Listens on port N and logs HTTP traffic.

Start Server:
httplog serve`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("running on port: %d\n", viper.GetInt("port"))

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	fmt.Println("Executing Command")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	fmt.Println("Initializing")

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.httplog.yaml)")

	rootCmd.PersistentFlags().IntP("port", "p", 80, "port numbert (default is 80")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	fmt.Println("Initializing Configuration")

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".httplog" (without extension).
		viper.AddConfigPath("./")
		viper.AddConfigPath(home)
		viper.SetConfigName(".httplog")
	}

	fmt.Println("Binding Flags")
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		fmt.Println("Error binding root flags:", err)
	}

	viper.SetEnvPrefix("HTTPLOG")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
