package ciscoise

import (
	"encoding/json"
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

func responseInterfaceToSliceString(v interface{}) []string {
	value, ok := v.([]interface{})
	if !ok {
		return nil
	}
	newValue := []string{}
	for _, i := range value {
		newValue = append(newValue, responseInterfaceToString(i))
	}
	return newValue
}

func responseInterfaceToString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprint(v)
	}
	return fmt.Sprint(string(b))
}
