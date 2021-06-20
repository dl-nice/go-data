package encoding

import (
	"encoding/json"
	"reflect"
)

var _ Codec = (*jsonCodec)(nil)

const jsonType = "json"

type jsonCodec struct{}

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonCodec) Unmarshal(d []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() && rv.CanInterface() {
			rv.Set(reflect.New(rv.Type().Elem()))
			v = rv.Interface()
		}
		rv = rv.Elem()
	}
	return json.Unmarshal(d, v)
}

func (jsonCodec) Type() string {
	return jsonType
}

func init() {
	RegisterCodec(jsonCodec{})
}
