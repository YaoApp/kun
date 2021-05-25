package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
