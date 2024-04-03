package zookeeper

import (
	"bytes"
	"github.com/go-zookeeper/zk"
	"log"
	"time"
)

const (
	ConfigPath  = "/config/"
	ServicePath = "/service/"
)

type ZkCli struct {
	zkConn *zk.Conn
}

func Init() *ZkCli {
	zkAddr := "172.30.3.50:2181"
	conn, _, err := zk.Connect([]string{zkAddr}, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	return &ZkCli{
		zkConn: conn,
	}

}

func (cli *ZkCli) GetConfigList() ([]string, error) {
	dir, _, err := cli.zkConn.Children(ConfigPath)
	if err != nil {
		log.Println("Get config list err", err)
		return nil, err
	}
	return dir, nil
}

func (cli *ZkCli) GetConfigDetail(configId string) ([]*Prop, error) {
	path := ConfigPath + configId
	dir, _, err := cli.zkConn.Children(path)
	if err != nil {
		log.Println("Get config list err", err)
		return nil, err
	}

	var props []*Prop
	for _, key := range dir {
		data, _, err := cli.zkConn.Get(path + "/" + key)
		if err != nil {
			return nil, err
		}
		buffer := bytes.Buffer{}
		buffer.Write(data)
		props = append(props, &Prop{
			Key:   key,
			Value: buffer.String(),
		})
	}

	return props, nil
}

func (cli *ZkCli) GetServiceList() ([]string, error) {
	path := "/services"
	dir, _, err := cli.zkConn.Children(path)
	if err != nil {
		log.Println("Get config list err", err)
		return nil, err
	}
	return dir, nil
}

func (cli *ZkCli) GetInstanceList(service string) ([]string, error) {
	path := "/services/" + service
	dir, _, err := cli.zkConn.Children(path)
	if err != nil {
		log.Println("Get config list err", err)
		return nil, err
	}
	return dir, nil
}

func (cli *ZkCli) GetInstanceDetail(service, instanceId string) ([]byte, error) {
	path := "/services/" + service + "/" + instanceId
	data, _, err := cli.zkConn.Get(path)
	if err != nil {
		log.Println("Get config list err", err)
		return nil, err
	}
	return data, nil
}
