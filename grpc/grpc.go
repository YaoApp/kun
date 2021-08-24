package grpc

import (
	"github.com/yaoapp/kun/grpc/proto"
	"golang.org/x/net/context"
)

// ClientGRPC is an implementation of KV that talks over RPC.
type ClientGRPC struct{ client proto.ModelClient }

// Put Model Put
func (m *ClientGRPC) Put(name string, payload []byte) error {
	_, err := m.client.Put(context.Background(), &proto.PutRequest{
		Name:    name,
		Payload: payload,
	})
	return err
}

// Get Model get
func (m *ClientGRPC) Get(name string, payload []byte) ([]byte, error) {
	resp, err := m.client.Get(context.Background(), &proto.GetRequest{
		Name:    name,
		Payload: payload,
	})
	if err != nil {
		return nil, err
	}

	return resp.Response, nil
}

// ServerGRPC Here is the gRPC server that ClientGRPC talks to.
type ServerGRPC struct {
	// This is the real implementation
	Impl Model
}

// Put Server Put
func (m *ServerGRPC) Put(
	ctx context.Context,
	req *proto.PutRequest) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Put(req.Name, req.Payload)
}

// Get ServerGet
func (m *ServerGRPC) Get(
	ctx context.Context,
	req *proto.GetRequest) (*proto.GetResponse, error) {
	v, err := m.Impl.Get(req.Name, req.Payload)
	return &proto.GetResponse{Response: v}, err
}
