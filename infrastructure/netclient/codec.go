package netclient

import (
	"io"

	"google.golang.org/protobuf/proto"
)

type protoBufCodec struct {
	w io.Writer
	r io.Reader
}

type ProtoRW interface {
	ProtoBufEncode
	ProtoBufDecode
}

type ProtoBufEncode interface {
	Encode(msg proto.Message) error
}

type ProtoBufDecode interface {
	Decode(msg proto.Message) error
}

func NewProtoRW(r io.Reader, w io.Writer) ProtoRW {
	return &protoBufCodec{w: w, r: r}
}

func (rw *protoBufCodec) Encode(msg proto.Message) error {
	return WriteProto(rw.w, msg)
}

func (rw *protoBufCodec) Decode(msg proto.Message) error {
	return ReadProto(rw.r, msg)
}
