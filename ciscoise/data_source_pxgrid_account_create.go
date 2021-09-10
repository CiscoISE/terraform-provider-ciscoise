package ciscoise

import (
	"context"

	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourcePxgridAccountCreate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePxgridAccountCreateRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourcePxgridAccountCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateAccount")
		request1 := expandRequestPxgridAccountCreateCreateAccount(ctx, "", d)

		response1, err := client.Consumer.CreateAccount(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateAccount", err,
				"Failure at CreateAccount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateAccount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPxgridAccountCreateCreateAccount(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestClearThreatsAndVulnerabilitiesCreateAccount {
	request := isegosdk.RequestClearThreatsAndVulnerabilitiesCreateAccount{}
	if v, ok := d.GetOkExists(key + ".node_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".node_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".node_name"))) {
		request.NodeName = interfaceToString(v)
	}
	return &request
}
