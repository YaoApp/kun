package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/kun/interfaces"
)

func TestStrAnyMake(t *testing.T) {
	m1 := Map{
		"foo": "bar",
		"nested": Map{
			"foo": "bar",
		},
	}
	assert.Equal(t, "bar", m1.Get("foo"))
	assert.Equal(t, Map{"foo": "bar"}, m1.Get("nested"))

	m2 := Of(map[string]interface{}{
		"foo": "bar",
		"nested": map[string]interface{}{
			"foo": "bar",
		},
	})
	assert.Equal(t, "bar", m2.Get("foo"))
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, m2.Get("nested"))
	assert.IsType(t, MapStrAny{}, Make())
	assert.IsType(t, MapStrAny{}, MakeMap())
	assert.IsType(t, MapStrAny{}, MakeMapStr())
	assert.IsType(t, MapStrAny{}, MakeStr())
	assert.IsType(t, MapStrAny{}, MakeStrAny())
	assert.IsType(t, MapStrAny{}, MakeMapStrAny())
	assert.IsType(t, MapStrAny{}, Of(map[string]interface{}{"foo": "bar"}))
	assert.IsType(t, MapStrAny{}, MapOf(map[string]interface{}{"foo": "bar"}))
	assert.IsType(t, MapStrAny{}, MapStrOf(map[string]interface{}{"foo": "bar"}))
	assert.IsType(t, MapStrAny{}, StrOf(map[string]interface{}{"foo": "bar"}))
	assert.IsType(t, MapStrAny{}, StrAnyOf(map[string]interface{}{"foo": "bar"}))
	assert.IsType(t, MapStrAny{}, MapStrAnyOf(map[string]interface{}{"foo": "bar"}))
}

func TestStrAnySetBasic(t *testing.T) {
	basic, _, _, _, _ := prepareTestingData()
	m := MakeMapStrAny()
	for key, value := range basic {
		m.Set(key, value)
	}
	if assert.Equal(t, 16, m.Len(), "The length of map should be 16") {
		checkBaiscValues(t, m)
	}
}

