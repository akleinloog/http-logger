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

// rootCmd is the base command, executed when called without sub-commands
var rootCmd = &cobra.Command{
	Use:   "httplog",
	Short: "HTTP Logger",
	Long: `Simple HTTP Logger
`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

// init is called before Execute is called to execute the command.
func init() {

	cobra.OnInitialize(initConfig)

	// Persistent flags are global for all commands.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.httplog.yaml)")

	rootCmd.PersistentFlags().IntP("port", "p", 80, "port numbert (default is 80")

	// Local flags are specific for this command.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Execute executes the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig initializes the viper configuration.
func initConfig() {

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

	// Binding Flags
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
