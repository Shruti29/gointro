package main

import "fmt"

func main() {
  i, j := 10, 5
  sum := func () int{
    return i+j
  }
  fmt.Println(sum)
}
