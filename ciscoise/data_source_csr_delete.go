package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceCsrDelete() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCsrDeleteRead,
		Schema: map[string]*schema.Schema{
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCsrDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vHostName := d.Get("host_name")
	vID := d.Get("id")
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 2: DeleteCsrByID")
		vvHostName := vHostName.(string)
		vvID := vID.(string)

		response2, _, err := client.Certificates.DeleteCsrByID(vvHostName, vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeleteCsrByID", err,
				"Failure at DeleteCsrByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenCertificatesDeleteCsrByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeleteCsrByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags
	}
	return diags
}

func flattenCertificatesDeleteCsrByIDItem(item *isegosdk.ResponseCertificatesDeleteCsrByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
