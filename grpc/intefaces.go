package grpc

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/go-plugin"
	"github.com/yaoapp/kun/grpc/proto"
	"github.com/yaoapp/kun/maps"
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
	Exec(name string, args ...interface{}) (*Response, error)
}

// Response GRPC Response
type Response struct {
	Bytes []byte
}

// Bind bind struct
func (res Response) Bind(v interface{}) error {
	return json.Unmarshal(res.Bytes, &v)
}

// Map cast to map
func (res Response) Map() (maps.MapStrAny, error) {
	v := maps.MakeMapStrAny()
	err := json.Unmarshal(res.Bytes, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// MustMap cast to map
func (res Response) MustMap() maps.MapStrAny {
	v := maps.MakeMapStrAny()
	err := json.Unmarshal(res.Bytes, &v)
	if err != nil {
		panic(err)
	}
	return v
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
