package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGuestUserBulkMonitorStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGuestUserBulkMonitorStatusRead,
		Schema: map[string]*schema.Schema{
			"bulkid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bulk_status": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bulk_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"media_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"execution_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"operation_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"resources_count": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"success_count": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"fail_count": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"resources_status": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"resource_execution_status": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": &schema.Schema{
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
		},
	}
}

func dataSourceGuestUserBulkMonitorStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vBulkid := d.Get("bulkid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: MonitorBulkStatusGuestUser")
		vvBulkid := vBulkid.(string)

		response1, _, err := client.GuestUser.MonitorBulkStatusGuestUser(vvBulkid)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing MonitorBulkStatusGuestUser", err,
				"Failure at MonitorBulkStatusGuestUser, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenGuestUserMonitorBulkStatusGuestUserItem(&response1.BulkStatus)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting MonitorBulkStatusGuestUser response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenGuestUserMonitorBulkStatusGuestUserItem(item *isegosdk.ResponseGuestUserMonitorBulkStatusGuestUserBulkStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["bulk_id"] = item.BulkID
	respItem["media_type"] = item.MediaType
	respItem["execution_status"] = item.ExecutionStatus
	respItem["operation_type"] = item.OperationType
	respItem["start_time"] = item.StartTime
	respItem["resources_count"] = item.ResourcesCount
	respItem["success_count"] = item.SuccessCount
	respItem["fail_count"] = item.FailCount
	respItem["resources_status"] = flattenGuestUserMonitorBulkStatusGuestUserItemResourcesStatus(item.ResourcesStatus)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenGuestUserMonitorBulkStatusGuestUserItemResourcesStatus(items []isegosdk.ResponseGuestUserMonitorBulkStatusGuestUserBulkStatusResourcesStatus) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["resource_execution_status"] = item.ResourceExecutionStatus
		respItem["status"] = item.Status
	}
	return respItems

}
