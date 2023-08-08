package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBackupRestore() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Backup And Restore.
- Triggers a configuration DB restore job on the ISE node. The API returns the task ID. Use the Task Service status API
to get the status of the backup job
`,

		CreateContext: resourceBackupRestoreCreate,
		ReadContext:   resourceBackupRestoreRead,
		DeleteContext: resourceBackupRestoreDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_encryption_key": &schema.Schema{
							Description: `The encryption key which was provided at the time of taking backup.`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"repository_name": &schema.Schema{
							Description: `Name of the configred repository where the backup file exists.`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"restore_file": &schema.Schema{
							Description: `Name of the backup file to be restored on ISE node.`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"restore_include_adeos": &schema.Schema{
							Description: `Determines whether the ADE-OS configure is restored. Possible values true, false`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func resourceBackupRestoreCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning RestoreConfigBackup create")
	log.Printf("[DEBUG] Missing RestoreConfigBackup create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestBackupRestoreRestoreConfigBackup(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, restyResp1, err := client.BackupAndRestore.RestoreConfigBackup(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing RestoreConfigBackup", err, restyResp1.String(),
				"Failure at RestoreConfigBackup, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RestoreConfigBackup", err,
			"Failure at RestoreConfigBackup, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenBackupAndRestoreRestoreConfigBackupItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RestoreConfigBackup response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceBackupRestoreRead(ctx, d, m)
}

func resourceBackupRestoreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceBackupRestoreDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BackupRestore delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing BackupRestore delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
func expandRequestBackupRestoreRestoreConfigBackup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestBackupAndRestoreRestoreConfigBackup {
	request := isegosdk.RequestBackupAndRestoreRestoreConfigBackup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".backup_encryption_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".backup_encryption_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".backup_encryption_key")))) {
		request.BackupEncryptionKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".repository_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".repository_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".repository_name")))) {
		request.RepositoryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".restore_file")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".restore_file")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".restore_file")))) {
		request.RestoreFile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".restore_include_adeos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".restore_include_adeos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".restore_include_adeos")))) {
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
