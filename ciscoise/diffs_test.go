package ciscoise

import "testing"

func TestDiffsDiffSupressHotpatchName(t *testing.T) {
	cases := map[string]struct {
		Old, New           string
		ExpectDiffSuppress bool
	}{
		"same hotpatch name": {
			Old:                "ise-apply-CSCvz53724_3.2.x_patchall-SPA.tar.gz",
			New:                "ise-apply-CSCvz53724_3.2.x_patchall-SPA.tar.gz",
			ExpectDiffSuppress: true,
		},
		"contains hotpatch name 1 ": {
			Old:                "ise-apply-CSCvz53724_3.2.x_patchall-SPA.tar.gz",
			New:                "CSCvz53724_3.2.x_patchall",
			ExpectDiffSuppress: true,
		},
		"contains hotpatch name 2": {
			Old:                "CSCvz53724_3.2.x_patchall",
			New:                "ise-rollback-CSCvz53724_3.2.x_patchall-SPA.tar.gz",
			ExpectDiffSuppress: true,
		},
		"does not contain": {
			Old:                "ise-apply-CSCvz53724_3.2.x_patchall-SPA.tar.gz",
			New:                "ise-rollback-CSCvz53724_3.2.x_patchall-SPA.tar.gz",
			ExpectDiffSuppress: false,
		},
	}
	for tn, tc := range cases {
		if diffSupressHotpatchName()("key", tc.Old, tc.New, nil) != tc.ExpectDiffSuppress {
			t.Errorf("bad: %s, '%s' => '%s' expect DiffSuppress to return %t", tn, tc.Old, tc.New, tc.ExpectDiffSuppress)
		}
	}
}

func TestDiffsDiffSupressMacAddress(t *testing.T) {
	cases := map[string]struct {
		Old, New           string
		ExpectDiffSuppress bool
	}{
		"same mac address, F8-40-3E-57-2C-75 vs F8403E572C75": {
			Old:                "F8-40-3E-57-2C-75",
			New:                "F8403E572C75",
			ExpectDiffSuppress: true,
		},
		"same mac address, F8.40.3E.57.2C.75 vs F8403E572C75": {
			Old:                "F8.40.3E.57.2C.75",
			New:                "F8403E572C75",
			ExpectDiffSuppress: true,
		},
		"same mac address, F8-40-3E-57-2C-75 vs F8.40.3E.57.2C.75": {
			Old:                "F8-40-3E-57-2C-75",
			New:                "F8.40.3E.57.2C.75",
			ExpectDiffSuppress: true,
		},
		"same mac address, f8-40-3e-57-2c-75 vs F8.40.3E.57.2C.75": {
			Old:                "f8-40-3e-57-2c-75",
			New:                "F8.40.3E.57.2C.75",
			ExpectDiffSuppress: true,
		},
		"same mac address, F8:40:3E:57:2C:75 vs F8.40.3E.57.2C.75": {
			Old:                "F8:40:3E:57:2C:75",
			New:                "F8.40.3E.57.2C.75",
			ExpectDiffSuppress: true,
		},
		"same mac address, F8:40:3E:57:2C:75 vs F8403E572C75": {
			Old:                "F8:40:3E:57:2C:75",
			New:                "F8403E572C75",
			ExpectDiffSuppress: true,
		},
		"different mac address, F8-40-3E-57-2C-75 vs 91-58-C8-5A-FB-B1": {
			Old:                "F8-40-3E-57-2C-75",
			New:                "91-58-C8-5A-FB-B1",
			ExpectDiffSuppress: false,
		},
	}
	for tn, tc := range cases {
		if diffSupressMacAddress()("key", tc.Old, tc.New, nil) != tc.ExpectDiffSuppress {
			t.Errorf("bad: %s, '%s' => '%s' expect DiffSuppress to return %t", tn, tc.Old, tc.New, tc.ExpectDiffSuppress)
		}
	}
}

