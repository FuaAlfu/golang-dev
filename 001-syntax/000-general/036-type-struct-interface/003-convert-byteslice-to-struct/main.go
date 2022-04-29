package main

import(
	"fmt"
	"unsafe"
)

type(
	Movment struct{
		Opcode int
		X      float64
		Y      float64
		Z      float64
	}
)

const movmentSize = int(unsafe.Sizeof(Movment{}))

func main() {
	m := Movment{
		Opcode: 32,
		X:      56.9,
		Y:      11.89,
		Z:      70.70,
	}

	packet := (*[movmentSize]byte)(unsafe.Pointer(&m))[:]
	unsafe := *(*Movment)(unsafe.Pointer(&packet[0])) 

	fmt.Println(m)
	fmt.Println(len(packet))
	fmt.Println(packet)
	fmt.Println(unsafe)
}
