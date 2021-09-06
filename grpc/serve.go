package grpc

import (
	"io"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// Plugin Here is a real implementation of Pluine that writes to a local file with
// the key name and the contents are the value of the key.
type Plugin struct {
	Logger hclog.Logger
}

// SetLogger set logger output
func (plugin *Plugin) SetLogger(output io.Writer, level Level) {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Level(level),
		Output:     output,
		JSONFormat: true,
	})
	plugin.Logger = logger
}

// Serve GRPC Server
func Serve(model Model) {
	pluginMap := map[string]plugin.Plugin{
		"model": &ModelGRPCPlugin{Impl: model},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}
