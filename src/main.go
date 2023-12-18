package main

import (
	"encoding/json"
	"fmt"
	"gotranspile/proto"
	"gotranspile/types"
	"gotranspile/utils"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
   
   if len(os.Args) < 2 {
      log.Fatal("Usage: go run main.go [path_to_json_file]")
      return
  }
  var jsonObj map[string]interface{}
  filePath := os.Args[1]
  filename := strings.Split(filePath, ".")[0]
  data, err := os.ReadFile(filePath)
  if err != nil {
      log.Fatalf("Error reading file: %s", err)
  }
   protoStorageTypes := make(map[string]types.Item)
  jsonStr := string(data)
  er2 := json.Unmarshal([]byte(jsonStr), &jsonObj)
  fmt.Print(jsonStr)
    if er2 != nil {
        log.Fatal(err)
    }
    for key, value := range jsonObj {
      j := utils.GetValueType(value)
      if num, ok := value.(float64); ok {
        if num == float64(int64(num)) {
            j = reflect.TypeOf(int(0))
        }
    }
    print(value)
      protoStorageTypes[key] = types.Item{Value: value, Type: j}
  }
  for d , j := range protoStorageTypes {
    fmt.Println(d ,j.Value, j.Type)
  }

    protostring :=  proto.ConvertJsonTOProto(protoStorageTypes,filename)
    fmt.Println(protostring)

    
}



