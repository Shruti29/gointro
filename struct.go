package main

import "fmt"

type Unicorn struct {
  Name string
  Age int
}

func main() {
    fmt.Println(Unicorn{"Harry", 55})
    fmt.Println(Unicorn{Name: "Harry", Age: 55})
    fmt.Println(Unicorn{Name: "Harry"})
    fmt.Println(&Unicorn{Name: "Harry", Age: 55})
    u := Unicorn{Name: "Harry", Age: 55}
    fmt.Println(u.Name)
    up := &u
    fmt.Println(up.Age)
    up.Age = 51
    fmt.Println(up.Age)
}
