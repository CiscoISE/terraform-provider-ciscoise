package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTransportGatewaySettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on telemetry.

- Transport Gateway acts a proxy for the communication between the ISE servers in your network and the Telemetry servers
in case of air-gapped network.
`,

		ReadContext: dataSourceTransportGatewaySettingsRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enable_transport_gateway": &schema.Schema{
							Description: `Indicates whether transport gateway is enabled or not.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"url": &schema.Schema{
							Description: `URL of transport gateway`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTransportGatewaySettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTransportGateway")

		response1, restyResp1, err := client.Telemetry.GetTransportGateway()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTransportGateway", err,
				"Failure at GetTransportGateway, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTelemetryGetTransportGatewayItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTransportGateway response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTelemetryGetTransportGatewayItem(item *isegosdk.ResponseTelemetryGetTransportGatewayResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enable_transport_gateway"] = boolPtrToString(item.EnableTransportGateway)
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
