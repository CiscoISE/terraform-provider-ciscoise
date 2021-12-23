package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIDStoreSequence() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on IdentitySequence.

- This data source allows the client to get an identity sequence by name.

- This data source allows the client to get an identity sequence by ID.

- This data source allows the client to get all the identity sequences.
`,

		ReadContext: dataSourceIDStoreSequenceRead,
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

						"break_on_store_fail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_authentication_profile": &schema.Schema{
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
						"id_seq_item": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"idstore": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"order": &schema.Schema{
										Type:     schema.TypeInt,
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
						"parent": &schema.Schema{
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

						"break_on_store_fail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_authentication_profile": &schema.Schema{
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
						"id_seq_item": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"idstore": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"order": &schema.Schema{
										Type:     schema.TypeInt,
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
						"parent": &schema.Schema{
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

func dataSourceIDStoreSequenceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetIDentitySequence")
		queryParams1 := isegosdk.GetIDentitySequenceQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, restyResp1, err := client.IDentitySequence.GetIDentitySequence(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetIDentitySequence", err,
				"Failure at GetIDentitySequence, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseIDentitySequenceGetIDentitySequenceSearchResultResources
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
				response1, _, err = client.IDentitySequence.GetIDentitySequence(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenIDentitySequenceGetIDentitySequenceItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIDentitySequence response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetIDentitySequenceByName")
		vvName := vName.(string)

		response2, restyResp2, err := client.IDentitySequence.GetIDentitySequenceByName(vvName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetIDentitySequenceByName", err,
				"Failure at GetIDentitySequenceByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemName2 := flattenIDentitySequenceGetIDentitySequenceByNameItemName(response2.IDStoreSequence)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIDentitySequenceByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetIDentitySequenceByID")
		vvID := vID.(string)

		response3, restyResp3, err := client.IDentitySequence.GetIDentitySequenceByID(vvID)

		if err != nil || response3 == nil {
			if restyResp3 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp3.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetIDentitySequenceByID", err,
				"Failure at GetIDentitySequenceByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response3))

		vItemID3 := flattenIDentitySequenceGetIDentitySequenceByIDItemID(response3.IDStoreSequence)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIDentitySequenceByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIDentitySequenceGetIDentitySequenceItems(items *[]isegosdk.ResponseIDentitySequenceGetIDentitySequenceSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenIDentitySequenceGetIDentitySequenceItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIDentitySequenceGetIDentitySequenceItemsLink(item *isegosdk.ResponseIDentitySequenceGetIDentitySequenceSearchResultResourcesLink) []map[string]interface{} {
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

func flattenIDentitySequenceGetIDentitySequenceByNameItemName(item *isegosdk.ResponseIDentitySequenceGetIDentitySequenceByNameIDStoreSequence) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["parent"] = item.Parent
	respItem["id_seq_item"] = flattenIDentitySequenceGetIDentitySequenceByNameItemNameIDSeqItem(item.IDSeqItem)
	respItem["certificate_authentication_profile"] = item.CertificateAuthenticationProfile
	respItem["break_on_store_fail"] = boolPtrToString(item.BreakOnStoreFail)
	respItem["link"] = flattenIDentitySequenceGetIDentitySequenceByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIDentitySequenceGetIDentitySequenceByNameItemNameIDSeqItem(items *[]isegosdk.ResponseIDentitySequenceGetIDentitySequenceByNameIDStoreSequenceIDSeqItem) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["idstore"] = item.IDstore
		respItem["order"] = item.Order
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIDentitySequenceGetIDentitySequenceByNameItemNameLink(item *isegosdk.ResponseIDentitySequenceGetIDentitySequenceByNameIDStoreSequenceLink) []map[string]interface{} {
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

func flattenIDentitySequenceGetIDentitySequenceByIDItemID(item *isegosdk.ResponseIDentitySequenceGetIDentitySequenceByIDIDStoreSequence) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["parent"] = item.Parent
	respItem["id_seq_item"] = flattenIDentitySequenceGetIDentitySequenceByIDItemIDIDSeqItem(item.IDSeqItem)
	respItem["certificate_authentication_profile"] = item.CertificateAuthenticationProfile
	respItem["break_on_store_fail"] = boolPtrToString(item.BreakOnStoreFail)
	respItem["link"] = flattenIDentitySequenceGetIDentitySequenceByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIDentitySequenceGetIDentitySequenceByIDItemIDIDSeqItem(items *[]isegosdk.ResponseIDentitySequenceGetIDentitySequenceByIDIDStoreSequenceIDSeqItem) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["idstore"] = item.IDstore
		respItem["order"] = item.Order
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIDentitySequenceGetIDentitySequenceByIDItemIDLink(item *isegosdk.ResponseIDentitySequenceGetIDentitySequenceByIDIDStoreSequenceLink) []map[string]interface{} {
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
