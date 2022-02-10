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

func resourceSxpLocalBindings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SXPLocalBindings.

- This resource allows the client to update a SXP local binding.

- This resource deletes a SXP local binding.

- This resource creates a SXP local binding.
`,

		CreateContext: resourceSxpLocalBindingsCreate,
		ReadContext:   resourceSxpLocalBindingsRead,
		UpdateContext: resourceSxpLocalBindingsUpdate,
		DeleteContext: resourceSxpLocalBindingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"binding_name": &schema.Schema{
							Description: `This field is depricated from Cisco ISE 3.0`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address_or_host": &schema.Schema{
							Description: `IP address for static mapping (hostname is not supported)`,
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
						"sgt": &schema.Schema{
							Description: `SGT name or ID`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sxp_vpn": &schema.Schema{
							Description: `List of SXP Domains, separated with comma. At least one of: sxpVpn or vns should be defined`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"vns": &schema.Schema{
							Description: `List of Virtual Networks, separated with comma. At least one of: sxpVpn or vns should be defined`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"binding_name": &schema.Schema{
							Description: `This field is depricated from Cisco ISE 3.0`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_address_or_host": &schema.Schema{
							Description: `IP address for static mapping (hostname is not supported)`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"sgt": &schema.Schema{
							Description:      `SGT name or ID`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSuppressSgt(),
						},
						"sxp_vpn": &schema.Schema{
							Description: `List of SXP Domains, separated with comma. At least one of: sxpVpn or vns should be defined`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"vns": &schema.Schema{
							Description: `List of Virtual Networks, separated with comma. At least one of: sxpVpn or vns should be defined`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSxpLocalBindingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SxpLocalBindings create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSxpLocalBindingsCreateSxpLocalBindings(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okID && vvID != "" {
		getResponse2, _, err := client.SxpLocalBindings.GetSxpLocalBindingsByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceSxpLocalBindingsRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetSxpLocalBindingsQueryParams{}

		response2, _, err := client.SxpLocalBindings.GetSxpLocalBindings(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsSxpLocalBindingsGetSxpLocalBindings(m, response2, &queryParams2)
			item2, err := searchSxpLocalBindingsGetSxpLocalBindings(m, items2, "", vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				d.SetId(joinResourceID(resourceMap))
				return resourceSxpLocalBindingsRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.SxpLocalBindings.CreateSxpLocalBindings(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSxpLocalBindings", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSxpLocalBindings", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceSxpLocalBindingsRead(ctx, d, m)
}

func resourceSxpLocalBindingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SxpLocalBindings read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetSxpLocalBindings")
		queryParams1 := isegosdk.GetSxpLocalBindingsQueryParams{}

		response1, restyResp1, err := client.SxpLocalBindings.GetSxpLocalBindings(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsSxpLocalBindingsGetSxpLocalBindings(m, response1, &queryParams1)
		item1, err := searchSxpLocalBindingsGetSxpLocalBindings(m, items1, "", vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenSxpLocalBindingsGetSxpLocalBindingsByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSxpLocalBindings search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSxpLocalBindingsByID")
		vvID := vID

		response2, restyResp2, err := client.SxpLocalBindings.GetSxpLocalBindingsByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSxpLocalBindingsGetSxpLocalBindingsByIDItem(response2.ERSSxpLocalBindings)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSxpLocalBindingsByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSxpLocalBindingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SxpLocalBindings update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestSxpLocalBindingsUpdateSxpLocalBindingsByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.SxpLocalBindings.UpdateSxpLocalBindingsByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSxpLocalBindingsByID", err, restyResp1.String(),
					"Failure at UpdateSxpLocalBindingsByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSxpLocalBindingsByID", err,
				"Failure at UpdateSxpLocalBindingsByID, unexpected response", ""))
			return diags
		}
		d.Set("last_updated", getUnixTimeString())
	}

	return resourceSxpLocalBindingsRead(ctx, d, m)
}

func resourceSxpLocalBindingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SxpLocalBindings delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {
		queryParams1 := isegosdk.GetSxpLocalBindingsQueryParams{}

		getResp1, _, err := client.SxpLocalBindings.GetSxpLocalBindings(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSxpLocalBindingsGetSxpLocalBindings(m, getResp1, &queryParams1)
		item1, err := searchSxpLocalBindingsGetSxpLocalBindings(m, items1, "", vID)
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
		getResp, _, err := client.SxpLocalBindings.GetSxpLocalBindingsByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.SxpLocalBindings.DeleteSxpLocalBindingsByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSxpLocalBindingsByID", err, restyResp1.String(),
				"Failure at DeleteSxpLocalBindingsByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSxpLocalBindingsByID", err,
			"Failure at DeleteSxpLocalBindingsByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSxpLocalBindingsCreateSxpLocalBindings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpLocalBindingsCreateSxpLocalBindings {
	request := isegosdk.RequestSxpLocalBindingsCreateSxpLocalBindings{}
	request.ERSSxpLocalBindings = expandRequestSxpLocalBindingsCreateSxpLocalBindingsERSSxpLocalBindings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSxpLocalBindingsCreateSxpLocalBindingsERSSxpLocalBindings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpLocalBindingsCreateSxpLocalBindingsERSSxpLocalBindings {
	request := isegosdk.RequestSxpLocalBindingsCreateSxpLocalBindingsERSSxpLocalBindings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding_name")))) {
		request.BindingName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_or_host")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_or_host")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_or_host")))) {
		request.IPAddressOrHost = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sxp_vpn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sxp_vpn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sxp_vpn")))) {
		request.SxpVpn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt")))) {
		first, _ := replaceRegExStrings(interfaceToString(v), "", `\s*\(.*\)$`, "")
		request.Sgt = first
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vns")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vns")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vns")))) {
		request.Vns = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSxpLocalBindingsUpdateSxpLocalBindingsByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpLocalBindingsUpdateSxpLocalBindingsByID {
	request := isegosdk.RequestSxpLocalBindingsUpdateSxpLocalBindingsByID{}
	request.ERSSxpLocalBindings = expandRequestSxpLocalBindingsUpdateSxpLocalBindingsByIDERSSxpLocalBindings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSxpLocalBindingsUpdateSxpLocalBindingsByIDERSSxpLocalBindings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpLocalBindingsUpdateSxpLocalBindingsByIDERSSxpLocalBindings {
	request := isegosdk.RequestSxpLocalBindingsUpdateSxpLocalBindingsByIDERSSxpLocalBindings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding_name")))) {
		request.BindingName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_or_host")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_or_host")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_or_host")))) {
		request.IPAddressOrHost = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sxp_vpn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sxp_vpn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sxp_vpn")))) {
		request.SxpVpn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt")))) {
		first, _ := replaceRegExStrings(interfaceToString(v), "", `\s*\(.*\)$`, "")
		request.Sgt = first
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vns")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vns")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vns")))) {
		request.Vns = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsSxpLocalBindingsGetSxpLocalBindings(m interface{}, response *isegosdk.ResponseSxpLocalBindingsGetSxpLocalBindings, queryParams *isegosdk.GetSxpLocalBindingsQueryParams) []isegosdk.ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultResources
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
			response, _, err = client.SxpLocalBindings.GetSxpLocalBindings(queryParams)
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

func searchSxpLocalBindingsGetSxpLocalBindings(m interface{}, items []isegosdk.ResponseSxpLocalBindingsGetSxpLocalBindingsSearchResultResources, name string, id string) (*isegosdk.ResponseSxpLocalBindingsGetSxpLocalBindingsByIDERSSxpLocalBindings, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSxpLocalBindingsGetSxpLocalBindingsByIDERSSxpLocalBindings
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSxpLocalBindingsGetSxpLocalBindingsByID
			getItem, _, err = client.SxpLocalBindings.GetSxpLocalBindingsByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSxpLocalBindingsByID")
			}
			foundItem = getItem.ERSSxpLocalBindings
			return foundItem, err
		}
	}
	return foundItem, err
}
