package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	enc  *gob.Encoder
	dec  *gob.Decoder
}

// 这一段是什么意思？
var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buff := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buff,
		enc:  gob.NewEncoder(buff),
		dec:  gob.NewDecoder(conn),
	}
}

func (g *GobCodec) ReadHeader(h *Header) error {
	return g.dec.Decode(h)
}

func (g *GobCodec) ReadBody(body interface{}) error {
	return g.dec.Decode(body)
}

func (g *GobCodec) Close() error {
	return g.conn.Close()
}

func (g *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = g.buf.Flush()
		if err != nil {
			_ = g.Close()
		}
	}()
	if err = g.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob error encoding header:", err)
		return err
	}

	if err = g.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}

	return nil
}
