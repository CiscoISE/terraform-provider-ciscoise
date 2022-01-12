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

func resourcePortalTheme() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on PortalTheme.

- This resource allows the client to update a portal theme by ID.

- This resource deletes a portal theme by ID.

- This resource creates a portal theme.
`,

		CreateContext: resourcePortalThemeCreate,
		ReadContext:   resourcePortalThemeRead,
		UpdateContext: resourcePortalThemeUpdate,
		DeleteContext: resourcePortalThemeDelete,
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
						"theme_data": &schema.Schema{
							Description: `Portal Theme for all portals`,
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

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"theme_data": &schema.Schema{
							Description: `Portal Theme for all portals`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourcePortalThemeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PortalTheme create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestPortalThemeCreatePortalTheme(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.PortalTheme.GetPortalThemeByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourcePortalThemeRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetPortalThemesQueryParams{}

		response2, _, err := client.PortalTheme.GetPortalThemes(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsPortalThemeGetPortalThemes(m, response2, &queryParams2)
			item2, err := searchPortalThemeGetPortalThemes(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourcePortalThemeRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.PortalTheme.CreatePortalTheme(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreatePortalTheme", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreatePortalTheme", err))
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
	return resourcePortalThemeRead(ctx, d, m)
}

func resourcePortalThemeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PortalTheme read for id=[%s]", d.Id())
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
		log.Printf("[DEBUG] Selected method: GetPortalThemes")
		queryParams1 := isegosdk.GetPortalThemesQueryParams{}

		response1, restyResp1, err := client.PortalTheme.GetPortalThemes(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsPortalThemeGetPortalThemes(m, response1, &queryParams1)
		item1, err := searchPortalThemeGetPortalThemes(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenPortalThemeGetPortalThemeByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortalThemes search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPortalThemeByID")
		vvID := vID

		response2, restyResp2, err := client.PortalTheme.GetPortalThemeByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenPortalThemeGetPortalThemeByIDItem(response2.PortalTheme)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortalThemeByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourcePortalThemeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PortalTheme update for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetPortalThemesQueryParams{}

		getResp1, _, err := client.PortalTheme.GetPortalThemes(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsPortalThemeGetPortalThemes(m, getResp1, &queryParams1)
			item1, err := searchPortalThemeGetPortalThemes(m, items1, vName, vID)
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
		request1 := expandRequestPortalThemeUpdatePortalThemeByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.PortalTheme.UpdatePortalThemeByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatePortalThemeByID", err, restyResp1.String(),
					"Failure at UpdatePortalThemeByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatePortalThemeByID", err,
				"Failure at UpdatePortalThemeByID, unexpected response", ""))
			return diags
		}
	}

	return resourcePortalThemeRead(ctx, d, m)
}

func resourcePortalThemeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PortalTheme delete for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetPortalThemesQueryParams{}

		getResp1, _, err := client.PortalTheme.GetPortalThemes(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsPortalThemeGetPortalThemes(m, getResp1, &queryParams1)
		item1, err := searchPortalThemeGetPortalThemes(m, items1, vName, vID)
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
		getResp, _, err := client.PortalTheme.GetPortalThemeByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.PortalTheme.DeletePortalThemeByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletePortalThemeByID", err, restyResp1.String(),
				"Failure at DeletePortalThemeByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletePortalThemeByID", err,
			"Failure at DeletePortalThemeByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestPortalThemeCreatePortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPortalThemeCreatePortalTheme {
	request := isegosdk.RequestPortalThemeCreatePortalTheme{}
	request.PortalTheme = expandRequestPortalThemeCreatePortalThemePortalTheme(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPortalThemeCreatePortalThemePortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPortalThemeCreatePortalThemePortalTheme {
	request := isegosdk.RequestPortalThemeCreatePortalThemePortalTheme{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".theme_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".theme_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".theme_data")))) {
		request.ThemeData = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPortalThemeUpdatePortalThemeByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPortalThemeUpdatePortalThemeByID {
	request := isegosdk.RequestPortalThemeUpdatePortalThemeByID{}
	request.PortalTheme = expandRequestPortalThemeUpdatePortalThemeByIDPortalTheme(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPortalThemeUpdatePortalThemeByIDPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPortalThemeUpdatePortalThemeByIDPortalTheme {
	request := isegosdk.RequestPortalThemeUpdatePortalThemeByIDPortalTheme{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".theme_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".theme_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".theme_data")))) {
		request.ThemeData = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsPortalThemeGetPortalThemes(m interface{}, response *isegosdk.ResponsePortalThemeGetPortalThemes, queryParams *isegosdk.GetPortalThemesQueryParams) []isegosdk.ResponsePortalThemeGetPortalThemesSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponsePortalThemeGetPortalThemesSearchResultResources
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
			response, _, err = client.PortalTheme.GetPortalThemes(queryParams)
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

func searchPortalThemeGetPortalThemes(m interface{}, items []isegosdk.ResponsePortalThemeGetPortalThemesSearchResultResources, name string, id string) (*isegosdk.ResponsePortalThemeGetPortalThemeByIDPortalTheme, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponsePortalThemeGetPortalThemeByIDPortalTheme
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponsePortalThemeGetPortalThemeByID
			getItem, _, err = client.PortalTheme.GetPortalThemeByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetPortalThemeByID")
			}
			foundItem = getItem.PortalTheme
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponsePortalThemeGetPortalThemeByID
			getItem, _, err = client.PortalTheme.GetPortalThemeByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetPortalThemeByID")
			}
			foundItem = getItem.PortalTheme
			return foundItem, err
		}
	}
	return foundItem, err
}
