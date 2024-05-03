package ciscoise

import (
	"context"
	"log"

	"reflect"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceDuoMfaTestconnection() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Duo-Mfa.

- Duo-MFA Verify the Auth and Admin API keys of the Duo Host.
`,

		CreateContext: resourceDuoMfaTestconnectionCreate,
		ReadContext:   resourceDuoMfaTestconnectionRead,
		DeleteContext: resourceDuoMfaTestconnectionDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
						"connection_name": &schema.Schema{
							Description: `connectionName path parameter. This name is used to retrieve secret keys for testing connection of the specified Duo-MFA configuration in case none are specified.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"admin_api": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ikey": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"s_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"api_host_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"authentication_api": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ikey": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"s_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceDuoMfaTestconnectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))

	vConnectionName := resourceItem["connection_name"]

	vvConnectionName := vConnectionName.(string)
	request1 := expandRequestDuoMfaTestconnectionTestConnection(ctx, "parameters.0", d)

	response1, err := client.DuoMfa.TestConnection(vvConnectionName, request1)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing LoadGroupsFromDomain", err,
			"Failure at LoadGroupsFromDomain, unexpected response", ""))
		return diags
	}

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting TestConnection response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags
}

func expandRequestDuoMfaTestconnectionTestConnection(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaTestConnection {
	request := isegosdk.RequestDuoMfaTestConnection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_api")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_api")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_api")))) {
		request.AdminAPI = expandRequestDuoMfaTestconnectionTestConnectionAdminAPI(ctx, key+".admin_api.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".api_host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".api_host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".api_host_name")))) {
		request.APIHostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_api")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_api")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_api")))) {
		request.AuthenticationAPI = expandRequestDuoMfaTestconnectionTestConnectionAuthenticationAPI(ctx, key+".authentication_api.0", d)
	}
	return &request
}

func expandRequestDuoMfaTestconnectionTestConnectionAdminAPI(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaTestConnectionAdminAPI {
	request := isegosdk.RequestDuoMfaTestConnectionAdminAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ikey")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ikey")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ikey")))) {
		request.Ikey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".s_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".s_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".s_key")))) {
		request.SKey = interfaceToString(v)
	}
	return &request
}

func expandRequestDuoMfaTestconnectionTestConnectionAuthenticationAPI(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaTestConnectionAuthenticationAPI {
	request := isegosdk.RequestDuoMfaTestConnectionAuthenticationAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ikey")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ikey")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ikey")))) {
		request.Ikey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".s_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".s_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".s_key")))) {
		request.SKey = interfaceToString(v)
	}
	return &request
}

func resourceDuoMfaTestconnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceDuoMfaTestconnectionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceDuoMfaTestconnectionRead(ctx, d, m)
}

func resourceDuoMfaTestconnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
