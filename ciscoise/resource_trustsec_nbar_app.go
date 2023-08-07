package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrustsecNbarApp() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on nbarApp.

- Create NBAR application

- Update NBAR Application

- Delete NBAR Application
`,

		CreateContext: resourceTrustsecNbarAppCreate,
		ReadContext:   resourceTrustsecNbarAppRead,
		UpdateContext: resourceTrustsecNbarAppUpdate,
		DeleteContext: resourceTrustsecNbarAppDelete,
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

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_identities": &schema.Schema{
							Description: `Array of NIs`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ports": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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

						"id": &schema.Schema{
							Description:      `id path parameter.`,
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: diffSupressOptional(),
						},
						"ports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
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

func resourceTrustsecNbarAppCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecNbarApp create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	isEnableAutoImport := clientConfig.EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTrustsecNbarAppCreateNbarApp(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse2, _, err := client.NbarApp.GetNbarAppByID(vvID)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceTrustsecNbarAppRead(ctx, d, m)
			}
		}
		if okName && vvName != "" {
			queryParams2 := isegosdk.GetNbarAppsQueryParams{}

			response2, _, err := client.NbarApp.GetNbarApps(&queryParams2)
			if response2 != nil && err == nil {
				items2 := getAllItemsNbarAppGetNbarApps(m, response2, &queryParams2)
				item2, err := searchNbarAppGetNbarApps(m, items2, vvName, vvID)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["id"] = vvID
					resourceMap["name"] = vvName
					d.SetId(joinResourceID(resourceMap))
					return resourceTrustsecNbarAppRead(ctx, d, m)
				}
			}
		}
	}
	resp1, restyResp1, err := client.NbarApp.CreateNbarApp(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNbarApp", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNbarApp", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceTrustsecNbarAppRead(ctx, d, m)
}

func resourceTrustsecNbarAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecNbarApp read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
		log.Printf("[DEBUG] Selected method: GetNbarApps")
		queryParams1 := isegosdk.GetNbarAppsQueryParams{}

		response1, restyResp1, err := client.NbarApp.GetNbarApps(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsNbarAppGetNbarApps(m, response1, nil)
		item1, err := searchNbarAppGetNbarApps(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenNbarAppGetNbarAppByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNbarApps search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNbarApps search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNbarAppByID")
		vvID := vID

		response2, restyResp2, err := client.NbarApp.GetNbarAppByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNbarAppGetNbarAppByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNbarAppByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNbarAppByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTrustsecNbarAppUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecNbarApp update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
	}
	if selectedMethod == 2 {
		vvName = vName

		queryParams1 := isegosdk.GetNbarAppsQueryParams{}
		getResp, _, err := client.NbarApp.GetNbarApps(&queryParams1)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNbarApps", err,
				"Failure at GetNbarApps, unexpected response", ""))
			return diags
		}
		items1 := getAllItemsNbarAppGetNbarApps(m, getResp, nil)
		item1, err := searchNbarAppGetNbarApps(m, items1, vvName, vvID)
		//Set value vvID = getResp.
		if item1 != nil && len(*item1) > 0 {
			vvID = (*item1)[0].ID
		}
	}
	// NOTE: Added getAllItems and search function to get missing params
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestTrustsecNbarAppUpdateNbarAppByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.NbarApp.UpdateNbarAppByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNbarAppByID", err, restyResp1.String(),
					"Failure at UpdateNbarAppByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNbarAppByID", err,
				"Failure at UpdateNbarAppByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceTrustsecNbarAppRead(ctx, d, m)
}

func resourceTrustsecNbarAppDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecNbarApp delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
		queryParams1 := isegosdk.GetNbarAppsQueryParams{}

		getResp1, _, err := client.NbarApp.GetNbarApps(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNbarAppGetNbarApps(m, getResp1, &queryParams1)
		item1, err := searchNbarAppGetNbarApps(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if item1 != nil && len(*item1) > 0 {
			if vID != (*item1)[0].ID {
				vID = (*item1)[0].ID
			} else {
				vvID = vID
			}
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.NbarApp.GetNbarAppByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NbarApp.DeleteNbarAppByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNbarAppByID", err, restyResp1.String(),
				"Failure at DeleteNbarAppByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNbarAppByID", err,
			"Failure at DeleteNbarAppByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTrustsecNbarAppCreateNbarApp(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNbarAppCreateNbarApp {
	request := isegosdk.RequestNbarAppCreateNbarApp{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_identities")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_identities")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_identities")))) {
		request.NetworkIDentities = expandRequestTrustsecNbarAppCreateNbarAppNetworkIDentitiesArray(ctx, key+".network_identities", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTrustsecNbarAppCreateNbarAppNetworkIDentitiesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNbarAppCreateNbarAppNetworkIDentities {
	request := []isegosdk.RequestNbarAppCreateNbarAppNetworkIDentities{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestTrustsecNbarAppCreateNbarAppNetworkIDentities(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTrustsecNbarAppCreateNbarAppNetworkIDentities(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNbarAppCreateNbarAppNetworkIDentities {
	request := isegosdk.RequestNbarAppCreateNbarAppNetworkIDentities{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ports")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ports")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ports")))) {
		request.Ports = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTrustsecNbarAppUpdateNbarAppByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNbarAppUpdateNbarAppByID {
	request := isegosdk.RequestNbarAppUpdateNbarAppByID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_identities")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_identities")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_identities")))) {
		request.NetworkIDentities = expandRequestTrustsecNbarAppUpdateNbarAppByIDNetworkIDentitiesArray(ctx, key+".network_identities", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTrustsecNbarAppUpdateNbarAppByIDNetworkIDentitiesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNbarAppUpdateNbarAppByIDNetworkIDentities {
	request := []isegosdk.RequestNbarAppUpdateNbarAppByIDNetworkIDentities{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestTrustsecNbarAppUpdateNbarAppByIDNetworkIDentities(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTrustsecNbarAppUpdateNbarAppByIDNetworkIDentities(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNbarAppUpdateNbarAppByIDNetworkIDentities {
	request := isegosdk.RequestNbarAppUpdateNbarAppByIDNetworkIDentities{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ports")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ports")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ports")))) {
		request.Ports = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNbarAppGetNbarApps(m interface{}, response *isegosdk.ResponseNbarAppGetNbarApps, queryParams *isegosdk.GetNbarAppsQueryParams) []isegosdk.ResponseNbarAppGetNbarAppsResponse {
	var respItems []isegosdk.ResponseNbarAppGetNbarAppsResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNbarAppGetNbarApps(m interface{}, items []isegosdk.ResponseNbarAppGetNbarAppsResponse, name string, id string) (*[]isegosdk.ResponseNbarAppGetNbarAppByIDResponse, error) {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var err error
	var foundItem *[]isegosdk.ResponseNbarAppGetNbarAppByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNbarAppGetNbarAppByID
			getItem, _, err = client.NbarApp.GetNbarAppByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNbarAppByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNbarAppGetNbarAppByID
			getItem, _, err = client.NbarApp.GetNbarAppByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNbarAppByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
