package main 
import "fmt"
func main() { 
    var p int = 50
    var q int = 25
        fmt.Println("logical operators")
    if(p!=q && p<=q){  
        fmt.Println("True") 
    } 
        
    if(p!=q || p<=q){  
        fmt.Println("True") 
    } 
        
    if(!(p==q)){  
        fmt.Println("True") 
    } 
        
} 
