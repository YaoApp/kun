package any

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	v := Make()
	assert.Empty(t, v.String())
}

func TestOf(t *testing.T) {
	v := Of("hello")
	assert.Equal(t, "hello", v.String())
}

func TestSet(t *testing.T) {
	v := Of("hello")
	assert.Equal(t, "hello", v.Get())
	assert.Equal(t, "hello", v.Val())
	assert.Equal(t, "hello", v.Interface())

	v.Set(1)
	assert.Equal(t, "1", v.CString())
}

func TestGet(t *testing.T) {
	v := Of("hello")
	v.Set(1)
	assert.Equal(t, 1, v.Get())
	assert.Equal(t, 1, v.Val())
	assert.Equal(t, 1, v.Interface())
}

func TestString(t *testing.T) {
	v := Of("hello")
	assert.Equal(t, "hello", v.String())

	v.Set(nil)
	assert.Equal(t, "", v.String())

	v.Set(1)
	assert.Panics(t, func() {
		fmt.Println(v.String())
	})
}

func TestCString(t *testing.T) {
	v := Of("hello")
	assert.Equal(t, "hello", v.CString())

	v.Set(nil)
	assert.Equal(t, "", v.String())

	v.Set(1)
	assert.Equal(t, "1", v.CString())
}

func TestStrings(t *testing.T) {
	v := Of([]string{"hello", "world"})
	assert.Equal(t, []string{"hello", "world"}, v.Strings())

	v.Set(nil)
	assert.Equal(t, []string{}, v.Strings())

	v.Set("hello")
	assert.Panics(t, func() {
		fmt.Println(v.Strings())
	})
}

func TestCStrings(t *testing.T) {
	v := Of([]interface{}{"hello", 1, true})
	assert.Equal(t, []string{"hello", "1", "true"}, v.CStrings())

	v.Set(nil)
	assert.Equal(t, []string{}, v.Strings())

	v.Set([]string{"hello", "world"})
	assert.Equal(t, []string{"hello", "world"}, v.CStrings())

	v.Set("hello")
	assert.Equal(t, []string{"hello"}, v.CStrings())
}

func TestInt(t *testing.T) {
	v := Of(10)
	assert.Equal(t, 10, v.Int())

	v.Set(nil)
	assert.Equal(t, 0, v.Int())

	v.Set("hello")
	assert.Panics(t, func() {
		fmt.Println(v.Int())
	})
}

func TestCInt(t *testing.T) {
	v := Of(10)
	assert.Equal(t, 10, v.CInt())

	v.Set(nil)
	assert.Equal(t, 0, v.Int())

	v.Set("20")
	assert.Equal(t, 20, v.CInt())
}
