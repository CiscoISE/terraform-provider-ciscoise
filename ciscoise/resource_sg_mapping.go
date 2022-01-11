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

func resourceSgMapping() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on IPToSGTMapping.

- This resource allows the client to update an IP to SGT mapping by ID.

- This resource deletes an IP to SGT mapping.

- This resource creates an IP to SGT mapping.
`,

		CreateContext: resourceSgMappingCreate,
		ReadContext:   resourceSgMappingRead,
		UpdateContext: resourceSgMappingUpdate,
		DeleteContext: resourceSgMappingDelete,
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

						"deploy_to": &schema.Schema{
							Description: `Mandatory unless mappingGroup is set or unless deployType=ALL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"deploy_type": &schema.Schema{
							Description: `Allowed values:
- ALL,
- ND,
- NDG`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_ip": &schema.Schema{
							Description: `Mandatory if hostName is empty -- valid IP`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"host_name": &schema.Schema{
							Description: `Mandatory if hostIp is empty`,
							Type:        schema.TypeString,
							Computed:    true,
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
						"mapping_group": &schema.Schema{
							Description: `Mapping Group Id. Mandatory unless sgt and deployTo and deployType are set`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sgt": &schema.Schema{
							Description: `Mandatory unless mappingGroup is set`,
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

						"deploy_to": &schema.Schema{
							Description: `Mandatory unless mappingGroup is set or unless deployType=ALL`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"deploy_type": &schema.Schema{
							Description: `Allowed values:
- ALL,
- ND,
- NDG`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"host_ip": &schema.Schema{
							Description: `Mandatory if hostName is empty -- valid IP`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"host_name": &schema.Schema{
							Description: `Mandatory if hostIp is empty`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mapping_group": &schema.Schema{
							Description: `Mapping Group Id. Mandatory unless sgt and deployTo and deployType are set`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sgt": &schema.Schema{
							Description:      `Mandatory unless mappingGroup is set`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSuppressSgt(),
						},
					},
				},
			},
		},
	}
}

func resourceSgMappingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMapping Create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSgMappingCreateIPToSgtMapping(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.IPToSgtMapping.GetIPToSgtMappingByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceSgMappingRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetIPToSgtMappingQueryParams{}

		response2, _, err := client.IPToSgtMapping.GetIPToSgtMapping(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsIPToSgtMappingGetIPToSgtMapping(m, response2, &queryParams2)
			item2, err := searchIPToSgtMappingGetIPToSgtMapping(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceSgMappingRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.IPToSgtMapping.CreateIPToSgtMapping(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateIPToSgtMapping", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateIPToSgtMapping", err))
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
	return resourceSgMappingRead(ctx, d, m)
}

func resourceSgMappingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMapping Read for id=[%s]", d.Id())
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
		log.Printf("[DEBUG] Selected method: GetIPToSgtMapping")
		queryParams1 := isegosdk.GetIPToSgtMappingQueryParams{}

		response1, restyResp1, err := client.IPToSgtMapping.GetIPToSgtMapping(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsIPToSgtMappingGetIPToSgtMapping(m, response1, &queryParams1)
		item1, err := searchIPToSgtMappingGetIPToSgtMapping(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenIPToSgtMappingGetIPToSgtMappingByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPToSgtMapping search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetIPToSgtMappingByID")
		vvID := vID

		response2, restyResp2, err := client.IPToSgtMapping.GetIPToSgtMappingByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenIPToSgtMappingGetIPToSgtMappingByIDItem(response2.SgMapping)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPToSgtMappingByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSgMappingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMapping Update for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetIPToSgtMappingQueryParams{}
		getResp1, _, err := client.IPToSgtMapping.GetIPToSgtMapping(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsIPToSgtMappingGetIPToSgtMapping(m, getResp1, &queryParams1)
			item1, err := searchIPToSgtMappingGetIPToSgtMapping(m, items1, vName, vID)
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
		request1 := expandRequestSgMappingUpdateIPToSgtMappingByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.IPToSgtMapping.UpdateIPToSgtMappingByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateIPToSgtMappingByID", err, restyResp1.String(),
					"Failure at UpdateIPToSgtMappingByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateIPToSgtMappingByID", err,
				"Failure at UpdateIPToSgtMappingByID, unexpected response", ""))
			return diags
		}
	}

	return resourceSgMappingRead(ctx, d, m)
}

func resourceSgMappingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMapping Delete for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetIPToSgtMappingQueryParams{}

		getResp1, _, err := client.IPToSgtMapping.GetIPToSgtMapping(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsIPToSgtMappingGetIPToSgtMapping(m, getResp1, &queryParams1)
		item1, err := searchIPToSgtMappingGetIPToSgtMapping(m, items1, vName, vID)
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
		getResp, _, err := client.IPToSgtMapping.GetIPToSgtMappingByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.IPToSgtMapping.DeleteIPToSgtMappingByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteIPToSgtMappingByID", err, restyResp1.String(),
				"Failure at DeleteIPToSgtMappingByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteIPToSgtMappingByID", err,
			"Failure at DeleteIPToSgtMappingByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSgMappingCreateIPToSgtMapping(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingCreateIPToSgtMapping {
	request := isegosdk.RequestIPToSgtMappingCreateIPToSgtMapping{}
	request.SgMapping = expandRequestSgMappingCreateIPToSgtMappingSgMapping(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgMappingCreateIPToSgtMappingSgMapping(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingCreateIPToSgtMappingSgMapping {
	request := isegosdk.RequestIPToSgtMappingCreateIPToSgtMappingSgMapping{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt")))) {
		first, _ := replaceRegExStrings(interfaceToString(v), "", `\s*\(.*\)$`, "")
		request.Sgt = first
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deploy_to")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deploy_to")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deploy_to")))) {
		request.DeployTo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deploy_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deploy_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deploy_type")))) {
		request.DeployType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_ip")))) {
		request.HostIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mapping_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mapping_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mapping_group")))) {
		request.MappingGroup = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgMappingUpdateIPToSgtMappingByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingUpdateIPToSgtMappingByID {
	request := isegosdk.RequestIPToSgtMappingUpdateIPToSgtMappingByID{}
	request.SgMapping = expandRequestSgMappingUpdateIPToSgtMappingByIDSgMapping(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgMappingUpdateIPToSgtMappingByIDSgMapping(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingUpdateIPToSgtMappingByIDSgMapping {
	request := isegosdk.RequestIPToSgtMappingUpdateIPToSgtMappingByIDSgMapping{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt")))) {
		first, _ := replaceRegExStrings(interfaceToString(v), "", `\s*\(.*\)$`, "")
		request.Sgt = first
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deploy_to")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deploy_to")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deploy_to")))) {
		request.DeployTo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deploy_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deploy_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deploy_type")))) {
		request.DeployType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_ip")))) {
		request.HostIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mapping_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mapping_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mapping_group")))) {
		request.MappingGroup = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsIPToSgtMappingGetIPToSgtMapping(m interface{}, response *isegosdk.ResponseIPToSgtMappingGetIPToSgtMapping, queryParams *isegosdk.GetIPToSgtMappingQueryParams) []isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingSearchResultResources
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
			response, _, err = client.IPToSgtMapping.GetIPToSgtMapping(queryParams)
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

func searchIPToSgtMappingGetIPToSgtMapping(m interface{}, items []isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingSearchResultResources, name string, id string) (*isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingByIDSgMapping, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingByIDSgMapping
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingByID
			getItem, _, err = client.IPToSgtMapping.GetIPToSgtMappingByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetIPToSgtMappingByID")
			}
			foundItem = getItem.SgMapping
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingByID
			getItem, _, err = client.IPToSgtMapping.GetIPToSgtMappingByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetIPToSgtMappingByID")
			}
			foundItem = getItem.SgMapping
			return foundItem, err
		}
	}
	return foundItem, err
}
