package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceActiveDirectory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceActiveDirectoryRead,
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
						"domain": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_domain_allowed_list": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enable_domain_white_list": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"adgroups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"groups": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"sid": &schema.Schema{
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
						"advanced_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_pass_change": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_machine_auth": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_machine_access": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"aging_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"enable_dialin_permission_check": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_callback_for_dialin_client": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"plaintext_auth": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_failed_auth_protection": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"auth_protection_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"failed_auth_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"identity_not_in_ad_behaviour": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"unreachable_domains_behaviour": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_rewrites": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"rewrite_rules": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"row_id": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rewrite_match": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"rewrite_result": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"department": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"organizational_unit": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"job_title": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"locality": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"email": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"state_or_province": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"telephone": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"country": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"street_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"schema": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"ad_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"internal_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"ad_scopes_names": &schema.Schema{
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
			"item_id": &schema.Schema{
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
						"domain": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_domain_white_list": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enable_domain_allowed_list": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"adgroups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"groups": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"sid": &schema.Schema{
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
						"advanced_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_pass_change": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_machine_auth": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_machine_access": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"aging_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"enable_dialin_permission_check": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_callback_for_dialin_client": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"plaintext_auth": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_failed_auth_protection": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"auth_protection_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"failed_auth_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"identity_not_in_ad_behaviour": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"unreachable_domains_behaviour": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_rewrites": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"rewrite_rules": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"row_id": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rewrite_match": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"rewrite_result": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"department": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"organizational_unit": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"job_title": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"locality": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"email": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"state_or_province": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"telephone": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"country": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"street_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"schema": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"ad_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"internal_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"ad_scopes_names": &schema.Schema{
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

func dataSourceActiveDirectoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetActiveDirectory")
		queryParams1 := isegosdk.GetActiveDirectoryQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, _, err := client.ActiveDirectory.GetActiveDirectory(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetActiveDirectory", err,
				"Failure at GetActiveDirectory, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseActiveDirectoryGetActiveDirectorySearchResultResources
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
				response1, _, err = client.ActiveDirectory.GetActiveDirectory(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenActiveDirectoryGetActiveDirectoryItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetActiveDirectory response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetActiveDirectoryByName")
		vvName := vName.(string)

		response2, _, err := client.ActiveDirectory.GetActiveDirectoryByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetActiveDirectoryByName", err,
				"Failure at GetActiveDirectoryByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemName2 := flattenActiveDirectoryGetActiveDirectoryByNameItemName(&response2.ERSActiveDirectory)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetActiveDirectoryByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetActiveDirectoryByID")
		vvID := vID.(string)

		response3, _, err := client.ActiveDirectory.GetActiveDirectoryByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetActiveDirectoryByID", err,
				"Failure at GetActiveDirectoryByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

		vItemID3 := flattenActiveDirectoryGetActiveDirectoryByIDItemID(&response3.ERSActiveDirectory)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetActiveDirectoryByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenActiveDirectoryGetActiveDirectoryItems(items *[]isegosdk.ResponseActiveDirectoryGetActiveDirectorySearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenActiveDirectoryGetActiveDirectoryItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenActiveDirectoryGetActiveDirectoryItemsLink(item isegosdk.ResponseActiveDirectoryGetActiveDirectorySearchResultResourcesLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemName(item *isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectory) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["domain"] = item.Domain
	respItem["enable_domain_allowed_list"] = item.EnableDomainAllowedList
	respItem["enable_domain_white_list"] = item.EnableDomainWhiteList
	respItem["adgroups"] = flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdgroups(item.Adgroups)
	respItem["advanced_settings"] = flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdvancedSettings(item.AdvancedSettings)
	respItem["ad_attributes"] = flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdAttributes(item.AdAttributes)
	respItem["ad_scopes_names"] = item.AdScopesNames
	respItem["link"] = flattenActiveDirectoryGetActiveDirectoryByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdgroups(item isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdgroups) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["groups"] = flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdgroupsGroups(item.Groups)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdgroupsGroups(items []isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdgroupsGroups) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["sid"] = item.Sid
		respItem["type"] = item.Type
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdvancedSettings(item isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdvancedSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enable_pass_change"] = item.EnablePassChange
	respItem["enable_machine_auth"] = item.EnableMachineAuth
	respItem["enable_machine_access"] = item.EnableMachineAccess
	respItem["aging_time"] = item.AgingTime
	respItem["enable_dialin_permission_check"] = item.EnableDialinPermissionCheck
	respItem["enable_callback_for_dialin_client"] = item.EnableCallbackForDialinClient
	respItem["plaintext_auth"] = item.PlaintextAuth
	respItem["enable_failed_auth_protection"] = item.EnableFailedAuthProtection
	respItem["auth_protection_type"] = item.AuthProtectionType
	respItem["failed_auth_threshold"] = item.FailedAuthThreshold
	respItem["identity_not_in_ad_behaviour"] = item.IDentityNotInAdBehaviour
	respItem["unreachable_domains_behaviour"] = item.UnreachableDomainsBehaviour
	respItem["enable_rewrites"] = item.EnableRewrites
	respItem["rewrite_rules"] = flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdvancedSettingsRewriteRules(item.RewriteRules)
	respItem["first_name"] = item.FirstName
	respItem["department"] = item.Department
	respItem["last_name"] = item.LastName
	respItem["organizational_unit"] = item.OrganizationalUnit
	respItem["job_title"] = item.JobTitle
	respItem["locality"] = item.Locality
	respItem["email"] = item.Email
	respItem["state_or_province"] = item.StateOrProvince
	respItem["telephone"] = item.Telephone
	respItem["country"] = item.Country
	respItem["street_address"] = item.StreetAddress
	respItem["schema"] = item.Schema

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdvancedSettingsRewriteRules(items []isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdvancedSettingsRewriteRules) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["row_id"] = item.RowID
		respItem["rewrite_match"] = item.RewriteMatch
		respItem["rewrite_result"] = item.RewriteResult
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdAttributes(item isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdAttributes) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["attributes"] = flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdAttributesAttributes(item.Attributes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdAttributesAttributes(items []isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdAttributesAttributes) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["internal_name"] = item.InternalName
		respItem["default_value"] = item.DefaultValue
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameLink(item isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemID(item *isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectory) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["domain"] = item.Domain
	respItem["enable_domain_white_list"] = item.EnableDomainWhiteList
	respItem["enable_domain_allowed_list"] = item.EnableDomainAllowedList
	respItem["adgroups"] = flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdgroups(item.Adgroups)
	respItem["advanced_settings"] = flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdvancedSettings(item.AdvancedSettings)
	respItem["ad_attributes"] = flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdAttributes(item.AdAttributes)
	respItem["ad_scopes_names"] = item.AdScopesNames
	respItem["link"] = flattenActiveDirectoryGetActiveDirectoryByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdgroups(item isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdgroups) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["groups"] = flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdgroupsGroups(item.Groups)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdgroupsGroups(items []isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdgroupsGroups) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["sid"] = item.Sid
		respItem["type"] = item.Type
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdvancedSettings(item isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdvancedSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enable_pass_change"] = item.EnablePassChange
	respItem["enable_machine_auth"] = item.EnableMachineAuth
	respItem["enable_machine_access"] = item.EnableMachineAccess
	respItem["aging_time"] = item.AgingTime
	respItem["enable_dialin_permission_check"] = item.EnableDialinPermissionCheck
	respItem["enable_callback_for_dialin_client"] = item.EnableCallbackForDialinClient
	respItem["plaintext_auth"] = item.PlaintextAuth
	respItem["enable_failed_auth_protection"] = item.EnableFailedAuthProtection
	respItem["auth_protection_type"] = item.AuthProtectionType
	respItem["failed_auth_threshold"] = item.FailedAuthThreshold
	respItem["identity_not_in_ad_behaviour"] = item.IDentityNotInAdBehaviour
	respItem["unreachable_domains_behaviour"] = item.UnreachableDomainsBehaviour
	respItem["enable_rewrites"] = item.EnableRewrites
	respItem["rewrite_rules"] = flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdvancedSettingsRewriteRules(item.RewriteRules)
	respItem["first_name"] = item.FirstName
	respItem["department"] = item.Department
	respItem["last_name"] = item.LastName
	respItem["organizational_unit"] = item.OrganizationalUnit
	respItem["job_title"] = item.JobTitle
	respItem["locality"] = item.Locality
	respItem["email"] = item.Email
	respItem["state_or_province"] = item.StateOrProvince
	respItem["telephone"] = item.Telephone
	respItem["country"] = item.Country
	respItem["street_address"] = item.StreetAddress
	respItem["schema"] = item.Schema

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdvancedSettingsRewriteRules(items []isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdvancedSettingsRewriteRules) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["row_id"] = item.RowID
		respItem["rewrite_match"] = item.RewriteMatch
		respItem["rewrite_result"] = item.RewriteResult
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdAttributes(item isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdAttributes) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["attributes"] = flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdAttributesAttributes(item.Attributes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdAttributesAttributes(items []isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdAttributesAttributes) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["internal_name"] = item.InternalName
		respItem["default_value"] = item.DefaultValue
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDLink(item isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
