package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/log"
)

func AuthorityFailed(c *gin.Context) {
	c.JSON(403, gin.H{
		"code":    403,
		"message": "no auhtorized",
	})
	c.Abort()
}

func ParserError(c *gin.Context, err error) {
	log.Errorf("request paramter error,%s\n", err.Error())
	c.JSON(400, gin.H{
		"code":    400,
		"message": "the request's paramter error",
	})
	c.Abort()
}

func ParameterNull(c *gin.Context) {
	c.JSON(400, gin.H{
		"code":    400,
		"message": "one parameters is null",
	})
	c.Abort()
}

func OperationFailed(c *gin.Context, err error) {
	c.JSON(611, gin.H{
		"code":    611,
		"message": fmt.Sprintf("%s\n", err),
	})
	c.Abort()
}
