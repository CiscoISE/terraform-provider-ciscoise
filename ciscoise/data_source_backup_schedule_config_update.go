package ciscoise

import (
	"context"

	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceBackupScheduleConfigUpdate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBackupScheduleConfigUpdateRead,
		Schema: map[string]*schema.Schema{
			"backup_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_encryption_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_date": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"month_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_date": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"week_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceBackupScheduleConfigUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateScheduledConfigBackup")
		request1 := expandRequestBackupScheduleConfigUpdateScheduledConfigBackup(ctx, "", d)

		response1, _, err := client.BackupAndRestore.UpdateScheduledConfigBackup(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateScheduledConfigBackup", err,
				"Failure at UpdateScheduledConfigBackup, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenBackupAndRestoreUpdateScheduledConfigBackupItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateScheduledConfigBackup response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestBackupScheduleConfigUpdateScheduledConfigBackup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestBackupAndRestoreUpdateScheduledConfigBackup {
	request := isegosdk.RequestBackupAndRestoreUpdateScheduledConfigBackup{}
	if v, ok := d.GetOkExists(key + ".backup_description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".backup_description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".backup_description"))) {
		request.BackupDescription = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".backup_encryption_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".backup_encryption_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".backup_encryption_key"))) {
		request.BackupEncryptionKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".backup_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".backup_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".backup_name"))) {
		request.BackupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".end_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_date"))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".frequency"); !isEmptyValue(reflect.ValueOf(d.Get(key+".frequency"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".frequency"))) {
		request.Frequency = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".month_day"); !isEmptyValue(reflect.ValueOf(d.Get(key+".month_day"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".month_day"))) {
		request.MonthDay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".repository_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".repository_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".repository_name"))) {
		request.RepositoryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_date"))) {
		request.StartDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".status"); !isEmptyValue(reflect.ValueOf(d.Get(key+".status"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".status"))) {
		request.Status = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".time"))) {
		request.Time = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".week_day"); !isEmptyValue(reflect.ValueOf(d.Get(key+".week_day"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".week_day"))) {
		request.WeekDay = interfaceToString(v)
	}
	return &request
}

func flattenBackupAndRestoreUpdateScheduledConfigBackupItem(item *isegosdk.ResponseBackupAndRestoreUpdateScheduledConfigBackupResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	respItem["link"] = flattenBackupAndRestoreUpdateScheduledConfigBackupItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenBackupAndRestoreUpdateScheduledConfigBackupItemLink(item *isegosdk.ResponseBackupAndRestoreUpdateScheduledConfigBackupResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
