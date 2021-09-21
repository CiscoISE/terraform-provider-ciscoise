package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSxpVpns() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SXPVPNs.

- This resource deletes a SXP VPN.

- This resource creates a SXP VPN.
`,

		CreateContext: resourceSxpVpnsCreate,
		ReadContext:   resourceSxpVpnsRead,
		UpdateContext: resourceSxpVpnsUpdate,
		DeleteContext: resourceSxpVpnsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
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
						"sxp_vpn_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSxpVpnsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestSxpVpnsCreateSxpVpn(ctx, "item.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vName, _ := resourceItem["sxp_vpn_name"]
	vvID := interfaceToString(vID)
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.SxpVpns.GetSxpVpnByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		queryParams2 := isegosdk.GetSxpVpnsQueryParams{}

		response2, _, err := client.SxpVpns.GetSxpVpns(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsSxpVpnsGetSxpVpns(m, response2, &queryParams2)
			item2, err := searchSxpVpnsGetSxpVpns(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	restyResp1, err := client.SxpVpns.CreateSxpVpn(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSxpVpn", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSxpVpn", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["sxp_vpn_name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSxpVpnsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, _ := resourceMap["sxp_vpn_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)
	vvName := vName
	vvID := vID

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSxpVpns")
		queryParams1 := isegosdk.GetSxpVpnsQueryParams{}

		response1, _, err := client.SxpVpns.GetSxpVpns(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSxpVpns", err,
				"Failure at GetSxpVpns, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsSxpVpnsGetSxpVpns(m, response1, &queryParams1)
		item1, err := searchSxpVpnsGetSxpVpns(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetSxpVpns response", err,
				"Failure when searching item from GetSxpVpns, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSxpVpns search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSxpVpnByID")

		response2, _, err := client.SxpVpns.GetSxpVpnByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSxpVpnByID", err,
				"Failure at GetSxpVpnByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSxpVpnsGetSxpVpnByIDItem(response2.ERSSxpVpn)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSxpVpnByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSxpVpnsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSxpVpnsRead(ctx, d, m)
}

func resourceSxpVpnsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, _ := resourceMap["sxp_vpn_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {
		queryParams1 := isegosdk.GetSxpVpnsQueryParams{}

		getResp1, _, err := client.SxpVpns.GetSxpVpns(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSxpVpnsGetSxpVpns(m, getResp1, &queryParams1)
		item1, err := searchSxpVpnsGetSxpVpns(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 2 {
		vvID = vID
		getResp, _, err := client.SxpVpns.GetSxpVpnByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.SxpVpns.DeleteSxpVpnByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSxpVpnByID", err, restyResp1.String(),
				"Failure at DeleteSxpVpnByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSxpVpnByID", err,
			"Failure at DeleteSxpVpnByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSxpVpnsCreateSxpVpn(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpVpnsCreateSxpVpn {
	request := isegosdk.RequestSxpVpnsCreateSxpVpn{}
	request.ERSSxpVpn = expandRequestSxpVpnsCreateSxpVpnERSSxpVpn(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSxpVpnsCreateSxpVpnERSSxpVpn(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpVpnsCreateSxpVpnERSSxpVpn {
	request := isegosdk.RequestSxpVpnsCreateSxpVpnERSSxpVpn{}
	if v, ok := d.GetOkExists(key + ".sxp_vpn_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_vpn_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_vpn_name"))) {
		request.SxpVpnName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsSxpVpnsGetSxpVpns(m interface{}, response *isegosdk.ResponseSxpVpnsGetSxpVpns, queryParams *isegosdk.GetSxpVpnsQueryParams) []isegosdk.ResponseSxpVpnsGetSxpVpnsSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseSxpVpnsGetSxpVpnsSearchResultResources
	for response.SearchResult != nil && response.SearchResult.Resources != nil && len(*response.SearchResult.Resources) > 0 {
		respItems = append(respItems, *response.SearchResult.Resources...)
		if response.SearchResult.NextPage != nil && response.SearchResult.NextPage.Rel == "next" {
			href := response.SearchResult.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			if queryParams != nil {
				queryParams.Page = page
				queryParams.Size = size
			}
			response, _, err = client.SxpVpns.GetSxpVpns(queryParams)
			if err != nil {
				break
			}
			// All is good, continue to the next page
			continue
		}
		// Does not have next page finish iteration
		break
	}
	return respItems
}

func searchSxpVpnsGetSxpVpns(m interface{}, items []isegosdk.ResponseSxpVpnsGetSxpVpnsSearchResultResources, name string, id string) (*isegosdk.ResponseSxpVpnsGetSxpVpnByIDERSSxpVpn, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSxpVpnsGetSxpVpnByIDERSSxpVpn
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSxpVpnsGetSxpVpnByID
			getItem, _, err = client.SxpVpns.GetSxpVpnByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSxpVpnByID")
			}
			foundItem = getItem.ERSSxpVpn
			return foundItem, err
		}
	}
	return foundItem, err
}
