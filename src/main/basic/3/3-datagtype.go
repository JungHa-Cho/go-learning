package main

import "fmt"

func main() {
	// bool
	var exampleBool bool = true
	fmt.Println(exampleBool)

	// string
	var exampleString string = "example"
	fmt.Println(exampleString)

	// String is Immutable, ?
	exampleString = "not fix"
	fmt.Println(exampleString)

	// integer, number is bit
	var exampleInt int = 1
	var exampleInt8 int8 = 2   // 8 bit is 1 byte
	var exampleInt16 int16 = 3 // 16 bit is 2 byte
	var exampleInt32 int32 = 4 // 32 bit is 4 byte
	var exampleInt64 int64 = 5 // 64 bit is 8 byte

	var exampleUint uint = 1
	var exampleUint8 uint8 = 2   // 8 bit is 1 byte
	var exampleUint16 uint16 = 3 // 16 bit is 2 byte
	var exampleUint32 uint32 = 4 // 32 bit is 4 byte
	var exampleUint64 uint64 = 5 // 64 bit is 8 byte

	//var exampleUintptr uintptr = exampleUint64
	fmt.Println(exampleInt, exampleInt8, exampleInt16, exampleInt32, exampleInt64, exampleUint, exampleUint8, exampleUint16, exampleUint32, exampleUint64)

	// var exampleFloat32 float32 = 3.14
	// var exampleFloat64 float64 = 3.141
	// var complex64 complex64
	// var complex64 complex128

	// byte = uint8
	// rune = int32
}
