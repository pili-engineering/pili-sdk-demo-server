package main

import (
	"cli"
	"config"
	// "errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	pili2 "github.com/pili-engineering/pili-sdk-go.v2/pili"
	"github.com/qiniu/log"
	"model"
	"os"
	"runtime"
	"time"
	// "util"
)

const (
	VERSION = "2.0"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var confFile string
	flag.StringVar(&confFile, "c", "", "config file for the service")

	flag.Usage = func() {
		fmt.Println(`
Usage of qasync:
    -c="": config file for the service

version ` + VERSION)
	}

	flag.Parse()

	if confFile == "" {
		fmt.Println("Err: no config file specified")
		os.Exit(1)
	}

	_, statErr := os.Stat(confFile)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			fmt.Println("Err: config file not found")
		} else {
			fmt.Println(statErr)
		}
		os.Exit(1)
	}

	//load config
	cfg, cfgErr := config.LoadConfig(confFile)
	if cfgErr != nil {
		fmt.Println(cfgErr)
		os.Exit(1)
	}

	//init log
	lErr := initLog(cfg.App.QLogLevel, cfg.App.LogFile)
	if lErr != nil {
		fmt.Println("init log error,", lErr)
		os.Exit(1)
	}

	//init orm
	ormErr := cli.InitOrm(&cfg.Orm)
	if ormErr != nil {
		fmt.Println(ormErr)
		os.Exit(1)
	}

	mac := &pili2.MAC{cfg.App.AccessKey, []byte(cfg.App.SecretKey)}
	// client := pili2.New(mac, nil)
	// hub := client.Hub(cfg.App.Hub)

	startTimer(mac)

	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	//************************** monitor ******************************//
	router.GET("/pili/v1/server", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
		})
	})

	model.UserServer(router, cfg)
	model.RoomServer(router, mac)
	model.StreamServer(router, cfg)
	model.SallerServer(router)

	router.Run(fmt.Sprintf(":%d", cfg.Server.ListenPort))
}

//定时删除
func startTimer(mac *pili2.MAC) {
	go func() {
		for {
			ttl := time.Now().Unix() - (30 * 24 * 60 * 60) //second
			cli.DeleteUserByTimer(mac, ttl)
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 23, 59, 59, 59, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

func initLog(logLevel int, logFile string) (err error) {
	log.Info("init log")
	log.SetOutputLevel(logLevel)

	logFp, openErr := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if openErr != nil {
		err = openErr
		return
	}

	log.SetOutput(logFp)

	return
}
