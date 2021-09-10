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
func dataSourceSystemCertificateImport() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSystemCertificateImportRead,
		Schema: map[string]*schema.Schema{
			"admin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_extended_validity": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_out_of_date_cert": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_replacement_of_certificates": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_replacement_of_portal_group_tag": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_sha1_certificates": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_wild_card_certificates": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eap": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ims": &schema.Schema{
				Type:     schema.TypeBool,
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
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"portal": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"portal_group_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_key_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pxgrid": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"radius": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"saml": &schema.Schema{
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

func dataSourceSystemCertificateImportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ImportSystemCertificate")
		request1 := expandRequestSystemCertificateImportImportSystemCertificate(ctx, "", d)

		response1, _, err := client.Certificates.ImportSystemCertificate(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportSystemCertificate", err,
				"Failure at ImportSystemCertificate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenCertificatesImportSystemCertificateItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportSystemCertificate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSystemCertificateImportImportSystemCertificate(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesImportSystemCertificate {
	request := isegosdk.RequestCertificatesImportSystemCertificate{}
	if v, ok := d.GetOkExists(key + ".admin"); !isEmptyValue(reflect.ValueOf(d.Get(key+".admin"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".admin"))) {
		request.Admin = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_extended_validity"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_extended_validity"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_extended_validity"))) {
		request.AllowExtendedValidity = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_out_of_date_cert"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_out_of_date_cert"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_out_of_date_cert"))) {
		request.AllowOutOfDateCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_replacement_of_certificates"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_replacement_of_certificates"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_replacement_of_certificates"))) {
		request.AllowReplacementOfCertificates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_replacement_of_portal_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_replacement_of_portal_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_replacement_of_portal_group_tag"))) {
		request.AllowReplacementOfPortalGroupTag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_sha1_certificates"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_sha1_certificates"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_sha1_certificates"))) {
		request.AllowSHA1Certificates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_wild_card_certificates"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_wild_card_certificates"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_wild_card_certificates"))) {
		request.AllowWildCardCertificates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".eap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".eap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".eap"))) {
		request.Eap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".ims"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ims"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ims"))) {
		request.Ims = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal"))) {
		request.Portal = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_group_tag"))) {
		request.PortalGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".private_key_data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".private_key_data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".private_key_data"))) {
		request.PrivateKeyData = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".pxgrid"); !isEmptyValue(reflect.ValueOf(d.Get(key+".pxgrid"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".pxgrid"))) {
		request.Pxgrid = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".radius"); !isEmptyValue(reflect.ValueOf(d.Get(key+".radius"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".radius"))) {
		request.Radius = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".saml"); !isEmptyValue(reflect.ValueOf(d.Get(key+".saml"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".saml"))) {
		request.Saml = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".validate_certificate_extensions"); !isEmptyValue(reflect.ValueOf(d.Get(key+".validate_certificate_extensions"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".validate_certificate_extensions"))) {
		request.ValidateCertificateExtensions = interfaceToBoolPtr(v)
	}
	return &request
}

func flattenCertificatesImportSystemCertificateItem(item *isegosdk.ResponseCertificatesImportSystemCertificateResponse) []map[string]interface{} {
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
