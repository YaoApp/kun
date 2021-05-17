package maps

import (
	"fmt"
	"testing"
)

func TestStrAnySet(t *testing.T) {
	m := MakeStrAny()
	m.Set("foo", "bar")
	value := m.Get("foo")
	fmt.Printf("%v", value)
}
