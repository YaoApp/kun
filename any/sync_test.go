package any

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncMake(t *testing.T) {
	v := MakeSync()
	assert.Empty(t, v.String())
}

func TestSyncOf(t *testing.T) {
	v := SyncOf("hello")
	assert.Equal(t, "hello", v.String())
}

func TestSyncSet(t *testing.T) {
	v := SyncOf("hello")
	assert.Equal(t, "hello", v.CString())

	v.Set(1)
	assert.Equal(t, 1, v.Get())
	assert.Equal(t, "1", v.CString())
}

func TestSyncGet(t *testing.T) {

	v := SyncOf("hello")
	assert.Equal(t, "hello", v.Get())
	assert.Equal(t, "hello", v.Val())
	assert.Equal(t, "hello", v.Interface())

	v.Set(1)
	assert.Equal(t, 1, v.Get())
	assert.Equal(t, 1, v.Val())
	assert.Equal(t, 1, v.Interface())
}

func TestSyncString(t *testing.T) {
	v := SyncOf("hello")
	assert.Equal(t, "hello", v.String())
	v.Set(1)
	assert.Panics(t, func() {
		fmt.Printf(v.String())
	})
}

func TestSyncCString(t *testing.T) {
	v := SyncOf("hello")
	assert.Equal(t, "hello", v.CString())
	v.Set(1)
	assert.Equal(t, "1", v.CString())
}
