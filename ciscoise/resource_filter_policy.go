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

func resourceFilterPolicy() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on FilterPolicy.

- This resource allows the client to update a filter policy.

- This resource deletes a filter policy.

- This resource creates a filter policy.
`,

		CreateContext: resourceFilterPolicyCreate,
		ReadContext:   resourceFilterPolicyRead,
		UpdateContext: resourceFilterPolicyUpdate,
		DeleteContext: resourceFilterPolicyDelete,
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

						"domains": &schema.Schema{
							Description: `List of SXP Domains, separated with comma`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sgt": &schema.Schema{
							Description: `SGT name or ID. At least one of subnet or sgt or vn should be defined`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"subnet": &schema.Schema{
							Description: `Subnet for filter policy (hostname is not supported).
At least one of subnet or sgt or vn should be defined`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"vn": &schema.Schema{
							Description: `Virtual Network.
At least one of subnet or sgt or vn should be defined`,
							Type:     schema.TypeString,
							Computed: true,
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

						"domains": &schema.Schema{
							Description:      `List of SXP Domains, separated with comma`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"id": &schema.Schema{
							Description:      `id path parameter.`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
						},
						"sgt": &schema.Schema{
							Description:      `SGT name or ID. At least one of subnet or sgt or vn should be defined`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"subnet": &schema.Schema{
							Description: `Subnet for filter policy (hostname is not supported).
		At least one of subnet or sgt or vn should be defined`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"vn": &schema.Schema{
							Description: `Virtual Network.
		At least one of subnet or sgt or vn should be defined`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func resourceFilterPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning FilterPolicy create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := clientConfig.EnableAutoImport

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestFilterPolicyCreateFilterPolicy(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vSgt, _ := resourceItem["sgt"]
	vSubnet, _ := resourceItem["subnet"]
	vVn, _ := resourceItem["vn"]
	vvID := interfaceToString(vID)
	vvSgt := interfaceToString(vSgt)
	vvSubnet := interfaceToString(vSubnet)
	vvVn := interfaceToString(vVn)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse2, _, err := client.FilterPolicy.GetFilterPolicyByID(vvID)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["sgt"] = vvSgt
				resourceMap["subnet"] = vvSubnet
				resourceMap["vn"] = vvVn
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		} else {
			queryParams2 := isegosdk.GetFilterPolicyQueryParams{}

			response2, _, err := client.FilterPolicy.GetFilterPolicy(&queryParams2)
			if response2 != nil && err == nil {
				items2 := getAllItemsFilterPolicyGetFilterPolicy(m, response2, &queryParams2)
				item2, nID, err := searchFilterPolicyGetFilterPolicy(m, items2, vvSgt, vvSubnet, vvVn, vvID)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["id"] = item2.Sgt
					if nID != "" && nID != vvID {
						resourceMap["id"] = nID
					}
					resourceMap["sgt"] = vvSgt
					resourceMap["subnet"] = vvSubnet
					resourceMap["vn"] = vvVn
					d.SetId(joinResourceID(resourceMap))
					return diags
				}
			}
		}
	}
	restyResp1, err := client.FilterPolicy.CreateFilterPolicy(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateFilterPolicy", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateFilterPolicy", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["sgt"] = vvSgt
	resourceMap["subnet"] = vvSubnet
	resourceMap["vn"] = vvVn
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceFilterPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning FilterPolicy read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vSgt, okSgt := resourceMap["sgt"]
	vSubnet, okSubnet := resourceMap["subnet"]
	vVn, okVn := resourceMap["vn"]
	vvSgt := vSgt
	vvSubnet := vSubnet
	vvVn := vVn
	vvID := vID
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okSgt, okSubnet, okVn}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetFilterPolicy")
		queryParams1 := isegosdk.GetFilterPolicyQueryParams{}

		response1, restyResp1, err := client.FilterPolicy.GetFilterPolicy(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsFilterPolicyGetFilterPolicy(m, response1, &queryParams1)
		item1, _, err := searchFilterPolicyGetFilterPolicy(m, items1, vvSgt, vvSubnet, vvVn, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenFilterPolicyGetFilterPolicyByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFilterPolicy search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFilterPolicy search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFilterPolicyByID")

		response2, restyResp2, err := client.FilterPolicy.GetFilterPolicyByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenFilterPolicyGetFilterPolicyByIDItem(response2.ERSFilterPolicy)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFilterPolicyByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFilterPolicyByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceFilterPolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning FilterPolicy update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vSgt, okSgt := resourceMap["sgt"]
	vSubnet, okSubnet := resourceMap["subnet"]
	vVn, okVn := resourceMap["vn"]
	vvSgt := vSgt
	vvSubnet := vSubnet
	vvVn := vVn
	vvID := vID
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okSgt, okSubnet, okVn}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetFilterPolicyQueryParams{}

		getResp1, _, err := client.FilterPolicy.GetFilterPolicy(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsFilterPolicyGetFilterPolicy(m, getResp1, &queryParams1)
			item1, nID, err := searchFilterPolicyGetFilterPolicy(m, items1, vvSgt, vvSubnet, vvVn, vvID)
			if err == nil && item1 != nil && nID != "" {
				vvID = nID
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestFilterPolicyUpdateFilterPolicyByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.FilterPolicy.UpdateFilterPolicyByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateFilterPolicyByID", err, restyResp1.String(),
					"Failure at UpdateFilterPolicyByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateFilterPolicyByID", err,
				"Failure at UpdateFilterPolicyByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceFilterPolicyRead(ctx, d, m)
}

func resourceFilterPolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning FilterPolicy delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vSgt, okSgt := resourceMap["sgt"]
	vSubnet, okSubnet := resourceMap["subnet"]
	vVn, okVn := resourceMap["vn"]
	vvSgt := vSgt
	vvSubnet := vSubnet
	vvVn := vVn
	vvID := vID
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okSgt, okSubnet, okVn}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetFilterPolicyQueryParams{}

		getResp1, _, err := client.FilterPolicy.GetFilterPolicy(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsFilterPolicyGetFilterPolicy(m, getResp1, &queryParams1)
		item1, nID, err := searchFilterPolicyGetFilterPolicy(m, items1, vvSgt, vvSubnet, vvVn, vvID)
		if err != nil || item1 == nil || nID == "" {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 1 {
		getResp, _, err := client.FilterPolicy.GetFilterPolicyByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.FilterPolicy.DeleteFilterPolicyByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteFilterPolicyByID", err, restyResp1.String(),
				"Failure at DeleteFilterPolicyByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteFilterPolicyByID", err,
			"Failure at DeleteFilterPolicyByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestFilterPolicyCreateFilterPolicy(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestFilterPolicyCreateFilterPolicy {
	request := isegosdk.RequestFilterPolicyCreateFilterPolicy{}
	request.ERSFilterPolicy = expandRequestFilterPolicyCreateFilterPolicyERSFilterPolicy(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestFilterPolicyCreateFilterPolicyERSFilterPolicy(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestFilterPolicyCreateFilterPolicyERSFilterPolicy {
	request := isegosdk.RequestFilterPolicyCreateFilterPolicyERSFilterPolicy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subnet")))) {
		request.Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains")))) {
		request.Domains = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt")))) {
		first, _ := replaceRegExStrings(interfaceToString(v), "", `\s*\(.*\)$`, "")
		request.Sgt = first
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn")))) {
		request.Vn = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestFilterPolicyUpdateFilterPolicyByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestFilterPolicyUpdateFilterPolicyByID {
	request := isegosdk.RequestFilterPolicyUpdateFilterPolicyByID{}
	request.ERSFilterPolicy = expandRequestFilterPolicyUpdateFilterPolicyByIDERSFilterPolicy(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestFilterPolicyUpdateFilterPolicyByIDERSFilterPolicy(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestFilterPolicyUpdateFilterPolicyByIDERSFilterPolicy {
	request := isegosdk.RequestFilterPolicyUpdateFilterPolicyByIDERSFilterPolicy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subnet")))) {
		request.Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains")))) {
		request.Domains = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt")))) {
		first, _ := replaceRegExStrings(interfaceToString(v), "", `\s*\(.*\)$`, "")
		request.Sgt = first
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn")))) {
		request.Vn = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsFilterPolicyGetFilterPolicy(m interface{}, response *isegosdk.ResponseFilterPolicyGetFilterPolicy, queryParams *isegosdk.GetFilterPolicyQueryParams) []isegosdk.ResponseFilterPolicyGetFilterPolicySearchResultResources {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var respItems []isegosdk.ResponseFilterPolicyGetFilterPolicySearchResultResources
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
			response, _, err = client.FilterPolicy.GetFilterPolicy(queryParams)
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

func searchFilterPolicyGetFilterPolicy(m interface{}, items []isegosdk.ResponseFilterPolicyGetFilterPolicySearchResultResources, sgt string, subnet string, vn string, id string) (*isegosdk.ResponseFilterPolicyGetFilterPolicyByIDERSFilterPolicy, string, error) {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var err error
	var foundID string
	var foundItem *isegosdk.ResponseFilterPolicyGetFilterPolicyByIDERSFilterPolicy
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseFilterPolicyGetFilterPolicyByID
			getItem, _, err = client.FilterPolicy.GetFilterPolicyByID(id)
			if err != nil {
				return foundItem, foundID, err
			}
			if getItem == nil {
				return foundItem, foundID, fmt.Errorf("Empty response from %s", "GetFilterPolicyByID")
			}
			foundItem = getItem.ERSFilterPolicy
			foundID = item.ID
			return foundItem, foundID, err
		} else {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseFilterPolicyGetFilterPolicyByID
			getItem, _, err = client.FilterPolicy.GetFilterPolicyByID(item.ID)
			if err != nil || getItem == nil {
				// Not found for some reason skip it
				continue
			}
			if getItem.ERSFilterPolicy != nil {
				hasSameVn := getItem.ERSFilterPolicy.Vn == vn
				hasSameSubnet := getItem.ERSFilterPolicy.Subnet == subnet
				hasSameSgt := compareSGT(getItem.ERSFilterPolicy.Sgt, sgt)
				if hasSameVn && hasSameSubnet && hasSameSgt {
					foundID = item.ID
					foundItem = getItem.ERSFilterPolicy
					return foundItem, foundID, err
				}
			}
		}
	}
	return foundItem, foundID, err
}
