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
func dataSourceBackupConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBackupConfigRead,
		Schema: map[string]*schema.Schema{
			"backup_encryption_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
				Type:     schema.TypeString,
				Optional: true,
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

		response1, _, err := client.BackupAndRestore.ConfigBackup(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ConfigBackup", err,
				"Failure at ConfigBackup, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

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
	if v, ok := d.GetOkExists(key + ".backup_encryption_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".backup_encryption_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".backup_encryption_key"))) {
		request.BackupEncryptionKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".backup_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".backup_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".backup_name"))) {
		request.BackupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".repository_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".repository_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".repository_name"))) {
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
