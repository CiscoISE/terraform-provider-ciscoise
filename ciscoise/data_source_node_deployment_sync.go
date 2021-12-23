package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceNodeDeploymentSync() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Node Deployment.

- Performing a manual synchronization involves a reload of the target node, but not the primary PAN.
 Approximate execution time 300 seconds.
`,

		ReadContext: dataSourceNodeDeploymentSyncRead,
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Description: `hostname path parameter. Hostname of the node.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `ID which can be used to track the status of the deployment task using the Task Service API.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNodeDeploymentSyncRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vHostname := d.Get("hostname")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: SyncNode")
		vvHostname := vHostname.(string)

		response1, restyResp1, err := client.NodeDeployment.SyncNode(vvHostname)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SyncNode", err,
				"Failure at SyncNode, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNodeDeploymentSyncNodeItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SyncNode response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNodeDeploymentSyncNodeItem(item *isegosdk.ResponseNodeDeploymentSyncNodeResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
