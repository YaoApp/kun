package grpc

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/yaoapp/kun/grpc/proto"
	"google.golang.org/grpc"
)

// Level represents a log level.
type Level int32

const (
	// NoLevel is a special level used to indicate that no level has been
	// set and allow for a default to be used.
	NoLevel Level = 0

	// Trace is the most verbose level. Intended to be used for the tracing
	// of actions in code, such as function enters/exits, etc.
	Trace Level = 1

	// Debug information for programmer lowlevel analysis.
	Debug Level = 2

	// Info information about steady state operations.
	Info Level = 3

	// Warn information about rare but handled events.
	Warn Level = 4

	// Error information about unrecoverable events.
	Error Level = 5

	// Off disables all logging output.
	Off Level = 6
)

// Model is the interface that we're exposing as a plugin.
type Model interface {
	Exec(name string, args ...interface{}) (*Response, error)
}

// ModelGRPCPlugin This is the implementation of plugin.Plugin so we can serve/consume this.
type ModelGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Model
}

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "GOU_MODEL_PLUGIN",
	MagicCookieValue: "GOU VER0.6.0",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"model": &ModelGRPCPlugin{},
}

// Response GRPC Response
type Response struct {
	Bytes []byte
	Type  string
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
