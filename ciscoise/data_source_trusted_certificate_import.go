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
func dataSourceTrustedCertificateImport() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Certificates.



Import an X509 certificate as a trust certificate.

NOTE:
Request Parameters accepting True and False as input can be replaced by 1 and 0 respectively.

Following Parameters are used in POST body




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate

Trust Certificate



description

Description of the certificate

Passw***



data

Plain-text contents of the certificate file (required)

Trust Certificate in escaped format



allowOutOfDateCert

Allow out of date certificates (required)

false



allowSHA1Certificates

Allow SHA1 based certificates (required)

false




trustForIseAuth

Trust for authentication within ISE

false



trustForClientAuth

Trust for client authentication and Syslog

false



trustForCertificateBasedAdminAuth

Trust for Certificate based Admin authentication

false



trustForCiscoServicesAuth

Trust for authentication of Cisco Services

false



validateCertificateExtensions

Validate extensions for trust certificate

false




NOTE:
If name is not set, a default name of the following format will be generated:
common-name#issuer#nnnnn

    where
"nnnnn"
 is a unique number. You can always change the friendly name later by editing the certificate.


    You must choose how this certificate will be trusted in ISE. The objective here is to distinguish between
certificates that are used for trust within an ISE deployment and public certificates that are used to trust Cisco
services. Typically, you will not want to use a given certificate for both purposes.




Trusted For

Usage





Authentication within ISE

Use
"trustForIseAuth":true
 if the certificate is used for trust within ISE, such as for secure communication between ISE nodes



Client authentication and Syslog

Use
"trustForClientAuth":true
 if the certificate is to be used for authentication of endpoints that contact ISE over the EAP protocol. Also check
this box if certificate is used to trust a Syslog server. Make sure to have keyCertSign bit asserted under KeyUsage
extension for this certificate.
Note:
 "" can be set true only if the "trustForIseAuth" has been set true.



Certificate based admin authentication

Use
"trustForCertificateBasedAdminAuth":true
 if the certificate is used for trust within ISE, such as for secure communication between ISE nodes
Note:
 "trustForCertificateBasedAdminAuth" can be set true only if "trustForIseAuth" and "trustForClientAuth" are true.



Authentication of Cisco Services

 Use
"trustForCiscoServicesAuth":true
 if the certificate is to be used for trusting external Cisco services, such as Feed Service.



 `,

		ReadContext: dataSourceTrustedCertificateImportRead,
		Schema: map[string]*schema.Schema{
			"allow_basic_constraint_cafalse": &schema.Schema{
				Description: `Allow Certificates with Basic Constraints CA Field as False (required)`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"allow_out_of_date_cert": &schema.Schema{
				Description: `Allow out of date certificates (required)`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"allow_sha1_certificates": &schema.Schema{
				Description: `Allow SHA1 based certificates (required)`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"data": &schema.Schema{
				Description: `Certificate content (required)`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": &schema.Schema{
				Description: `Description of the certificate`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `ID of the imported trust certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Response message on import of system or trust certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status": &schema.Schema{
							Description: `HTTP response status after import`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Description: `Name of the certificate`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"trust_for_certificate_based_admin_auth": &schema.Schema{
				Description: `Trust for Certificate based Admin authentication`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"trust_for_cisco_services_auth": &schema.Schema{
				Description: `Trust for authentication of Cisco Services`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"trust_for_client_auth": &schema.Schema{
				Description: `Trust for client authentication and Syslog`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"trust_for_ise_auth": &schema.Schema{
				Description: `Trust for authentication within ISE`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"validate_certificate_extensions": &schema.Schema{
				Description: `Validate trust certificate extension`,
				Type:        schema.TypeBool,
				Optional:    true,
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
