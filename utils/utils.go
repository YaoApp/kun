package utils

import (
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
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
		var res interface{}
		if err, ok := v.(error); ok {
			fmt.Printf("%s\n", err.Error())
			continue
		}

		// to Map
		if value, ok := v.(interface{ Map() map[string]interface{} }); ok {
			v = value.Map()
		}

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
