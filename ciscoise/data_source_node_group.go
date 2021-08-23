package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodeGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNodeGroupRead,
		Schema: map[string]*schema.Schema{
			"node_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mar_cache": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"replication_timeout": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"replication_attempts": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"query_timeout": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"query_attempts": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mar_cache": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"replication_timeout": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"replication_attempts": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"query_timeout": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"query_attempts": &schema.Schema{
										Type:     schema.TypeInt,
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

func dataSourceNodeGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vNodeGroupName, okNodeGroupName := d.GetOk("node_group_name")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okNodeGroupName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNodeGroups")

		response1, _, err := client.NodeGroup.GetNodeGroups()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeGroups", err,
				"Failure at GetNodeGroups, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenNodeGroupGetNodeGroupsItems(&response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeGroups response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNodeGroup")
		vvNodeGroupName := vNodeGroupName.(string)

		response2, _, err := client.NodeGroup.GetNodeGroup(vvNodeGroupName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeGroup", err,
				"Failure at GetNodeGroup, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNodeGroupGetNodeGroupItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeGroup response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNodeGroupGetNodeGroupsItems(items *[]isegosdk.ResponseNodeGroupGetNodeGroupsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["mar_cache"] = flattenNodeGroupGetNodeGroupsItemsMarCache(item.MarCache)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNodeGroupGetNodeGroupsItemsMarCache(item isegosdk.ResponseNodeGroupGetNodeGroupsResponseMarCache) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["replication_timeout"] = item.ReplicationTimeout
	respItem["replication_attempts"] = item.ReplicationAttempts
	respItem["query_timeout"] = item.QueryTimeout
	respItem["query_attempts"] = item.QueryAttempts

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeGroupGetNodeGroupItem(item *isegosdk.ResponseNodeGroupGetNodeGroup) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["mar_cache"] = flattenNodeGroupGetNodeGroupItemMarCache(item.MarCache)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeGroupGetNodeGroupItemMarCache(item isegosdk.ResponseNodeGroupGetNodeGroupMarCache) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["replication_timeout"] = item.ReplicationTimeout
	respItem["replication_attempts"] = item.ReplicationAttempts
	respItem["query_timeout"] = item.QueryTimeout
	respItem["query_attempts"] = item.QueryAttempts

	return []map[string]interface{}{
		respItem,
	}

}
