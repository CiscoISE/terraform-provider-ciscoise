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

func pickMethod(m1 []bool, m2 []bool) int {
	lenM1 := len(m1) + 1 // Start with 1 to avoid divide by zero
	lenM2 := len(m2) + 1
	priorityM1 := lenM1 <= lenM2 // Give priority to m1
	countM1 := 1                 // Start with 1 to avoid divide by zero
	countM2 := 1
	for _, em1 := range m1 {
		if em1 {
			countM1 += 1
		}
	}
	for _, em2 := range m2 {
		if em2 {
			countM2 += 1
		}
	}
	var percentM1 float64 = float64(countM1) / float64(lenM1)
	var percentM2 float64 = float64(countM2) / float64(lenM2)
	if percentM1 == percentM2 {
		if priorityM1 {
			return 1
		}
		return 2
	}
	if percentM1 < percentM2 {
		return 2
	}
	return 1
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
