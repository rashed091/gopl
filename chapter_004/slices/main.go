package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(months)
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	empty := make([]int, 3, 9)
	empty = append(empty, 1)
	fmt.Println(empty)

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
