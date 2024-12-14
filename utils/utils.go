package utils

import (
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	jsoniter "github.com/json-iterator/go"
)

// DD The DD function dumps the given variables and ends execution of the script
func DD(values ...interface{}) {
	Dump(values...)
	os.Exit(0)
}

// String the values as string
func String(values ...interface{}) string {
	text := ""
	f := colorjson.NewFormatter()
	f.Indent = 4
	for _, v := range values {
		if err, ok := v.(error); ok {
			text += color.RedString(err.Error())
			continue
		}

		switch v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
			text += color.CyanString(fmt.Sprintf("%v", v))
			continue
		case string, []byte:
			text += color.GreenString(fmt.Sprintf("%s", v))
			continue
		default:
			var res interface{}
			txt, err := jsoniter.Marshal(v)
			if err != nil {
				text += color.RedString(err.Error())
				continue
			}
			jsoniter.Unmarshal(txt, &res)
			s, _ := f.Marshal(res)
			text += color.YellowString(string(s))
		}
	}
	return text
}

// Dump The Dump function dumps the given variables:
func Dump(values ...interface{}) {
	fmt.Println(String(values...))

}
