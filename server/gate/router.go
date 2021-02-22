package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	// 路由到指定模块
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.UserRegisterRequest{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.UserLoginRequest{}, login.ChanRPC)
}
