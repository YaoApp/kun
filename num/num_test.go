package num

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	v := Make()
	assert.Equal(t, 0.0, v.Float())
}
func TestOf(t *testing.T) {
	v := Of(0.618)
	assert.Equal(t, 0.618, v.Float())
}

func TestSet(t *testing.T) {
	v := Of(0.618)
	assert.Equal(t, 0.618, v.Float())

	v.Set(0.852)
	assert.Equal(t, 0.852, v.Float())
}

func TestFloat(t *testing.T) {
	assert.Equal(t, 0.0, Of(nil).Float())
	assert.Equal(t, 0.618, Of(0.618).Float())
	assert.Equal(t, float64(float32(0.618)), Of(float32(0.618)).Float())
	assert.Equal(t, 0.618, Of("0.618").Float())
	assert.Panics(t, func() {
		fmt.Println(Of("error").Float())
	})
}

func TestFloat32(t *testing.T) {
	assert.Equal(t, float32(0.0), Of(nil).Float32())
	assert.Equal(t, float32(0.618), Of(0.618).Float32())
	assert.Equal(t, float32(0.618), Of(float32(0.618)).Float32())
	assert.Equal(t, float32(0.618), Of("0.618").Float32())
	assert.Panics(t, func() {
		fmt.Println(Of("error").Float32())
	})
}

func TestToFixed(t *testing.T) {
	assert.Equal(t, "0.62", Of(0.618).ToFixed(2))
	assert.Equal(t, "0.61", Of(0.614).ToFixed(2))
	assert.Equal(t, "0", Of(0.4).ToFixed(0))
	assert.Equal(t, "1", Of(0.51).ToFixed(0))
}

func TestInt(t *testing.T) {
	assert.Equal(t, 0, Of(nil).Int())
	assert.Equal(t, 10, Of(10).Int())
	assert.Equal(t, 0, Of(0.312).Int())
	assert.Equal(t, 1, Of(0.618).Int())
	assert.Equal(t, 1, Of(float32(0.618)).Int())
	assert.Equal(t, 10, Of("10").Int())
	assert.Equal(t, 0, Of("0.312").Int())
	assert.Equal(t, 1, Of("0.618").Int())
	assert.Equal(t, 1, Of("1.24").Int())
	assert.Panics(t, func() {
		fmt.Println(Of("error").Int())
	})
}

func TestInt8(t *testing.T) {
	assert.Equal(t, int8(1), Of(int8(1)).Int8())
	assert.Equal(t, int8(1), Of(1).Int8())
}
func TestInt16(t *testing.T) {
	assert.Equal(t, int16(1), Of(int16(1)).Int16())
	assert.Equal(t, int16(1), Of(1).Int16())
}

func TestInt32(t *testing.T) {
	assert.Equal(t, int32(1), Of(int32(1)).Int32())
	assert.Equal(t, int32(1), Of(1).Int32())
}

func TestInt64(t *testing.T) {
	assert.Equal(t, int64(1), Of(int64(1)).Int64())
	assert.Equal(t, int64(1), Of(1).Int64())
}

func TestUint(t *testing.T) {
	assert.Equal(t, uint(1), Of(uint(1)).Uint())
	assert.Equal(t, uint(1), Of(1).Uint())
}

func TestUint8(t *testing.T) {
	assert.Equal(t, uint8(1), Of(uint8(1)).Uint8())
	assert.Equal(t, uint8(1), Of(1).Uint8())
}

func TestUint16(t *testing.T) {
	assert.Equal(t, uint16(1), Of(uint16(1)).Uint16())
	assert.Equal(t, uint16(1), Of(1).Uint16())
}

func TestUint32(t *testing.T) {
	assert.Equal(t, uint32(1), Of(uint32(1)).Uint32())
	assert.Equal(t, uint32(1), Of(1).Uint32())
}

func TestUint64(t *testing.T) {
	assert.Equal(t, uint64(1), Of(uint64(1)).Uint64())
	assert.Equal(t, uint64(1), Of(1).Uint64())
}

func TestUintptr(t *testing.T) {
	assert.Equal(t, uintptr(1), Of(uintptr(1)).Uintptr())
	assert.Equal(t, uintptr(1), Of(1).Uintptr())
}

func TestIsInt(t *testing.T) {
	assert.Equal(t, true, Of(1).IsInt())
	assert.Equal(t, true, Of(int8(1)).IsInt())
	assert.Equal(t, false, Of(true).IsInt())
}

func TestIsFloat(t *testing.T) {
	assert.Equal(t, true, Of(1.618).IsFloat())
	assert.Equal(t, true, Of(float32(1.382)).IsFloat())
	assert.Equal(t, false, Of(5).IsFloat())
}

func TestIsSet(t *testing.T) {
	assert.Equal(t, true, Of(1).IsSet())
	assert.Equal(t, false, Of(nil).IsSet())
}

func TestIsNil(t *testing.T) {
	assert.Equal(t, false, Of(1).IsNil())
	assert.Equal(t, true, Of(nil).IsNil())
}
