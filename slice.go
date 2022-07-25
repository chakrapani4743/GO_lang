/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

//Create myslice :
/*
func main() {
    myslice := []int {1,2,3,4,5,6}
    fmt.Println(myslice)
}*/

// Modify and access slice elements:
/*
func main(){
    myslice := [] int {1,2,3,4,5,6}
    myslice[3] = 10
    fmt.Println(myslice)
}*/
// Create slice from array
func main(){
    var arr1 = []int {1,2,3,4,5,6}
    myslice := arr1[1:4]
    fmt.Println(myslice)
}
