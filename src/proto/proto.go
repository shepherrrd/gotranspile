package proto

import (
	"fmt"
	"gotranspile/types"
	"log"
	"os"
	"strings"
)

func ConvertJsonTOProto(protoStorageTypes map[string]types.Item, protoname string) string {
    var result strings.Builder
    result.WriteString(fmt.Sprintf("syntax = 'proto3';\nmessage %s {\n", protoname))
    
    fieldCount := 1
    for key, valuepair := range protoStorageTypes {
        if valuepair.HasCHildren {
            // Define a nested message for children
            childMessageName := key // Using key as the message name
            result.WriteString(fmt.Sprintf("    message %s {\n", childMessageName))

            childCount := 1
            for _, child := range valuepair.Children {
                protobufType := getProtobufType(child.Type)
                result.WriteString(fmt.Sprintf("        %s %s = %d;\n", protobufType, key, childCount))
                childCount++
            }
            result.WriteString("    }\n")
            result.WriteString(fmt.Sprintf("    repeated %s %s = %d;\n", childMessageName, key, fieldCount))
        } else {
            protobufType := getProtobufType(valuepair.Type)
            result.WriteString(fmt.Sprintf("    %s %s = %d;\n", protobufType, key, fieldCount))
        }
        fieldCount++
    }
    result.WriteString("}\n")
	saveProtoFile(result.String(), protoname)
    return result.String()
}

func saveProtoFile(protoContent, protoname string) {
	data := []byte(protoContent)
	parentDir := "../results" // Parent directory
	err := os.MkdirAll(parentDir, os.ModePerm) // Create the directory if it doesn't exist
	if err != nil {
		log.Fatalf("Failed to create directory: %s", err)
		return
	}

	path := fmt.Sprintf("../results/%s.proto", protoname)
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %s", err)
		return
	}
	log.Printf("Protobuf file %s created successfully\n", path)
}

func getProtobufType(goType interface{}) string {
    switch goType.(type) {
    case int, int32, int64:
        return "int32"
    case float32, float64:
        return "float"
    case string:
        return "string"
    case bool:
        return "bool"
    default:
        return "string" 
    }
}