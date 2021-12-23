package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEndpoint() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on endpoint.

- This data source allows the client to get an endpoint by name.

- This data source allows the client to get an endpoint by ID.

- This data source allows the client to get all the endpoints.

Filter:
Filters can be used to filter out Endpoints based on a set of attributes. This data source currently provides the
following filters:
[logicalProfileName, portalUser, staticProfileAssignment, profileId, profile, groupId, staticGroupAssignment, mac]

Example 1:

The
logicalProfileName
 filter can be used to get Enpoints that belong  to a specific Logical Profile. The supported operator for
logicalProfileNamefilter is EQ (equal to). The syntax to invoke the API with this filter:

/ers/config/endpoint?filter={filter name}.{operator}.{logical profile name}

Example:

https://{ise-ip}:9060/ers/config/endpoint?filter=logicalProfileName.EQ.LP_Apple

Example 2:

Sorting:
[name, description]
`,

		ReadContext: dataSourceEndpointRead,
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

						"custom_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_attributes": &schema.Schema{
										Description: `Key value map`,
										// CHECK: The type of this param
										// Replaced List to Map
										Type:     schema.TypeMap,
										Computed: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_store": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_store_id": &schema.Schema{
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
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mdm_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"mdm_compliance_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_encrypted": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_enrolled": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_ime_i": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_jail_broken": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_manufacturer": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_model": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_os": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_phone_number": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_pinlock": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_reachable": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_serial": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_server_name": &schema.Schema{
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
						"portal_user": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"static_group_assignment": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"static_profile_assignment": &schema.Schema{
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

						"custom_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_attributes": &schema.Schema{
										Description: `Key value map`,
										// CHECK: The type of this param
										// Replaced List to Map
										Type:     schema.TypeMap,
										Computed: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_store": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_store_id": &schema.Schema{
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
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mdm_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"mdm_compliance_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_encrypted": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_enrolled": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_ime_i": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_jail_broken": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_manufacturer": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_model": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_os": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_phone_number": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_pinlock": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_reachable": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_serial": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_server_name": &schema.Schema{
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
						"portal_user": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"static_group_assignment": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"static_profile_assignment": &schema.Schema{
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

func dataSourceEndpointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEndpoints")
		queryParams1 := isegosdk.GetEndpointsQueryParams{}

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

		response1, restyResp1, err := client.Endpoint.GetEndpoints(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEndpoints", err,
				"Failure at GetEndpoints, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseEndpointGetEndpointsSearchResultResources
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
				response1, _, err = client.Endpoint.GetEndpoints(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenEndpointGetEndpointsItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEndpoints response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetEndpointByName")
		vvName := vName.(string)

		response2, restyResp2, err := client.Endpoint.GetEndpointByName(vvName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEndpointByName", err,
				"Failure at GetEndpointByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemName2 := flattenEndpointGetEndpointByNameItemName(response2.ERSEndPoint)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEndpointByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method: GetEndpointByID")
		vvID := vID.(string)

		response3, restyResp3, err := client.Endpoint.GetEndpointByID(vvID)

		if err != nil || response3 == nil {
			if restyResp3 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp3.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEndpointByID", err,
				"Failure at GetEndpointByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response3))

		vItemID3 := flattenEndpointGetEndpointByIDItemID(response3.ERSEndPoint)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEndpointByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEndpointGetEndpointsItems(items *[]isegosdk.ResponseEndpointGetEndpointsSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenEndpointGetEndpointsItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEndpointGetEndpointsItemsLink(item *isegosdk.ResponseEndpointGetEndpointsSearchResultResourcesLink) []map[string]interface{} {
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

func flattenEndpointGetEndpointByNameItemName(item *isegosdk.ResponseEndpointGetEndpointByNameERSEndPoint) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["mac"] = item.Mac
	respItem["profile_id"] = item.ProfileID
	respItem["static_profile_assignment"] = boolPtrToString(item.StaticProfileAssignment)
	respItem["group_id"] = item.GroupID
	respItem["static_group_assignment"] = boolPtrToString(item.StaticGroupAssignment)
	respItem["portal_user"] = item.PortalUser
	respItem["identity_store"] = item.IDentityStore
	respItem["identity_store_id"] = item.IDentityStoreID
	respItem["mdm_attributes"] = flattenEndpointGetEndpointByNameItemNameMdmAttributes(item.MdmAttributes)
	respItem["custom_attributes"] = flattenEndpointGetEndpointByNameItemNameCustomAttributes(item.CustomAttributes)
	respItem["link"] = flattenEndpointGetEndpointByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEndpointGetEndpointByNameItemNameMdmAttributes(item *isegosdk.ResponseEndpointGetEndpointByNameERSEndPointMdmAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["mdm_server_name"] = item.MdmServerName
	respItem["mdm_reachable"] = boolPtrToString(item.MdmReachable)
	respItem["mdm_enrolled"] = boolPtrToString(item.MdmEnrolled)
	respItem["mdm_compliance_status"] = boolPtrToString(item.MdmComplianceStatus)
	respItem["mdm_os"] = item.MdmOS
	respItem["mdm_manufacturer"] = item.MdmManufacturer
	respItem["mdm_model"] = item.MdmModel
	respItem["mdm_serial"] = item.MdmSerial
	respItem["mdm_encrypted"] = boolPtrToString(item.MdmEncrypted)
	respItem["mdm_pinlock"] = boolPtrToString(item.MdmPinlock)
	respItem["mdm_jail_broken"] = boolPtrToString(item.MdmJailBroken)
	respItem["mdm_ime_i"] = item.MdmIMEI
	respItem["mdm_phone_number"] = item.MdmPhoneNumber

	return []map[string]interface{}{
		respItem,
	}

}

func flattenEndpointGetEndpointByNameItemNameCustomAttributes(item *isegosdk.ResponseEndpointGetEndpointByNameERSEndPointCustomAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["custom_attributes"] = flattenEndpointGetEndpointByNameItemNameCustomAttributesCustomAttributes(item.CustomAttributes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenEndpointGetEndpointByNameItemNameCustomAttributesCustomAttributes(item *isegosdk.ResponseEndpointGetEndpointByNameERSEndPointCustomAttributesCustomAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return respItem

}

func flattenEndpointGetEndpointByNameItemNameLink(item *isegosdk.ResponseEndpointGetEndpointByNameERSEndPointLink) []map[string]interface{} {
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

func flattenEndpointGetEndpointByIDItemID(item *isegosdk.ResponseEndpointGetEndpointByIDERSEndPoint) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["mac"] = item.Mac
	respItem["profile_id"] = item.ProfileID
	respItem["static_profile_assignment"] = boolPtrToString(item.StaticProfileAssignment)
	respItem["group_id"] = item.GroupID
	respItem["static_group_assignment"] = boolPtrToString(item.StaticGroupAssignment)
	respItem["portal_user"] = item.PortalUser
	respItem["identity_store"] = item.IDentityStore
	respItem["identity_store_id"] = item.IDentityStoreID
	respItem["mdm_attributes"] = flattenEndpointGetEndpointByIDItemIDMdmAttributes(item.MdmAttributes)
	respItem["custom_attributes"] = flattenEndpointGetEndpointByIDItemIDCustomAttributes(item.CustomAttributes)
	respItem["link"] = flattenEndpointGetEndpointByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEndpointGetEndpointByIDItemIDMdmAttributes(item *isegosdk.ResponseEndpointGetEndpointByIDERSEndPointMdmAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["mdm_server_name"] = item.MdmServerName
	respItem["mdm_reachable"] = boolPtrToString(item.MdmReachable)
	respItem["mdm_enrolled"] = boolPtrToString(item.MdmEnrolled)
	respItem["mdm_compliance_status"] = boolPtrToString(item.MdmComplianceStatus)
	respItem["mdm_os"] = item.MdmOS
	respItem["mdm_manufacturer"] = item.MdmManufacturer
	respItem["mdm_model"] = item.MdmModel
	respItem["mdm_serial"] = item.MdmSerial
	respItem["mdm_encrypted"] = boolPtrToString(item.MdmEncrypted)
	respItem["mdm_pinlock"] = boolPtrToString(item.MdmPinlock)
	respItem["mdm_jail_broken"] = boolPtrToString(item.MdmJailBroken)
	respItem["mdm_ime_i"] = item.MdmIMEI
	respItem["mdm_phone_number"] = item.MdmPhoneNumber

	return []map[string]interface{}{
		respItem,
	}

}

func flattenEndpointGetEndpointByIDItemIDCustomAttributes(item *isegosdk.ResponseEndpointGetEndpointByIDERSEndPointCustomAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["custom_attributes"] = flattenEndpointGetEndpointByIDItemIDCustomAttributesCustomAttributes(item.CustomAttributes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenEndpointGetEndpointByIDItemIDCustomAttributesCustomAttributes(item *isegosdk.ResponseEndpointGetEndpointByIDERSEndPointCustomAttributesCustomAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return respItem

}

func flattenEndpointGetEndpointByIDItemIDLink(item *isegosdk.ResponseEndpointGetEndpointByIDERSEndPointLink) []map[string]interface{} {
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
