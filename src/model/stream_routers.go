package model

import (
	"config"
	"github.com/gin-gonic/gin"
	"github.com/pili-engineering/pili-sdk-go.v2/pili"
	"util"
)

func StreamServer(router *gin.Engine, cfg *config.RtcConfig) {
	mac := &pili.MAC{cfg.App.AccessKey, []byte(cfg.App.SecretKey)}
	client := pili.New(mac, nil)
	hub := client.Hub(cfg.App.Hub)

	//************************* Stream *******************************//
	//create stream
	router.POST("/pili/v1/stream/:id", func(c *gin.Context) {
		//check authorization
		token := c.Request.Header.Get("Authorization")
		_, authErr := util.Authority(token)
		if authErr != nil {
			AuthorityFailed(c)
			return
		}

		id := c.Params.ByName("id")
		//先创建流
		hub.Create(id)
		url := pili.RTMPPublishURL("pili-publish.ps.pdex-service.com", cfg.App.Hub, id, mac, 3600)
		c.JSON(200, gin.H{
			"code": 200,
			"url":  url,
		})
		return
	})

	//query live url
	router.GET("/pili/v1/stream/query/:id", func(c *gin.Context) {
		//check authorization
		token := c.Request.Header.Get("Authorization")
		_, authErr := util.Authority(token)
		if authErr != nil {
			AuthorityFailed(c)
			return
		}

		key := c.Params.ByName("id")
		rtmpurl := pili.RTMPPlayURL("pili-live-rtmp.ps.pdex-service.com", cfg.App.Hub, key)
		hlsurl := pili.HLSPlayURL("pili-live-hdl.ps.pdex-service.com", cfg.App.Hub, key)
		hdlurl := pili.HDLPlayURL("pili-live-hls.ps.pdex-service.com", cfg.App.Hub, key)
		c.JSON(200, gin.H{
			"code": 200,
			"rtmp": rtmpurl,
			"hdl":  hdlurl,
			"hls":  hlsurl,
		})
		return
	})
}
