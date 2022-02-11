package ciscoise

import (
	"reflect"
	"testing"
)

func TestConvertMapInterfaceToMapString(t *testing.T) {
	tests := []struct {
		val         map[string]interface{}
		out         map[string]string
		errExpected bool
	}{
		{
			map[string]interface{}{"string": "string", "int": 1, "float": 1.10, "bool": true},
			map[string]string{"string": "string", "int": "1", "float": "1.1", "bool": "true"},
			false,
		},
		{
			map[string]interface{}{"string": []string{"string"}, "int": []int{1}, "float": []float64{1.10}, "bool": []bool{true}},
			map[string]string{"string": "[string]", "int": "[1]", "float": "[1.1]", "bool": "[true]"},
			false,
		},
	}

	for _, test := range tests {
		out := mapInterfaceToMapString(test.val)
		hasError := false
		for k, v := range out {
			if tv, tok := test.out[k]; tok {
				if tv != v {
					hasError = hasError || true
					break
				}
			}
		}
		if hasError && !test.errExpected {
			t.Errorf("Mismatch on val %#v: expected %#v but got %#v", test.val, test.out, out)
		}
		if !hasError && test.errExpected {
			t.Errorf("Expected error on val %#v: expected %#v, got %#v", test.val, test.out, out)
		}
	}
}

func TestConvertInterfaceToIntPtr(t *testing.T) {
	tests := []struct {
		val         interface{}
		out         int
		nilExpected bool
	}{
		{
			"1",
			1,
			false,
		},
		{
			"-1",
			-1,
			false,
		},
		{
			"-1.32",
			0,
			true,
		},
		{
			"string",
			0,
			true,
		},
	}

	for _, test := range tests {
		out := interfaceToIntPtr(test.val)
		if test.nilExpected {
			if out != nil {
				t.Errorf("Mismatch on val %#v: expected nil but got %#v", test.val, out)
			}
		} else {
			if out == nil {
				t.Errorf("Mismatch on val %#v: expected %#v but got nil", test.val, test.out)
			} else {
				o := *out
				if o != test.out {
					t.Errorf("Mismatch on val %#v: expected %#v but got pointer to %#v", test.val, test.out, o)
				}
			}
		}
	}
}

func TestConvertBoolPtrToString(t *testing.T) {
	boolTrue := true
	boolFalse := false
	tests := []struct {
		val *bool
		out string
	}{
		{
			&boolTrue,
			"true",
		},
		{
			&boolFalse,
			"false",
		},
		{
			nil,
			"",
		},
	}

	for _, test := range tests {
		out := boolPtrToString(test.val)
		if out != test.out {
			t.Errorf("Mismatch on val %#v: expected %#v but got %#v", test.val, test.out, out)
		}
	}
}

func TestConvertInterfaceToBoolPtr(t *testing.T) {
	tests := []struct {
		val         interface{}
		out         bool
		nilExpected bool
	}{
		{
			"",
			false,
			true,
		},
		{
			"on",
			true,
			false,
		},
		{
			"true",
			true,
			false,
		},
		{
			"off",
			false,
			false,
		},
		{
			"false",
			false,
			false,
		},
	}

	for _, test := range tests {
		out := interfaceToBoolPtr(test.val)
		if test.nilExpected {
			if out != nil {
				t.Errorf("Mismatch on val %#v: expected nil but got %#v", test.val, out)
			}
		} else {
			if out == nil {
				t.Errorf("Mismatch on val %#v: expected %#v but got nil", test.val, test.out)
			} else {
				o := *out
				if o != test.out {
					t.Errorf("Mismatch on val %#v: expected %#v but got pointer to %#v", test.val, test.out, o)
				}
			}
		}
	}
}

func TestConvertGetResourceItem(t *testing.T) {
	tests := []struct {
		val         interface{}
		nilExpected bool
	}{
		{
			[]interface{}{map[string]interface{}{"key": "1"}},
			false,
		},
		{
			[]map[string]interface{}{},
			true,
		},
		{
			[]interface{}{},
			true,
		},
		{
			[]interface{}{"key"},
			true,
		},
	}

	for _, test := range tests {
		out := getResourceItem(test.val)
		if test.nilExpected {
			if out != nil {
				t.Errorf("Mismatch on val %#v: expected nil but got %#v", test.val, out)
			}
		} else {
			if out == nil {
				t.Errorf("Mismatch on val %#v: not expected nil but got nil", test.val)
			}
		}
	}
}

