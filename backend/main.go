package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500"}
	r.Use(cors.New(config))

	r.GET("/get_opinion", GetOpinionHandler)

	r.Run()
}
