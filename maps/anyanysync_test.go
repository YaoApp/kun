package maps

import (
	"fmt"
	"testing"
)

func TestAnyAnySyncSet(t *testing.T) {
	m := MakeAnyAnySync()
	m.Set("foo", "bar")
	value := m.Get("foo")
	fmt.Printf("%v", value)
}
