package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBindSignedCertificate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Certificates.
- Bind CA Signed Certificate.

NOTE:
This resource requires an existing Certificate Signing Request, and the root certificate must already be
trusted.

NOTE:
The certificate may have a validity period longer than 398 days. It may be untrusted by many browsers.

NOTE:
Request parameters accepting True and False as input can be replaced by 1 and 0 respectively.
`,
		CreateContext: resourceBindSignedCertificateCreate,
		ReadContext:   resourceBindSignedCertificateRead,
		DeleteContext: resourceBindSignedCertificateDelete,

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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin": &schema.Schema{
							Description:  ` Use certificate to authenticate the Cisco ISE Admin Portal`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_extended_validity": &schema.Schema{
							Description:  `Allow import of certificates with validity greater than 398 days (required)`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_out_of_date_cert": &schema.Schema{
							Description:  `Allow out of date certificates (required)`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_replacement_of_certificates": &schema.Schema{
							Description:  `Allow Replacement of certificates (required)`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_replacement_of_portal_group_tag": &schema.Schema{
							Description:  `Allow Replacement of Portal Group Tag (required)`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"data": &schema.Schema{
							Description: `Signed certificate data (required)`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"eap": &schema.Schema{
							Description:  `Use certificate for EAP protocols that use SSL/TLS tunneling`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"host_name": &schema.Schema{
							Description: `Name of Host whose CSR ID has been provided (required)`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"id": &schema.Schema{
							Description: `ID of the generated CSR (required)`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"ims": &schema.Schema{
							Description:  `Use certificate for the Cisco ISE Messaging Service`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"name": &schema.Schema{
							Description: `Friendly Name of the certificate`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"portal": &schema.Schema{
							Description:  `Use for portal`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"portal_group_tag": &schema.Schema{
							Description: `Set Group tag`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"pxgrid": &schema.Schema{
							Description:  `Use certificate for the pxGrid Controller`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"radius": &schema.Schema{
							Description:  `Use certificate for the RADSec server`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"saml": &schema.Schema{
							Description:  `Use certificate for SAML Signing`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"validate_certificate_extensions": &schema.Schema{
							Description:  `Validate Certificate Extensions`,
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

func resourceBindSignedCertificateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BindCsr create")
	log.Printf("[DEBUG] Missing BindCsr create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	request1 := expandRequestBindSignedCertificateBindCsr(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, restyResp1, err := client.Certificates.BindCsr(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing BindCsr", err, restyResp1.String(),
				"Failure at BindCsr, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BindCsr", err,
			"Failure at BindCsr, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenCertificatesBindCsrItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BindCsr response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceBindSignedCertificateRead(ctx, d, m)
}

func resourceBindSignedCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceBindSignedCertificateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BindSignedCertificate delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing BindSignedCertificate delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestBindSignedCertificateBindCsr(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesBindCsr {
	request := isegosdk.RequestCertificatesBindCsr{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin")))) {
		request.Admin = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_extended_validity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_extended_validity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_extended_validity")))) {
		request.AllowExtendedValidity = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_out_of_date_cert")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_out_of_date_cert")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_out_of_date_cert")))) {
		request.AllowOutOfDateCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_replacement_of_certificates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_replacement_of_certificates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_replacement_of_certificates")))) {
		request.AllowReplacementOfCertificates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_replacement_of_portal_group_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_replacement_of_portal_group_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_replacement_of_portal_group_tag")))) {
		request.AllowReplacementOfPortalGroupTag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap")))) {
		request.Eap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ims")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ims")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ims")))) {
		request.Ims = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal")))) {
		request.Portal = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_group_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_group_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_group_tag")))) {
		request.PortalGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pxgrid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pxgrid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pxgrid")))) {
		request.Pxgrid = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radius")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radius")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radius")))) {
		request.Radius = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".saml")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".saml")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".saml")))) {
		request.Saml = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".validate_certificate_extensions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".validate_certificate_extensions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".validate_certificate_extensions")))) {
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
