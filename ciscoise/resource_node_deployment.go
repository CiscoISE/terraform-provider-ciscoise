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

func resourceNodeDeployment() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Node Deployment.

- This resource registers a Cisco ISE node to form a multi-node deployment.
 Approximate execution time 300 seconds.

- This resource updates the configuration of the Cisco ISE node with the configuration provided.
 Approximate execution time 300 seconds.

- The deregistered node becomes a standalone Cisco ISE node.
 It retains the last configuration that it received from the primary PAN and assumes the default roles and services of a
standalone node.
 Approximate execution time 300 seconds.
`,

		CreateContext: resourceNodeDeploymentCreate,
		ReadContext:   resourceNodeDeploymentRead,
		UpdateContext: resourceNodeDeploymentUpdate,
		DeleteContext: resourceNodeDeploymentDelete,
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

						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": &schema.Schema{
							Description: `Roles can be empty or have many values for a node. `,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"services": &schema.Schema{
							Description: `Services can be empty or have many values for a node. `,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

						"allow_cert_import": &schema.Schema{
							Description:  `Consent to import the self-signed certificate of the registering node. `,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"hostname": &schema.Schema{
							Description: `hostname path parameter. Hostname of the deployed node.`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"roles": &schema.Schema{
							Description: `Roles can be empty or have many values for a node. `,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"services": &schema.Schema{
							Description: `Services can be empty or have many values for a node. `,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceNodeDeploymentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeDeployment create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNodeDeploymentRegisterNode(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vHostname, okHostname := resourceItem["hostname"]
	vvHostname := interfaceToString(vHostname)
	vFQDN, _ := resourceItem["fqdn"]
	vvFQDN := interfaceToString(vFQDN)
	if okHostname && vvHostname != "" {
		getResponse2, _, err := client.NodeDeployment.GetNodeDetails(vvHostname)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["hostname"] = vvHostname
			resourceMap["fqdn"] = vvFQDN
			d.SetId(joinResourceID(resourceMap))
			return resourceNodeDeploymentRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.NodeDeployment.GetDeploymentNodes(nil)
		if response2 != nil && err == nil {
			items2 := getAllItemsNodeDeploymentGetDeploymentNodes(m, response2)
			item2, err := searchNodeDeploymentGetDeploymentNodes(m, items2, vvHostname, vvFQDN, "")
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["hostname"] = vvHostname
				resourceMap["fqdn"] = vvFQDN
				d.SetId(joinResourceID(resourceMap))
				return resourceNodeDeploymentRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.NodeDeployment.RegisterNode(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing RegisterNode", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing RegisterNode", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["hostname"] = vvHostname
	resourceMap["fqdn"] = vvFQDN
	d.SetId(joinResourceID(resourceMap))
	return resourceNodeDeploymentRead(ctx, d, m)
}

func resourceNodeDeploymentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeDeployment read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vHostname, okHostname := resourceMap["hostname"]
	vFQDN, _ := resourceMap["fqdn"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostname}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeploymentNodes")
		queryParams1 := isegosdk.GetDeploymentNodesQueryParams{}

		response1, restyResp1, err := client.NodeDeployment.GetDeploymentNodes(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsNodeDeploymentGetDeploymentNodes(m, response1)
		item1, err := searchNodeDeploymentGetDeploymentNodes(m, items1, vHostname, vFQDN, "")
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenNodeDeploymentGetNodeDetailsItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeploymentNodes search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNodeDetails")
		vvHostname := vHostname

		response2, restyResp2, err := client.NodeDeployment.GetNodeDetails(vvHostname)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNodeDeploymentGetNodeDetailsItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeDetails response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNodeDeploymentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeDeployment update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vHostname, okHostname := resourceMap["hostname"]
	vFQDN, _ := resourceMap["fqdn"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostname}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvHostname string
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 1 {
		getResp1, _, err := client.NodeDeployment.GetDeploymentNodes(nil)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsNodeDeploymentGetDeploymentNodes(m, getResp1)
			item1, err := searchNodeDeploymentGetDeploymentNodes(m, items1, vHostname, vFQDN, "")
			if err == nil && item1 != nil {
				if vHostname != item1.Hostname {
					vvHostname = item1.Hostname
				} else {
					vvHostname = vHostname
				}
			}
		}
	}
	if selectedMethod == 2 {
		vvHostname = vHostname
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvHostname)
		request1 := expandRequestNodeDeploymentUpdateDeploymentNode(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.NodeDeployment.UpdateDeploymentNode(vvHostname, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDeploymentNode", err, restyResp1.String(),
					"Failure at UpdateDeploymentNode, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeploymentNode", err,
				"Failure at UpdateDeploymentNode, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceNodeDeploymentRead(ctx, d, m)
}

func resourceNodeDeploymentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeDeployment delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vHostname, okHostname := resourceMap["hostname"]
	vFQDN, _ := resourceMap["fqdn"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostname}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvHostname string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.NodeDeployment.GetDeploymentNodes(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNodeDeploymentGetDeploymentNodes(m, getResp1)
		item1, err := searchNodeDeploymentGetDeploymentNodes(m, items1, vHostname, vFQDN, "")
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vHostname != item1.Hostname {
			vvHostname = item1.Hostname
		} else {
			vvHostname = vHostname
		}
	}
	if selectedMethod == 2 {
		vvHostname = vHostname
		getResp, _, err := client.NodeDeployment.GetNodeDetails(vvHostname)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NodeDeployment.DeleteDeploymentNode(vvHostname)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeploymentNode", err, restyResp1.String(),
				"Failure at DeleteDeploymentNode, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeploymentNode", err,
			"Failure at DeleteDeploymentNode, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNodeDeploymentRegisterNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNode {
	request := isegosdk.RequestNodeDeploymentRegisterNode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_cert_import")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_cert_import")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_cert_import")))) {
		request.AllowCertImport = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".roles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".roles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".roles")))) {
		request.Roles = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services")))) {
		request.Services = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateDeploymentNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateDeploymentNode {
	request := isegosdk.RequestNodeDeploymentUpdateDeploymentNode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".roles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".roles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".roles")))) {
		request.Roles = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services")))) {
		request.Services = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNodeDeploymentGetDeploymentNodes(m interface{}, response *isegosdk.ResponseNodeDeploymentGetDeploymentNodes) []isegosdk.ResponseNodeDeploymentGetDeploymentNodesResponse {
	var respItems []isegosdk.ResponseNodeDeploymentGetDeploymentNodesResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNodeDeploymentGetDeploymentNodes(m interface{}, items []isegosdk.ResponseNodeDeploymentGetDeploymentNodesResponse, name string, fqdn string, id string) (*isegosdk.ResponseNodeDeploymentGetNodeDetailsResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNodeDeploymentGetNodeDetailsResponse
	for _, item := range items {
		if fqdn != "" && item.Fqdn == fqdn {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNodeDeploymentGetNodeDetails
			getItem, _, err = client.NodeDeployment.GetNodeDetails(item.Hostname)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNodeDetails")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
		if name != "" && item.Hostname == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNodeDeploymentGetNodeDetails
			getItem, _, err = client.NodeDeployment.GetNodeDetails(name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNodeDetails")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
