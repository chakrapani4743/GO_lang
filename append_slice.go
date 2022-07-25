package main
import "fmt"

func main() {
    myslice1 := [] int {10,20,30,40,50}
    fmt.Println(myslice1)
    fmt.Println(len(myslice1))
    fmt.Println(cap(myslice1))
    
    myslice2 := append(myslice1,60,70)
    fmt.Println(myslice2)
    fmt.Println(len(myslice2))
    fmt.Println(cap(myslice2))
}
