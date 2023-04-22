package codec

import (
	"bufio"
	"encoding/gob"
	"io"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  bufio.Writer
	enc  *gob.Encoder
	dec  gob.Decoder
}

// 这一段是什么意思？
var _ Codec = (*GobCodec)(nil)
