package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSgToVnToVLAN() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SecurityGroupToVirtualNetwork.

- This data source allows the client to get a security group to virtual network by ID.

- This data source allows the client to get all the security group ACL to virtual networks.

Filter:

[sgtId]

`,

		ReadContext: dataSourceSgToVnToVLANRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Description: `filter query parameter. 

**Simple filtering** should be available through the filter query string parameter. The structure of a filter is
a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator
common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query
string parameter. Each resource Data model description should specify if an attribute is a filtered field.



              Operator    | Description 

              ------------|----------------

              EQ          | Equals 

              NEQ         | Not Equals 

              GT          | Greater Than 

              LT          | Less Then 

              STARTSW     | Starts With 

              NSTARTSW    | Not Starts With 

              ENDSW       | Ends With 

              NENDSW      | Not Ends With 

              CONTAINS	  | Contains 

              NCONTAINS	  | Not Contains 

`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter_type": &schema.Schema{
				Description: `filterType query parameter. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"page": &schema.Schema{
				Description: `page query parameter. Page number`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. Number of objects returned per page`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"item": &schema.Schema{
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
						"sgt_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtualnetworklist": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default_virtual_network": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"vlans": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"data": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"max_value": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"name": &schema.Schema{
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

func dataSourceSgToVnToVLANRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSecurityGroupsToVnToVLAN")
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

		response1, restyResp1, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLAN(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSecurityGroupsToVnToVLAN", err,
				"Failure at GetSecurityGroupsToVnToVLAN, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResources
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
		log.Printf("[DEBUG] Selected method: GetSecurityGroupsToVnToVLANByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLANByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSecurityGroupsToVnToVLANByID", err,
				"Failure at GetSecurityGroupsToVnToVLANByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItem(response2.SgtVnVLANContainer)
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

func flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANItemsLink(item *isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResourcesLink) []map[string]interface{} {
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

func flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItemVirtualnetworklist(items *[]isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["default_virtual_network"] = boolPtrToString(item.DefaultVirtualNetwork)
		respItem["vlans"] = flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItemVirtualnetworklistVLANs(item.VLANs)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItemVirtualnetworklistVLANs(items *[]isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["default_vlan"] = boolPtrToString(item.DefaultVLAN)
		respItem["max_value"] = item.MaxValue
		respItem["data"] = boolPtrToString(item.Data)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItemLink(item *isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainerLink) []map[string]interface{} {
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
