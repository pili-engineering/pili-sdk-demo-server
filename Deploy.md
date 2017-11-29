# 安装mysql
服务端程序需要两个数据表来记录账号信息：

* 客户端账号（users）
* 管理后台账号（sallers）

具体数据库可以使用 mydb.sql 文件创建

# 编译代码
首先需要修改 Makefile 文件中的执行路径，然后执行 make 编译源代码
>> export GOPATH=$(GOPATH):/Users/Misty/Workspace/go/src/pili-server-demo

>> make

# 修改 config 文件
```
{
	"server":{
		"listen_host" : "0.0.0.0", # 监听地址
		"listen_port" : 9091, # 监听端口号
		"read_timeout" : 10, # 读取超时时间
		"write_timeout" : 10, # 响应超时时间
		"max_header_bytes" : 4096 # header最大长度
	},
	"app":{
		"alert" : 2.5,
		"access_key":"<AccessKey>", # ak
		"secret_key":"<SecretKey>", # sk
		"hub":"<Hub>", # 直播空间名
 		"prescription":1, # 暂时没用，可以不用管
		"log_file" : "server.log", # log 文件
		"log_level" : "info" # log 输出的等级
	},
	"orm":{
		"driver_name":"mysql", # 数据库
		"data_source":"root:root@tcp(localhost:3306)/pili_server_users?charset=utf8&loc=Asia%2FShanghai", # 连接 mysql 数据库
		"max_idle_conn":30,
		"max_open_conn":50,
		"debug_mode":true
	}
}
```

主要修改的地方是 ak/sk 和 hub 名称，其他均可使用默认值

# 运行
>>./main -c config.conf

出现下面的信息则表示运行成功

```
2017/11/28 18:19:51 [INFO][pili-server-demo] main.go:119: init log
table `users` already exists, skip
table `sallers` already exists, skip
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /assets/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /assets/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[ORM] - 2017-11-28 18:19:51 - [Queries/default] - [  OK /    db.Query /     0.7ms] - [SELECT T0.`name`, T0.`password`, T0.`room`, T0.`deadline` FROM `users` T0 WHERE T0.`deadline` <= ? LIMIT 1000] - `1509272391`
0
[GIN-debug] Loaded HTML Templates (7): 
	- publisher2.tmpl
	- error.tmpl
	- index.tmpl
	- index2.tmpl
	- player.tmpl
	- player2.tmpl
	- publisher.tmpl

[GIN-debug] GET    /pili/v1/server           --> main.main.func2 (3 handlers)
[GIN-debug] POST   /pili/v1/user/new         --> model.UserServer.func1 (3 handlers)
[GIN-debug] GET    /pili/v1/user/query/:name --> model.UserServer.func2 (3 handlers)
[GIN-debug] POST   /pili/v1/user/update/:name --> model.UserServer.func3 (3 handlers)
[GIN-debug] POST   /pili/v1/user/delete/:name --> model.UserServer.func4 (3 handlers)
[GIN-debug] POST   /pili/v1/login            --> model.UserServer.func5 (3 handlers)
[GIN-debug] GET    /pili/v1/room/query/:id   --> model.RoomServer.func1 (3 handlers)
[GIN-debug] POST   /pili/v1/room/new         --> model.RoomServer.func2 (3 handlers)
[GIN-debug] POST   /pili/v1/room/delete/:id  --> model.RoomServer.func3 (3 handlers)
[GIN-debug] POST   /pili/v1/room/token       --> model.RoomServer.func4 (3 handlers)
[GIN-debug] POST   /pili/v1/stream/:id       --> model.StreamServer.func1 (3 handlers)
[GIN-debug] GET    /pili/v1/stream/query/:id --> model.StreamServer.func2 (3 handlers)
[GIN-debug] POST   /pili/v1/saller/new       --> model.SallerServer.func1 (3 handlers)
[GIN-debug] POST   /pili/v1/saller/login     --> model.SallerServer.func2 (3 handlers)
[GIN-debug] Listening and serving HTTP on :9091
```