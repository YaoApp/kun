package maps

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/yaoapp/kun/interfaces"
)

// MapStrAny type of map[string}]inteface{}
type MapStrAny map[string]interface{}

// Map alias of MapStrAny
type Map = MapStrAny

// MapStr alias of MapStrAny
type MapStr = MapStrAny

// Make create a new instance (the default type of map)
func Make() MapStrAny {
	return MakeMapStrAny()
}

// MakeMap create a new instance
func MakeMap() MapStrAny {
	return MakeMapStrAny()
}

// MakeMapStr create a new instance
func MakeMapStr() MapStrAny {
	return MakeMapStrAny()
}

// MakeStr create a new instance
func MakeStr() MapStrAny {
	return MakeMapStrAny()
}

// MakeStrAny create a new instance
func MakeStrAny() MapStrAny {
	return MakeMapStrAny()
}

// MakeMapStrAny create a new instance
func MakeMapStrAny() MapStrAny {
	return MapStrAny{}
}

// Of create a new instance (the default type of map)
func Of(values map[string]interface{}) MapStrAny {
	return MapStrAnyOf(values)
}

// MapOf create a new instance
func MapOf(values map[string]interface{}) MapStrAny {
	return MapStrAnyOf(values)
}

// MapStrOf create a new instance
func MapStrOf(values map[string]interface{}) MapStrAny {
	return MapStrAnyOf(values)
}

// StrOf create a new instance
func StrOf(values map[string]interface{}) MapStrAny {
	return MapStrAnyOf(values)
}

// StrAnyOf create a new instance
func StrAnyOf(values map[string]interface{}) MapStrAny {
	return MapStrAnyOf(values)
}

// MapStrAnyOf create a new instance (the default type of map)
func MapStrAnyOf(values map[string]interface{}) MapStrAny {
	m := MakeMapStrAny()
	for key, value := range values {
		m.Set(key, value)
	}
	return m
}

// Flatten The Flatten method is alias of Dot, to flatten a multi-dimensional map[string]inteface{} into a single level  map[string]inteface{}
// that uses "dot" notation to indicate depth
func (m MapStrAny) Flatten() MapStrAny {
	return m.Dot()
}

// Dot The Dot method flattens a multi-dimensional map[string]inteface{} into a single level  map[string]inteface{}
// that uses "dot" notation to indicate depth
func (m MapStrAny) Dot() MapStrAny {
	res := MakeMapStrAny()
	m.Range(func(key string, value interface{}) bool {
		res.dotSet(key, value)
		return true
	})
	return res
}

// dotSet set the value for a key uses "dot" notation
func (m MapStrAny) dotSet(key string, value interface{}) {

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
	// } else if subMap, ok := value.(MapStrAny); ok { // map[string]interface{}
	// 	subMap.Range(func(sub string, val interface{}) bool {
	// 		m.dotSet(sub, val)
	// 		return true
	// 	})
	// }
}

// Set set the value for a key
func (m MapStrAny) Set(key string, value interface{}) {
	m[key] = value
}

// Get turns the value stored in the map for a key, or nil if no value is present.
func (m MapStrAny) Get(key string) interface{} {
	return m[key]
}

// Has return true whether value was found in the map.
func (m MapStrAny) Has(key string) bool {
	_, has := m[key]
	return has
}

// Del deletes the value for a key.
func (m MapStrAny) Del(key string) {
	delete(m, key)
}

// GetOrSet returns the existing value for the key if present. Otherwise, it stores and returns the given value.
func (m MapStrAny) GetOrSet(key string, value interface{}) interface{} {
	if res, has := m[key]; has {
		return res
	}
	m.Set(key, value)
	return value
}

// GetAndDel deletes the value for a key, returning the previous value if any. The loaded result reports whether the key was present.
func (m MapStrAny) GetAndDel(key string) interface{} {
	if res, has := m[key]; has {
		m.Del(key)
		return res
	}
	return nil
}

// Len returns the length of the map.
func (m MapStrAny) Len() int {
	return len(m)
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m MapStrAny) Range(cb func(key string, value interface{}) bool) {
	for key, value := range m {
		if !cb(key, value) {
			break
		}
	}
}

// Keys returns all keys of the map as a slice.
func (m MapStrAny) Keys() []string {
	keys := []string{}
	m.Range(func(key string, value interface{}) bool {
		keys = append(keys, key)
		return true
	})
	sort.Strings(keys)
	return keys
}

// Values returns all values of the map as a slice.
func (m MapStrAny) Values() []interface{} {
	values := []interface{}{}
	keys := m.Keys()
	for _, key := range keys {
		values = append(values, m.Get(key))
	}
	return values
}

// IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.
func (m MapStrAny) IsEmpty() bool {
	return len(m) == 0
}

// Merge merges hash maps
func (m MapStrAny) Merge(maps ...interfaces.MapStr) {
	for _, new := range maps {
		new.Range(func(key string, value interface{}) bool {
			m.Set(key, value)
			return true
		})
	}
}
