package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSgToVnToVLAN() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSgToVnToVLANRead,
		Schema: map[string]*schema.Schema{
			"page": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
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
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"href": &schema.Schema{
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
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"sgt_vnvlan_container": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"sgt_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"virtualnetworklist": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_virtual_network": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"vlans": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"default_vlan": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"max_value": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"data": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"link": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"rel": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"href": &schema.Schema{
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceSgToVnToVLANRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSecurityGroupsToVnToVLAN")
		queryParams1 := isegosdk.GetSecurityGroupsToVnToVLANQueryParams{}

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

		response1, _, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLAN(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSecurityGroupsToVnToVLAN", err,
				"Failure at GetSecurityGroupsToVnToVLAN, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResources
		for len(response1.SearchResult.Resources) > 0 {
			items1 = append(items1, response1.SearchResult.Resources...)
			if response1.SearchResult.NextPage.Rel == "next" {
				href := response1.SearchResult.NextPage.Href
				page, size, err := getNextPageAndSizeParams(href)
				if err != nil {
					break
				}
				queryParams1.Page = page
				queryParams1.Size = size
				response1, _, err = client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLAN(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityGroupsToVnToVLAN response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetSecurityGroupsToVnToVLANByID")
		vvID := vID.(string)

		response2, _, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLANByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSecurityGroupsToVnToVLANByID", err,
				"Failure at GetSecurityGroupsToVnToVLANByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItem(&response2.SgtVnVLANContainer)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityGroupsToVnToVLANByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANItems(items *[]isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANItemsLink(item isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResourcesLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItem(item *isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["sgt_id"] = item.SgtID
	respItem["virtualnetworklist"] = flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItemVirtualnetworklist(item.Virtualnetworklist)
	respItem["link"] = flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItemVirtualnetworklist(items []isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["default_virtual_network"] = item.DefaultVirtualNetwork
		respItem["vlans"] = flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItemVirtualnetworklistVLANs(item.VLANs)
	}
	return respItems

}

func flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItemVirtualnetworklistVLANs(items []isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["default_vlan"] = item.DefaultVLAN
		respItem["max_value"] = item.MaxValue
		respItem["data"] = item.Data
	}
	return respItems

}

func flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItemLink(item isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
