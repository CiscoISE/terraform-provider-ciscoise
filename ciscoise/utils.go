package ciscoise

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// listNicely listNicely
/* Converts []string to string, by adding quotes and separate values by comma
@param values
*/
func listNicely(values []string) string {
	pvalues := fmt.Sprintf("%q", values)
	pvalues = pvalues[1 : len(pvalues)-1]
	return strings.Join(strings.Split(pvalues, " "), ", ")
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
