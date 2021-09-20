package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessPolicySet() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Access - Policy Set.

- Get all network access policy sets.

- Network Access Get policy set attributes.
`,

		ReadContext: dataSourceNetworkAccessPolicySetRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Policy id`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"condition": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_id": &schema.Schema{
										Description: `Dictionary attribute id (Optional), used for additional verification`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"attribute_name": &schema.Schema{
										Description: `Dictionary attribute name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"attribute_value": &schema.Schema{
										Description: `<ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"children": &schema.Schema{
										Description: `In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"condition_type": &schema.Schema{
													Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"is_negate": &schema.Schema{
													Description: `Indicates whereas this condition is in negate mode`,
													// Type:        schema.TypeBool,
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
											},
										},
									},
									"condition_type": &schema.Schema{
										Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"dates_range": &schema.Schema{
										Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_date": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"dates_range_exception": &schema.Schema{
										Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_date": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"description": &schema.Schema{
										Description: `Condition description`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"dictionary_name": &schema.Schema{
										Description: `Dictionary name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"dictionary_value": &schema.Schema{
										Description: `Dictionary value`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"hours_range": &schema.Schema{
										Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_time": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"hours_range_exception": &schema.Schema{
										Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_time": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_negate": &schema.Schema{
										Description: `Indicates whereas this condition is in negate mode`,
										// Type:        schema.TypeBool,
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
										Description: `Condition name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"operator": &schema.Schema{
										Description: `Equality operator`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"week_days": &schema.Schema{
										Description: `<p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"week_days_exception": &schema.Schema{
										Description: `<p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>`,
										Type:        schema.TypeList,
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
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `The description for the policy set`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"hit_counts": &schema.Schema{
							Description: `The amount of times the policy was matched`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier for the policy set`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"is_proxy": &schema.Schema{
							Description: `Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'`,
							// Type:        schema.TypeBool,
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
							Description: `Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"rank": &schema.Schema{
							Description: `The rank(priority) in relation to other policy set. Lower rank is higher priority.`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"service_name": &schema.Schema{
							Description: `Policy set service identifier - Allowed Protocols,Server Sequence..`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"state": &schema.Schema{
							Description: `The state that the policy set is in. A disabled policy set cannot be matched.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"condition": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_id": &schema.Schema{
										Description: `Dictionary attribute id (Optional), used for additional verification`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"attribute_name": &schema.Schema{
										Description: `Dictionary attribute name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"attribute_value": &schema.Schema{
										Description: `<ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"children": &schema.Schema{
										Description: `In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"condition_type": &schema.Schema{
													Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"is_negate": &schema.Schema{
													Description: `Indicates whereas this condition is in negate mode`,
													// Type:        schema.TypeBool,
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
											},
										},
									},
									"condition_type": &schema.Schema{
										Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"dates_range": &schema.Schema{
										Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_date": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"dates_range_exception": &schema.Schema{
										Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_date": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"description": &schema.Schema{
										Description: `Condition description`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"dictionary_name": &schema.Schema{
										Description: `Dictionary name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"dictionary_value": &schema.Schema{
										Description: `Dictionary value`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"hours_range": &schema.Schema{
										Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_time": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"hours_range_exception": &schema.Schema{
										Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_time": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_negate": &schema.Schema{
										Description: `Indicates whereas this condition is in negate mode`,
										// Type:        schema.TypeBool,
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
										Description: `Condition name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"operator": &schema.Schema{
										Description: `Equality operator`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"week_days": &schema.Schema{
										Description: `<p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"week_days_exception": &schema.Schema{
										Description: `<p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>`,
										Type:        schema.TypeList,
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
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `The description for the policy set`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"hit_counts": &schema.Schema{
							Description: `The amount of times the policy was matched`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier for the policy set`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"is_proxy": &schema.Schema{
							Description: `Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'`,
							// Type:        schema.TypeBool,
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
							Description: `Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"rank": &schema.Schema{
							Description: `The rank(priority) in relation to other policy set. Lower rank is higher priority.`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"service_name": &schema.Schema{
							Description: `Policy set service identifier - Allowed Protocols,Server Sequence..`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"state": &schema.Schema{
							Description: `The state that the policy set is in. A disabled policy set cannot be matched.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkAccessPolicySetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNetworkAccessPolicySets")

		response1, _, err := client.NetworkAccessPolicySet.GetNetworkAccessPolicySets()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessPolicySets", err,
				"Failure at GetNetworkAccessPolicySets, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessPolicySets response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNetworkAccessPolicySetByID")
		vvID := vID.(string)

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
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItems(items *[]isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsCondition(item.Condition)
		respItem["default"] = boolPtrToString(item.Default)
		respItem["description"] = item.Description
		respItem["hit_counts"] = item.HitCounts
		respItem["id"] = item.ID
		respItem["is_proxy"] = boolPtrToString(item.IsProxy)
		respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsLink(item.Link)
		respItem["name"] = item.Name
		respItem["rank"] = item.Rank
		respItem["service_name"] = item.ServiceName
		respItem["state"] = item.State
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsCondition(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = boolPtrToString(item.IsNegate)
	respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionChildren(item.Children)
	respItem["dates_range"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionChildren(items *[]isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = boolPtrToString(item.IsNegate)
		respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionChildrenLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionChildrenLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildrenLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionDatesRange(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionDatesRange) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionDatesRangeException(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionDatesRangeException) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionHoursRange(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionHoursRange) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionHoursRangeException(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionHoursRangeException) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItem(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemCondition(item.Condition)
	respItem["default"] = boolPtrToString(item.Default)
	respItem["description"] = item.Description
	respItem["hit_counts"] = item.HitCounts
	respItem["id"] = item.ID
	respItem["is_proxy"] = boolPtrToString(item.IsProxy)
	respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemLink(item.Link)
	respItem["name"] = item.Name
	respItem["rank"] = item.Rank
	respItem["service_name"] = item.ServiceName
	respItem["state"] = item.State
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemCondition(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = boolPtrToString(item.IsNegate)
	respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionChildren(item.Children)
	respItem["dates_range"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionChildren(items *[]isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = boolPtrToString(item.IsNegate)
		respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionChildrenLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionChildrenLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionChildrenLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionDatesRange(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionDatesRange) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionDatesRangeException(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionDatesRangeException) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionHoursRange(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionHoursRange) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionHoursRangeException(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionHoursRangeException) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
