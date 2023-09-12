package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkAccessNetworkCondition() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Access - Network Conditions.

- Network Access Creates network condition.

- Network Access Update network condition.

- Network Access Delete network condition.
`,

		CreateContext: resourceNetworkAccessNetworkConditionCreate,
		ReadContext:   resourceNetworkAccessNetworkConditionRead,
		UpdateContext: resourceNetworkAccessNetworkConditionUpdate,
		DeleteContext: resourceNetworkAccessNetworkConditionDelete,
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

						"condition_type": &schema.Schema{
							Description: `This field determines the content of the conditions field`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"conditions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cli_dnis_list": &schema.Schema{
										Description: `<p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"condition_type": &schema.Schema{
										Description: `This field determines the content of the conditions field`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_group_list": &schema.Schema{
										Description: `<p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_list": &schema.Schema{
										Description: `<p>This field should contain Device-Name,port-number. The device name must be the same as the name field in a Network Device object.<br> Line format - Device Name,Port</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_addr_list": &schema.Schema{
										Description: `<p>This field should contain IP-address-or-subnet,port number<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet,Port</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
									"mac_addr_list": &schema.Schema{
										Description: `<p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": &schema.Schema{
										Description: `Network Condition name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
							Description: `Network Condition name`,
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

						"condition_type": &schema.Schema{
							Description:      `This field determines the content of the conditions field`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"conditions": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cli_dnis_list": &schema.Schema{
										Description:      `<p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>`,
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"condition_type": &schema.Schema{
										Description:      `This field determines the content of the conditions field`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"description": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"device_group_list": &schema.Schema{
										Description:      `<p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>`,
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_list": &schema.Schema{
										Description:      `<p>This field should contain Device-Name,port-number. The device name must be the same as the name field in a Network Device object.<br> Line format - Device Name,Port</p>`,
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"ip_addr_list": &schema.Schema{
										Description:      `<p>This field should contain IP-address-or-subnet,port number<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet,Port</p>`,
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"mac_addr_list": &schema.Schema{
										Description:      `<p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>`,
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": &schema.Schema{
										Description:      `Network Condition name`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},

						"name": &schema.Schema{
							Description:      `Network Condition name`,
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

func resourceNetworkAccessNetworkConditionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessNetworkCondition create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	isEnableAutoImport := clientConfig.EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkCondition(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse2, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(vvID)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceNetworkAccessNetworkConditionRead(ctx, d, m)
			}
		} else {
			response2, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditions()
			if response2 != nil && err == nil {
				items2 := getAllItemsNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, response2)
				item2, err := searchNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, items2, vvName, vvID)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["id"] = item2.ID
					resourceMap["name"] = vvName
					d.SetId(joinResourceID(resourceMap))
					return resourceNetworkAccessNetworkConditionRead(ctx, d, m)
				}
			}
		}
	}
	resp1, restyResp1, err := client.NetworkAccessNetworkConditions.CreateNetworkAccessNetworkCondition(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkAccessNetworkCondition", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkAccessNetworkCondition", err))
		return diags
	}
	if vvID != resp1.Response.ID {
		vvID = resp1.Response.ID
	}
	if vvName != resp1.Response.Name {
		vvName = resp1.Response.Name
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceNetworkAccessNetworkConditionRead(ctx, d, m)
}

func resourceNetworkAccessNetworkConditionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessNetworkCondition read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
		log.Printf("[DEBUG] Selected method: GetNetworkAccessNetworkConditions")

		response1, restyResp1, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditions()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, response1)
		item1, err := searchNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessNetworkConditions search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessNetworkConditions search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessNetworkConditionByID")

		response2, restyResp2, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessNetworkConditionByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessNetworkConditionByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkAccessNetworkConditionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessNetworkCondition update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		getResp1, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditions()
		if err == nil && getResp1 != nil {
			items1 := getAllItemsNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, getResp1)
			item1, err := searchNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, items1, vvName, vvID)
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
		request1 := expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.NetworkAccessNetworkConditions.UpdateNetworkAccessNetworkConditionByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkAccessNetworkConditionByID", err, restyResp1.String(),
					"Failure at UpdateNetworkAccessNetworkConditionByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkAccessNetworkConditionByID", err,
				"Failure at UpdateNetworkAccessNetworkConditionByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceNetworkAccessNetworkConditionRead(ctx, d, m)
}

func resourceNetworkAccessNetworkConditionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessNetworkCondition delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {

		getResp1, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditions()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, getResp1)
		item1, err := searchNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, items1, vvName, vvID)
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
		getResp, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NetworkAccessNetworkConditions.DeleteNetworkAccessNetworkConditionByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkAccessNetworkConditionByID", err, restyResp1.String(),
				"Failure at DeleteNetworkAccessNetworkConditionByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkAccessNetworkConditionByID", err,
			"Failure at DeleteNetworkAccessNetworkConditionByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition {
	request := isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_type")))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".conditions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".conditions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".conditions")))) {
		request.Conditions = expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionConditionsArray(ctx, key+".conditions", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionLink {
	request := isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionConditionsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionConditions {
	request := []isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionConditions{}
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
		i := expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionConditions(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionConditions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionConditions {
	request := isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionConditions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_dnis_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_dnis_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_dnis_list")))) {
		request.CliDnisList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_type")))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_addr_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_addr_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_addr_list")))) {
		request.IPAddrList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_addr_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_addr_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_addr_list")))) {
		request.MacAddrList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_group_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_group_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_group_list")))) {
		request.DeviceGroupList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_list")))) {
		request.DeviceList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionConditionsLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionConditionsLink {
	request := isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionConditionsLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID {
	request := isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_type")))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".conditions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".conditions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".conditions")))) {
		request.Conditions = expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDConditionsArray(ctx, key+".conditions", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDLink {
	request := isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDConditionsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDConditions {
	request := []isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDConditions{}
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
		i := expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDConditions(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDConditions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDConditions {
	request := isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDConditions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_dnis_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_dnis_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_dnis_list")))) {
		request.CliDnisList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_type")))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_addr_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_addr_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_addr_list")))) {
		request.IPAddrList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_addr_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_addr_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_addr_list")))) {
		request.MacAddrList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_group_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_group_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_group_list")))) {
		request.DeviceGroupList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_list")))) {
		request.DeviceList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDConditionsLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDConditionsLink {
	request := isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDConditionsLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m interface{}, response *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions) []isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponse {
	var respItems []isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m interface{}, items []isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponse, name string, id string) (*isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponse, error) {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var err error
	var foundItem *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByID
			getItem, _, err = client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessNetworkConditionByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByID
			getItem, _, err = client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessNetworkConditionByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