func TestStrAnySetAll(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := MakeMapStrAny()
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

func TestStrAnyKeys(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	keys := Of(all).Keys()
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

func TestStrAnyKeysValues(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := Of(all)
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

func TestStrAnyRange(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := Of(all)
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

func TestStrAnyFlatten(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	flatten := Of(all).Flatten()

	if assert.Equal(t, 69, flatten.Len(), "The length of map should be 69") {
		values := flatten.Values()
		for i, key := range flatten.Keys() {
			assert.Equal(t, values[i], flatten.Get(key))
		}
	}
}

func TestStrAnyHas(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	flatten := Of(all).Flatten()
	keys := flatten.Keys()
	if assert.Equal(t, 69, len(keys), "The length of keys should be 69") {
		for _, key := range keys {
			assert.True(t, flatten.Has(key))
		}
	}
	assert.False(t, flatten.Has("not_existed_key"))
}

func TestStrAnyDel(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := Of(all)
	keys := m.Keys()
	if assert.Equal(t, 22, len(keys), "The length of keys should be 22") {
		for _, key := range keys {
			m.Del(key)
		}
	}
	assert.Equal(t, 0, m.Len())
}

func TestStrAnyGetAndDel(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := Of(all)
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

func TestStrAnyGetOrSet(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := Of(all)
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

func TestStrAnyIsEmpty(t *testing.T) {
	_, _, _, _, all := prepareTestingData()
	m := Of(all)
	assert.False(t, m.IsEmpty())

	keys := m.Keys()
	if assert.Equal(t, 22, len(keys), "The length of keys should be 22") {
		for _, key := range keys {
			m.Del(key)
		}
	}
	assert.True(t, m.IsEmpty())
}

func TestStrAnyMerge(t *testing.T) {
	basic, array, slice, _, _ := prepareTestingData()
	m := Of(basic)

	var new interfaces.MapStr = Of(array)
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

func TestStrAnyUnDot(t *testing.T) {
	m1 := Map{
		"foo":             "bar",
		"dot.hello.world": "hello world",
		"dot.bingo.world": "hello bingo",
		"nested": Map{
			"foo": "bar",
		},
	}
	res := m1.UnFlatten()
	dot, ok := res.Get("dot").(Map)
	assert.True(t, ok)
	if ok {
		assert.Equal(t, dot.Get("hello"), Map{"world": "hello world"})
		assert.Equal(t, dot.Get("bingo"), Map{"world": "hello bingo"})
	}
}

func checkArrayValues(t *testing.T, m interfaces.MapStrAny) {
	assert.Equal(t, [2]int64{64, 64}, m.Get("arrayint64"))
}

func checkSliceValues(t *testing.T, m interfaces.MapStrAny) {
	assert.Equal(t, []int64{64, 64, 64}, m.Get("sliceint64"))
}

func checkMapValues(t *testing.T, m interfaces.MapStrAny) {
	assert.Equal(t, map[int64]interface{}{64: "hello"}, m.Get("mapint64"))
}

func checkStructValues(t *testing.T, m interfaces.MapStrAny) {
	assert.Equal(t,
		struct {
			Name  string
			Value interface{}
		}{Name: "unit-test", Value: "hello"}, m.Get("struct"))
}

func checkNestedValues(t *testing.T, m interfaces.MapStrAny) {
	assert.Equal(t, map[string]interface{}{
		"basic": map[string]interface{}{
			"int64":   int64(64),
			"int32":   int32(32),
			"int16":   int16(16),
			"int8":    int8(8),
			"int":     1,
			"uint64":  uint64(64),
			"uint32":  uint32(32),
			"uint16":  uint16(16),
			"uint8":   uint8(8),
			"uint":    uint(1),
			"float64": float64(9.65),
			"float32": float32(9.65),
			"byte":    byte(55),
			"bool":    true,
			"uintptr": uintptr(19),
			"string":  "string",
		},
	}, m.Get("nested"))
}

func checkBaiscValues(t *testing.T, m interfaces.MapStrAny) {
	assert.Equal(t, int64(64), m.Get("int64"))
	assert.Equal(t, int32(32), m.Get("int32"))
	assert.Equal(t, int16(16), m.Get("int16"))
	assert.Equal(t, int8(8), m.Get("int8"))
	assert.Equal(t, int(1), m.Get("int"))
	assert.Equal(t, uint64(64), m.Get("uint64"))
	assert.Equal(t, uint32(32), m.Get("uint32"))
	assert.Equal(t, uint16(16), m.Get("uint16"))
	assert.Equal(t, uint8(8), m.Get("uint8"))
	assert.Equal(t, uint(1), m.Get("uint"))
	assert.Equal(t, float64(9.65), m.Get("float64"))
	assert.Equal(t, float32(9.65), m.Get("float32"))
	assert.Equal(t, byte(55), m.Get("byte"))
	assert.Equal(t, true, m.Get("bool"))
	assert.Equal(t, uintptr(19), m.Get("uintptr"))
	assert.Equal(t, "string", m.Get("string"))
}

// prepareTestingData prepare the data for testing
// baiscValues, arrayValues, sliceValues, mapValues, allValues
func prepareTestingData() (baiscValues, arrayValues, sliceValues, mapValues, allValues map[string]interface{}) {

	var structValue = struct {
		Name  string
		Value interface{}
	}{Name: "unit-test", Value: "hello"}

	baiscValues = map[string]interface{}{
		"int64":   int64(64),
		"int32":   int32(32),
		"int16":   int16(16),
		"int8":    int8(8),
		"int":     1,
		"uint64":  uint64(64),
		"uint32":  uint32(32),
		"uint16":  uint16(16),
		"uint8":   uint8(8),
		"uint":    uint(1),
		"float64": float64(9.65),
		"float32": float32(9.65),
		"byte":    byte(55),
		"bool":    true,
		"uintptr": uintptr(19),
		"string":  "string",
	}

	arrayValues = map[string]interface{}{
		"int64": [2]int64{64, 64},
	}

	sliceValues = map[string]interface{}{
		"int64": []int64{64, 64, 64},
	}

	mapValues = map[string]interface{}{
		"int64": map[int64]interface{}{64: "hello"},
	}

	allValues = map[string]interface{}{}
	for key, value := range baiscValues {
		allValues[key] = value
	}
	for key, value := range arrayValues {
		allValues["array"+key] = value
	}
	for key, value := range sliceValues {
		allValues["slice"+key] = value
	}
	for key, value := range mapValues {
		allValues["map"+key] = value
	}

	allValues["struct"] = structValue
	allValues["nested"] = map[string]interface{}{}
	for key, value := range baiscValues {
		allValues[key] = value
	}
	allValues["nested"].(map[string]interface{})["basic"] = map[string]interface{}{}
	for key, value := range baiscValues {
		allValues["nested"].(map[string]interface{})["basic"].(map[string]interface{})[key] = value
	}

	allValues["nested2"] = Map{
		"basic": Of(baiscValues),
	}

	return baiscValues, arrayValues, sliceValues, mapValues, allValues
}
