package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTacacsServerSequence() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTacacsServerSequenceRead,
		Schema: map[string]*schema.Schema{
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
						"local_accounting": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_strip": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"remote_accounting": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"server_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"suffix_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"suffix_strip": &schema.Schema{
							Type:     schema.TypeBool,
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
						"local_accounting": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_strip": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"remote_accounting": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"server_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"suffix_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"suffix_strip": &schema.Schema{
							Type:     schema.TypeBool,
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

func dataSourceTacacsServerSequenceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %q", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTacacsServerSequence")
		queryParams1 := isegosdk.GetTacacsServerSequenceQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, _, err := client.TacacsServerSequence.GetTacacsServerSequence(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsServerSequence", err,
				"Failure at GetTacacsServerSequence, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultResources
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
				response1, _, err = client.TacacsServerSequence.GetTacacsServerSequence(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenTacacsServerSequenceGetTacacsServerSequenceItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsServerSequence response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetTacacsServerSequenceByName")
		vvName := vName.(string)

		response2, _, err := client.TacacsServerSequence.GetTacacsServerSequenceByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsServerSequenceByName", err,
				"Failure at GetTacacsServerSequenceByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemName2 := flattenTacacsServerSequenceGetTacacsServerSequenceByNameItemName(response2.TacacsServerSequence)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsServerSequenceByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetTacacsServerSequenceByID")
		vvID := vID.(string)

		response3, _, err := client.TacacsServerSequence.GetTacacsServerSequenceByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsServerSequenceByID", err,
				"Failure at GetTacacsServerSequenceByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

		vItemID3 := flattenTacacsServerSequenceGetTacacsServerSequenceByIDItemID(response3.TacacsServerSequence)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsServerSequenceByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTacacsServerSequenceGetTacacsServerSequenceItems(items *[]isegosdk.ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenTacacsServerSequenceGetTacacsServerSequenceItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTacacsServerSequenceGetTacacsServerSequenceItemsLink(item *isegosdk.ResponseTacacsServerSequenceGetTacacsServerSequenceSearchResultResourcesLink) []map[string]interface{} {
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

func flattenTacacsServerSequenceGetTacacsServerSequenceByNameItemName(item *isegosdk.ResponseTacacsServerSequenceGetTacacsServerSequenceByNameTacacsServerSequence) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["server_list"] = item.ServerList
	respItem["local_accounting"] = item.LocalAccounting
	respItem["remote_accounting"] = item.RemoteAccounting
	respItem["prefix_strip"] = item.PrefixStrip
	respItem["prefix_delimiter"] = item.PrefixDelimiter
	respItem["suffix_strip"] = item.SuffixStrip
	respItem["suffix_delimiter"] = item.SuffixDelimiter
	respItem["link"] = flattenTacacsServerSequenceGetTacacsServerSequenceByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTacacsServerSequenceGetTacacsServerSequenceByNameItemNameLink(item *isegosdk.ResponseTacacsServerSequenceGetTacacsServerSequenceByNameTacacsServerSequenceLink) []map[string]interface{} {
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

func flattenTacacsServerSequenceGetTacacsServerSequenceByIDItemID(item *isegosdk.ResponseTacacsServerSequenceGetTacacsServerSequenceByIDTacacsServerSequence) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["server_list"] = item.ServerList
	respItem["local_accounting"] = item.LocalAccounting
	respItem["remote_accounting"] = item.RemoteAccounting
	respItem["prefix_strip"] = item.PrefixStrip
	respItem["prefix_delimiter"] = item.PrefixDelimiter
	respItem["suffix_strip"] = item.SuffixStrip
	respItem["suffix_delimiter"] = item.SuffixDelimiter
	respItem["link"] = flattenTacacsServerSequenceGetTacacsServerSequenceByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTacacsServerSequenceGetTacacsServerSequenceByIDItemIDLink(item *isegosdk.ResponseTacacsServerSequenceGetTacacsServerSequenceByIDTacacsServerSequenceLink) []map[string]interface{} {
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
