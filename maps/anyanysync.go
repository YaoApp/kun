package maps

import (
	"sync"

	"github.com/yaoapp/kun/interfaces"
)

// MapAnyAnySync type of sync.Map
type MapAnyAnySync struct {
	*sync.Map
}

// MakeAnyAnySync create a new instance
func MakeAnyAnySync() MapAnyAnySync {
	return MapAnyAnySync{
		Map: &sync.Map{},
	}
}

// Set set the value for a key
func (m MapAnyAnySync) Set(key, value interface{}) {
	m.Store(key, value)
}

// Get get the value of the given key
func (m MapAnyAnySync) Get(key interface{}) interface{} {
	value, has := m.Load(key)
	if has {
		return value
	}
	return nil
}

// Has return true whether value was found in the map.
func (m MapAnyAnySync) Has(key interface{}) bool {
	_, has := m.Map.Load(key)
	return has
}

// Del deletes the value for a key.
func (m MapAnyAnySync) Del(key interface{}) {
	m.Delete(key)
}

// GetOrSet returns the existing value for the key if present. Otherwise, it stores and returns the given value.
func (m MapAnyAnySync) GetOrSet(key, value interface{}) interface{} {
	value, load := m.LoadOrStore(key, value)
	if load {
		return value
	}

	return value
}

// GetAndDel deletes the value for a key, returning the previous value if any. The loaded result reports whether the key was present.
func (m MapAnyAnySync) GetAndDel(key interface{}) interface{} {
	value := m.Get(key)
	m.Del(key)
	return value
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m MapAnyAnySync) Range(cb func(key, value interface{}) bool) {
	m.Map.Range(cb)
}

//IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.
func (m MapAnyAnySync) IsEmpty() bool {
	has := false
	m.Range(func(key, value interface{}) bool {
		has = true
		return false
	})
	return has
}

// Merge merges hash maps
func (m MapAnyAnySync) Merge(maps ...interfaces.MapAnyAny) {
	for _, new := range maps {
		new.Range(func(key, value interface{}) bool {
			m.Set(key, value)
			return true
		})
	}
}
