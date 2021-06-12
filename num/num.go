package num

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// Number type of numberic
type Number struct {
	value interface{}
}

// Make create a empty instance
func Make() *Number {
	return &Number{value: nil}
}

// Of make a new number
func Of(value interface{}) *Number {
	return &Number{value: value}
}

// Set set <value> to <v>, and returns the old value.
func (n *Number) Set(value interface{}) (old interface{}) {
	old = n.value
	n.value = value
	return old
}

// ToFixed the return value is the type of float64 and keeps the given decimal places
func (n *Number) ToFixed(places int) string {
	format := fmt.Sprintf("%%.%df", places)
	return fmt.Sprintf(format, n.Float64())
}

// Float is alias of Float64 converts and returns as float64
func (n *Number) Float() float64 {
	return n.Float64()
}

// Float64 converts and returns as float64
func (n *Number) Float64() float64 {
	if n.value == nil {
		return 0.0
	}

	switch n.value.(type) {
	case float64:
		return n.value.(float64)
	case float32:
		return float64(n.value.(float32))
	}

	value, err := strconv.ParseFloat(fmt.Sprintf("%v", n.value), 64)
	if err != nil {
		panic(err.Error())
	}
	return value
}

// Float32 converts and returns as float32
func (n *Number) Float32() float32 {
	if n.value == nil {
		return 0.0
	}

	switch n.value.(type) {
	case float64:
		return float32(n.value.(float64))
	case float32:
		return n.value.(float32)
	}

	value, err := strconv.ParseFloat(fmt.Sprintf("%v", n.value), 32)
	if err != nil {
		panic(err.Error())
	}
	return float32(value)
}

// Complex128 the return value is the type of complex128
func (n Number) Complex128() complex128 {
	return 0
}

// Complex64 the return value is the type of complex64
func (n Number) Complex64() complex64 {
	return 0
}

// Int64 the return value is the type of int64 and remove the decimal
func (n Number) Int64() int64 {
	value, ok := n.value.(int64)
	if ok {
		return value
	}
	return int64(n.Int())
}

// Int32 the return value is the type of int32 and remove the decimal
func (n *Number) Int32() int32 {
	value, ok := n.value.(int32)
	if ok {
		return value
	}
	return int32(n.Int())
}

// Int16 the return value is the type of int16 and remove the decimal
func (n *Number) Int16() int16 {
	value, ok := n.value.(int16)
	if ok {
		return value
	}
	return int16(n.Int())
}

// Int8 converts and returns as Int8
func (n *Number) Int8() int8 {
	value, ok := n.value.(int8)
	if ok {
		return value
	}
	return int8(n.Int())
}

// Int converts and returns as Int
func (n *Number) Int() int {
	if n.value == nil {
		return 0
	}

	value, ok := n.value.(int)
	if ok {
		return value
	}

	value, err := strconv.Atoi(fmt.Sprintf("%.0f", n.Float64()))
	if err != nil {
		panic(err.Error())
	}
	return value
}

// Uint64 the return value is the type of uint64 and remove the decimal
func (n Number) Uint64() uint64 {
	return 0
}

// Uint32 the return value is the type of uint32 and remove the decimal
func (n Number) Uint32() uint32 {
	return 0
}

// Uint16 the return value is the type of uint16 and remove the decimal
func (n Number) Uint16() uint16 {
	return 0
}

// Uint8 the return value is the type of uint8 and remove the decimal
func (n Number) Uint8() uint8 {
	return 0
}

// Uint the return value is the type of uint and remove the decimal
func (n Number) Uint() uint {
	return 0
}

// Uintptr the return value is the type of uintptr
func (n Number) Uintptr() uintptr {
	return 0
}

// IsFloat checks whether <v> is type of float.
func (n *Number) IsFloat() bool {
	switch n.value.(type) {
	case float32, float64:
		return true
	default:
		return false
	}
}

// Scan for db scan
func (n *Number) Scan(src interface{}) error {
	*n = *Of(src)
	return nil
}

// Value for db driver value
func (n *Number) Value() (driver.Value, error) {
	return n.value, nil
}

// MarshalJSON for json marshalJSON
func (n *Number) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.value)
}

// UnmarshalJSON for json marshalJSON
func (n *Number) UnmarshalJSON(data []byte) error {
	var v float64
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*n = *Of(v)
	return nil
}
