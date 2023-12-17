package proto

import (
	"fmt"
	"gotranspile/types"
	"log"
	"os"
)
func ConvertJsonTOProto(protoStorageTypes map[string]types.Item, protoname string) string{
	var result string = "sytax = 'proto3;'\n"
	result += fmt.Sprintf("message %s {\n",protoname)
	for key,valuepair := range protoStorageTypes{
		count := 1
		result += fmt.Sprintf("    %v %s = %d;\n", valuepair.Type, key, count)
		count++
	}
	result += "}\n"
	data := []byte(result)
	parentDir := "../results" // Parent directory
    errr := os.MkdirAll(parentDir, os.ModePerm) // Create the directory if it doesn't exist
    if errr != nil {
        log.Fatal(errr)
    }
	path := fmt.Sprintf("../results/%s.proto",protoname)
	err := os.WriteFile(path, data, 0644)
    if err != nil {
        log.Fatal(err)
		return "Could Not Save File"
    }

	return result
}