package main

import (
  "fmt"
  corelang "github.com/kinga-ai/kinga-core/lang/go"
)

func main() {
  fmt.Println("Kweli dev bootstrap")
  fmt.Println("Core-supported languages:", corelang.Codes())
}
