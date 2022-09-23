package main

import "fmt"


type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int

	fmt.Println(a)
	fmt.Println(len(a))

	for i, v := range a {
		fmt.Printf("%d - %d\n", i, v)
	}

	var q [3]int = [3]int{1, 4, 7}
	r := [...]int{2, 6, 8}
	fmt.Println(q)
	fmt.Println(r)

	symbol := [...]string{USD: "$"}
	fmt.Println(USD, symbol[USD])

}

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
