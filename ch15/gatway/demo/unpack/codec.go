package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

const Msg_Header = "12345678"

// Encode 信息压缩
func Encode(bytesBuffer io.Writer, content string) error {
	// msf_header + content_len + content
	// 8+4+content_len
	// 写入massHead
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(Msg_Header)); err != nil {
		return err
	}
	// 写入msg的长度
	clen := int32(len([]byte(content)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}
	// 写入数据
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return nil
	}
	return nil
}

// Decode 信息解压
func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	MagicBuf := make([]byte, len(Msg_Header))
	if _, err = io.ReadFull(bytesBuffer, MagicBuf); err != nil {
		return nil, err
	}
	if string(MagicBuf) != Msg_Header {
		return nil, errors.New("msg_header error")
	}

	lengthBuf := make([]byte, 4)
	if _, err = io.ReadFull(bytesBuffer, lengthBuf); err != nil {
		return nil, err
	}

	lenth := binary.BigEndian.Uint32(lengthBuf)
	bodyBuf = make([]byte, lenth)
	if _, err = io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}

	return bodyBuf, err
}
