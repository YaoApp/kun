package num

import (
	"database/sql/driver"
	"encoding/json"
)

// Number type of numberic
type Number struct {
	value interface{}
}

// Of make a new number
func Of(value interface{}) Number {
	return Number{value: value}
}

// ToFixed the return value is the type of float64 and keeps the given decimal places
func (n Number) ToFixed(places int) float64 {
	return 0
}

// Float64 the return value is the type of float64
func (n Number) Float64() float64 {
	return 0
}

// Float32 the return value is the type of float32
func (n Number) Float32() float32 {
	return 0
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
	return 0
}

// Int32 the return value is the type of int32 and remove the decimal
func (n Number) Int32() int32 {
	return 0
}

// Int16 the return value is the type of int16 and remove the decimal
func (n Number) Int16() int16 {
	return 0
}

// Int8 the return value is the type of int8 and remove the decimal
func (n Number) Int8() int8 {
	return 0
}

// Int the return value is the type of int and remove the decimal
func (n Number) Int() int {
	return 0
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

// Scan for db scan
func (n *Number) Scan(src interface{}) error {
	*n = Of(src)
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
	*n = Of(v)
	return nil
}
