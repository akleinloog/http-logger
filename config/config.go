package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var cfgFile string

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
