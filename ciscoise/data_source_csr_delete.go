package ciscoise

import (
	"context"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceCsrDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Certificates.

- This data source action deletes a Certificate Signing Request of a particular node based on a given HostName and ID.
`,

		ReadContext: dataSourceCsrDeleteRead,
		Schema: map[string]*schema.Schema{
			"host_name": &schema.Schema{
				Description: `hostName path parameter. Name of the host of which CSR's should be deleted`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. The ID of the Certificate Signing Request to be deleted`,
				Type:        schema.TypeString,
				Required:    true,
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
	vHostName, okHostName := d.GetOk("host_name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okHostName, okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okHostName, okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetCsrByID")
		vvHostName := vHostName.(string)
		vvID := vID.(string)

		response1, restyResp1, err := client.Certificates.GetCsrByID(vvHostName, vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetCsrByID", err,
				"Failure at GetCsrByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	if selectedMethod == 2 {
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

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

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
