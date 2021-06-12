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
	assert.Equal(t, []string{}, v.CStrings())

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
	assert.Equal(t, 0, v.CInt())

	v.Set("20")
	assert.Equal(t, 20, v.CInt())

	v.Set("error")
	assert.Panics(t, func() {
		fmt.Println(v.CInt())
	})
}

func TestInts(t *testing.T) {
	v := Of([]int{1, 2})
	assert.Equal(t, []int{1, 2}, v.Ints())

	v.Set(nil)
	assert.Equal(t, []int{}, v.Ints())

	v.Set("hello")
	assert.Panics(t, func() {
		fmt.Println(v.Ints())
	})
}

func TestCInts(t *testing.T) {
	v := Of([]interface{}{1, 2, "3"})
	assert.Equal(t, []int{1, 2, 3}, v.CInts())

	v.Set(nil)
	assert.Equal(t, []int{}, v.CInts())

	v.Set([]string{"5", "6"})
	assert.Equal(t, []int{5, 6}, v.CInts())

	v.Set(7)
	assert.Equal(t, []int{7}, v.CInts())
}

func TestFloat(t *testing.T) {
	v := Of(10.056)
	assert.Equal(t, 10.056, v.Float())

	v.Set(nil)
	assert.Equal(t, 0.0, v.Float())

	v.Set("hello")
	assert.Panics(t, func() {
		fmt.Println(v.Float())
	})
}

func TestCFloat(t *testing.T) {
	v := Of(10.056)
	assert.Equal(t, 10.056, v.CFloat())

	v.Set(nil)
	assert.Equal(t, 0.0, v.CFloat())

	v.Set("20.018")
	assert.Equal(t, 20.018, v.CFloat())

	v.Set("error")
	assert.Panics(t, func() {
		fmt.Println(v.CFloat())
	})
}

func TestFloats(t *testing.T) {
	v := Of([]float64{1.618, 2.154})
	assert.Equal(t, []float64{1.618, 2.154}, v.Floats())

	v.Set(nil)
	assert.Equal(t, []float64{}, v.Floats())

	v.Set("hello")
	assert.Panics(t, func() {
		fmt.Println(v.Floats())
	})
}

func TestCFloats(t *testing.T) {
	v := Of([]interface{}{1.618, 2.154, "3.617"})
	assert.Equal(t, []float64{1.618, 2.154, 3.617}, v.CFloats())

	v.Set(nil)
	assert.Equal(t, []float64{}, v.CFloats())

	v.Set([]string{"5.10", "6.18"})
	assert.Equal(t, []float64{5.10, 6.18}, v.CFloats())

	v.Set(7.46)
	assert.Equal(t, []float64{7.46}, v.CFloats())
}
