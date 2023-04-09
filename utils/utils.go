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

// Dump The Dump function dumps the given variables:
func Dump(values ...interface{}) {

	f := colorjson.NewFormatter()
	f.Indent = 4
	for _, v := range values {

		if err, ok := v.(error); ok {
			color.Red(err.Error())
			continue
		}

		switch v.(type) {

		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
			color.Cyan(fmt.Sprintf("%v", v))
			return

		case string, []byte:
			color.Green(fmt.Sprintf("%s", v))
			return

		default:
			var res interface{}
			txt, err := jsoniter.Marshal(v)
			if err != nil {
				fmt.Printf("%#v\n%s\n", v, err)
				continue
			}
			jsoniter.Unmarshal(txt, &res)
			s, _ := f.Marshal(res)
			fmt.Printf("%s\n", s)
		}
	}
}
