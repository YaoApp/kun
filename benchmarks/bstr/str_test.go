package bstr

import (
	"testing"

	"github.com/yaoapp/kun/str"
)

func BenchmarkASCII(t *testing.B) {
	str.Of("www").ASCII()
}
