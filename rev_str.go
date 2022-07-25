package main
import "fmt"

func main() {
    str := "RAMA"
    len1 := len(str)
    fmt.Println(len1)
   // len1--
    for i := len1 -1;i >= 0;i--{
    fmt.Printf("%c",str[i])
    }
}
