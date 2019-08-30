package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"flag"

	"github.com/gin-gonic/gin"
	"github.com/hellojukay/gors-server/config"
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
}

var c = config.LoadConfig("", nil)

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
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
