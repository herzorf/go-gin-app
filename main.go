package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"githun.com/herzorf/go-gin-app/common"
	"os"
)

func main() {
	initConfig()
	common.InitDB()

	route := gin.Default()
	port := viper.GetString("server.port")
	err := route.Run(":" + port)
	if err != nil {
		panic(err)
	}
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
