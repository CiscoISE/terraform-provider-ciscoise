package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIseRootCaRegenerate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Certificates.
- This data source action initiates regeneration of Cisco ISE root CA certificate chain. Response contains ID which can
be used to track the status.
  Setting "removeExistingISEIntermediateCSR" to true removes existing Cisco ISE Intermediate CSR
`,

		CreateContext: resourceIseRootCaRegenerateCreate,
		ReadContext:   resourceIseRootCaRegenerateRead,
		DeleteContext: resourceIseRootCaRegenerateDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `ID which can be used to track status of Cisco ISE root CA chain regeneration`,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"remove_existing_ise_intermediate_csr": &schema.Schema{
							Description:  `Setting this attribute to true removes existing Cisco ISE Intermediate CSR`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
					},
				},
			},
		},
	}
}

func resourceIseRootCaRegenerateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning RegenerateIseRootCa create")
	log.Printf("[DEBUG] Missing RegenerateIseRootCa create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	request1 := expandRequestIseRootCaRegenerateRegenerateIseRootCa(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, restyResp1, err := client.Certificates.RegenerateIseRootCa(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing RegenerateIseRootCa", err, restyResp1.String(),
				"Failure at RegenerateIseRootCa, unexpected response", ""))
			return diags
		}
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
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceIseRootCaRegenerateRead(ctx, d, m)
}

func resourceIseRootCaRegenerateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceIseRootCaRegenerateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning IseRootCaRegenerate delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing IseRootCaRegenerate delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
func expandRequestIseRootCaRegenerateRegenerateIseRootCa(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesRegenerateIseRootCa {
	request := isegosdk.RequestCertificatesRegenerateIseRootCa{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remove_existing_ise_intermediate_csr")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remove_existing_ise_intermediate_csr")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remove_existing_ise_intermediate_csr")))) {
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
