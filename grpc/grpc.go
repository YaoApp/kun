package grpc

import (
	"encoding/json"

	"github.com/yaoapp/kun/grpc/proto"
	"golang.org/x/net/context"
)

// ClientGRPC is an implementation of KV that talks over RPC.
type ClientGRPC struct{ client proto.ModelClient }

// Exec execute the plugin model
func (m *ClientGRPC) Exec(name string, args ...interface{}) (*Response, error) {

	payload, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	res, err := m.client.Exec(context.Background(), &proto.Request{
		Name:    name,
		Payload: payload,
	})
	if err != nil {
		return nil, err
	}

	return &Response{Bytes: res.Response}, nil
}

// ServerGRPC Here is the gRPC server that ClientGRPC talks to.
type ServerGRPC struct {
	// This is the real implementation
	Impl Model
}

// Exec ServerGet
func (m *ServerGRPC) Exec(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	args := []interface{}{}
	err := json.Unmarshal(req.Payload, &args)
	if err != nil {
		return nil, err
	}
	v, err := m.Impl.Exec(req.Name, args...)
	return &proto.Response{Response: v.Bytes}, err
}
