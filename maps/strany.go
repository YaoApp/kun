package maps

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/yaoapp/kun/interfaces"
	"github.com/yaoapp/kun/share"
)

// Map alias of MapStrAny
type Map = MapStrAny

// Str alias of MapStrAny
type Str = MapStrAny

// StrAny alias of MapStrAny
type StrAny = MapStrAny

// MapStr alias of MapStrAny
type MapStr = MapStrAny

// MapStrAny type of map[string}]inteface{}
type MapStrAny map[string]interface{}

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

// UnFlatten The UnFlatten method unflatten a single level map[string]inteface{} into  multi-dimensional  map[string]inteface{}
// that uses "dot" notation to indicate depth
func (m MapStrAny) UnFlatten() MapStrAny {
	return m.UnDot()
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

	reflectValue := reflect.ValueOf(value)
	reflectValue = reflect.Indirect(reflectValue)
	valueKind := reflectValue.Kind()
	_, bytes := value.([]byte)

	if (valueKind == reflect.Slice || valueKind == reflect.Array) && !bytes { // Slice || Array
		for i := 0; i < reflectValue.Len(); i++ {
			m.dotSet(fmt.Sprintf("%s.%d", key, i), reflectValue.Index(i).Interface())  // xxx.0
			m.dotSet(fmt.Sprintf("%s[%d]", key, i), reflectValue.Index(i).Interface()) // xxx[0]
		}
	} else if valueKind == reflect.Map { // Map
		for _, sub := range reflectValue.MapKeys() {
			m.dotSet(fmt.Sprintf("%s.%v", key, sub), reflectValue.MapIndex(sub).Interface())
		}

	} else if valueKind == reflect.Struct { // Struct

		if toMap := reflect.ValueOf(value).MethodByName("ToMap"); toMap.IsValid() {
			args := []reflect.Value{}
			values := toMap.Call(args)
			if len(values) == 1 {
				v, ok := values[0].Interface().(map[string]interface{})
				if ok {
					m.dotSet(key, v)
				}
			}
			return
		}

		// auto struct
		typeOfS := reflectValue.Type()
		for i := 0; i < reflectValue.NumField(); i++ {
			sub := share.GetTagName(typeOfS.Field(i), "json")
			if reflectValue.Field(i).CanInterface() {
				v := reflectValue.Field(i).Interface()
				m.dotSet(fmt.Sprintf("%s.%v", key, sub), v)
			}
		}
	}

	m.Set(key, value)
}

// UnDot The UnDot method unflatten a single level map[string]inteface{} into  multi-dimensional  map[string]inteface{}
// that uses "dot" notation to indicate depth
func (m MapStrAny) UnDot() MapStrAny {
	res := MakeMapStrAny()
	m.Range(func(key string, value interface{}) bool {
		res.SetUnDot(key, value)
		return true
	})
	return res
}

// SetUnDot set the value for a key uses "dot" notation
func (m MapStrAny) SetUnDot(key string, value interface{}) {
	if !strings.Contains(key, ".") {
		m.Set(key, value)
		return
	}

	keys := strings.Split(key, ".")
	tail := strings.Join(keys[1:], ".")
	v, ok := m.Get(keys[0]).(MapStrAny)
	if !ok {
		v = MapStrAny{}
	}
	v[tail] = value

	v.Range(func(key string, value interface{}) bool {
		v.SetUnDot(key, value)
		return true
	})

	m.Set(keys[0], v)
	m.Del(key)
}

// Set set the value for a key
func (m MapStrAny) Set(key string, value interface{}) {
	m[key] = value
}

// Get returns the value stored in the map for a key, or nil if no value is present.
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
func (m MapStrAny) Merge(maps ...interfaces.MapStrAny) {
	for _, new := range maps {
		new.Range(func(key string, value interface{}) bool {
			m.Set(key, value)
			return true
		})
	}
}
