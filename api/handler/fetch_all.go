package handler

import (
	"log"
	"net/http"

	"github.com/dingqing/registry/global"
	"github.com/gin-gonic/gin"
)

func FetchAllHandler(c *gin.Context) {
	log.Println("request api/fetchall...")

	data := global.Discovery.Registry.FetchAll()
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "",
		"data":    data,
	})
}
