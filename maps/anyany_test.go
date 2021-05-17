package maps

import (
	"fmt"
	"testing"
)

func TestAnyAnySet(t *testing.T) {
	m := MakeAnyAny()
	m.Set("foo", "bar")
	value := m.Get("foo")
	fmt.Printf("%v", value)
}

func TestAnyAnyMerge(t *testing.T) {
	m1 := MakeAnyAny()
	m1.Set("foo1", "bar1")

	m2 := MakeAnyAny()
	m2.Set("foo2", "bar1")

	m1.Merge(m2)
	fmt.Printf("%#v\n", m1)
}
