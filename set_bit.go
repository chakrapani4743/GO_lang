package main
import "fmt"

func set_bit(n int,k uint)int{
    n |= 1 << k
    return n
}
func main() {
    var n int = 10
    var k uint = 2 
     fmt.Printf("%d",set_bit(n,k))
}
