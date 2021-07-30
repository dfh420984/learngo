package main

import (
	"log"

	"dfhgrpc.168.cn/ch13/13.2/protocol/dfhgrpc.168.cn/protocol"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := &protocol.UserInfo{
		Message: *proto.String("testInfo"),
		Length: *proto.Int32(10),
	}
	data, err := proto.Marshal(u)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newInfo := &protocol.UserInfo{}
	err = proto.Unmarshal(data, newInfo)
	if err != nil {
		log.Fatal("Unmarshal error: ", err)
	}
	log.Fatalln(newInfo)
}
