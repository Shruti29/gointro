package main

import "fmt"

func main() {
  // Declare an array of strings initalized to zero value
  var strings [5]string

  // Decalre an array of ints initialized to some values
  intArray := [4]int{10,11,12,13}

  strings[0] = "Batman"
  strings[1] = "Superman"
  strings[2] = "Wonderwoman"
  strings[3] = "Ironman"
  strings[4] = "Shaktiman"

  for i,super := range intArray {
    fmt.Println(i, super)
  }

  for i,v := range strings {
    fmt.Printf("Value [%s] Address [%p] IndexArr [%p]\n", v, &v, &strings[i])
  }
}
