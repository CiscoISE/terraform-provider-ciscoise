package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodeServicesSxpInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Node Services.

- This data source retrieves the SXP interface.
`,

		ReadContext: dataSourceNodeServicesSxpInterfacesRead,
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Description: `hostname path parameter. Hostname of the node.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNodeServicesSxpInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vHostname := d.Get("hostname")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSxpInterface")
		vvHostname := vHostname.(string)

		response1, restyResp1, err := client.NodeServices.GetSxpInterface(vvHostname)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSxpInterface", err,
				"Failure at GetSxpInterface, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNodeServicesGetSxpInterfaceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSxpInterface response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNodeServicesGetSxpInterfaceItem(item *isegosdk.ResponseNodeServicesGetSxpInterfaceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interface"] = item.Interface
	return []map[string]interface{}{
		respItem,
	}
}
