package maps

import (
	"fmt"
	"testing"
)

func TestStrStrSet(t *testing.T) {
	m := MakeStrStr()
	m.Set("foo", "bar")
	value := m.Get("foo")
	fmt.Printf("%v", value)
}
