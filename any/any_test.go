package any

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/kun/day"
	"github.com/yaoapp/kun/maps"
	"github.com/yaoapp/kun/num"
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

func TestInterfaces(t *testing.T) {
	v := Of([]interface{}{"hello", 1, true})
	assert.Equal(t, []interface{}{"hello", 1, true}, v.Interfaces())

	v.Set(nil)
	assert.Equal(t, []interface{}{}, v.Interfaces())

	v.Set([]string{"hello", "world"})
	assert.Equal(t, []interface{}{"hello", "world"}, v.Interfaces())

	v.Set("hello")
	assert.Equal(t, []interface{}{"hello"}, v.Interfaces())
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

func TestBool(t *testing.T) {
	v := Of(false)
	assert.Equal(t, false, v.Bool())

	v.Set(true)
	assert.Equal(t, true, v.Bool())

	v.Set(nil)
	assert.Equal(t, false, v.Bool())

	v.Set("hello")
	assert.Panics(t, func() {
		fmt.Println(v.Bool())
	})
}

func TestCBool(t *testing.T) {
	v := Of(false)
	assert.Equal(t, false, v.CBool())

	v.Set(true)
	assert.Equal(t, true, v.CBool())

	v.Set(nil)
	assert.Equal(t, false, v.CBool())
	v.Set("true")
	assert.Equal(t, true, v.CBool())

	v.Set("T")
	assert.Equal(t, true, v.CBool())

	v.Set("TRUE")
	assert.Equal(t, true, v.CBool())

	v.Set("1")
	assert.Equal(t, true, v.CBool())

	v.Set(1)
	assert.Equal(t, true, v.CBool())

	v.Set("false")
	assert.Equal(t, false, v.CBool())

	v.Set("FALSE")
	assert.Equal(t, false, v.CBool())

	v.Set("F")
	assert.Equal(t, false, v.CBool())

	v.Set("0")
	assert.Equal(t, false, v.CBool())

	v.Set(0)
	assert.Equal(t, false, v.CBool())

	v.Set("-10")
	assert.Panics(t, func() {
		fmt.Println(v.CBool())
	})
}

func TestNumber(t *testing.T) {
	num1 := Of(0.618).Number()
	assert.Equal(t, 0.618, num1.Float())

	num2 := Of(1.618 + 0.532i).Number()
	assert.Equal(t, 1.618+0.532i, num2.Complex())

	num3 := Of(num.Of(0.618)).Number()
	assert.Equal(t, 1, num3.Int())

	num4 := Of(*num.Of(0.618)).Number()
	assert.Equal(t, 1, num4.Int())
}

func TestDatetime(t *testing.T) {
	day.Timezone("Beijing", 8*60*60)
	assert.Equal(t, 31, Of("2019-12-31 08:20:55").Datetime().Day())
	assert.Equal(t, 31, Of(day.Of("2019-12-31 08:20:55")).Datetime().Day())
	assert.Equal(t, 31, Of(*day.Of("2019-12-31 08:20:55")).Datetime().Day())

	name, offset := Of("2019-12-31 08:20:55").Datetime().Zone()
	assert.Equal(t, "Beijing", name)
	assert.Equal(t, 8*60*60, offset)
}

func TestMap(t *testing.T) {
	map1 := Of(map[string]interface{}{
		"hello": "world",
		"foo":   1,
		"array": []string{"one", "two", "three"},
		"name": map[string]interface{}{
			"first":  "Join",
			"second": "Cool",
		}}).Map()
	assert.Equal(t, "world", map1.Get("hello"))
	assert.Equal(t, "world", map1.Any("hello").String())
	assert.Equal(t, "Join", map1.Dot().Get("name.first"))
	assert.Equal(t, "Join", map1.Dot().Any("name.first").String())
	assert.Equal(t, nil, map1.Dot().Get("name.notexist"))
	assert.Equal(t, nil, map1.Dot().Any("name.notexist").Interface())
	assert.Equal(t, true, map1.Flatten().Any("name.notexist").IsNil())
	assert.Equal(t, "two", map1.Flatten().Any("array.1").String())

	map1.Set("title", "CEO")
	map2 := Of(map1).Map()
	assert.Equal(t, "CEO", map2.Get("title"))
	assert.Equal(t, "CEO", map2.Any("title").String())

	map3 := Of(map[int]string{0: "0", 1: "1"}).Map()
	assert.Equal(t, "0", map3.Get("0"))
	assert.Equal(t, "0", map3.Any("0").String())
	assert.Equal(t, "1", map3.Get("1"))
	assert.Equal(t, "1", map3.Any("1").String())

	map4 := Of(maps.Of(map[string]interface{}{"a": "b"})).Map()
	assert.Equal(t, "b", map4.Get("a"))

	assert.Panics(t, func() {
		Of([]string{"hello", "world"}).Map()
	})
}

func TestIsDatetime(t *testing.T) {
	assert.Equal(t, true, Of(day.Of("2000-01-01")).IsDatetime())
	assert.Equal(t, true, Of(time.Now()).IsDatetime())
	assert.Equal(t, false, Of(1).IsDatetime())
}

func TestIsNumber(t *testing.T) {
	assert.Equal(t, true, Of(1).IsNumber())
	assert.Equal(t, true, Of(0.618).IsNumber())
	assert.Equal(t, true, Of(1+0.618i).IsNumber())
	assert.Equal(t, false, Of([]string{"hello", "world"}).IsNumber())
}

func TestIsMap(t *testing.T) {
	assert.Equal(t, true, Of(map[string]interface{}{"hello": "world"}).IsMap())
	assert.Equal(t, true, Of(map[int]string{0: "0", 1: "1"}).IsMap())
	assert.Equal(t, false, Of([]string{"hello", "world"}).IsMap())
}

func TestIsBool(t *testing.T) {
	assert.Equal(t, true, Of(true).IsBool())
	assert.Equal(t, true, Of(false).IsBool())
	assert.Equal(t, false, Of(0).IsBool())
	assert.Equal(t, false, Of(1).IsBool())
}

func TestIsInt(t *testing.T) {
	assert.Equal(t, true, Of(1).IsInt())
	assert.Equal(t, true, Of(int8(1)).IsInt())
	assert.Equal(t, false, Of(true).IsInt())
}

func TestIsFloat(t *testing.T) {
	assert.Equal(t, true, Of(1.618).IsFloat())
	assert.Equal(t, true, Of(float32(1.382)).IsFloat())
	assert.Equal(t, false, Of(5).IsFloat())
}

func TestIsSlice(t *testing.T) {
	assert.Equal(t, true, Of([]string{"hello", "world"}).IsSlice())
	assert.Equal(t, false, Of([2]string{"hello", "world"}).IsSlice())
	assert.Equal(t, false, Of(5).IsSlice())
}

func TestIsArray(t *testing.T) {
	assert.Equal(t, true, Of([2]string{"hello", "world"}).IsArray())
	assert.Equal(t, false, Of([]string{"hello", "world"}).IsArray())
	assert.Equal(t, false, Of(5).IsArray())
}

func TestIsCollection(t *testing.T) {
	assert.Equal(t, true, Of([2]string{"hello", "world"}).IsCollection())
	assert.Equal(t, true, Of([]string{"hello", "world"}).IsCollection())
	assert.Equal(t, false, Of(5).IsCollection())
}

func TestIsNotNil(t *testing.T) {
	assert.Equal(t, true, Of(1).IsNotNil())
	assert.Equal(t, false, Of(nil).IsNotNil())
}

func TestIsNil(t *testing.T) {
	assert.Equal(t, false, Of(1).IsNil())
	assert.Equal(t, true, Of(nil).IsNil())
}

func TestIsEmpty(t *testing.T) {
	assert.Equal(t, true, Of(nil).IsEmpty())
	assert.Equal(t, true, Of("0").IsEmpty())
	assert.Equal(t, true, Of("F").IsEmpty())
	assert.Equal(t, true, Of("false").IsEmpty())
	assert.Equal(t, true, Of("").IsEmpty())
	assert.Equal(t, true, Of(0).IsEmpty())
	assert.Equal(t, true, Of(0.0).IsEmpty())
	assert.Equal(t, true, Of(false).IsEmpty())
	assert.Equal(t, true, Of([]string{}).IsEmpty())
	assert.Equal(t, true, Of([0]string{}).IsEmpty())
	assert.Equal(t, false, Of(true).IsEmpty())
}
