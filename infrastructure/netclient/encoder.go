package netclient

import (
	"encoding/binary"
	"errors"
	"io"

	"google.golang.org/protobuf/proto"
)

// type ProtoBufCodec struct {
// 	w io.Writer
// 	r io.Reader
// }

// type ProtoBufEncode interface {
// 	WriteProto(msg proto.Message) error
// }

// type ProtoBufDecode interface {
// 	ReadProto(msg proto.Message) error
// }(w io.Writer, msg proto.Message)

func (pb *ProtoBufCodec) Encode(w io.Writer, msg proto.Message) error {
	return writeProto(pb.w, msg)

}

func (pb *ProtoBufCodec) Decode(r io.Reader, msg proto.Message) error {
	return readProto(pb.r, msg)
}

func writeProto(w io.Writer, msg proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	lengthBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuf, uint32(len(data)))

	if _, err := w.Write(lengthBuf); err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		return err
	}
	return nil
}

func readProto(r io.Reader, msg proto.Message) error {
	lengthBuf := make([]byte, 4)
	if _, err := io.ReadFull(r, lengthBuf); err != nil {
		return err
	}

	length := binary.BigEndian.Uint32(lengthBuf)
	if length == 0 {
		return errors.New("zero-length message")
	}
	data := make([]byte, length)
	if _, err := io.ReadFull(r, data); err != nil {
		return err
	}

	return proto.Unmarshal(data, msg)
}
