package main
import "fmt"
func main(){
	var x, y = 62, 20
    fmt.Println("operations")
	fmt.Printf("add = %d\n", x+y)
	fmt.Printf("sub = %d\n", x-y)
	fmt.Printf("mul = %d\n", x*y)
	fmt.Printf("div = %d\n", x/y)
	fmt.Printf("mod = %d\n", x%y)
	x++
	fmt.Printf("inc= %d\n", x)
	y--
	fmt.Printf("dec = %d\n", y)
}