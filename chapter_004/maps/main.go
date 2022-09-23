package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["Rashed"] = 32
	ages["Arna"] = 26

	fmt.Println(ages)

	address := map[string]string{
		"Dhaka":   "Mahammadpur",
		"Kushtia": "Kamlapur",
	}

	fmt.Println(address)

	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}

	if age, ok := ages["Rashed"]; !ok {
		fmt.Println("Wrong key")
	} else {
		fmt.Printf("Rashed exist. Age: %d\n", age)
	}

}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
