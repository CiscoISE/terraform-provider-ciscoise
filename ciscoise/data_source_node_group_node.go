package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodeGroupNode() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Node Group.

- This data source retrieves the list of nodes associated with a node group in the cluster with a given node group name.
`,

		ReadContext: dataSourceNodeGroupNodeRead,
		Schema: map[string]*schema.Schema{
			"node_group_name": &schema.Schema{
				Description: `nodeGroupName path parameter. Name of the existing node group.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"items": &schema.Schema{
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
	}
}

func dataSourceNodeGroupNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vNodeGroupName := d.Get("node_group_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNodes")
		vvNodeGroupName := vNodeGroupName.(string)

		response1, restyResp1, err := client.NodeGroup.GetNodes(vvNodeGroupName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodes", err,
				"Failure at GetNodes, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNodeGroupGetNodesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodes response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNodeGroupGetNodesItems(items *[]isegosdk.ResponseNodeGroupGetNodesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["hostname"] = item.Hostname
		respItems = append(respItems, respItem)
	}
	return respItems
}
