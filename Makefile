all: env
	export GOPATH=$(GOPATH):/Users/Misty/Workspace/go/src/pili-server-demo
	go build *.go
env:
	go get "github.com/gin-gonic/gin"
	go get "github.com/go-sql-driver/mysql"
	go get "github.com/pili-engineering/pili-sdk-go.v2/pili"
	go get "github.com/qiniu/log"
	go get "github.com/astaxie/beego/orm"

build-linux: env
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pili-server-demo *.go