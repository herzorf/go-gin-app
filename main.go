package main

import (
	"github.com/spf13/viper"
	"os"
)

func main() {
	initConfig()
}

func initConfig() {
	dir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(dir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
