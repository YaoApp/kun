package interfaces

// Map the map interface (map[string]inteface{})
type Map = MapStrAny

// MapStr the map interface (map[string]inteface{})
type MapStr = MapStrAny

// MapStrAny the map interface (map[string]inteface{})
type MapStrAny interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Del(key string)
	GetOrSet(key string, value interface{}) interface{}
	GetAndDel(key string) interface{}
	Range(func(key string, value interface{}) bool)
	Has(key string) bool
	Len() int
	Keys() []string
	Values() []interface{}
	IsEmpty() bool
	Merge(maps ...MapStr)
}

// MapAny the map interface (map[inteface{}]inteface{})
type MapAny = MapAnyAny

// MapAnyAny the map interface (map[inteface{}]inteface{})
type MapAnyAny interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
	Del(key interface{})
	GetOrSet(key, value interface{}) interface{}
	GetAndDel(key interface{}) interface{}
	Range(func(key, value interface{}) bool)
	IsEmpty() bool
	Merge(maps ...MapStrStr)
}

// MapStrStr the map interface (map[string]string)
type MapStrStr interface {
	Set(key string, value string)
	Get(key string) string
	Del(key string)
	GetOrSet(key string, value string) string
	GetAndDel(key string) string
	Range(func(key string, value string) bool)
	IsEmpty() bool
	Merge(maps ...MapStrStr)
}
