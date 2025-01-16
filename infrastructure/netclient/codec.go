package netclient

import (
	"io"

	"google.golang.org/protobuf/proto"
)

type ProtoBufCodec struct {
	w io.Writer
	r io.Reader
}

type ProtoBufEncode interface {
	WriteProto(w io.Writer, msg proto.Message) error
}

type ProtoBufDecode interface {
	ReadProto(r io.Reader, msg proto.Message) error
}
