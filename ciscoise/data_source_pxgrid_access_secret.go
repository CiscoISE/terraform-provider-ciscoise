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
func dataSourcePxgridAccessSecret() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Consumer.

🚧 AccessSecret`,

		ReadContext: dataSourcePxgridAccessSecretRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourcePxgridAccessSecretRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: AccessSecret")
		request1 := expandRequestPxgridAccessSecretAccessSecret(ctx, "", d)

		response1, err := client.Consumer.AccessSecret(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AccessSecret", err,
				"Failure at AccessSecret, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting AccessSecret response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPxgridAccessSecretAccessSecret(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestClearThreatsAndVulnerabilitiesAccessSecret {
	request := isegosdk.RequestClearThreatsAndVulnerabilitiesAccessSecret{}
	if v, ok := d.GetOkExists(key + ".peer_node_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".peer_node_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".peer_node_name"))) {
		request.PeerNodeName = interfaceToString(v)
	}
	return &request
}
