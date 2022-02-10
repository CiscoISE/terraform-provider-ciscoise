package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLicensingRegistration() *schema.Resource {
	return &schema.Resource{
		Description: `It manages DEREGISTER, REGISTER, RENEW and UPDATE operations on License - registration information.

- License Configure registration information.
`,

		CreateContext: resourceLicensingRegistrationCreate,
		ReadContext:   resourceLicensingRegistrationRead,
		UpdateContext: resourceLicensingRegistrationUpdate,
		DeleteContext: resourceLicensingRegistrationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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

						"connection_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"registration_state": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssm_on_prem_server": &schema.Schema{
							Description: `If connection type is selected as SSM_ONPREM_SERVER, then  IP address or the hostname (or FQDN) of the SSM On-Prem server Host.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tier": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_type": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringHasValueFunc([]string{"", "HTTP_DIRECT", "PROXY", "SSM_ONPREM_SERVER", "TRANSPORT_GATEWAY"}),
						},
						"registration_type": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringHasValueFunc([]string{"", "DEREGISTER", "REGISTER", "RENEW", "UPDATE"}),
						},
						"ssm_on_prem_server": &schema.Schema{
							Description: `If connection type is selected as SSM_ONPREM_SERVER, then  IP address or the hostname (or FQDN) of the SSM On-Prem server Host.`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"tier": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validateStringHasValueFunc([]string{"", "ADVANTAGE", "DEVICEADMIN", "ESSENTIAL", "PREMIER", "VM"}),
							},
						},
						"token": &schema.Schema{
							Description: `token`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceLicensingRegistrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning LicenseRegistration create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestLicensingRegistrationCreateCreateRegistrationInfo(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	restyResp1, err := client.Licensing.CreateRegistrationInfo(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateRegistrationInfo", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateRegistrationInfo", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["connection_type"] = interfaceToString(resourceItem["connection_type"])
	resourceMap["registration_type"] = interfaceToString(resourceItem["registration_type"])
	resourceMap["ssm_on_prem_server"] = interfaceToString(resourceItem["ssm_on_prem_server"])
	resourceMap["tier"] = interfaceToString(resourceItem["tier"])
	d.SetId(joinResourceID(resourceMap))
	return resourceLicensingRegistrationRead(ctx, d, m)
}

func resourceLicensingRegistrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning LicenseRegistration read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	response1, restyResp1, err := client.Licensing.GetRegistrationInfo()

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenLicensingGetRegistrationInfoItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetRegistrationInfo response",
			err))
		return diags
	}
	if err := d.Set("parameters", remove_parameters(vItem1)); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetRegistrationInfo response to parameters",
			err))
		return diags
	}
	return diags
}

func resourceLicensingRegistrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning LicenseRegistration update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestLicensingRegistrationCreateCreateRegistrationInfo(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		restyResp1, err := client.Licensing.CreateRegistrationInfo(request1)
		if err != nil {
			if restyResp1 != nil {
				diags = append(diags, diagErrorWithResponse(
					"Failure when executing CreateRegistrationInfo", err, restyResp1.String()))
				return diags
			}
			diags = append(diags, diagError(
				"Failure when executing CreateRegistrationInfo", err))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}
	return resourceLicensingRegistrationRead(ctx, d, m)
}

func resourceLicensingRegistrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning LicenseRegistration delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing LicenseRegistration delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestLicensingRegistrationCreateCreateRegistrationInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLicensingCreateRegistrationInfo {
	request := isegosdk.RequestLicensingCreateRegistrationInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_type")))) {
		request.ConnectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".registration_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".registration_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".registration_type")))) {
		request.RegistrationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssm_on_prem_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssm_on_prem_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssm_on_prem_server")))) {
		request.SsmOnPremServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tier")))) {
		request.Tier = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".token")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".token")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".token")))) {
		request.Token = interfaceToString(v)
	}
	return &request
}
