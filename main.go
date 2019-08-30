package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"flag"

	"github.com/gin-gonic/gin"
	"github.com/hellojukay/gors-server/config"
)

var filename string
var cf *config.Config

func init() {
	flag.StringVar(&filename, "config", "./config.json", "config file")
	if !flag.Parsed() {
		flag.Parse()
	}
	c, err := config.LoadConfig(filename)
	if err != nil {
		fmt.Printf("parse config error : %s ", err)
		os.Exit(1)
	}
	cf = c
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 100 << 20 // 100 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		filename := filepath.Base(file.Filename)
		dir := filepath.Join(cf.Data.Dir, fmt.Sprintf("%d", time.Now().Unix()))
		os.MkdirAll(dir, 0733)
		filename = filepath.Join(dir, filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
	})
	router.Run(fmt.Sprintf(":%d", cf.Server.Port)) // listen and serve on 0.0.0.0:8080
}
