package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationConditionsForAuthorizationRule() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceAdministrationConditionsForAuthorizationRuleRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"children": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_negate": &schema.Schema{
										Type:     schema.TypeBool,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"dates_range": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
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
							Type:     schema.TypeList,
							Computed: true,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"dictionary_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dictionary_value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hours_range": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
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
							Type:     schema.TypeList,
							Computed: true,
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
							Type:     schema.TypeBool,
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
						"operator": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"week_days": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"week_days_exception": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
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

func dataSourceDeviceAdministrationConditionsForAuthorizationRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminConditionsForAuthorizationRules")

		response1, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionsForAuthorizationRules()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminConditionsForAuthorizationRules", err,
				"Failure at GetDeviceAdminConditionsForAuthorizationRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminConditionsForAuthorizationRules response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItems(items *[]isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsLink(item.Link)
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["attribute_name"] = item.AttributeName
		respItem["attribute_id"] = item.AttributeID
		respItem["attribute_value"] = item.AttributeValue
		respItem["dictionary_name"] = item.DictionaryName
		respItem["dictionary_value"] = item.DictionaryValue
		respItem["operator"] = item.Operator
		respItem["children"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsChildren(item.Children)
		respItem["dates_range"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsDatesRange(item.DatesRange)
		respItem["dates_range_exception"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsDatesRangeException(item.DatesRangeException)
		respItem["hours_range"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsHoursRange(item.HoursRange)
		respItem["hours_range_exception"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsHoursRangeException(item.HoursRangeException)
		respItem["week_days"] = item.WeekDays
		respItem["week_days_exception"] = item.WeekDaysException
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsLink(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsChildren(items *[]isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsChildrenLink(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsDatesRange(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsDatesRangeException(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsHoursRange(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesItemsHoursRangeException(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsForAuthorizationRulesResponseHoursRangeException) []map[string]interface{} {
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
