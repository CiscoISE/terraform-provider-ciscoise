package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePanHa() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on PAN HA.

- In a high availability configuration, the primary PAN is in active state. The secondary PAN (backup PAN) is in standby
state, which means that it receives all the configuration updates from the primary PAN, but is not active in the Cisco
ISE cluster. You can configure Cisco ISE to automatically promote the secondary PAN when the primary PAN becomes
unavailable.
`,

		ReadContext: dataSourcePanHaRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"failed_attempts": &schema.Schema{
							Description: `Failover occurs if the primary PAN is down for the specified number of failure polls. Count (2 - 60).<br> The default value is 5. `,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"is_enabled": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"polling_interval": &schema.Schema{
							Description: `Administration nodes are checked after each interval. Seconds (30 - 300) <br> The default value is 120. `,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"primary_health_check_node": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"secondary_health_check_node": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourcePanHaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetPanHaStatus")

		response1, restyResp1, err := client.PanHa.GetPanHaStatus()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPanHaStatus", err,
				"Failure at GetPanHaStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPanHaGetPanHaStatusItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPanHaStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPanHaGetPanHaStatusItem(item *isegosdk.ResponsePanHaGetPanHaStatusResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["failed_attempts"] = item.FailedAttempts
	respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
	respItem["polling_interval"] = item.PollingInterval
	respItem["primary_health_check_node"] = flattenPanHaGetPanHaStatusItemPrimaryHealthCheckNode(item.PrimaryHealthCheckNode)
	respItem["secondary_health_check_node"] = flattenPanHaGetPanHaStatusItemSecondaryHealthCheckNode(item.SecondaryHealthCheckNode)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenPanHaGetPanHaStatusItemPrimaryHealthCheckNode(item *isegosdk.ResponsePanHaGetPanHaStatusResponsePrimaryHealthCheckNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["hostname"] = item.Hostname

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPanHaGetPanHaStatusItemSecondaryHealthCheckNode(item *isegosdk.ResponsePanHaGetPanHaStatusResponseSecondaryHealthCheckNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["hostname"] = item.Hostname

	return []map[string]interface{}{
		respItem,
	}

}
