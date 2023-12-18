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
    
        var item types.Item
        item.Type = j
    
        // Check if value is a slice (i.e., []interface{})
        if slice, ok := value.([]interface{}); ok {
            item.HasCHildren = true
            for _, childValue := range slice {
                var childItem types.Item
    
                // Here you would determine the structure of each child item
                if childMap, ok := childValue.(map[string]interface{}); ok {
                    childItem.HasCHildren = false // Assuming child items don't have further nested children
                    for childKey, childValue := range childMap {
                        fmt.Print(childKey)
                        childItemType := utils.GetValueType(childValue)
                        childItem.Children = append(childItem.Children, types.Item{Value: childValue, Type: childItemType})
                    }
                } else {
                    // If the child is not a map, handle it as a simple field
                    childItemType := utils.GetValueType(childValue)
                    childItem = types.Item{Value: childValue, Type: childItemType}
                }
    
                item.Children = append(item.Children, childItem)
            }
        } else {
            // Handle non-slice values as before
            if num, ok := value.(float64); ok && num == float64(int64(num)) {
                item.Type = reflect.TypeOf(int(0))
            }
            item.Value = value
        }
    
        protoStorageTypes[key] = item
    }

    protostring :=  proto.ConvertJsonTOProto(protoStorageTypes,filename)
    fmt.Println(protostring)

    
}



