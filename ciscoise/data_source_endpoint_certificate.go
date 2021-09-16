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
func dataSourceEndpointCertificate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on EndpointCertificate.

This data source action allows the client to create an endpoint certificate.`,

		ReadContext: dataSourceEndpointCertificateRead,
		Schema: map[string]*schema.Schema{
			"dirpath": &schema.Schema{
				Description: `Directory absolute path in which to save the file.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"cert_template_name": &schema.Schema{
				Description: `Name of an Internal CA template`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"certificate_request": &schema.Schema{
				Description: `Key value map. Must have CN and SAN entries`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cn": &schema.Schema{
							Description: `Matches the requester's User Name, unless the Requester is an ERS Admin.
ERS Admins are allowed to create requests for any CN`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"san": &schema.Schema{
							Description: `Valid MAC Address, delimited by '-'`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"format": &schema.Schema{
				Description: `Allowed values:
- PKCS12,
- PKCS12_CHAIN,
- PKCS8,
- PKCS8_CHAIN`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Description: `Protects the private key. Must have more than 8 characters, less than 15 characters,
at least one upper case letter, at least one lower case letter, at least one digit,
and can only contain [A-Z][a-z][0-9]_#`,
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceEndpointCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateEndpointCertificate")
		request1 := expandRequestEndpointCertificateCreateEndpointCertificate(ctx, "", d)

		response1, _, err := client.EndpointCertificate.CreateEndpointCertificate(request1)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing CreateEndpointCertificate", err))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response")

		vvDirpath := d.Get("dirpath").(string)
		err = response1.SaveDownload(vvDirpath)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when downloading file", err))
			return diags
		}
		log.Printf("[DEBUG] Downloaded file %s", vvDirpath)

	}
	return diags
}

func expandRequestEndpointCertificateCreateEndpointCertificate(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCertificateCreateEndpointCertificate {
	request := isegosdk.RequestEndpointCertificateCreateEndpointCertificate{}
	request.ERSEndPointCert = expandRequestEndpointCertificateCreateEndpointCertificateERSEndPointCert(ctx, key, d)
	return &request
}

func expandRequestEndpointCertificateCreateEndpointCertificateERSEndPointCert(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCertificateCreateEndpointCertificateERSEndPointCert {
	request := isegosdk.RequestEndpointCertificateCreateEndpointCertificateERSEndPointCert{}
	if v, ok := d.GetOkExists("cert_template_name"); !isEmptyValue(reflect.ValueOf(d.Get("cert_template_name"))) && (ok || !reflect.DeepEqual(v, d.Get("cert_template_name"))) {
		request.CertTemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("format"); !isEmptyValue(reflect.ValueOf(d.Get("format"))) && (ok || !reflect.DeepEqual(v, d.Get("format"))) {
		request.Format = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("password"); !isEmptyValue(reflect.ValueOf(d.Get("password"))) && (ok || !reflect.DeepEqual(v, d.Get("password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("certificate_request"); !isEmptyValue(reflect.ValueOf(d.Get("certificate_request"))) && (ok || !reflect.DeepEqual(v, d.Get("certificate_request"))) {
		request.CertificateRequest = expandRequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest(ctx, key+".certificate_request.0", d)
	}
	return &request
}

func expandRequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest {
	request := isegosdk.RequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest{}
	if v, ok := d.GetOkExists("san"); !isEmptyValue(reflect.ValueOf(d.Get("san"))) && (ok || !reflect.DeepEqual(v, d.Get("san"))) {
		request.San = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("cn"); !isEmptyValue(reflect.ValueOf(d.Get("cn"))) && (ok || !reflect.DeepEqual(v, d.Get("cn"))) {
		request.Cn = interfaceToString(v)
	}
	return &request
}
