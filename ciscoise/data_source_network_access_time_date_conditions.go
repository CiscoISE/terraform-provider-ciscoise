package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessTimeDateConditions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkAccessTimeDateConditionsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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
						"name": &schema.Schema{
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

func dataSourceNetworkAccessTimeDateConditionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNetworkAccessTimeConditions")

		response1, _, err := client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditions()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessTimeConditions", err,
				"Failure at GetNetworkAccessTimeConditions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessTimeConditions response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNetworkAccessTimeConditionByID")
		vvID := vID.(string)

		response2, _, err := client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditionByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessTimeConditionByID", err,
				"Failure at GetNetworkAccessTimeConditionByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItem(&response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessTimeConditionByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsItems(items *isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = interfaceToSliceString(items.Response)
	respItem["version"] = items.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItem(item *isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dates_range"] = flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItemDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItemDatesRangeException(item.DatesRangeException)
	respItem["description"] = item.Description
	respItem["hours_range"] = flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItemHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItemHoursRangeException(item.HoursRangeException)
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItemDatesRange(item isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseDatesRange) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItemDatesRangeException(item isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseDatesRangeException) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItemHoursRange(item isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseHoursRange) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItemHoursRangeException(item isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponseHoursRangeException) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}
