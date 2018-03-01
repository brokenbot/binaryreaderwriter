package binaryreaderwriter

import (
	"encoding/binary"
	"io"
)

//Writer struct to create simple binaryWriter
type Writer struct {
	ByteStream io.Writer
	Endianess  binary.ByteOrder
}

//BinaryWriter Returns instance of Writer
func BinaryWriter(ByteStream io.Writer) *Writer {
	return &Writer{ByteStream: ByteStream, Endianess: binary.LittleEndian}
}

//WriteInt32 write an int32 to Writer.ByteStream
func (Writer *Writer) WriteInt32(i int32) {
	binary.Write(Writer.ByteStream, Writer.Endianess, i)
}

//Write7BitEncodedInt write a int32 as a 7bit encoded int to Writer.ByteStream
func (Writer *Writer) Write7BitEncodedInt(i int32) {
	b := uint32(i)
	for b > 0x80 {
		binary.Write(Writer.ByteStream, Writer.Endianess, byte(b|0x80))
		b = b >> 7
	}
	binary.Write(Writer.ByteStream, Writer.Endianess, byte(b))
}

//WriteString write a string prefixed with the length as a 7bit encoded int to Writer.ByteStream
func (Writer *Writer) WriteString(s string) {
	size := len(s)
	Writer.Write7BitEncodedInt(int32(size))
	//binary.Write(Writer.ByteStream, Writer.Endianess, s)
	Writer.ByteStream.Write([]byte(s))
}
