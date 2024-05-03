package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEndpoints() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on endpoints.

- Get all endpoints

- Get endpoint by id or MAC
`,

		ReadContext: dataSourceEndpointsRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Description: `filter query parameter. 
 
 
 
Simple filtering
 should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the 
'filterType=or'
 query string parameter. Each resource Data model description should specify if an attribute is a filtered field. 
 
 
 
 
 
OPERATOR
 
DESCRIPTION
 
 
 
 
 
EQ
 
Equals
 
 
 
NEQ
 
Not Equals
 
 
 
GT
 
Greater Than
 
 
 
LT
 
Less Then
 
 
 
STARTSW
 
Starts With
 
 
 
NSTARTSW
 
Not Starts With
 
 
 
ENDSW
 
Ends With
 
 
 
NENDSW
 
Not Ends With
 
 
 
CONTAINS
 
Contains
 
 
 
NCONTAINS
 
Not Contains
 
 
 
 `,
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter_type": &schema.Schema{
				Description: `filterType query parameter. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter`,
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
			"sort": &schema.Schema{
				Description: `sort query parameter. sort type asc or desc`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. sort column by which objects needs to be sorted`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"value": &schema.Schema{
				Description: `value path parameter. The id or MAC of the endpoint`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connected_links": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"custom_attributes": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hardware_revision": &schema.Schema{
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
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mdm_attributes": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"portal_user": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"product_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_revision": &schema.Schema{
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
						"vendor": &schema.Schema{
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

						"connected_links": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"custom_attributes": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hardware_revision": &schema.Schema{
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
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mdm_attributes": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"portal_user": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"product_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_revision": &schema.Schema{
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
						"vendor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEndpointsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSort, okSort := d.GetOk("sort")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vValue, okValue := d.GetOk("value")

	method1 := []bool{okPage, okSize, okSort, okSortBy, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okValue}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: List1")
		queryParams1 := isegosdk.List1QueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}
		if okSort {
			queryParams1.Sort = vSort.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okFilter {
			queryParams1.Filter = vFilter.(string)
		}
		if okFilterType {
			queryParams1.FilterType = vFilterType.(string)
		}

		response1, restyResp1, err := client.Endpoints.List1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 List1", err,
				"Failure at List1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEndpointsList1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting List1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: Get1")
		vvValue := vValue.(string)

		response2, restyResp2, err := client.Endpoints.Get1(vvValue)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 Get1", err,
				"Failure at Get1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenEndpointsGet1Item(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting Get1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEndpointsList1Items(items *isegosdk.ResponseEndpointsList1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["connected_links"] = flattenEndpointsList1ItemsConnectedLinks(item.ConnectedLinks)
		respItem["custom_attributes"] = flattenEndpointsList1ItemsCustomAttributes(item.CustomAttributes)
		respItem["description"] = item.Description
		respItem["device_type"] = item.DeviceType
		respItem["group_id"] = item.GroupID
		respItem["hardware_revision"] = item.HardwareRevision
		respItem["id"] = item.ID
		respItem["identity_store"] = item.IDentityStore
		respItem["identity_store_id"] = item.IDentityStoreID
		respItem["ip_address"] = item.IPAddress
		respItem["mac"] = item.Mac
		respItem["mdm_attributes"] = flattenEndpointsList1ItemsMdmAttributes(item.MdmAttributes)
		respItem["name"] = item.Name
		respItem["portal_user"] = item.PortalUser
		respItem["product_id"] = item.ProductID
		respItem["profile_id"] = item.ProfileID
		respItem["protocol"] = item.Protocol
		respItem["serial_number"] = item.SerialNumber
		respItem["software_revision"] = item.SoftwareRevision
		respItem["static_group_assignment"] = boolPtrToString(item.StaticGroupAssignment)
		respItem["static_profile_assignment"] = boolPtrToString(item.StaticProfileAssignment)
		respItem["vendor"] = item.Vendor
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEndpointsList1ItemsConnectedLinks(item *isegosdk.ResponseItemEndpointsList1ConnectedLinks) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenEndpointsList1ItemsCustomAttributes(item *isegosdk.ResponseItemEndpointsList1CustomAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenEndpointsList1ItemsMdmAttributes(item *isegosdk.ResponseItemEndpointsList1MdmAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenEndpointsGet1Item(item *isegosdk.ResponseEndpointsGet1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connected_links"] = flattenEndpointsGet1ItemConnectedLinks(item.ConnectedLinks)
	respItem["custom_attributes"] = flattenEndpointsGet1ItemCustomAttributes(item.CustomAttributes)
	respItem["description"] = item.Description
	respItem["device_type"] = item.DeviceType
	respItem["group_id"] = item.GroupID
	respItem["hardware_revision"] = item.HardwareRevision
	respItem["id"] = item.ID
	respItem["identity_store"] = item.IDentityStore
	respItem["identity_store_id"] = item.IDentityStoreID
	respItem["ip_address"] = item.IPAddress
	respItem["mac"] = item.Mac
	respItem["mdm_attributes"] = flattenEndpointsGet1ItemMdmAttributes(item.MdmAttributes)
	respItem["name"] = item.Name
	respItem["portal_user"] = item.PortalUser
	respItem["product_id"] = item.ProductID
	respItem["profile_id"] = item.ProfileID
	respItem["protocol"] = item.Protocol
	respItem["serial_number"] = item.SerialNumber
	respItem["software_revision"] = item.SoftwareRevision
	respItem["static_group_assignment"] = boolPtrToString(item.StaticGroupAssignment)
	respItem["static_profile_assignment"] = boolPtrToString(item.StaticProfileAssignment)
	respItem["vendor"] = item.Vendor
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEndpointsGet1ItemConnectedLinks(item *isegosdk.ResponseEndpointsGet1ConnectedLinks) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenEndpointsGet1ItemCustomAttributes(item *isegosdk.ResponseEndpointsGet1CustomAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenEndpointsGet1ItemMdmAttributes(item *isegosdk.ResponseEndpointsGet1MdmAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
