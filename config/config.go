package config

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("server_addr", ":8080")
	viper.SetDefault("database_url", "sqlite3://shitpost.db")
}

func ServerAddr() string {
	return viper.GetString("server_addr")
}

func DatabaseUrl() string {
	return viper.GetString("database_url")
}
