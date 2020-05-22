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
package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	// configFile points to the configuration file.
	configFile string

	// DefaultConfig holds the default configuration.
	DefaultConfig = AppConfig()
)

type Config struct {
	Debug  bool `env:"DEBUG, defaults to false"`
	Server serverConf
	//Db     dbConf
}

type serverConf struct {
	Port int `env:"HTTPLOG_PORT, defaults to 80"`
	//TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	//TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	//TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

//type dbConf struct {
//	Host     string `env:"DB_HOST,required"`
//	Port     int    `env:"DB_PORT,required"`
//	Username string `env:"DB_USER,required"`
//	Password string `env:"DB_PASS,required"`
//	DbName   string `env:"DB_NAME,required"`
//}

// AppConfig returns the application configuration.
func AppConfig() *Config {

	var config Config

	config.Debug = viper.GetBool("debug")

	config.Server.Port = viper.GetInt("port")

	return &config
}

// Initialize initializes the viper configuration.
func Initialize(rootCommand *cobra.Command) {

	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
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
	if err := viper.BindPFlags(rootCommand.Flags()); err != nil {
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
