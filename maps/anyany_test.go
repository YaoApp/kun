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
	m2.Set("foo2", "bar2")

	m3 := MakeAnyAnySync()
	m3.Set("foo3", "bar3")

	m1.Merge(m2, m3)
	fmt.Printf("%#v\n", m1)

	m3.Merge(m1, m2)
	m3.Range(func(key, value interface{}) bool {
		fmt.Printf("%v=%v\n", key, value)
		return true
	})

}
