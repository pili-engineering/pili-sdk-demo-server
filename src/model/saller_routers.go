package model

import (
	"cli"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/log"
)

func SallerServer(router *gin.Engine) {
	//create saller
	router.POST("/pili/v1/saller/new", func(c *gin.Context) {
		var reqInfo cli.Sallers
		rErr := c.BindJSON(&reqInfo)
		if rErr != nil {
			ParserError(c, rErr)
			return
		}
		iErr := cli.InsertSaller(&reqInfo)
		if iErr != nil {
			log.Errorf("create saller %s error, %s\n", reqInfo.Name, iErr)
			OperationFailed(c, iErr)
			return
		}
		c.JSON(200, gin.H{
			"code": 200,
			"name": reqInfo.Name,
			"room": reqInfo.Password,
		})
	})

	//saller login
	router.POST("/pili/v1/saller/login", func(c *gin.Context) {
		var reqInfo ReqLoginBody
		rErr := c.BindJSON(&reqInfo)
		if rErr != nil {
			ParserError(c, rErr)
			return
		}
		err := cli.QuerySaller(reqInfo.Name, reqInfo.Password)
		if err != nil {
			OperationFailed(c, err)
			return
		}
		c.JSON(200, gin.H{
			"code":    200,
			"message": "ok",
		})
		return
	})
}
