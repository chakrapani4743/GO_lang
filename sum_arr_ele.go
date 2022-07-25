package main
import "fmt"

func main(){
    arr :=[] int {1,2,3,4,5,6,7}
    len1 := len(arr)
    sum := 0
    for i := 0;i < len1;i++{
    sum += arr[i]
    }
    fmt.Println(sum)
}
