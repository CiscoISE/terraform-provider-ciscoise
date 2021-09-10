package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSxpLocalBindings() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSxpLocalBindingsCreate,
		ReadContext:   resourceSxpLocalBindingsRead,
		UpdateContext: resourceSxpLocalBindingsUpdate,
		DeleteContext: resourceSxpLocalBindingsDelete,
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

						"binding_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_address_or_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:             schema.TypeList,
							DiffSuppressFunc: diffSuppressAlways(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
									"rel": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
									"type": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
								},
							},
						},
						"sgt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sxp_vpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vns": &schema.Schema{
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

func resourceSxpLocalBindingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestSxpLocalBindingsCreateSxpLocalBindings(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okID && vvID != "" {
		getResponse2, _, err := client.SxpLocalBindings.GetSxpLocalBindingsByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return diags
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
				return diags
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
	return diags
}

func resourceSxpLocalBindingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetSxpLocalBindings")
		queryParams1 := isegosdk.GetSxpLocalBindingsQueryParams{}

		response1, _, err := client.SxpLocalBindings.GetSxpLocalBindings(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSxpLocalBindings", err,
				"Failure at GetSxpLocalBindings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsSxpLocalBindingsGetSxpLocalBindings(m, response1, &queryParams1)
		item1, err := searchSxpLocalBindingsGetSxpLocalBindings(m, items1, "", vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetSxpLocalBindings response", err,
				"Failure when searching item from GetSxpLocalBindings, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSxpLocalBindings search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSxpLocalBindingsByID")
		vvID := vID

		response2, _, err := client.SxpLocalBindings.GetSxpLocalBindingsByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSxpLocalBindingsByID", err,
				"Failure at GetSxpLocalBindingsByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

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
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	if selectedMethod == 2 {
		vvID = vID
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestSxpLocalBindingsUpdateSxpLocalBindingsByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SxpLocalBindings.UpdateSxpLocalBindingsByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
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
	}

	return resourceSxpLocalBindingsRead(ctx, d, m)
}

func resourceSxpLocalBindingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
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
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
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
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".binding_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".binding_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".binding_name"))) {
		request.BindingName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ip_address_or_host"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ip_address_or_host"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ip_address_or_host"))) {
		request.IPAddressOrHost = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_vpn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_vpn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_vpn"))) {
		request.SxpVpn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sgt"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sgt"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sgt"))) {
		request.Sgt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".vns"); !isEmptyValue(reflect.ValueOf(d.Get(key+".vns"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".vns"))) {
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
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".binding_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".binding_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".binding_name"))) {
		request.BindingName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ip_address_or_host"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ip_address_or_host"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ip_address_or_host"))) {
		request.IPAddressOrHost = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_vpn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_vpn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_vpn"))) {
		request.SxpVpn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sgt"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sgt"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sgt"))) {
		request.Sgt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".vns"); !isEmptyValue(reflect.ValueOf(d.Get(key+".vns"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".vns"))) {
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
