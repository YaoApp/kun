package maps

import (
	"sync"

	"github.com/yaoapp/kun/interfaces"
)

// AnyAnySync type of sync.Map
type AnyAnySync struct {
	*sync.Map
}

// MakeMapSync create a new instance
func MakeMapSync() AnyAnySync {
	return AnyAnySync{
		Map: &sync.Map{},
	}
}

// MakeAnyAnySync create a new instance
func MakeAnyAnySync() AnyAnySync {
	return AnyAnySync{
		Map: &sync.Map{},
	}
}

// Set set the value for a key
func (m AnyAnySync) Set(key, value interface{}) {
	m.Store(key, value)
}

// Get get the value of the given key
func (m AnyAnySync) Get(key interface{}) interface{} {
	value, has := m.Load(key)
	if has {
		return value
	}
	return nil
}

// Del deletes the value for a key.
func (m AnyAnySync) Del(key interface{}) {
	m.Delete(key)
}

// GetOrSet returns the existing value for the key if present. Otherwise, it stores and returns the given value.
func (m AnyAnySync) GetOrSet(key, value interface{}) interface{} {
	value, load := m.LoadOrStore(key, value)
	if load {
		return value
	}

	return value
}

// GetAndDel deletes the value for a key, returning the previous value if any. The loaded result reports whether the key was present.
func (m AnyAnySync) GetAndDel(key interface{}) interface{} {
	value := m.Get(key)
	m.Del(key)
	return value
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m AnyAnySync) Range(cb func(key, value interface{}) bool) {
	m.Map.Range(cb)
}

//IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.
func (m AnyAnySync) IsEmpty() bool {
	has := false
	m.Range(func(key, value interface{}) bool {
		has = true
		return false
	})
	return has
}

// Merge merges hash maps
func (m AnyAnySync) Merge(maps ...interfaces.Map) {
	for _, new := range maps {
		new.Range(func(key, value interface{}) bool {
			m.Set(key, value)
			return true
		})
	}
}
