package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSystemCertificateCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on SystemCertificate.

- This data source action allows the client to create a system certificate.
`,

		ReadContext: dataSourceSystemCertificateCreateRead,
		Schema: map[string]*schema.Schema{
			"ers_local_cert_stub": &schema.Schema{
				Description: `Inputs for certificate creation`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allow_wildcard_certs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"certificate_policies": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"certificate_san_dns": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"certificate_san_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"certificate_san_uri": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"digest": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ers_subject_stub": &schema.Schema{
							Description: `Subject data of certificate`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"common_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"country_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"locality_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"organization_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"organizational_unit_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"state_or_province_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"expiration_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_tag_dd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_length": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"saml_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"selected_expiration_ttl_unit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"xgrid_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_id": &schema.Schema{
				Description: `NodeId of Cisco ISE application`,
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceSystemCertificateCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CreateSystemCertificate")
		request1 := expandRequestSystemCertificateCreateCreateSystemCertificate(ctx, "", d)

		response1, err := client.SystemCertificate.CreateSystemCertificate(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateSystemCertificate", err,
				"Failure at CreateSystemCertificate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateSystemCertificate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSystemCertificateCreateCreateSystemCertificate(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSystemCertificateCreateSystemCertificate {
	request := isegosdk.RequestSystemCertificateCreateSystemCertificate{}
	request.ERSSystemCertificate = expandRequestSystemCertificateCreateCreateSystemCertificateERSSystemCertificate(ctx, key, d)
	return &request
}

func expandRequestSystemCertificateCreateCreateSystemCertificateERSSystemCertificate(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSystemCertificateCreateSystemCertificateERSSystemCertificate {
	request := isegosdk.RequestSystemCertificateCreateSystemCertificateERSSystemCertificate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".node_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".node_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".node_id")))) {
		request.NodeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ers_local_cert_stub")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ers_local_cert_stub")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ers_local_cert_stub")))) {
		request.ErsLocalCertStub = expandRequestSystemCertificateCreateCreateSystemCertificateERSSystemCertificateErsLocalCertStub(ctx, key+".ers_local_cert_stub.0", d)
	}
	return &request
}

func expandRequestSystemCertificateCreateCreateSystemCertificateERSSystemCertificateErsLocalCertStub(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStub {
	request := isegosdk.RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStub{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".friendly_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".friendly_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".friendly_name")))) {
		request.FriendlyName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ers_subject_stub")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ers_subject_stub")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ers_subject_stub")))) {
		request.ErsSubjectStub = expandRequestSystemCertificateCreateCreateSystemCertificateERSSystemCertificateErsLocalCertStubErsSubjectStub(ctx, key+".ers_subject_stub.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key_length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key_length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key_length")))) {
		request.KeyLength = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".xgrid_certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".xgrid_certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".xgrid_certificate")))) {
		request.XgridCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_tag_dd")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_tag_dd")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_tag_dd")))) {
		request.GroupTagDD = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".saml_certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".saml_certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".saml_certificate")))) {
		request.SamlCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key_type")))) {
		request.KeyType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".digest")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".digest")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".digest")))) {
		request.Digest = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_policies")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_policies")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_policies")))) {
		request.CertificatePolicies = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".expiration_ttl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".expiration_ttl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".expiration_ttl")))) {
		request.ExpirationTTL = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selected_expiration_ttl_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selected_expiration_ttl_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selected_expiration_ttl_unit")))) {
		request.SelectedExpirationTTLUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_wildcard_certs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_wildcard_certs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_wildcard_certs")))) {
		request.AllowWildcardCerts = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_san_dns")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_san_dns")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_san_dns")))) {
		request.CertificateSanDNS = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_san_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_san_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_san_ip")))) {
		request.CertificateSanIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_san_uri")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_san_uri")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_san_uri")))) {
		request.CertificateSanURI = interfaceToString(v)
	}
	return &request
}

func expandRequestSystemCertificateCreateCreateSystemCertificateERSSystemCertificateErsLocalCertStubErsSubjectStub(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStubErsSubjectStub {
	request := isegosdk.RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStubErsSubjectStub{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".common_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".common_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".common_name")))) {
		request.CommonName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".organizational_unit_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".organizational_unit_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".organizational_unit_name")))) {
		request.OrganizationalUnitName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".organization_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".organization_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".organization_name")))) {
		request.OrganizationName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".locality_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".locality_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".locality_name")))) {
		request.LocalityName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state_or_province_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state_or_province_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state_or_province_name")))) {
		request.StateOrProvinceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".country_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".country_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".country_name")))) {
		request.CountryName = interfaceToString(v)
	}
	return &request
}
