package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRenewCertificate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Certificates.
		- This resource initiates regeneration of certificates. Response contains ID which can be used to track the
		status
		`,

		CreateContext: resourceRenewCertificateCreate,
		ReadContext:   resourceRenewCertificateRead,
		DeleteContext: resourceRenewCertificateDelete,

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
							Description: `ID which can be used to track status of certificate regeneration`,
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
						"cert_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceRenewCertificateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning RenewCerts create")
	log.Printf("[DEBUG] Missing RenewCerts create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestRenewCertificateRenewCerts(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, restyResp1, err := client.Certificates.RenewCerts(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing RenewCerts", err, restyResp1.String(),
				"Failure at RenewCerts, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RenewCerts", err,
			"Failure at RenewCerts, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenCertificatesRenewCertsItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RenewCerts response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceRenewCertificateRead(ctx, d, m)
}

func resourceRenewCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceRenewCertificateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning RenewCertificate delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing RenewCertificate delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
func expandRequestRenewCertificateRenewCerts(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesRenewCerts {
	request := isegosdk.RequestCertificatesRenewCerts{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cert_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cert_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cert_type")))) {
		request.CertType = interfaceToString(v)
	}
	return &request
}

func flattenCertificatesRenewCertsItem(item *isegosdk.ResponseCertificatesRenewCertsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["link"] = flattenCertificatesRenewCertsItemLink(item.Link)
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}

func flattenCertificatesRenewCertsItemLink(item *isegosdk.ResponseCertificatesRenewCertsResponseLink) []map[string]interface{} {
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
