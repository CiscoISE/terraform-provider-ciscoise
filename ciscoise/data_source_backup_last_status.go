package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBackupLastStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBackupLastStatusRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"action": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"details": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"initiated_from": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"just_complete": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"percent_complete": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"repository": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"scheduled": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
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
	}
}

func dataSourceBackupLastStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetLastConfigBackupStatus")

		response1, _, err := client.BackupAndRestore.GetLastConfigBackupStatus()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetLastConfigBackupStatus", err,
				"Failure at GetLastConfigBackupStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenBackupAndRestoreGetLastConfigBackupStatusItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLastConfigBackupStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenBackupAndRestoreGetLastConfigBackupStatusItem(item *isegosdk.ResponseBackupAndRestoreGetLastConfigBackupStatusResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["action"] = item.Action
	respItem["details"] = item.Details
	respItem["error"] = item.Error
	respItem["host_name"] = item.HostName
	respItem["initiated_from"] = item.InitiatedFrom
	respItem["just_complete"] = item.JustComplete
	respItem["message"] = item.Message
	respItem["name"] = item.Name
	respItem["percent_complete"] = item.PercentComplete
	respItem["repository"] = item.Repository
	respItem["scheduled"] = item.Scheduled
	respItem["start_date"] = item.StartDate
	respItem["status"] = item.Status
	respItem["type"] = item.Type
	return []map[string]interface{}{
		respItem,
	}
}
