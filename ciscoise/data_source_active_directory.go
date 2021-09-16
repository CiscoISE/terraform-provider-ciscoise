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
		Description: `It performs read operation on ActiveDirectory.

- This data source allows the client to get Active Directory by name.
- This data source fetchs the join point details by ID. The ID can be retrieved with the Get All operation.
- This data source lists all the join points for Active Directory domains in Cisco ISE.`,

		ReadContext: dataSourceActiveDirectoryRead,
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

						"ad_attributes": &schema.Schema{
							Description: `Holds list of AD Attributes`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Description: `List of Attributes`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_value": &schema.Schema{
													Description: `Required for each attribute in the attribute list. Can contain an empty string. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"internal_name": &schema.Schema{
													Description: `Required for each attribute in the attribute list. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Required for each attribute in the attribute list with no duplication between attributes. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Required for each group in the group list. Allowed values: STRING, IP, BOOLEAN, INT, OCTET_STRING`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"ad_scopes_names": &schema.Schema{
							Description: `String that contains the names of the scopes that the active directory belongs to. Names are separated by comma. Alphanumeric, underscore (_) characters are allowed`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"adgroups": &schema.Schema{
							Description: `Holds list of AD Groups`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"groups": &schema.Schema{
										Description: `List of Groups`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Required for each group in the group list with no duplication between groups. All characters are allowed except %`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"sid": &schema.Schema{
													Description: `Cisco ISE uses security identifiers (SIDs) for optimization of group membership evaluation. SIDs are useful for efficiency (speed) when the groups are evaluated. All characters are allowed except %`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `No character restriction`,
													Type:        schema.TypeString,
													Computed:    true,
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

									"aging_time": &schema.Schema{
										Description: `Range 1-8760 hours`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"auth_protection_type": &schema.Schema{
										Description: `Enable prevent AD account lockout. Allowed values:
- WIRELESS,
- WIRED,
- BOTH`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"country": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"department": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"email": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"enable_callback_for_dialin_client": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_dialin_permission_check": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_failed_auth_protection": &schema.Schema{
										Description: `Enable prevent AD account lockout due to too many bad password attempts`,
										Type:        schema.TypeBool,
										Computed:    true,
									},
									"enable_machine_access": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_machine_auth": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_pass_change": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_rewrites": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"failed_auth_threshold": &schema.Schema{
										Description: `Number of bad password attempts`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"first_name": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"identity_not_in_ad_behaviour": &schema.Schema{
										Description: `Allowed values: REJECT, SEARCH_JOINED_FOREST, SEARCH_ALL`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"job_title": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"last_name": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"locality": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"organizational_unit": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"plaintext_auth": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"rewrite_rules": &schema.Schema{
										Description: `Identity rewrite is an advanced feature that directs Cisco ISE to manipulate the identity
before it is passed to the external Active Directory system. You can create rules to change
the identity to a desired format that includes or excludes a domain prefix and/or suffix or
other additional markup of your choice`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"rewrite_match": &schema.Schema{
													Description: `Required for each rule in the list with no duplication between rules. All characters are allowed except %"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"rewrite_result": &schema.Schema{
													Description: `Required for each rule in the list. All characters are allowed except %"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"row_id": &schema.Schema{
													Description: `Required for each rule in the list in serial order`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
											},
										},
									},
									"schema": &schema.Schema{
										Description: `Allowed values: ACTIVE_DIRECTORY, CUSTOM.
Choose ACTIVE_DIRECTORY schema when the AD attributes defined in AD can be copied to relevant attributes
in Cisco ISE. If customization is needed, choose CUSTOM schema. All User info attributes are always set to
default value if schema is ACTIVE_DIRECTORY. Values can be changed only for CUSTOM schema`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"state_or_province": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"street_address": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"telephone": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"unreachable_domains_behaviour": &schema.Schema{
										Description: `Allowed values: PROCEED, DROP`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Description: `No character restriction`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"domain": &schema.Schema{
							Description: `The AD domain. Alphanumeric, hyphen (-) and dot (.) characters are allowed`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"enable_domain_allowed_list": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enable_domain_white_list": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Resource UUID value`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `Resource Name. Maximum 32 characters allowed. Allowed characters are alphanumeric and .-_/\\ characters`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ad_attributes": &schema.Schema{
							Description: `Holds list of AD Attributes`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Description: `List of Attributes`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_value": &schema.Schema{
													Description: `Required for each attribute in the attribute list. Can contain an empty string. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"internal_name": &schema.Schema{
													Description: `Required for each attribute in the attribute list. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Required for each attribute in the attribute list with no duplication between attributes. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Required for each group in the group list. Allowed values: STRING, IP, BOOLEAN, INT, OCTET_STRING`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"ad_scopes_names": &schema.Schema{
							Description: `String that contains the names of the scopes that the active directory belongs to. Names are separated by comma. Alphanumeric, underscore (_) characters are allowed`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"adgroups": &schema.Schema{
							Description: `Holds list of AD Groups`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"groups": &schema.Schema{
										Description: `List of Groups`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Required for each group in the group list with no duplication between groups. All characters are allowed except %`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"sid": &schema.Schema{
													Description: `Cisco ISE uses security identifiers (SIDs) for optimization of group membership evaluation. SIDs are useful for efficiency (speed) when the groups are evaluated. All characters are allowed except %`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `No character restriction`,
													Type:        schema.TypeString,
													Computed:    true,
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

									"aging_time": &schema.Schema{
										Description: `Range 1-8760 hours`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"auth_protection_type": &schema.Schema{
										Description: `Enable prevent AD account lockout. Allowed values:
- WIRELESS,
- WIRED,
- BOTH`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"country": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"department": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"email": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"enable_callback_for_dialin_client": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_dialin_permission_check": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_failed_auth_protection": &schema.Schema{
										Description: `Enable prevent AD account lockout due to too many bad password attempts`,
										Type:        schema.TypeBool,
										Computed:    true,
									},
									"enable_machine_access": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_machine_auth": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_pass_change": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_rewrites": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"failed_auth_threshold": &schema.Schema{
										Description: `Number of bad password attempts`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"first_name": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"identity_not_in_ad_behaviour": &schema.Schema{
										Description: `Allowed values: REJECT, SEARCH_JOINED_FOREST, SEARCH_ALL`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"job_title": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"last_name": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"locality": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"organizational_unit": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"plaintext_auth": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"rewrite_rules": &schema.Schema{
										Description: `Identity rewrite is an advanced feature that directs Cisco ISE to manipulate the identity
before it is passed to the external Active Directory system. You can create rules to change
the identity to a desired format that includes or excludes a domain prefix and/or suffix or
other additional markup of your choice`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"rewrite_match": &schema.Schema{
													Description: `Required for each rule in the list with no duplication between rules. All characters are allowed except %"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"rewrite_result": &schema.Schema{
													Description: `Required for each rule in the list. All characters are allowed except %"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"row_id": &schema.Schema{
													Description: `Required for each rule in the list in serial order`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
											},
										},
									},
									"schema": &schema.Schema{
										Description: `Allowed values: ACTIVE_DIRECTORY, CUSTOM.
Choose ACTIVE_DIRECTORY schema when the AD attributes defined in AD can be copied to relevant attributes
in Cisco ISE. If customization is needed, choose CUSTOM schema. All User info attributes are always set to
default value if schema is ACTIVE_DIRECTORY. Values can be changed only for CUSTOM schema`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"state_or_province": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"street_address": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"telephone": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"unreachable_domains_behaviour": &schema.Schema{
										Description: `Allowed values: PROCEED, DROP`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Description: `No character restriction`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"domain": &schema.Schema{
							Description: `The AD domain. Alphanumeric, hyphen (-) and dot (.) characters are allowed`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"enable_domain_allowed_list": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enable_domain_white_list": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Resource UUID value`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `Resource Name. Maximum 32 characters allowed. Allowed characters are alphanumeric and .-_/\\ characters`,
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

func dataSourceActiveDirectoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %v", method3)

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

		vItemName2 := flattenActiveDirectoryGetActiveDirectoryByNameItemName(response2.ERSActiveDirectory)
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

		vItemID3 := flattenActiveDirectoryGetActiveDirectoryByIDItemID(response3.ERSActiveDirectory)
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

func flattenActiveDirectoryGetActiveDirectoryItemsLink(item *isegosdk.ResponseActiveDirectoryGetActiveDirectorySearchResultResourcesLink) []map[string]interface{} {
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

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdgroups(item *isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdgroups) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["groups"] = flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdgroupsGroups(item.Groups)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdgroupsGroups(items *[]isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdgroupsGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["sid"] = item.Sid
		respItem["type"] = item.Type
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdvancedSettings(item *isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdvancedSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
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

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdvancedSettingsRewriteRules(items *[]isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdvancedSettingsRewriteRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["row_id"] = item.RowID
		respItem["rewrite_match"] = item.RewriteMatch
		respItem["rewrite_result"] = item.RewriteResult
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdAttributes(item *isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attributes"] = flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdAttributesAttributes(item.Attributes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameAdAttributesAttributes(items *[]isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryAdAttributesAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["internal_name"] = item.InternalName
		respItem["default_value"] = item.DefaultValue
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByNameItemNameLink(item *isegosdk.ResponseActiveDirectoryGetActiveDirectoryByNameERSActiveDirectoryLink) []map[string]interface{} {
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

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdgroups(item *isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdgroups) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["groups"] = flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdgroupsGroups(item.Groups)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdgroupsGroups(items *[]isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdgroupsGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["sid"] = item.Sid
		respItem["type"] = item.Type
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdvancedSettings(item *isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdvancedSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
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

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdvancedSettingsRewriteRules(items *[]isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdvancedSettingsRewriteRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["row_id"] = item.RowID
		respItem["rewrite_match"] = item.RewriteMatch
		respItem["rewrite_result"] = item.RewriteResult
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdAttributes(item *isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attributes"] = flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdAttributesAttributes(item.Attributes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDAdAttributesAttributes(items *[]isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryAdAttributesAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["internal_name"] = item.InternalName
		respItem["default_value"] = item.DefaultValue
	}
	return respItems

}

func flattenActiveDirectoryGetActiveDirectoryByIDItemIDLink(item *isegosdk.ResponseActiveDirectoryGetActiveDirectoryByIDERSActiveDirectoryLink) []map[string]interface{} {
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
