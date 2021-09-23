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
func dataSourceIseRootCaRegenerate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Certificates.

- This data source action will initiate regeneration of ISE root CA certificate chain. Response contains id which can be
used to track the status.
  Setting "removeExistingISEIntermediateCSR" to true will remove existing ISE Intermediate CSR
`,

		ReadContext: dataSourceIseRootCaRegenerateRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Id which can be used to track status of ISE root CA chain regeneration`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"remove_existing_ise_intermediate_csr": &schema.Schema{
				Description: `Setting this attribute to true will remove existing ISE Intermediate CSR`,
				// Type:        schema.TypeBool,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
		},
	}
}

func dataSourceIseRootCaRegenerateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RegenerateIseRootCa")
		request1 := expandRequestIseRootCaRegenerateRegenerateIseRootCa(ctx, "", d)

		response1, _, err := client.Certificates.RegenerateIseRootCa(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RegenerateIseRootCa", err,
				"Failure at RegenerateIseRootCa, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenCertificatesRegenerateIseRootCaItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RegenerateIseRootCa response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestIseRootCaRegenerateRegenerateIseRootCa(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesRegenerateIseRootCa {
	request := isegosdk.RequestCertificatesRegenerateIseRootCa{}
	if v, ok := d.GetOkExists(key + ".remove_existing_ise_intermediate_csr"); !isEmptyValue(reflect.ValueOf(d.Get(key+".remove_existing_ise_intermediate_csr"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".remove_existing_ise_intermediate_csr"))) {
		request.RemoveExistingIseIntermediateCsr = interfaceToBoolPtr(v)
	}
	return &request
}

func flattenCertificatesRegenerateIseRootCaItem(item *isegosdk.ResponseCertificatesRegenerateIseRootCaResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["link"] = flattenCertificatesRegenerateIseRootCaItemLink(item.Link)
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}

func flattenCertificatesRegenerateIseRootCaItemLink(item *isegosdk.ResponseCertificatesRegenerateIseRootCaResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
