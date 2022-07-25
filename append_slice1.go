/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

func main() {
    myslice1 :=[]int {1,2,3,4}
    myslice2 := [] int {5,6,7,8}
    myslice3 := append(myslice1,myslice2...)
    fmt.Println(myslice1)
    fmt.Println(myslice2)
    fmt.Println(myslice3)
}
