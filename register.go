package encoding

import "strings"

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(d []byte, v interface{}) error
	Type() string
}

var registeredCodecs = make(map[string]Codec)

func RegisterCodec(c Codec) {
	if c == nil {
		panic("cannot register a nil Codec")
	}
	if c.Type() == "" {
		panic("cannot register Codec with empty string result for Type()")
	}
	registeredCodecs[strings.ToLower(c.Type())] = c
}

func GetCodec(contentSubtype string) Codec {
	return registeredCodecs[contentSubtype]
}
