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
func dataSourceSystemCertificateCreate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSystemCertificateCreateRead,
		Schema: map[string]*schema.Schema{
			"ers_local_cert_stub": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
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
							Type:     schema.TypeList,
							Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceSystemCertificateCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateSystemCertificate")
		request1 := expandRequestSystemCertificateCreateCreateSystemCertificate(ctx, "", d)

		response1, err := client.SystemCertificate.CreateSystemCertificate(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateSystemCertificate", err,
				"Failure at CreateSystemCertificate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

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
	if v, ok := d.GetOkExists("node_id"); !isEmptyValue(reflect.ValueOf(d.Get("node_id"))) && (ok || !reflect.DeepEqual(v, d.Get("node_id"))) {
		request.NodeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("ers_local_cert_stub"); !isEmptyValue(reflect.ValueOf(d.Get("ers_local_cert_stub"))) && (ok || !reflect.DeepEqual(v, d.Get("ers_local_cert_stub"))) {
		request.ErsLocalCertStub = expandRequestSystemCertificateCreateCreateSystemCertificateERSSystemCertificateErsLocalCertStub(ctx, key+".ers_local_cert_stub.0", d)
	}
	return &request
}

func expandRequestSystemCertificateCreateCreateSystemCertificateERSSystemCertificateErsLocalCertStub(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStub {
	request := isegosdk.RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStub{}
	if v, ok := d.GetOkExists("friendly_name"); !isEmptyValue(reflect.ValueOf(d.Get("friendly_name"))) && (ok || !reflect.DeepEqual(v, d.Get("friendly_name"))) {
		request.FriendlyName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("ers_subject_stub"); !isEmptyValue(reflect.ValueOf(d.Get("ers_subject_stub"))) && (ok || !reflect.DeepEqual(v, d.Get("ers_subject_stub"))) {
		request.ErsSubjectStub = expandRequestSystemCertificateCreateCreateSystemCertificateERSSystemCertificateErsLocalCertStubErsSubjectStub(ctx, key+".ers_subject_stub.0", d)
	}
	if v, ok := d.GetOkExists("key_length"); !isEmptyValue(reflect.ValueOf(d.Get("key_length"))) && (ok || !reflect.DeepEqual(v, d.Get("key_length"))) {
		request.KeyLength = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("xgrid_certificate"); !isEmptyValue(reflect.ValueOf(d.Get("xgrid_certificate"))) && (ok || !reflect.DeepEqual(v, d.Get("xgrid_certificate"))) {
		request.XgridCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("group_tag_dd"); !isEmptyValue(reflect.ValueOf(d.Get("group_tag_dd"))) && (ok || !reflect.DeepEqual(v, d.Get("group_tag_dd"))) {
		request.GroupTagDD = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("saml_certificate"); !isEmptyValue(reflect.ValueOf(d.Get("saml_certificate"))) && (ok || !reflect.DeepEqual(v, d.Get("saml_certificate"))) {
		request.SamlCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("key_type"); !isEmptyValue(reflect.ValueOf(d.Get("key_type"))) && (ok || !reflect.DeepEqual(v, d.Get("key_type"))) {
		request.KeyType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("digest"); !isEmptyValue(reflect.ValueOf(d.Get("digest"))) && (ok || !reflect.DeepEqual(v, d.Get("digest"))) {
		request.Digest = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("certificate_policies"); !isEmptyValue(reflect.ValueOf(d.Get("certificate_policies"))) && (ok || !reflect.DeepEqual(v, d.Get("certificate_policies"))) {
		request.CertificatePolicies = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("expiration_ttl"); !isEmptyValue(reflect.ValueOf(d.Get("expiration_ttl"))) && (ok || !reflect.DeepEqual(v, d.Get("expiration_ttl"))) {
		request.ExpirationTTL = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists("selected_expiration_ttl_unit"); !isEmptyValue(reflect.ValueOf(d.Get("selected_expiration_ttl_unit"))) && (ok || !reflect.DeepEqual(v, d.Get("selected_expiration_ttl_unit"))) {
		request.SelectedExpirationTTLUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("allow_wildcard_certs"); !isEmptyValue(reflect.ValueOf(d.Get("allow_wildcard_certs"))) && (ok || !reflect.DeepEqual(v, d.Get("allow_wildcard_certs"))) {
		request.AllowWildcardCerts = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("certificate_san_dns"); !isEmptyValue(reflect.ValueOf(d.Get("certificate_san_dns"))) && (ok || !reflect.DeepEqual(v, d.Get("certificate_san_dns"))) {
		request.CertificateSanDNS = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("certificate_san_ip"); !isEmptyValue(reflect.ValueOf(d.Get("certificate_san_ip"))) && (ok || !reflect.DeepEqual(v, d.Get("certificate_san_ip"))) {
		request.CertificateSanIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("certificate_san_uri"); !isEmptyValue(reflect.ValueOf(d.Get("certificate_san_uri"))) && (ok || !reflect.DeepEqual(v, d.Get("certificate_san_uri"))) {
		request.CertificateSanURI = interfaceToString(v)
	}
	return &request
}

func expandRequestSystemCertificateCreateCreateSystemCertificateERSSystemCertificateErsLocalCertStubErsSubjectStub(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStubErsSubjectStub {
	request := isegosdk.RequestSystemCertificateCreateSystemCertificateERSSystemCertificateErsLocalCertStubErsSubjectStub{}
	if v, ok := d.GetOkExists("common_name"); !isEmptyValue(reflect.ValueOf(d.Get("common_name"))) && (ok || !reflect.DeepEqual(v, d.Get("common_name"))) {
		request.CommonName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("organizational_unit_name"); !isEmptyValue(reflect.ValueOf(d.Get("organizational_unit_name"))) && (ok || !reflect.DeepEqual(v, d.Get("organizational_unit_name"))) {
		request.OrganizationalUnitName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("organization_name"); !isEmptyValue(reflect.ValueOf(d.Get("organization_name"))) && (ok || !reflect.DeepEqual(v, d.Get("organization_name"))) {
		request.OrganizationName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("locality_name"); !isEmptyValue(reflect.ValueOf(d.Get("locality_name"))) && (ok || !reflect.DeepEqual(v, d.Get("locality_name"))) {
		request.LocalityName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("state_or_province_name"); !isEmptyValue(reflect.ValueOf(d.Get("state_or_province_name"))) && (ok || !reflect.DeepEqual(v, d.Get("state_or_province_name"))) {
		request.StateOrProvinceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("country_name"); !isEmptyValue(reflect.ValueOf(d.Get("country_name"))) && (ok || !reflect.DeepEqual(v, d.Get("country_name"))) {
		request.CountryName = interfaceToString(v)
	}
	return &request
}
