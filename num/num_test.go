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
	assert.Equal(t, 0.618, Of(0.618).Float())
	assert.Equal(t, float64(float32(0.618)), Of(float32(0.618)).Float())
	assert.Equal(t, 0.618, Of("0.618").Float())
	assert.Panics(t, func() {
		fmt.Println(Of("error").Float())
	})
}

func TestFloat32(t *testing.T) {
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
