package maps

import "github.com/yaoapp/kun/interfaces"

// MapStrStr type of map[string}]string
type MapStrStr map[string]string

// MakeMapStrStr create a new instance
func MakeMapStrStr() MapStrStr {
	return MapStrStr{}
}

// MakeStrStr create a new instance
func MakeStrStr() MapStrStr {
	return MapStrStr{}
}

// Set set the value for a key
func (m MapStrStr) Set(key string, value string) {
	m[key] = value
}

// Get turns the value stored in the map for a key, or nil if no value is present.
func (m MapStrStr) Get(key string) interface{} {
	return m[key]
}

// Has return true whether value was found in the map.
func (m MapStrStr) Has(key string) bool {
	_, has := m[key]
	return has
}

// Del deletes the value for a key.
func (m MapStrStr) Del(key string) {
	delete(m, key)
}

// GetOrSet returns the existing value for the key if present. Otherwise, it stores and returns the given value.
func (m MapStrStr) GetOrSet(key, value string) string {
	if res, has := m[key]; has {
		return res
	}
	m[key] = value
	return value
}

// GetAndDel deletes the value for a key, returning the previous value if any. The loaded result reports whether the key was present.
func (m MapStrStr) GetAndDel(key string) string {
	if res, has := m[key]; has {
		delete(m, key)
		return res
	}
	return ""
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m MapStrStr) Range(cb func(key string, value string) bool) {
	for key, value := range m {
		if !cb(key, value) {
			break
		}
	}
}

// IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.
func (m MapStrStr) IsEmpty() bool {
	return len(m) == 0
}

// Merge merges hash maps
func (m MapStrStr) Merge(maps ...interfaces.MapStrStr) {
	for _, new := range maps {
		new.Range(func(key string, value string) bool {
			m.Set(key, value)
			return true
		})
	}
}
