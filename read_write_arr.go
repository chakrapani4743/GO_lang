package main
import "fmt"

func main(){
    var arr [5]int
    fmt.Println("Enter arr values")
    for i := 0;i < 5;i++{
    fmt.Scanf("%d",&arr[i])
    }
    fmt.Println("Display arr values")
    for i := 0;i < 5;i++{
    fmt.Printf("%d ",arr[i])
    }
}
