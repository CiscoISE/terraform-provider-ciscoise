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
func dataSourceThreatVulnerabilitiesClear() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ClearThreatsAndVulnerabilities.

- This data source action allows the client to delete the ThreatContext and Threat events that are associated with the
given MAC Address.`,

		ReadContext: dataSourceThreatVulnerabilitiesClearRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mac_addresses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceThreatVulnerabilitiesClearRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ClearThreatsAndVulnerabilities")
		request1 := expandRequestThreatVulnerabilitiesClearClearThreatsAndVulnerabilities(ctx, "", d)

		response1, err := client.ClearThreatsAndVulnerabilities.ClearThreatsAndVulnerabilities(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ClearThreatsAndVulnerabilities", err,
				"Failure at ClearThreatsAndVulnerabilities, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ClearThreatsAndVulnerabilities response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestThreatVulnerabilitiesClearClearThreatsAndVulnerabilities(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilities {
	request := isegosdk.RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilities{}
	request.ERSIrfThreatContext = expandRequestThreatVulnerabilitiesClearClearThreatsAndVulnerabilitiesERSIrfThreatContext(ctx, key, d)
	return &request
}

func expandRequestThreatVulnerabilitiesClearClearThreatsAndVulnerabilitiesERSIrfThreatContext(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilitiesERSIrfThreatContext {
	request := isegosdk.RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilitiesERSIrfThreatContext{}
	if v, ok := d.GetOkExists("mac_addresses"); !isEmptyValue(reflect.ValueOf(d.Get("mac_addresses"))) && (ok || !reflect.DeepEqual(v, d.Get("mac_addresses"))) {
		request.MacAddresses = interfaceToString(v)
	}
	return &request
}
