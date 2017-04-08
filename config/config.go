// provides access to config files magically and fixed the issue of
// loading config file before init() and wrapper from github.com/spf13/viper

package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	// Name of the file default is config
	Name = "config"
)

func init() {
	// do not find config file on test
	// if flag.Lookup("test.v") != nil {
	// 	return
	// }

	// in production config file must be along side with the binary
	if err := New("."); err != nil {
		log.Fatalln(err)
	}
}

func New(path string) error {
	viper.SetConfigName(Name)
	viper.AddConfigPath(path)

	return viper.ReadInConfig()
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
