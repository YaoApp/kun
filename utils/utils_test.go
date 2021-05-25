package utils

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/kun/maps"
)

func TestDD(t *testing.T) {
	if os.Getenv("BE_EXIT") == "1" {
		DD("hello world")
		return
	}
	args := os.Args
	args = append(args, "-test.run=TestDD")
	args = args[1:]
	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = append(os.Environ(), "BE_EXIT=1")
	bytes, err := cmd.Output()
	out := string(bytes)
	assert.Nil(t, err, "the command should be executed success")
	assert.True(t, strings.Contains(out, "hello world"), "the command return value should be have hello...")
}

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
