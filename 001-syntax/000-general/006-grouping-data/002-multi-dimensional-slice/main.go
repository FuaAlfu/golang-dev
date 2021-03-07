package main

import "fmt"

func main() {

	//a Whole slice Type...
	fA := []string{"Fahd", "Alfu"}
	fmt.Println(fA)

	mu := []string{"Mua", "Alfu"}
	fmt.Println(mu)

	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Nata", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	fmt.Println("======================================================")

	//Multi Dimintional..Only Accepted two slices ...[as we know]
	xp := [][]string{jb, mp}
	fmt.Println(xp)
	println("---")

	xpO := [][]string{fA, mp}
	fmt.Println(xpO)
	println("---")

	//wrong Dimintional..Not working like that..yet);
	//xpp := [][][]string{jb,mp,fA}
	//fmt.Println(xpp)

}
