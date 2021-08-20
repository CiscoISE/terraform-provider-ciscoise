package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInternalUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceInternalUserRead,
		Schema: map[string]*schema.Schema{
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

						"internal_user": &schema.Schema{
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
									"enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"email": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"change_password": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"identity_groups": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"expiry_date_enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"expiry_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_password": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"custom_attributes": &schema.Schema{
										Type:     schema.TypeMap,
										Computed: true,
									},
									"password_idstore": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
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

						"internal_user": &schema.Schema{
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
									"enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"email": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"change_password": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"identity_groups": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"expiry_date_enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"expiry_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_password": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"custom_attributes": &schema.Schema{
										Type:     schema.TypeMap,
										Computed: true,
									},
									"password_idstore": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
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

func dataSourceInternalUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %q", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetInternalUser")
		queryParams1 := isegosdk.GetInternalUserQueryParams{}

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

		response1, _, err := client.InternalUser.GetInternalUser(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetInternalUser", err,
				"Failure at GetInternalUser, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseInternalUserGetInternalUserSearchResultResources
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
				response1, _, err = client.InternalUser.GetInternalUser(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenInternalUserGetInternalUserItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInternalUser response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetInternalUserByName")
		vvName := vName.(string)

		response2, _, err := client.InternalUser.GetInternalUserByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetInternalUserByName", err,
				"Failure at GetInternalUserByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemName2 := flattenInternalUserGetInternalUserByNameItemName(&response2.InternalUser)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInternalUserByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetInternalUserByID")
		vvID := vID.(string)

		response3, _, err := client.InternalUser.GetInternalUserByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetInternalUserByID", err,
				"Failure at GetInternalUserByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

		vItemID3 := flattenInternalUserGetInternalUserByIDItemID(&response3.InternalUser)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInternalUserByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenInternalUserGetInternalUserItems(items *[]isegosdk.ResponseInternalUserGetInternalUserSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenInternalUserGetInternalUserItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenInternalUserGetInternalUserItemsLink(item isegosdk.ResponseInternalUserGetInternalUserSearchResultResourcesLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenInternalUserGetInternalUserByNameItemName(item *isegosdk.ResponseInternalUserGetInternalUserByNameInternalUser) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["enabled"] = item.Enabled
	respItem["email"] = item.Email
	respItem["password"] = item.Password
	respItem["first_name"] = item.FirstName
	respItem["last_name"] = item.LastName
	respItem["change_password"] = item.ChangePassword
	respItem["identity_groups"] = item.IDentityGroups
	respItem["expiry_date_enabled"] = item.ExpiryDateEnabled
	respItem["expiry_date"] = item.ExpiryDate
	respItem["enable_password"] = item.EnablePassword
	respItem["custom_attributes"] = item.CustomAttributes
	respItem["password_idstore"] = item.PasswordIDStore
	respItem["link"] = flattenInternalUserGetInternalUserByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenInternalUserGetInternalUserByNameItemNameLink(item isegosdk.ResponseInternalUserGetInternalUserByNameInternalUserLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenInternalUserGetInternalUserByIDItemID(item *isegosdk.ResponseInternalUserGetInternalUserByIDInternalUser) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["enabled"] = item.Enabled
	respItem["email"] = item.Email
	respItem["password"] = item.Password
	respItem["first_name"] = item.FirstName
	respItem["last_name"] = item.LastName
	respItem["change_password"] = item.ChangePassword
	respItem["identity_groups"] = item.IDentityGroups
	respItem["expiry_date_enabled"] = item.ExpiryDateEnabled
	respItem["expiry_date"] = item.ExpiryDate
	respItem["enable_password"] = item.EnablePassword
	respItem["custom_attributes"] = item.CustomAttributes
	respItem["password_idstore"] = item.PasswordIDStore
	respItem["link"] = flattenInternalUserGetInternalUserByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenInternalUserGetInternalUserByIDItemIDLink(item isegosdk.ResponseInternalUserGetInternalUserByIDInternalUserLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
