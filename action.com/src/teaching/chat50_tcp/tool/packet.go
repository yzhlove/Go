package tool

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Packet struct {
	Size uint16
	Body []byte
}

func readPacket(dataReader io.Reader) (Packet, error) {
	var (
		sizeBuf = make([]byte, 2)
		pkt     Packet
		err     error
	)
	if _, err = io.ReadFull(dataReader, sizeBuf); err != nil {
		return Packet{}, err
	}
	sizeReader := bytes.NewBuffer(sizeBuf)
	if err = binary.Read(sizeReader, binary.LittleEndian, &pkt.Size); err != nil {
		return Packet{}, err
	}
	pkt.Body = make([]byte, pkt.Size)
	if _, err = io.ReadFull(dataReader, pkt.Body); err != nil {
		return Packet{}, err
	}
	return pkt, nil
}

func writePacket(dataWrite io.Writer, data []byte) error {
	var (
		buf bytes.Buffer
		err error
	)
	if err = binary.Write(&buf, binary.LittleEndian, uint16(len(data))); err != nil {
		return err
	}
	if _, err = buf.Write(data); err != nil {
		return err
	}
	if _, err = dataWrite.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
