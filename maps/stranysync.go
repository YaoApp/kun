package maps

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"sync"

	"github.com/yaoapp/kun/interfaces"
)

// MapSync alias of MapStrAnySync
type MapSync = MapStrAnySync

// MapStrSync alias of MapStrAnySync
type MapStrSync = MapStrAnySync

// StrSync alias of MapStrAnySync
type StrSync = MapStrAnySync

// StrAnySync alias of MapStrAnySync
type StrAnySync = MapStrAnySync

// MapStrAnySync type of sync.Map
type MapStrAnySync struct {
	*sync.Map
}

// MakeSync create a new instance
func MakeSync() MapStrAnySync {
	return MakeMapStrAnySync()
}

// MakeMapSync create a new instance
func MakeMapSync() MapStrAnySync {
	return MakeMapStrAnySync()
}

// MakeMapStrSync create a new instance
func MakeMapStrSync() MapStrAnySync {
	return MakeMapStrAnySync()
}

// MakeStrSync create a new instance
func MakeStrSync() MapStrAnySync {
	return MakeMapStrAnySync()
}

// MakeStrAnySync create a new instance
func MakeStrAnySync() MapStrAnySync {
	return MakeMapStrAnySync()
}

// MakeMapStrAnySync create a new instance
func MakeMapStrAnySync() MapStrAnySync {
	return MapStrAnySync{
		Map: &sync.Map{},
	}
}

// SyncOf create a new instance
func SyncOf(data map[string]interface{}) MapStrAnySync {
	return MapStrAnySyncOf(data)
}

// MapSyncOf create a new instance
func MapSyncOf(data map[string]interface{}) MapStrAnySync {
	return MapStrAnySyncOf(data)
}

// MapStrSyncOf create a new instance
func MapStrSyncOf(data map[string]interface{}) MapStrAnySync {
	return MapStrAnySyncOf(data)
}

// StrSyncOf create a new instance
func StrSyncOf(data map[string]interface{}) MapStrAnySync {
	return MapStrAnySyncOf(data)
}

// StrAnySyncOf create a new instance
func StrAnySyncOf(data map[string]interface{}) MapStrAnySync {
	return MapStrAnySyncOf(data)
}

// MapStrAnySyncOf create a new instance
func MapStrAnySyncOf(data map[string]interface{}) MapStrAnySync {
	m := MakeMapStrAnySync()
	for key, value := range data {
		m.Set(key, value)
	}
	return m
}

// Flatten The Flatten method is alias of Dot, to flatten a multi-dimensional map[string]inteface{} into a single level  map[string]inteface{}
// that uses "dot" notation to indicate depth
func (m MapStrAnySync) Flatten() MapStrAnySync {
	return m.Dot()
}

// Dot The Dot method flattens a multi-dimensional map[string]inteface{} into a single level  map[string]inteface{}
// that uses "dot" notation to indicate depth
func (m MapStrAnySync) Dot() MapStrAnySync {
	res := MakeMapStrAnySync()
	m.Range(func(key string, value interface{}) bool {
		res.dotSet(key, value)
		return true
	})
	return res
}

// dotSet set the value for a key uses "dot" notation
func (m MapStrAnySync) dotSet(key string, value interface{}) {

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

// Set set the value for a key
func (m MapStrAnySync) Set(key string, value interface{}) {
	m.Store(key, value)
}

// Get get the value of the given key
func (m MapStrAnySync) Get(key string) interface{} {
	value, has := m.Load(key)
	if has {
		return value
	}
	return nil
}

// Has return true whether value was found in the map.
func (m MapStrAnySync) Has(key string) bool {
	_, has := m.Map.Load(key)
	return has
}

// Del deletes the value for a key.
func (m MapStrAnySync) Del(key string) {
	m.Delete(key)
}

// GetOrSet returns the existing value for the key if present. Otherwise, it stores and returns the given value.
func (m MapStrAnySync) GetOrSet(key string, value interface{}) interface{} {
	value, load := m.LoadOrStore(key, value)
	if load {
		return value
	}

	return value
}

// GetAndDel deletes the value for a key, returning the previous value if any. The loaded result reports whether the key was present.
func (m MapStrAnySync) GetAndDel(key string) interface{} {
	value := m.Get(key)
	m.Del(key)
	return value
}

// Len returns the length of the map.
func (m MapStrAnySync) Len() int {
	length := 0
	m.Map.Range(func(key, value interface{}) bool {
		length++
		return true
	})
	return length
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m MapStrAnySync) Range(cb func(key string, value interface{}) bool) {
	m.Map.Range(func(key, value interface{}) bool {
		keyStr := key.(string)
		return cb(keyStr, value)
	})
}

// Keys returns all keys of the map as a slice.
func (m MapStrAnySync) Keys() []string {
	keys := []string{}
	m.Range(func(key string, value interface{}) bool {
		keys = append(keys, key)
		return true
	})
	sort.Strings(keys)
	return keys
}

// Values returns all values of the map as a slice.
func (m MapStrAnySync) Values() []interface{} {
	values := []interface{}{}
	keys := m.Keys()
	for _, key := range keys {
		values = append(values, m.Get(key))
	}
	return values
}

// IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.
func (m MapStrAnySync) IsEmpty() bool {
	empty := true
	m.Range(func(key string, value interface{}) bool {
		empty = false
		return false
	})
	return empty
}

// Merge merges hash maps
func (m MapStrAnySync) Merge(maps ...interfaces.MapStrAny) {
	for _, new := range maps {
		new.Range(func(key string, value interface{}) bool {
			m.Set(key, value)
			return true
		})
	}
}

// MarshalJSON for json marshalJSON
func (m MapStrAnySync) MarshalJSON() ([]byte, error) {
	res := map[string]interface{}{}
	m.Range(func(key string, value interface{}) bool {
		res[key] = value
		return true
	})
	return json.Marshal(res)
}
