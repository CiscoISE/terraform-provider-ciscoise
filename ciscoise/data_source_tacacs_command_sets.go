package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTacacsCommandSets() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on TACACSCommandSets.

- This data source allows the client to get TACACS command sets by name.

- This data source allows the client to get TACACS command sets by ID.

- This data source allows the client to get all the TACACS command sets.
`,

		ReadContext: dataSourceTacacsCommandSetsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name path parameter.`,
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
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"commands": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"command_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"arguments": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"command": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"grant": &schema.Schema{
													Description: `Allowed values: PERMIT, DENY, DENY_ALWAYS`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
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
						"permit_unmatched": &schema.Schema{
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

						"commands": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"command_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"arguments": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"command": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"grant": &schema.Schema{
													Description: `Allowed values: PERMIT, DENY, DENY_ALWAYS`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
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
						"permit_unmatched": &schema.Schema{
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

func dataSourceTacacsCommandSetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetTacacsCommandSets")
		queryParams1 := isegosdk.GetTacacsCommandSetsQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, restyResp1, err := client.TacacsCommandSets.GetTacacsCommandSets(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsCommandSets", err,
				"Failure at GetTacacsCommandSets, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultResources
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
				response1, _, err = client.TacacsCommandSets.GetTacacsCommandSets(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenTacacsCommandSetsGetTacacsCommandSetsItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsCommandSets response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetTacacsCommandSetsByName")
		vvName := vName.(string)

		response2, _, err := client.TacacsCommandSets.GetTacacsCommandSetsByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsCommandSetsByName", err,
				"Failure at GetTacacsCommandSetsByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemName2 := flattenTacacsCommandSetsGetTacacsCommandSetsByNameItemName(response2.TacacsCommandSets)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsCommandSetsByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetTacacsCommandSetsByID")
		vvID := vID.(string)

		response3, _, err := client.TacacsCommandSets.GetTacacsCommandSetsByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsCommandSetsByID", err,
				"Failure at GetTacacsCommandSetsByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response3))

		vItemID3 := flattenTacacsCommandSetsGetTacacsCommandSetsByIDItemID(response3.TacacsCommandSets)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsCommandSetsByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTacacsCommandSetsGetTacacsCommandSetsItems(items *[]isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenTacacsCommandSetsGetTacacsCommandSetsItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTacacsCommandSetsGetTacacsCommandSetsItemsLink(item *isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsSearchResultResourcesLink) []map[string]interface{} {
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

func flattenTacacsCommandSetsGetTacacsCommandSetsByNameItemName(item *isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSets) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["permit_unmatched"] = boolPtrToString(item.PermitUnmatched)
	respItem["commands"] = flattenTacacsCommandSetsGetTacacsCommandSetsByNameItemNameCommands(item.Commands)
	respItem["link"] = flattenTacacsCommandSetsGetTacacsCommandSetsByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTacacsCommandSetsGetTacacsCommandSetsByNameItemNameCommands(item *isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSetsCommands) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["command_list"] = flattenTacacsCommandSetsGetTacacsCommandSetsByNameItemNameCommandsCommandList(item.CommandList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenTacacsCommandSetsGetTacacsCommandSetsByNameItemNameCommandsCommandList(items *[]isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSetsCommandsCommandList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["grant"] = item.Grant
		respItem["command"] = item.Command
		respItem["arguments"] = item.Arguments
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenTacacsCommandSetsGetTacacsCommandSetsByNameItemNameLink(item *isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsByNameTacacsCommandSetsLink) []map[string]interface{} {
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

func flattenTacacsCommandSetsGetTacacsCommandSetsByIDItemID(item *isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSets) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["permit_unmatched"] = boolPtrToString(item.PermitUnmatched)
	respItem["commands"] = flattenTacacsCommandSetsGetTacacsCommandSetsByIDItemIDCommands(item.Commands)
	respItem["link"] = flattenTacacsCommandSetsGetTacacsCommandSetsByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTacacsCommandSetsGetTacacsCommandSetsByIDItemIDCommands(item *isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSetsCommands) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["command_list"] = flattenTacacsCommandSetsGetTacacsCommandSetsByIDItemIDCommandsCommandList(item.CommandList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenTacacsCommandSetsGetTacacsCommandSetsByIDItemIDCommandsCommandList(items *[]isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["grant"] = item.Grant
		respItem["command"] = item.Command
		respItem["arguments"] = item.Arguments
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenTacacsCommandSetsGetTacacsCommandSetsByIDItemIDLink(item *isegosdk.ResponseTacacsCommandSetsGetTacacsCommandSetsByIDTacacsCommandSetsLink) []map[string]interface{} {
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
