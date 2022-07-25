package main
import "fmt"

func swap(x int,y int){
    var temp int = x
    x = y
    y = temp
    fmt.Println("After swap x y :")
    fmt.Println(x,y)
}
func main() {
    var x ,y = 10,20
    fmt.Println("Before swap x y :")
    fmt.Println(x,y)
    swap(x,y)
}
