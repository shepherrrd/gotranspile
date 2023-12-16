package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
   
   if len(os.Args) < 2 {
      log.Fatal("Usage: go run main.go [path_to_json_file]")
  }
  filePath := os.Args[1]
  data, err := os.ReadFile(filePath)
  if err != nil {
      log.Fatalf("Error reading file: %s", err)
  }
  jsonStr := string(data)
  fmt.Println(jsonStr)
}