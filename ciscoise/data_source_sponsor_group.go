package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSponsorGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSponsorGroupRead,
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

						"auto_notification": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"create_permissions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"can_create_random_accounts": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_import_multiple_accounts": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_set_future_start_date": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_specify_username_prefix": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"default_username_prefix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"import_batch_size_limit": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"random_batch_size_limit": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"start_date_future_limit_days": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"guest_types": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default_group": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_enabled": &schema.Schema{
							Type:     schema.TypeBool,
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
						"locations": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"manage_permission": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"member_groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"other_permissions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"can_access_via_rest": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_approve_selfreg_guests": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_delete_guest_accounts": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_extend_guest_accounts": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_reinstate_suspended_accounts": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_reset_guest_passwords": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_send_sms_notifications": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_suspend_guest_accounts": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_update_guest_contact_info": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"can_view_guest_passwords": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"limit_approval_to_sponsors_guests": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"require_suspension_reason": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
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

func dataSourceSponsorGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetSponsorGroup")
		queryParams1 := isegosdk.GetSponsorGroupQueryParams{}

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

		response1, _, err := client.SponsorGroup.GetSponsorGroup(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSponsorGroup", err,
				"Failure at GetSponsorGroup, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseSponsorGroupGetSponsorGroupSearchResultResources
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
				response1, _, err = client.SponsorGroup.GetSponsorGroup(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenSponsorGroupGetSponsorGroupItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsorGroup response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetSponsorGroupByID")
		vvID := vID.(string)

		response2, _, err := client.SponsorGroup.GetSponsorGroupByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSponsorGroupByID", err,
				"Failure at GetSponsorGroupByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenSponsorGroupGetSponsorGroupByIDItem(response2.SponsorGroup)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsorGroupByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSponsorGroupGetSponsorGroupItems(items *[]isegosdk.ResponseSponsorGroupGetSponsorGroupSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenSponsorGroupGetSponsorGroupItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSponsorGroupGetSponsorGroupItemsLink(item *isegosdk.ResponseSponsorGroupGetSponsorGroupSearchResultResourcesLink) []map[string]interface{} {
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

func flattenSponsorGroupGetSponsorGroupByIDItem(item *isegosdk.ResponseSponsorGroupGetSponsorGroupByIDSponsorGroup) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["is_enabled"] = item.IsEnabled
	respItem["is_default_group"] = item.IsDefaultGroup
	respItem["member_groups"] = item.MemberGroups
	respItem["guest_types"] = item.GuestTypes
	respItem["locations"] = item.Locations
	respItem["auto_notification"] = item.AutoNotification
	respItem["create_permissions"] = flattenSponsorGroupGetSponsorGroupByIDItemCreatePermissions(item.CreatePermissions)
	respItem["manage_permission"] = item.ManagePermission
	respItem["other_permissions"] = flattenSponsorGroupGetSponsorGroupByIDItemOtherPermissions(item.OtherPermissions)
	respItem["link"] = flattenSponsorGroupGetSponsorGroupByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSponsorGroupGetSponsorGroupByIDItemCreatePermissions(item *isegosdk.ResponseSponsorGroupGetSponsorGroupByIDSponsorGroupCreatePermissions) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["can_import_multiple_accounts"] = item.CanImportMultipleAccounts
	respItem["import_batch_size_limit"] = item.ImportBatchSizeLimit
	respItem["can_create_random_accounts"] = item.CanCreateRandomAccounts
	respItem["random_batch_size_limit"] = item.RandomBatchSizeLimit
	respItem["default_username_prefix"] = item.DefaultUsernamePrefix
	respItem["can_specify_username_prefix"] = item.CanSpecifyUsernamePrefix
	respItem["can_set_future_start_date"] = item.CanSetFutureStartDate
	respItem["start_date_future_limit_days"] = item.StartDateFutureLimitDays

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorGroupGetSponsorGroupByIDItemOtherPermissions(item *isegosdk.ResponseSponsorGroupGetSponsorGroupByIDSponsorGroupOtherPermissions) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["can_update_guest_contact_info"] = item.CanUpdateGuestContactInfo
	respItem["can_view_guest_passwords"] = item.CanViewGuestPasswords
	respItem["can_send_sms_notifications"] = item.CanSendSmsNotifications
	respItem["can_reset_guest_passwords"] = item.CanResetGuestPasswords
	respItem["can_extend_guest_accounts"] = item.CanExtendGuestAccounts
	respItem["can_delete_guest_accounts"] = item.CanDeleteGuestAccounts
	respItem["can_suspend_guest_accounts"] = item.CanSuspendGuestAccounts
	respItem["require_suspension_reason"] = item.RequireSuspensionReason
	respItem["can_reinstate_suspended_accounts"] = item.CanReinstateSuspendedAccounts
	respItem["can_approve_selfreg_guests"] = item.CanApproveSelfregGuests
	respItem["limit_approval_to_sponsors_guests"] = item.LimitApprovalToSponsorsGuests
	respItem["can_access_via_rest"] = item.CanAccessViaRest

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorGroupGetSponsorGroupByIDItemLink(item *isegosdk.ResponseSponsorGroupGetSponsorGroupByIDSponsorGroupLink) []map[string]interface{} {
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
