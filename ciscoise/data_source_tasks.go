package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTasks() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on tasks.

get all task status
Monitor task status`,

		ReadContext: dataSourceTasksRead,
		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Description: `taskId path parameter. The id of the task executed before`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"detail_status": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"execution_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"fail_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"module_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"resources_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"success_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"detail_status": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"execution_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"fail_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"module_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"resources_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"success_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTasksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vTaskID, okTaskID := d.GetOk("task_id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okTaskID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTaskStatus")

		response1, _, err := client.Tasks.GetTaskStatus()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskStatus", err,
				"Failure at GetTaskStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenTasksGetTaskStatusItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTaskStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetTaskStatusByID")
		vvTaskID := vTaskID.(string)

		response2, _, err := client.Tasks.GetTaskStatusByID(vvTaskID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskStatusByID", err,
				"Failure at GetTaskStatusByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenTasksGetTaskStatusByIDItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTaskStatusByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTasksGetTaskStatusItems(items *[]isegosdk.ResponseTasksGetTaskStatus) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["execution_status"] = item.ExecutionStatus
		respItem["module_type"] = item.ModuleType
		respItem["start_time"] = item.StartTime
		respItem["resources_count"] = item.ResourcesCount
		respItem["success_count"] = item.SuccessCount
		respItem["fail_count"] = item.FailCount
		respItem["detail_status"] = responseInterfaceToSliceString(item.DetailStatus)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTasksGetTaskStatusByIDItem(item *isegosdk.ResponseTasksGetTaskStatusByID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["execution_status"] = item.ExecutionStatus
	respItem["module_type"] = item.ModuleType
	respItem["start_time"] = item.StartTime
	respItem["resources_count"] = item.ResourcesCount
	respItem["success_count"] = item.SuccessCount
	respItem["fail_count"] = item.FailCount
	respItem["detail_status"] = responseInterfaceToSliceString(item.DetailStatus)
	return []map[string]interface{}{
		respItem,
	}
}
