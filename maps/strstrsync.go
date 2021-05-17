package maps

import (
	"sync"

	"github.com/yaoapp/kun/interfaces"
)

// MapStrStrSync type of sync.Map
type MapStrStrSync struct {
	*sync.Map
}

// MakeStrStrSync create a new instance
func MakeStrStrSync() MapStrStrSync {
	return MapStrStrSync{
		Map: &sync.Map{},
	}
}

// Set set the value for a key
func (m MapStrStrSync) Set(key string, value string) {
	m.Store(key, value)
}

// Get get the value of the given key
func (m MapStrStrSync) Get(key string) string {
	value, has := m.Load(key)
	if has {
		return value.(string)
	}
	return ""
}

// Del deletes the value for a key.
func (m MapStrStrSync) Del(key string) {
	m.Delete(key)
}

// GetOrSet returns the existing value for the key if present. Otherwise, it stores and returns the given value.
func (m MapStrStrSync) GetOrSet(key string, value string) string {
	res, load := m.LoadOrStore(key, value)
	if load {
		return res.(string)
	}

	return value
}

// GetAndDel deletes the value for a key, returning the previous value if any. The loaded result reports whether the key was present.
func (m MapStrStrSync) GetAndDel(key string) string {
	value := m.Get(key)
	m.Del(key)
	return value
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m MapStrStrSync) Range(cb func(key string, value string) bool) {
	m.Map.Range(func(key, value interface{}) bool {
		keyStr := key.(string)
		valueStr := value.(string)
		return cb(keyStr, valueStr)
	})
}

//IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.
func (m MapStrStrSync) IsEmpty() bool {
	has := false
	m.Range(func(key string, value string) bool {
		has = true
		return false
	})
	return has
}

// Merge merges hash maps
func (m MapStrStrSync) Merge(maps ...interfaces.MapStrStr) {
	for _, new := range maps {
		new.Range(func(key string, value string) bool {
			m.Set(key, value)
			return true
		})
	}
}
