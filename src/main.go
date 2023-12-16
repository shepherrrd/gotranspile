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
    if er2 != nil {
        log.Fatal(err)
    }
    for key, value := range jsonObj {
      j := getValueType(value)
      protoStorageTypes[key] = Item{Value: value, Type: j}
  }

  for key,value := range protoStorageTypes{
   fmt.Printf("The Name of the proto data is %s and the value is %s and type is %s \n",key,value.Value,value.Type)
  }
  print(reflect.TypeOf(30))
}
func getValueType(variableToCheck interface{}) reflect.Type{
 return reflect.TypeOf(variableToCheck)
}
func print(args interface{}){
   fmt.Println(args)
}
