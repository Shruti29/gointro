package main

import "fmt"

type user struct {
  name string
  surname string
}
type Data map[string]string

func main() {
  users := map[string]user { "Einstein": {"Albert", "Einstein"}, "Faraday": {"Michael", "Faraday"}, "Feynman": {"Richard", "Feynman"},}
  for key, value := range users {
    fmt.Printf("Key : %v\tValue : %v\n", key,value)
    }
  u := make(map[string]Data)
  u["clients"] = Data{"123": "Henry", "345" : "Osborn"}
  for key, data := range u {
    fmt.Println("\n",key)
    for key, value := range data {
      fmt.Println("\t", key, value)
    }
  }
}
