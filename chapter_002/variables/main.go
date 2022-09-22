package main

import "fmt"

func main() {
	var name string = "Md Rasheduzzaman"
	var exe string
	fmt.Println(name)
	fmt.Println(exe)

	var b, f, s = true, 2.3, "four"
	fmt.Printf("%t - %f - %s\n", b, f, s)

	x := 1
	p := &x
	fmt.Println(p)
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	var y, z int
	fmt.Println(&y == &y, &y == &z, &y == nil)

	fmt.Println(ff() == ff())

}


func ff() *int {
	v := 1
	return &v
}
