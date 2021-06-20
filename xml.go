package encoding

import (
	"encoding/xml"
)

var _ Codec = (*xmlCodec)(nil)

const xmlType = "xml"

type xmlCodec struct{}

func (xmlCodec) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (xmlCodec) Unmarshal(d []byte, v interface{}) error {
	return xml.Unmarshal(d, v)
}

func (xmlCodec) Type() string {
	return xmlType
}

func init() {
	RegisterCodec(xmlCodec{})
}
