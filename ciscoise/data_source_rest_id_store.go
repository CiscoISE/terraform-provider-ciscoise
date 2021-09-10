package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRestIDStore() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRestIDStoreRead,
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
			"sortasc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sortdsc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ers_rest_idstore_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"headers": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"predefined": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"root_url": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"username_suffix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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
			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ers_rest_idstore_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"headers": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"predefined": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"root_url": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"username_suffix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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

func dataSourceRestIDStoreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSortasc, okSortasc := d.GetOk("sortasc")
	vSortdsc, okSortdsc := d.GetOk("sortdsc")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okSortasc, okSortdsc, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %v", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetRestIDStore")
		queryParams1 := isegosdk.GetRestIDStoreQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}
		if okSortasc {
			queryParams1.Sortasc = vSortasc.(string)
		}
		if okSortdsc {
			queryParams1.Sortdsc = vSortdsc.(string)
		}
		if okFilter {
			queryParams1.Filter = interfaceToSliceString(vFilter)
		}
		if okFilterType {
			queryParams1.FilterType = vFilterType.(string)
		}

		response1, _, err := client.RestidStore.GetRestIDStore(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRestIDStore", err,
				"Failure at GetRestIDStore, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseRestidStoreGetRestIDStoreSearchResultResources
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
				response1, _, err = client.RestidStore.GetRestIDStore(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenRestidStoreGetRestIDStoreItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRestIDStore response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetRestIDStoreByName")
		vvName := vName.(string)

		response2, _, err := client.RestidStore.GetRestIDStoreByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRestIDStoreByName", err,
				"Failure at GetRestIDStoreByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemName2 := flattenRestidStoreGetRestIDStoreByNameItemName(response2.ERSRestIDStore)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRestIDStoreByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetRestIDStoreByID")
		vvID := vID.(string)

		response3, _, err := client.RestidStore.GetRestIDStoreByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRestIDStoreByID", err,
				"Failure at GetRestIDStoreByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

		vItemID3 := flattenRestidStoreGetRestIDStoreByIDItemID(response3.ERSRestIDStore)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRestIDStoreByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenRestidStoreGetRestIDStoreItems(items *[]isegosdk.ResponseRestidStoreGetRestIDStoreSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenRestidStoreGetRestIDStoreItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenRestidStoreGetRestIDStoreItemsLink(item *isegosdk.ResponseRestidStoreGetRestIDStoreSearchResultResourcesLink) []map[string]interface{} {
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

func flattenRestidStoreGetRestIDStoreByNameItemName(item *isegosdk.ResponseRestidStoreGetRestIDStoreByNameERSRestIDStore) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["ers_rest_idstore_attributes"] = flattenRestidStoreGetRestIDStoreByNameItemNameErsRestIDStoreAttributes(item.ErsRestIDStoreAttributes)
	respItem["link"] = flattenRestidStoreGetRestIDStoreByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenRestidStoreGetRestIDStoreByNameItemNameErsRestIDStoreAttributes(item *isegosdk.ResponseRestidStoreGetRestIDStoreByNameERSRestIDStoreErsRestIDStoreAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["username_suffix"] = item.UsernameSuffix
	respItem["root_url"] = item.RootURL
	respItem["predefined"] = item.Predefined
	respItem["headers"] = flattenRestidStoreGetRestIDStoreByNameItemNameErsRestIDStoreAttributesHeaders(item.Headers)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenRestidStoreGetRestIDStoreByNameItemNameErsRestIDStoreAttributesHeaders(items *[]isegosdk.ResponseRestidStoreGetRestIDStoreByNameERSRestIDStoreErsRestIDStoreAttributesHeaders) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
	}
	return respItems

}

func flattenRestidStoreGetRestIDStoreByNameItemNameLink(item *isegosdk.ResponseRestidStoreGetRestIDStoreByNameERSRestIDStoreLink) []map[string]interface{} {
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

func flattenRestidStoreGetRestIDStoreByIDItemID(item *isegosdk.ResponseRestidStoreGetRestIDStoreByIDERSRestIDStore) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["ers_rest_idstore_attributes"] = flattenRestidStoreGetRestIDStoreByIDItemIDErsRestIDStoreAttributes(item.ErsRestIDStoreAttributes)
	respItem["link"] = flattenRestidStoreGetRestIDStoreByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenRestidStoreGetRestIDStoreByIDItemIDErsRestIDStoreAttributes(item *isegosdk.ResponseRestidStoreGetRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["username_suffix"] = item.UsernameSuffix
	respItem["root_url"] = item.RootURL
	respItem["predefined"] = item.Predefined
	respItem["headers"] = flattenRestidStoreGetRestIDStoreByIDItemIDErsRestIDStoreAttributesHeaders(item.Headers)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenRestidStoreGetRestIDStoreByIDItemIDErsRestIDStoreAttributesHeaders(items *[]isegosdk.ResponseRestidStoreGetRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
	}
	return respItems

}

func flattenRestidStoreGetRestIDStoreByIDItemIDLink(item *isegosdk.ResponseRestidStoreGetRestIDStoreByIDERSRestIDStoreLink) []map[string]interface{} {
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
