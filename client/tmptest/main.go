package main

import (
	"encoding/binary"
	"fmt"
	"net"

	"client/msg"

	"github.com/golang/protobuf/proto"
)

func sendUserRegisterRequest(conn net.Conn) {
	test := &msg.UserRegisterRequest{
		// 使用辅助函数设置域的值
		Name: proto.String("LeafName"),
		Pass: proto.String("LeafPass"),
	}
	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling send msg error: ", err)
	}
	// len + id + data
	m := make([]byte, 4+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)+2))
	binary.BigEndian.PutUint16(m[2:], uint16(1))
	copy(m[4:], data)
	// 发送消息
	conn.Write(m)
}

func sendUserLoginRequest(conn net.Conn) {
	test := &msg.UserLoginRequest{
		// 使用辅助函数设置域的值
		Name: proto.String("LeafName"),
		Pass: proto.String("LeafPass"),
	}
	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling send msg error: ", err)
	}
	// len + id + data
	m := make([]byte, 4+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)+2))
	binary.BigEndian.PutUint16(m[2:], uint16(2))
	copy(m[4:], data)
	// 发送消息
	conn.Write(m)
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	test := &msg.Hello{
		// 使用辅助函数设置域的值
		Name: proto.String("Leaf"),
	}
	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling send msg error: ", err)
	}
	// len + id + data
	m := make([]byte, 4+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)+2))
	copy(m[4:], data)
	// 发送消息
	conn.Write(m)
	buf := make([]byte, 32)
	// 接收消息
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read buf error: ", err)
	}
	recv := &msg.Hello{}

	len := binary.BigEndian.Uint16(buf)
	id := binary.BigEndian.Uint16(buf[2:])
	err = proto.Unmarshal(buf[4:n], recv)
	if err != nil {
		fmt.Println("marshaling recv msg error: ", err)
	}
	fmt.Println(len, id)
	fmt.Println(recv.GetName())
	fmt.Println(recv.String())

	sendUserRegisterRequest(conn)
	sendUserLoginRequest(conn)

}
