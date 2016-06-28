package main

import "fmt"

func main() {
  // Declare a string slice with nil value
  var slice1 []string
  // Print length and capacity of the slice
  fmt.Printf("Length of slice1 %v\nCapacity of slice1 %v\n", len(slice1), cap(slice1))
  // Declare a slice of ints with different capacity and value
  var slice2 = make([]int, 5, 8)
  slice2 = []int {3,4,5,6,7}
  // Call change slice for slice 2
  ChangeSlice(slice2)
  // Print changed slice
  for _,v := range slice2 {
    fmt.Printf("%v\t", v)
  }
}

func ChangeSlice(slice2 []int) {
  for i,_ := range slice2 {
    slice2[i] = slice2[i]+ 1
  }
}
