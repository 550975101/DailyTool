package main

import "fmt"

func swap(pa *int, pb *int) {
	*pa, *pb = *pb, *pa
}

func main() {
	a := 10
	b := 20

	fmt.Println("a = ", a, "b = ", b)
	// swap
	swap(&a, &b)

	fmt.Println("a = ", a, "b = ", b)

	params := make(map[string]interface{}, 10)
	params["key"] = "values"
	fmt.Println(params)
}
