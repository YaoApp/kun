package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/kun/maps"
)

func TestBindMapStrAny(t *testing.T) {
	data := maps.MapStrAny{
		"foo": "bar",
		"num": 101,
		"extra": maps.MapStrAny{
			"sex":    "男",
			"weight": 198,
		},
	}

	content := Of("{{foo}} #{{num}} Sex:{{extra.sex}} Weight:{{extra.weight}}kg  {{extra}}").Bind(data)
	assert.Equal(t, content, "bar #101 Sex:男 Weight:198kg  map[sex:男 weight:198]")
}
