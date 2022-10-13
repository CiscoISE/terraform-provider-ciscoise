package ciscoise

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func compareHotpatchName(old, new string) bool {
	return len(old) > 0 && len(new) > 0 && (strings.Contains(old, new) || strings.Contains(new, old))
}

func remove_parameters(respItems []map[string]interface{}, parameters ...string) []map[string]interface{} {
	for i := range respItems {
		for j, element := range respItems[i] {
			for _, parameter := range parameters {
				if element != nil && strings.Contains(reflect.TypeOf(element).String(), "[]map[") {
					paramMap := element.([]map[string]interface{})
					respItems[i][j] = remove_parameters(paramMap, parameter)
				} else if element != nil && strings.Contains(reflect.TypeOf(element).String(), "map[") {
					var paramMap []map[string]interface{}
					paramMap = append(paramMap, element.(map[string]interface{}))
					respItems[i][j] = remove_parameters(paramMap, parameter)
				} else {
					for k := range parameters {
						_, ok := respItems[i][parameters[k]]
						if ok {
							delete(respItems[i], parameters[k])
						}
					}
				}
			}
		}
	}
	return respItems
}

func compareMacAddress(old_mac_address, new_mac_address string) bool {
	rexp := `([-.:])`
	oldClear, newClear := replaceRegExStrings(old_mac_address, new_mac_address, rexp, "")
	return strings.ToLower(oldClear) == strings.ToLower(newClear)
}

func compareBoolean(new string, old string) bool {
	if old == "" && new != "" {
		return false
	}
	oldBool, err := strconv.ParseBool(old)
	if err != nil {
		return true
	}
	newBool, err := strconv.ParseBool(new)
	if err != nil {
		return true
	}

	if oldBool == newBool {
		return true
	}
	return false
}

func compareOptional(new_value string, old_value string) bool {
	return new_value == "" || new_value == old_value
}
func fixKeyAccess(key string) string {
	return strings.Trim(key, ".")
}

func compareSGT(first_sgt, second_sgt string) bool {
	rexp := `\s*\(.*\)$`
	oldClear, newClear := replaceRegExStrings(first_sgt, second_sgt, rexp, "")
	return oldClear == newClear
}

func replaceRegExStrings(first_str, second_str, regex_src, sub_repl string) (string, string) {
	m1 := regexp.MustCompile(regex_src)
	first_repl := m1.ReplaceAllString(first_str, sub_repl)
	second_repl := m1.ReplaceAllString(second_str, sub_repl)
	return first_repl, second_repl
}

func replaceAllStr(original_str string, old string, new string) string {
	return strings.ReplaceAll(original_str, old, new)
}

func getLocationID(location string) string {
	var locationID string
	URL_SEPARATOR := "/"
	locationFragments := strings.Split(location, URL_SEPARATOR)
	if len(locationFragments) > 0 {
		return locationFragments[len(locationFragments)-1]
	}
	return locationID
}

func isEmptyValue(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Struct:
		return v.IsZero()
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func joinResourceID(result_params map[string]string) string {
	var PARAMS_SEPARATOR string = "\\"
	var PARAM_VALUE_SEPARATOR string = ":="
	ID := ""
	params := []string{}
	for key, value := range result_params {
		if value != "" {
			params = append(params, fmt.Sprintf("%s%s%s", key, PARAM_VALUE_SEPARATOR, value))
		}
	}
	sort.Strings(params) // Sort params
	ID = strings.Join(params, PARAMS_SEPARATOR)
	return ID
}

func separateResourceID(ID string) map[string]string {
	var PARAMS_SEPARATOR string = "\\"
	var PARAM_VALUE_SEPARATOR string = ":="
	params := strings.Split(ID, PARAMS_SEPARATOR)
	sort.Strings(params) // Sort params
	result_params := make(map[string]string)
	for _, param := range params {
		param_key_value := strings.Split(param, PARAM_VALUE_SEPARATOR)
		if len(param_key_value) == 2 {
			if param_key_value[1] != "" {
				result_params[param_key_value[0]] = param_key_value[1]
			}
		}
	}
	return result_params
}

// listNicely listNicely
/* Converts []string to string, by adding quotes and separate values by comma
@param values
*/
func listNicely(values []string) string {
	pvalues := []string{}
	for i := range values {
		pvalues = append(pvalues, fmt.Sprintf("\"%s\"", values[i]))
	}
	return strings.Join(pvalues, ", ")
}

func pickMethodAux(method []bool) float64 {
	lenM := len(method)
	countM := 0
	for _, em := range method {
		if em {
			countM += 1
		}
	}
	var percentM float64 = float64(countM) / float64(lenM)
	return percentM
}

func pickMethod(methods [][]bool) int {
	methodN := 0
	maxPercentM := 0.0
	for i, method := range methods {
		percentM := pickMethodAux(method)
		if maxPercentM < percentM {
			methodN = i
			maxPercentM = percentM
		}
	}
	// Add 1 to match number method and not index
	return methodN + 1
}

func diagError(summaryErr string, err error) diag.Diagnostic {
	diagErrResponse := diag.Diagnostic{Severity: diag.Error, Summary: summaryErr}
	if err != nil {
		diagErrResponse.Detail = err.Error()
		return diagErrResponse
	}
	return diagErrResponse
}

func diagErrorWithResponse(summaryErr string, err error, restyResponse string) diag.Diagnostic {
	diagErrResponse := diag.Diagnostic{Severity: diag.Error, Summary: summaryErr}
	if err != nil {
		diagErrResponse.Detail = fmt.Sprintf("%s\n%v", err.Error(), restyResponse)
		return diagErrResponse
	}
	diagErrResponse.Detail = restyResponse
	return diagErrResponse
}

func diagErrorWithAltAndResponse(summaryErr string, err error, restyResponse string, summaryAlt string, detail string) diag.Diagnostic {
	diagErrResponse := diag.Diagnostic{Severity: diag.Error}
	if err != nil {
		diagErrResponse.Summary = summaryErr
		diagErrResponse.Detail = fmt.Sprintf("%s\n%v", err.Error(), restyResponse)
		return diagErrResponse
	}
	diagErrResponse.Summary = summaryAlt
	if detail != "" {
		diagErrResponse.Detail = detail
	}
	return diagErrResponse
}

func diagErrorWithAlt(summaryErr string, err error, summaryAlt string, detail string) diag.Diagnostic {
	diagErrResponse := diag.Diagnostic{Severity: diag.Error}
	if err != nil {
		diagErrResponse.Summary = summaryErr
		diagErrResponse.Detail = err.Error()
		return diagErrResponse
	}
	diagErrResponse.Summary = summaryAlt
	if detail != "" {
		diagErrResponse.Detail = detail
	}
	return diagErrResponse
}

func getUnixTimeString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func getNextPageAndSizeParams(href string) (page int, size int, err error) {
	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(href)
	if err != nil {
		return
	}

	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return
	}
	var pageStr string
	var sizeStr string
	if v, ok := m["page"]; ok && len(v) > 0 {
		pageStr = v[0]
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return
		}
	}
	if v, ok := m["size"]; ok && len(v) > 0 {
		sizeStr = v[0]
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			return
		}
	}

	return
}
