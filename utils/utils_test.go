package utils

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/kun/maps"
)

func TestDump(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	Dump(maps.Str{
		"foo": "bar",
		"nested": maps.Str{
			"foo": "bar",
		},
	})
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	assert.True(t, strings.Contains(string(out), "foo"), "the command return value should be have foo...")
}

func TestDumpString(t *testing.T) {
	Dump("hello world")
}

func TestDumpNumber(t *testing.T) {
	Dump(1024)
}
