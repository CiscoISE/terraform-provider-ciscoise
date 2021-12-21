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
func dataSourceNodeGroupNodeDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Node Group.

- Purpose of this API is to remove a node from a node group in the cluster. Removing node from the node group does not
delete the node, but failover is no longer carried out if the node is not part any node group.

`,

		ReadContext: dataSourceNodeGroupNodeDeleteRead,
		Schema: map[string]*schema.Schema{
			"node_group_name": &schema.Schema{
				Description: `nodeGroupName path parameter. Name of the existing node group.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"success": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"message": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNodeGroupNodeDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vNodeGroupName := d.Get("node_group_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RemoveNode")
		vvNodeGroupName := vNodeGroupName.(string)
		request1 := expandRequestNodeGroupNodeDeleteRemoveNode(ctx, "", d)

		response1, restyResp1, err := client.NodeGroup.RemoveNode(vvNodeGroupName, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RemoveNode", err,
				"Failure at RemoveNode, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNodeGroupRemoveNodeItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RemoveNode response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestNodeGroupNodeDeleteRemoveNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeGroupRemoveNode {
	request := isegosdk.RequestNodeGroupRemoveNode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	return &request
}

func flattenNodeGroupRemoveNodeItem(item *isegosdk.ResponseNodeGroupRemoveNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = flattenNodeGroupRemoveNodeItemSuccess(item.Success)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeGroupRemoveNodeItemSuccess(item *isegosdk.ResponseNodeGroupRemoveNodeSuccess) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message

	return []map[string]interface{}{
		respItem,
	}

}
