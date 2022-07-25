package main
import "fmt"

func main() {
   fmt.Println("Choose any option :\n1. Add\n2. Sub\n3. Mul\n4. Div\n")
   var x,y int = 10,5
 //  fmt.Scanf("%d",&x)
   //fmt.Scanf("%d",&y)
    opt := 1
   //fmt.Scanf("%d",&opt)
   switch opt {
       case 1:
       sum := x + y 
       fmt.Println(sum)
       case 2:
       Sub := x - y 
       fmt.Println(Sub)
       case 3:
       Mul := x * y 
       fmt.Println(Mul)
       case 4:
       Div := x / y 
       fmt.Println(Div)
     //  default :
       fmt.Println("Invalid")
   }
}
