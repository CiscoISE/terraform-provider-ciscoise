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
func dataSourcePatchRollback() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Patching.

- Triggers patch rollback on the Cisco ISE node. A task ID is returned which can be used to monitor the progress of the
patch rollback process. As the patch   rollback triggers the Cisco ISE to restart, the task API becomes unavailable for
a certain period of time.
`,

		ReadContext: dataSourcePatchRollbackRead,
		Schema: map[string]*schema.Schema{
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
			"patch_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func dataSourcePatchRollbackRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RollbackPatch")
		request1 := expandRequestPatchRollbackRollbackPatch(ctx, "", d)

		response1, restyResp1, err := client.Patching.RollbackPatch(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RollbackPatch", err,
				"Failure at RollbackPatch, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPatchingRollbackPatchItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RollbackPatch response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPatchRollbackRollbackPatch(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPatchingRollbackPatch {
	request := isegosdk.RequestPatchingRollbackPatch{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".patch_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".patch_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".patch_number")))) {
		request.PatchNumber = interfaceToIntPtr(v)
	}
	return &request
}

func flattenPatchingRollbackPatchItem(item *isegosdk.ResponsePatchingRollbackPatchResponse) []map[string]interface{} {
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
