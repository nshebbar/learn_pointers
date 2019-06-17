package main

import "fmt"

func setTo10Fail(a *int) {
	a = new(int)
	*a = 10
}
func main() {
//var d *int

	a := 20
	b:= &a
	c := a
	d := new(int)
	
	fmt.Println(a)
	setTo10Fail(&a)
	fmt.Println(a)



	fmt.Println(a, b, *b, c)

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	fmt.Println(d)
	fmt.Println(*d)
}