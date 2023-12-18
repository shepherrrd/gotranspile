package utils

import (
	"reflect"
)

func GetValueType(variableToCheck interface{}) reflect.Type{
	return reflect.TypeOf(variableToCheck)
   }