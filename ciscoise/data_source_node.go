package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNode() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNodeRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"display_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"gate_way": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"in_deployment": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_addresses": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_service_types": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"other_pap_fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pap_node": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"pass_word": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"primary_pap_node": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"px_grid_node": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sxp_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"display_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"gate_way": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"in_deployment": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_addresses": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_service_types": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"other_pap_fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pap_node": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"pass_word": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"primary_pap_node": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"px_grid_node": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sxp_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %q", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNodeDetails")
		queryParams1 := isegosdk.GetNodeDetailsQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}
		if okFilter {
			queryParams1.Filter = interfaceToSliceString(vFilter)
		}
		if okFilterType {
			queryParams1.FilterType = vFilterType.(string)
		}

		response1, _, err := client.NodeDetails.GetNodeDetails(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeDetails", err,
				"Failure at GetNodeDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseNodeDetailsGetNodeDetailsSearchResultResources
		for response1.SearchResult != nil && response1.SearchResult.Resources != nil && len(*response1.SearchResult.Resources) > 0 {
			items1 = append(items1, *response1.SearchResult.Resources...)
			if response1.SearchResult.NextPage != nil && response1.SearchResult.NextPage.Rel == "next" {
				href := response1.SearchResult.NextPage.Href
				page, size, err := getNextPageAndSizeParams(href)
				if err != nil {
					break
				}
				queryParams1.Page = page
				queryParams1.Size = size
				response1, _, err = client.NodeDetails.GetNodeDetails(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenNodeDetailsGetNodeDetailsItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNodeDetailByName")
		vvName := vName.(string)

		response2, _, err := client.NodeDetails.GetNodeDetailByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeDetailByName", err,
				"Failure at GetNodeDetailByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemName2 := flattenNodeDetailsGetNodeDetailByNameItemName(response2.Node)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeDetailByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetNodeDetailByID")
		vvID := vID.(string)

		response3, _, err := client.NodeDetails.GetNodeDetailByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeDetailByID", err,
				"Failure at GetNodeDetailByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

		vItemID3 := flattenNodeDetailsGetNodeDetailByIDItemID(response3.Node)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeDetailByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNodeDetailsGetNodeDetailsItems(items *[]isegosdk.ResponseNodeDetailsGetNodeDetailsSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenNodeDetailsGetNodeDetailsItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNodeDetailsGetNodeDetailsItemsLink(item *isegosdk.ResponseNodeDetailsGetNodeDetailsSearchResultResourcesLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDetailsGetNodeDetailByNameItemName(item *isegosdk.ResponseNodeDetailsGetNodeDetailByNameNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["gate_way"] = item.GateWay
	respItem["user_name"] = item.UserName
	respItem["pass_word"] = item.PassWord
	respItem["display_name"] = item.DisplayName
	respItem["in_deployment"] = item.InDeployment
	respItem["other_pap_fqdn"] = item.OtherPapFqdn
	respItem["ip_addresses"] = item.IPAddresses
	respItem["ip_address"] = item.IPAddress
	respItem["sxp_ip_address"] = item.SxpIPAddress
	respItem["node_service_types"] = item.NodeServiceTypes
	respItem["fqdn"] = item.Fqdn
	respItem["pap_node"] = item.PapNode
	respItem["primary_pap_node"] = item.PrimaryPapNode
	respItem["px_grid_node"] = item.PxGridNode
	respItem["link"] = flattenNodeDetailsGetNodeDetailByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeDetailsGetNodeDetailByNameItemNameLink(item *isegosdk.ResponseNodeDetailsGetNodeDetailByNameNodeLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDetailsGetNodeDetailByIDItemID(item *isegosdk.ResponseNodeDetailsGetNodeDetailByIDNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["gate_way"] = item.GateWay
	respItem["user_name"] = item.UserName
	respItem["pass_word"] = item.PassWord
	respItem["display_name"] = item.DisplayName
	respItem["in_deployment"] = item.InDeployment
	respItem["other_pap_fqdn"] = item.OtherPapFqdn
	respItem["ip_addresses"] = item.IPAddresses
	respItem["ip_address"] = item.IPAddress
	respItem["sxp_ip_address"] = item.SxpIPAddress
	respItem["node_service_types"] = item.NodeServiceTypes
	respItem["fqdn"] = item.Fqdn
	respItem["pap_node"] = item.PapNode
	respItem["primary_pap_node"] = item.PrimaryPapNode
	respItem["px_grid_node"] = item.PxGridNode
	respItem["link"] = flattenNodeDetailsGetNodeDetailByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeDetailsGetNodeDetailByIDItemIDLink(item *isegosdk.ResponseNodeDetailsGetNodeDetailByIDNodeLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
