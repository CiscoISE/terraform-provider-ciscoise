package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAdminUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAdminUserRead,
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

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
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
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"change_password": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"include_system_alarms_in_email": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"external_user": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"inactive_account_never_disabled": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"admin_groups": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_attributes": &schema.Schema{
							Type:     schema.TypeMap,
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

func dataSourceAdminUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSortasc, okSortasc := d.GetOk("sortasc")
	vSortdsc, okSortdsc := d.GetOk("sortdsc")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okSortasc, okSortdsc, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAdminUsers")
		queryParams1 := isegosdk.GetAdminUsersQueryParams{}

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

		response1, _, err := client.AdminUser.GetAdminUsers(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAdminUsers", err,
				"Failure at GetAdminUsers, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseAdminUserGetAdminUsersSearchResultResources
		for response1.SearchResult != nil && response1.SearchResult.Resources != nil && len(*response1.SearchResult.Resources) > 0 {
			items1 = append(items1, *response1.SearchResult.Resources...)
			if response1.SearchResult.NextPage.Rel == "next" {
				href := response1.SearchResult.NextPage.Href
				page, size, err := getNextPageAndSizeParams(href)
				if err != nil {
					break
				}
				queryParams1.Page = page
				queryParams1.Size = size
				response1, _, err = client.AdminUser.GetAdminUsers(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenAdminUserGetAdminUsersItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAdminUsers response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetAdminUserByID")
		vvID := vID.(string)

		response2, _, err := client.AdminUser.GetAdminUserByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAdminUserByID", err,
				"Failure at GetAdminUserByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenAdminUserGetAdminUserByIDItem(response2.AdminUser)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAdminUserByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAdminUserGetAdminUsersItems(items *[]isegosdk.ResponseAdminUserGetAdminUsersSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenAdminUserGetAdminUsersItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenAdminUserGetAdminUsersItemsLink(item *isegosdk.ResponseAdminUserGetAdminUsersSearchResultResourcesLink) []map[string]interface{} {
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

func flattenAdminUserGetAdminUserByIDItem(item *isegosdk.ResponseAdminUserGetAdminUserByIDAdminUser) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["id"] = item.ID
	respItem["description"] = item.Description
	respItem["enabled"] = item.Enabled
	respItem["password"] = item.Password
	respItem["change_password"] = item.ChangePassword
	respItem["include_system_alarms_in_email"] = item.IncludeSystemAlarmsInEmail
	respItem["external_user"] = item.ExternalUser
	respItem["inactive_account_never_disabled"] = item.InactiveAccountNeverDisabled
	respItem["admin_groups"] = item.AdminGroups
	respItem["custom_attributes"] = item.CustomAttributes
	respItem["link"] = flattenAdminUserGetAdminUserByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenAdminUserGetAdminUserByIDItemLink(item *isegosdk.ResponseAdminUserGetAdminUserByIDAdminUserLink) []map[string]interface{} {
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