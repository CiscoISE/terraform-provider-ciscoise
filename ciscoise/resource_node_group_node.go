package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNodeGroupNode() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and delete operation on Node Group.

- This resource action adds a node to the node group in the cluster. When a node that belongs to a node group fails,
another node in the same node group issues a Change of Authorization (CoA) for all the URL-redirected sessions on the
failed node.

- Purpose of this API is to remove a node from a node group in the cluster. Removing node from the node group does not
delete the node, but failover is no longer carried out if the node is not part any node group.

`,

		CreateContext: resourceNodeGroupNodeCreate,
		ReadContext:   resourceNodeGroupNodeRead,
		UpdateContext: resourceNodeGroupNodeCreate,
		DeleteContext: resourceNodeGroupNodeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_group_name": &schema.Schema{
							Description: `nodeGroupName path parameter. Name of the existing node group.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"success": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"message": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceNodeGroupNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeGroupNode read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vNodeGroupName := resourceMap["node_group_name"]
	vHostname := resourceMap["hostname"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNodes")
		vvNodeGroupName := interfaceToString(vNodeGroupName)
		vvHostname := interfaceToString(vHostname)

		response1, restyResp1, err := client.NodeGroup.GetNodes(vvNodeGroupName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		response_nodes, err := searchNodeGroupGetNodes(m, response1.Response, vvHostname)
		if err != nil || response_nodes == nil || len(*response_nodes) == 0 {
			d.SetId("")
			return diags
		}
		vItems1 := flattenNodeGroupGetNodesItems(response_nodes)
		if err := d.Set("item", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodes response",
				err))
			return diags
		}
		return diags
	}
	return diags
}

func resourceNodeGroupNodeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeGroupNode create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vNodeGroupName := resourceItem["node_group_name"]
	vHostname := resourceItem["hostname"]

	vvNodeGroupName := vNodeGroupName.(string)
	vvHostname := vHostname.(string)

	selectedMethod := 1
	if selectedMethod == 1 {
		response1, _, err := client.NodeGroup.GetNodes(vvNodeGroupName)
		if err == nil && response1 != nil {
			response_nodes, err := searchNodeGroupGetNodes(m, response1.Response, vvHostname)
			if err == nil && response_nodes != nil && len(*response_nodes) > 0 {
				resourceMap := make(map[string]string)
				resourceMap["node_group_name"] = vvNodeGroupName
				resourceMap["hostname"] = vvHostname
				d.SetId(joinResourceID(resourceMap))
				return resourceNodeGroupNodeRead(ctx, d, m)
			}
		}
	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: AddNode")
		request1 := expandRequestNodeGroupNodeCreateAddNode(ctx, "parameters.0", d)
		response1, restyResp1, err := client.NodeGroup.AddNode(vvNodeGroupName, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AddNode", err,
				"Failure at AddNode, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNodeGroupAddNodeItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting AddNode response",
				err))
			return diags
		}
		resourceMap := make(map[string]string)
		resourceMap["node_group_name"] = vvNodeGroupName
		resourceMap["hostname"] = vvHostname
		d.SetId(joinResourceID(resourceMap))
		return resourceNodeGroupNodeRead(ctx, d, m)
	}
	return diags
}

func resourceNodeGroupNodeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vNodeGroupName := resourceItem["node_group_name"]
	vHostname := resourceItem["hostname"]

	vvNodeGroupName := vNodeGroupName.(string)
	vvHostname := vHostname.(string)

	selectedMethod := 1
	if selectedMethod == 1 {
		response1, _, err := client.NodeGroup.GetNodes(vvNodeGroupName)
		if err != nil || response1 == nil {
			// Assume that element it is already gone
			return diags
		}
		response_nodes, err := searchNodeGroupGetNodes(m, response1.Response, vvHostname)
		if err != nil || response_nodes == nil || len(*response_nodes) == 0 {
			// Assume that element it is already gone
			return resourceNodeGroupNodeRead(ctx, d, m)
		}
	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RemoveNode")
		vvNodeGroupName := vNodeGroupName.(string)
		request1 := expandRequestNodeGroupNodeDeleteRemoveNode(ctx, "parameters.0", d)

		response1, restyResp1, err := client.NodeGroup.RemoveNode(vvNodeGroupName, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RemoveNode", err,
				"Failure at RemoveNode, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNodeGroupRemoveNodeItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RemoveNode response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func expandRequestNodeGroupNodeDeleteRemoveNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeGroupRemoveNode {
	request := isegosdk.RequestNodeGroupRemoveNode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	return &request
}

func flattenNodeGroupRemoveNodeItem(item *isegosdk.ResponseNodeGroupRemoveNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = flattenNodeGroupRemoveNodeItemSuccess(item.Success)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeGroupRemoveNodeItemSuccess(item *isegosdk.ResponseNodeGroupRemoveNodeSuccess) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message

	return []map[string]interface{}{
		respItem,
	}

}

func expandRequestNodeGroupNodeCreateAddNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeGroupAddNode {
	request := isegosdk.RequestNodeGroupAddNode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	return &request
}

func flattenNodeGroupAddNodeItem(item *isegosdk.ResponseNodeGroupAddNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = flattenNodeGroupAddNodeItemSuccess(item.Success)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeGroupAddNodeItemSuccess(item *isegosdk.ResponseNodeGroupAddNodeSuccess) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message

	return []map[string]interface{}{
		respItem,
	}

}

func searchNodeGroupGetNodes(m interface{}, items *[]isegosdk.ResponseNodeGroupGetNodesResponse, name string) (*[]isegosdk.ResponseNodeGroupGetNodesResponse, error) {
	var err error
	var foundItems []isegosdk.ResponseNodeGroupGetNodesResponse
	if items == nil {
		return nil, err
	}
	for _, item := range *items {
		if name != "" && item.Hostname == name {
			// Call get by _ method and set value to foundItem and return
			foundItems = append(foundItems, item)
			return &foundItems, err
		}
	}
	return nil, err
}
