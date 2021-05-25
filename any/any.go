package any

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
