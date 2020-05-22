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
	"github.com/akleinloog/http-logger/app"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP Logger",
	Long:  `The HTTP logger logs all requests that are received and answers with status 200 and Request Received.`,
	Run: func(cmd *cobra.Command, args []string) {

		server := app.Instance()
		logger := server.Logger()
		config := server.Config()

		logger.Info().Msgf("Starting server on port: %d\n", config.Server.Port)

		address := fmt.Sprintf("%s:%d", "", config.Server.Port)

		s := &http.Server{
			Addr:         address,
			Handler:      server.Router(),
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
		}

		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Error occurred while listening and serving")
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
