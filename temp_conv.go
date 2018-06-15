package main

import "fmt"

func main() {
  fmt.Println(`Celsius to Fahrenheit converter.`)
  fmt.Print("Please enter the temperature to convert: ")
  var input float64
  fmt.Scanf("%f", &input)
  
  output := 32 + (input * 1.8)

  fmt.Println(output)
}
