package maps

import (
	"fmt"
	"testing"
)

func TestStrStrSyncSet(t *testing.T) {
	m := MakeStrStrSync()
	m.Set("foo", "bar")
	value := m.Get("foo")
	fmt.Printf("%v", value)
}
