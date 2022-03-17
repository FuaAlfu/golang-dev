package main

import(
	"fmt"
    "golang.org/x/example/stringutil"
)

func reverse(s string)string{
	return s
}
func main() {
	fmt.Println(stringutil.Reverse(reverse("fua")))
	fmt.Println(stringutil.ToUpper("fua"))
}