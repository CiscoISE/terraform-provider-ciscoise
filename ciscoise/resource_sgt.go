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

func resourceSgt() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SecurityGroups.

- This resource allows the client to update a security group.

- This resource deletes a security group.

- This resource creates a security group.
`,

		CreateContext: resourceSgtCreate,
		ReadContext:   resourceSgtRead,
		UpdateContext: resourceSgtUpdate,
		DeleteContext: resourceSgtDelete,
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

						"default_sgacls": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"generation_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_read_only": &schema.Schema{
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
						"propogate_to_apic": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": &schema.Schema{
							Description: `Value range: 2 ot 65519 or -1 to auto-generate`,
							Type:        schema.TypeInt,
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

						"default_sgacls": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"generation_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"is_read_only": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"propogate_to_apic": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"value": &schema.Schema{
							Description: `Value range: 2 ot 65519 or -1 to auto-generate`,
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSgtCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Sgt create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSgtCreateSecurityGroup(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

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
			return resourceSgtRead(ctx, d, m)
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
				return resourceSgtRead(ctx, d, m)
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
	return resourceSgtRead(ctx, d, m)
}

func resourceSgtRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Sgt read for id=[%s]", d.Id())
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

		response1, restyResp1, err := client.SecurityGroups.GetSecurityGroups(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsSecurityGroupsGetSecurityGroups(m, response1, &queryParams1)
		item1, err := searchSecurityGroupsGetSecurityGroups(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenSecurityGroupsGetSecurityGroupByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityGroups search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSecurityGroupByID")
		vvID := vID

		response2, restyResp2, err := client.SecurityGroups.GetSecurityGroupByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

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
	log.Printf("[DEBUG] Beginning Sgt update for id=[%s]", d.Id())
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
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestSgtUpdateSecurityGroupByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.SecurityGroups.UpdateSecurityGroupByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
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
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceSgtRead(ctx, d, m)
}

func resourceSgtDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Sgt delete for id=[%s]", d.Id())
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
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".generation_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".generation_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".generation_id")))) {
		request.GenerationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_read_only")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_read_only")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_read_only")))) {
		request.IsReadOnly = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".propogate_to_apic")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".propogate_to_apic")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".propogate_to_apic")))) {
		request.PropogateToAPIc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_sgacls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_sgacls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_sgacls")))) {
		request.DefaultSgACLs = expandRequestSgtCreateSecurityGroupSgtDefaultSgACLsArray(ctx, key+".default_sgacls", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgtCreateSecurityGroupSgtDefaultSgACLsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSecurityGroupsCreateSecurityGroupSgtDefaultSgACLs {
	request := []isegosdk.RequestSecurityGroupsCreateSecurityGroupSgtDefaultSgACLs{}
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
		i := expandRequestSgtCreateSecurityGroupSgtDefaultSgACLs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgtCreateSecurityGroupSgtDefaultSgACLs(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsCreateSecurityGroupSgtDefaultSgACLs {
	var request isegosdk.RequestSecurityGroupsCreateSecurityGroupSgtDefaultSgACLs
	keyValue := d.Get(fixKeyAccess(key))
	request = requestStringToInterface(interfaceToString(keyValue))
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".generation_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".generation_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".generation_id")))) {
		request.GenerationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_read_only")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_read_only")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_read_only")))) {
		request.IsReadOnly = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".propogate_to_apic")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".propogate_to_apic")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".propogate_to_apic")))) {
		request.PropogateToAPIc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_sgacls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_sgacls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_sgacls")))) {
		request.DefaultSgACLs = expandRequestSgtUpdateSecurityGroupByIDSgtDefaultSgACLsArray(ctx, key+".default_sgacls", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgtUpdateSecurityGroupByIDSgtDefaultSgACLsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSecurityGroupsUpdateSecurityGroupByIDSgtDefaultSgACLs {
	request := []isegosdk.RequestSecurityGroupsUpdateSecurityGroupByIDSgtDefaultSgACLs{}
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
		i := expandRequestSgtUpdateSecurityGroupByIDSgtDefaultSgACLs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgtUpdateSecurityGroupByIDSgtDefaultSgACLs(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsUpdateSecurityGroupByIDSgtDefaultSgACLs {
	var request isegosdk.RequestSecurityGroupsUpdateSecurityGroupByIDSgtDefaultSgACLs
	keyValue := d.Get(fixKeyAccess(key))
	request = requestStringToInterface(interfaceToString(keyValue))
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
