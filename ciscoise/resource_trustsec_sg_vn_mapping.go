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

func resourceTrustsecSgVnMapping() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on sgVnMapping.

- Create Security Group and Virtual Network mapping

- Update Security Group and Virtual Network mapping

- Delete Security Group and Virtual Network mapping
`,

		CreateContext: resourceTrustsecSgVnMappingCreate,
		ReadContext:   resourceTrustsecSgVnMappingRead,
		UpdateContext: resourceTrustsecSgVnMappingUpdate,
		DeleteContext: resourceTrustsecSgVnMappingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description:      `Identifier of the SG-VN mapping`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"last_update": &schema.Schema{
							Description:      `Timestamp for the last update of the SG-VN mapping`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"sg_name": &schema.Schema{
							Description:      `Name of the associated Security Group to be used for identity if id is not provided`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"sgt_id": &schema.Schema{
							Description:      `Identifier of the associated Security Group which is required unless its name is provided`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"vn_id": &schema.Schema{
							Description:      `Identifier for the associated Virtual Network which is required unless its name is provided`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"vn_name": &schema.Schema{
							Description:      `Name of the associated Virtual Network to be used for identity if id is not provided`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Identifier of the SG-VN mapping`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_update": &schema.Schema{
							Description: `Timestamp for the last update of the SG-VN mapping`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sg_name": &schema.Schema{
							Description: `Name of the associated Security Group to be used for identity if id is not provided`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sgt_id": &schema.Schema{
							Description: `Identifier of the associated Security Group which is required unless its name is provided`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"vn_id": &schema.Schema{
							Description: `Identifier for the associated Virtual Network which is required unless its name is provided`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"vn_name": &schema.Schema{
							Description: `Name of the associated Virtual Network to be used for identity if id is not provided`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceTrustsecSgVnMappingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecSgVnMapping create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	isEnableAutoImport := clientConfig.EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTrustsecSgVnMappingCreateSgVnMapping(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vSgName, _ := resourceItem["sg_name"]
	vvSgName := interfaceToString(vSgName)
	vSgtID, _ := resourceItem["sgt_id"]
	vvSgtID := interfaceToString(vSgtID)
	vVnID, _ := resourceItem["vn_id"]
	vvVnID := interfaceToString(vVnID)
	vVnName, _ := resourceItem["vn_name"]
	vvVnName := interfaceToString(vVnName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse2, _, err := client.SgVnMapping.GetSgVnMappingByID(vvID)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["sg_name"] = vvSgName
				resourceMap["sgt_id"] = vvSgtID
				resourceMap["vn_id"] = vvVnID
				resourceMap["vn_name"] = vvVnName
				d.SetId(joinResourceID(resourceMap))
				return resourceTrustsecSgVnMappingRead(ctx, d, m)
			}
		} else {
			queryParams2 := isegosdk.GetSgVnMappingsQueryParams{}

			response2, _, err := client.SgVnMapping.GetSgVnMappings(&queryParams2)
			if response2 != nil && err == nil {
				items2 := getAllItemsSgVnMappingGetSgVnMappings(m, response2, &queryParams2)
				item2, err := searchSgVnMappingGetSgVnMappings(m, items2, vvSgName, vvSgtID, vvVnID, vvVnName, vvID)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["id"] = vvID
					resourceMap["sg_name"] = vvSgName
					resourceMap["sgt_id"] = vvSgtID
					resourceMap["vn_id"] = vvVnID
					resourceMap["vn_name"] = vvVnName
					d.SetId(joinResourceID(resourceMap))
					return resourceTrustsecSgVnMappingRead(ctx, d, m)
				}
			}
		}
	}
	resp1, restyResp1, err := client.SgVnMapping.CreateSgVnMapping(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSgVnMapping", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSgVnMapping", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["sg_name"] = vvSgName
	resourceMap["sgt_id"] = vvSgtID
	resourceMap["vn_id"] = vvVnID
	resourceMap["vn_name"] = vvVnName
	d.SetId(joinResourceID(resourceMap))
	return resourceTrustsecSgVnMappingRead(ctx, d, m)
}

func resourceTrustsecSgVnMappingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecSgVnMapping read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vSgName, _ := resourceMap["sg_name"]
	vSgtID, _ := resourceMap["sgt_id"]
	vVnID, _ := resourceMap["vn_id"]
	vVnName, _ := resourceMap["vn_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		vvID := vID
		vvSgName := vSgName
		vvSgtID := vSgtID
		vvVnID := vVnID
		vvVnName := vVnName
		log.Printf("[DEBUG] Selected method: GetSgVnMappings")
		queryParams1 := isegosdk.GetSgVnMappingsQueryParams{}

		response1, restyResp1, err := client.SgVnMapping.GetSgVnMappings(&queryParams1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsSgVnMappingGetSgVnMappings(m, response1, nil)
		item1, err := searchSgVnMappingGetSgVnMappings(m, items1, vvSgName, vvSgtID, vvVnID, vvVnName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenSgVnMappingGetSgVnMappingByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSgVnMappings search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSgVnMappings search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSgVnMappingByID")
		vvID := vID

		response2, restyResp2, err := client.SgVnMapping.GetSgVnMappingByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSgVnMappingGetSgVnMappingByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSgVnMappingByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSgVnMappingByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTrustsecSgVnMappingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecSgVnMapping update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vSgName, _ := resourceMap["sg_name"]
	vSgtID, _ := resourceMap["sgt_id"]
	vVnID, _ := resourceMap["vn_id"]
	vVnName, _ := resourceMap["vn_name"]
	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 1 {
		vvID = vID
		vvSgName := vSgName
		vvSgtID := vSgtID
		vvVnID := vVnID
		vvVnName := vVnName

		log.Printf("[DEBUG] Selected method: GetSgVnMappings")
		queryParams2 := isegosdk.GetSgVnMappingsQueryParams{}
		response1, _, err := client.SgVnMapping.GetSgVnMappings(&queryParams2)

		if err == nil && response1 != nil {
			items1 := getAllItemsSgVnMappingGetSgVnMappings(m, response1, nil)
			item1, err := searchSgVnMappingGetSgVnMappings(m, items1, vvSgName, vvSgtID, vvVnID, vvVnName, vvID)
			if err == nil && item1 != nil {
				if len(*item1) > 0 {
					vvID = (*item1)[0].ID
				}
			}

		}
	}
	if selectedMethod == 2 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestTrustsecSgVnMappingUpdateSgVnMappingByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.SgVnMapping.UpdateSgVnMappingByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSgVnMappingByID", err, restyResp1.String(),
					"Failure at UpdateSgVnMappingByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSgVnMappingByID", err,
				"Failure at UpdateSgVnMappingByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceTrustsecSgVnMappingRead(ctx, d, m)
}

func resourceTrustsecSgVnMappingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecSgVnMapping delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vSgName, _ := resourceMap["sg_name"]
	vSgtID, _ := resourceMap["sgt_id"]
	vVnID, _ := resourceMap["vn_id"]
	vVnName, _ := resourceMap["vn_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {
		vvID = vID
		vvSgName := vSgName
		vvSgtID := vSgtID
		vvVnID := vVnID
		vvVnName := vVnName

		queryParams1 := isegosdk.GetSgVnMappingsQueryParams{}

		getResp1, _, err := client.SgVnMapping.GetSgVnMappings(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSgVnMappingGetSgVnMappings(m, getResp1, &queryParams1)
		item1, err := searchSgVnMappingGetSgVnMappings(m, items1, vvSgName, vvSgtID, vvVnID, vvVnName, vvID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if item1 != nil && len(*item1) > 0 {
			if vvID != (*item1)[0].ID {
				vvID = (*item1)[0].ID
			} else {
				vvID = vID
			}
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 2 {
		vvID = vID
		getResp, _, err := client.SgVnMapping.GetSgVnMappingByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.SgVnMapping.DeleteSgVnMappingByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSgVnMappingByID", err, restyResp1.String(),
				"Failure at DeleteSgVnMappingByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSgVnMappingByID", err,
			"Failure at DeleteSgVnMappingByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTrustsecSgVnMappingCreateSgVnMapping(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSgVnMappingCreateSgVnMapping {
	request := isegosdk.RequestSgVnMappingCreateSgVnMapping{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sg_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sg_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sg_name")))) {
		request.SgName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt_id")))) {
		request.SgtID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_id")))) {
		request.VnID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_name")))) {
		request.VnName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTrustsecSgVnMappingUpdateSgVnMappingByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSgVnMappingUpdateSgVnMappingByID {
	request := isegosdk.RequestSgVnMappingUpdateSgVnMappingByID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sg_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sg_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sg_name")))) {
		request.SgName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt_id")))) {
		request.SgtID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_id")))) {
		request.VnID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_name")))) {
		request.VnName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func getAllItemsSgVnMappingGetSgVnMappings(m interface{}, response *isegosdk.ResponseSgVnMappingGetSgVnMappings, queryParams *isegosdk.GetSgVnMappingsQueryParams) []isegosdk.ResponseSgVnMappingGetSgVnMappingsResponse {
	var respItems []isegosdk.ResponseSgVnMappingGetSgVnMappingsResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchSgVnMappingGetSgVnMappings(m interface{}, items []isegosdk.ResponseSgVnMappingGetSgVnMappingsResponse, sgName string, sgtID string, vnID string, vnName string, id string) (*[]isegosdk.ResponseSgVnMappingGetSgVnMappingByIDResponse, error) {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var err error
	var foundItem *[]isegosdk.ResponseSgVnMappingGetSgVnMappingByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSgVnMappingGetSgVnMappingByID
			getItem, _, err = client.SgVnMapping.GetSgVnMappingByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSgVnMappingByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSgVnMappingGetSgVnMappingByID
			getItem, _, err = client.SgVnMapping.GetSgVnMappingByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSgVnMappingByID")
			}
			hasSgName := (sgName != "" && sgName == item.SgName)
			hasSgtID := (sgtID != "" && sgtID == item.SgtID)
			hasVnID := (vnID != "" && vnID == item.VnID)
			hasVnName := (vnName != "" && vnName == item.VnName)
			isEqual := (hasSgName || hasSgtID) && (hasVnID || hasVnName)
			if isEqual {
				foundItem = getItem.Response
				return foundItem, err
			}
		}
	}
	return foundItem, err
}
