package session

import (
	"bytes"
	"encoding/gob"
)

//定义rpc交互的数据格式
type RPCData struct {
	//访问时的函数
	Name string
	//访问的参数
	Args []interface{}
}

//编码
func encode(data RPCData) ([]byte, error) {
	var buf bytes.Buffer
	//得到字节数组编码器
	bufEnc := gob.NewEncoder(&buf)
	//对数据进行编码
	if err := bufEnc.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//解码
func decode(b []byte) (RPCData, error) {
	buf := bytes.NewBuffer(b)
	//字节数组解码器
	bufDec := gob.NewDecoder(buf)
	var data RPCData
	if err := bufDec.Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}
