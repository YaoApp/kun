package maps

import (
	"fmt"
	"testing"
)

func TestStrAnySyncSet(t *testing.T) {
	m := MakeStrAnySync()
	m.Set("foo", "bar")
	value := m.Get("foo")
	fmt.Printf("%v", value)
}
