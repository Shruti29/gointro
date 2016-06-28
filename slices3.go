package main
import "fmt"
type person struct {
  name string
  age int
}
func main() {
  p1 := person { name : "Federer", age : 34,}
  p2 := person { name: "Nadal", age: 30,}
  display(p1,p2)
  p3 := []person{ {"Roddick", 33}, {"Djokovic", 29}, }
  display(p3...)
}
func display(p ...person) {
  for _,i := range p {
          fmt.Printf("%+v\n", i)
  }
}
