package server

import (
	"github.com/gin-gonic/gin"
	"zk-ui/pkg/zookeeper"
	"zk-ui/server/config"
	"zk-ui/server/register"
)

type Server struct {
	ZkCli *zookeeper.ZkCli
	Route *gin.Engine
}

func NewService() {
	r := gin.Default()
	zkCli := zookeeper.Init()

	config.Init(zkCli, r)
	register.Init(zkCli, r)
	//监听端口默认为8080
	r.Run(":8080")

}
