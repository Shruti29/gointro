package main

import (
  "fmt"
)
func main() {
  b1 := true         // type is bool
  n1 := 100         // int
  n2 := 154.678     // float32/64
  n3 := 2e64         // float32/64
  n4 := uint8(100)   // uint
  n5 := float32(100) // float32
  s1 := `Hello world`
  s2 := "Hello again"

  fmt.Printf("%T (%v)\n", b1, b1)
  fmt.Printf("%T (%v)\n", n1, n1)
  fmt.Printf("%T (%v)\n", n2, n2)
  fmt.Printf("%T (%v)\n", n3, n3)
  fmt.Printf("%T (%v)\n", n4, n4)
  fmt.Printf("%T (%v)\n", n5, n5)
  fmt.Printf("%T (%v)\n", s1, s1)
  fmt.Printf("%T (%v)\n", s2, s2)
}
