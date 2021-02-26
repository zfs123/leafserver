package internal

import (
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.UserRegisterRequest{}, handleUserRegister)
	handleMsg(&msg.UserLoginRequest{}, handleUserLogin)
}

func handleUserRegister(args []interface{}) {
	m := args[0].(*msg.UserRegisterRequest)
	//a := args[1].(gate.Agent)

	log.Debug("register username: %v passwd: %v", m.GetName(), m.GetPass())

}

func handleUserLogin(args []interface{}) {
	m := args[0].(*msg.UserLoginRequest)
	//a := args[1].(gate.Agent)

	log.Debug("Login %v %v", m.GetName(), m.GetPass())
}
