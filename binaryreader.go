package binaryreaderwriter

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Reader instance of BinaryReaderWriter, use BinaryReader()
type Reader struct {
	ByteStream io.Reader
	Endianess  binary.ByteOrder
}

//BinaryReader creates a new binary reader instance with default values
func BinaryReader(ByteStream io.Reader) *Reader {
	return &Reader{ByteStream: ByteStream, Endianess: binary.LittleEndian}
}

//ReadByte reads one Byte from Reader.ByteStream
func (Reader *Reader) ReadByte() (b byte, err error) {
	err = binary.Read(Reader.ByteStream, Reader.Endianess, &b)
	return
}

//ReadBytes reads i number of bytes Reader.ByteStream
func (Reader *Reader) ReadBytes(i int) (bytes []byte, err error) {
	bytes = make([]byte, i)
	_, err = Reader.ByteStream.Read(bytes)
	return
}

//ReadChar reads 1 char from Reader.ByteStream
func (Reader *Reader) ReadChar() (c rune, err error) {
	err = binary.Read(Reader.ByteStream, Reader.Endianess, &c)
	return
}

//ReadChars read i chars from Reader.ByteStream
func (Reader *Reader) ReadChars(i int) (chars []rune, err error) {
	chars = make([]rune, i)
	for n := 0; n < i; n++ {
		chars[n], err = Reader.ReadChar()
	}
	return
}

//ReadInt8 read single 8bit integer from BinaryReader.ByteStream
func (Reader *Reader) ReadInt8() (i int8, err error) {
	err = binary.Read(Reader.ByteStream, Reader.Endianess, &i)
	return
}

// ReadInt32 read a single 32bit from BinaryReader.ByteStream
func (Reader *Reader) ReadInt32() (i int32, err error) {
	err = binary.Read(Reader.ByteStream, Reader.Endianess, &i)
	return
}

// Read7BitEncodedInt read 7bit encoded int32 from BinaryReader.ByteStream
func (Reader *Reader) Read7BitEncodedInt() (i int32, err error) {
	var b uint8
	for n := 1; ; n++ {
		if n > 6 {
			err = fmt.Errorf("Invalid 7bit Encoded Int")
		}
		err = binary.Read(Reader.ByteStream, Reader.Endianess, &b)
		i = i | int32(b&0x7F)
		if b >= 0x80 {
			i = i << 7
		} else {
			break
		}
	}
	return
}

//ReadString reads a string prefix with the length in a 7bit encoded int
func (Reader *Reader) ReadString() (s string, err error) {
	size, err := Reader.Read7BitEncodedInt()
	if err != nil {
		return
	}
	bString := make([]byte, size)
	_, err = Reader.ByteStream.Read(bString)
	if err != nil {
		return
	}
	s = string(bString)
	return
}
