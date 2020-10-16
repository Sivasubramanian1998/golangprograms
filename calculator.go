package main

import "fmt"

func add(n1, n2 int) int {
	var ans int
	ans = n1 + n2
	return ans
}
func sub(n1, n2 int) int {
	var answe int
	answe = n1 - n2
	return answe
}
func mul(n1, n2 int) int {
	var answer int
	answer = n1 * n2
	return answer
}

func main() {

	var c int
	var a, b int
	fmt.Scan(&c)

	switch c {
	case 1:
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Println("addition: %d", add(a, b))
	case 2:
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Println("Subtraction: %d", sub(a, b))
	case 3:
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Printf("multiplication: %d", mul(a, b))
	default:
		fmt.Printf("INVALID")
	}

}
