package main    
import "fmt"  
func main(){ 
   a:= 100
   b:= 50   
   sum:= a + b 
   fmt.Printf("sum of a + b = %d", sum)      
   sub:= a - b 
   fmt.Printf("\nsub of a - b = %d", sub)   
   mult:= a * b 
   fmt.Printf("\nmult of a * b = %d", mult) 
   div:= a / b 
   fmt.Printf("\ndiv of a / b = %d", div)  
   mod:= a % b 
   fmt.Printf("\nmod of a %% b = %d", mod) 
} 