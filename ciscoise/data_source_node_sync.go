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
func dataSourceNodeSync() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sync ISE Node.

Performing a manual sync will involve a reload of the target node, but not the primary PAN. There might be situations
where if the node has been out of sync for a long time, it may not be possible to recover via a manual sync.`,

		ReadContext: dataSourceNodeSyncRead,
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"code": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"root_cause": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNodeSyncRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SyncNode")
		request1 := expandRequestNodeSyncSyncNode(ctx, "", d)

		response1, _, err := client.SyncIseNode.SyncNode(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SyncNode", err,
				"Failure at SyncNode, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenSyncIseNodeSyncNodeItem(response1)
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

func expandRequestNodeSyncSyncNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSyncIseNodeSyncNode {
	request := isegosdk.RequestSyncIseNodeSyncNode{}
	if v, ok := d.GetOkExists(key + ".hostname"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hostname"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hostname"))) {
		request.Hostname = interfaceToString(v)
	}
	return &request
}

func flattenSyncIseNodeSyncNodeItem(item *isegosdk.ResponseSyncIseNodeSyncNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["code"] = item.Code
	respItem["message"] = item.Message
	respItem["root_cause"] = item.RootCause
	return []map[string]interface{}{
		respItem,
	}
}
