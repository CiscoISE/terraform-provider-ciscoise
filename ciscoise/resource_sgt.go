package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSgt() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSgtCreate,
		ReadContext:   resourceSgtRead,
		UpdateContext: resourceSgtUpdate,
		DeleteContext: resourceSgtDelete,
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

						"default_sgacls": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"generation_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_read_only": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
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
							Optional: true,
							Computed: true,
						},
						"propogate_to_apic": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSgtCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestSgtCreateSecurityGroup(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.SecurityGroups.GetSecurityGroupByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		queryParams2 := isegosdk.GetSecurityGroupsQueryParams{}

		response2, _, err := client.SecurityGroups.GetSecurityGroups(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsSecurityGroupsGetSecurityGroups(m, response2, &queryParams2)
			item2, err := searchSecurityGroupsGetSecurityGroups(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	restyResp1, err := client.SecurityGroups.CreateSecurityGroup(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSecurityGroup", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSecurityGroup", err))
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
	return diags
}

func resourceSgtRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetSecurityGroups")
		queryParams1 := isegosdk.GetSecurityGroupsQueryParams{}

		response1, _, err := client.SecurityGroups.GetSecurityGroups(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSecurityGroups", err,
				"Failure at GetSecurityGroups, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsSecurityGroupsGetSecurityGroups(m, response1, &queryParams1)
		item1, err := searchSecurityGroupsGetSecurityGroups(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetSecurityGroups response", err,
				"Failure when searching item from GetSecurityGroups, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityGroups search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSecurityGroupByID")
		vvID := vID

		response2, _, err := client.SecurityGroups.GetSecurityGroupByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSecurityGroupByID", err,
				"Failure at GetSecurityGroupByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenSecurityGroupsGetSecurityGroupByIDItem(response2.Sgt)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityGroupByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSgtUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetSecurityGroupsQueryParams{}

		getResp1, _, err := client.SecurityGroups.GetSecurityGroups(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsSecurityGroupsGetSecurityGroups(m, getResp1, &queryParams1)
			item1, err := searchSecurityGroupsGetSecurityGroups(m, items1, vName, vID)
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
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestSgtUpdateSecurityGroupByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SecurityGroups.UpdateSecurityGroupByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSecurityGroupByID", err, restyResp1.String(),
					"Failure at UpdateSecurityGroupByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSecurityGroupByID", err,
				"Failure at UpdateSecurityGroupByID, unexpected response", ""))
			return diags
		}
	}

	return resourceSgtRead(ctx, d, m)
}

func resourceSgtDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		queryParams1 := isegosdk.GetSecurityGroupsQueryParams{}

		getResp1, _, err := client.SecurityGroups.GetSecurityGroups(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSecurityGroupsGetSecurityGroups(m, getResp1, &queryParams1)
		item1, err := searchSecurityGroupsGetSecurityGroups(m, items1, vName, vID)
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
		getResp, _, err := client.SecurityGroups.GetSecurityGroupByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.SecurityGroups.DeleteSecurityGroupByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSecurityGroupByID", err, restyResp1.String(),
				"Failure at DeleteSecurityGroupByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSecurityGroupByID", err,
			"Failure at DeleteSecurityGroupByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSgtCreateSecurityGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsCreateSecurityGroup {
	request := isegosdk.RequestSecurityGroupsCreateSecurityGroup{}
	request.Sgt = expandRequestSgtCreateSecurityGroupSgt(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgtCreateSecurityGroupSgt(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsCreateSecurityGroupSgt {
	request := isegosdk.RequestSecurityGroupsCreateSecurityGroupSgt{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".generation_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".generation_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".generation_id"))) {
		request.GenerationID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".is_read_only"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_read_only"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_read_only"))) {
		request.IsReadOnly = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".propogate_to_apic"); !isEmptyValue(reflect.ValueOf(d.Get(key+".propogate_to_apic"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".propogate_to_apic"))) {
		request.PropogateToAPIc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".default_sgacls"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_sgacls"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_sgacls"))) {
		defaultSgACLs := v.([]interface{})
		request.DefaultSgACLs = &defaultSgACLs
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgtUpdateSecurityGroupByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsUpdateSecurityGroupByID {
	request := isegosdk.RequestSecurityGroupsUpdateSecurityGroupByID{}
	request.Sgt = expandRequestSgtUpdateSecurityGroupByIDSgt(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgtUpdateSecurityGroupByIDSgt(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsUpdateSecurityGroupByIDSgt {
	request := isegosdk.RequestSecurityGroupsUpdateSecurityGroupByIDSgt{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".generation_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".generation_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".generation_id"))) {
		request.GenerationID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".is_read_only"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_read_only"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_read_only"))) {
		request.IsReadOnly = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".propogate_to_apic"); !isEmptyValue(reflect.ValueOf(d.Get(key+".propogate_to_apic"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".propogate_to_apic"))) {
		request.PropogateToAPIc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".default_sgacls"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_sgacls"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_sgacls"))) {
		defaultSgACLs := v.([]interface{})
		request.DefaultSgACLs = &defaultSgACLs
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsSecurityGroupsGetSecurityGroups(m interface{}, response *isegosdk.ResponseSecurityGroupsGetSecurityGroups, queryParams *isegosdk.GetSecurityGroupsQueryParams) []isegosdk.ResponseSecurityGroupsGetSecurityGroupsSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseSecurityGroupsGetSecurityGroupsSearchResultResources
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
			response, _, err = client.SecurityGroups.GetSecurityGroups(queryParams)
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

func searchSecurityGroupsGetSecurityGroups(m interface{}, items []isegosdk.ResponseSecurityGroupsGetSecurityGroupsSearchResultResources, name string, id string) (*isegosdk.ResponseSecurityGroupsGetSecurityGroupByIDSgt, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSecurityGroupsGetSecurityGroupByIDSgt
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSecurityGroupsGetSecurityGroupByID
			getItem, _, err = client.SecurityGroups.GetSecurityGroupByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSecurityGroupByID")
			}
			foundItem = getItem.Sgt
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSecurityGroupsGetSecurityGroupByID
			getItem, _, err = client.SecurityGroups.GetSecurityGroupByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSecurityGroupByID")
			}
			foundItem = getItem.Sgt
			return foundItem, err
		}
	}
	return foundItem, err
}
