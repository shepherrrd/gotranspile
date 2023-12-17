package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
   type Item struct {
      Type interface{}
      Value  interface{}
  }
   if len(os.Args) < 2 {
      log.Fatal("Usage: go run main.go [path_to_json_file]")
      return
  }
  var jsonObj map[string]interface{}
  filePath := os.Args[1]
  data, err := os.ReadFile(filePath)
  if err != nil {
      log.Fatalf("Error reading file: %s", err)
  }
   protoStorageTypes := make(map[string]Item)
  jsonStr := string(data)
  er2 := json.Unmarshal([]byte(jsonStr), &jsonObj)
  fmt.Print(jsonStr)
    if er2 != nil {
        log.Fatal(err)
    }
    for key, value := range jsonObj {
      j := getValueType(value)
      if num, ok := value.(float64); ok {
        if num == float64(int64(num)) {
            j = reflect.TypeOf(int(0))
        }
    }
    print(value)
      protoStorageTypes[key] = Item{Value: value, Type: j}
  }
  for d , j := range protoStorageTypes {
    fmt.Println(d ,j.Value, j.Type)
  }
}
func getValueType(variableToCheck interface{}) reflect.Type{
 return reflect.TypeOf(variableToCheck)
}
func print(args interface{}){
   fmt.Println(args)
}
