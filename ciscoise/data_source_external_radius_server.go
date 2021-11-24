package ciscoise

import (
	"context"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceExternalRadiusServer() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ExternalRADIUSServer.

- This data source allows the client to get an external RADIUS server by name.

- This data source allows the client to get an external RADIUS server by ID.

- This data source allows the client to get all the external RADIUS servers.
`,

		ReadContext: dataSourceExternalRadiusServerRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"page": &schema.Schema{
				Description: `page query parameter. Page number`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. Number of objects returned per page`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accounting_port": &schema.Schema{
							Description: `Valid Range 1 to 65535`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"authentication_port": &schema.Schema{
							Description: `Valid Range 1 to 65535`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"authenticator_key": &schema.Schema{
							Description: `The authenticatorKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty.
The maximum length is 20 ASCII characters or 40 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_key_wrap": &schema.Schema{
							Description: `KeyWrap may only be enabled if it is supported on the device.
When running in FIPS mode this option should be enabled for such devices`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"encryption_key": &schema.Schema{
							Description: `The encryptionKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty.
The maximum length is 16 ASCII characters or 32 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_ip": &schema.Schema{
							Description: `The IP of the host - must be a valid IPV4 address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_input_format": &schema.Schema{
							Description: `Specifies the format of the input for fields 'encryptionKey' and 'authenticatorKey'.
Allowed Values:
- ASCII
- HEXADECIMAL`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `Resource Name. Allowed charactera are alphanumeric and _ (underscore).`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"proxy_timeout": &schema.Schema{
							Description: `Valid Range 1 to 600`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"retries": &schema.Schema{
							Description: `Valid Range 1 to 9`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"shared_secret": &schema.Schema{
							Description: `Shared secret maximum length is 128 characters`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"timeout": &schema.Schema{
							Description: `Valid Range 1 to 120`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accounting_port": &schema.Schema{
							Description: `Valid Range 1 to 65535`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"authentication_port": &schema.Schema{
							Description: `Valid Range 1 to 65535`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"authenticator_key": &schema.Schema{
							Description: `The authenticatorKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty.
The maximum length is 20 ASCII characters or 40 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_key_wrap": &schema.Schema{
							Description: `KeyWrap may only be enabled if it is supported on the device.
When running in FIPS mode this option should be enabled for such devices`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"encryption_key": &schema.Schema{
							Description: `The encryptionKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty.
The maximum length is 16 ASCII characters or 32 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_ip": &schema.Schema{
							Description: `The IP of the host - must be a valid IPV4 address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_input_format": &schema.Schema{
							Description: `Specifies the format of the input for fields 'encryptionKey' and 'authenticatorKey'.
Allowed Values:
- ASCII
- HEXADECIMAL`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `Resource Name. Allowed charactera are alphanumeric and _ (underscore).`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"proxy_timeout": &schema.Schema{
							Description: `Valid Range 1 to 600`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"retries": &schema.Schema{
							Description: `Valid Range 1 to 9`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"shared_secret": &schema.Schema{
							Description: `Shared secret maximum length is 128 characters`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"timeout": &schema.Schema{
							Description: `Valid Range 1 to 120`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceExternalRadiusServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %q", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetExternalRadiusServer")
		queryParams1 := isegosdk.GetExternalRadiusServerQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, restyResp1, err := client.ExternalRadiusServer.GetExternalRadiusServer(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExternalRadiusServer", err,
				"Failure at GetExternalRadiusServer, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseExternalRadiusServerGetExternalRadiusServerSearchResultResources
		for response1.SearchResult != nil && response1.SearchResult.Resources != nil && len(*response1.SearchResult.Resources) > 0 {
			items1 = append(items1, *response1.SearchResult.Resources...)
			if response1.SearchResult.NextPage != nil && response1.SearchResult.NextPage.Rel == "next" {
				href := response1.SearchResult.NextPage.Href
				page, size, err := getNextPageAndSizeParams(href)
				if err != nil {
					break
				}
				queryParams1.Page = page
				queryParams1.Size = size
				response1, _, err = client.ExternalRadiusServer.GetExternalRadiusServer(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenExternalRadiusServerGetExternalRadiusServerItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExternalRadiusServer response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetExternalRadiusServerByName")
		vvName := vName.(string)

		response2, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExternalRadiusServerByName", err,
				"Failure at GetExternalRadiusServerByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemName2 := flattenExternalRadiusServerGetExternalRadiusServerByNameItemName(response2.ExternalRadiusServer)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExternalRadiusServerByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetExternalRadiusServerByID")
		vvID := vID.(string)

		response3, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExternalRadiusServerByID", err,
				"Failure at GetExternalRadiusServerByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response3))

		vItemID3 := flattenExternalRadiusServerGetExternalRadiusServerByIDItemID(response3.ExternalRadiusServer)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExternalRadiusServerByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenExternalRadiusServerGetExternalRadiusServerItems(items *[]isegosdk.ResponseExternalRadiusServerGetExternalRadiusServerSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenExternalRadiusServerGetExternalRadiusServerItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenExternalRadiusServerGetExternalRadiusServerItemsLink(item *isegosdk.ResponseExternalRadiusServerGetExternalRadiusServerSearchResultResourcesLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenExternalRadiusServerGetExternalRadiusServerByNameItemName(item *isegosdk.ResponseExternalRadiusServerGetExternalRadiusServerByNameExternalRadiusServer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["host_ip"] = item.HostIP
	respItem["shared_secret"] = item.SharedSecret
	respItem["enable_key_wrap"] = boolPtrToString(item.EnableKeyWrap)
	respItem["encryption_key"] = item.EncryptionKey
	respItem["authenticator_key"] = item.AuthenticatorKey
	respItem["key_input_format"] = item.KeyInputFormat
	respItem["authentication_port"] = item.AuthenticationPort
	respItem["accounting_port"] = item.AccountingPort
	respItem["timeout"] = item.Timeout
	respItem["retries"] = item.Retries
	respItem["proxy_timeout"] = item.ProxyTimeout
	respItem["link"] = flattenExternalRadiusServerGetExternalRadiusServerByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenExternalRadiusServerGetExternalRadiusServerByNameItemNameLink(item *isegosdk.ResponseExternalRadiusServerGetExternalRadiusServerByNameExternalRadiusServerLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenExternalRadiusServerGetExternalRadiusServerByIDItemID(item *isegosdk.ResponseExternalRadiusServerGetExternalRadiusServerByIDExternalRadiusServer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["host_ip"] = item.HostIP
	respItem["shared_secret"] = item.SharedSecret
	respItem["enable_key_wrap"] = boolPtrToString(item.EnableKeyWrap)
	respItem["encryption_key"] = item.EncryptionKey
	respItem["authenticator_key"] = item.AuthenticatorKey
	respItem["key_input_format"] = item.KeyInputFormat
	respItem["authentication_port"] = item.AuthenticationPort
	respItem["accounting_port"] = item.AccountingPort
	respItem["timeout"] = item.Timeout
	respItem["retries"] = item.Retries
	respItem["proxy_timeout"] = item.ProxyTimeout
	respItem["link"] = flattenExternalRadiusServerGetExternalRadiusServerByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenExternalRadiusServerGetExternalRadiusServerByIDItemIDLink(item *isegosdk.ResponseExternalRadiusServerGetExternalRadiusServerByIDExternalRadiusServerLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
