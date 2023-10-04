package main

import (
  "fmt"
)

func main() {

  menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

  fmt.Println("The menu:")
  // WRITE LOOP GOING THROUGH MENU BELOW THIS LINE
  for number,item := range menu {
    fmt.Println(number , item)
  }

  orders := map[string]string{
    "John": "Cheeseburgers",
    "Janet": "Hamburgers",
    "Jordan": "Fries",
  }
  // A late order
  orders["James"] = "Chicken Sandwich"
  
    fmt.Println("\nThe friend's orders:")
  // WRITE LOOP GOING THROUGH ORDERS BELOW THIS LINE
  for key,value := range orders {
    fmt.Println(key,value)
  }
  

}  
