package utils

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/kun/exception"
	"github.com/yaoapp/kun/maps"
)

func TestDump(t *testing.T) {
	var buf bytes.Buffer
	exception.SetWriter(&buf)
	defer exception.SetWriter(nil) // reset to os.Stdout via GetWriter fallback

	Dump(maps.Str{
		"foo": "bar",
		"nested": maps.Str{
			"foo": "bar",
		},
	})
	assert.True(t, strings.Contains(buf.String(), "foo"), "the command return value should be have foo...")
}

func TestDumpString(t *testing.T) {
	Dump("hello world", "foo", "bar")
}

func TestDumpNumber(t *testing.T) {
	Dump(1024, 0.618)
}

func TestDumpMix(t *testing.T) {
	Dump(
		"hello world",
		1024, 0.618,
		map[string]interface{}{"foo": "bar", "number": 1024},
		[]interface{}{"foo", "bar", 1024, 0.618, map[string]interface{}{"foo": "bar", "number": 1024}},
	)
}
