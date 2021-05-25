package maps

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/kun/interfaces"
)

func TestStrAnySyncMake(t *testing.T) {

	m1 := SyncOf(map[string]interface{}{
		"foo": "bar",
		"nested": map[string]interface{}{
			"foo": "bar",
		},
	})
	assert.Equal(t, "bar", m1.Get("foo"))
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, m1.Get("nested"))

	assert.IsType(t, MapStrAnySync{}, MakeSync())
	assert.IsType(t, MapStrAnySync{}, MakeMapSync())
	assert.IsType(t, MapStrAnySync{}, MakeMapStrSync())
	assert.IsType(t, MapStrAnySync{}, MakeStrSync())
	assert.IsType(t, MapStrAnySync{}, MakeStrAnySync())
	assert.IsType(t, MapStrAnySync{}, MakeMapStrAnySync())
	assert.IsType(t, MapStrAnySync{}, SyncOf(map[string]interface{}{"foo": "bar"}))
	assert.IsType(t, MapStrAnySync{}, MapSyncOf(map[string]interface{}{"foo": "bar"}))
	assert.IsType(t, MapStrAnySync{}, MapStrSyncOf(map[string]interface{}{"foo": "bar"}))
	assert.IsType(t, MapStrAnySync{}, StrSyncOf(map[string]interface{}{"foo": "bar"}))
	assert.IsType(t, MapStrAnySync{}, StrAnySyncOf(map[string]interface{}{"foo": "bar"}))
	assert.IsType(t, MapStrAnySync{}, MapStrAnySyncOf(map[string]interface{}{"foo": "bar"}))
}

func TestStrAnySyncSetBasic(t *testing.T) {
	basic, _, _, _, _ := prepareTestingData()
	m := MakeMapStrAnySync()
	for key, value := range basic {
		m.Set(key, value)
	}
	if assert.Equal(t, 16, m.Len(), "The length of map should be 16") {
		checkBaiscValues(t, m)
	}
}

func TestStrAnySyncSetBasicConcurrent(t *testing.T) {
	basic, _, _, _, _ := prepareTestingData()
	m := MakeMapStrAnySync()
	var wg sync.WaitGroup
	for key, value := range basic {
		wg.Add(1)
		go func(k string, v interface{}) {
			defer wg.Done()
			m.Set(k, v)
		}(key, value)
	}
	wg.Wait()
	if assert.Equal(t, 16, m.Len(), "The length of map should be 16") {
		checkBaiscValues(t, m)
	}
}

func TestStrAnySyncSetAll(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := MakeMapStrAnySync()
	for key, value := range all {
		m.Set(key, value)
	}
	if assert.Equal(t, 22, m.Len(), "The length of map should be 22") {
		checkBaiscValues(t, m)
		checkArrayValues(t, m)
		checkSliceValues(t, m)
		checkMapValues(t, m)
		checkNestedValues(t, m)
		checkStructValues(t, m)
	}
}

func TestStrAnySyncSetAllConcurrent(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := MakeMapStrAnySync()
	var wg sync.WaitGroup
	for key, value := range all {
		wg.Add(1)
		go func(k string, v interface{}) {
			defer wg.Done()
			m.Set(k, v)
		}(key, value)
	}
	wg.Wait()
	if assert.Equal(t, 22, m.Len(), "The length of map should be 22") {
		checkBaiscValues(t, m)
		checkArrayValues(t, m)
		checkSliceValues(t, m)
		checkMapValues(t, m)
		checkNestedValues(t, m)
		checkStructValues(t, m)
	}
}

func TestStrAnySyncKeys(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	keys := SyncOf(all).Keys()
	assert.Equal(t, []string{
		"arrayint64",
		"bool",
		"byte",
		"float32",
		"float64",
		"int",
		"int16",
		"int32",
		"int64",
		"int8",
		"mapint64",
		"nested",
		"nested2",
		"sliceint64",
		"string",
		"struct",
		"uint",
		"uint16",
		"uint32",
		"uint64",
		"uint8",
		"uintptr"}, keys)
}

