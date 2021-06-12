package bmaps

import (
	"testing"

	"github.com/yaoapp/kun/maps"
)

func BenchmarkMapSet(t *testing.B) {
	m := maps.MakeMap()
	m.Set("foo", "bar")
}

func BenchmarkMapStrSet(t *testing.B) {
	m := maps.MakeMapStr()
	m.Set("foo", "bar")
}
