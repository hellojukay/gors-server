package routers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hellojukay/gors-server/config"
	"github.com/hellojukay/gors-server/database"
)

func Upload(c *gin.Context) {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	originFilename := filepath.Base(file.Filename)
	timePath := fmt.Sprintf("%d", time.Now().Unix())
	originFilePath := filepath.Join(timePath, originFilename)
	dir := filepath.Join(config.C.Data.Dir, timePath)
	os.MkdirAll(dir, 0733)
	filename := filepath.Join(dir, originFilename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	database.SaveRecord(originFilename, originFilePath)
	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
}
