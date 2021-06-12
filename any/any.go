package any

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/yaoapp/kun/maps"
	"github.com/yaoapp/kun/num"
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
	value, _ = strconv.Atoi(fmt.Sprintf("%.0f", v.CFloat64()))
	// if err != nil {
	// 	panic(err.Error())
	// }
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

// Bool returns <v> as bool
func (v *Any) Bool() bool {
	if v.value == nil {
		return false
	}

	value, ok := v.value.(bool)
	if !ok {
		panic("v is not a type of bool")
	}
	return value
}

// CBool converts and returns <v> as bool
func (v *Any) CBool() bool {

	if v.value == nil {
		return false
	}

	value, ok := v.value.(bool)
	if ok {
		return value
	}

	value, err := strconv.ParseBool(fmt.Sprintf("%v", v.value))
	if err != nil {
		panic(err.Error())
	}
	return value
}

// Number converts and returns <v> as num.Number
func (v *Any) Number() *num.Number {
	switch v.value.(type) {
	case *num.Number:
		return v.value.(*num.Number)
	case num.Number:
		value := v.value.(num.Number)
		return &value
	default:
		return num.Of(v.value)
	}
}

// Map converts and returns <v> as maps.Map
func (v *Any) Map() maps.Map {
	switch v.value.(type) {
	case maps.Map:
		return v.value.(maps.Map)
	case map[string]interface{}:
		return maps.Of(v.value.(map[string]interface{}))
	}

	// converts to map
	values := reflect.ValueOf(v.value)
	values = reflect.Indirect(values)
	if values.Kind() == reflect.Map {
		valuesMap := map[string]interface{}{}
		for _, key := range values.MapKeys() {
			k := fmt.Sprintf("%v", key)
			valuesMap[k] = values.MapIndex(key).Interface()
		}
		return maps.Of(valuesMap)
	}

	panic("v is not a type of map")
}

// IsNumber checks whether <v> is type of number.
func (v *Any) IsNumber() bool {
	switch v.value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		return true
	default:
		return false
	}
}

// IsMap checks whether <v> is type of map.
func (v *Any) IsMap() bool {
	switch v.value.(type) {
	case map[string]interface{}, maps.Map:
		return true
	default:
		typeof := reflect.TypeOf(v.value)
		return typeof.Kind() == reflect.Map
	}
}

// IsBool checks whether <v> is type of bool.
func (v *Any) IsBool() bool {
	switch v.value.(type) {
	case bool:
		return true
	default:
		return false
	}
}

// IsInt checks whether <v> is type of int.
func (v *Any) IsInt() bool {
	switch v.value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	default:
		return false
	}
}

// IsFloat checks whether <v> is type of float.
func (v *Any) IsFloat() bool {
	switch v.value.(type) {
	case float32, float64:
		return true
	default:
		return false
	}
}

// IsSlice checks whether <v> is type of slice.
func (v *Any) IsSlice() bool {
	values := reflect.ValueOf(v.value)
	values = reflect.Indirect(values)
	return values.Kind() == reflect.Slice
}

// IsArray checks whether <v> is type of array.
func (v *Any) IsArray() bool {
	values := reflect.ValueOf(v.value)
	values = reflect.Indirect(values)
	return values.Kind() == reflect.Array
}

// IsCollection checks whether <v> is type of array or slice.
func (v *Any) IsCollection() bool {
	values := reflect.ValueOf(v.value)
	values = reflect.Indirect(values)
	kind := values.Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

// IsSet checks whether <v> is not nil.
func (v *Any) IsSet() bool {
	return v.value != nil
}

// IsNil checks whether <v> is nil.
func (v *Any) IsNil() bool {
	return v.value == nil
}

// IsEmpty checks whether <v> is empty.
func (v *Any) IsEmpty() bool {
	if v.value == nil {
		return true
	}

	if v.IsInt() {
		return v.CInt() == 0
	}

	if v.IsFloat() {
		return v.CFloat64() == 0.0
	}

	if v.IsBool() {
		return v.Bool() == false
	}

	if v.IsCollection() {
		return len(v.Interfaces()) == 0
	}

	str := v.CString()
	return str == "" || str == "0" || strings.ToLower(str) == "false" || strings.ToLower(str) == "f"
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
