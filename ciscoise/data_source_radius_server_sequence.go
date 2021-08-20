package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRadiusServerSequence() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRadiusServerSequenceRead,
		Schema: map[string]*schema.Schema{
			"page": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
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

						"radius_server_sequence": &schema.Schema{
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
									"strip_prefix": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"strip_suffix": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"prefix_separator": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"suffix_separator": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"remote_accounting": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"local_accounting": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"use_attr_set_on_request": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"use_attr_set_before_acc": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"continue_authorz_policy": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"radius_server_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"on_request_attr_manipulator_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"action": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"dictionary_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"attribute_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"changed_val": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"before_accept_attr_manipulators_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"action": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"dictionary_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"attribute_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"changed_val": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
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

func dataSourceRadiusServerSequenceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetRadiusServerSequence")
		queryParams1 := isegosdk.GetRadiusServerSequenceQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, _, err := client.RadiusServerSequence.GetRadiusServerSequence(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRadiusServerSequence", err,
				"Failure at GetRadiusServerSequence, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultResources
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
				response1, _, err = client.RadiusServerSequence.GetRadiusServerSequence(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenRadiusServerSequenceGetRadiusServerSequenceItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRadiusServerSequence response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetRadiusServerSequenceByID")
		vvID := vID.(string)

		response2, _, err := client.RadiusServerSequence.GetRadiusServerSequenceByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRadiusServerSequenceByID", err,
				"Failure at GetRadiusServerSequenceByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenRadiusServerSequenceGetRadiusServerSequenceByIDItem(&response2.RadiusServerSequence)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRadiusServerSequenceByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenRadiusServerSequenceGetRadiusServerSequenceItems(items *[]isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenRadiusServerSequenceGetRadiusServerSequenceItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenRadiusServerSequenceGetRadiusServerSequenceItemsLink(item isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultResourcesLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenRadiusServerSequenceGetRadiusServerSequenceByIDItem(item *isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequence) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["strip_prefix"] = item.StripPrefix
	respItem["strip_suffix"] = item.StripSuffix
	respItem["prefix_separator"] = item.PrefixSeparator
	respItem["suffix_separator"] = item.SuffixSeparator
	respItem["remote_accounting"] = item.RemoteAccounting
	respItem["local_accounting"] = item.LocalAccounting
	respItem["use_attr_set_on_request"] = item.UseAttrSetOnRequest
	respItem["use_attr_set_before_acc"] = item.UseAttrSetBeforeAcc
	respItem["continue_authorz_policy"] = item.ContinueAuthorzPolicy
	respItem["radius_server_list"] = item.RadiusServerList
	respItem["on_request_attr_manipulator_list"] = flattenRadiusServerSequenceGetRadiusServerSequenceByIDItemOnRequestAttrManipulatorList(item.OnRequestAttrManipulatorList)
	respItem["before_accept_attr_manipulators_list"] = flattenRadiusServerSequenceGetRadiusServerSequenceByIDItemBeforeAcceptAttrManipulatorsList(item.BeforeAcceptAttrManipulatorsList)
	respItem["link"] = flattenRadiusServerSequenceGetRadiusServerSequenceByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenRadiusServerSequenceGetRadiusServerSequenceByIDItemOnRequestAttrManipulatorList(items []isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["action"] = item.Action
		respItem["dictionary_name"] = item.DictionaryName
		respItem["attribute_name"] = item.AttributeName
		respItem["value"] = item.Value
		respItem["changed_val"] = item.ChangedVal
	}
	return respItems

}

func flattenRadiusServerSequenceGetRadiusServerSequenceByIDItemBeforeAcceptAttrManipulatorsList(items []isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["action"] = item.Action
		respItem["dictionary_name"] = item.DictionaryName
		respItem["attribute_name"] = item.AttributeName
		respItem["value"] = item.Value
		respItem["changed_val"] = item.ChangedVal
	}
	return respItems

}

func flattenRadiusServerSequenceGetRadiusServerSequenceByIDItemLink(item isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequenceLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
