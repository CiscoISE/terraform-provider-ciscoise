package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkAccessPolicySet() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Access - Policy Set.
  
  - Network Access Create a new policy set:
  
  
  
   Policy must include name , service identifier (either server sequence or allowed protocol) and a condition.
  
  
   Condition has hierarchical structure which define a set of condition for which policy could be match.
  
  
   Condition can be either reference to a stored Library condition, using model
  ConditionReference
  ,
  or, dynamically built conditions which are not stored in the conditions Library, using models
  ConditionAttributes, ConditionAndBlock, ConditionOrBlock
  .
  
  
  
  
  - Network Access Update a policy set.
  - Network Access Delete a policy set.`,

		CreateContext: resourceNetworkAccessPolicySetCreate,
		ReadContext:   resourceNetworkAccessPolicySetRead,
		UpdateContext: resourceNetworkAccessPolicySetUpdate,
		DeleteContext: resourceNetworkAccessPolicySetDelete,
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

						"condition": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_id": &schema.Schema{
										Description: `Dictionary attribute id (Optional), used for additional verification`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"attribute_name": &schema.Schema{
										Description: `Dictionary attribute name`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"attribute_value": &schema.Schema{
										Description: `<ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"children": &schema.Schema{
										Description: `In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"condition_type": &schema.Schema{
													Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"is_negate": &schema.Schema{
													Description: `Indicates whereas this condition is in negate mode`,
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"condition_type": &schema.Schema{
										Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"dates_range": &schema.Schema{
										Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_date": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"start_date": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"dates_range_exception": &schema.Schema{
										Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_date": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"start_date": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"description": &schema.Schema{
										Description: `Condition description`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"dictionary_name": &schema.Schema{
										Description: `Dictionary name`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"dictionary_value": &schema.Schema{
										Description: `Dictionary value`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"hours_range": &schema.Schema{
										Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_time": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"hours_range_exception": &schema.Schema{
										Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_time": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_negate": &schema.Schema{
										Description: `Indicates whereas this condition is in negate mode`,
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Condition name`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"operator": &schema.Schema{
										Description: `Equality operator`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"week_days": &schema.Schema{
										Description: `<p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"week_days_exception": &schema.Schema{
										Description: `<p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"default": &schema.Schema{
							Description: `Flag which indicates if this policy set is the default one`,
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `The description for the policy set`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"hit_counts": &schema.Schema{
							Description: `The amount of times the policy was matched`,
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier for the policy set`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"is_proxy": &schema.Schema{
							Description: `Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'`,
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"rank": &schema.Schema{
							Description: `The rank(priority) in relation to other policy set. Lower rank is higher priority.`,
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"service_name": &schema.Schema{
							Description: `Policy set service identifier - Allowed Protocols,Server Sequence..`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"state": &schema.Schema{
							Description: `The state that the policy set is in. A disabled policy set cannot be matched.`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkAccessPolicySetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySet(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.NetworkAccessPolicySet.GetNetworkAccessPolicySetByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.NetworkAccessPolicySet.GetNetworkAccessPolicySets()
		if response2 != nil && err == nil {
			items2 := getAllItemsNetworkAccessPolicySetGetNetworkAccessPolicySets(m, response2)
			item2, err := searchNetworkAccessPolicySetGetNetworkAccessPolicySets(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	resp1, restyResp1, err := client.NetworkAccessPolicySet.CreateNetworkAccessPolicySet(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkAccessPolicySet", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkAccessPolicySet", err))
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
	return diags
}

func resourceNetworkAccessPolicySetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetNetworkAccessPolicySets")

		response1, _, err := client.NetworkAccessPolicySet.GetNetworkAccessPolicySets()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessPolicySets", err,
				"Failure at GetNetworkAccessPolicySets, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsNetworkAccessPolicySetGetNetworkAccessPolicySets(m, response1)
		item1, err := searchNetworkAccessPolicySetGetNetworkAccessPolicySets(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetNetworkAccessPolicySets response", err,
				"Failure when searching item from GetNetworkAccessPolicySets, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessPolicySets search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessPolicySetByID")

		response2, _, err := client.NetworkAccessPolicySet.GetNetworkAccessPolicySetByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessPolicySetByID", err,
				"Failure at GetNetworkAccessPolicySetByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessPolicySetByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkAccessPolicySetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		getResp1, _, err := client.NetworkAccessPolicySet.GetNetworkAccessPolicySets()
		if err == nil && getResp1 != nil {
			items1 := getAllItemsNetworkAccessPolicySetGetNetworkAccessPolicySets(m, getResp1)
			item1, err := searchNetworkAccessPolicySetGetNetworkAccessPolicySets(m, items1, vvName, vvID)
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
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkAccessPolicySet.UpdateNetworkAccessPolicySetByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkAccessPolicySetByID", err, restyResp1.String(),
					"Failure at UpdateNetworkAccessPolicySetByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkAccessPolicySetByID", err,
				"Failure at UpdateNetworkAccessPolicySetByID, unexpected response", ""))
			return diags
		}
	}

	return resourceNetworkAccessPolicySetRead(ctx, d, m)
}

func resourceNetworkAccessPolicySetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {

		getResp1, _, err := client.NetworkAccessPolicySet.GetNetworkAccessPolicySets()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkAccessPolicySetGetNetworkAccessPolicySets(m, getResp1)
		item1, err := searchNetworkAccessPolicySetGetNetworkAccessPolicySets(m, items1, vvName, vvID)
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
		getResp, _, err := client.NetworkAccessPolicySet.GetNetworkAccessPolicySetByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NetworkAccessPolicySet.DeleteNetworkAccessPolicySetByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkAccessPolicySetByID", err, restyResp1.String(),
				"Failure at DeleteNetworkAccessPolicySetByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkAccessPolicySetByID", err,
			"Failure at DeleteNetworkAccessPolicySetByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySet(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySet {
	request := isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySet{}
	if v, ok := d.GetOkExists(key + ".condition"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition"))) {
		request.Condition = expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition(ctx, key+".condition.0", d)
	}
	if v, ok := d.GetOkExists(key + ".default"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default"))) {
		request.Default = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".hit_counts"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hit_counts"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hit_counts"))) {
		request.HitCounts = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_proxy"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_proxy"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_proxy"))) {
		request.IsProxy = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rank"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rank"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rank"))) {
		request.Rank = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".service_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".service_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".service_name"))) {
		request.ServiceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".state"); !isEmptyValue(reflect.ValueOf(d.Get(key+".state"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".state"))) {
		request.State = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition {
	request := isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_name"))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_id"))) {
		request.AttributeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_value"))) {
		request.AttributeValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_name"))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_value"))) {
		request.DictionaryValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".operator"); !isEmptyValue(reflect.ValueOf(d.Get(key+".operator"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".operator"))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".children"); !isEmptyValue(reflect.ValueOf(d.Get(key+".children"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".children"))) {
		request.Children = expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildrenArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".week_days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".week_days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".week_days"))) {
		request.WeekDays = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".week_days_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".week_days_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".week_days_exception"))) {
		request.WeekDaysException = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionLink {
	request := isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildren {
	request := []isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildren{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildren {
	request := isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildrenLink(ctx, key+".link.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildrenLink {
	request := isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionChildrenLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRange {
	request := isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRange{}
	if v, ok := d.GetOkExists(key + ".end_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_date"))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_date"))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRangeException {
	request := isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionDatesRangeException{}
	if v, ok := d.GetOkExists(key + ".end_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_date"))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_date"))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRange {
	request := isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRange{}
	if v, ok := d.GetOkExists(key + ".end_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_time"))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_time"))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRangeException {
	request := isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetConditionHoursRangeException{}
	if v, ok := d.GetOkExists(key + ".end_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_time"))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_time"))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetCreateNetworkAccessPolicySetLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetLink {
	request := isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID {
	request := isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID{}
	if v, ok := d.GetOkExists(key + ".condition"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition"))) {
		request.Condition = expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition(ctx, key+".condition.0", d)
	}
	if v, ok := d.GetOkExists(key + ".default"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default"))) {
		request.Default = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".hit_counts"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hit_counts"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hit_counts"))) {
		request.HitCounts = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_proxy"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_proxy"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_proxy"))) {
		request.IsProxy = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rank"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rank"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rank"))) {
		request.Rank = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".service_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".service_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".service_name"))) {
		request.ServiceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".state"); !isEmptyValue(reflect.ValueOf(d.Get(key+".state"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".state"))) {
		request.State = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition {
	request := isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_name"))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_id"))) {
		request.AttributeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_value"))) {
		request.AttributeValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_name"))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_value"))) {
		request.DictionaryValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".operator"); !isEmptyValue(reflect.ValueOf(d.Get(key+".operator"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".operator"))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".children"); !isEmptyValue(reflect.ValueOf(d.Get(key+".children"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".children"))) {
		request.Children = expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildrenArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".week_days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".week_days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".week_days"))) {
		request.WeekDays = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".week_days_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".week_days_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".week_days_exception"))) {
		request.WeekDaysException = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionLink {
	request := isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildren {
	request := []isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildren{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildren {
	request := isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildrenLink(ctx, key+".link.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildrenLink {
	request := isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionChildrenLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRange {
	request := isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRange{}
	if v, ok := d.GetOkExists(key + ".end_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_date"))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_date"))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRangeException {
	request := isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionDatesRangeException{}
	if v, ok := d.GetOkExists(key + ".end_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_date"))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_date"))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRange {
	request := isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRange{}
	if v, ok := d.GetOkExists(key + ".end_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_time"))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_time"))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRangeException {
	request := isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDConditionHoursRangeException{}
	if v, ok := d.GetOkExists(key + ".end_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_time"))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_time"))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDLink {
	request := isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNetworkAccessPolicySetGetNetworkAccessPolicySets(m interface{}, response *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySets) []isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponse {
	var respItems []isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNetworkAccessPolicySetGetNetworkAccessPolicySets(m interface{}, items []isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponse, name string, id string) (*isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByID
			getItem, _, err = client.NetworkAccessPolicySet.GetNetworkAccessPolicySetByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessPolicySetByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByID
			getItem, _, err = client.NetworkAccessPolicySet.GetNetworkAccessPolicySetByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessPolicySetByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
