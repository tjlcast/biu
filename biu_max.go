package biu

import "errors"

var (
	BitConvertError = errors.New("BitConvertError")
)

func GetInt8Max() int8 {
	var b int8
	err := ReadBinaryString("[01111111]", &b)
	if err != nil {
		panic(err)
	}
	return b
}

func GetUint8Max() uint8 {
	var b uint8
	err := ReadBinaryString("[11111111]", &b)
	if err != nil {
		panic(err)
	}
	return b
}

func GetInt16Max() int16 {
	var b int16
	err := ReadBinaryString("[01111111 11111111]", &b)
	if err != nil {
		panic(err)
	}
	return b
}

func GetUint16Max() uint16 {
	var b uint16
	err := ReadBinaryString("[11111111 11111111]", &b)
	if err != nil {
		panic(err)
	}
	return b
}

func GetInt32Max() int32 {
	var b int32
	err := ReadBinaryString("[01111111 11111111 11111111 11111111]", &b)
	if err != nil {
		panic(err)
	}
	return b
}

func GetUint32Max() uint32 {
	var b uint32
	err := ReadBinaryString("[11111111 11111111 11111111 11111111]", &b)
	if err != nil {
		panic(err)
	}
	return b
}

func GetInt64Max() int64 {
	var b int64
	err := ReadBinaryString("[01111111 11111111 11111111 11111111 11111111 11111111 11111111 11111111]", &b)
	if err != nil {
		panic(err)
	}
	return b
}

func GetUint64Max() uint64 {
	var b uint64
	err := ReadBinaryString("[11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111111]", &b)
	if err != nil {
		panic(err)
	}
	return b
}
