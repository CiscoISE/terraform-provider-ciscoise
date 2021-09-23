package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGuestType() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on GuestType.

- This data source allows the client to get a guest type by ID.

- This data source allows the client to get all the guest types.

Filter:

[name]

Sorting:

[name, description]
`,

		ReadContext: dataSourceGuestTypeRead,
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

						"access_time": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_access_on_specific_days_times": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"day_time_limits": &schema.Schema{
										Description: `List of Time Ranges for account access`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"days": &schema.Schema{
													Description: `List of Days
Values should be one of Week day.
Allowed values are:
- Sunday,
- Monday,
- Tuesday,
- Wednesday,
- Thursday,
- Friday,
- Saturday`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"end_time": &schema.Schema{
													Description: `End time in HH:mm format`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"start_time": &schema.Schema{
													Description: `Start time in HH:mm format`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"default_duration": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"duration_time_unit": &schema.Schema{
										Description: `Allowed values are:
- DAYS,
- HOURS,
- MINUTES`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"from_first_login": &schema.Schema{
										Description: `When Account Duration starts from first login or specified date`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_account_duration": &schema.Schema{
										Description: `Maximum value of Account Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"expiration_notification": &schema.Schema{
							Description: `Expiration Notification Settings`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advance_notification_duration": &schema.Schema{
										Description: `Send Account Expiration Notification Duration before ( Days, Hours, Minutes )`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"advance_notification_units": &schema.Schema{
										Description: `Allowed values are:
- DAYS,
- HOURS,
- MINUTES`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"email_text": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_notification": &schema.Schema{
										Description: `Enable Notification settings`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"send_email_notification": &schema.Schema{
										Description: `Enable Email Notification`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"send_sms_notification": &schema.Schema{
										Description: `Maximum devices guests can register`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"sms_text": &schema.Schema{
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
						"is_default_type": &schema.Schema{
							// Type:     schema.TypeBool,
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
						"login_options": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_guest_portal_bypass": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"failure_action": &schema.Schema{
										Description: `When Guest Exceeds limit this action will be invoked.
Allowed values are:
- Disconnect_Oldest_Connection,
- Disconnect_Newest_Connection`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"identity_group_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"limit_simultaneous_logins": &schema.Schema{
										Description: `Enable Simultaneous Logins`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_registered_devices": &schema.Schema{
										Description: `Maximum devices guests can register`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"max_simultaneous_logins": &schema.Schema{
										Description: `Number of Simultaneous Logins`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sponsor_groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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

func dataSourceGuestTypeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetGuestType")
		queryParams1 := isegosdk.GetGuestTypeQueryParams{}

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

		response1, _, err := client.GuestType.GetGuestType(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestType", err,
				"Failure at GetGuestType, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseGuestTypeGetGuestTypeSearchResultResources
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
				response1, _, err = client.GuestType.GetGuestType(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenGuestTypeGetGuestTypeItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestType response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetGuestTypeByID")
		vvID := vID.(string)

		response2, _, err := client.GuestType.GetGuestTypeByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestTypeByID", err,
				"Failure at GetGuestTypeByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenGuestTypeGetGuestTypeByIDItem(response2.GuestType)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestTypeByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenGuestTypeGetGuestTypeItems(items *[]isegosdk.ResponseGuestTypeGetGuestTypeSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenGuestTypeGetGuestTypeItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenGuestTypeGetGuestTypeItemsLink(item *isegosdk.ResponseGuestTypeGetGuestTypeSearchResultResourcesLink) []map[string]interface{} {
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

func flattenGuestTypeGetGuestTypeByIDItem(item *isegosdk.ResponseGuestTypeGetGuestTypeByIDGuestType) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["is_default_type"] = boolPtrToString(item.IsDefaultType)
	respItem["access_time"] = flattenGuestTypeGetGuestTypeByIDItemAccessTime(item.AccessTime)
	respItem["login_options"] = flattenGuestTypeGetGuestTypeByIDItemLoginOptions(item.LoginOptions)
	respItem["expiration_notification"] = flattenGuestTypeGetGuestTypeByIDItemExpirationNotification(item.ExpirationNotification)
	respItem["sponsor_groups"] = item.SponsorGroups
	respItem["link"] = flattenGuestTypeGetGuestTypeByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenGuestTypeGetGuestTypeByIDItemAccessTime(item *isegosdk.ResponseGuestTypeGetGuestTypeByIDGuestTypeAccessTime) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["from_first_login"] = boolPtrToString(item.FromFirstLogin)
	respItem["max_account_duration"] = item.MaxAccountDuration
	respItem["duration_time_unit"] = item.DurationTimeUnit
	respItem["default_duration"] = item.DefaultDuration
	respItem["allow_access_on_specific_days_times"] = boolPtrToString(item.AllowAccessOnSpecificDaysTimes)
	respItem["day_time_limits"] = flattenGuestTypeGetGuestTypeByIDItemAccessTimeDayTimeLimits(item.DayTimeLimits)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenGuestTypeGetGuestTypeByIDItemAccessTimeDayTimeLimits(items *[]isegosdk.ResponseGuestTypeGetGuestTypeByIDGuestTypeAccessTimeDayTimeLimits) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["start_time"] = item.StartTime
		respItem["end_time"] = item.EndTime
		respItem["days"] = item.Days
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenGuestTypeGetGuestTypeByIDItemLoginOptions(item *isegosdk.ResponseGuestTypeGetGuestTypeByIDGuestTypeLoginOptions) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["limit_simultaneous_logins"] = boolPtrToString(item.LimitSimultaneousLogins)
	respItem["max_simultaneous_logins"] = item.MaxSimultaneousLogins
	respItem["failure_action"] = item.FailureAction
	respItem["max_registered_devices"] = item.MaxRegisteredDevices
	respItem["identity_group_id"] = item.IDentityGroupID
	respItem["allow_guest_portal_bypass"] = boolPtrToString(item.AllowGuestPortalBypass)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenGuestTypeGetGuestTypeByIDItemExpirationNotification(item *isegosdk.ResponseGuestTypeGetGuestTypeByIDGuestTypeExpirationNotification) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enable_notification"] = boolPtrToString(item.EnableNotification)
	respItem["advance_notification_duration"] = item.AdvanceNotificationDuration
	respItem["advance_notification_units"] = item.AdvanceNotificationUnits
	respItem["send_email_notification"] = boolPtrToString(item.SendEmailNotification)
	respItem["email_text"] = item.EmailText
	respItem["send_sms_notification"] = boolPtrToString(item.SendSmsNotification)
	respItem["sms_text"] = item.SmsText

	return []map[string]interface{}{
		respItem,
	}

}

func flattenGuestTypeGetGuestTypeByIDItemLink(item *isegosdk.ResponseGuestTypeGetGuestTypeByIDGuestTypeLink) []map[string]interface{} {
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
