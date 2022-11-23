package core

import (
	"app_ws/global"
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Snowflake() {
	conf := global.GVA_CONFIG.Snowflake
	var st time.Time
	st, err := time.Parse("2006-01-02", conf.StartTime)
	if err != nil {
		fmt.Printf("init snowflake failed,err:%v\n", err)
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(conf.MachineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
