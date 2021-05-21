package maps

import "github.com/yaoapp/kun/interfaces"

// MapStrAny type of map[string}]inteface{}
type MapStrAny map[string]interface{}

// Map alias of MapStrAny
type Map = MapStrAny

// MapStr alias of MapStrAny
type MapStr = MapStrAny

// Make create a new instance (the default type of map)
func Make() MapStrAny {
	return MakeMap()
}

// MakeMap create a new instance (the default type of map)
func MakeMap() MapStrAny {
	return MapStrAny{}
}

// Of create a new instance (the default type of map)
func Of(values map[string]interface{}) MapStrAny {
	return MapOf(values)
}

// MapOf create a new instance (the default type of map)
func MapOf(values map[string]interface{}) MapStrAny {
	return MapStrAny(values)
}

// MapStrOf create a new instance (the default type of map)
func MapStrOf(values map[string]interface{}) MapStrAny {
	return MapStrAny(values)
}

// MakeMapStr create a new instance
func MakeMapStr() MapStrAny {
	return MapStrAny{}
}

// MakeStrAny create a new instance
func MakeStrAny() MapStrAny {
	return MapStrAny{}
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
func (m MapStrAny) GetOrSet(key, value string) interface{} {
	if res, has := m[key]; has {
		return res
	}
	m[key] = value
	return value
}

// GetAndDel deletes the value for a key, returning the previous value if any. The loaded result reports whether the key was present.
func (m MapStrAny) GetAndDel(key string) interface{} {
	if res, has := m[key]; has {
		delete(m, key)
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
	return keys
}

// Values returns all values of the map as a slice.
func (m MapStrAny) Values() []interface{} {
	values := []interface{}{}
	m.Range(func(key string, value interface{}) bool {
		values = append(values, value)
		return true
	})
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
