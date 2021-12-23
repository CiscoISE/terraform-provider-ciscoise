package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceHotpatchInstall() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Patching.

- Triggers hot patch installation on the Cisco ISE node. A task ID is returned which  can be used to monitor the
progress of the hot patch installation process.  As hot patch installation triggers the Cisco ISE to restart, the task
API becomes  unavailable for a certain period of time.
`,

		ReadContext: dataSourceHotpatchInstallRead,
		Schema: map[string]*schema.Schema{
			"hotpatch_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `ID which can be used to track the status of install / rollback of patch and hotpatch.`,
							Type:        schema.TypeString,
							Computed:    true,
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

func dataSourceHotpatchInstallRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: InstallHotpatch")
		request1 := expandRequestHotpatchInstallInstallHotpatch(ctx, "", d)

		response1, restyResp1, err := client.Patching.InstallHotpatch(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing InstallHotpatch", err,
				"Failure at InstallHotpatch, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPatchingInstallHotpatchItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting InstallHotpatch response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestHotpatchInstallInstallHotpatch(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPatchingInstallHotpatch {
	request := isegosdk.RequestPatchingInstallHotpatch{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hotpatch_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hotpatch_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hotpatch_name")))) {
		request.HotpatchName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".repository_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".repository_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".repository_name")))) {
		request.RepositoryName = interfaceToString(v)
	}
	return &request
}

func flattenPatchingInstallHotpatchItem(item *isegosdk.ResponsePatchingInstallHotpatchResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
