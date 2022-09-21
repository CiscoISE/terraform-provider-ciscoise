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

func resourceRadiusServerSequence() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on RADIUSServerSequence.

- This resource allows the client to update a RADIUS server sequence.

- This resource deletes a RADIUS server sequence.

- This resource creates a RADIUS server sequence.
`,

		CreateContext: resourceRadiusServerSequenceCreate,
		ReadContext:   resourceRadiusServerSequenceRead,
		UpdateContext: resourceRadiusServerSequenceUpdate,
		DeleteContext: resourceRadiusServerSequenceDelete,
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

						"before_accept_attr_manipulators_list": &schema.Schema{
							Description: `The beforeAcceptAttrManipulators is required only if useAttrSetBeforeAcc is true`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action": &schema.Schema{
										Description: `Allowed Values:
- ADD,
- UPDATE,
- REMOVE,
- REMOVEANY`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"attribute_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"changed_val": &schema.Schema{
										Description: `The changedVal is required only if the action equals to 'UPDATE'`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"dictionary_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"on_request_attr_manipulator_list": &schema.Schema{
							Description: `The onRequestAttrManipulators is required only if useAttrSetOnRequest is true`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action": &schema.Schema{
										Description: `Allowed Values:
- ADD,
- UPDATE,
- REMOVE,
- REMOVEANY`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"attribute_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"changed_val": &schema.Schema{
										Description: `The changedVal is required only if the action equals to 'UPDATE'`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"dictionary_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"radius_server_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"continue_authorz_policy": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
						"local_accounting": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_separator": &schema.Schema{
							Description: `The prefixSeparator is required only if stripPrefix is true. The maximum length is 1 character`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"remote_accounting": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"strip_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"strip_suffix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"suffix_separator": &schema.Schema{
							Description: `The suffixSeparator is required only if stripSuffix is true. The maximum length is 1 character`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"use_attr_set_before_acc": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"use_attr_set_on_request": &schema.Schema{
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

						"before_accept_attr_manipulators_list": &schema.Schema{
							Description:      `The beforeAcceptAttrManipulators is required only if useAttrSetBeforeAcc is true`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action": &schema.Schema{
										Description: `Allowed Values:
		- ADD,
		- UPDATE,
		- REMOVE,
		- REMOVEANY`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"attribute_name": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"changed_val": &schema.Schema{
										Description:      `The changedVal is required only if the action equals to 'UPDATE'`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"dictionary_name": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"value": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"on_request_attr_manipulator_list": &schema.Schema{
							Description:      `The onRequestAttrManipulators is required only if useAttrSetOnRequest is true`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action": &schema.Schema{
										Description: `Allowed Values:
		- ADD,
		- UPDATE,
		- REMOVE,
		- REMOVEANY`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"attribute_name": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"changed_val": &schema.Schema{
										Description:      `The changedVal is required only if the action equals to 'UPDATE'`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"dictionary_name": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"value": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"radius_server_list": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"continue_authorz_policy": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
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
						"local_accounting": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"name": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"prefix_separator": &schema.Schema{
							Description:      `The prefixSeparator is required only if stripPrefix is true. The maximum length is 1 character`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"remote_accounting": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"strip_prefix": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"strip_suffix": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"suffix_separator": &schema.Schema{
							Description:      `The suffixSeparator is required only if stripSuffix is true. The maximum length is 1 character`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"use_attr_set_before_acc": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"use_attr_set_on_request": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func resourceRadiusServerSequenceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning RadiusServerSequence create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	isEnableAutoImport := clientConfig.EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestRadiusServerSequenceCreateRadiusServerSequence(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vName, _ := resourceItem["name"]
	vvID := interfaceToString(vID)
	vvName := interfaceToString(vName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse2, _, err := client.RadiusServerSequence.GetRadiusServerSequenceByID(vvID)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceRadiusServerSequenceRead(ctx, d, m)
			}
		} else {
			queryParams2 := isegosdk.GetRadiusServerSequenceQueryParams{}

			response2, _, err := client.RadiusServerSequence.GetRadiusServerSequence(&queryParams2)
			if response2 != nil && err == nil {
				items2 := getAllItemsRadiusServerSequenceGetRadiusServerSequence(m, response2, &queryParams2)
				item2, err := searchRadiusServerSequenceGetRadiusServerSequence(m, items2, vvName, vvID)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["id"] = item2.ID
					resourceMap["name"] = vvName
					d.SetId(joinResourceID(resourceMap))
					return resourceRadiusServerSequenceRead(ctx, d, m)
				}
			}
		}
	}
	restyResp1, err := client.RadiusServerSequence.CreateRadiusServerSequence(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateRadiusServerSequence", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateRadiusServerSequence", err))
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
	return resourceRadiusServerSequenceRead(ctx, d, m)
}

func resourceRadiusServerSequenceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning RadiusServerSequence read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
		vvID := vID
		vvName := vName
		log.Printf("[DEBUG] Selected method: GetRadiusServerSequence")
		queryParams1 := isegosdk.GetRadiusServerSequenceQueryParams{}

		response1, restyResp1, err := client.RadiusServerSequence.GetRadiusServerSequence(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsRadiusServerSequenceGetRadiusServerSequence(m, response1, &queryParams1)
		item1, err := searchRadiusServerSequenceGetRadiusServerSequence(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenRadiusServerSequenceGetRadiusServerSequenceByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRadiusServerSequence search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRadiusServerSequence search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetRadiusServerSequenceByID")
		vvID := vID

		response2, restyResp2, err := client.RadiusServerSequence.GetRadiusServerSequenceByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenRadiusServerSequenceGetRadiusServerSequenceByIDItem(response2.RadiusServerSequence)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRadiusServerSequenceByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRadiusServerSequenceByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceRadiusServerSequenceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning RadiusServerSequence update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
		queryParams1 := isegosdk.GetRadiusServerSequenceQueryParams{}

		getResp1, _, err := client.RadiusServerSequence.GetRadiusServerSequence(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsRadiusServerSequenceGetRadiusServerSequence(m, getResp1, &queryParams1)
			item1, err := searchRadiusServerSequenceGetRadiusServerSequence(m, items1, vName, vID)
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
		request1 := expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.RadiusServerSequence.UpdateRadiusServerSequenceByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateRadiusServerSequenceByID", err, restyResp1.String(),
					"Failure at UpdateRadiusServerSequenceByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateRadiusServerSequenceByID", err,
				"Failure at UpdateRadiusServerSequenceByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceRadiusServerSequenceRead(ctx, d, m)
}

func resourceRadiusServerSequenceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning RadiusServerSequence delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
		queryParams1 := isegosdk.GetRadiusServerSequenceQueryParams{}

		getResp1, _, err := client.RadiusServerSequence.GetRadiusServerSequence(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsRadiusServerSequenceGetRadiusServerSequence(m, getResp1, &queryParams1)
		item1, err := searchRadiusServerSequenceGetRadiusServerSequence(m, items1, vName, vID)
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
		getResp, _, err := client.RadiusServerSequence.GetRadiusServerSequenceByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.RadiusServerSequence.DeleteRadiusServerSequenceByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteRadiusServerSequenceByID", err, restyResp1.String(),
				"Failure at DeleteRadiusServerSequenceByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteRadiusServerSequenceByID", err,
			"Failure at DeleteRadiusServerSequenceByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestRadiusServerSequenceCreateRadiusServerSequence(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequence {
	request := isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequence{}
	request.RadiusServerSequence = expandRequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequence(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequence(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequence {
	request := isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequence{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".strip_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".strip_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".strip_prefix")))) {
		request.StripPrefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".strip_suffix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".strip_suffix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".strip_suffix")))) {
		request.StripSuffix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prefix_separator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prefix_separator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prefix_separator")))) {
		request.PrefixSeparator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".suffix_separator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".suffix_separator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".suffix_separator")))) {
		request.SuffixSeparator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_accounting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_accounting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_accounting")))) {
		request.RemoteAccounting = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_accounting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_accounting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_accounting")))) {
		request.LocalAccounting = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_attr_set_on_request")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_attr_set_on_request")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_attr_set_on_request")))) {
		request.UseAttrSetOnRequest = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_attr_set_before_acc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_attr_set_before_acc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_attr_set_before_acc")))) {
		request.UseAttrSetBeforeAcc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".continue_authorz_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".continue_authorz_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".continue_authorz_policy")))) {
		request.ContinueAuthorzPolicy = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radius_server_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radius_server_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radius_server_list")))) {
		request.RadiusServerList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".on_request_attr_manipulator_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".on_request_attr_manipulator_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".on_request_attr_manipulator_list")))) {
		request.OnRequestAttrManipulatorList = expandRequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceOnRequestAttrManipulatorListArray(ctx, key+".on_request_attr_manipulator_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".before_accept_attr_manipulators_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".before_accept_attr_manipulators_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".before_accept_attr_manipulators_list")))) {
		request.BeforeAcceptAttrManipulatorsList = expandRequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceBeforeAcceptAttrManipulatorsListArray(ctx, key+".before_accept_attr_manipulators_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceOnRequestAttrManipulatorListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceOnRequestAttrManipulatorList {
	request := []isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceOnRequestAttrManipulatorList{}
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
		i := expandRequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceOnRequestAttrManipulatorList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceOnRequestAttrManipulatorList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceOnRequestAttrManipulatorList {
	request := isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceOnRequestAttrManipulatorList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".action")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".action")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".action")))) {
		request.Action = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_name")))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_name")))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".changed_val")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".changed_val")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".changed_val")))) {
		request.ChangedVal = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceBeforeAcceptAttrManipulatorsListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceBeforeAcceptAttrManipulatorsList {
	request := []isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceBeforeAcceptAttrManipulatorsList{}
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
		i := expandRequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceBeforeAcceptAttrManipulatorsList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceBeforeAcceptAttrManipulatorsList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceBeforeAcceptAttrManipulatorsList {
	request := isegosdk.RequestRadiusServerSequenceCreateRadiusServerSequenceRadiusServerSequenceBeforeAcceptAttrManipulatorsList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".action")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".action")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".action")))) {
		request.Action = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_name")))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_name")))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".changed_val")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".changed_val")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".changed_val")))) {
		request.ChangedVal = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByID {
	request := isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByID{}
	request.RadiusServerSequence = expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequence(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequence(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequence {
	request := isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequence{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".strip_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".strip_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".strip_prefix")))) {
		request.StripPrefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".strip_suffix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".strip_suffix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".strip_suffix")))) {
		request.StripSuffix = interfaceToBoolPtr(v)
	}
	vStripPrefix, okStripPrefix := d.GetOk(fixKeyAccess(key + ".strip_prefix"))
	vvStripPrefix := interfaceToBoolPtr(vStripPrefix)
	if okStripPrefix && vvStripPrefix != nil && *vvStripPrefix {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".prefix_separator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prefix_separator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prefix_separator")))) {
			request.PrefixSeparator = interfaceToString(v)
		}
	}
	vStripSuffix, okStripSuffix := d.GetOk(fixKeyAccess(key + ".strip_suffix"))
	vvStripSuffix := interfaceToBoolPtr(vStripSuffix)
	if okStripSuffix && vvStripSuffix != nil && *vvStripSuffix {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".suffix_separator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".suffix_separator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".suffix_separator")))) {
			request.SuffixSeparator = interfaceToString(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_accounting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_accounting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_accounting")))) {
		request.RemoteAccounting = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_accounting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_accounting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_accounting")))) {
		request.LocalAccounting = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_attr_set_on_request")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_attr_set_on_request")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_attr_set_on_request")))) {
		request.UseAttrSetOnRequest = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_attr_set_before_acc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_attr_set_before_acc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_attr_set_before_acc")))) {
		request.UseAttrSetBeforeAcc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".continue_authorz_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".continue_authorz_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".continue_authorz_policy")))) {
		request.ContinueAuthorzPolicy = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radius_server_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radius_server_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radius_server_list")))) {
		request.RadiusServerList = interfaceToSliceString(v)
	}

	vUseAttrSetBeforeAcc, okUseAttrSetBeforeAcc := d.GetOk(fixKeyAccess(key + ".use_attr_set_before_acc"))
	vvUseAttrSetBeforeAcc := interfaceToBoolPtr(vUseAttrSetBeforeAcc)
	if okUseAttrSetBeforeAcc && vvUseAttrSetBeforeAcc != nil && *vvUseAttrSetBeforeAcc {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".on_request_attr_manipulator_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".on_request_attr_manipulator_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".on_request_attr_manipulator_list")))) {
			request.OnRequestAttrManipulatorList = expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorListArray(ctx, fixKeyAccess(key+".on_request_attr_manipulator_list"), d)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".before_accept_attr_manipulators_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".before_accept_attr_manipulators_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".before_accept_attr_manipulators_list")))) {
			request.BeforeAcceptAttrManipulatorsList = expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsListArray(ctx, fixKeyAccess(key+".before_accept_attr_manipulators_list"), d)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList {
	request := []isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList{}
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
		i := expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList {
	request := isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceOnRequestAttrManipulatorList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".action")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".action")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".action")))) {
		request.Action = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_name")))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_name")))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	vAction, okAction := d.GetOk(fixKeyAccess(key + ".action"))
	vvAction := interfaceToString(vAction)
	if okAction && vvAction == "UPDATE" {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".changed_val")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".changed_val")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".changed_val")))) {
			request.ChangedVal = interfaceToString(v)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList {
	request := []isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList{}
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
		i := expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList {
	request := isegosdk.RequestRadiusServerSequenceUpdateRadiusServerSequenceByIDRadiusServerSequenceBeforeAcceptAttrManipulatorsList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".action")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".action")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".action")))) {
		request.Action = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_name")))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_name")))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	vAction, okAction := d.GetOk(fixKeyAccess(key + ".action"))
	vvAction := interfaceToString(vAction)
	if okAction && vvAction == "UPDATE" {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".changed_val")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".changed_val")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".changed_val")))) {
			request.ChangedVal = interfaceToString(v)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsRadiusServerSequenceGetRadiusServerSequence(m interface{}, response *isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequence, queryParams *isegosdk.GetRadiusServerSequenceQueryParams) []isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultResources {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var respItems []isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultResources
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
			response, _, err = client.RadiusServerSequence.GetRadiusServerSequence(queryParams)
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

func searchRadiusServerSequenceGetRadiusServerSequence(m interface{}, items []isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceSearchResultResources, name string, id string) (*isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequence, error) {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var err error
	var foundItem *isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceByIDRadiusServerSequence
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceByID
			getItem, _, err = client.RadiusServerSequence.GetRadiusServerSequenceByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetRadiusServerSequenceByID")
			}
			foundItem = getItem.RadiusServerSequence
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseRadiusServerSequenceGetRadiusServerSequenceByID
			getItem, _, err = client.RadiusServerSequence.GetRadiusServerSequenceByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetRadiusServerSequenceByID")
			}
			foundItem = getItem.RadiusServerSequence
			return foundItem, err
		}
	}
	return foundItem, err
}
