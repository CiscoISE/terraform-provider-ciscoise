package ciscoise

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func diffSupressHotpatchName() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		return compareHotpatchName(old, new)
	}
}

func diffSupressOptional() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		return compareOptional(new, old)
	}
}
func diffSupressBool() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		return compareBoolean(new, old)
	}
}

func diffSupressMacAddress() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		return compareMacAddress(old, new)
	}
}

func diffSuppressSgt() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		return compareSGT(old, new)
	}
}

func diffSuppressAlways() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		return true
	}
}

func caseInsensitive() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		return strings.EqualFold(old, new)
	}
}

func diffSuppressBooleans() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		if old == "off" {
			return old == new || "false" == new
		}
		if old == "false" {
			return old == new || "off" == new
		}
		if old == "on" {
			return old == new || "true" == new
		}
		if old == "true" {
			return old == new || "on" == new
		}
		return true
	}
}
