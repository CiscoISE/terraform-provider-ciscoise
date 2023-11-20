package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationServiceNames() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Administration - Service Names.

- Returns list of Allowed Protocols and Server Sequences for Device Admin Policy Set results.
 'isLocalAuthorization' property is available only for Network Access Policy Set results of type Server Sequence.
 (Other CRUD APIs available throught ERS)
`,

		ReadContext: dataSourceDeviceAdministrationServiceNamesRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_local_authorization": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_type": &schema.Schema{
							Description: `Allowed Protocols OR Server Sequence`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceAdministrationServiceNamesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminServiceNames")

		response1, restyResp1, err := client.DeviceAdministrationServiceNames.GetDeviceAdminServiceNames()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminServiceNames", err,
				"Failure at GetDeviceAdminServiceNames, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceAdministrationServiceNamesGetDeviceAdminServiceNamesItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminServiceNames response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationServiceNamesGetDeviceAdminServiceNamesItems(items *isegosdk.ResponseDeviceAdministrationServiceNamesGetDeviceAdminServiceNames) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["is_local_authorization"] = boolPtrToString(item.IsLocalAuthorization)
		respItem["name"] = item.Name
		respItem["service_type"] = item.ServiceType
		respItems = append(respItems, respItem)
	}
	return respItems
}
