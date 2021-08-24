package grpc

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/yaoapp/kun/grpc/proto"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "MODEL_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"model": &ModelGRPCPlugin{},
}

// Model is the interface that we're exposing as a plugin.
type Model interface {
	Get(name string, payload []byte) ([]byte, error)
	Put(name string, payload []byte) error
}

// ModelGRPCPlugin This is the implementation of plugin.Plugin so we can serve/consume this.
type ModelGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Model
}

// GRPCServer the GRPC Server
func (p *ModelGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterModelServer(s, &ServerGRPC{Impl: p.Impl})
	return nil
}

// GRPCClient the GRPC client
func (p *ModelGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &ClientGRPC{client: proto.NewModelClient(c)}, nil
}
