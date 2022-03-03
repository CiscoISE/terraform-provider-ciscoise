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

func resourceDownloadableACL() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on DownloadableACL.

- This resource allows the client to update a downloadable ACL.

- This resource deletes a downloadable ACL.

- This resource creates a downloadable ACL.
`,

		CreateContext: resourceDownloadableACLCreate,
		ReadContext:   resourceDownloadableACLRead,
		UpdateContext: resourceDownloadableACLUpdate,
		DeleteContext: resourceDownloadableACLDelete,
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

						"dacl": &schema.Schema{
							Description: `The DACL Content. Use the string \\n for a newline`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"dacl_type": &schema.Schema{
							Description: `Allowed values:
- IPV4,
- IPV6,
- IP_AGNOSTIC`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `Use the string \\n for a newline`,
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
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"dacl": &schema.Schema{
							Description: `The DACL Content. Use the string \\n for a newline`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"dacl_type": &schema.Schema{
							Description: `Allowed values:
- IPV4,
- IPV6,
- IP_AGNOSTIC`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Description: `Use the string \\n for a newline`,
							Type:        schema.TypeString,
							Optional:    true,
						},
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

func resourceDownloadableACLCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DownloadableACL create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDownloadableACLCreateDownloadableACL(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.DownloadableACL.GetDownloadableACLByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceDownloadableACLRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetDownloadableACLQueryParams{}

		response2, _, err := client.DownloadableACL.GetDownloadableACL(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsDownloadableACLGetDownloadableACL(m, response2, &queryParams2)
			item2, err := searchDownloadableACLGetDownloadableACL(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceDownloadableACLRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.DownloadableACL.CreateDownloadableACL(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateDownloadableACL", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateDownloadableACL", err))
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
	return resourceDownloadableACLRead(ctx, d, m)
}

func resourceDownloadableACLRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DownloadableACL read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	vvID := vID
	vvName := vName

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDownloadableACL")
		queryParams1 := isegosdk.GetDownloadableACLQueryParams{}

		response1, restyResp1, err := client.DownloadableACL.GetDownloadableACL(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsDownloadableACLGetDownloadableACL(m, response1, &queryParams1)
		item1, err := searchDownloadableACLGetDownloadableACL(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenDownloadableACLGetDownloadableACLByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDownloadableACL search response",
				err))
			return diags
		}
		if err := d.Set("parameters", remove_parameters(vItem1, "link")); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDownloadableACL response to parameters",
				err))
			return diags
		}
		return diags
	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDownloadableACLByID")
		vvID := vID

		response2, restyResp2, err := client.DownloadableACL.GetDownloadableACLByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDownloadableACLGetDownloadableACLByIDItem(response2.DownloadableACL)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDownloadableACLByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", remove_parameters(vItem2, "link")); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDownloadableACLByID response to parameters",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceDownloadableACLUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DownloadableACL update for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetDownloadableACLQueryParams{}

		getResp1, _, err := client.DownloadableACL.GetDownloadableACL(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsDownloadableACLGetDownloadableACL(m, getResp1, &queryParams1)
			item1, err := searchDownloadableACLGetDownloadableACL(m, items1, vName, vID)
			if err == nil && getResp1 != nil {
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
		request1 := expandRequestDownloadableACLUpdateDownloadableACLByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.DownloadableACL.UpdateDownloadableACLByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDownloadableACLByID", err, restyResp1.String(),
					"Failure at UpdateDownloadableACLByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDownloadableACLByID", err,
				"Failure at UpdateDownloadableACLByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceDownloadableACLRead(ctx, d, m)
}

func resourceDownloadableACLDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DownloadableACL delete for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetDownloadableACLQueryParams{}

		getResp1, _, err := client.DownloadableACL.GetDownloadableACL(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsDownloadableACLGetDownloadableACL(m, getResp1, &queryParams1)
		item1, err := searchDownloadableACLGetDownloadableACL(m, items1, vName, vID)
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
		getResp, _, err := client.DownloadableACL.GetDownloadableACLByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.DownloadableACL.DeleteDownloadableACLByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDownloadableACLByID", err, restyResp1.String(),
				"Failure at DeleteDownloadableACLByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDownloadableACLByID", err,
			"Failure at DeleteDownloadableACLByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDownloadableACLCreateDownloadableACL(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDownloadableACLCreateDownloadableACL {
	request := isegosdk.RequestDownloadableACLCreateDownloadableACL{}
	request.DownloadableACL = expandRequestDownloadableACLCreateDownloadableACLDownloadableACL(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDownloadableACLCreateDownloadableACLDownloadableACL(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDownloadableACLCreateDownloadableACLDownloadableACL {
	request := isegosdk.RequestDownloadableACLCreateDownloadableACLDownloadableACL{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dacl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dacl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dacl")))) {
		request.Dacl = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dacl_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dacl_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dacl_type")))) {
		request.DaclType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDownloadableACLUpdateDownloadableACLByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDownloadableACLUpdateDownloadableACLByID {
	request := isegosdk.RequestDownloadableACLUpdateDownloadableACLByID{}
	request.DownloadableACL = expandRequestDownloadableACLUpdateDownloadableACLByIDDownloadableACL(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDownloadableACLUpdateDownloadableACLByIDDownloadableACL(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDownloadableACLUpdateDownloadableACLByIDDownloadableACL {
	request := isegosdk.RequestDownloadableACLUpdateDownloadableACLByIDDownloadableACL{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dacl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dacl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dacl")))) {
		request.Dacl = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dacl_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dacl_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dacl_type")))) {
		request.DaclType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsDownloadableACLGetDownloadableACL(m interface{}, response *isegosdk.ResponseDownloadableACLGetDownloadableACL, queryParams *isegosdk.GetDownloadableACLQueryParams) []isegosdk.ResponseDownloadableACLGetDownloadableACLSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseDownloadableACLGetDownloadableACLSearchResultResources
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
			response, _, err = client.DownloadableACL.GetDownloadableACL(queryParams)
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

func searchDownloadableACLGetDownloadableACL(m interface{}, items []isegosdk.ResponseDownloadableACLGetDownloadableACLSearchResultResources, name string, id string) (*isegosdk.ResponseDownloadableACLGetDownloadableACLByIDDownloadableACL, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseDownloadableACLGetDownloadableACLByIDDownloadableACL
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDownloadableACLGetDownloadableACLByID
			getItem, _, err = client.DownloadableACL.GetDownloadableACLByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDownloadableACLByID")
			}
			foundItem = getItem.DownloadableACL
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDownloadableACLGetDownloadableACLByID
			getItem, _, err = client.DownloadableACL.GetDownloadableACLByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDownloadableACLByID")
			}
			foundItem = getItem.DownloadableACL
			return foundItem, err
		}
	}
	return foundItem, err
}
