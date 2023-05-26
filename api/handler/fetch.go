package handler

import (
	"log"
	"net/http"

	"github.com/dingqing/registry/global"
	"github.com/dingqing/registry/model"
	"github.com/dingqing/registry/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func FetchHandler(c *gin.Context) {
	log.Println("request api/fetch...")
	var req model.RequestFetch
	if e := c.ShouldBindJSON(&req); e != nil {
		err := errcode.ParamError
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code(),
			"message": err.Error(),
		})
		return
	}

	//fetch
	fetchData, err := global.Discovery.Registry.Fetch(req.Env, req.AppId, req.Status, 0)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code(),
			"data":    "",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    fetchData,
		"message": "",
	})
}
