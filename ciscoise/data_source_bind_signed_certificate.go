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
func dataSourceBindSignedCertificate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Certificates.



Bind CA Signed Certificate.

NOTE:
This data source action requires an existing Certificate Signing Request, and the root certificate must already be
trusted.

NOTE:
The certificate may have a validity period longer than 398 days. It may be untrusted by many browsers.

NOTE:
Request Parameters accepting True and False as input can be replaced by 1 and 0 respectively.


Following Parameters are used in POST body




PARAMETER

DESCRIPTION

EXAMPLE





name

Friendly name of the certificate.

Signed Certificate



data

Plain-text contents of the certificate file (required)

Signed Certificate in escaped format



allowExtendedValidity

Allow the certificates greater than validity of 398 days (required)

false



allowOutOfDateCert

Allow out of date certificates (required)

false



allowReplacementOfCertificates

Allow Replacement of certificates (required)

false



allowReplacementOfPortalGroupTag

Allow Replacement of Portal Group Tag (required)

false


admin

Use certificate to authenticate the ISE Admin Portal

false


eap

Use certificate for EAP protocols that use SSL/TLS tunneling

false



radius

Use certificate for RADSec server

false



pxgrid

Use certificate for the pxGrid Controller

false



ims

Use certificate for the ISE Messaging Service

false



saml

Use certificate for SAML Signing

false



portal

Use certificate for portal

false



portalGroupTag

Portal Group Tag for using certificate with portal role

Default Portal Certificate Group



validateCertificateExtensions

Validate Certificate Extensions

false





Following Roles can be used in any combinations




ROLE

DEFAULT

WARNING





Admin

False

Enabling Admin role for this certificate will cause an application server restart on the selected node.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



EAP Authentication

False

Only one system certificate can be used for EAP. Assigning EAP to this certificate will remove the assignment from
another certificate.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



RADIUS DTLS

False

Only one system certificate can be used for DTLS. Assigning DTLS to this certificate will remove the assignment from
another certificate.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



SAML

False

SAML cannot be used with other Usage. Enabling SAML will uncheck all other Usage.
Note:
 Make sure required Certificate Chain is imported under Trusted Certificates



 `,

		ReadContext: dataSourceBindSignedCertificateRead,
		Schema: map[string]*schema.Schema{
			"admin": &schema.Schema{
				Description: ` Use certificate to authenticate the ISE Admin Portal`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"allow_extended_validity": &schema.Schema{
				Description: `Allow import of certificates with validity greater than 398 days`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"allow_out_of_date_cert": &schema.Schema{
				Description: `Allow out of date certificates (required)`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"allow_replacement_of_certificates": &schema.Schema{
				Description: `Allow Replacement of certificates (required)`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"allow_replacement_of_portal_group_tag": &schema.Schema{
				Description: `Allow Replacement of Portal Group Tag (required)`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"data": &schema.Schema{
				Description: `Signed Certificate in escaped format`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"eap": &schema.Schema{
				Description: `Use certificate for EAP protocols that use SSL/TLS tunneling`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"host_name": &schema.Schema{
				Description: `Name of Host whose CSR ID has been provided`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": &schema.Schema{
				Description: `ID of the generated CSR`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ims": &schema.Schema{
				Description: `Use certificate for the ISE Messaging Service`,
				Type:        schema.TypeBool,
				Optional:    true,
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
						"status": &schema.Schema{
							Description: `Response status after import`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Description: `Friendly Name of the certificate`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"portal": &schema.Schema{
				Description: `Use for portal`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"portal_group_tag": &schema.Schema{
				Description: `Set Group tag`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pxgrid": &schema.Schema{
				Description: `Use certificate for the pxGrid Controller`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"radius": &schema.Schema{
				Description: `Use certificate for the RADSec server`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"saml": &schema.Schema{
				Description: `Use certificate for SAML Signing`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"validate_certificate_extensions": &schema.Schema{
				Description: `Validate Certificate Extensions`,
				Type:        schema.TypeBool,
				Optional:    true,
			},
		},
	}
}

func dataSourceBindSignedCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BindCsr")
		request1 := expandRequestBindSignedCertificateBindCsr(ctx, "", d)

		response1, _, err := client.Certificates.BindCsr(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BindCsr", err,
				"Failure at BindCsr, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenCertificatesBindCsrItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BindCsr response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestBindSignedCertificateBindCsr(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesBindCsr {
	request := isegosdk.RequestCertificatesBindCsr{}
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
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".eap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".eap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".eap"))) {
		request.Eap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".host_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".host_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".host_name"))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ims"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ims"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ims"))) {
		request.Ims = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal"))) {
		request.Portal = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_group_tag"))) {
		request.PortalGroupTag = interfaceToString(v)
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

func flattenCertificatesBindCsrItem(item *isegosdk.ResponseCertificatesBindCsrResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	respItem["status"] = item.Status
	return []map[string]interface{}{
		respItem,
	}
}
