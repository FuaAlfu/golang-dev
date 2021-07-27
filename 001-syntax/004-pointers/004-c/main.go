package main

func main() {
	p1 := 31
	p2 := &p1
	println("p1:\tvalue of[", p1, "]\t\taddr of [", &p1, "]")
	println("p2:\tvalue of[", p2, "]\t\taddr of [", &p2, "]")
	println("-----")

	*p2 = 88
	println("p2:\tvalue of[", p2, "]\t\taddr of [", &p2, "]")
}
