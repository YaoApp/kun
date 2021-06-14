package any

import (
	"fmt"
	"reflect"

	"github.com/yaoapp/kun/maps"
)

// Map the replacement for Maps.Map
type Map struct {
	maps.MapStrAny
}

// MapOf create a new instance (the default type of map)
func MapOf(values map[string]interface{}) Map {
	return Map{MapStrAny: maps.MapStrAnyOf(values)}
}

// Any returns the value stored in the map for a key.
func (m Map) Any(key string) *Any {
	return Of(m.Get(key))
}

// Flatten The Flatten method is alias of Dot, to flatten a multi-dimensional map[string]inteface{} into a single level  map[string]inteface{}
// that uses "dot" notation to indicate depth
func (m Map) Flatten() Map {
	return m.Dot()
}

// Dot The Dot method flattens a multi-dimensional map[string]inteface{} into a single level  map[string]inteface{}
// that uses "dot" notation to indicate depth
func (m Map) Dot() Map {
	res := MapOf(map[string]interface{}{})
	m.Range(func(key string, value interface{}) bool {
		res.dotSet(key, value)
		return true
	})
	return res
}

// dotSet set the value for a key uses "dot" notation
func (m Map) dotSet(key string, value interface{}) {

	m.Set(key, value)

	reflectValue := reflect.ValueOf(value)
	reflectValue = reflect.Indirect(reflectValue)
	valueKind := reflectValue.Kind()

	if valueKind == reflect.Slice || valueKind == reflect.Array { // Slice || Array
		for i := 0; i < reflectValue.Len(); i++ {
			m.dotSet(fmt.Sprintf("%s.%d", key, i), reflectValue.Index(i).Interface())
		}

	} else if valueKind == reflect.Map { // Map
		for _, sub := range reflectValue.MapKeys() {
			m.dotSet(fmt.Sprintf("%s.%v", key, sub), reflectValue.MapIndex(sub).Interface())
		}
	}
}
