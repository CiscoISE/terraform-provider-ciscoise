package ciscoise

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"ciscoise": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("ISE_BASE_URL"); err == "" {
		t.Fatal("ISE_BASE_URL must be set for acceptance tests")
	}
	if err := os.Getenv("ISE_USERNAME"); err == "" {
		t.Fatal("ISE_USERNAME must be set for acceptance tests")
	}
	if err := os.Getenv("ISE_PASSWORD"); err == "" {
		t.Fatal("ISE_PASSWORD must be set for acceptance tests")
	}
	if err := os.Getenv("ISE_SSL_VERIFY"); err == "" {
		t.Fatal("ISE_SSL_VERIFY must be set for acceptance tests")
	}
	if err := os.Getenv("ISE_USE_API_GATEWAY"); err == "" {
		t.Fatal("ISE_USE_API_GATEWAY must be set for acceptance tests")
	}
}
