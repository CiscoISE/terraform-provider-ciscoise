package ciscoise

import (
	"reflect"
	"testing"
)

func TestUtilsCompareMacAddress(t *testing.T) {
	cases := map[string]struct {
		Old, New     string
		ExpectResult bool
	}{
		"same mac address, F8-40-3E-57-2C-75 vs F8403E572C75": {
			Old:          "F8-40-3E-57-2C-75",
			New:          "F8403E572C75",
			ExpectResult: true,
		},
		"same mac address, F8.40.3E.57.2C.75 vs F8403E572C75": {
			Old:          "F8.40.3E.57.2C.75",
			New:          "F8403E572C75",
			ExpectResult: true,
		},
		"same mac address, F8-40-3E-57-2C-75 vs F8.40.3E.57.2C.75": {
			Old:          "F8-40-3E-57-2C-75",
			New:          "F8.40.3E.57.2C.75",
			ExpectResult: true,
		},
		"same mac address, f8-40-3e-57-2c-75 vs F8.40.3E.57.2C.75": {
			Old:          "f8-40-3e-57-2c-75",
			New:          "F8.40.3E.57.2C.75",
			ExpectResult: true,
		},
		"same mac address, F8:40:3E:57:2C:75 vs F8.40.3E.57.2C.75": {
			Old:          "F8:40:3E:57:2C:75",
			New:          "F8.40.3E.57.2C.75",
			ExpectResult: true,
		},
		"same mac address, F8:40:3E:57:2C:75 vs F8403E572C75": {
			Old:          "F8:40:3E:57:2C:75",
			New:          "F8403E572C75",
			ExpectResult: true,
		},
		"different mac address, F8-40-3E-57-2C-75 vs 91-58-C8-5A-FB-B1": {
			Old:          "F8-40-3E-57-2C-75",
			New:          "91-58-C8-5A-FB-B1",
			ExpectResult: false,
		},
	}
	for tn, tc := range cases {
		if compareMacAddress(tc.Old, tc.New) != tc.ExpectResult {
			t.Errorf("bad: %s, '%s' => '%s' expect compareMacAddress to return %t", tn, tc.Old, tc.New, tc.ExpectResult)
		}
	}
}

func TestUtilsCompareSGT(t *testing.T) {
	cases := map[string]struct {
		Old, New     string
		ExpectResult bool
	}{
		"same sgt, Auditors vs Auditors": {
			Old:          "Auditors",
			New:          "Auditors",
			ExpectResult: true,
		},
		"different sgt, Auditors vs auditors": {
			Old:          "Auditors",
			New:          "auditors",
			ExpectResult: false,
		},
		"same sgt, Auditors vs Auditors (16)": {
			Old:          "Auditors",
			New:          "Auditors (16)",
			ExpectResult: true,
		},
		"different sgt, Auditors vs auditors (16)": {
			Old:          "Auditors",
			New:          "auditors (16)",
			ExpectResult: false,
		},
	}
	for tn, tc := range cases {
		if compareSGT(tc.Old, tc.New) != tc.ExpectResult {
			t.Errorf("bad: %s, '%s' => '%s' expect compareSGT to return %t", tn, tc.Old, tc.New, tc.ExpectResult)
		}
	}
}

func TestUtilsCaseInsensitive(t *testing.T) {
	cases := map[string]struct {
		Old, New     string
		ExpectResult bool
	}{
		"trim leading .": {
			Old:          ".string",
			New:          "string",
			ExpectResult: true,
		},
		"trim trailing .": {
			Old:          "string.",
			New:          "string",
			ExpectResult: true,
		},
		"trim leading and trailing .": {
			Old:          ".string.",
			New:          "string",
			ExpectResult: true,
		},
		"trim many leading and trailing .": {
			Old:          "...string..",
			New:          "string",
			ExpectResult: true,
		},
		"different strings, spaces": {
			Old:          "..string ",
			New:          "string",
			ExpectResult: false,
		},
	}
	for tn, tc := range cases {
		if (fixKeyAccess(tc.Old) == tc.New) != tc.ExpectResult {
			t.Errorf("bad: %s, fixKeyAccess('%s') == '%s' expect fixKeyAccess comparisson to return %t", tn, tc.Old, tc.New, tc.ExpectResult)
		}
	}
}

func TestUtilsGetLocationID(t *testing.T) {
	cases := map[string]struct {
		Value        string
		ExpectResult string
	}{
		"location with multiple /": {
			Value:        "1/2/30",
			ExpectResult: "30",
		},
		"location with one /": {
			Value:        "/3",
			ExpectResult: "3",
		},
		"location with none /": {
			Value:        "3",
			ExpectResult: "3",
		},
		"different location with one /": {
			Value:        "3/",
			ExpectResult: "",
		},
	}
	for tn, tc := range cases {
		if getLocationID(tc.Value) != tc.ExpectResult {
			t.Errorf("bad: %s, '%s' expect getLocationID to return %s", tn, tc.Value, tc.ExpectResult)
		}
	}
}

func TestUtilsListNicely(t *testing.T) {
	cases := map[string]struct {
		Old          []string
		New          string
		ExpectResult bool
	}{
		"list with len 0": {
			Old:          []string{},
			New:          "",
			ExpectResult: true,
		},
		"list with len 1": {
			Old:          []string{"1"},
			New:          "\"1\"",
			ExpectResult: true,
		},
		"list with len 2": {
			Old:          []string{"1", "2 3"},
			New:          "\"1\", \"2 3\"",
			ExpectResult: true,
		},
	}
	for tn, tc := range cases {
		if (listNicely(tc.Old) == tc.New) != tc.ExpectResult {
			t.Errorf("bad: %s, '%s' == '%s' expect listNicely comparisson to return %t", tn, listNicely(tc.Old), tc.New, tc.ExpectResult)
		}
	}
}

