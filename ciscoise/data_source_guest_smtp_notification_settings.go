package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGuestSmtpNotificationSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGuestSmtpNotificationSettingsRead,
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
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connection_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_from_address": &schema.Schema{
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
						"notification_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"smtp_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"smtp_server": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"use_default_from_address": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"use_password_authentication": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"use_tlsor_ssl_encryption": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"user_name": &schema.Schema{
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
					},
				},
			},
		},
	}
}

func dataSourceGuestSmtpNotificationSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGuestSmtpNotificationSettings")
		queryParams1 := isegosdk.GetGuestSmtpNotificationSettingsQueryParams{}

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

		response1, _, err := client.GuestSmtpNotificationConfiguration.GetGuestSmtpNotificationSettings(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestSmtpNotificationSettings", err,
				"Failure at GetGuestSmtpNotificationSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultResources
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
				response1, _, err = client.GuestSmtpNotificationConfiguration.GetGuestSmtpNotificationSettings(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestSmtpNotificationSettings response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetGuestSmtpNotificationSettingsByID")
		vvID := vID.(string)

		response2, _, err := client.GuestSmtpNotificationConfiguration.GetGuestSmtpNotificationSettingsByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestSmtpNotificationSettingsByID", err,
				"Failure at GetGuestSmtpNotificationSettingsByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDItem(response2.ERSGuestSmtpNotificationSettings)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestSmtpNotificationSettingsByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsItems(items *[]isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["link"] = flattenGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsItemsLink(item *isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultResourcesLink) []map[string]interface{} {
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

func flattenGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDItem(item *isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["smtp_server"] = item.SmtpServer
	respItem["notification_enabled"] = item.NotificationEnabled
	respItem["use_default_from_address"] = item.UseDefaultFromAddress
	respItem["default_from_address"] = item.DefaultFromAddress
	respItem["smtp_port"] = item.SmtpPort
	respItem["connection_timeout"] = item.ConnectionTimeout
	respItem["use_tlsor_ssl_encryption"] = item.UseTLSorSSLEncryption
	respItem["use_password_authentication"] = item.UsePasswordAuthentication
	respItem["user_name"] = item.UserName
	respItem["password"] = item.Password
	respItem["link"] = flattenGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDItemLink(item *isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettingsLink) []map[string]interface{} {
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
