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
func dataSourceLicensingRegistrationCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Licensing.

- License Configure registration information.
`,

		ReadContext: dataSourceLicensingRegistrationCreateRead,
		Schema: map[string]*schema.Schema{
			"connection_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"registration_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
					Type: schema.TypeString,
				},
			},
			"token": &schema.Schema{
				Description: `token`,
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceLicensingRegistrationCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateRegistrationInfo")
		request1 := expandRequestLicensingRegistrationCreateCreateRegistrationInfo(ctx, "", d)

		response1, err := client.Licensing.CreateRegistrationInfo(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateRegistrationInfo", err,
				"Failure at CreateRegistrationInfo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateRegistrationInfo response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
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
