package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDuoMfa() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Duo-Mfa.

- Duo-MFA List of Duo-MFA configurations

- Duo-MFA Get the Duo-MFA configuration specified in the connectionName.
`,

		ReadContext: dataSourceDuoMfaRead,
		Schema: map[string]*schema.Schema{
			"connection_name": &schema.Schema{
				Description: `connectionName path parameter. This name is used to update, delete or retrieve the specific Duo-MFA configuration.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_configurations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"admin_api": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ikey": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"s_key": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"api_host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"authentication_api": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ikey": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"connection_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_sync": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"api_host_name": &schema.Schema{
							Description: `Duo API HostName`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"identity_sync": &schema.Schema{
							Description: `Name of the Identity Sync configuration`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name of the Duo-MFA configuration`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"provider": &schema.Schema{
							Description: `Name of the Mfa provider`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": &schema.Schema{
							Description: `Protocol type for which Mfa can be used`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDuoMfaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vConnectionName, okConnectionName := d.GetOk("connection_name")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okConnectionName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetMfa")

		response1, restyResp1, err := client.DuoMfa.GetMfa()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetMfa", err,
				"Failure at GetMfa, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDuoMfaGetMfaItemsResponse(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMfa response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetMfaByconnectionName")
		vvConnectionName := vConnectionName.(string)

		response2, restyResp2, err := client.DuoMfa.GetMfaByconnectionName(vvConnectionName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetMfaByconnectionName", err,
				"Failure at GetMfaByconnectionName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDuoMfaGetMfaByconnectionNameItemResponse(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMfaByconnectionName response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDuoMfaGetMfaItemsResponse(items *[]isegosdk.ResponseDuoMfaGetMfaResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["api_host_name"] = item.APIHostName
		respItem["identity_sync"] = item.IDentitySync
		respItem["name"] = item.Name
		respItem["provider"] = item.Provider
		respItem["type"] = item.Type
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDuoMfaGetMfaByconnectionNameItemResponse(item *isegosdk.ResponseDuoMfaGetMfaByconnectionNameResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["mfa"] = flattenDuoMfaGetMfaByconnectionNameItemResponseMfa(item.Mfa)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDuoMfaGetMfaByconnectionNameItemResponseMfa(item *isegosdk.ResponseDuoMfaGetMfaByconnectionNameResponseMfa) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["account_configurations"] = flattenDuoMfaGetMfaByconnectionNameItemResponseMfaAccountConfigurations(item.AccountConfigurations)
	respItem["connection_name"] = item.ConnectionName
	respItem["description"] = item.Description
	respItem["identity_sync"] = item.IDentitySync
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDuoMfaGetMfaByconnectionNameItemResponseMfaAccountConfigurations(item *isegosdk.ResponseDuoMfaGetMfaByconnectionNameResponseMfaAccountConfigurations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_api"] = flattenDuoMfaGetMfaByconnectionNameItemResponseMfaAccountConfigurationsAdminAPI(item.AdminAPI)
	respItem["api_host_name"] = item.APIHostName
	respItem["authentication_api"] = flattenDuoMfaGetMfaByconnectionNameItemResponseMfaAccountConfigurationsAuthenticationAPI(item.AuthenticationAPI)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDuoMfaGetMfaByconnectionNameItemResponseMfaAccountConfigurationsAdminAPI(item *isegosdk.ResponseDuoMfaGetMfaByconnectionNameResponseMfaAccountConfigurationsAdminAPI) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ikey"] = item.Ikey
	respItem["s_key"] = item.SKey

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDuoMfaGetMfaByconnectionNameItemResponseMfaAccountConfigurationsAuthenticationAPI(item *isegosdk.ResponseDuoMfaGetMfaByconnectionNameResponseMfaAccountConfigurationsAuthenticationAPI) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ikey"] = item.Ikey
	respItem["s_key"] = item.SKey

	return []map[string]interface{}{
		respItem,
	}

}
