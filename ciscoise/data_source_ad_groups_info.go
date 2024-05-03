package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAdGroupsInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ADGroups.

- Duo-IdentitySync  Get the list of all AD groups for the specified Active Directory
`,

		ReadContext: dataSourceAdGroupsInfoRead,
		Schema: map[string]*schema.Schema{
			"active_directory": &schema.Schema{
				Description: `activeDirectory path parameter. List of AD groups for the specified Active Directory`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Description: `Active Directory Group ID`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"source": &schema.Schema{
							Description: `Source of the Active Directory Group`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAdGroupsInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vActiveDirectory := d.Get("active_directory")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdgroups")
		vvActiveDirectory := vActiveDirectory.(string)

		response1, restyResp1, err := client.ADGroups.GetAdgroups(vvActiveDirectory)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAdgroups", err,
				"Failure at GetAdgroups, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenADGroupsGetAdgroupsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAdgroups response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenADGroupsGetAdgroupsItems(items *isegosdk.ResponseADGroupsGetAdgroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["source"] = item.Source
		respItems = append(respItems, respItem)
	}
	return respItems
}
