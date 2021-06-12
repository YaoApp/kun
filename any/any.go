package any

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// Any the replacement for interface{}
type Any struct {
	value interface{}
}

// Make create a empty instance
func Make() *Any {
	return &Any{value: nil}
}

// Of create a new instance
func Of(value interface{}) *Any {
	return &Any{value: value}
}

// Set set <value> to <v>, and returns the old value.
func (v *Any) Set(value interface{}) (old interface{}) {
	old = v.value
	v.value = value
	return old
}

// Get returns the current value of <v>.
func (v *Any) Get() interface{} {
	return v.value
}

// Val is alias of Get.
func (v *Any) Val() interface{} {
	return v.Get()
}

// Interface is alias of Get.
func (v *Any) Interface() interface{} {
	return v.Get()
}

// Interfaces converts and returns <v> as []interfaces{}.
func (v *Any) Interfaces() []interface{} {
	if v.value == nil {
		return []interface{}{}
	}
	values := reflect.ValueOf(v.value)
	values = reflect.Indirect(values)
	kind := values.Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return []interface{}{v.value}
	}
	res := []interface{}{}
	for i := 0; i < values.Len(); i++ {
		v := values.Index(i).Interface()
		res = append(res, v)
	}
	return res
}

// String returns <v> as string.
func (v *Any) String() string {
	if v.value == nil {
		return ""
	}
	value, ok := v.value.(string)
	if !ok {
		panic("v is not a type of string")
	}
	return value
}

// CString converts and returns <v> as string.
func (v *Any) CString() string {
	value, ok := v.value.(string)
	if ok {
		return value
	}
	return fmt.Sprintf("%v", v.value)
}

// Strings returns <v> as []string.
func (v *Any) Strings() []string {
	if v.value == nil {
		return []string{}
	}
	value, ok := v.value.([]string)
	if !ok {
		panic("v is not a type of []string")
	}
	return value
}

// CStrings converts and returns <v> as []string.
func (v *Any) CStrings() []string {
	if v.value == nil {
		return []string{}
	}

	values := reflect.ValueOf(v.value)
	values = reflect.Indirect(values)
	kind := values.Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return []string{Of(v.value).CString()}
	}
	res := []string{}
	for i := 0; i < values.Len(); i++ {
		v := values.Index(i).Interface()
		res = append(res, Of(v).CString())
	}
	return res
}

// Int returns <v> as int
func (v *Any) Int() int {
	if v.value == nil {
		return 0
	}

	value, ok := v.value.(int)
	if !ok {
		panic("v is not a type of int")
	}
	return value
}

// CInt converts and returns <v> as int
func (v *Any) CInt() int {

	if v.value == nil {
		return 0
	}

	value, ok := v.value.(int)
	if ok {
		return value
	}
	value, err := strconv.Atoi(fmt.Sprintf("%v", v.value))
	if err != nil {
		panic(err.Error())
	}
	return value
}

// Ints returns <v> as []int
func (v *Any) Ints() []int {
	if v.value == nil {
		return []int{}
	}
	value, ok := v.value.([]int)
	if !ok {
		panic("v is not a type of []int")
	}
	return value
}

// CInts converts and returns <v> as []int
func (v *Any) CInts() []int {

	if v.value == nil {
		return []int{}
	}

	values := reflect.ValueOf(v.value)
	values = reflect.Indirect(values)
	kind := values.Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return []int{Of(v.value).CInt()}
	}
	res := []int{}
	for i := 0; i < values.Len(); i++ {
		v := values.Index(i).Interface()
		res = append(res, Of(v).CInt())
	}
	return res
}

// Float is alias of Float64 returns <v> as float64
func (v *Any) Float() float64 {
	return v.Float64()
}

// Float64 returns <v> as float64
func (v *Any) Float64() float64 {
	if v.value == nil {
		return 0
	}

	value, ok := v.value.(float64)
	if !ok {
		panic("v is not a type of float64")
	}
	return value
}

// CFloat is alias of CFloat64 converts and returns <v> as float64
func (v *Any) CFloat() float64 {
	return v.CFloat64()
}

// CFloat64 converts and returns <v> as float64
func (v *Any) CFloat64() float64 {

	if v.value == nil {
		return 0
	}

	value, ok := v.value.(float64)
	if ok {
		return value
	}
	value, err := strconv.ParseFloat(fmt.Sprintf("%v", v.value), 64)
	if err != nil {
		panic(err.Error())
	}
	return value
}

// Floats is alias of Float64s returns <v> as []float64
func (v *Any) Floats() []float64 {
	return v.Float64s()
}

// Float64s returns <v> as []float64
func (v *Any) Float64s() []float64 {
	if v.value == nil {
		return []float64{}
	}
	value, ok := v.value.([]float64)
	if !ok {
		panic("v is not a type of []float64")
	}
	return value
}

// CFloats is alias of CFloat64s converts and returns <v> as []float64
func (v *Any) CFloats() []float64 {
	return v.CFloat64s()
}

// CFloat64s converts and returns <v> as []float64
func (v *Any) CFloat64s() []float64 {

	if v.value == nil {
		return []float64{}
	}

	values := reflect.ValueOf(v.value)
	values = reflect.Indirect(values)
	kind := values.Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return []float64{Of(v.value).CFloat64()}
	}
	res := []float64{}
	for i := 0; i < values.Len(); i++ {
		v := values.Index(i).Interface()
		res = append(res, Of(v).CFloat64())
	}
	return res
}

// Scan for db scan
func (v *Any) Scan(src interface{}) error {
	*v = *Of(src)
	return nil
}

// Value for db driver value
func (v *Any) Value() (driver.Value, error) {
	return v.value, nil
}

// MarshalJSON for json marshalJSON
func (v *Any) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON for json marshalJSON
func (v *Any) UnmarshalJSON(data []byte) error {
	*v = *Of(data)
	return nil
}
