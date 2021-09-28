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
func dataSourcePxgridServiceLookup() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Consumer.

- 🚧 ServiceLookup`,

		ReadContext: dataSourcePxgridServiceLookupRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourcePxgridServiceLookupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: LookupService")
		request1 := expandRequestPxgridServiceLookupLookupService(ctx, "", d)

		response1, err := client.Consumer.LookupService(request1)

		if err != nil || response1 == nil {
			if request1 != nil {
				log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing LookupService", err,
				"Failure at LookupService, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LookupService response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPxgridServiceLookupLookupService(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestClearThreatsAndVulnerabilitiesLookupService {
	request := isegosdk.RequestClearThreatsAndVulnerabilitiesLookupService{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
