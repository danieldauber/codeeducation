package main
import "fmt"

func plus(a int, b int) int {
    
    return a + b
}


func main() {

    res := plus(5, 5)
    fmt.Println("5+5 =", res)
}