package ciscoise

import (
	"fmt"
)

func interfaceToSliceString(v interface{}) []string {
	value, ok := v.([]interface{})
	if !ok {
		return nil
	}
	newValue := []string{}
	for _, i := range value {
		newValue = append(newValue, interfaceToString(i))
	}
	return newValue
}

func interfaceToString(v interface{}) string {
	return fmt.Sprint(v)
}
