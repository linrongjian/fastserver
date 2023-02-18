package fastgrpc

import (
	"fastgameserver/framework/network/gamerpc"
	"fastgameserver/framework/protocol/pbgrpc"

	"google.golang.org/grpc"
)

type grpcStreamChanConnector struct {
	conn   *grpc.ClientConn
	stream pbgrpc.PbGameRPC_StreamClient

	local  string
	remote string
}

func (g *grpcStreamChanConnector) Local() string {
	return g.local
}

func (g *grpcStreamChanConnector) Remote() string {
	return g.remote
}

func (g *grpcStreamChanConnector) Recv(m *gamerpc.GMessage) error {
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

func (g *grpcStreamChanConnector) Send(m *gamerpc.GMessage) error {
	if m == nil {
		return nil
	}

	return g.stream.Send(&pbgrpc.PbGMessage{
		Header: m.Header,
		Body:   m.Body,
	})
}

func (g *grpcStreamChanConnector) Close() error {
	return g.conn.Close()
}
