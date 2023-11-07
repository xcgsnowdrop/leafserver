package internal

import (
	"server/conf"
	"server/game"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

type Module struct {
	*gate.Gate
}

// 初始化gate对象，除一堆配置意外，最重要的是初始化Processor和AgentChanRPC
// Processor为实现了*network.Processor接口的对象，该接口包含Route(),Unmarshal(),Marshal()三个方法

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}
}
