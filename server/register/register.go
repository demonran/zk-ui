package register

import (
	"encoding/json"
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
	serviceApi := r.Group("/services")
	{
		serviceApi.GET("/", cli.list)
	}

}

func (cli *Cli) list(c *gin.Context) {
	list, err := cli.zkCli.GetServiceList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	var services []*Service
	for _, item := range list {
		instanceList, err := cli.zkCli.GetInstanceList(item)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		var instances []*ServiceInstance
		for _, i := range instanceList {

			data, err := cli.zkCli.GetInstanceDetail(item, i)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, err)
				return
			}
			var zookeeperInstance *ZookeeperInstance
			json.Unmarshal(data, &zookeeperInstance)
			instances = append(instances, &ServiceInstance{
				InstanceId: zookeeperInstance.Id,
				ServiceId:  item,
				Host:       zookeeperInstance.Address,
				Port:       zookeeperInstance.Port,
				MetaData:   zookeeperInstance.Payload.Metadata,
			})
		}
		services = append(services, &Service{
			Name:      item,
			Instances: instances,
		})
	}
	c.JSON(http.StatusOK, services)
}
