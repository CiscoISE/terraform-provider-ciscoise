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
func dataSourceBackupRestore() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Backup And Restore.

- Triggers a configuration DB restore job on the ISE node. The API returns the task ID. Use the Task Service status API to
get the status of the backup job`,

		ReadContext: dataSourceBackupRestoreRead,
		Schema: map[string]*schema.Schema{
			"backup_encryption_key": &schema.Schema{
				Description: `The encryption key which was provided at the time of taking backup.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Id which can be used to track the status of backup / restore of config DB.`,
							Type:        schema.TypeString,
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
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"repository_name": &schema.Schema{
				Description: `Name of the configred repository where the backup file exists.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"restore_file": &schema.Schema{
				Description: `Name of the backup file to be restored on ISE node.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"restore_include_adeos": &schema.Schema{
				Description: `Determines whether the ADE-OS configure is restored. Possible values true, false`,
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceBackupRestoreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RestoreConfigBackup")
		request1 := expandRequestBackupRestoreRestoreConfigBackup(ctx, "", d)

		response1, _, err := client.BackupAndRestore.RestoreConfigBackup(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RestoreConfigBackup", err,
				"Failure at RestoreConfigBackup, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenBackupAndRestoreRestoreConfigBackupItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RestoreConfigBackup response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestBackupRestoreRestoreConfigBackup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestBackupAndRestoreRestoreConfigBackup {
	request := isegosdk.RequestBackupAndRestoreRestoreConfigBackup{}
	if v, ok := d.GetOkExists(key + ".backup_encryption_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".backup_encryption_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".backup_encryption_key"))) {
		request.BackupEncryptionKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".repository_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".repository_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".repository_name"))) {
		request.RepositoryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".restore_file"); !isEmptyValue(reflect.ValueOf(d.Get(key+".restore_file"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".restore_file"))) {
		request.RestoreFile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".restore_include_adeos"); !isEmptyValue(reflect.ValueOf(d.Get(key+".restore_include_adeos"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".restore_include_adeos"))) {
		request.RestoreIncludeAdeos = interfaceToString(v)
	}
	return &request
}

func flattenBackupAndRestoreRestoreConfigBackupItem(item *isegosdk.ResponseBackupAndRestoreRestoreConfigBackupResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["message"] = item.Message
	respItem["link"] = flattenBackupAndRestoreRestoreConfigBackupItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenBackupAndRestoreRestoreConfigBackupItemLink(item *isegosdk.ResponseBackupAndRestoreRestoreConfigBackupResponseLink) []map[string]interface{} {
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
