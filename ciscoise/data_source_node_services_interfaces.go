package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodeServicesInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Node Services.

- This data source retrieves the list of interfaces on a node in a cluster.
`,

		ReadContext: dataSourceNodeServicesInterfacesRead,
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Description: `hostname path parameter. Hostname of the node.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"items": &schema.Schema{
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

func dataSourceNodeServicesInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vHostname := d.Get("hostname")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetInterfaces")
		vvHostname := vHostname.(string)

		response1, restyResp1, err := client.NodeServices.GetInterfaces(vvHostname)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetInterfaces", err,
				"Failure at GetInterfaces, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNodeServicesGetInterfacesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInterfaces response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNodeServicesGetInterfacesItems(items *[]isegosdk.ResponseNodeServicesGetInterfacesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItems = append(respItems, respItem)
	}
	return respItems
}
