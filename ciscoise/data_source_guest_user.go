package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGuestUser() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on GuestUser.

- This data source allows the client to get a guest user by name.

- This data source allows the client to get a guest user by ID.

- This data source allows the client to get all the guest users.

Filter:

[lastName, sponsor, creationTime, personBeingVisited, toDate, userName, firstName, emailAddress, phoneNumber, groupTag,
name, company, guestType, status]

Sorting:

[firstName, lastName, emailAddress, name, description]
`,

		ReadContext: dataSourceGuestUserRead,
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
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_fields": &schema.Schema{
							Description: `Key value map`,
							// CHECK: The type of this param
							// Replaced List to Map
							Type:     schema.TypeMap,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"guest_access_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"from_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"group_tag": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"location": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"to_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"valid_days": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"guest_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"company": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"creation_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"email_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": &schema.Schema{
										Description: `This field is only for Get operation not applicable for Create, Update operations`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"notification_language": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"phone_number": &schema.Schema{
										Description: `Phone number should be E.164 format`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"sms_service_provider": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_name": &schema.Schema{
										Description: `If account needs be created with mobile number, please provide mobile number here`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"guest_type": &schema.Schema{
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
						"portal_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reason_for_visit": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sponsor_user_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sponsor_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_reason": &schema.Schema{
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

						"custom_fields": &schema.Schema{
							Description: `Key value map`,
							// CHECK: The type of this param
							// Replaced List to Map
							Type:     schema.TypeMap,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"guest_access_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"from_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"group_tag": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"location": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"to_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"valid_days": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"guest_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"company": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"creation_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"email_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": &schema.Schema{
										Description: `This field is only for Get operation not applicable for Create, Update operations`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"notification_language": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"phone_number": &schema.Schema{
										Description: `Phone number should be E.164 format`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"sms_service_provider": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_name": &schema.Schema{
										Description: `If account needs be created with mobile number, please provide mobile number here`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"guest_type": &schema.Schema{
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
						"portal_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reason_for_visit": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sponsor_user_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sponsor_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_reason": &schema.Schema{
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

func dataSourceGuestUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetGuestUsers")
		queryParams1 := isegosdk.GetGuestUsersQueryParams{}

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

		response1, restyResp1, err := client.GuestUser.GetGuestUsers(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestUsers", err,
				"Failure at GetGuestUsers, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseGuestUserGetGuestUsersSearchResultResources
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
				response1, _, err = client.GuestUser.GetGuestUsers(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenGuestUserGetGuestUsersItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestUsers response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetGuestUserByName")
		vvName := vName.(string)

		response2, restyResp2, err := client.GuestUser.GetGuestUserByName(vvName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestUserByName", err,
				"Failure at GetGuestUserByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemName2 := flattenGuestUserGetGuestUserByNameItemName(response2.GuestUser)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestUserByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetGuestUserByID")
		vvID := vID.(string)

		response3, restyResp3, err := client.GuestUser.GetGuestUserByID(vvID)

		if err != nil || response3 == nil {
			if restyResp3 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp3.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestUserByID", err,
				"Failure at GetGuestUserByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response3))

		vItemID3 := flattenGuestUserGetGuestUserByIDItemID(response3.GuestUser)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestUserByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenGuestUserGetGuestUsersItems(items *[]isegosdk.ResponseGuestUserGetGuestUsersSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenGuestUserGetGuestUsersItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenGuestUserGetGuestUsersItemsLink(item *isegosdk.ResponseGuestUserGetGuestUsersSearchResultResourcesLink) []map[string]interface{} {
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

func flattenGuestUserGetGuestUserByNameItemName(item *isegosdk.ResponseGuestUserGetGuestUserByNameGuestUser) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["guest_type"] = item.GuestType
	respItem["status"] = item.Status
	respItem["status_reason"] = item.StatusReason
	respItem["reason_for_visit"] = item.ReasonForVisit
	respItem["sponsor_user_id"] = item.SponsorUserID
	respItem["sponsor_user_name"] = item.SponsorUserName
	respItem["guest_info"] = flattenGuestUserGetGuestUserByNameItemNameGuestInfo(item.GuestInfo)
	respItem["guest_access_info"] = flattenGuestUserGetGuestUserByNameItemNameGuestAccessInfo(item.GuestAccessInfo)
	respItem["portal_id"] = item.PortalID
	respItem["custom_fields"] = flattenGuestUserGetGuestUserByNameItemNameCustomFields(item.CustomFields)
	respItem["link"] = flattenGuestUserGetGuestUserByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenGuestUserGetGuestUserByNameItemNameGuestInfo(item *isegosdk.ResponseGuestUserGetGuestUserByNameGuestUserGuestInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["first_name"] = item.FirstName
	respItem["last_name"] = item.LastName
	respItem["company"] = item.Company
	respItem["creation_time"] = item.CreationTime
	respItem["notification_language"] = item.NotificationLanguage
	respItem["user_name"] = item.UserName
	respItem["email_address"] = item.EmailAddress
	respItem["phone_number"] = item.PhoneNumber
	respItem["password"] = item.Password
	respItem["enabled"] = boolPtrToString(item.Enabled)
	respItem["sms_service_provider"] = item.SmsServiceProvider

	return []map[string]interface{}{
		respItem,
	}

}

func flattenGuestUserGetGuestUserByNameItemNameGuestAccessInfo(item *isegosdk.ResponseGuestUserGetGuestUserByNameGuestUserGuestAccessInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["valid_days"] = item.ValidDays
	respItem["from_date"] = item.FromDate
	respItem["to_date"] = item.ToDate
	respItem["location"] = item.Location
	respItem["ssid"] = item.SSID
	respItem["group_tag"] = item.GroupTag

	return []map[string]interface{}{
		respItem,
	}

}

func flattenGuestUserGetGuestUserByNameItemNameCustomFields(item *isegosdk.ResponseGuestUserGetGuestUserByNameGuestUserCustomFields) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return respItem

}

func flattenGuestUserGetGuestUserByNameItemNameLink(item *isegosdk.ResponseGuestUserGetGuestUserByNameGuestUserLink) []map[string]interface{} {
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

func flattenGuestUserGetGuestUserByIDItemID(item *isegosdk.ResponseGuestUserGetGuestUserByIDGuestUser) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["guest_type"] = item.GuestType
	respItem["status"] = item.Status
	respItem["status_reason"] = item.StatusReason
	respItem["reason_for_visit"] = item.ReasonForVisit
	respItem["sponsor_user_id"] = item.SponsorUserID
	respItem["sponsor_user_name"] = item.SponsorUserName
	respItem["guest_info"] = flattenGuestUserGetGuestUserByIDItemIDGuestInfo(item.GuestInfo)
	respItem["guest_access_info"] = flattenGuestUserGetGuestUserByIDItemIDGuestAccessInfo(item.GuestAccessInfo)
	respItem["portal_id"] = item.PortalID
	respItem["custom_fields"] = flattenGuestUserGetGuestUserByIDItemIDCustomFields(item.CustomFields)
	respItem["link"] = flattenGuestUserGetGuestUserByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenGuestUserGetGuestUserByIDItemIDGuestInfo(item *isegosdk.ResponseGuestUserGetGuestUserByIDGuestUserGuestInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["first_name"] = item.FirstName
	respItem["last_name"] = item.LastName
	respItem["company"] = item.Company
	respItem["creation_time"] = item.CreationTime
	respItem["notification_language"] = item.NotificationLanguage
	respItem["user_name"] = item.UserName
	respItem["email_address"] = item.EmailAddress
	respItem["phone_number"] = item.PhoneNumber
	respItem["password"] = item.Password
	respItem["enabled"] = boolPtrToString(item.Enabled)
	respItem["sms_service_provider"] = item.SmsServiceProvider

	return []map[string]interface{}{
		respItem,
	}

}

func flattenGuestUserGetGuestUserByIDItemIDGuestAccessInfo(item *isegosdk.ResponseGuestUserGetGuestUserByIDGuestUserGuestAccessInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["valid_days"] = item.ValidDays
	respItem["from_date"] = item.FromDate
	respItem["to_date"] = item.ToDate
	respItem["location"] = item.Location
	respItem["ssid"] = item.SSID
	respItem["group_tag"] = item.GroupTag

	return []map[string]interface{}{
		respItem,
	}

}

func flattenGuestUserGetGuestUserByIDItemIDCustomFields(item *isegosdk.ResponseGuestUserGetGuestUserByIDGuestUserCustomFields) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return respItem

}

func flattenGuestUserGetGuestUserByIDItemIDLink(item *isegosdk.ResponseGuestUserGetGuestUserByIDGuestUserLink) []map[string]interface{} {
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
