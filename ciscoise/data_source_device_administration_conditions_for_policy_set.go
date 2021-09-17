package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationConditionsForPolicySet() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Administration - Conditions.

- Device Admin Returns list of library conditions for policy sets.`,

		ReadContext: dataSourceDeviceAdministrationConditionsForPolicySetRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
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
										Type:        schema.TypeBool,
										Computed:    true,
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
							Type:        schema.TypeBool,
							Computed:    true,
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
		},
	}
}

func dataSourceDeviceAdministrationConditionsForPolicySetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminConditionsForPolicySets")

		response1, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionsForPolicySets()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminConditionsForPolicySets", err,
				"Failure at GetDeviceAdminConditionsForPolicySets, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminConditionsForPolicySets response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItems(items *[]isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsLink(item.Link)
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["attribute_name"] = item.AttributeName
		respItem["attribute_id"] = item.AttributeID
		respItem["attribute_value"] = item.AttributeValue
		respItem["dictionary_name"] = item.DictionaryName
		respItem["dictionary_value"] = item.DictionaryValue
		respItem["operator"] = item.Operator
		respItem["children"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsChildren(item.Children)
		respItem["dates_range"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsDatesRange(item.DatesRange)
		respItem["dates_range_exception"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsDatesRangeException(item.DatesRangeException)
		respItem["hours_range"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsHoursRange(item.HoursRange)
		respItem["hours_range_exception"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsHoursRangeException(item.HoursRangeException)
		respItem["week_days"] = item.WeekDays
		respItem["week_days_exception"] = item.WeekDaysException
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsLink(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsChildren(items *[]isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsChildrenLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsChildrenLink(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsDatesRange(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsDatesRangeException(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsHoursRange(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsItemsHoursRangeException(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForPolicySetsResponseHoursRangeException) []map[string]interface{} {
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
