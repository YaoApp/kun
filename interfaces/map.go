package interfaces

// Map the map interface (map[inteface{}]inteface{})
type Map interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
	Del(key interface{})
	GetOrSet(key, value interface{}) interface{}
	GetAndDel(key interface{}) interface{}
	Range(func(key, value interface{}) bool)
	IsEmpty() bool
	Merge(maps ...Map)
}

// // MapString the map interface (map[string]inteface{})
// type MapString interface {
// 	Set(key string, value interface{})
// 	Get(key string) interface{}
// 	Del(key string)
// 	GetOrSet(key string) interface{}
// 	GetAndDel(key string) interface{}
// 	Range(func(key string, value interface{}) bool)
// 	IsEmpty() bool
// }

// // MapStringString the map interface (map[string]string)
// type MapStringString interface {
// 	Set(key string, value string)
// 	Get(key string) string
// 	Del(key string)
// 	GetOrSet(key string) string
// 	GetAndDel(key string) string
// 	Range(func(key string, value string) bool)
// 	IsEmpty() bool
// }
