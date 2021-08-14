package ciscoise

import (
	"fmt"
)

func interfaceToSliceString(v interface{}) []string {
	newValue := []string{}
	value := v.([]interface{})
	for _, i := range value {
		newValue = append(newValue, fmt.Sprintf("%v", i))
	}
	return newValue
}
