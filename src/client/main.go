package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	/**
	* *network.MsgParser.Read()读取到的数据如下：
	[]uint8 len: 39, cap: 39, [123,10,9,9,34,72,101,108,108,111,34,58,32,123,10,9,9,9,34,78,97,109,101,34,58,32,34,108,101,97,102,34,10,9,9,125,10,9,125]
	其string()表达如下：
	"{\n\t\t\"Hello\": {\n\t\t\t\"Name\": \"leaf\"\n\t\t}\n\t}"
	*/
	data := []byte(`{
		"Hello": {
			"Name": "leaf"
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)

	for {
		fmt.Println("sleep...")
		time.Sleep(time.Second * 60)
	}
}
