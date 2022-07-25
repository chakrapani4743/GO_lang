package main
import "fmt"

func swap(x int,y int){
    x = x + y 
    y = x - y 
    x = x - y 
    fmt.Println("After swap x y :")
    fmt.Println(x,y)
}
func main(){
    x := 10
    y := 20
    fmt.Println("Before swap x y :")
    fmt.Println(x,y)
    swap(x,y)
}
