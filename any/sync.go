package any

import (
	"fmt"
	"reflect"
	"sync/atomic"
)

// Sync the replacement for concurrent safe interface
type Sync struct {
	value *atomic.Value
}

// MakeSync create a empty instance
func MakeSync() *Sync {
	return &Sync{
		value: &atomic.Value{},
	}
}

// SyncOf create a new instance
func SyncOf(value interface{}) *Sync {
	v := MakeSync()
	v.value.Store(value)
	return v
}

// Set set <value> to <v>, and returns the old value.
func (v *Sync) Set(value interface{}) (old interface{}) {
	old = v.Get()
	if reflect.TypeOf(old) != reflect.TypeOf(value) {
		*v = *SyncOf(value)
		return old
	}
	v.value.Store(value)
	return old
}

// Get returns the current value of <v>.
func (v *Sync) Get() interface{} {
	return v.value.Load()
}

// Val is alias of Get.
func (v *Sync) Val() interface{} {
	return v.Get()
}

// Interface is alias of Get.
func (v *Sync) Interface() interface{} {
	return v.Get()
}

// String returns <v> as string.
func (v *Sync) String() string {
	val := v.value.Load()
	if val == nil {
		return ""
	}
	value, ok := val.(string)
	if !ok {
		panic("v is not a type of string")
	}
	return value
}

// CString converts and returns <v> as string.
func (v *Sync) CString() string {
	val := v.value.Load()
	value, ok := val.(string)
	if ok {
		return value
	}
	return fmt.Sprintf("%v", val)
}
