package main

import (
  "github.com/djhworld/go-lambda-invoke/golambdainvoke"
  "fmt"
)

func main() {
  res, err := golambdainvoke.Run(8001, "payload")
  fmt.Print(string(res), ":aaa:", err)
}
