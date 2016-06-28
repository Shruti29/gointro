package main

import "fmt"

func main() {
  // Declare a slice with 10 ints
  slice := []int{10,20,30,40,50,60,70,80,90,100}

  // Point to the seconds element of the slice
  second := &slice[1]

  // Append a new value : append is one of the many builtin functions provided by Go
  slice = append(slice, 110)

  // Modify second element
  slice[1] = 25

  fmt.Println("Pointer:", *second, "Element", slice[1])

}
