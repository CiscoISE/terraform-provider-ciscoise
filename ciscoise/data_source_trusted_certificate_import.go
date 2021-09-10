package ciscoise

import (
	"context"

	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceTrustedCertificateImport() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTrustedCertificateImportRead,
		Schema: map[string]*schema.Schema{
			"allow_basic_constraint_cafalse": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_out_of_date_cert": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_sha1_certificates": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trust_for_certificate_based_admin_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"trust_for_cisco_services_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"trust_for_client_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"trust_for_ise_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"validate_certificate_extensions": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func dataSourceTrustedCertificateImportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ImportTrustCertificate")
		request1 := expandRequestTrustedCertificateImportImportTrustCertificate(ctx, "", d)

		response1, _, err := client.Certificates.ImportTrustCertificate(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportTrustCertificate", err,
				"Failure at ImportTrustCertificate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenCertificatesImportTrustCertificateItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportTrustCertificate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestTrustedCertificateImportImportTrustCertificate(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesImportTrustCertificate {
	request := isegosdk.RequestCertificatesImportTrustCertificate{}
	if v, ok := d.GetOkExists(key + ".allow_basic_constraint_cafalse"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_basic_constraint_cafalse"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_basic_constraint_cafalse"))) {
		request.AllowBasicConstraintCaFalse = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_out_of_date_cert"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_out_of_date_cert"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_out_of_date_cert"))) {
		request.AllowOutOfDateCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_sha1_certificates"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_sha1_certificates"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_sha1_certificates"))) {
		request.AllowSHA1Certificates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".validate_certificate_extensions"); !isEmptyValue(reflect.ValueOf(d.Get(key+".validate_certificate_extensions"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".validate_certificate_extensions"))) {
		request.ValidateCertificateExtensions = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".trust_for_certificate_based_admin_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".trust_for_certificate_based_admin_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".trust_for_certificate_based_admin_auth"))) {
		request.TrustForCertificateBasedAdminAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".trust_for_cisco_services_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".trust_for_cisco_services_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".trust_for_cisco_services_auth"))) {
		request.TrustForCiscoServicesAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".trust_for_client_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".trust_for_client_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".trust_for_client_auth"))) {
		request.TrustForClientAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".trust_for_ise_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".trust_for_ise_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".trust_for_ise_auth"))) {
		request.TrustForIseAuth = interfaceToBoolPtr(v)
	}
	return &request
}

func flattenCertificatesImportTrustCertificateItem(item *isegosdk.ResponseCertificatesImportTrustCertificateResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["message"] = item.Message
	respItem["status"] = item.Status
	return []map[string]interface{}{
		respItem,
	}
}
