package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zk-ui/pkg/zookeeper"
)

type Cli struct {
	zkCli *zookeeper.ZkCli
}

func Init(zkCli *zookeeper.ZkCli, r *gin.Engine) {
	cli := &Cli{
		zkCli: zkCli,
	}
	configApi := r.Group("/configs")
	{
		configApi.GET("/", cli.list)
		configApi.GET("/:configId", cli.detail)
	}

}

func (cli *Cli) list(c *gin.Context) {
	list, err := cli.zkCli.GetConfigList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, list)
}

func (cli *Cli) detail(c *gin.Context) {
	configId := c.Param("configId")

	props, err := cli.zkCli.GetConfigDetail(configId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	
	content := ""
	for _, prop := range props {
		content += prop.String() + "\n"
	}
	resp := &DetailResp{
		Content: content,
	}

	c.JSON(http.StatusOK, resp)
}
