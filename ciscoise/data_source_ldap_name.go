package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLdapName() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ldap.

- Get-By-Name
`,

		ReadContext:   dataSourceLdapNameRead,
		Schema: map[string]*schema.Schema{
      "name": &schema.Schema{
        Description: `name path parameter.`,
        Type:        schema.TypeString,
        Required:    true,
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
                    Elem:        &schema.Schema{
                      Type:      schema.TypeString,
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
            "enable_password_change_lda_p": &schema.Schema{
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
                    Elem:        &schema.Schema{
                      Type:      schema.TypeString,
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
		},
	}
}

func dataSourceLdapNameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vName := d.Get("name")


	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetLdapNameName")
		vvName := vName.(string)

		response1, restyResp1, err := client.Ldap.GetLdapNameName(vvName)

	
	
		if err != nil || response1 == nil {
		  if restyResp1 != nil {
		    log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		  }
		  diags = append(diags, diagErrorWithAlt(
		    "Failure when executing 2 GetLdapNameName", err,
		    "Failure at GetLdapNameName, unexpected response", ""))
		  return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLdapGetLdapNameNameItem(response1.ERSLdap)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLdapNameName response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
  return diags
}



func flattenLdapGetLdapNameNameItem(item *isegosdk.ResponseLdapGetLdapNameNameERSLdap) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["general_settings"] = flattenLdapGetLdapNameNameItemGeneralSettings(item.GeneralSettings)
	respItem["connection_settings"] = flattenLdapGetLdapNameNameItemConnectionSettings(item.ConnectionSettings)
	respItem["directory_organization"] = flattenLdapGetLdapNameNameItemDirectoryOrganization(item.DirectoryOrganization)
	respItem["groups"] = flattenLdapGetLdapNameNameItemGroups(item.Groups)
	respItem["attributes"] = flattenLdapGetLdapNameNameItemAttributes(item.Attributes)
	respItem["enable_password_change_lda_p"] = boolPtrToString(item.EnablePasswordChangeLDAP)
	respItem["name"] = item.Name
	respItem["id"] = item.ID
	respItem["description"] = item.Description
	respItem["link"] = flattenLdapGetLdapNameNameItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}


func flattenLdapGetLdapNameNameItemGeneralSettings(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapGeneralSettings) []map[string]interface{} {
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
	respItem["schema"] = flattenLdapGetLdapNameNameItemGeneralSettingsSchema(item.Schema)
	respItem["group_member_reference"] = flattenLdapGetLdapNameNameItemGeneralSettingsGroupMemberReference(item.GroupMemberReference)
	respItem["user_info_attributes"] = flattenLdapGetLdapNameNameItemGeneralSettingsUserInfoAttributes(item.UserInfoAttributes)

	return []map[string]interface{}{
		respItem,
	}

}


func flattenLdapGetLdapNameNameItemGeneralSettingsSchema(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapGeneralSettingsSchema) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}



func flattenLdapGetLdapNameNameItemGeneralSettingsGroupMemberReference(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapGeneralSettingsGroupMemberReference) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}



func flattenLdapGetLdapNameNameItemGeneralSettingsUserInfoAttributes(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapGeneralSettingsUserInfoAttributes) []map[string]interface{} {
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



func flattenLdapGetLdapNameNameItemConnectionSettings(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapConnectionSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["primary_server"] = flattenLdapGetLdapNameNameItemConnectionSettingsPrimaryServer(item.PrimaryServer)
	respItem["secondary_server"] = flattenLdapGetLdapNameNameItemConnectionSettingsSecondaryServer(item.SecondaryServer)
	respItem["ldap_node_data"] = flattenLdapGetLdapNameNameItemConnectionSettingsLdapNodeData(item.LdapNodeData)
	respItem["failover_to_secondary"] = boolPtrToString(item.FailoverToSecondary)
	respItem["failback_retry_delay"] = item.FailbackRetryDelay
	respItem["specify_server_for_each_ise_node"] = boolPtrToString(item.SpecifyServerForEachIseNode)
	respItem["always_access_primary_first"] = boolPtrToString(item.AlwaysAccessPrimaryFirst)

	return []map[string]interface{}{
		respItem,
	}

}


func flattenLdapGetLdapNameNameItemConnectionSettingsPrimaryServer(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapConnectionSettingsPrimaryServer) []map[string]interface{} {
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



func flattenLdapGetLdapNameNameItemConnectionSettingsSecondaryServer(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapConnectionSettingsSecondaryServer) []map[string]interface{} {
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



func flattenLdapGetLdapNameNameItemConnectionSettingsLdapNodeData(items *[]isegosdk.ResponseLdapGetLdapNameNameERSLdapConnectionSettingsLdapNodeData) []map[string]interface{} {
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



func flattenLdapGetLdapNameNameItemDirectoryOrganization(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapDirectoryOrganization) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["user_directory_subtree"] = item.UserDirectorySubtree
	respItem["group_directory_subtree"] = item.GroupDirectorySubtree
	respItem["mac_format"] = flattenLdapGetLdapNameNameItemDirectoryOrganizationMacFormat(item.MacFormat)
	respItem["strip_prefix"] = boolPtrToString(item.StripPrefix)
	respItem["strip_suffix"] = item.StripSuffix
	respItem["prefix_separator"] = item.PrefixSeparator
	respItem["suffix_separator"] = item.SuffixSeparator

	return []map[string]interface{}{
		respItem,
	}

}


func flattenLdapGetLdapNameNameItemDirectoryOrganizationMacFormat(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapDirectoryOrganizationMacFormat) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}



func flattenLdapGetLdapNameNameItemGroups(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapGroups) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["groups_names"] = item.GroupsNames

	return []map[string]interface{}{
		respItem,
	}

}



func flattenLdapGetLdapNameNameItemAttributes(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attributes"] = item.Attributes

	return []map[string]interface{}{
		respItem,
	}

}



func flattenLdapGetLdapNameNameItemLink(item *isegosdk.ResponseLdapGetLdapNameNameERSLdapLink) []map[string]interface{} {
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