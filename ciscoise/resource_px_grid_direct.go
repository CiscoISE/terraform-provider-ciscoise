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

func resourcePxGridDirect() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on pxGrid Direct.

- pxGrid Direct Configure connectorconfig information.

- pxGrid Direct update Configure connectorConfig information based on ConnectorName.

- pxGrid Direct Delete Configure connectorConfig information based on ConnectorName.
`,

		CreateContext: resourcePxGridDirectCreate,
		ReadContext:   resourcePxGridDirectRead,
		UpdateContext: resourcePxGridDirectUpdate,
		DeleteContext: resourcePxGridDirectDelete,
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
						"connector": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:     schema.TypeMap,
										Computed: true,
									},
									"attributes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute_mapping": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dictionary_attribute": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"include_in_dictionary": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"json_attribute": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"correlation_identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"top_level_object": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"unique_identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"version_identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"connector_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connector_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"deltasync_schedule": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"interval": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"interval_unit": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"fullsync_schedule": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"interval": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"interval_unit": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"protocol": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"skip_certificate_validations": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"bulk_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"incremental_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"password": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"user_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_properties": &schema.Schema{
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"attributes": &schema.Schema{
							Description:      `connectorName`,
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"connector_name": &schema.Schema{
							Description:      `connectorName`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"connector_type": &schema.Schema{
							Description:      `connector Type list`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"deltasync_schedule": &schema.Schema{
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"description": &schema.Schema{
							Description:      `description`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"enabled": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"fullsync_schedule": &schema.Schema{
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"protocol": &schema.Schema{
							Description:      `protocol`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"response": &schema.Schema{
							Description: `ConnectorConfig information format`,
							Type:        schema.TypeMap,
							Computed:    true,
						},
						"skip_certificate_validations": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"url": &schema.Schema{
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
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

func resourcePxGridDirectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := m.(ClientConfig).EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestPxGridDirectCreateConnectorConfig(ctx, "parameters.0", d)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vConnectorName, okConnectorName := resourceItem["connector_name"]
	vvConnectorName := interfaceToString(vConnectorName)
	if isEnableAutoImport {
		if okConnectorName && vvConnectorName != "" {
			getResponse2, _, err := client.PxGridDirect.GetConnectorConfigByConnectorName(vvConnectorName)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["connector_name"] = vvConnectorName
				d.SetId(joinResourceID(resourceMap))
				return resourcePxGridDirectRead(ctx, d, m)
			}
		} else {
			response2, _, err := client.PxGridDirect.GetConnectorConfig()
			if response2 != nil && err == nil {
				items2 := getAllItemsPxGridDirectGetConnectorConfig(m, response2)
				item2, err := searchPxGridDirectGetConnectorConfig(m, items2, vvConnectorName)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["connector_name"] = vvConnectorName
					d.SetId(joinResourceID(resourceMap))
					return resourcePxGridDirectRead(ctx, d, m)
				}
			}
		}
	}
	resp1, err := client.PxGridDirect.CreateConnectorConfig(request1)
	if err != nil || resp1 == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateConnectorConfig", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["connector_name"] = vvConnectorName
	d.SetId(joinResourceID(resourceMap))
	return resourcePxGridDirectRead(ctx, d, m)
}

func resourcePxGridDirectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vConnectorName, okConnectorName := resourceMap["connector_name"]
	vvConnectorName := vConnectorName

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okConnectorName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetConnectorConfig")

		response1, restyResp1, err := client.PxGridDirect.GetConnectorConfig()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsPxGridDirectGetConnectorConfig(m, response1)
		item1, err := searchPxGridDirectGetConnectorConfig(m, items1, vvConnectorName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponse(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConnectorConfig search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConnectorConfig search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetConnectorConfigByConnectorName")

		response2, restyResp2, err := client.PxGridDirect.GetConnectorConfigByConnectorName(vvConnectorName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponse(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConnectorConfigByConnectorName response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConnectorConfigByConnectorName response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourcePxGridDirectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vConnectorName, _ := resourceMap["connector_name"]
	vvConnectorName := vConnectorName

	if d.HasChange("parameters") {

		log.Printf("[DEBUG] Name used for update operation %s", vvConnectorName)

		request1 := expandRequestPxGridDirectUpdateConnectorConfigByConnectorName(ctx, "parameters.0", d)

		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

		response1, err := client.PxGridDirect.UpdateConnectorConfigByConnectorName(vvConnectorName, request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(

				"Failure when executing UpdateConnectorConfigByConnectorName", err,

				"Failure at UpdateConnectorConfigByConnectorName, unexpected response", ""))

			return diags

		}

	}

	return resourcePxGridDirectRead(ctx, d, m)
}

func resourcePxGridDirectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vConnectorName, okConnectorName := resourceMap["connector_name"]
	vvConnectorName := vConnectorName

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okConnectorName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.PxGridDirect.GetConnectorConfig()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsPxGridDirectGetConnectorConfig(m, getResp1)
		item1, err := searchPxGridDirectGetConnectorConfig(m, items1, vvConnectorName)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vConnectorName != item1.Connector.ConnectorName {
			vvConnectorName = item1.Connector.ConnectorName
		} else {
			vvConnectorName = vConnectorName
		}
	}
	if selectedMethod == 2 {
		getResp, _, err := client.PxGridDirect.GetConnectorConfigByConnectorName(vvConnectorName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, err := client.PxGridDirect.DeleteConnectorConfigByConnectorName(vvConnectorName)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteConnectorConfigByConnectorName", err,
			"Failure at DeleteConnectorConfigByConnectorName, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestPxGridDirectCreateConnectorConfig(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectCreateConnectorConfig {
	request := isegosdk.RequestPxGridDirectCreateConnectorConfig{}
	request.Connector = expandRequestPxGridDirectCreateConnectorConfigConnector(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectCreateConnectorConfigConnector(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectCreateConnectorConfigConnector {
	request := isegosdk.RequestPxGridDirectCreateConnectorConfigConnector{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_properties")))) {
		request.AdditionalProperties = expandRequestPxGridDirectCreateConnectorConfigConnectorAdditionalProperties(ctx, key+".additional_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = expandRequestPxGridDirectCreateConnectorConfigConnectorAttributes(ctx, key+".attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_name")))) {
		request.ConnectorName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deltasync_schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deltasync_schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deltasync_schedule")))) {
		request.DeltasyncSchedule = expandRequestPxGridDirectCreateConnectorConfigConnectorDeltasyncSchedule(ctx, key+".deltasync_schedule.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enabled")))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fullsync_schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fullsync_schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fullsync_schedule")))) {
		request.FullsyncSchedule = expandRequestPxGridDirectCreateConnectorConfigConnectorFullsyncSchedule(ctx, key+".fullsync_schedule.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".skip_certificate_validations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".skip_certificate_validations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".skip_certificate_validations")))) {
		request.SkipCertificateValidations = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = expandRequestPxGridDirectCreateConnectorConfigConnectorURL(ctx, key+".url.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectCreateConnectorConfigConnectorAdditionalProperties(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorAdditionalProperties {
	var request isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorAdditionalProperties
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectCreateConnectorConfigConnectorAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorAttributes {
	request := isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_mapping")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_mapping")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_mapping")))) {
		request.AttributeMapping = expandRequestPxGridDirectCreateConnectorConfigConnectorAttributesAttributeMappingArray(ctx, key+".attribute_mapping", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".correlation_identifier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".correlation_identifier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".correlation_identifier")))) {
		request.CorrelationIDentifier = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".top_level_object")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".top_level_object")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".top_level_object")))) {
		request.TopLevelObject = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".unique_identifier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".unique_identifier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".unique_identifier")))) {
		request.UniqueIDentifier = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version_identifier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version_identifier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version_identifier")))) {
		request.VersionIDentifier = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectCreateConnectorConfigConnectorAttributesAttributeMappingArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorAttributesAttributeMapping {
	request := []isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorAttributesAttributeMapping{}
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
		i := expandRequestPxGridDirectCreateConnectorConfigConnectorAttributesAttributeMapping(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectCreateConnectorConfigConnectorAttributesAttributeMapping(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorAttributesAttributeMapping {
	request := isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorAttributesAttributeMapping{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_attribute")))) {
		request.DictionaryAttribute = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_in_dictionary")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_in_dictionary")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_in_dictionary")))) {
		request.IncludeInDictionary = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".json_attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".json_attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".json_attribute")))) {
		request.JSONAttribute = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectCreateConnectorConfigConnectorDeltasyncSchedule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorDeltasyncSchedule {
	request := isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorDeltasyncSchedule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval")))) {
		request.Interval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval_unit")))) {
		request.IntervalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectCreateConnectorConfigConnectorFullsyncSchedule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorFullsyncSchedule {
	request := isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorFullsyncSchedule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval")))) {
		request.Interval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval_unit")))) {
		request.IntervalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectCreateConnectorConfigConnectorURL(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorURL {
	request := isegosdk.RequestPxGridDirectCreateConnectorConfigConnectorURL{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_type")))) {
		request.AuthenticationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bulk_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bulk_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bulk_url")))) {
		request.BulkURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".incremental_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".incremental_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".incremental_url")))) {
		request.IncrementalURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectUpdateConnectorConfigByConnectorName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorName {
	request := isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorName{}
	request.Connector = expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnector(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnector(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnector {
	request := isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnector{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_properties")))) {
		request.AdditionalProperties = expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAdditionalProperties(ctx, key+".additional_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributes(ctx, key+".attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_name")))) {
		request.ConnectorName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deltasync_schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deltasync_schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deltasync_schedule")))) {
		request.DeltasyncSchedule = expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorDeltasyncSchedule(ctx, key+".deltasync_schedule.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enabled")))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fullsync_schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fullsync_schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fullsync_schedule")))) {
		request.FullsyncSchedule = expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorFullsyncSchedule(ctx, key+".fullsync_schedule.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".skip_certificate_validations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".skip_certificate_validations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".skip_certificate_validations")))) {
		request.SkipCertificateValidations = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorURL(ctx, key+".url.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAdditionalProperties(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAdditionalProperties {
	var request isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAdditionalProperties
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributes {
	request := isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_mapping")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_mapping")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_mapping")))) {
		request.AttributeMapping = expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMappingArray(ctx, key+".attribute_mapping", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".correlation_identifier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".correlation_identifier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".correlation_identifier")))) {
		request.CorrelationIDentifier = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".top_level_object")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".top_level_object")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".top_level_object")))) {
		request.TopLevelObject = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".unique_identifier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".unique_identifier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".unique_identifier")))) {
		request.UniqueIDentifier = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version_identifier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version_identifier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version_identifier")))) {
		request.VersionIDentifier = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMappingArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMapping {
	request := []isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMapping{}
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
		i := expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMapping(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMapping(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMapping {
	request := isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorAttributesAttributeMapping{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_attribute")))) {
		request.DictionaryAttribute = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_in_dictionary")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_in_dictionary")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_in_dictionary")))) {
		request.IncludeInDictionary = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".json_attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".json_attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".json_attribute")))) {
		request.JSONAttribute = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorDeltasyncSchedule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorDeltasyncSchedule {
	request := isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorDeltasyncSchedule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval")))) {
		request.Interval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval_unit")))) {
		request.IntervalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorFullsyncSchedule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorFullsyncSchedule {
	request := isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorFullsyncSchedule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval")))) {
		request.Interval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval_unit")))) {
		request.IntervalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorURL(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorURL {
	request := isegosdk.RequestPxGridDirectUpdateConnectorConfigByConnectorNameConnectorURL{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_type")))) {
		request.AuthenticationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bulk_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bulk_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bulk_url")))) {
		request.BulkURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".incremental_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".incremental_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".incremental_url")))) {
		request.IncrementalURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsPxGridDirectGetConnectorConfig(m interface{}, response *isegosdk.ResponsePxGridDirectGetConnectorConfig) []isegosdk.ResponsePxGridDirectGetConnectorConfigResponse {
	var respItems []isegosdk.ResponsePxGridDirectGetConnectorConfigResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchPxGridDirectGetConnectorConfig(m interface{}, items []isegosdk.ResponsePxGridDirectGetConnectorConfigResponse, connectorName string) (*isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorNameResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorNameResponse
	for _, item := range items {
		if connectorName != "" && item.Connector.ConnectorName == connectorName {
			var getItem *isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorName
			getItem, _, err = client.PxGridDirect.GetConnectorConfigByConnectorName(connectorName)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetConnectorConfigByConnectorName")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
