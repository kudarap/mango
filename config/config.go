// provides access to config files magically and fixed the issue of
// loading config file before init() and wrapper from github.com/spf13/viper

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	// Name of the file default is config
	Name = "config"
	// Path of the file default is .
	Path = "."
)

func init() {
	viper.SetConfigName(Name)
	viper.AddConfigPath(Path)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// GetString from config
func GetString(n string) string {
	return viper.GetString(n)
}

// GetBool from config
func GetBool(n string) bool {
	return viper.GetBool(n)
}

// GetInt from config
func GetInt(n string) int {
	return viper.GetInt(n)
}
