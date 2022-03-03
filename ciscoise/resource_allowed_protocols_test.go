package ciscoise

import (
	"fmt"
	"os"
	"strings"

	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformCiscoISEAllowedProtocolsExample(t *testing.T) {
	if v := os.Getenv("TF_ACC"); v != "1" {
		t.Skip("TF_ACC not enabled")
	}
	t.Parallel()

	number := random.RandomInt([]int{1, 2, 3, 4})
	name := strings.Title(fmt.Sprintf("New Allowed Protocol %d", number))
	description := strings.ToLower(fmt.Sprintf("New Allowed Protocol Service %d", number))
	descriptionUpdate := strings.ToLower(fmt.Sprintf("New-Allowed-Protocol-Service-%d", number))

	// website::tag::1:: Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/samples/resources/ciscoise_allowed_protocols",
		Vars: map[string]interface{}{
			"name":          name,
			"description":   description,
			"allow_eap_tls": "false",
		},
	}
	terraformOptionsUpdate := &terraform.Options{
		TerraformDir: "../examples/samples/resources/ciscoise_allowed_protocols",
		Vars: map[string]interface{}{
			"name":          name,
			"description":   descriptionUpdate,
			"allow_eap_tls": "true",
		},
	}

	// website::tag::5:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::2:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// website::tag::3:: Run `terraform output` to get the values of output variables
	itemDescription := terraform.Output(t, terraformOptions, "ciscoise_allowed_protocols_response_item_description")
	ID := terraform.Output(t, terraformOptions, "ciscoise_allowed_protocols_response_id")
	item := terraform.OutputListOfObjects(t, terraformOptions, "ciscoise_allowed_protocols_response_item")

	/* Alternatives to Test UpdateContext

	Alternative 1
		Test using only the resource definition.

		// Requires double to get the value from here using direct output. [Unexpected behavior]
		terraform.Apply(t, terraformOptionsUpdate)
		terraform.Apply(t, terraformOptionsUpdate)

		var itemDescriptionUpdate string

		// Alternative 1.A
		// Get value directly from output with Tf conditionals [https://www.terraform.io/language/expressions/conditionals]
		// and function length https://www.terraform.io/language/functions/length
		itemDescriptionUpdate = terraform.Output(t, terraformOptionsUpdate, "ciscoise_allowed_protocols_response_item_description")

		// Alternative 1.B
		// Get value using terratest/modules/terraform functions and in-site conversions
		response := terraform.OutputListOfObjects(t, terraformOptionsUpdate, "ciscoise_allowed_protocols_response_item")
		if len(response) > 0 {
			itemDescriptionUpdate = response[0]["description"].(string)
		}

	Alternative 2
		Test using resource and data source definition.

		terraform.Apply(t, terraformOptionsUpdate)

		var itemDescriptionUpdate string

		// Alternative 2.A
		// Get value directly from output with Tf conditionals [https://www.terraform.io/language/expressions/conditionals]
		// and function length https://www.terraform.io/language/functions/length
		itemDescriptionUpdate = terraform.Output(t, terraformOptionsUpdate, "ciscoise_allowed_protocols_response_item_description_datasource")

		// Alternative 2.B
		// Get value using terratest/modules/terraform functions and in-site conversions
		response := terraform.OutputListOfObjects(t, terraformOptionsUpdate, "ciscoise_allowed_protocols_response_item_datasource")
		if len(response) > 0 {
			itemDescriptionUpdate = response[0]["description"].(string)
		}
	*/

	// website::tag::2:: Run `terraform apply`. Fail the test if there are any errors.
	terraform.Apply(t, terraformOptionsUpdate)

	// website::tag::3:: Run `terraform output` to get the values of output variables
	var itemDescriptionUpdate string
	var itemEapTlsInterface map[string]interface{}
	response := terraform.OutputListOfObjects(t, terraformOptionsUpdate, "ciscoise_allowed_protocols_response_item_datasource")
	if len(response) > 0 {
		itemDescriptionUpdate = response[0]["description"].(string)
		// Using https://www.terraform.io/language/expressions/dynamic-blocks to add condition to block
		eap_tls_interface, _ := response[0]["eap_tls"]
		// t.Logf("eap_tls_interface ??? %#v", eap_tls_interface)
		v, ok := eap_tls_interface.([]map[string]interface{})
		if ok && len(v) > 0 {
			itemEapTlsInterface = v[0]
		}
	}

	// website::tag::4:: Assert
	assert := assert.New(t)
	assert.Contains(ID, fmt.Sprintf("name:=%s", name), "[ERR 1]")
	assert.Contains(ID, "id:=", "[ERR 2]")
	assert.NotEmpty(item, "[ERR 3]")
	assert.Equal(description, itemDescription, "[ERR 4]")
	assert.Equal(descriptionUpdate, itemDescriptionUpdate, "[ERR 5]")
	assert.NotEmpty(itemEapTlsInterface, "[ERR 6]")
	assert.Equal(itemEapTlsInterface["allow_eap_tls_auth_of_expired_certs"], "false", "[ERR 7]")
}
