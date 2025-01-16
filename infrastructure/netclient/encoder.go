package netclient

import (
	"encoding/binary"
	"errors"
	"io"

	"google.golang.org/protobuf/proto"
)

func WriteProto(w io.Writer, msg proto.Message) error {
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

func ReadProto(r io.Reader, msg proto.Message) error {
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
