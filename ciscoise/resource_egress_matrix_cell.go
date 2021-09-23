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

func resourceEgressMatrixCell() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on EgressMatrixCell.

- This resource allows the client to update an egress matrix cell.

- This resource deletes an egress matrix cell.

- This resource creates an egress matrix cell.
`,

		CreateContext: resourceEgressMatrixCellCreate,
		ReadContext:   resourceEgressMatrixCellRead,
		UpdateContext: resourceEgressMatrixCellUpdate,
		DeleteContext: resourceEgressMatrixCellDelete,
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

						"default_rule": &schema.Schema{
							Description: `Allowed values:
- NONE,
- DENY_IP,
- PERMIT_IP`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination_sgt_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
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
						"matrix_cell_status": &schema.Schema{
							Description: `Allowed values:
- DISABLED,
- ENABLED,
- MONITOR`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sgacls": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"source_sgt_id": &schema.Schema{
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

func resourceEgressMatrixCellCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestEgressMatrixCellCreateEgressMatrixCell(ctx, "item.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.EgressMatrixCell.GetEgressMatrixCellByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		queryParams2 := isegosdk.GetEgressMatrixCellQueryParams{}

		response2, _, err := client.EgressMatrixCell.GetEgressMatrixCell(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsEgressMatrixCellGetEgressMatrixCell(m, response2, &queryParams2)
			item2, err := searchEgressMatrixCellGetEgressMatrixCell(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	restyResp1, err := client.EgressMatrixCell.CreateEgressMatrixCell(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateEgressMatrixCell", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateEgressMatrixCell", err))
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

func resourceEgressMatrixCellRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	vvName := vName
	vvID := vID
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetEgressMatrixCell")
		queryParams1 := isegosdk.GetEgressMatrixCellQueryParams{}

		response1, _, err := client.EgressMatrixCell.GetEgressMatrixCell(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEgressMatrixCell", err,
				"Failure at GetEgressMatrixCell, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsEgressMatrixCellGetEgressMatrixCell(m, response1, &queryParams1)
		item1, err := searchEgressMatrixCellGetEgressMatrixCell(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetEgressMatrixCell response", err,
				"Failure when searching item from GetEgressMatrixCell, unexpected response", ""))
			return diags
		}
		vItem1 := flattenEgressMatrixCellGetEgressMatrixCellByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEgressMatrixCell search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEgressMatrixCellByID")

		response2, _, err := client.EgressMatrixCell.GetEgressMatrixCellByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEgressMatrixCellByID", err,
				"Failure at GetEgressMatrixCellByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenEgressMatrixCellGetEgressMatrixCellByIDItem(response2.EgressMatrixCell)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEgressMatrixCellByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceEgressMatrixCellUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		queryParams1 := isegosdk.GetEgressMatrixCellQueryParams{}

		getResp1, _, err := client.EgressMatrixCell.GetEgressMatrixCell(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsEgressMatrixCellGetEgressMatrixCell(m, getResp1, &queryParams1)
			item1, err := searchEgressMatrixCellGetEgressMatrixCell(m, items1, vName, vID)
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
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestEgressMatrixCellUpdateEgressMatrixCellByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.EgressMatrixCell.UpdateEgressMatrixCellByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateEgressMatrixCellByID", err, restyResp1.String(),
					"Failure at UpdateEgressMatrixCellByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateEgressMatrixCellByID", err,
				"Failure at UpdateEgressMatrixCellByID, unexpected response", ""))
			return diags
		}
	}

	return resourceEgressMatrixCellRead(ctx, d, m)
}

func resourceEgressMatrixCellDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		queryParams1 := isegosdk.GetEgressMatrixCellQueryParams{}

		getResp1, _, err := client.EgressMatrixCell.GetEgressMatrixCell(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsEgressMatrixCellGetEgressMatrixCell(m, getResp1, &queryParams1)
		item1, err := searchEgressMatrixCellGetEgressMatrixCell(m, items1, vName, vID)
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
		getResp, _, err := client.EgressMatrixCell.GetEgressMatrixCellByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.EgressMatrixCell.DeleteEgressMatrixCellByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteEgressMatrixCellByID", err, restyResp1.String(),
				"Failure at DeleteEgressMatrixCellByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteEgressMatrixCellByID", err,
			"Failure at DeleteEgressMatrixCellByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestEgressMatrixCellCreateEgressMatrixCell(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEgressMatrixCellCreateEgressMatrixCell {
	request := isegosdk.RequestEgressMatrixCellCreateEgressMatrixCell{}
	request.EgressMatrixCell = expandRequestEgressMatrixCellCreateEgressMatrixCellEgressMatrixCell(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEgressMatrixCellCreateEgressMatrixCellEgressMatrixCell(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEgressMatrixCellCreateEgressMatrixCellEgressMatrixCell {
	request := isegosdk.RequestEgressMatrixCellCreateEgressMatrixCellEgressMatrixCell{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".source_sgt_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".source_sgt_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".source_sgt_id"))) {
		request.SourceSgtID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".destination_sgt_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".destination_sgt_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".destination_sgt_id"))) {
		request.DestinationSgtID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".matrix_cell_status"); !isEmptyValue(reflect.ValueOf(d.Get(key+".matrix_cell_status"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".matrix_cell_status"))) {
		request.MatrixCellStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_rule"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_rule"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_rule"))) {
		request.DefaultRule = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sgacls"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sgacls"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sgacls"))) {
		request.Sgacls = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEgressMatrixCellUpdateEgressMatrixCellByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEgressMatrixCellUpdateEgressMatrixCellByID {
	request := isegosdk.RequestEgressMatrixCellUpdateEgressMatrixCellByID{}
	request.EgressMatrixCell = expandRequestEgressMatrixCellUpdateEgressMatrixCellByIDEgressMatrixCell(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEgressMatrixCellUpdateEgressMatrixCellByIDEgressMatrixCell(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEgressMatrixCellUpdateEgressMatrixCellByIDEgressMatrixCell {
	request := isegosdk.RequestEgressMatrixCellUpdateEgressMatrixCellByIDEgressMatrixCell{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".source_sgt_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".source_sgt_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".source_sgt_id"))) {
		request.SourceSgtID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".destination_sgt_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".destination_sgt_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".destination_sgt_id"))) {
		request.DestinationSgtID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".matrix_cell_status"); !isEmptyValue(reflect.ValueOf(d.Get(key+".matrix_cell_status"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".matrix_cell_status"))) {
		request.MatrixCellStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_rule"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_rule"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_rule"))) {
		request.DefaultRule = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sgacls"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sgacls"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sgacls"))) {
		request.Sgacls = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsEgressMatrixCellGetEgressMatrixCell(m interface{}, response *isegosdk.ResponseEgressMatrixCellGetEgressMatrixCell, queryParams *isegosdk.GetEgressMatrixCellQueryParams) []isegosdk.ResponseEgressMatrixCellGetEgressMatrixCellSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseEgressMatrixCellGetEgressMatrixCellSearchResultResources
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
			response, _, err = client.EgressMatrixCell.GetEgressMatrixCell(queryParams)
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

func searchEgressMatrixCellGetEgressMatrixCell(m interface{}, items []isegosdk.ResponseEgressMatrixCellGetEgressMatrixCellSearchResultResources, name string, id string) (*isegosdk.ResponseEgressMatrixCellGetEgressMatrixCellByIDEgressMatrixCell, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseEgressMatrixCellGetEgressMatrixCellByIDEgressMatrixCell
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseEgressMatrixCellGetEgressMatrixCellByID
			getItem, _, err = client.EgressMatrixCell.GetEgressMatrixCellByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetEgressMatrixCellByID")
			}
			foundItem = getItem.EgressMatrixCell
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseEgressMatrixCellGetEgressMatrixCellByID
			getItem, _, err = client.EgressMatrixCell.GetEgressMatrixCellByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetEgressMatrixCellByID")
			}
			foundItem = getItem.EgressMatrixCell
			return foundItem, err
		}
	}
	return foundItem, err
}