func TestUtilsPickMethod(t *testing.T) {
	cases := map[string]struct {
		Old          [][]bool
		New          int
		ExpectResult bool
	}{
		"list with len 0": {
			Old:          [][]bool{},
			New:          1,
			ExpectResult: true,
		},
		"list with len 1 and true": {
			Old:          [][]bool{{true}},
			New:          1,
			ExpectResult: true,
		},
		"list with len 1 and false": {
			Old:          [][]bool{{false}},
			New:          1,
			ExpectResult: true,
		},
		"list with len 2. Each of len 1": {
			Old:          [][]bool{{false}, {true}},
			New:          2,
			ExpectResult: true,
		},
		"list with len 3. Each of len 2": {
			Old:          [][]bool{{false, true}, {true, true}, {false, true}},
			New:          2,
			ExpectResult: true,
		},
		"list with len 3. Each of len 2 [1]": {
			Old:          [][]bool{{false, true}, {true, true}, {true, true}},
			New:          2,
			ExpectResult: true,
		},
		"list with len 3. Descending lens from 2 to 4 almost true": {
			Old:          [][]bool{{false, true}, {false, true, true}, {false, true, true, true}},
			New:          3,
			ExpectResult: true,
		},
	}
	for tn, tc := range cases {
		if (pickMethod(tc.Old) == tc.New) != tc.ExpectResult {
			t.Errorf("bad: %s, '%d' == '%d' expect pickMethod comparisson to return %t", tn, pickMethod(tc.Old), tc.New, tc.ExpectResult)
		}
	}
}

func TestUtilsGetNextPageAndSizeParams(t *testing.T) {
	result_to_list_interface := func(page int, size int, err error) []interface{} {
		if err != nil {
			return []interface{}{page, size, err.Error()}
		} else {
			return []interface{}{page, size, ""}
		}
	}
	cases := map[string]struct {
		Value        string
		ExpectResult []interface{}
	}{
		"url: dev?page=3&size=2": {
			Value:        "dev?page=3&size=2",
			ExpectResult: []interface{}{3, 2, ""},
		},
		"url: https://ex.dev?page=1&size=2": {
			Value:        "https://ex.dev?page=1&size=2",
			ExpectResult: []interface{}{1, 2, ""},
		},
		"url: dev?page=3&size=3&size=2": {
			Value:        "dev?page=3&size=3&size=2",
			ExpectResult: []interface{}{3, 3, ""},
		},
		"url: dev?size=3&page=4": {
			Value:        "dev?size=3&page=4",
			ExpectResult: []interface{}{4, 3, ""},
		},
		"url: dev?size=3": {
			Value:        "dev?size=3",
			ExpectResult: []interface{}{0, 3, ""},
		},
		"url: dev?page=3": {
			Value:        "dev?page=3",
			ExpectResult: []interface{}{3, 0, ""},
		},
	}
	for tn, tc := range cases {
		v := result_to_list_interface(getNextPageAndSizeParams(tc.Value))
		if !reflect.DeepEqual(v, tc.ExpectResult) {
			t.Errorf("bad: %s, %#v expect to return %#v", tn, v, tc.ExpectResult)
		}
	}
}

func TestUtilsRemoveParameters(t *testing.T) {
	cases := map[string]struct {
		Value1       []map[string]interface{}
		Value2       []string
		ExpectResult []map[string]interface{}
	}{
		"map 0": {
			Value1:       []map[string]interface{}{{"test": 1, "url": nil, "op": "string"}},
			Value2:       []string{},
			ExpectResult: []map[string]interface{}{{"test": 1, "url": nil, "op": "string"}},
		},
		"map 1": {
			Value1:       []map[string]interface{}{{"test": 1, "url": nil, "op": "string"}},
			Value2:       []string{"yest"},
			ExpectResult: []map[string]interface{}{{"test": 1, "url": nil, "op": "string"}},
		},
		"map 2": {
			Value1:       []map[string]interface{}{{"test": 1, "url": nil, "op": "string"}},
			Value2:       []string{"url"},
			ExpectResult: []map[string]interface{}{{"test": 1, "op": "string"}},
		},
		"map 3": {
			Value1:       []map[string]interface{}{{"test": 1, "url": nil, "op": "string"}},
			Value2:       []string{"url", "op"},
			ExpectResult: []map[string]interface{}{{"test": 1}},
		},
		"map 4": {
			Value1:       []map[string]interface{}{{"test": 1, "url": nil, "op": "string"}},
			Value2:       []string{"url", "op", "test"},
			ExpectResult: []map[string]interface{}{{}},
		},
		"map 5": {
			Value1:       []map[string]interface{}{{"test": 1, "url": nil, "op": "string"}},
			Value2:       []string{"url", "op", "test", "other"},
			ExpectResult: []map[string]interface{}{{}},
		},
	}
	for tn, tc := range cases {
		v := remove_parameters(tc.Value1, tc.Value2...)
		if !reflect.DeepEqual(v, tc.ExpectResult) {
			t.Errorf("bad: %s, %#v expect to return %#v", tn, v, tc.ExpectResult)
		}
	}
}
