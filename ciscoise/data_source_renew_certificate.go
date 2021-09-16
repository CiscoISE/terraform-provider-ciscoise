package ciscoise

import (
	"context"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceRenewCertificate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Certificates.

- This data source action will initiate regeneration of certificates. Response contains id which can be used to track the
status`,

		ReadContext: dataSourceRenewCertificateRead,
		Schema: map[string]*schema.Schema{
			"cert_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Id which can be used to track status of certificate regeneration`,
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
		},
	}
}

func dataSourceRenewCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RenewCertificates")
		request1 := expandRequestRenewCertificateRenewCertificates(ctx, "", d)

		response1, _, err := client.Certificates.RenewCertificates(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RenewCertificates", err,
				"Failure at RenewCertificates, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenCertificatesRenewCertificatesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RenewCertificates response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestRenewCertificateRenewCertificates(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesRenewCertificates {
	request := isegosdk.RequestCertificatesRenewCertificates{}
	if v, ok := d.GetOkExists(key + ".cert_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".cert_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".cert_type"))) {
		request.CertType = interfaceToString(v)
	}
	return &request
}

func flattenCertificatesRenewCertificatesItem(item *isegosdk.ResponseCertificatesRenewCertificatesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["link"] = flattenCertificatesRenewCertificatesItemLink(item.Link)
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}

func flattenCertificatesRenewCertificatesItemLink(item *isegosdk.ResponseCertificatesRenewCertificatesResponseLink) []map[string]interface{} {
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
