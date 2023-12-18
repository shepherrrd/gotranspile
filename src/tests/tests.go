package tests

import (
	"gotranspile/proto"
	"gotranspile/types"
	"gotranspile/utils"
	"reflect"
	"testing"
)

func TestGetValueType(t *testing.T) {
    testCases := []struct {
        name string
        input interface{}
        expected reflect.Type
    }{
        {"Int", 1, reflect.TypeOf(int(0))},
        {"Float", 1.1, reflect.TypeOf(float64(0))},
        {"String", "test", reflect.TypeOf("")},
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := utils.GetValueType(tc.input)
            if result != tc.expected {
                t.Errorf("Expected %v, got %v", tc.expected, result)
            }
        })
    }
}

func TestConvertJsonTOProto(t *testing.T) {
    testCases := []struct {
        name string
        input map[string]types.Item
        protoname string
        expected string
    }{
        {"Empty", map[string]types.Item{}, "test", "sytax = 'proto3;'\nmessage test {\n}\n"},
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := proto.ConvertJsonTOProto(tc.input, tc.protoname)
            if result != tc.expected {
                t.Errorf("Expected '%s', got '%s'", tc.expected, result)
            }
        })
    }
}