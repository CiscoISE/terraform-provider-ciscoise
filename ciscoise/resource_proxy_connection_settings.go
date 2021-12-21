package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProxyConnectionSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on proxy.

- The following functionalities are impacted by the proxy settings:


Partner Mobile Management

Endpoint Profiler Feed Service Update

Endpoint Posture Update

Endpoint Posture Agent Resources Download

CRL (Certificate Revocation List) Download

SMS Message Transmission

Social Login

Rest Auth Service Azure AD

pxGrid Cloud


`,

		CreateContext: resourceProxyConnectionSettingsCreate,
		ReadContext:   resourceProxyConnectionSettingsRead,
		UpdateContext: resourceProxyConnectionSettingsUpdate,
		DeleteContext: resourceProxyConnectionSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bypass_hosts": &schema.Schema{
							Description: `Bypass hosts for the proxy connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"fqdn": &schema.Schema{
							Description: `proxy IP address or DNS-resolvable host name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password": &schema.Schema{
							Description: `Password for the proxy connection`,
							Type:        schema.TypeString,
							Sensitive:   true,
							Computed:    true,
						},
						"password_required": &schema.Schema{
							Description: `Indicates whether password configuration is required for Proxy.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"port": &schema.Schema{
							Description: `Port for proxy connection. should be between 1 and 65535`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"user_name": &schema.Schema{
							Description: `User name for the proxy connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bypass_hosts": &schema.Schema{
							Description: `Bypass hosts for the proxy connection`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"fqdn": &schema.Schema{
							Description: `proxy IP address or DNS-resolvable host name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"password": &schema.Schema{
							Description: `Password for the proxy connection`,
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
						},
						"password_required": &schema.Schema{
							Description:  `Indicates whether password configuration is required for Proxy.`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"port": &schema.Schema{
							Description: `Port for proxy connection. should be between 1 and 65535`,
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"user_name": &schema.Schema{
							Description: `User name for the proxy connection`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceProxyConnectionSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["fqdn"] = interfaceToString(resourceItem["fqdn"])
	d.SetId(joinResourceID(resourceMap))
	return resourceProxyConnectionSettingsRead(ctx, d, m)
}

func resourceProxyConnectionSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetProxyConnection")

		response1, restyResp1, err := client.Proxy.GetProxyConnection()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetProxyConnection", err,
				"Failure at GetProxyConnection, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenProxyGetProxyConnectionItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetProxyConnection response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceProxyConnectionSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	// NOTE: Consider adding getAllItems and search function to get missing params

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation")
		request1 := expandRequestProxyConnectionSettingsUpdateProxyConnection(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Proxy.UpdateProxyConnection(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateProxyConnection", err, restyResp1.String(),
					"Failure at UpdateProxyConnection, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateProxyConnection", err,
				"Failure at UpdateProxyConnection, unexpected response", ""))
			return diags
		}
	}

	return resourceProxyConnectionSettingsRead(ctx, d, m)
}

func resourceProxyConnectionSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete ProxyConnectionSettings on Cisco ISE
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestProxyConnectionSettingsUpdateProxyConnection(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestProxyUpdateProxyConnection {
	request := isegosdk.RequestProxyUpdateProxyConnection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bypass_hosts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bypass_hosts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bypass_hosts")))) {
		request.BypassHosts = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_required")))) {
		request.PasswordRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
