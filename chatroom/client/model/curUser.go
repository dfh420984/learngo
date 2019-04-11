package model

import (
	"net"
	"learngo/chatroom/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}