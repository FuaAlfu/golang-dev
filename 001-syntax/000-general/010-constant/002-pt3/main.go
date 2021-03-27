package main

import "fmt"

type (
	fua struct{ f string }
)

func main() {
	/*
		iota gives us a block
		we coud say, its a syntax of block, and its work only insid a block {}
	*/
	const (
		//not a good practice
		A1 = iota //0 :Start at 0
		B1 = iota //1 :Increment by 1
		C1 = iota //2 :Increment by 1
	)
	fmt.Println("1: ", A1, B1, C1)
	println("---")

	const (
		//this is how we should write it
		A2 = iota //0 :Start at 0
		B2        //1 :Increment by 1
		C2        //2 :Increment by 1
	)
	fmt.Println("2: ", A2, B2, C2)
	println("---")

	const (
		//here we were performing a math
		A3 = iota + 1 //0 :Start at 1
		B3            //1 :Increment by 1
		C3            //2 :Increment by 1
	)
	fmt.Println("3: ", A3, B3, C3)
	println("---")

	const (
		//bitwise operations (shifting) :: we could see it inside log pkg ::set of constat -> flages
		Ldate         = 1 << iota //1: shift 1 to the left 0. 0000 0001
		Ltime                     //2: shift 1 to the left 1. 0000 0010
		Lmicroseconds             //4: shift 1 to the left 2. 0000 0100
		Llongfile                 //8: shift 1 to the left 3. 0000 1000
		Lshortfile                //16: shift 1 to the left 4. 0001 0000
		LUTC                      //32: shift 1 to the left 5. 0010 0000
	)
	fmt.Println("log: ", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
	println("---")

	f := fua{
		f: "written by fua",
	}
	fmt.Printf("f:[%d]", f.f)

}
