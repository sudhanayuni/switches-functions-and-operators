package main
import "fmt"
func main() {
	var x uint = 25
	var y uint = 62
	var result uint
	fmt.Println(x)
	fmt.Println(y)
	result = x ^ y
	fmt.Println("x ^ y  =", result)

	result = x & y
	fmt.Println("x & y  =", result)

	result = x | y
	fmt.Println("x | y  =", result)

	result = x << 1
	fmt.Println("x << 1 =", result)

	result = x >> 1
	fmt.Println("x >> 1 =", result)
}