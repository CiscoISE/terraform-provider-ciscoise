package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourcePxgridAccountActivate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Consumer.

🚧 AccountActivate`,

		ReadContext: dataSourcePxgridAccountActivateRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePxgridAccountActivateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ActivateAccount")
		request1 := expandRequestPxgridAccountActivateActivateAccount(ctx, "", d)

		response1, err := client.Consumer.ActivateAccount(request1)

		if err != nil || response1 == nil {
			if request1 != nil {
				log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ActivateAccount", err,
				"Failure at ActivateAccount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ActivateAccount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPxgridAccountActivateActivateAccount(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestClearThreatsAndVulnerabilitiesActivateAccount {
	request := isegosdk.RequestClearThreatsAndVulnerabilitiesActivateAccount{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	return &request
}
