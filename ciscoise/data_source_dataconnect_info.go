package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDataconnectInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Dataconnect Services.

- This data source retrieves the Dataconnect ODBC details.
`,

		ReadContext: dataSourceDataconnectInfoRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_name": &schema.Schema{
							Description: `hostname value`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"port": &schema.Schema{
							Description: `port value`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"service_name": &schema.Schema{
							Description: `servicename value`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"user_name": &schema.Schema{
							Description: `username value`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDataconnectInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOdbcDetail")

		response1, restyResp1, err := client.DataconnectServices.GetOdbcDetail()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetOdbcDetail", err,
				"Failure at GetOdbcDetail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDataconnectServicesGetOdbcDetailItemResponse(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetOdbcDetail response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDataconnectServicesGetOdbcDetailItemResponse(item *isegosdk.ResponseDataconnectServicesGetOdbcDetailResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_name"] = item.Hostname
	respItem["port"] = item.Port
	respItem["servicename"] = item.Servicename
	respItem["user_name"] = item.Username

	return []map[string]interface{}{
		respItem,
	}

}
