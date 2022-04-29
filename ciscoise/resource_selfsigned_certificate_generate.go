package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSelfsignedCertificateGenerate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Certificates.
- Generate Self-signed Certificate
NOTE:
The certificate may have a validity period longer than 398 days. It may be untrusted by many browsers.
NOTE:
Request parameters accepting True and False as input can be replaced by 1 and 0 respectively.
NOTE:
Wildcard certificate and SAML certificate can be generated only on PPAN or Standalone
`,

		CreateContext: resourceSelfsignedCertificateGenerateCreate,
		ReadContext:   resourceSelfsignedCertificateGenerateRead,
		DeleteContext: resourceSelfsignedCertificateGenerateDelete,

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
							Description: `ID of the generated sefl signed system certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Response message on generation of self-signed system certificate`,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin": &schema.Schema{
							Description:  `Use certificate to authenticate the Cisco ISE Admin Portal`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_extended_validity": &schema.Schema{
							Description:  `Allow generation of self-signed certificate with validity greater than 398 days`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_portal_tag_transfer_for_same_subject": &schema.Schema{
							Description:  `Allow overwriting the portal tag from matching certificate of same subject`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_replacement_of_certificates": &schema.Schema{
							Description:  `Allow Replacement of certificates`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_replacement_of_portal_group_tag": &schema.Schema{
							Description:  `Allow Replacement of Portal Group Tag`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_role_transfer_for_same_subject": &schema.Schema{
							Description:  `Allow transfer of roles for certificate with matching subject`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_san_dns_bad_name": &schema.Schema{
							Description:  `Allow usage of SAN DNS Bad name`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_san_dns_non_resolvable": &schema.Schema{
							Description:  `Allow use of non resolvable Common Name or SAN Values`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"allow_wild_card_certificates": &schema.Schema{
							Description:  `Allow Wildcard Certificates`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"certificate_policies": &schema.Schema{
							Description: `Certificate Policies`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"digest_type": &schema.Schema{
							Description: `Digest to sign with`,
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
						"expiration_ttl": &schema.Schema{
							Description: `Certificate expiration value`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
						},
						"expiration_ttl_unit": &schema.Schema{
							Description: `Certificate expiration unit`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"host_name": &schema.Schema{
							Description: `Hostname of the Cisco ISE node in which self-signed certificate should be generated.`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"key_length": &schema.Schema{
							Description: `Bit size of public key`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"key_type": &schema.Schema{
							Description: `Algorithm to use for certificate public key creation`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"name": &schema.Schema{
							Description: `Friendly name of the certificate.`,
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
						"san_dns": &schema.Schema{
							Description: `Array of SAN (Subject Alternative Name) DNS entries`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"san_ip": &schema.Schema{
							Description: `Array of SAN IP entries`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"san_uri": &schema.Schema{
							Description: `Array of SAN URI entries`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"subject_city": &schema.Schema{
							Description: `Certificate city or locality (L)`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"subject_common_name": &schema.Schema{
							Description: `Certificate common name (CN)`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"subject_country": &schema.Schema{
							Description: `Certificate country (C)`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"subject_org": &schema.Schema{
							Description: `Certificate organization (O)`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"subject_org_unit": &schema.Schema{
							Description: `Certificate organizational unit (OU)`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"subject_state": &schema.Schema{
							Description: `Certificate state (ST)`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSelfsignedCertificateGenerateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GenerateSelfSignedCertificate create")
	log.Printf("[DEBUG] Missing GenerateSelfSignedCertificate create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	request1 := expandRequestSelfsignedCertificateGenerateGenerateSelfSignedCertificate(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, restyResp1, err := client.Certificates.GenerateSelfSignedCertificate(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing GenerateSelfSignedCertificate", err, restyResp1.String(),
				"Failure at GenerateSelfSignedCertificate, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GenerateSelfSignedCertificate", err,
			"Failure at GenerateSelfSignedCertificate, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenCertificatesGenerateSelfSignedCertificateItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GenerateSelfSignedCertificate response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceSelfsignedCertificateGenerateRead(ctx, d, m)
}

func resourceSelfsignedCertificateGenerateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSelfsignedCertificateGenerateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SelfsignedCertificateGenerate delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing SelfsignedCertificateGenerate delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestSelfsignedCertificateGenerateGenerateSelfSignedCertificate(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesGenerateSelfSignedCertificate {
	request := isegosdk.RequestCertificatesGenerateSelfSignedCertificate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin")))) {
		request.Admin = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_extended_validity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_extended_validity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_extended_validity")))) {
		request.AllowExtendedValidity = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_portal_tag_transfer_for_same_subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_portal_tag_transfer_for_same_subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_portal_tag_transfer_for_same_subject")))) {
		request.AllowPortalTagTransferForSameSubject = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_replacement_of_certificates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_replacement_of_certificates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_replacement_of_certificates")))) {
		request.AllowReplacementOfCertificates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_replacement_of_portal_group_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_replacement_of_portal_group_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_replacement_of_portal_group_tag")))) {
		request.AllowReplacementOfPortalGroupTag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_role_transfer_for_same_subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_role_transfer_for_same_subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_role_transfer_for_same_subject")))) {
		request.AllowRoleTransferForSameSubject = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_san_dns_bad_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_san_dns_bad_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_san_dns_bad_name")))) {
		request.AllowSanDNSBadName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_san_dns_non_resolvable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_san_dns_non_resolvable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_san_dns_non_resolvable")))) {
		request.AllowSanDNSNonResolvable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_wild_card_certificates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_wild_card_certificates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_wild_card_certificates")))) {
		request.AllowWildCardCertificates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_policies")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_policies")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_policies")))) {
		request.CertificatePolicies = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".digest_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".digest_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".digest_type")))) {
		request.DigestType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap")))) {
		request.Eap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".expiration_ttl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".expiration_ttl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".expiration_ttl")))) {
		request.ExpirationTTL = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".expiration_ttl_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".expiration_ttl_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".expiration_ttl_unit")))) {
		request.ExpirationTTLUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key_length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key_length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key_length")))) {
		request.KeyLength = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key_type")))) {
		request.KeyType = interfaceToString(v)
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".san_dns")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".san_dns")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".san_dns")))) {
		request.SanDNS = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".san_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".san_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".san_ip")))) {
		request.SanIP = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".san_uri")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".san_uri")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".san_uri")))) {
		request.SanURI = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject_city")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject_city")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject_city")))) {
		request.SubjectCity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject_common_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject_common_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject_common_name")))) {
		request.SubjectCommonName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject_country")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject_country")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject_country")))) {
		request.SubjectCountry = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject_org")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject_org")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject_org")))) {
		request.SubjectOrg = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject_org_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject_org_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject_org_unit")))) {
		request.SubjectOrgUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject_state")))) {
		request.SubjectState = interfaceToString(v)
	}
	return &request
}

func flattenCertificatesGenerateSelfSignedCertificateItem(item *isegosdk.ResponseCertificatesGenerateSelfSignedCertificateResponse) []map[string]interface{} {
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
