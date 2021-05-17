package maps

import "github.com/yaoapp/kun/interfaces"

// AnyAny type of map[interface{}]inteface{}
type AnyAny map[interface{}]interface{}

// MakeMap create a new instance
func MakeMap() AnyAny {
	return AnyAny{}
}

// MakeAnyAny create a new instance
func MakeAnyAny() AnyAny {
	return AnyAny{}
}

// Set set the value for a key
func (m AnyAny) Set(key, value interface{}) {
	m[key] = value
}

// Get turns the value stored in the map for a key, or nil if no value is present.
func (m AnyAny) Get(key interface{}) interface{} {
	return m[key]
}

// Has return true whether value was found in the map.
func (m AnyAny) Has(key interface{}) bool {
	_, has := m[key]
	return has
}

// Del deletes the value for a key.
func (m AnyAny) Del(key interface{}) {
	delete(m, key)
}

// GetOrSet returns the existing value for the key if present. Otherwise, it stores and returns the given value.
func (m AnyAny) GetOrSet(key, value interface{}) interface{} {
	if res, has := m[key]; has {
		return res
	}
	m[key] = value
	return value
}

// GetAndDel deletes the value for a key, returning the previous value if any. The loaded result reports whether the key was present.
func (m AnyAny) GetAndDel(key interface{}) interface{} {
	if res, has := m[key]; has {
		delete(m, key)
		return res
	}
	return nil
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m AnyAny) Range(cb func(key, value interface{}) bool) {
	for key, value := range m {
		if !cb(key, value) {
			break
		}
	}
}

// IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.
func (m AnyAny) IsEmpty() bool {
	return len(m) == 0
}

// Merge merges hash maps
func (m AnyAny) Merge(maps ...interfaces.Map) {
	for _, new := range maps {
		new.Range(func(key, value interface{}) bool {
			m.Set(key, value)
			return true
		})
	}
}
