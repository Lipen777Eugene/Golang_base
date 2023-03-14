package main

func main() {
	var a int8 = -1     //-128 <> 127 , 1 byte, 8 bit
	var b int16 = 30200 // -32768 <> 32767 , 2 bytes, 16 bit
	var c int32 = 65323 // -2 bil <> 2 bil, 4 byte
	var d int64 = 32    // -9pent <> 9pent, 8 byte
	var e uint8 = 1     // 0 <> 255 , 1 byte, 8 bit
	var f uint16 = 257
	var g uint32 = 54432
	var h uint64 = 342232
	var i byte = 101  // synonym uint8
	var j rune = 4424 // synonym int32
	var k int = -2442
	var m uint = 234422

	var a1 float32 = 1.8
	var b1 float64 = 3.42412412

	var c1 complex64 = 1 + 2i
	var d1 complex128 = 2 + 90i

	var b2 bool = true
	var b3 bool = false

	var name string = "Eugene"

	var name1 = "Eugene"
	name2 := "Eugene1"

	var (
		firstName = "Eugene"
		age       = 31
	)

}
