package num

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"regexp"
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
func (n Number) ToFixed(places int) string {
	format := fmt.Sprintf("%%.%df", places)
	return fmt.Sprintf(format, n.Float64())
}

// Float is alias of Float64 converts and returns as float64
func (n Number) Float() float64 {
	return n.Float64()
}

// Float64 converts and returns as float64
func (n Number) Float64() float64 {
	if n.value == nil {
		return 0.0
	}

	switch n.value.(type) {
	case float64:
		return n.value.(float64)
	case float32:
		return float64(n.value.(float32))
	case complex128:
		return real(n.value.(complex128))
	case complex64:
		return float64(real(n.value.(complex64)))
	}

	value, err := strconv.ParseFloat(fmt.Sprintf("%v", n.value), 64)
	if err != nil {
		panic(err.Error())
	}
	return value
}

// Float32 converts and returns as float32
func (n Number) Float32() float32 {
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

// Complex is alias of Complex128 converts and returns as complex128
func (n Number) Complex() complex128 {
	return n.Complex128()
}

// Complex128 converts and returns as complex128
func (n Number) Complex128() complex128 {
	value, ok := n.value.(complex128)
	if ok {
		return value
	}

	value64, ok := n.value.(complex64)
	if ok {
		return complex128(value64)
	}

	valueStr, ok := n.value.(string)
	if ok {
		// 1.56+2.48i
		re := regexp.MustCompile(`[ \()]*([0-9\.]+)[ ]*\+[ ]*([0-9\.]+)i[ \)]*`)
		matched := re.FindStringSubmatch(valueStr)
		if len(matched) > 0 {
			return complex(Of(matched[1]).Float64(), Of(matched[2]).Float64())
		}

		// (1.56,2.48i)
		re = regexp.MustCompile(`\([ ]*([0-9\.]+)[ ]*,[ ]*([0-9\.]+)[ ]*\)`)
		matched = re.FindStringSubmatch(valueStr)
		if len(matched) > 0 {
			return complex(Of(matched[1]).Float64(), Of(matched[2]).Float64())
		}
	}
	return complex(n.Float64(), 0)
}

// Complex64 converts and returns as Complex64
func (n Number) Complex64() complex64 {
	value, ok := n.value.(complex64)
	if ok {
		return value
	}

	value128, ok := n.value.(complex128)
	if ok {
		return complex64(value128)
	}

	valueStr, ok := n.value.(string)
	if ok {
		// 1.56+2.48i
		re := regexp.MustCompile(`[ \()]*([0-9\.]+)[ ]*\+[ ]*([0-9\.]+)i[ \)]*`)
		matched := re.FindStringSubmatch(valueStr)
		if len(matched) > 0 {
			return complex(Of(matched[1]).Float32(), Of(matched[2]).Float32())
		}

		// (1.56,2.48i)
		re = regexp.MustCompile(`\([ ]*([0-9\.]+)[ ]*,[ ]*([0-9\.]+)[ ]*\)`)
		matched = re.FindStringSubmatch(valueStr)
		if len(matched) > 0 {
			return complex(Of(matched[1]).Float32(), Of(matched[2]).Float32())
		}
	}
	return complex(n.Float32(), 0.0)
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
func (n Number) Int32() int32 {
	value, ok := n.value.(int32)
	if ok {
		return value
	}
	return int32(n.Int())
}

// Int16 the return value is the type of int16 and remove the decimal
func (n Number) Int16() int16 {
	value, ok := n.value.(int16)
	if ok {
		return value
	}
	return int16(n.Int())
}

// Int8 converts and returns as Int8
func (n Number) Int8() int8 {
	value, ok := n.value.(int8)
	if ok {
		return value
	}
	return int8(n.Int())
}

// Int converts and returns as Int
func (n Number) Int() int {
	if n.value == nil {
		return 0
	}

	value, ok := n.value.(int)
	if ok {
		return value
	}

	value, _ = strconv.Atoi(fmt.Sprintf("%.0f", n.Float64()))
	return value
}

// Uint64 the return value is the type of uint64 and remove the decimal
func (n Number) Uint64() uint64 {
	value, ok := n.value.(uint64)
	if ok {
		return value
	}
	return uint64(n.Int())
}

// Uint32 the return value is the type of uint32 and remove the decimal
func (n Number) Uint32() uint32 {
	value, ok := n.value.(uint32)
	if ok {
		return value
	}
	return uint32(n.Int())
}

// Uint16 the return value is the type of uint16 and remove the decimal
func (n Number) Uint16() uint16 {
	value, ok := n.value.(uint16)
	if ok {
		return value
	}
	return uint16(n.Int())
}

// Uint8 the return value is the type of uint8 and remove the decimal
func (n Number) Uint8() uint8 {
	value, ok := n.value.(uint8)
	if ok {
		return value
	}
	return uint8(n.Int())
}

// Uint the return value is the type of uint and remove the decimal
func (n Number) Uint() uint {
	value, ok := n.value.(uint)
	if ok {
		return value
	}
	return uint(n.Int())
}

// Uintptr the return value is the type of uintptr
func (n Number) Uintptr() uintptr {
	value, ok := n.value.(uintptr)
	if ok {
		return value
	}
	return uintptr(n.Int())
}

// IsSet checks whether <v> is not nil.
func (n Number) IsSet() bool {
	return n.value != nil
}

// IsNil checks whether <v> is nil.
func (n Number) IsNil() bool {
	return n.value == nil
}

// IsInt checks whether <v> is type of int.
func (n Number) IsInt() bool {
	switch n.value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	default:
		return false
	}
}

// IsFloat checks whether <v> is type of float.
func (n Number) IsFloat() bool {
	switch n.value.(type) {
	case float32, float64:
		return true
	default:
		return false
	}
}

// IsComplex checks whether <v> is type of complex.
func (n Number) IsComplex() bool {
	switch n.value.(type) {
	case complex128, complex64:
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
