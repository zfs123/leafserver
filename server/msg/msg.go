package msg
import (
	"github.com/name5566/leaf/network/protobuf"
)

var (
	Processor = protobuf.NewProcessor()
)

func init() {	// 这里我们注册 protobuf 消息)
    Processor.SetByteOrder(true)
    Processor.Register(&Hello{})
    Processor.Register(&UserRegisterRequest{})
    Processor.Register(&UserLoginRequest{})
    Processor.Register(&Response{})
    Processor.Register(&PlayerChuPai{})

}