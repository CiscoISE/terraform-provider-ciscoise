package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEndpointCertificate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on EndpointCertificate.
- This resource allows the client to create an endpoint certificate.
`,

		CreateContext: resourceEndpointCertificateCreate,
		ReadContext:   resourceEndpointCertificateRead,
		DeleteContext: resourceEndpointCertificateDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dirpath": &schema.Schema{
							Description: `Directory absolute path in which to save the file.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"cert_template_name": &schema.Schema{
							Description: `Name of an Internal CA template`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"certificate_request": &schema.Schema{
							Description: `Key value map. Must have CN and SAN entries`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cn": &schema.Schema{
										Description: `Matches the requester's User Name, unless the Requester is an ERS Admin.
			ERS Admins are allowed to create requests for any CN`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"san": &schema.Schema{
										Description: `Valid MAC Address, delimited by '-'`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
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
							ForceNew: true,
						},
						"password": &schema.Schema{
							Description: `Protects the private key. Must have more than 8 characters, less than 15 characters,
			at least one upper case letter, at least one lower case letter, at least one digit,
			and can only contain [A-Z][a-z][0-9]_#`,
							Type:      schema.TypeString,
							Optional:  true,
							ForceNew:  true,
							Sensitive: true,
						},
					},
				},
			},
		},
	}
}

func resourceEndpointCertificateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning CreateEndpointCertificate create")
	log.Printf("[DEBUG] Missing CreateEndpointCertificate create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	log.Printf("[DEBUG] Selected method: CreateEndpointCertificate")
	request1 := expandRequestEndpointCertificateCreateEndpointCertificate(ctx, "parameters.0", d)

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
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceEndpointCertificateRead(ctx, d, m)
}

func resourceEndpointCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceEndpointCertificateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning EndpointCertificate delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing EndpointCertificate delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestEndpointCertificateCreateEndpointCertificate(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCertificateCreateEndpointCertificate {
	request := isegosdk.RequestEndpointCertificateCreateEndpointCertificate{}
	request.ERSEndPointCert = expandRequestEndpointCertificateCreateEndpointCertificateERSEndPointCert(ctx, key, d)
	return &request
}

func expandRequestEndpointCertificateCreateEndpointCertificateERSEndPointCert(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCertificateCreateEndpointCertificateERSEndPointCert {
	request := isegosdk.RequestEndpointCertificateCreateEndpointCertificateERSEndPointCert{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cert_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cert_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cert_template_name")))) {
		request.CertTemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".format")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".format")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".format")))) {
		request.Format = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_request")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_request")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_request")))) {
		request.CertificateRequest = expandRequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest(ctx, key+".certificate_request.0", d)
	}
	return &request
}

func expandRequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest {
	request := isegosdk.RequestEndpointCertificateCreateEndpointCertificateERSEndPointCertCertificateRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".san")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".san")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".san")))) {
		request.San = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cn")))) {
		request.Cn = interfaceToString(v)
	}
	return &request
}
