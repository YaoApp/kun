package grpc

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/yaoapp/kun/any"
	"github.com/yaoapp/kun/maps"
)

// Bind bind struct
func (res Response) Bind(v interface{}) error {
	return jsoniter.Unmarshal(res.Bytes, &v)
}

// MustBind bind struct
func (res Response) MustBind(v interface{}) {
	err := res.Bind(&v)
	if err != nil {
		panic(err)
	}
}

// Interface bind struct
func (res Response) Interface() (interface{}, error) {
	var v interface{}
	err := jsoniter.Unmarshal(res.Bytes, &v)
	return v, err
}

// MustInterface bind struct
func (res Response) MustInterface() interface{} {
	v, err := res.Interface()
	if err != nil {
		panic(err)
	}
	return v
}

// Map cast to map
func (res Response) Map() (maps.MapStrAny, error) {
	v := maps.Map{}
	err := jsoniter.Unmarshal(res.Bytes, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// MustMap cast to map
func (res Response) MustMap() maps.MapStrAny {
	v, err := res.Map()
	if err != nil {
		panic(err)
	}
	return v
}

// Array cast to array | slice
func (res Response) Array() ([]interface{}, error) {
	v := []interface{}{}
	err := jsoniter.Unmarshal(res.Bytes, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// MustArray cast to array | slice
func (res Response) MustArray() []interface{} {
	v, err := res.Array()
	if err != nil {
		panic(err)
	}
	return v
}

// Value get the response value
func (res Response) Value() (interface{}, error) {
	switch res.Type {
	case "interface":
		return res.Interface()
	case "string":
		return string(res.Bytes), nil
	case "integer", "int":
		return any.Of(string(res.Bytes)).CInt(), nil
	case "float", "double":
		return any.Of(string(res.Bytes)).CFloat(), nil
	case "map":
		return res.Map()
	case "array", "slice":
		return res.Array()
	default:
		return res.Bytes, nil
	}
}

// MustValue get the response value
func (res Response) MustValue() interface{} {
	v, err := res.Value()
	if err != nil {
		panic(err)
	}
	return v
}
