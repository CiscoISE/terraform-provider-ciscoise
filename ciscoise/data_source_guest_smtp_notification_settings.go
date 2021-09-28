package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGuestSmtpNotificationSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on GuestSMTPNotificationConfiguration.

- This data source allows the client to get a guest SMTP notification configuration by ID.

- This data source allows the client to get all the guest SMTP notification configurations.

Filter:

[name]

Sorting:

[name, description]
`,

		ReadContext: dataSourceGuestSmtpNotificationSettingsRead,
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
			"sortasc": &schema.Schema{
				Description: `sortasc query parameter. sort asc`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sortdsc": &schema.Schema{
				Description: `sortdsc query parameter. sort desc`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connection_timeout": &schema.Schema{
							Description: `Interval in seconds for all the SMTP client connections`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"default_from_address": &schema.Schema{
							Description: `The default from email address to be used to send emails from`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `Indicates if the email notification service is to be enabled`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password": &schema.Schema{
							Description: `Password of Secure SMTP server`,
							Type:        schema.TypeString,
							Sensitive:   true,
							Computed:    true,
						},
						"smtp_port": &schema.Schema{
							Description: `Port at which SMTP Secure Server is listening`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"smtp_server": &schema.Schema{
							Description: `The SMTP server ip address or fqdn such as outbound.mycompany.com`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"use_default_from_address": &schema.Schema{
							Description: `If the default from address should be used rather than using a sponsor user email address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"use_password_authentication": &schema.Schema{
							Description: `If configured to true, SMTP server authentication will happen using username/password`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"use_tlsor_ssl_encryption": &schema.Schema{
							Description: `If configured to true, SMTP server authentication will happen using TLS/SSL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"user_name": &schema.Schema{
							Description: `Username of Secure SMTP server`,
							Type:        schema.TypeString,
							Computed:    true,
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
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

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

		response1, restyResp1, err := client.GuestSmtpNotificationConfiguration.GetGuestSmtpNotificationSettings(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestSmtpNotificationSettings", err,
				"Failure at GetGuestSmtpNotificationSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

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

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

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
	respItem["notification_enabled"] = boolPtrToString(item.NotificationEnabled)
	respItem["use_default_from_address"] = boolPtrToString(item.UseDefaultFromAddress)
	respItem["default_from_address"] = item.DefaultFromAddress
	respItem["smtp_port"] = item.SmtpPort
	respItem["connection_timeout"] = item.ConnectionTimeout
	respItem["use_tlsor_ssl_encryption"] = boolPtrToString(item.UseTLSorSSLEncryption)
	respItem["use_password_authentication"] = boolPtrToString(item.UsePasswordAuthentication)
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
