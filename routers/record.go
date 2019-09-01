package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hellojukay/gors-server/database"
)

func Records(c *gin.Context) {
	recodrs, err := database.AllRecords()
	if err != nil {
		c.String(500, "读取记录失败")
		c.Abort()
		return
	}
	c.JSON(200, recodrs)
}
