package log

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type output struct{ data []byte }

func (out output) String() string {
	return string(out.data)
}

func (out output) Map() map[string]interface{} {
	v := map[string]interface{}{}
	json.Unmarshal(out.data, &v)
	return v
}

func (out *output) Write(p []byte) (n int, err error) {
	out.data = p
	return len(out.data), nil
}

func (out *output) Clean() {
	out.data = nil
}

var testout = &output{}

func TestSetting(t *testing.T) {
	SetOutput(testout)
	SetFormatter(JSON)
	SetLevel(TraceLevel)
	Trace("hello %s", "world")
	data := testout.Map()
	assert.Equal(t, "trace", data["level"])
	assert.Equal(t, "hello world", data["msg"])

	SetFormatter(TEXT)
	Trace("hello %s", "world")
	assert.Contains(t, testout.String(), `msg="hello world"`)
	testout.Clean()
}

func TestWith(t *testing.T) {
	SetOutput(testout)
	SetFormatter(JSON)
	With(F{"foo": "bar"}).Info("hello %s", "world")
	data := testout.Map()
	assert.Equal(t, "info", data["level"])
	assert.Equal(t, "hello world", data["msg"])
	assert.Equal(t, "bar", data["foo"])
	testout.Clean()
}

func TestTrace(t *testing.T) {
	SetOutput(testout)
	SetFormatter(JSON)
	SetLevel(TraceLevel)
	Trace("hello %s", "world")
	data := testout.Map()
	assert.Equal(t, "trace", data["level"])
	assert.Equal(t, "hello world", data["msg"])
	testout.Clean()
}

func TestDebug(t *testing.T) {
	SetOutput(testout)
	SetFormatter(JSON)
	SetLevel(DebugLevel)
	Trace("hello %s", "world")
	assert.Nil(t, testout.data)

	Debug("hello %s", "world")
	data := testout.Map()
	assert.Equal(t, "debug", data["level"])
	assert.Equal(t, "hello world", data["msg"])
	testout.Clean()
}

func TestInfo(t *testing.T) {
	SetOutput(testout)
	SetFormatter(JSON)
	SetLevel(InfoLevel)
	Trace("hello %s", "world")
	assert.Nil(t, testout.data)

	Debug("hello %s", "world")
	assert.Nil(t, testout.data)

	Info("hello %s", "world")
	data := testout.Map()
	assert.Equal(t, "info", data["level"])
	assert.Equal(t, "hello world", data["msg"])
	testout.Clean()
}

func TestWarn(t *testing.T) {
	SetOutput(testout)
	SetFormatter(JSON)
	SetLevel(WarnLevel)
	Trace("hello %s", "world")
	assert.Nil(t, testout.data)

	Debug("hello %s", "world")
	assert.Nil(t, testout.data)

	Info("hello %s", "world")
	assert.Nil(t, testout.data)

	Warn("hello %s", "world")
	data := testout.Map()
	assert.Equal(t, "warning", data["level"])
	assert.Equal(t, "hello world", data["msg"])
	testout.Clean()
}

func TestError(t *testing.T) {
	SetOutput(testout)
	SetFormatter(JSON)
	SetLevel(ErrorLevel)
	Trace("hello %s", "world")
	assert.Nil(t, testout.data)

	Debug("hello %s", "world")
	assert.Nil(t, testout.data)

	Info("hello %s", "world")
	assert.Nil(t, testout.data)

	Warn("hello %s", "world")
	assert.Nil(t, testout.data)

	Error("hello %s", "world")
	data := testout.Map()
	assert.Equal(t, "error", data["level"])
	assert.Equal(t, "hello world", data["msg"])
	testout.Clean()
}

func TestPanics(t *testing.T) { //  Calls panic() after logging
	SetOutput(testout)
	SetFormatter(JSON)
	SetLevel(PanicLevel)
	Trace("hello %s", "world")
	assert.Nil(t, testout.data)

	Debug("hello %s", "world")
	assert.Nil(t, testout.data)

	Info("hello %s", "world")
	assert.Nil(t, testout.data)

	Warn("hello %s", "world")
	assert.Nil(t, testout.data)

	Error("hello %s", "world")
	assert.Nil(t, testout.data)

	assert.Panics(t, func() {
		Panic("hello %s", "world")
	})
	data := testout.Map()
	assert.Equal(t, "panic", data["level"])
	assert.Equal(t, "hello world", data["msg"])
	testout.Clean()
}
