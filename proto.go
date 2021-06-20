package encoding

import "google.golang.org/protobuf/proto"

var _ Codec = (*protoCodec)(nil)

const protoType = "proto"

type protoCodec struct{}

func (protoCodec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (protoCodec) Unmarshal(d []byte, v interface{}) error {
	return proto.Unmarshal(d, v.(proto.Message))
}

func (protoCodec) Type() string {
	return protoType
}

func init() {
	RegisterCodec(protoCodec{})
}
