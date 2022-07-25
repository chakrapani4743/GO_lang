package main
import "fmt"
func add(x int,y int)int{
    return x + y
}
func main() {
   var x int = 10
   var y int = 20
   var z int = add(x,y)
   fmt.Println(z)
}
