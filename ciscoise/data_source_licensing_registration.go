package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicensingRegistration() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licensing.

- Get registration information
`,

		ReadContext: dataSourceLicensingRegistrationRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connection_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"registration_state": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssm_on_prem_server": &schema.Schema{
							Description: `If connection type is selected as SSM_ONPREM_SERVER, then  IP address or the hostname (or FQDN) of the SSM On-Prem server Host.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tier": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceLicensingRegistrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetRegistrationInfo")

		response1, restyResp1, err := client.Licensing.GetRegistrationInfo()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRegistrationInfo", err,
				"Failure at GetRegistrationInfo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensingGetRegistrationInfoItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRegistrationInfo response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensingGetRegistrationInfoItem(item *isegosdk.ResponseLicensingGetRegistrationInfoResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connection_type"] = item.ConnectionType
	respItem["registration_state"] = item.RegistrationState
	respItem["ssm_on_prem_server"] = item.SsmOnPremServer
	respItem["tier"] = item.Tier
	return []map[string]interface{}{
		respItem,
	}
}
