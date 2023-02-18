package fastgrpc

import (
	"fastgameserver/framework/network/gamerpc"
	"fastgameserver/framework/protocol/pbgrpc"
)

type grpcStreamChanAcceptor struct {
	stream pbgrpc.PbGameRPC_StreamServer
	local  string
	remote string
}

func (g *grpcStreamChanAcceptor) Local() string {
	return g.local
}

func (g *grpcStreamChanAcceptor) Remote() string {
	return g.remote
}

func (g *grpcStreamChanAcceptor) Recv(m *gamerpc.GMessage) error {
	if m == nil {
		return nil
	}

	msg, err := g.stream.Recv()
	if err != nil {
		return err
	}

	m.Header = msg.Header
	m.Body = msg.Body
	return nil
}

func (g *grpcStreamChanAcceptor) Send(m *gamerpc.GMessage) error {
	if m == nil {
		return nil
	}

	return g.stream.Send(&pbgrpc.PbGMessage{
		Header: m.Header,
		Body:   m.Body,
	})
}

func (g *grpcStreamChanAcceptor) Close() error {
	return nil
}
