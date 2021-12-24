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
func dataSourceNodeGroupNodeCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Node Group.

- This data source action adds a node to the node group in the cluster. When a node that belongs to a node group fails,
another node in the same node group issues a Change of Authorization (CoA) for all the URL-redirected sessions on the
failed node.

`,

		ReadContext: dataSourceNodeGroupNodeCreateRead,
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

func dataSourceNodeGroupNodeCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vNodeGroupName := d.Get("node_group_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: AddNode")
		vvNodeGroupName := vNodeGroupName.(string)
		request1 := expandRequestNodeGroupNodeCreateAddNode(ctx, "", d)

		response1, restyResp1, err := client.NodeGroup.AddNode(vvNodeGroupName, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AddNode", err,
				"Failure at AddNode, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNodeGroupAddNodeItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting AddNode response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestNodeGroupNodeCreateAddNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeGroupAddNode {
	request := isegosdk.RequestNodeGroupAddNode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	return &request
}

func flattenNodeGroupAddNodeItem(item *isegosdk.ResponseNodeGroupAddNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = flattenNodeGroupAddNodeItemSuccess(item.Success)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeGroupAddNodeItemSuccess(item *isegosdk.ResponseNodeGroupAddNodeSuccess) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message

	return []map[string]interface{}{
		respItem,
	}

}
