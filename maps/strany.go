package maps

import "github.com/yaoapp/kun/interfaces"

// MapStrAny type of map[string}]inteface{}
type MapStrAny map[string]interface{}

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

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m MapStrAny) Range(cb func(key string, value interface{}) bool) {
	for key, value := range m {
		if !cb(key, value) {
			break
		}
	}
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
