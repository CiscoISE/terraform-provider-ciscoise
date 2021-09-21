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

func resourceNetworkAccessDictionaryAttribute() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Access - Dictionary Attribute.

- Create a new Dictionary Attribute for an existing Dictionary.

- Update a Dictionary Attribute

- Delete a Dictionary Attribute.
`,

		CreateContext: resourceNetworkAccessDictionaryAttributeCreate,
		ReadContext:   resourceNetworkAccessDictionaryAttributeRead,
		UpdateContext: resourceNetworkAccessDictionaryAttributeUpdate,
		DeleteContext: resourceNetworkAccessDictionaryAttributeDelete,
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

						"allowed_values": &schema.Schema{
							Description: `all of the allowed values for the dictionary attribute`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"is_default": &schema.Schema{
										Description: `true if this key value is the default between the allowed values of the dictionary attribute`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"key": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"data_type": &schema.Schema{
							Description: `the data type for the dictionary attribute`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `The description of the Dictionary attribute`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"dictionary_name": &schema.Schema{
							Description: `the name of the dictionary which the dictionary attribute belongs to`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"direction_type": &schema.Schema{
							Description: `the direction for the useage of the dictionary attribute`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier for the dictionary attribute`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"internal_name": &schema.Schema{
							Description: `the internal name of the dictionary attribute`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `The dictionary attribute's name`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkAccessDictionaryAttributeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vDictionaryName, okDictionaryName := resourceItem["dictionary_name"]
	vvDictionaryName := interfaceToString(vDictionaryName)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okName && vvName != "" && okDictionaryName && vvDictionaryName != "" {
		getResponse2, _, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributeByName(vvName, vvDictionaryName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["dictionary_name"] = vvDictionaryName
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributesByDictionaryName(vvDictionaryName)
		if response2 != nil && err == nil {
			items2 := getAllItemsNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName(m, response2, vvDictionaryName)
			item2, err := searchNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName(m, items2, vvName, vvDictionaryName)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["dictionary_name"] = vvDictionaryName
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	resp1, restyResp1, err := client.NetworkAccessDictionaryAttribute.CreateNetworkAccessDictionaryAttribute(vvDictionaryName, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkAccessDictionaryAttribute", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkAccessDictionaryAttribute", err))
		return diags
	}
	if vvName != resp1.Response.Name {
		vvName = resp1.Response.Name
	}
	resourceMap := make(map[string]string)
	resourceMap["dictionary_name"] = vvDictionaryName
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNetworkAccessDictionaryAttributeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDictionaryName, okDictionaryName := resourceMap["dictionary_name"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okDictionaryName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName, okDictionaryName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessDictionaryAttributesByDictionaryName")
		vvDictionaryName := vDictionaryName
		vvName := vName

		response1, _, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributesByDictionaryName(vvDictionaryName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionaryAttributesByDictionaryName", err,
				"Failure at GetNetworkAccessDictionaryAttributesByDictionaryName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName(m, response1, vvDictionaryName)
		item1, err := searchNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName(m, items1, vvName, vvDictionaryName)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetNetworkAccessDictionaryAttributesByDictionaryName response", err,
				"Failure when searching item from GetNetworkAccessDictionaryAttributesByDictionaryName, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaryAttributesByDictionaryName search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessDictionaryAttributeByName")
		vvName := vName
		vvDictionaryName := vDictionaryName

		response2, _, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributeByName(vvName, vvDictionaryName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionaryAttributeByName", err,
				"Failure at GetNetworkAccessDictionaryAttributeByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaryAttributeByName response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkAccessDictionaryAttributeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDictionaryName, okDictionaryName := resourceMap["dictionary_name"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okDictionaryName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName, okDictionaryName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvDictionaryName string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 1 {
		vvDictionaryName = vDictionaryName

		getResp1, _, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributesByDictionaryName(vvDictionaryName)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName(m, getResp1, vvDictionaryName)
			item1, err := searchNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName(m, items1, vName, vDictionaryName)
			if err == nil && item1 != nil {
				if vName != item1.Name {
					vvName = item1.Name
				} else {
					vvName = vName
				}
			}
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		vvDictionaryName = vDictionaryName
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvName %s", vvName)
		request1 := expandRequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkAccessDictionaryAttribute.UpdateNetworkAccessDictionaryAttributeByName(vvName, vvDictionaryName, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkAccessDictionaryAttributeByName", err, restyResp1.String(),
					"Failure at UpdateNetworkAccessDictionaryAttributeByName, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkAccessDictionaryAttributeByName", err,
				"Failure at UpdateNetworkAccessDictionaryAttributeByName, unexpected response", ""))
			return diags
		}
	}

	return resourceNetworkAccessDictionaryAttributeRead(ctx, d, m)
}

func resourceNetworkAccessDictionaryAttributeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDictionaryName, okDictionaryName := resourceMap["dictionary_name"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okDictionaryName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName, okDictionaryName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvDictionaryName string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {
		vvDictionaryName = vDictionaryName

		getResp1, _, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributesByDictionaryName(vvDictionaryName)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName(m, getResp1, vvDictionaryName)
		item1, err := searchNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName(m, items1, vName, vDictionaryName)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vName != item1.Name {
			vvName = item1.Name
		} else {
			vvName = vName
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		vvDictionaryName = vDictionaryName
		getResp, _, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributeByName(vvName, vvDictionaryName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NetworkAccessDictionaryAttribute.DeleteNetworkAccessDictionaryAttributeByName(vvName, vvDictionaryName)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkAccessDictionaryAttributeByName", err, restyResp1.String(),
				"Failure at DeleteNetworkAccessDictionaryAttributeByName, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkAccessDictionaryAttributeByName", err,
			"Failure at DeleteNetworkAccessDictionaryAttributeByName, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute {
	request := isegosdk.RequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttribute{}
	if v, ok := d.GetOkExists(key + ".allowed_values"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allowed_values"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allowed_values"))) {
		request.AllowedValues = expandRequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeAllowedValuesArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".data_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data_type"))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_name"))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".direction_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".direction_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".direction_type"))) {
		request.DirectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".internal_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".internal_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".internal_name"))) {
		request.InternalName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeAllowedValuesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeAllowedValues {
	request := []isegosdk.RequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeAllowedValues{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeAllowedValues(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeAllowedValues(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeAllowedValues {
	request := isegosdk.RequestNetworkAccessDictionaryAttributeCreateNetworkAccessDictionaryAttributeAllowedValues{}
	if v, ok := d.GetOkExists(key + ".is_default"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_default"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_default"))) {
		request.IsDefault = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key"))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName {
	request := isegosdk.RequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByName{}
	if v, ok := d.GetOkExists(key + ".allowed_values"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allowed_values"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allowed_values"))) {
		request.AllowedValues = expandRequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameAllowedValuesArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".data_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data_type"))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_name"))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".direction_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".direction_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".direction_type"))) {
		request.DirectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".internal_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".internal_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".internal_name"))) {
		request.InternalName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameAllowedValuesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameAllowedValues {
	request := []isegosdk.RequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameAllowedValues{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameAllowedValues(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameAllowedValues(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameAllowedValues {
	request := isegosdk.RequestNetworkAccessDictionaryAttributeUpdateNetworkAccessDictionaryAttributeByNameAllowedValues{}
	if v, ok := d.GetOkExists(key + ".is_default"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_default"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_default"))) {
		request.IsDefault = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key"))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName(m interface{}, response *isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName, dictionaryname string) []isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameResponse {
	var respItems []isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryName(m interface{}, items []isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameResponse, name string, id string) (*isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameResponse
	for _, item := range items {
		if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByName
			getItem, _, err = client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributeByName(name, id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessDictionaryAttributeByName")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
