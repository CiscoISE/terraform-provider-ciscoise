package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationTimeDateConditions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Administration - Time/Date Conditions.

- Device Admin Returns a list of time and date conditions.
- Device Admin Returns a network condition.`,

		ReadContext: dataSourceDeviceAdministrationTimeDateConditionsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Condition id`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
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

func dataSourceDeviceAdministrationTimeDateConditionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminTimeConditions")

		response1, _, err := client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditions()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminTimeConditions", err,
				"Failure at GetDeviceAdminTimeConditions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminTimeConditions response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceAdminTimeConditionByID")
		vvID := vID.(string)

		response2, _, err := client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditionByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminTimeConditionByID", err,
				"Failure at GetDeviceAdminTimeConditionByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminTimeConditionByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItems(items *[]isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsLink(item.Link)
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["attribute_name"] = item.AttributeName
		respItem["attribute_id"] = item.AttributeID
		respItem["attribute_value"] = item.AttributeValue
		respItem["dictionary_name"] = item.DictionaryName
		respItem["dictionary_value"] = item.DictionaryValue
		respItem["operator"] = item.Operator
		respItem["children"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsChildren(item.Children)
		respItem["dates_range"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsDatesRange(item.DatesRange)
		respItem["dates_range_exception"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsDatesRangeException(item.DatesRangeException)
		respItem["hours_range"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsHoursRange(item.HoursRange)
		respItem["hours_range_exception"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsHoursRangeException(item.HoursRangeException)
		respItem["week_days"] = item.WeekDays
		respItem["week_days_exception"] = item.WeekDaysException
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsLink(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsChildren(items *[]isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsChildrenLink(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsDatesRange(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsDatesRangeException(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsHoursRange(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsItemsHoursRangeException(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponseHoursRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItem(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemLink(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemChildren(items *[]isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemChildrenLink(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemDatesRange(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemDatesRangeException(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemHoursRange(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItemHoursRangeException(item *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponseHoursRangeException) []map[string]interface{} {
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
