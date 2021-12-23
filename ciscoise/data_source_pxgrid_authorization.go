package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourcePxgridAuthorization() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Provider.

- 🚧 Authorization
`,

		ReadContext: dataSourcePxgridAuthorizationRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePxgridAuthorizationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: Authorization")

		response1, err := client.Provider.Authorization()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing Authorization", err,
				"Failure at Authorization, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting Authorization response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
