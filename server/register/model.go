package register

type Service struct {
	Name      string             `json:"name"`
	Instances []*ServiceInstance `json:"instances"`
}

type ServiceInstance struct {
	InstanceId string            `json:"instanceId"`
	ServiceId  string            `json:"serviceId"`
	Host       string            `json:"host"`
	Port       int               `json:"port"`
	MetaData   map[string]string `json:"metadata"`
}

type ZookeeperInstance struct {
	Id      string           `json:"id"`
	Address string           `json:"address"`
	Port    int              `json:"port"`
	Payload ZookeeperPayload `json:"payload"`
}

type ZookeeperPayload struct {
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	Metadata map[string]string `json:"metadata"`
}
