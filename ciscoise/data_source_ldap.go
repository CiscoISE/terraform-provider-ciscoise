package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLdap() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ldap.

- Get-All

- Get-By-Id
`,

		ReadContext: dataSourceLdapRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"page": &schema.Schema{
				Description: `page query parameter. Page Number (0...N)`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. Items by Page`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Description: `List of Attributes`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"connection_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"always_access_primary_first": &schema.Schema{
										Description: `alwaysAccessPrimaryFirst`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"failback_retry_delay": &schema.Schema{
										Description: `failbackRetryDelay`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"failover_to_secondary": &schema.Schema{
										Description: `failoverToSecondary`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"ldap_node_data": &schema.Schema{
										Description: `ldapNodeData`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"primary_hostname": &schema.Schema{
													Description: `ipaddress`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"primary_port": &schema.Schema{
													Description: `primaryPort`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"secondary_hostname": &schema.Schema{
													Description: `ipaddress`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"secondary_port": &schema.Schema{
													Description: `secondaryPort`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
									"primary_server": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"admin_dn": &schema.Schema{
													Description: `adminDN`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"admin_password": &schema.Schema{
													Description: `adminPassword`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"enable_force_reconnect": &schema.Schema{
													Description: `enableForceReconnect`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"enable_secure_connection": &schema.Schema{
													Description: `enableSecureConnection`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"enable_server_identity_check": &schema.Schema{
													Description: `enableServerIdentityCheck`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"force_reconnect_time": &schema.Schema{
													Description: `forceReconnectTime`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"host_name": &schema.Schema{
													Description: `hostName`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"issuer_cacertificate": &schema.Schema{
													Description: `issuerCACertificate`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"max_connections": &schema.Schema{
													Description: `maxConnections`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"port": &schema.Schema{
													Description: `port`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"server_timeout": &schema.Schema{
													Description: `serverTimeout`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"trust_certificate": &schema.Schema{
													Description: `trustCertificate`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"use_admin_access": &schema.Schema{
													Description: `useAdminAccess`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"secondary_server": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"admin_dn": &schema.Schema{
													Description: `adminDN`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"admin_password": &schema.Schema{
													Description: `adminPassword`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"enable_force_reconnect": &schema.Schema{
													Description: `enableForceReconnect`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"enable_secure_connection": &schema.Schema{
													Description: `enableSecureConnection`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"enable_server_identity_check": &schema.Schema{
													Description: `enableServerIdentityCheck`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"force_reconnect_time": &schema.Schema{
													Description: `forceReconnectTime`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"host_name": &schema.Schema{
													Description: `hostName`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"issuer_cacertificate": &schema.Schema{
													Description: `issuerCACertificate`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"max_connections": &schema.Schema{
													Description: `maxConnections`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"port": &schema.Schema{
													Description: `port`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"server_timeout": &schema.Schema{
													Description: `serverTimeout`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"trust_certificate": &schema.Schema{
													Description: `trustCertificate`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"use_admin_access": &schema.Schema{
													Description: `useAdminAccess`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"specify_server_for_each_ise_node": &schema.Schema{
										Description: `specifyServerForEachISENode`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"directory_organization": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"group_directory_subtree": &schema.Schema{
										Description: `groupDirectorySubtree`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"mac_format": &schema.Schema{
										Description: `macFormat`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},
									"prefix_separator": &schema.Schema{
										Description: `prefixSeparator`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"strip_prefix": &schema.Schema{
										Description: `stripPrefix`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"strip_suffix": &schema.Schema{
										Description: `stripSuffix`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"suffix_separator": &schema.Schema{
										Description: `suffixSeparator`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"user_directory_subtree": &schema.Schema{
										Description: `userDirectorySubtree`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"enable_password_change_ldap": &schema.Schema{
							Description: `enablePasswordChangeLDAP`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"general_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"certificate": &schema.Schema{
										Description: `certificate`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"group_map_attribute_name": &schema.Schema{
										Description: `groupMapAttributeName`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"group_member_reference": &schema.Schema{
										Description: `groupMemberReference`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},
									"group_name_attribute": &schema.Schema{
										Description: `groupNameAttribute`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"group_object_class": &schema.Schema{
										Description: `groupObjectClass`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"schema": &schema.Schema{
										Description: `schema`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},
									"user_info_attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"additional_attribute": &schema.Schema{
													Description: `additionalAttribute`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"country": &schema.Schema{
													Description: `country`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"department": &schema.Schema{
													Description: `department`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"email": &schema.Schema{
													Description: `email`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"first_name": &schema.Schema{
													Description: `firstName`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"job_title": &schema.Schema{
													Description: `jobTitle`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"last_name": &schema.Schema{
													Description: `lastName`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"locality": &schema.Schema{
													Description: `locality`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"organizational_unit": &schema.Schema{
													Description: `organizationalUnit`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"state_or_province": &schema.Schema{
													Description: `stateOrProvince`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"street_address": &schema.Schema{
													Description: `streetAddress`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"telephone": &schema.Schema{
													Description: `telephone`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"user_name_attribute": &schema.Schema{
										Description: `userNameAttribute`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"user_object_class": &schema.Schema{
										Description: `userObjectClass.`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"groups_names": &schema.Schema{
										Description: `List of groups`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `name`,
							Type:        schema.TypeString,
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
							Description: `description`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Id description`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `name description`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLdapRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vID, okID := d.GetOk("id")

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okPage, okSize}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetLdap")
		queryParams1 := isegosdk.GetLdapQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, restyResp1, err := client.Ldap.GetLdap(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetLdap", err,
				"Failure at GetLdap, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseLdapGetLdapSearchResultResources
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
				response1, _, err = client.Ldap.GetLdap(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenLdapGetLdapItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLdap response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetLdapid")
		vvID := vID.(string)

		response2, restyResp2, err := client.Ldap.GetLdapid(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetLdapid", err,
				"Failure at GetLdapid, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenLdapGetLdapidItem(response2.ERSLdap)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLdapid response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLdapGetLdapItems(items *[]isegosdk.ResponseLdapGetLdapSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenLdapGetLdapItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLdapGetLdapItemsLink(item *isegosdk.ResponseLdapGetLdapSearchResultResourcesLink) []map[string]interface{} {
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

func flattenLdapGetLdapidItem(item *isegosdk.ResponseLdapGetLdapidERSLdap) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["general_settings"] = flattenLdapGetLdapidItemGeneralSettings(item.GeneralSettings)
	respItem["connection_settings"] = flattenLdapGetLdapidItemConnectionSettings(item.ConnectionSettings)
	respItem["directory_organization"] = flattenLdapGetLdapidItemDirectoryOrganization(item.DirectoryOrganization)
	respItem["groups"] = flattenLdapGetLdapidItemGroups(item.Groups)
	respItem["attributes"] = flattenLdapGetLdapidItemAttributes(item.Attributes)
	respItem["enable_password_change_ldap"] = boolPtrToString(item.EnablePasswordChangeLDAP)
	respItem["name"] = item.Name
	respItem["id"] = item.ID
	respItem["description"] = item.Description
	respItem["link"] = flattenLdapGetLdapidItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenLdapGetLdapidItemGeneralSettings(item *isegosdk.ResponseLdapGetLdapidERSLdapGeneralSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["user_object_class"] = item.UserObjectClass
	respItem["user_name_attribute"] = item.UserNameAttribute
	respItem["group_name_attribute"] = item.GroupNameAttribute
	respItem["group_object_class"] = item.GroupObjectClass
	respItem["group_map_attribute_name"] = item.GroupMapAttributeName
	respItem["certificate"] = item.Certificate
	respItem["schema"] = flattenLdapGetLdapidItemGeneralSettingsSchema(item.Schema)
	respItem["group_member_reference"] = flattenLdapGetLdapidItemGeneralSettingsGroupMemberReference(item.GroupMemberReference)
	respItem["user_info_attributes"] = flattenLdapGetLdapidItemGeneralSettingsUserInfoAttributes(item.UserInfoAttributes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLdapGetLdapidItemGeneralSettingsSchema(item *isegosdk.ResponseLdapGetLdapidERSLdapGeneralSettingsSchema) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenLdapGetLdapidItemGeneralSettingsGroupMemberReference(item *isegosdk.ResponseLdapGetLdapidERSLdapGeneralSettingsGroupMemberReference) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenLdapGetLdapidItemGeneralSettingsUserInfoAttributes(item *isegosdk.ResponseLdapGetLdapidERSLdapGeneralSettingsUserInfoAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["first_name"] = item.FirstName
	respItem["department"] = item.Department
	respItem["last_name"] = item.LastName
	respItem["organizational_unit"] = item.OrganizationalUnit
	respItem["job_title"] = item.JobTitle
	respItem["locality"] = item.Locality
	respItem["email"] = item.Email
	respItem["state_or_province"] = item.StateOrProvince
	respItem["telephone"] = item.Telephone
	respItem["country"] = item.Country
	respItem["street_address"] = item.StreetAddress
	respItem["additional_attribute"] = item.AdditionalAttribute

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLdapGetLdapidItemConnectionSettings(item *isegosdk.ResponseLdapGetLdapidERSLdapConnectionSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["primary_server"] = flattenLdapGetLdapidItemConnectionSettingsPrimaryServer(item.PrimaryServer)
	respItem["secondary_server"] = flattenLdapGetLdapidItemConnectionSettingsSecondaryServer(item.SecondaryServer)
	respItem["ldap_node_data"] = flattenLdapGetLdapidItemConnectionSettingsLdapNodeData(item.LdapNodeData)
	respItem["failover_to_secondary"] = boolPtrToString(item.FailoverToSecondary)
	respItem["failback_retry_delay"] = item.FailbackRetryDelay
	respItem["specify_server_for_each_ise_node"] = boolPtrToString(item.SpecifyServerForEachIseNode)
	respItem["always_access_primary_first"] = boolPtrToString(item.AlwaysAccessPrimaryFirst)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLdapGetLdapidItemConnectionSettingsPrimaryServer(item *isegosdk.ResponseLdapGetLdapidERSLdapConnectionSettingsPrimaryServer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_name"] = item.HostName
	respItem["port"] = item.Port
	respItem["max_connections"] = item.MaxConnections
	respItem["server_timeout"] = item.ServerTimeout
	respItem["use_admin_access"] = boolPtrToString(item.UseAdminAccess)
	respItem["admin_dn"] = item.AdminDN
	respItem["admin_password"] = item.AdminPassword
	respItem["enable_secure_connection"] = boolPtrToString(item.EnableSecureConnection)
	respItem["enable_server_identity_check"] = boolPtrToString(item.EnableServerIDentityCheck)
	respItem["trust_certificate"] = item.TrustCertificate
	respItem["issuer_cacertificate"] = item.IssuerCaCertificate
	respItem["enable_force_reconnect"] = boolPtrToString(item.EnableForceReconnect)
	respItem["force_reconnect_time"] = item.ForceReconnectTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLdapGetLdapidItemConnectionSettingsSecondaryServer(item *isegosdk.ResponseLdapGetLdapidERSLdapConnectionSettingsSecondaryServer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_name"] = item.HostName
	respItem["port"] = item.Port
	respItem["max_connections"] = item.MaxConnections
	respItem["server_timeout"] = item.ServerTimeout
	respItem["use_admin_access"] = boolPtrToString(item.UseAdminAccess)
	respItem["admin_dn"] = item.AdminDN
	respItem["admin_password"] = item.AdminPassword
	respItem["enable_secure_connection"] = boolPtrToString(item.EnableSecureConnection)
	respItem["enable_server_identity_check"] = boolPtrToString(item.EnableServerIDentityCheck)
	respItem["trust_certificate"] = item.TrustCertificate
	respItem["issuer_cacertificate"] = item.IssuerCaCertificate
	respItem["enable_force_reconnect"] = boolPtrToString(item.EnableForceReconnect)
	respItem["force_reconnect_time"] = item.ForceReconnectTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLdapGetLdapidItemConnectionSettingsLdapNodeData(items *[]isegosdk.ResponseLdapGetLdapidERSLdapConnectionSettingsLdapNodeData) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["primary_hostname"] = item.PrimaryHostname
		respItem["secondary_hostname"] = item.SecondaryHostname
		respItem["primary_port"] = item.PrimaryPort
		respItem["secondary_port"] = item.SecondaryPort
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLdapGetLdapidItemDirectoryOrganization(item *isegosdk.ResponseLdapGetLdapidERSLdapDirectoryOrganization) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["user_directory_subtree"] = item.UserDirectorySubtree
	respItem["group_directory_subtree"] = item.GroupDirectorySubtree
	respItem["mac_format"] = flattenLdapGetLdapidItemDirectoryOrganizationMacFormat(item.MacFormat)
	respItem["strip_prefix"] = boolPtrToString(item.StripPrefix)
	respItem["strip_suffix"] = item.StripSuffix
	respItem["prefix_separator"] = item.PrefixSeparator
	respItem["suffix_separator"] = item.SuffixSeparator

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLdapGetLdapidItemDirectoryOrganizationMacFormat(item *isegosdk.ResponseLdapGetLdapidERSLdapDirectoryOrganizationMacFormat) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenLdapGetLdapidItemGroups(item *isegosdk.ResponseLdapGetLdapidERSLdapGroups) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["groups_names"] = item.GroupsNames

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLdapGetLdapidItemAttributes(item *isegosdk.ResponseLdapGetLdapidERSLdapAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attributes"] = item.Attributes

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLdapGetLdapidItemLink(item *isegosdk.ResponseLdapGetLdapidERSLdapLink) []map[string]interface{} {
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
