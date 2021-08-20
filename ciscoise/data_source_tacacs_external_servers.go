package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTacacsExternalServers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTacacsExternalServersRead,
		Schema: map[string]*schema.Schema{
			"page": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"tacacs_external_server": &schema.Schema{
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
									"host_ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"connection_port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"single_connect": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"timeout": &schema.Schema{
										Type:     schema.TypeInt,
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
					},
				},
			},
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"tacacs_external_server": &schema.Schema{
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
									"host_ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"connection_port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"single_connect": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"timeout": &schema.Schema{
										Type:     schema.TypeInt,
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
					},
				},
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
		},
	}
}

func dataSourceTacacsExternalServersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetTacacsExternalServers")
		queryParams1 := isegosdk.GetTacacsExternalServersQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, _, err := client.TacacsExternalServers.GetTacacsExternalServers(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsExternalServers", err,
				"Failure at GetTacacsExternalServers, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseTacacsExternalServersGetTacacsExternalServersSearchResultResources
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
				response1, _, err = client.TacacsExternalServers.GetTacacsExternalServers(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenTacacsExternalServersGetTacacsExternalServersItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsExternalServers response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetTacacsExternalServersByName")
		vvName := vName.(string)

		response2, _, err := client.TacacsExternalServers.GetTacacsExternalServersByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsExternalServersByName", err,
				"Failure at GetTacacsExternalServersByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemName2 := flattenTacacsExternalServersGetTacacsExternalServersByNameItemName(&response2.TacacsExternalServer)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsExternalServersByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetTacacsExternalServersByID")
		vvID := vID.(string)

		response3, _, err := client.TacacsExternalServers.GetTacacsExternalServersByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsExternalServersByID", err,
				"Failure at GetTacacsExternalServersByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

		vItemID3 := flattenTacacsExternalServersGetTacacsExternalServersByIDItemID(&response3.TacacsExternalServer)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsExternalServersByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTacacsExternalServersGetTacacsExternalServersItems(items *[]isegosdk.ResponseTacacsExternalServersGetTacacsExternalServersSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenTacacsExternalServersGetTacacsExternalServersItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTacacsExternalServersGetTacacsExternalServersItemsLink(item isegosdk.ResponseTacacsExternalServersGetTacacsExternalServersSearchResultResourcesLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenTacacsExternalServersGetTacacsExternalServersByNameItemName(item *isegosdk.ResponseTacacsExternalServersGetTacacsExternalServersByNameTacacsExternalServer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["host_ip"] = item.HostIP
	respItem["connection_port"] = item.ConnectionPort
	respItem["single_connect"] = item.SingleConnect
	respItem["shared_secret"] = item.SharedSecret
	respItem["timeout"] = item.Timeout
	respItem["link"] = flattenTacacsExternalServersGetTacacsExternalServersByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTacacsExternalServersGetTacacsExternalServersByNameItemNameLink(item isegosdk.ResponseTacacsExternalServersGetTacacsExternalServersByNameTacacsExternalServerLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenTacacsExternalServersGetTacacsExternalServersByIDItemID(item *isegosdk.ResponseTacacsExternalServersGetTacacsExternalServersByIDTacacsExternalServer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["host_ip"] = item.HostIP
	respItem["connection_port"] = item.ConnectionPort
	respItem["single_connect"] = item.SingleConnect
	respItem["shared_secret"] = item.SharedSecret
	respItem["timeout"] = item.Timeout
	respItem["link"] = flattenTacacsExternalServersGetTacacsExternalServersByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTacacsExternalServersGetTacacsExternalServersByIDItemIDLink(item isegosdk.ResponseTacacsExternalServersGetTacacsExternalServersByIDTacacsExternalServerLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
