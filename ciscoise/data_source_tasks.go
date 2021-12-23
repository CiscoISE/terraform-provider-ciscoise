package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTasks() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on tasks.

- get all task status

- Monitor task status
`,

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
							// Replaced List to String
							Type:     schema.TypeString,
							Computed: true,
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
							// Replaced List to String
							Type:     schema.TypeString,
							Computed: true,
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
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllTaskStatus")

		response1, restyResp1, err := client.Tasks.GetAllTaskStatus()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAllTaskStatus", err,
				"Failure at GetAllTaskStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenTasksGetAllTaskStatusItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllTaskStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetTaskStatus")
		vvTaskID := vTaskID.(string)

		response2, restyResp2, err := client.Tasks.GetTaskStatus(vvTaskID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskStatus", err,
				"Failure at GetTaskStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenTasksGetTaskStatusItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTaskStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTasksGetAllTaskStatusItems(items *isegosdk.ResponseTasksGetAllTaskStatus) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["detail_status"] = flattenTasksGetAllTaskStatusItemsDetailStatus(item.DetailStatus)
		respItem["execution_status"] = item.ExecutionStatus
		respItem["fail_count"] = item.FailCount
		respItem["id"] = item.ID
		respItem["module_type"] = item.ModuleType
		respItem["resources_count"] = item.ResourcesCount
		respItem["start_time"] = item.StartTime
		respItem["success_count"] = item.SuccessCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTasksGetAllTaskStatusItemsDetailStatus(items *[]isegosdk.ResponseItemTasksGetAllTaskStatusDetailStatus) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenTasksGetTaskStatusItem(item *isegosdk.ResponseTasksGetTaskStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["detail_status"] = flattenTasksGetTaskStatusItemDetailStatus(item.DetailStatus)
	respItem["execution_status"] = item.ExecutionStatus
	respItem["fail_count"] = item.FailCount
	respItem["id"] = item.ID
	respItem["module_type"] = item.ModuleType
	respItem["resources_count"] = item.ResourcesCount
	respItem["start_time"] = item.StartTime
	respItem["success_count"] = item.SuccessCount
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTasksGetTaskStatusItemDetailStatus(items *[]isegosdk.ResponseTasksGetTaskStatusDetailStatus) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}
