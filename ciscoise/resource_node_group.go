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

func resourceNodeGroup() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Node Group.

- Developers need to create node group in the system.Node Group is a group of PSNs, mainly used for terminating posture
pending sessions when a PSN in local node group fails.Node group members can communicate over TCP/7800.

- API updates an existing node group in the system.

- Developers need to delete node group in the system.
`,

		CreateContext: resourceNodeGroupCreate,
		ReadContext:   resourceNodeGroupRead,
		UpdateContext: resourceNodeGroupUpdate,
		DeleteContext: resourceNodeGroupDelete,
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
						"mar_cache": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enabled": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"query_attempts": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"query_timeout": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"replication_attempts": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"replication_timeout": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_group_name": &schema.Schema{
							Description: `node-group-name path parameter. ID of the existing node group.`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceNodeGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestNodeGroupCreateNodeGroup(ctx, "item.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vNodeGroupName, okNodeGroupName := resourceItem["node_group_name"]
	vvNodeGroupName := interfaceToString(vNodeGroupName)
	if okNodeGroupName && vvNodeGroupName != "" {
		getResponse2, _, err := client.NodeGroup.GetNodeGroup(vvNodeGroupName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["node_group_name"] = vvNodeGroupName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.NodeGroup.GetNodeGroups()
		if response2 != nil && err == nil {
			items2 := getAllItemsNodeGroupGetNodeGroups(m, response2)
			item2, err := searchNodeGroupGetNodeGroups(m, items2, vvNodeGroupName, "")
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["node_group_name"] = vvNodeGroupName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	resp1, restyResp1, err := client.NodeGroup.CreateNodeGroup(vvNodeGroupName, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNodeGroup", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNodeGroup", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["node_group_name"] = vvNodeGroupName
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNodeGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vNodeGroupName, okNodeGroupName := resourceMap["node_group_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okNodeGroupName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		vvNodeGroupName := vNodeGroupName
		log.Printf("[DEBUG] Selected method: GetNodeGroups")

		response1, _, err := client.NodeGroup.GetNodeGroups()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeGroups", err,
				"Failure at GetNodeGroups, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsNodeGroupGetNodeGroups(m, response1)
		item1, err := searchNodeGroupGetNodeGroups(m, items1, vvNodeGroupName, "")
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetNodeGroups response", err,
				"Failure when searching item from GetNodeGroups, unexpected response", ""))
			return diags
		}
		vItem1 := flattenNodeGroupGetNodeGroupItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeGroups search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNodeGroup")
		vvNodeGroupName := vNodeGroupName

		response2, _, err := client.NodeGroup.GetNodeGroup(vvNodeGroupName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeGroup", err,
				"Failure at GetNodeGroup, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNodeGroupGetNodeGroupItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeGroup response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNodeGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vNodeGroupName, okNodeGroupName := resourceMap["node_group_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okNodeGroupName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvNodeGroupName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if selectedMethod == 2 {
		vvNodeGroupName = vNodeGroupName
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] NodeGroupName used for update operation %s", vvNodeGroupName)
		request1 := expandRequestNodeGroupUpdateNodeGroup(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NodeGroup.UpdateNodeGroup(vvNodeGroupName, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNodeGroup", err, restyResp1.String(),
					"Failure at UpdateNodeGroup, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNodeGroup", err,
				"Failure at UpdateNodeGroup, unexpected response", ""))
			return diags
		}
	}

	return resourceNodeGroupRead(ctx, d, m)
}

func resourceNodeGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vNodeGroupName, okNodeGroupName := resourceMap["node_group_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okNodeGroupName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvNodeGroupName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.NodeGroup.GetNodeGroups()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNodeGroupGetNodeGroups(m, getResp1)
		item1, err := searchNodeGroupGetNodeGroups(m, items1, vvNodeGroupName, "")
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		vvNodeGroupName = vNodeGroupName
	}
	if selectedMethod == 2 {
		vvNodeGroupName = vNodeGroupName
		getResp, _, err := client.NodeGroup.GetNodeGroup(vvNodeGroupName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NodeGroup.DeleteNodeGroup(vvNodeGroupName)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNodeGroup", err, restyResp1.String(),
				"Failure at DeleteNodeGroup, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNodeGroup", err,
			"Failure at DeleteNodeGroup, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNodeGroupCreateNodeGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeGroupCreateNodeGroup {
	request := isegosdk.RequestNodeGroupCreateNodeGroup{}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".mar_cache"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mar_cache"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mar_cache"))) {
		request.MarCache = expandRequestNodeGroupCreateNodeGroupMarCache(ctx, key+".mar_cache.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeGroupCreateNodeGroupMarCache(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeGroupCreateNodeGroupMarCache {
	request := isegosdk.RequestNodeGroupCreateNodeGroupMarCache{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".replication_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".replication_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".replication_timeout"))) {
		request.ReplicationTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".replication_attempts"); !isEmptyValue(reflect.ValueOf(d.Get(key+".replication_attempts"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".replication_attempts"))) {
		request.ReplicationAttempts = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".query_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".query_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".query_timeout"))) {
		request.QueryTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".query_attempts"); !isEmptyValue(reflect.ValueOf(d.Get(key+".query_attempts"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".query_attempts"))) {
		request.QueryAttempts = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeGroupUpdateNodeGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeGroupUpdateNodeGroup {
	request := isegosdk.RequestNodeGroupUpdateNodeGroup{}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".mar_cache"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mar_cache"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mar_cache"))) {
		request.MarCache = expandRequestNodeGroupUpdateNodeGroupMarCache(ctx, key+".mar_cache.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeGroupUpdateNodeGroupMarCache(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeGroupUpdateNodeGroupMarCache {
	request := isegosdk.RequestNodeGroupUpdateNodeGroupMarCache{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".replication_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".replication_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".replication_timeout"))) {
		request.ReplicationTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".replication_attempts"); !isEmptyValue(reflect.ValueOf(d.Get(key+".replication_attempts"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".replication_attempts"))) {
		request.ReplicationAttempts = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".query_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".query_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".query_timeout"))) {
		request.QueryTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".query_attempts"); !isEmptyValue(reflect.ValueOf(d.Get(key+".query_attempts"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".query_attempts"))) {
		request.QueryAttempts = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNodeGroupGetNodeGroups(m interface{}, response *isegosdk.ResponseNodeGroupGetNodeGroups) []isegosdk.ResponseNodeGroupGetNodeGroupsResponse {
	var respItems []isegosdk.ResponseNodeGroupGetNodeGroupsResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNodeGroupGetNodeGroups(m interface{}, items []isegosdk.ResponseNodeGroupGetNodeGroupsResponse, name string, id string) (*isegosdk.ResponseNodeGroupGetNodeGroup, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNodeGroupGetNodeGroup
	for _, item := range items {
		if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNodeGroupGetNodeGroup
			getItem, _, err = client.NodeGroup.GetNodeGroup(name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNodeGroup")
			}
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
