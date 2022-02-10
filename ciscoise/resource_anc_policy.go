package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAncPolicy() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on AncPolicy.

- This resource allows the client to update an ANC policy.

- This resource allows the client to delete an ANC policy.

- This resource allows the client to create an ANC policy.
`,

		CreateContext: resourceAncPolicyCreate,
		ReadContext:   resourceAncPolicyRead,
		UpdateContext: resourceAncPolicyUpdate,
		DeleteContext: resourceAncPolicyDelete,
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

						"actions": &schema.Schema{
							Description: `- QUARANTINE: Allows you to use Exception policies (authorization policies) to limit or deny an endpoint access to the network.
- PORTBOUNCE: Resets the port on the network device to which the endpoint is connected.
- SHUTDOWN : Shuts down the port on the network device to which the endpoint is connected.
- RE_AUTHENTICATE: Re-authenticates the session from the endpoint.`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

						"actions": &schema.Schema{
							Description: `- QUARANTINE: Allows you to use Exception policies (authorization policies) to limit or deny an endpoint access to the network.
- PORTBOUNCE: Resets the port on the network device to which the endpoint is connected.
- SHUTDOWN : Shuts down the port on the network device to which the endpoint is connected.
- RE_AUTHENTICATE: Re-authenticates the session from the endpoint.`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAncPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AncPolicy create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestAncPolicyCreateAncPolicy(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.AncPolicy.GetAncPolicyByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceAncPolicyRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.AncPolicy.GetAncPolicyByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceAncPolicyRead(ctx, d, m)
		}
	}
	restyResp1, err := client.AncPolicy.CreateAncPolicy(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateAncPolicy", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateAncPolicy", err))
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
	return resourceAncPolicyRead(ctx, d, m)
}

func resourceAncPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AncPolicy read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetAncPolicyByName")
		vvName := vName

		response1, restyResp1, err := client.AncPolicy.GetAncPolicyByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenAncPolicyGetAncPolicyByNameItemName(response1.ErsAncPolicy)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAncPolicyByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAncPolicyByID")
		vvID := vID

		response2, restyResp2, err := client.AncPolicy.GetAncPolicyByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenAncPolicyGetAncPolicyByIDItemID(response2.ErsAncPolicy)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAncPolicyByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceAncPolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AncPolicy update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.AncPolicy.GetAncPolicyByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAncPolicyByName", err,
				"Failure at GetAncPolicyByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.ErsAncPolicy != nil {
			vvID = getResp.ErsAncPolicy.ID
		}
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestAncPolicyUpdateAncPolicyByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.AncPolicy.UpdateAncPolicyByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateAncPolicyByID", err, restyResp1.String(),
					"Failure at UpdateAncPolicyByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateAncPolicyByID", err,
				"Failure at UpdateAncPolicyByID, unexpected response", ""))
			return diags
		}
		d.Set("last_updated", getUnixTimeString())
	}

	return resourceAncPolicyRead(ctx, d, m)
}

func resourceAncPolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AncPolicy delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.AncPolicy.GetAncPolicyByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.AncPolicy.GetAncPolicyByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.ErsAncPolicy != nil {
			vvID = getResp.ErsAncPolicy.ID
		}
	}
	restyResp1, err := client.AncPolicy.DeleteAncPolicyByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteAncPolicyByID", err, restyResp1.String(),
				"Failure at DeleteAncPolicyByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteAncPolicyByID", err,
			"Failure at DeleteAncPolicyByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestAncPolicyCreateAncPolicy(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncPolicyCreateAncPolicy {
	request := isegosdk.RequestAncPolicyCreateAncPolicy{}
	request.ErsAncPolicy = expandRequestAncPolicyCreateAncPolicyErsAncPolicy(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAncPolicyCreateAncPolicyErsAncPolicy(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncPolicyCreateAncPolicyErsAncPolicy {
	request := isegosdk.RequestAncPolicyCreateAncPolicyErsAncPolicy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".actions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".actions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".actions")))) {
		request.Actions = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAncPolicyUpdateAncPolicyByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncPolicyUpdateAncPolicyByID {
	request := isegosdk.RequestAncPolicyUpdateAncPolicyByID{}
	request.ErsAncPolicy = expandRequestAncPolicyUpdateAncPolicyByIDErsAncPolicy(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAncPolicyUpdateAncPolicyByIDErsAncPolicy(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncPolicyUpdateAncPolicyByIDErsAncPolicy {
	request := isegosdk.RequestAncPolicyUpdateAncPolicyByIDErsAncPolicy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".actions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".actions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".actions")))) {
		request.Actions = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
