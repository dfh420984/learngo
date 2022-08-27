package session

import (
	"encoding/binary"
	"io"
	"net"
)

type Session struct {
	conn net.Conn
}

//创建新连接
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

//像链接会话中写数据 网络字节流 header 4 字节， data 数据
func (s *Session) Write(data []byte) error {
	//声明头部+消息体总的数据长度
	buf := make([]byte, 4+len(data))
	//像消息头中写入消息体的数据长度
	//binary 只认固定长度的类型
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	//写入数据
	copy(buf[4:], data)
	//像链接中写入数据
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

//读取链接会话中的数据
func (s *Session) Read() ([]byte, error) {
	//先读取头部长度
	header := make([]byte, 4)
	_, err := io.ReadFull(s.conn, header)
	if err != nil {
		return nil, err
	}
	//读取数据长度
	dataLen := binary.BigEndian.Uint32(header)
	//按长度获取数据
	data := make([]byte, dataLen)
	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
