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

func resourceSxpConnections() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SXPConnections.

- This resource allows the client to update a SXP connection.

- This resource deletes a SXP connection.

- This resource creates a SXP connection.
`,

		CreateContext: resourceSxpConnectionsCreate,
		ReadContext:   resourceSxpConnectionsRead,
		UpdateContext: resourceSxpConnectionsUpdate,
		DeleteContext: resourceSxpConnectionsDelete,
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

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enabled": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_address": &schema.Schema{
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
						"sxp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sxp_node": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sxp_peer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sxp_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sxp_vpn": &schema.Schema{
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

func resourceSxpConnectionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestSxpConnectionsCreateSxpConnections(ctx, "item.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okID && vvID != "" {
		getResponse2, _, err := client.SxpConnections.GetSxpConnectionsByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		queryParams2 := isegosdk.GetSxpConnectionsQueryParams{}

		response2, _, err := client.SxpConnections.GetSxpConnections(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsSxpConnectionsGetSxpConnections(m, response2, &queryParams2)
			item2, err := searchSxpConnectionsGetSxpConnections(m, items2, "", vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	restyResp1, err := client.SxpConnections.CreateSxpConnections(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSxpConnections", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSxpConnections", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSxpConnectionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		vvID := vID

		log.Printf("[DEBUG] Selected method: GetSxpConnections")
		queryParams1 := isegosdk.GetSxpConnectionsQueryParams{}

		response1, _, err := client.SxpConnections.GetSxpConnections(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSxpConnections", err,
				"Failure at GetSxpConnections, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsSxpConnectionsGetSxpConnections(m, response1, &queryParams1)
		item1, err := searchSxpConnectionsGetSxpConnections(m, items1, "", vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetSxpConnections response", err,
				"Failure when searching item from GetSxpConnections, unexpected response", ""))
			return diags
		}
		vItem1 := flattenSxpConnectionsGetSxpConnectionsByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSxpConnections search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSxpConnectionsByID")
		vvID := vID

		response2, _, err := client.SxpConnections.GetSxpConnectionsByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSxpConnectionsByID", err,
				"Failure at GetSxpConnectionsByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSxpConnectionsGetSxpConnectionsByIDItem(response2.ERSSxpConnection)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSxpConnectionsByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSxpConnectionsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if selectedMethod == 2 {
		vvID = vID
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestSxpConnectionsUpdateSxpConnectionsByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SxpConnections.UpdateSxpConnectionsByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSxpConnectionsByID", err, restyResp1.String(),
					"Failure at UpdateSxpConnectionsByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSxpConnectionsByID", err,
				"Failure at UpdateSxpConnectionsByID, unexpected response", ""))
			return diags
		}
	}

	return resourceSxpConnectionsRead(ctx, d, m)
}

func resourceSxpConnectionsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {
		queryParams1 := isegosdk.GetSxpConnectionsQueryParams{}

		getResp1, _, err := client.SxpConnections.GetSxpConnections(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSxpConnectionsGetSxpConnections(m, getResp1, &queryParams1)
		item1, err := searchSxpConnectionsGetSxpConnections(m, items1, "", vID)
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
	if selectedMethod == 2 {
		vvID = vID
		getResp, _, err := client.SxpConnections.GetSxpConnectionsByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.SxpConnections.DeleteSxpConnectionsByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSxpConnectionsByID", err, restyResp1.String(),
				"Failure at DeleteSxpConnectionsByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSxpConnectionsByID", err,
			"Failure at DeleteSxpConnectionsByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSxpConnectionsCreateSxpConnections(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpConnectionsCreateSxpConnections {
	request := isegosdk.RequestSxpConnectionsCreateSxpConnections{}
	request.ERSSxpConnection = expandRequestSxpConnectionsCreateSxpConnectionsERSSxpConnection(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSxpConnectionsCreateSxpConnectionsERSSxpConnection(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpConnectionsCreateSxpConnectionsERSSxpConnection {
	request := isegosdk.RequestSxpConnectionsCreateSxpConnectionsERSSxpConnection{}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_peer"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_peer"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_peer"))) {
		request.SxpPeer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_vpn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_vpn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_vpn"))) {
		request.SxpVpn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_node"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_node"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_node"))) {
		request.SxpNode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ip_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ip_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ip_address"))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_mode"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_mode"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_mode"))) {
		request.SxpMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_version"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_version"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_version"))) {
		request.SxpVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSxpConnectionsUpdateSxpConnectionsByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpConnectionsUpdateSxpConnectionsByID {
	request := isegosdk.RequestSxpConnectionsUpdateSxpConnectionsByID{}
	request.ERSSxpConnection = expandRequestSxpConnectionsUpdateSxpConnectionsByIDERSSxpConnection(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSxpConnectionsUpdateSxpConnectionsByIDERSSxpConnection(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpConnectionsUpdateSxpConnectionsByIDERSSxpConnection {
	request := isegosdk.RequestSxpConnectionsUpdateSxpConnectionsByIDERSSxpConnection{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_peer"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_peer"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_peer"))) {
		request.SxpPeer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_vpn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_vpn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_vpn"))) {
		request.SxpVpn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_node"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_node"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_node"))) {
		request.SxpNode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ip_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ip_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ip_address"))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_mode"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_mode"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_mode"))) {
		request.SxpMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sxp_version"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxp_version"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxp_version"))) {
		request.SxpVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsSxpConnectionsGetSxpConnections(m interface{}, response *isegosdk.ResponseSxpConnectionsGetSxpConnections, queryParams *isegosdk.GetSxpConnectionsQueryParams) []isegosdk.ResponseSxpConnectionsGetSxpConnectionsSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseSxpConnectionsGetSxpConnectionsSearchResultResources
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
			response, _, err = client.SxpConnections.GetSxpConnections(queryParams)
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

func searchSxpConnectionsGetSxpConnections(m interface{}, items []isegosdk.ResponseSxpConnectionsGetSxpConnectionsSearchResultResources, name string, id string) (*isegosdk.ResponseSxpConnectionsGetSxpConnectionsByIDERSSxpConnection, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSxpConnectionsGetSxpConnectionsByIDERSSxpConnection
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSxpConnectionsGetSxpConnectionsByID
			getItem, _, err = client.SxpConnections.GetSxpConnectionsByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSxpConnectionsByID")
			}
			foundItem = getItem.ERSSxpConnection
			return foundItem, err
		}
	}
	return foundItem, err
}
