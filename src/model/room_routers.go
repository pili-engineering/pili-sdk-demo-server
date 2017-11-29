package model

import (
	"cli"
	"github.com/gin-gonic/gin"
	"github.com/pili-engineering/pili-sdk-go.v2/pili"
	// "github.com/qiniu/log"
	"util"
)

func RoomServer(router *gin.Engine, mac *pili.MAC) {
	//**************************** Room ****************************
	//query room
	router.GET("/pili/v1/room/query/:id", func(c *gin.Context) {
		//get request paramter
		roomid := c.Param("id")
		if roomid == "" {
			ParameterNull(c)
			return
		}
		r, err := cli.RoomStatus(mac, roomid)
		if err != nil {
			OperationFailed(c, err)
			return
		}
		c.JSON(200, gin.H{
			"code":    200,
			"room":    r.Room,
			"ownerId": r.OwnerUserID,
			"userMax": r.UserMax,
			"status":  r.Status,
		})
	})

	//create room
	router.POST("/pili/v1/room/new", func(c *gin.Context) {
		//check authorization
		token := c.Request.Header.Get("Authorization")
		_, authErr := util.Authority(token)
		if authErr != nil {
			AuthorityFailed(c)
			return
		}
		var reqInfo ReqNewRoomBody
		rErr := c.BindJSON(&reqInfo)
		if rErr != nil {
			ParserError(c, rErr)
			return
		}
		if reqInfo.Room == "" || reqInfo.User == "" {
			ParameterNull(c)
			return
		}
		if reqInfo.Max <= 0 {
			reqInfo.Max = 99
		}
		ret, retErr := cli.RoomCreate(mac, reqInfo.Room, reqInfo.User, reqInfo.Max)
		if retErr != nil {
			OperationFailed(c, retErr)
			return
		}
		c.JSON(200, gin.H{
			"code": 200,
			"room": ret.Room,
		})
	})

	//delete room
	router.POST("/pili/v1/room/delete/:id", func(c *gin.Context) {
		//check authorization
		token := c.Request.Header.Get("Authorization")
		author, authErr := util.Authority(token)
		if authErr != nil {
			AuthorityFailed(c)
			return
		}

		roomid := c.Param("id")
		//query room status
		r, err := cli.RoomStatus(mac, roomid)
		if err != nil {
			OperationFailed(c, err)
			return
		}
		if author != r.OwnerUserID {
			AuthorityFailed(c)
			return
		}
		_, retErr := cli.RoomDelete(mac, roomid)
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

	//new roomtoken
	router.POST("/pili/v1/room/token", func(c *gin.Context) {
		//check authorization
		token := c.Request.Header.Get("Authorization")
		_, authErr := util.Authority(token)
		if authErr != nil {
			AuthorityFailed(c)
			return
		}

		var reqInfo ReqNewRoomTokenBody
		rErr := c.BindJSON(&reqInfo)
		if rErr != nil {
			ParserError(c, rErr)
			return
		}
		roomtoken := cli.CreateToken(mac, reqInfo.Room, reqInfo.User, reqInfo.Version)
		c.String(200, roomtoken)
	})

}
