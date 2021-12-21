package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePatch() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Patching.

- List all the installed patches in the system, with the patch number for rollback.
`,

		ReadContext: dataSourcePatchRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ise_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"patch_version": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"install_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"patch_number": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
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

func dataSourcePatchRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ListInstalledPatches")

		response1, restyResp1, err := client.Patching.ListInstalledPatches()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ListInstalledPatches", err,
				"Failure at ListInstalledPatches, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPatchingListInstalledPatchesItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ListInstalledPatches response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPatchingListInstalledPatchesItem(item *isegosdk.ResponsePatchingListInstalledPatches) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ise_version"] = item.IseVersion
	respItem["patch_version"] = flattenPatchingListInstalledPatchesItemPatchVersion(item.PatchVersion)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenPatchingListInstalledPatchesItemPatchVersion(items *[]isegosdk.ResponsePatchingListInstalledPatchesPatchVersion) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["install_date"] = item.InstallDate
		respItem["patch_number"] = item.PatchNumber
		respItems = append(respItems, respItem)
	}
	return respItems
}