func TestDiffsDiffSuppressSgt(t *testing.T) {
	cases := map[string]struct {
		Old, New           string
		ExpectDiffSuppress bool
	}{
		"same sgt, Auditors vs Auditors": {
			Old:                "Auditors",
			New:                "Auditors",
			ExpectDiffSuppress: true,
		},
		"different sgt, Auditors vs auditors": {
			Old:                "Auditors",
			New:                "auditors",
			ExpectDiffSuppress: false,
		},
		"same sgt, Auditors vs Auditors (16)": {
			Old:                "Auditors",
			New:                "Auditors (16)",
			ExpectDiffSuppress: true,
		},
		"different sgt, Auditors vs auditors (16)": {
			Old:                "Auditors",
			New:                "auditors (16)",
			ExpectDiffSuppress: false,
		},
	}
	for tn, tc := range cases {
		if diffSuppressSgt()("key", tc.Old, tc.New, nil) != tc.ExpectDiffSuppress {
			t.Errorf("bad: %s, '%s' => '%s' expect DiffSuppress to return %t", tn, tc.Old, tc.New, tc.ExpectDiffSuppress)
		}
	}
}

func TestDiffsCaseInsensitive(t *testing.T) {
	cases := map[string]struct {
		Old, New           string
		ExpectDiffSuppress bool
	}{
		"same strings": {
			Old:                "hel23lo",
			New:                "hel23lo",
			ExpectDiffSuppress: true,
		},
		"same strings ignoring case": {
			Old:                "hEl3_á= lo",
			New:                "HeL3_Á= LO",
			ExpectDiffSuppress: true,
		},
		"different strings": {
			Old:                "hel24lo",
			New:                "hel23lo",
			ExpectDiffSuppress: false,
		},
		"different strings ignoring case": {
			Old:                "hEl= lo",
			New:                "HeL3_= LO",
			ExpectDiffSuppress: false,
		},
		"different strings, spaces": {
			Old:                "hel23lo ",
			New:                "hel23lo",
			ExpectDiffSuppress: false,
		},
	}
	for tn, tc := range cases {
		if caseInsensitive()("key", tc.Old, tc.New, nil) != tc.ExpectDiffSuppress {
			t.Errorf("bad: %s, '%s' => '%s' expect DiffSuppress to return %t", tn, tc.Old, tc.New, tc.ExpectDiffSuppress)
		}
	}
}

func TestDiffsDiffSuppressBooleans(t *testing.T) {
	cases := map[string]struct {
		Old, New           string
		ExpectDiffSuppress bool
	}{
		"same bool True, equivalent": {
			Old:                "true",
			New:                "on",
			ExpectDiffSuppress: true,
		},
		"same bool False, equivalent": {
			Old:                "off",
			New:                "false",
			ExpectDiffSuppress: true,
		},
		"same bool True, true": {
			Old:                "true",
			New:                "true",
			ExpectDiffSuppress: true,
		},
		"same bool True, on": {
			Old:                "on",
			New:                "on",
			ExpectDiffSuppress: true,
		},
		"same bool False, false": {
			Old:                "false",
			New:                "false",
			ExpectDiffSuppress: true,
		},
		"same bool False, off": {
			Old:                "off",
			New:                "off",
			ExpectDiffSuppress: true,
		},
		"different bool off": {
			Old:                "off",
			New:                "on",
			ExpectDiffSuppress: false,
		},
		"different bool false": {
			Old:                "true",
			New:                "false",
			ExpectDiffSuppress: false,
		},
		"different bool true": {
			Old:                "true",
			New:                "off",
			ExpectDiffSuppress: false,
		},
		"different bool on": {
			Old:                "on",
			New:                "false",
			ExpectDiffSuppress: false,
		},
	}
	for tn, tc := range cases {
		if diffSuppressBooleans()("key", tc.Old, tc.New, nil) != tc.ExpectDiffSuppress {
			t.Errorf("bad: %s, '%s' => '%s' expect DiffSuppress to return %t", tn, tc.Old, tc.New, tc.ExpectDiffSuppress)
		}
	}
}
