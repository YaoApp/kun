package maps

import (
	"sync"

	"github.com/yaoapp/kun/interfaces"
)

// MapStrAnySync type of sync.Map
type MapStrAnySync struct {
	*sync.Map
}

// MapSync alias of MapStrAnySync
type MapSync = MapStrAnySync

// MapStrSync alias of MapStrAnySync
type MapStrSync = MapStrAnySync

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

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m MapStrAnySync) Range(cb func(key string, value interface{}) bool) {
	m.Map.Range(func(key, value interface{}) bool {
		keyStr := key.(string)
		return cb(keyStr, value)
	})
}

//IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.
func (m MapStrAnySync) IsEmpty() bool {
	has := false
	m.Range(func(key string, value interface{}) bool {
		has = true
		return false
	})
	return has
}

// Merge merges hash maps
func (m MapStrAnySync) Merge(maps ...interfaces.MapStr) {
	for _, new := range maps {
		new.Range(func(key string, value interface{}) bool {
			m.Set(key, value)
			return true
		})
	}
}
