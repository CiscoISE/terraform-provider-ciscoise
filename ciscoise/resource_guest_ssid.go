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

func resourceGuestSSID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on GuestSSID.

- This resource allows the client to update a guest SSID by ID.

- This resource deletes a guest SSID by ID.

- This resource creates a guest SSID.
`,

		CreateContext: resourceGuestSSIDCreate,
		ReadContext:   resourceGuestSSIDRead,
		UpdateContext: resourceGuestSSIDUpdate,
		DeleteContext: resourceGuestSSIDDelete,
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
						"name": &schema.Schema{
							Description: `Resource Name. Name may contain alphanumeric or any of the following characters [_.-]`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Resource Name. Name may contain alphanumeric or any of the following characters [_.-]`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceGuestSSIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestSSID create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGuestSSIDCreateGuestSSID(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.GuestSSID.GetGuestSSIDByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceGuestSSIDRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetGuestSSIDQueryParams{}

		response2, _, err := client.GuestSSID.GetGuestSSID(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsGuestSSIDGetGuestSSID(m, response2, &queryParams2)
			item2, err := searchGuestSSIDGetGuestSSID(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceGuestSSIDRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.GuestSSID.CreateGuestSSID(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateGuestSSID", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateGuestSSID", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceGuestSSIDRead(ctx, d, m)
}

func resourceGuestSSIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestSSID read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		vvName := vName
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetGuestSSID")
		queryParams1 := isegosdk.GetGuestSSIDQueryParams{}

		response1, restyResp1, err := client.GuestSSID.GetGuestSSID(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsGuestSSIDGetGuestSSID(m, response1, &queryParams1)
		item1, err := searchGuestSSIDGetGuestSSID(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenGuestSSIDGetGuestSSIDByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestSSID search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetGuestSSIDByID")
		vvID := vID

		response2, restyResp2, err := client.GuestSSID.GetGuestSSIDByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenGuestSSIDGetGuestSSIDByIDItem(response2.GuestSSID)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestSSIDByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceGuestSSIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestSSID update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetGuestSSIDQueryParams{}

		getResp1, _, err := client.GuestSSID.GetGuestSSID(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsGuestSSIDGetGuestSSID(m, getResp1, &queryParams1)
			item1, err := searchGuestSSIDGetGuestSSID(m, items1, vName, vID)
			if err == nil && item1 != nil {
				if vID != item1.ID {
					vvID = item1.ID
				} else {
					vvID = vID
				}
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestGuestSSIDUpdateGuestSSIDByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.GuestSSID.UpdateGuestSSIDByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateGuestSSIDByID", err, restyResp1.String(),
					"Failure at UpdateGuestSSIDByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGuestSSIDByID", err,
				"Failure at UpdateGuestSSIDByID, unexpected response", ""))
			return diags
		}
	}

	return resourceGuestSSIDRead(ctx, d, m)
}

func resourceGuestSSIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestSSID delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetGuestSSIDQueryParams{}

		getResp1, _, err := client.GuestSSID.GetGuestSSID(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsGuestSSIDGetGuestSSID(m, getResp1, &queryParams1)
		item1, err := searchGuestSSIDGetGuestSSID(m, items1, vName, vID)
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
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.GuestSSID.GetGuestSSIDByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.GuestSSID.DeleteGuestSSIDByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteGuestSSIDByID", err, restyResp1.String(),
				"Failure at DeleteGuestSSIDByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteGuestSSIDByID", err,
			"Failure at DeleteGuestSSIDByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestGuestSSIDCreateGuestSSID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestSSIDCreateGuestSSID {
	request := isegosdk.RequestGuestSSIDCreateGuestSSID{}
	request.GuestSSID = expandRequestGuestSSIDCreateGuestSSIDGuestSSID(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestSSIDCreateGuestSSIDGuestSSID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestSSIDCreateGuestSSIDGuestSSID {
	request := isegosdk.RequestGuestSSIDCreateGuestSSIDGuestSSID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestSSIDUpdateGuestSSIDByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestSSIDUpdateGuestSSIDByID {
	request := isegosdk.RequestGuestSSIDUpdateGuestSSIDByID{}
	request.GuestSSID = expandRequestGuestSSIDUpdateGuestSSIDByIDGuestSSID(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestSSIDUpdateGuestSSIDByIDGuestSSID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestSSIDUpdateGuestSSIDByIDGuestSSID {
	request := isegosdk.RequestGuestSSIDUpdateGuestSSIDByIDGuestSSID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsGuestSSIDGetGuestSSID(m interface{}, response *isegosdk.ResponseGuestSSIDGetGuestSSID, queryParams *isegosdk.GetGuestSSIDQueryParams) []isegosdk.ResponseGuestSSIDGetGuestSSIDSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseGuestSSIDGetGuestSSIDSearchResultResources
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
			response, _, err = client.GuestSSID.GetGuestSSID(queryParams)
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

func searchGuestSSIDGetGuestSSID(m interface{}, items []isegosdk.ResponseGuestSSIDGetGuestSSIDSearchResultResources, name string, id string) (*isegosdk.ResponseGuestSSIDGetGuestSSIDByIDGuestSSID, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseGuestSSIDGetGuestSSIDByIDGuestSSID
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseGuestSSIDGetGuestSSIDByID
			getItem, _, err = client.GuestSSID.GetGuestSSIDByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetGuestSSIDByID")
			}
			foundItem = getItem.GuestSSID
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseGuestSSIDGetGuestSSIDByID
			getItem, _, err = client.GuestSSID.GetGuestSSIDByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetGuestSSIDByID")
			}
			foundItem = getItem.GuestSSID
			return foundItem, err
		}
	}
	return foundItem, err
}