func TestStrAnySyncKeysValues(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := SyncOf(all)
	keys := m.Keys()
	values := m.Values()
	if assert.Equal(t, 22, len(keys), "The length of map should be 21") {
		for i := 0; i < 22; i++ {
			key := keys[i]
			value := values[i]
			assert.Equal(t, value, m.Get(key))
		}
	}
}

func TestStrAnySyncRange(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := SyncOf(all)
	values := []interface{}{}
	m.Range(func(key string, value interface{}) bool {
		values = append(values, value)
		return true
	})
	assert.Equal(t, m.Len(), len(values))

	values = []interface{}{}
	m.Range(func(key string, value interface{}) bool {
		values = append(values, value)
		return false
	})
	assert.Equal(t, 1, len(values))
}

func TestStrAnySyncFlatten(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	flatten := SyncOf(all).Flatten()
	if assert.Equal(t, 62, flatten.Len(), "The length of map should be 62") {
		values := flatten.Values()
		for i, key := range flatten.Keys() {
			assert.Equal(t, values[i], flatten.Get(key))
		}
	}
}

func TestStrAnySyncHas(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	flatten := SyncOf(all).Flatten()
	keys := flatten.Keys()
	if assert.Equal(t, 62, len(keys), "The length of keys should be 62") {
		for _, key := range keys {
			assert.True(t, flatten.Has(key))
		}
	}
	assert.False(t, flatten.Has("not_existed_key"))
}

func TestStrSyncAnyDel(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := SyncOf(all)
	keys := m.Keys()
	if assert.Equal(t, 22, len(keys), "The length of keys should be 22") {
		for _, key := range keys {
			m.Del(key)
		}
	}
	assert.Equal(t, 0, m.Len())
}

func TestStrAnySyncGetAndDel(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := SyncOf(all)
	keys := m.Keys()
	valuesBefore := m.Values()
	valuesAfter := []interface{}{}
	if assert.Equal(t, 22, len(keys), "The length of keys should be 22") {
		for _, key := range keys {
			valuesAfter = append(valuesAfter, m.GetAndDel(key))
		}
	}
	assert.Equal(t, 0, m.Len())
	assert.Equal(t, 0, len(m.Values()))
	assert.Equal(t, valuesBefore, valuesAfter)
	assert.Nil(t, m.GetAndDel("not-exists"))
}

func TestStrAnySyncGetOrSet(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := SyncOf(all)
	keys := m.Keys()
	valuesBefore := m.Values()
	valuesAfter := []interface{}{}
	if assert.Equal(t, 22, len(keys), "The length of keys should be 22") {
		for _, key := range keys {
			valuesAfter = append(valuesAfter, m.GetOrSet(key, "getorset"))
		}
	}
	assert.Equal(t, 22, m.Len())
	assert.Equal(t, valuesBefore, valuesAfter)

	value := m.GetOrSet("new-key", "getorset")
	assert.Equal(t, 23, m.Len())
	assert.Equal(t, value, "getorset")
	assert.Equal(t, value, m.Get("new-key"))
}

func TestStrAnySyncIsEmpty(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := SyncOf(all)
	assert.False(t, m.IsEmpty())

	keys := m.Keys()
	if assert.Equal(t, 22, len(keys), "The length of keys should be 22") {
		for _, key := range keys {
			m.Del(key)
		}
	}
	assert.True(t, m.IsEmpty())
}

func TestStrAnySyncMerge(t *testing.T) {
	basic, array, slice, _, _ := prepareTestingData()
	m := SyncOf(basic)

	var new interfaces.MapStr = SyncOf(array)
	m.Merge(new)
	if assert.Equal(t, 16, m.Len(), "The length of keys should be 22") {
		assert.Equal(t, int64(64), m.Dot().Get("int64.0"))
		assert.Equal(t, int64(64), m.Dot().Get("int64.1"))
	}

	var new2 interfaces.MapStr = Of(slice)
	m.Merge(new, new2)
	if assert.Equal(t, 16, m.Len(), "The length of keys should be 22") {
		assert.Equal(t, int64(64), m.Dot().Get("int64.0"))
		assert.Equal(t, int64(64), m.Dot().Get("int64.1"))
		assert.Equal(t, int64(64), m.Dot().Get("int64.2"))
	}
}