func TestConvertInterfaceToSliceString(t *testing.T) {
	tests := []struct {
		val         interface{}
		out         []string
		errExpected bool
	}{
		{
			"string",
			[]string{},
			false,
		},
		{
			1,
			[]string{},
			false,
		},
		{
			[]interface{}{},
			[]string{},
			false,
		},
		{
			[]interface{}{"1", "2"},
			[]string{"1", "2"},
			false,
		},
		{
			[]interface{}{"1", "2"},
			[]string{"2", "1"},
			true,
		},
	}

	for _, test := range tests {
		out := interfaceToSliceString(test.val)
		hasError := false
		if len(out) != len(test.out) {
			hasError = true
		} else {
			for i := range out {
				hasError = hasError || (out[i] != test.out[i])
			}
		}
		if hasError && !test.errExpected {
			t.Errorf("Mismatch on val %#v: expected %#v but got %#v", test.val, test.out, out)
		}
		if !hasError && test.errExpected {
			t.Errorf("Expected error on val %#v: expected %#v, got %#v", test.val, test.out, out)
		}
	}
}

func TestConvertResponseInterfaceToSliceString(t *testing.T) {
	tests := []struct {
		val         interface{}
		out         []string
		errExpected bool
	}{
		{
			"string",
			[]string{},
			false,
		},
		{
			1,
			[]string{},
			false,
		},
		{
			[]interface{}{},
			[]string{},
			false,
		},
		{
			[]interface{}{"1", "2"},
			[]string{"\"1\"", "\"2\""},
			false,
		},
		{
			[]interface{}{1, 2},
			[]string{"1", "2"},
			false,
		},
		{
			[]interface{}{"1", "2"},
			[]string{"1", "2"},
			true,
		},
	}

	for _, test := range tests {
		out := responseInterfaceToSliceString(test.val)
		hasError := false
		if len(out) != len(test.out) {
			hasError = true
		} else {
			for i := range out {
				hasError = hasError || (out[i] != test.out[i])
			}
		}
		if hasError && !test.errExpected {
			t.Errorf("Mismatch on val %#v: expected %#v but got %#v", test.val, test.out, out)
		}
		if !hasError && test.errExpected {
			t.Errorf("Expected error on val %#v: expected %#v, got %#v", test.val, test.out, out)
		}
	}
}

func TestConvertResponseInterfaceToString(t *testing.T) {
	tests := []struct {
		val         interface{}
		out         string
		errExpected bool
	}{
		{
			"string",
			"\"string\"",
			false,
		},
		{
			1,
			"1",
			false,
		},
		{
			1.50,
			"1.5",
			false,
		},
		{
			[]interface{}{"1", "2"},
			"[\"1\",\"2\"]",
			false,
		},
		{
			[]interface{}{1, "2", true},
			"[1,\"2\",true]",
			false,
		},
		{
			nil,
			"null",
			false,
		},
	}

	for _, test := range tests {
		out := responseInterfaceToString(test.val)
		hasError := false
		hasError = hasError || (out != test.out)
		if hasError && !test.errExpected {
			t.Errorf("Mismatch on val %#v: expected %#v but got %#v", test.val, test.out, out)
		}
		if !hasError && test.errExpected {
			t.Errorf("Expected error on val %#v: expected %#v, got %#v", test.val, test.out, out)
		}
	}
}

func TestConvertRequestStringToInterface(t *testing.T) {
	tests := []struct {
		val         string
		out         interface{}
		errExpected bool
	}{
		{
			"\"string\"",
			"string",
			false,
		},
		{
			"1",
			float64(1),
			false,
		},
		{
			"{\"t\": true, \"r\": \"e\", \"f\": [\"1\"], \"e\": -4}",
			map[string]interface{}{"e": float64(-4), "f": []interface{}{"1"}, "r": "e", "t": true},
			false,
		},
		{
			"1.5",
			1.5,
			false,
		},
		{
			"[\"1\",\"2\"]",
			[]interface{}{"1", "2"},
			false,
		},
		{
			"true",
			true,
			false,
		},
		{
			"null",
			nil,
			false,
		},
	}

	for _, test := range tests {
		out := requestStringToInterface(test.val)
		hasError := false
		hasError = hasError || !reflect.DeepEqual(out, test.out)
		if hasError && !test.errExpected {
			t.Errorf("Mismatch on val %#v: expected %#v but got %#v", test.val, test.out, out)
		}
		if !hasError && test.errExpected {
			t.Errorf("Expected error on val %#v: expected %#v, got %#v", test.val, test.out, out)
		}
	}
}
