package main

import (
	"fmt"
	"os"

	"flag"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/hellojukay/gors-server/config"
	"github.com/hellojukay/gors-server/routers"
)

var filename string

func init() {
	flag.StringVar(&filename, "config", "./config.json", "config file")
	if !flag.Parsed() {
		flag.Parse()
	}
	err := config.LoadConfig(filename)
	if err != nil {
		fmt.Printf("parse config error : %s \n", err)
		os.Exit(1)
	}
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 100 << 20 // 100 MiB

	router.Use(static.Serve("/", static.LocalFile("./data/", false)))
	router.Use(static.Serve("/static/", static.LocalFile("./public/", false)))
	router.POST("/upload", routers.Upload)
	router.GET("/records", routers.Records)
	router.Run(fmt.Sprintf(":%d", config.C.Server.Port)) // listen and serve on 0.0.0.0:8080
}
