package netclient

import (
	"io"
	"net"

	"github.com/ithinkiborkedit/GUSH-Client.git"
)

type TCPNetClient struct {
	conn  net.Conn
	codec ProtoRW
}

func NewTCPNetClient() *TCPNetClient {
	return &TCPNetClient{}
}

func (c *TCPNetClient) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	c.conn = conn

	reader := io.Reader(conn)
	writer := io.Writer(conn)

	c.codec = NewProtoRW(reader, writer)

	return nil
}

func (c *TCPNetClient) SendCommand(cmd *GUSH.Command) error {
	if c.conn == nil {
		return nil
	}

	return c.codec.Encode(cmd)
}

func (c *TCPNetClient) ReadLoop(callback func(*GUSH.ServerMessage, error)) {
	for {
		if c.conn == nil {
			callback(nil, nil)
			return
		}

		msg := &GUSH.ServerMessage{}
		err := c.codec.Decode(msg)
		if err != nil {
			callback(nil, err)
		}
		callback(msg, nil)
	}
}

func (c *TCPNetClient) Close() error {
	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}

	return nil
}
