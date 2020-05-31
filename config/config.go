package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config represents application connector credentials and db name
type Config struct {
	Server   string
	Database string
	LogPath  string
}

// var (
// // LogPath config logs destination
// )

// func init() {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("toml")
// 	viper.AddConfigPath(filepath.Dir("../config"))
// 	viper.ReadInConfig()

// 	// To add a default value :
// 	viper.SetDefault("LOG_PATH", "log/rest.log")
// 	//To get from the toml file or env var
// 	LogPath = viper.GetString("LOG_PATH")
// }

// Read and parse the config file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
