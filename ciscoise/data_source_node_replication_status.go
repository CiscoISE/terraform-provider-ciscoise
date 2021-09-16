package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodeReplicationStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Replication Status.

Retrives replication status of a node`,

		ReadContext: dataSourceNodeReplicationStatusRead,
		Schema: map[string]*schema.Schema{
			"node": &schema.Schema{
				Description: `node path parameter. ID of the existing node.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"node_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNodeReplicationStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vNode := d.Get("node")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNodeReplicationStatus")
		vvNode := vNode.(string)

		response1, _, err := client.ReplicationStatus.GetNodeReplicationStatus(vvNode)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeReplicationStatus", err,
				"Failure at GetNodeReplicationStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenReplicationStatusGetNodeReplicationStatusItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeReplicationStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenReplicationStatusGetNodeReplicationStatusItem(item *isegosdk.ResponseReplicationStatusGetNodeReplicationStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["node_status"] = item.NodeStatus
	return []map[string]interface{}{
		respItem,
	}
}
