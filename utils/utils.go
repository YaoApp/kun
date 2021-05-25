package utils

import (
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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
		txt, err := json.Marshal(v)
		if err != nil {
			fmt.Printf("%#v\n%s\n", v, err)
			continue
		}
		json.Unmarshal(txt, &res)
		s, _ := f.Marshal(res)
		fmt.Printf("%s\n", s)
	}
}
