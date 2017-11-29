package model

import (
	"cli"
	"config"
	"github.com/gin-gonic/gin"
	"github.com/pili-engineering/pili-sdk-go.v2/pili"
	"github.com/qiniu/log"
	"time"
	"util"
)

func UserServer(router *gin.Engine, cfg *config.RtcConfig) {
	mac := &pili.MAC{cfg.App.AccessKey, []byte(cfg.App.SecretKey)}
	//************************** User **********************************
	//create user
	router.POST("/pili/v1/user/new", func(c *gin.Context) {
		//check admin authorization
		token := c.Request.Header.Get("Authorization")
		_, authErr := util.AuthorityOfAdmin(token)
		if authErr != nil {
			AuthorityFailed(c)
			return
		}
		var reqInfo cli.Users
		rErr := c.BindJSON(&reqInfo)
		if rErr != nil {
			ParserError(c, rErr)
			return
		}
		if reqInfo.Room == "" {
			log.Errorf("create user ,the room must not be null")
			ParameterNull(c)
			return
		}
		reqInfo.Deadline = time.Now().Unix()
		//create room
		_, nrErr := cli.RoomCreate(mac, reqInfo.Room, reqInfo.Name, 99)
		if nrErr != nil {
			log.Errorf("create room failed, %s\n", nrErr)
			OperationFailed(c, nrErr)
			return
		}
		iErr := cli.InsertUser(&reqInfo)
		if iErr != nil {
			log.Errorf("create user %s error, %s\n", reqInfo.Name, iErr)
			//delete room
			cli.RoomDelete(mac, reqInfo.Room)
			OperationFailed(c, iErr)
			return
		}
		c.JSON(200, gin.H{
			"code": 200,
			"name": reqInfo.Name,
			"room": reqInfo.Room,
		})
	})

	//query user
	router.GET("/pili/v1/user/query/:name", func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			ParameterNull(c)
			return
		}
		u, uErr := cli.UserIsExisted(name)
		if uErr != nil {
			log.Errorf("query user failed, %s\n", name)
			OperationFailed(c, uErr)
			return
		}
		c.JSON(200, gin.H{
			"code": 200,
			"name": u.Name,
			"room": u.Room,
		})
		return
	})

	//update user
	router.POST("/pili/v1/user/update/:name", func(c *gin.Context) {
		//check authorization
		token := c.Request.Header.Get("Authorization")
		author, authErr := util.Authority(token)
		if authErr != nil {
			AuthorityFailed(c)
			return
		}
		//get request paramter
		name := c.Param("name")
		if name == "" {
			ParameterNull(c)
			return
		}
		if name != author {
			AuthorityFailed(c)
			return
		}
		var reqInfo ReqUpdateUser
		rErr := c.BindJSON(&reqInfo)
		if rErr != nil {
			ParserError(c, rErr)
			return
		}
		udErr := cli.UpdateUser(name, reqInfo.Password)
		if udErr != nil {
			OperationFailed(c, udErr)
			return
		}
		c.JSON(200, gin.H{
			"code":    200,
			"message": "ok",
		})
		return
	})

	//delete user
	router.POST("/pili/v1/user/delete/:name", func(c *gin.Context) {
		//check authorization
		token := c.Request.Header.Get("Authorization")
		author, authErr := util.Authority(token)
		if authErr != nil {
			AuthorityFailed(c)
			return
		}
		//get request paramter
		name := c.Param("name")
		if name == "" {
			ParameterNull(c)
			return
		}
		if name != author {
			AuthorityFailed(c)
			return
		}
		//删除用户
		_, retErr := cli.DeleteUser(mac, name)
		if retErr != nil {
			OperationFailed(c, retErr)
			return
		}
		c.JSON(200, gin.H{
			"code":    200,
			"message": "ok",
		})
		return
	})

	//login
	router.POST("/pili/v1/login", func(c *gin.Context) {
		var reqInfo ReqLoginBody
		rErr := c.BindJSON(&reqInfo)
		if rErr != nil {
			ParserError(c, rErr)
			return
		}
		err := cli.QueryUser(reqInfo.Name, reqInfo.Password)
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
