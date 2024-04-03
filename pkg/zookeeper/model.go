package zookeeper

import "fmt"

type Prop struct {
	Key   string
	Value any
}

func (p *Prop) String() string {
	return p.Key + ":" + fmt.Sprint(p.Value)
}
