package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceBackupConfig() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Backup And Restore.

- Triggers on demand configuration backup on the ISE node. The API returns the task ID. Use the Task Service status API
to get the status of the backup job.
`,

		ReadContext: dataSourceBackupConfigRead,
		Schema: map[string]*schema.Schema{
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
				Description: `Name of the configured repository where the generated backup file will get copied.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceBackupConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ConfigBackup")
		request1 := expandRequestBackupConfigConfigBackup(ctx, "", d)

		response1, restyResp1, err := client.BackupAndRestore.ConfigBackup(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ConfigBackup", err,
				"Failure at ConfigBackup, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenBackupAndRestoreConfigBackupItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ConfigBackup response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestBackupConfigConfigBackup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestBackupAndRestoreConfigBackup {
	request := isegosdk.RequestBackupAndRestoreConfigBackup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".backup_encryption_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".backup_encryption_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".backup_encryption_key")))) {
		request.BackupEncryptionKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".backup_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".backup_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".backup_name")))) {
		request.BackupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".repository_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".repository_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".repository_name")))) {
		request.RepositoryName = interfaceToString(v)
	}
	return &request
}

func flattenBackupAndRestoreConfigBackupItem(item *isegosdk.ResponseBackupAndRestoreConfigBackupResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["message"] = item.Message
	respItem["link"] = flattenBackupAndRestoreConfigBackupItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenBackupAndRestoreConfigBackupItemLink(item *isegosdk.ResponseBackupAndRestoreConfigBackupResponseLink) []map[string]interface{} {
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
