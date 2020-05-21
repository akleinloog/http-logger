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
	"github.com/akleinloog/http-logger/app"
	"github.com/akleinloog/http-logger/config"
	"github.com/spf13/cobra"
	"os"
)

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

	cobra.OnInitialize(func() {
		config.Initialize(rootCmd)
	})

	// Persistent flags are global for all commands.

	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is empty)")

	rootCmd.PersistentFlags().IntP("port", "p", 80, "port numbert (default is 80")

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode (default is false)")
	// Local flags are specific for this command.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Execute executes the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := rootCmd.Execute(); err != nil {

		logger := app.Instance().Logger()
		logger.Fatal().Err(err).Msg("Fatal during root command execution")

		os.Exit(1)
	}
}
