package server

import (
	"encoding/binary"
	"io"
	"net"
)

type Session struct {
	conn net.Conn
}

func (s *Session) Write(data []byte) (err error) {
	buf := make([]byte, 4+len(data))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	copy(buf[4:], data)
	_, err = s.conn.Write(buf)
	return err
}

func (s *Session) Read() ([]byte, error) {
	header := make([]byte, 4)
	if _, err := io.ReadFull(s.conn, header); err != nil {
		return nil, err
	}
	dataLength := binary.BigEndian.Uint32(header)
	data := make([]byte, dataLength)
	if _, err := io.ReadFull(s.conn, data); err != nil {
		return nil, err
	}
	return data, nil
}
