package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProxyConnectionSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on proxy.

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

		ReadContext: dataSourceProxyConnectionSettingsRead,
		Schema: map[string]*schema.Schema{
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
		},
	}
}

func dataSourceProxyConnectionSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenProxyGetProxyConnectionItem(item *isegosdk.ResponseProxyGetProxyConnectionResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["bypass_hosts"] = item.BypassHosts
	respItem["fqdn"] = item.Fqdn
	respItem["password"] = item.Password
	respItem["password_required"] = boolPtrToString(item.PasswordRequired)
	respItem["port"] = item.Port
	respItem["user_name"] = item.UserName
	return []map[string]interface{}{
		respItem,
	}
}
