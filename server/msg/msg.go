package msg

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network/protobuf"
)

// 使用 Protobuf 消息处理器
var Processor = protobuf.NewProcessor()

func init() {
	// 注册消息
	Processor.Register(&Hello{})
	id1 := Processor.Register(&UserRegisterRequest{})
	id2 := Processor.Register(&UserLoginRequest{})

	log.Debug("id1: %v", id1)
	log.Debug("id2: %v", id2)

}
