package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

// 雪花算法 生成ID
func Init(startTime string, nodeId int64) (err error) {
	var t time.Time
	t, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	sf.Epoch = t.UnixNano() / 1000000
	node, err = sf.NewNode(nodeId)
	if err != nil {
		return err
	}
	return
}

// 生成通用ID
func GenerateId() int64 {
	return node.Generate().Int64()
}
