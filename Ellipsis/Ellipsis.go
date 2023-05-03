//Ellipsis atau Variadic Functions

package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func mult(nums ...int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total

}

func main() {
	//contoh
	fmt.Println("ini", "contoh")

	//Call sum func
	fmt.Println(sum(1, 2, 3, 4))

	//Call mult func
	fmt.Println(mult(1, 2, 3, 4))
}
