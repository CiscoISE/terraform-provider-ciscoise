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
		Description: `It performs update operation on Backup And Restore.

Update the Schedule of the configuration backup on the ISE node as per the input parameters. This data source action
only helps in editing the schedule.`,

		ReadContext: dataSourceBackupScheduleConfigUpdateRead,
		Schema: map[string]*schema.Schema{
			"backup_description": &schema.Schema{
				Description: `Description of the backup.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"backup_encryption_key": &schema.Schema{
				Description: `The encyption key for the backed up file. Encryption key must satisfy the following criteria - Contains at least one uppercase letter [A-Z], Contains at least one lowercase letter [a-z], Contains at least one digit [0-9], Contain only [A-Z][a-z][0-9]_#, Has at least 8 characters, Has not more than 15 characters, Must not contain 'CcIiSsCco', Must not begin with`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"backup_name": &schema.Schema{
				Description: `The backup file will get saved with this name.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"end_date": &schema.Schema{
				Description: `End date of the scheduled backup job. Allowed format MM/DD/YYYY. End date is not required in case of ONE_TIME frequency.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"frequency": &schema.Schema{
				Description: `Frequency with which the backup will get scheduled in the ISE node. Allowed values - ONCE, DAILY, WEEKLY, MONTHLY`,
				Type:        schema.TypeString,
				Optional:    true,
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
							Description: `Response message on successful scheduling the backup job.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"month_day": &schema.Schema{
				Description: `Day of month you want backup to be performed on when scheduled frequency is MONTHLY. Allowed values - from 1 to 28.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repository_name": &schema.Schema{
				Description: `Name of the configured repository where the generated backup file will get copied.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"start_date": &schema.Schema{
				Description: `Start date for scheduling the backup job. Allowed format MM/DD/YYYY.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": &schema.Schema{
				Description: `Enable or disable scheduled backup.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"time": &schema.Schema{
				Description: `Time at which backup job get scheduled. example- 12:00 AM`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"week_day": &schema.Schema{
				Description: `Day of week you want backup to be performed on when scheduled frequency is WEEKLY. Allowed values - MON, TUE, WED, THU, FRI, SAT, SUN.`,
				Type:        schema.TypeString,
				Optional:    true,
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
