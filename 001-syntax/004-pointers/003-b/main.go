package main

func main() {
	/*
		main rule of pointer semantic:
		don't operate on your own copy of data..
	*/
	amount := 10
	println("amount:\tvalue of[", amount, "]\t\taddr of [", amount, "]")

	//pass the address of amount :: give increment  to a copy of own address to that integer
	increment(&amount)
	println("amount:\tvalue of[", amount, "]\t\taddr of [", amount, "]")

}

//always an address and points to values of type int :: pointers allow share access
func increment(inc *int) {
	//increment the value of amount that pointer points to..
	*inc++
	println("inc:\tvalue of [", inc, "]\taddress of[", &inc, "]")
}
