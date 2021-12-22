package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHotpatch() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Patching.

- List all the installed hot patches in the system.
`,

		ReadContext: dataSourceHotpatchRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"hotpatch_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"install_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceHotpatchRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ListInstalledHotpatches")

		response1, restyResp1, err := client.Patching.ListInstalledHotpatches()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ListInstalledHotpatches", err,
				"Failure at ListInstalledHotpatches, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenPatchingListInstalledHotpatchesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ListInstalledHotpatches response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPatchingListInstalledHotpatchesItems(items *[]isegosdk.ResponsePatchingListInstalledHotpatchesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["hotpatch_name"] = item.HotpatchName
		respItem["install_date"] = item.InstallDate
		respItems = append(respItems, respItem)
	}
	return respItems
}
