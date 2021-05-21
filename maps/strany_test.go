package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrAnySetBasic(t *testing.T) {
	basic, _, _, _, _ := prepareTestingData()
	m := MakeStrAny()
	for key, value := range basic {
		m.Set(key, value)
	}
	if assert.Equal(t, 16, m.Len(), "The length of map should be 16") {
		checkBaiscValues(t, m)
	}
}

func checkBaiscValues(t *testing.T, m MapStrAny) {
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
func prepareTestingData() (map[string]interface{}, map[string]interface{}, map[string]interface{}, map[string]interface{}, map[string]interface{}) {

	var structValue = struct {
		Name  string
		Value interface{}
	}{Name: "unit-test", Value: "hello"}

	var baiscValues = map[string]interface{}{
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

	var arrayValues = map[string]interface{}{
		"int64": [2]int64{64, 64},
	}

	var sliceValues = map[string]interface{}{
		"int64": []int64{64, 64, 64},
	}

	var mapValues = map[string]interface{}{
		"int64": map[int64]interface{}{64: "hello"},
	}

	var fullValues = map[string]interface{}{}
	for key, value := range baiscValues {
		fullValues[key] = value
	}
	for key, value := range arrayValues {
		fullValues["array"+key] = value
	}
	for key, value := range sliceValues {
		fullValues["slice"+key] = value
	}
	for key, value := range mapValues {
		fullValues["map"+key] = value
	}

	fullValues["struct"] = structValue
	fullValues["nested"] = map[string]interface{}{}
	for key, value := range baiscValues {
		fullValues[key] = value
	}
	fullValues["nested"].(map[string]interface{})["basic"] = map[string]interface{}{}
	for key, value := range baiscValues {
		fullValues["nested"].(map[string]interface{})["basic"].(map[string]interface{})[key] = value
	}

	return baiscValues, arrayValues, sliceValues, mapValues, fullValues
}
