package any

import (
	"fmt"
	"reflect"

	"github.com/yaoapp/kun/maps"
	"github.com/yaoapp/kun/share"
)

// Map the replacement for Maps.Map
type Map struct {
	maps.MapStrAny
}

// MakeMap create a new  map instance
func MakeMap() Map {
	return Map{MapStrAny: maps.MapStrAnyOf(map[string]interface{}{})}
}

// MapOf create a new map instance
func MapOf(values interface{}) Map {
	if values == nil {
		return MakeMap()
	}
	switch values.(type) {
	case map[string]interface{}:
		return Map{MapStrAny: maps.MapStrAnyOf(values.(map[string]interface{}))}
	}

	// converts to map
	reflectValues := reflect.ValueOf(values)
	reflectValues = reflect.Indirect(reflectValues)
	if reflectValues.Kind() == reflect.Map {
		valuesMap := map[string]interface{}{}
		for _, key := range reflectValues.MapKeys() {
			k := fmt.Sprintf("%v", key)
			valuesMap[k] = reflectValues.MapIndex(key).Interface()
		}
		return Map{MapStrAny: maps.MapStrAnyOf(valuesMap)}
	} else if reflectValues.Kind() == reflect.Struct {
		valuesMap := map[string]interface{}{}
		typeOfS := reflectValues.Type()
		for i := 0; i < reflectValues.NumField(); i++ {
			name := share.GetTagName(typeOfS.Field(i), "json")
			valuesMap[name] = reflectValues.Field(i).Interface()
		}
		return Map{MapStrAny: maps.MapStrAnyOf(valuesMap)}
	}

	panic(fmt.Sprintf("v is %s not a type of map", reflectValues.Kind()))
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
	res := MakeMap()
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
