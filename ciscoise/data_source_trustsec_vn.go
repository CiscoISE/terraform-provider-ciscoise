package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTrustsecVn() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on virtualNetwork.

- Get all Virtual Networks

- Get Virtual Network by id
`,

		ReadContext: dataSourceTrustsecVnRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Description: `filter query parameter. 
 
 
 
Simple filtering
 should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the 
"filterType=or"
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
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_attributes": &schema.Schema{
							Description: `JSON String of additional attributes for the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier of the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_update": &schema.Schema{
							Description: `Timestamp for the last update of the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name of the Virtual Network`,
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

						"additional_attributes": &schema.Schema{
							Description: `JSON String of additional attributes for the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier of the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_update": &schema.Schema{
							Description: `Timestamp for the last update of the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name of the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTrustsecVnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSort, okSort := d.GetOk("sort")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okSort, okSortBy, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetVirtualNetworks")
		queryParams1 := isegosdk.GetVirtualNetworksQueryParams{}

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
			queryParams1.Filter = interfaceToSliceString(vFilter)
		}
		if okFilterType {
			queryParams1.FilterType = vFilterType.(string)
		}

		response1, restyResp1, err := client.VirtualNetwork.GetVirtualNetworks(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetVirtualNetworks", err,
				"Failure at GetVirtualNetworks, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenVirtualNetworkGetVirtualNetworksItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVirtualNetworks response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetVirtualNetworkByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.VirtualNetwork.GetVirtualNetworkByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetVirtualNetworkByID", err,
				"Failure at GetVirtualNetworkByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenVirtualNetworkGetVirtualNetworkByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVirtualNetworkByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenVirtualNetworkGetVirtualNetworksItems(items *[]isegosdk.ResponseVirtualNetworkGetVirtualNetworksResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["additional_attributes"] = item.AdditionalAttributes
		respItem["id"] = item.ID
		respItem["last_update"] = item.LastUpdate
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenVirtualNetworkGetVirtualNetworkByIDItem(items *[]isegosdk.ResponseVirtualNetworkGetVirtualNetworkByIDResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["additional_attributes"] = item.AdditionalAttributes
		respItem["id"] = item.ID
		respItem["last_update"] = item.LastUpdate
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
