package main

import (
  "AdityaHegde/apricate-ai/src/main/go/apricate-clients/users"
  "fmt"
)

func main() {
  fmt.Println("Hello, World!")
  user, err := users.GetUser()
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(user)
  }
}
