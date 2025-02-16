package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLdap() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on ldap.

- Create

- Update

- Delete
`,

		CreateContext: resourceLdapCreate,
		ReadContext:   resourceLdapRead,
		UpdateContext: resourceLdapUpdate,
		DeleteContext: resourceLdapDelete,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attributes": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Description:      `List of Attributes`,
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"connection_settings": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"always_access_primary_first": &schema.Schema{
										Description:      `alwaysAccessPrimaryFirst`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"failback_retry_delay": &schema.Schema{
										Description:      `failbackRetryDelay`,
										Type:             schema.TypeFloat,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"failover_to_secondary": &schema.Schema{
										Description:      `failoverToSecondary`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"ldap_node_data": &schema.Schema{
										Description:      `ldapNodeData`,
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description:      `name`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"primary_hostname": &schema.Schema{
													Description:      `ipaddress`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"primary_port": &schema.Schema{
													Description:      `primaryPort`,
													Type:             schema.TypeFloat,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"secondary_hostname": &schema.Schema{
													Description:      `ipaddress`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"secondary_port": &schema.Schema{
													Description:      `secondaryPort`,
													Type:             schema.TypeFloat,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
											},
										},
									},
									"primary_server": &schema.Schema{
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"admin_dn": &schema.Schema{
													Description:      `adminDN`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"admin_password": &schema.Schema{
													Description:      `adminPassword`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"enable_force_reconnect": &schema.Schema{
													Description:      `enableForceReconnect`,
													Type:             schema.TypeString,
													ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:         true,
													DiffSuppressFunc: diffSupressBool(),
													Computed:         true,
												},
												"enable_secure_connection": &schema.Schema{
													Description:      `enableSecureConnection`,
													Type:             schema.TypeString,
													ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:         true,
													DiffSuppressFunc: diffSupressBool(),
													Computed:         true,
												},
												"enable_server_identity_check": &schema.Schema{
													Description:      `enableServerIdentityCheck`,
													Type:             schema.TypeString,
													ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:         true,
													DiffSuppressFunc: diffSupressBool(),
													Computed:         true,
												},
												"force_reconnect_time": &schema.Schema{
													Description:      `forceReconnectTime`,
													Type:             schema.TypeFloat,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"host_name": &schema.Schema{
													Description:      `hostName`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"issuer_cacertificate": &schema.Schema{
													Description:      `issuerCACertificate`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"max_connections": &schema.Schema{
													Description:      `maxConnections`,
													Type:             schema.TypeFloat,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"port": &schema.Schema{
													Description:      `port`,
													Type:             schema.TypeFloat,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"server_timeout": &schema.Schema{
													Description:      `serverTimeout`,
													Type:             schema.TypeFloat,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"trust_certificate": &schema.Schema{
													Description:      `trustCertificate`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"use_admin_access": &schema.Schema{
													Description:      `useAdminAccess`,
													Type:             schema.TypeString,
													ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:         true,
													DiffSuppressFunc: diffSupressBool(),
													Computed:         true,
												},
											},
										},
									},
									"secondary_server": &schema.Schema{
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"admin_dn": &schema.Schema{
													Description:      `adminDN`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"admin_password": &schema.Schema{
													Description:      `adminPassword`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"enable_force_reconnect": &schema.Schema{
													Description:      `enableForceReconnect`,
													Type:             schema.TypeString,
													ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:         true,
													DiffSuppressFunc: diffSupressBool(),
													Computed:         true,
												},
												"enable_secure_connection": &schema.Schema{
													Description:      `enableSecureConnection`,
													Type:             schema.TypeString,
													ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:         true,
													DiffSuppressFunc: diffSupressBool(),
													Computed:         true,
												},
												"enable_server_identity_check": &schema.Schema{
													Description:      `enableServerIdentityCheck`,
													Type:             schema.TypeString,
													ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:         true,
													DiffSuppressFunc: diffSupressBool(),
													Computed:         true,
												},
												"force_reconnect_time": &schema.Schema{
													Description:      `forceReconnectTime`,
													Type:             schema.TypeFloat,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"host_name": &schema.Schema{
													Description:      `hostName`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"issuer_cacertificate": &schema.Schema{
													Description:      `issuerCACertificate`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"max_connections": &schema.Schema{
													Description:      `maxConnections`,
													Type:             schema.TypeFloat,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"port": &schema.Schema{
													Description:      `port`,
													Type:             schema.TypeFloat,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"server_timeout": &schema.Schema{
													Description:      `serverTimeout`,
													Type:             schema.TypeFloat,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"trust_certificate": &schema.Schema{
													Description:      `trustCertificate`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"use_admin_access": &schema.Schema{
													Description:      `useAdminAccess`,
													Type:             schema.TypeString,
													ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:         true,
													DiffSuppressFunc: diffSupressBool(),
													Computed:         true,
												},
											},
										},
									},
									"specify_server_for_each_ise_node": &schema.Schema{
										Description:      `specifyServerForEachISENode`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Description:      `Description`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"directory_organization": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"group_directory_subtree": &schema.Schema{
										Description:      `groupDirectorySubtree`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"mac_format": &schema.Schema{
										Description:      `macFormat`,
										Type:             schema.TypeString, //TEST,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"prefix_separator": &schema.Schema{
										Description:      `prefixSeparator`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"strip_prefix": &schema.Schema{
										Description:      `stripPrefix`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"strip_suffix": &schema.Schema{
										Description:      `stripSuffix`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"suffix_separator": &schema.Schema{
										Description:      `suffixSeparator`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"user_directory_subtree": &schema.Schema{
										Description:      `userDirectorySubtree`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"enable_password_change_ldap": &schema.Schema{
							Description:      `enablePasswordChangeLDAP`,
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"general_settings": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"certificate": &schema.Schema{
										Description:      `certificate`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"group_map_attribute_name": &schema.Schema{
										Description:      `groupMapAttributeName`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"group_member_reference": &schema.Schema{
										Description:      `groupMemberReference`,
										Type:             schema.TypeString, //TEST,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"group_name_attribute": &schema.Schema{
										Description:      `groupNameAttribute`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"group_object_class": &schema.Schema{
										Description:      `groupObjectClass`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"schema": &schema.Schema{
										Description:      `schema`,
										Type:             schema.TypeString, //TEST,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"user_info_attributes": &schema.Schema{
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"additional_attribute": &schema.Schema{
													Description:      `additionalAttribute`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"country": &schema.Schema{
													Description:      `country`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"department": &schema.Schema{
													Description:      `department`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"email": &schema.Schema{
													Description:      `email`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"first_name": &schema.Schema{
													Description:      `firstName`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"job_title": &schema.Schema{
													Description:      `jobTitle`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"last_name": &schema.Schema{
													Description:      `lastName`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"locality": &schema.Schema{
													Description:      `locality`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"organizational_unit": &schema.Schema{
													Description:      `organizationalUnit`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"state_or_province": &schema.Schema{
													Description:      `stateOrProvince`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"street_address": &schema.Schema{
													Description:      `streetAddress`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"telephone": &schema.Schema{
													Description:      `telephone`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
											},
										},
									},
									"user_name_attribute": &schema.Schema{
										Description:      `userNameAttribute`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"user_object_class": &schema.Schema{
										Description:      `userObjectClass.`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"groups": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"groups_names": &schema.Schema{
										Description:      `List of groups`,
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"id": &schema.Schema{
							Description:      `Id`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
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
							Description:      `name`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func resourceLdapCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning LDAP create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := m.(ClientConfig).EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestLdapPostLdap(ctx, "parameters.0", d)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse2, _, err := client.Ldap.GetLdapid(vvID)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceLdapRead(ctx, d, m)
			}
		} else {
			queryParams2 := isegosdk.GetLdapQueryParams{}

			response2, _, err := client.Ldap.GetLdap(&queryParams2)
			if response2 != nil && err == nil {
				items2 := getAllItemsLdapGetLdap(m, response2, &queryParams2)
				item2, err := searchLdapGetLdap(m, items2, vvName, vvID)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["id"] = item2.ID
					resourceMap["name"] = vvName
					d.SetId(joinResourceID(resourceMap))
					return resourceLdapRead(ctx, d, m)
				}
			}
		}
	}
	resp1, restyResp1, err := client.Ldap.PostLdap(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing PostLdap", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing PostLdap", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceLdapRead(ctx, d, m)
}

func resourceLdapRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	vvID := vID
	vvName := vName

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetLdap")
		queryParams1 := isegosdk.GetLdapQueryParams{}

		response1, restyResp1, err := client.Ldap.GetLdap(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsLdapGetLdap(m, response1, &queryParams1)
		item1, err := searchLdapGetLdap(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenLdapGetLdapidItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLdap search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLdap search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetLdapid")

		response2, restyResp2, err := client.Ldap.GetLdapid(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
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
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLdapid response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceLdapUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	var vvID string

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})

	if selectedMethod == 2 {
		log.Printf("[DEBUG] Update selected method: GetLdap")
		queryParams1 := isegosdk.GetLdapQueryParams{}
		response1, _, err := client.Ldap.GetLdap(&queryParams1)
		if err == nil && response1 != nil {
			items1 := getAllItemsLdapGetLdap(m, response1, &queryParams1)
			item1, err := searchLdapGetLdap(m, items1, vName, vID)
			if err == nil && item1 != nil {
				if vID != item1.ID {
					vvID = item1.ID
				} else {
					vvID = vID
				}
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestLdapPutLdapid(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Ldap.PutLdapid(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing PutLdapid", err, restyResp1.String(),
					"Failure at PutLdapid, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing PutLdapid", err,
				"Failure at PutLdapid, unexpected response", ""))
			return diags
		}
	}
	return resourceLdapRead(ctx, d, m)
}

func resourceLdapDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning LDAP delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	var vvID string

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})

	if selectedMethod == 2 {
		log.Printf("[DEBUG] Update selected method: GetLdap")
		queryParams1 := isegosdk.GetLdapQueryParams{}
		response1, _, err := client.Ldap.GetLdap(&queryParams1)
		if err != nil || response1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsLdapGetLdap(m, response1, &queryParams1)
		item1, err := searchLdapGetLdap(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.Ldap.GetLdapid(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}

	response1, restyResp1, err := client.Ldap.DeleteLdapid(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteLdapid", err, restyResp1.String(),
				"Failure at DeleteLdapid, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteLdapid", err,
			"Failure at DeleteLdapid, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestLdapPostLdap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdap {
	request := isegosdk.RequestLdapPostLdap{}
	request.ERSLdap = expandRequestLdapPostLdapERSLdap(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdap {
	request := isegosdk.RequestLdapPostLdapERSLdap{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".general_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".general_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".general_settings")))) {
		request.GeneralSettings = expandRequestLdapPostLdapERSLdapGeneralSettings(ctx, key+".general_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_settings")))) {
		request.ConnectionSettings = expandRequestLdapPostLdapERSLdapConnectionSettings(ctx, key+".connection_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".directory_organization")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".directory_organization")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".directory_organization")))) {
		request.DirectoryOrganization = expandRequestLdapPostLdapERSLdapDirectoryOrganization(ctx, key+".directory_organization.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".groups")))) {
		request.Groups = expandRequestLdapPostLdapERSLdapGroups(ctx, key+".groups.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = expandRequestLdapPostLdapERSLdapAttributes(ctx, key+".attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password_change_ldap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password_change_ldap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password_change_ldap")))) {
		request.EnablePasswordChangeLDAP = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapGeneralSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapGeneralSettings {
	request := isegosdk.RequestLdapPostLdapERSLdapGeneralSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_object_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_object_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_object_class")))) {
		request.UserObjectClass = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name_attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name_attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name_attribute")))) {
		request.UserNameAttribute = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_name_attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_name_attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_name_attribute")))) {
		request.GroupNameAttribute = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_object_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_object_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_object_class")))) {
		request.GroupObjectClass = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_map_attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_map_attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_map_attribute_name")))) {
		request.GroupMapAttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".schema")))) {
		request.Schema = expandRequestLdapPostLdapERSLdapGeneralSettingsSchema(ctx, key+".schema.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_member_reference")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_member_reference")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_member_reference")))) {
		request.GroupMemberReference = expandRequestLdapPostLdapERSLdapGeneralSettingsGroupMemberReference(ctx, key+".group_member_reference.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_info_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_info_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_info_attributes")))) {
		request.UserInfoAttributes = expandRequestLdapPostLdapERSLdapGeneralSettingsUserInfoAttributes(ctx, key+".user_info_attributes.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapGeneralSettingsSchema(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapGeneralSettingsSchema {
	var request isegosdk.RequestLdapPostLdapERSLdapGeneralSettingsSchema
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapGeneralSettingsGroupMemberReference(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapGeneralSettingsGroupMemberReference {
	var request isegosdk.RequestLdapPostLdapERSLdapGeneralSettingsGroupMemberReference
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapGeneralSettingsUserInfoAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapGeneralSettingsUserInfoAttributes {
	request := isegosdk.RequestLdapPostLdapERSLdapGeneralSettingsUserInfoAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_name")))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".department")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".department")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".department")))) {
		request.Department = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_name")))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".organizational_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".organizational_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".organizational_unit")))) {
		request.OrganizationalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".job_title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".job_title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".job_title")))) {
		request.JobTitle = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".locality")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".locality")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".locality")))) {
		request.Locality = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email")))) {
		request.Email = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state_or_province")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state_or_province")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state_or_province")))) {
		request.StateOrProvince = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".telephone")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".telephone")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".telephone")))) {
		request.Telephone = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".country")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".country")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".country")))) {
		request.Country = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".street_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".street_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".street_address")))) {
		request.StreetAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_attribute")))) {
		request.AdditionalAttribute = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapConnectionSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapConnectionSettings {
	request := isegosdk.RequestLdapPostLdapERSLdapConnectionSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_server")))) {
		request.PrimaryServer = expandRequestLdapPostLdapERSLdapConnectionSettingsPrimaryServer(ctx, key+".primary_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_server")))) {
		request.SecondaryServer = expandRequestLdapPostLdapERSLdapConnectionSettingsSecondaryServer(ctx, key+".secondary_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ldap_node_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ldap_node_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ldap_node_data")))) {
		request.LdapNodeData = expandRequestLdapPostLdapERSLdapConnectionSettingsLdapNodeDataArray(ctx, key+".ldap_node_data", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failover_to_secondary")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failover_to_secondary")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failover_to_secondary")))) {
		request.FailoverToSecondary = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failback_retry_delay")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failback_retry_delay")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failback_retry_delay")))) {
		request.FailbackRetryDelay = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".specify_server_for_each_ise_node")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".specify_server_for_each_ise_node")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".specify_server_for_each_ise_node")))) {
		request.SpecifyServerForEachIseNode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".always_access_primary_first")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".always_access_primary_first")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".always_access_primary_first")))) {
		request.AlwaysAccessPrimaryFirst = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapConnectionSettingsPrimaryServer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapConnectionSettingsPrimaryServer {
	request := isegosdk.RequestLdapPostLdapERSLdapConnectionSettingsPrimaryServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_connections")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_connections")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_connections")))) {
		request.MaxConnections = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_timeout")))) {
		request.ServerTimeout = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_admin_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_admin_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_admin_access")))) {
		request.UseAdminAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_dn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_dn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_dn")))) {
		request.AdminDN = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_password")))) {
		request.AdminPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_secure_connection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_secure_connection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_secure_connection")))) {
		request.EnableSecureConnection = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_server_identity_check")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_server_identity_check")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_server_identity_check")))) {
		request.EnableServerIDentityCheck = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trust_certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trust_certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trust_certificate")))) {
		request.TrustCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".issuer_cacertificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".issuer_cacertificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".issuer_cacertificate")))) {
		request.IssuerCaCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_force_reconnect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_force_reconnect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_force_reconnect")))) {
		request.EnableForceReconnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".force_reconnect_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".force_reconnect_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".force_reconnect_time")))) {
		request.ForceReconnectTime = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapConnectionSettingsSecondaryServer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapConnectionSettingsSecondaryServer {
	request := isegosdk.RequestLdapPostLdapERSLdapConnectionSettingsSecondaryServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_connections")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_connections")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_connections")))) {
		request.MaxConnections = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_timeout")))) {
		request.ServerTimeout = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_admin_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_admin_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_admin_access")))) {
		request.UseAdminAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_dn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_dn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_dn")))) {
		request.AdminDN = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_password")))) {
		request.AdminPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_secure_connection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_secure_connection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_secure_connection")))) {
		request.EnableSecureConnection = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_server_identity_check")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_server_identity_check")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_server_identity_check")))) {
		request.EnableServerIDentityCheck = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trust_certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trust_certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trust_certificate")))) {
		request.TrustCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".issuer_cacertificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".issuer_cacertificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".issuer_cacertificate")))) {
		request.IssuerCaCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_force_reconnect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_force_reconnect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_force_reconnect")))) {
		request.EnableForceReconnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".force_reconnect_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".force_reconnect_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".force_reconnect_time")))) {
		request.ForceReconnectTime = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapConnectionSettingsLdapNodeDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestLdapPostLdapERSLdapConnectionSettingsLdapNodeData {
	request := []isegosdk.RequestLdapPostLdapERSLdapConnectionSettingsLdapNodeData{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestLdapPostLdapERSLdapConnectionSettingsLdapNodeData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapConnectionSettingsLdapNodeData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapConnectionSettingsLdapNodeData {
	request := isegosdk.RequestLdapPostLdapERSLdapConnectionSettingsLdapNodeData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_hostname")))) {
		request.PrimaryHostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_hostname")))) {
		request.SecondaryHostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_port")))) {
		request.PrimaryPort = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_port")))) {
		request.SecondaryPort = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapDirectoryOrganization(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapDirectoryOrganization {
	request := isegosdk.RequestLdapPostLdapERSLdapDirectoryOrganization{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_directory_subtree")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_directory_subtree")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_directory_subtree")))) {
		request.UserDirectorySubtree = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_directory_subtree")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_directory_subtree")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_directory_subtree")))) {
		request.GroupDirectorySubtree = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_format")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_format")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_format")))) {
		request.MacFormat = expandRequestLdapPostLdapERSLdapDirectoryOrganizationMacFormat(ctx, key+".mac_format.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".strip_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".strip_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".strip_prefix")))) {
		request.StripPrefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".strip_suffix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".strip_suffix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".strip_suffix")))) {
		request.StripSuffix = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prefix_separator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prefix_separator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prefix_separator")))) {
		request.PrefixSeparator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".suffix_separator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".suffix_separator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".suffix_separator")))) {
		request.SuffixSeparator = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapDirectoryOrganizationMacFormat(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapDirectoryOrganizationMacFormat {
	var request isegosdk.RequestLdapPostLdapERSLdapDirectoryOrganizationMacFormat
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapGroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapGroups {
	request := isegosdk.RequestLdapPostLdapERSLdapGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".groups_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".groups_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".groups_names")))) {
		request.GroupsNames = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPostLdapERSLdapAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPostLdapERSLdapAttributes {
	request := isegosdk.RequestLdapPostLdapERSLdapAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapid(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapid {
	request := isegosdk.RequestLdapPutLdapid{}
	request.ERSLdap = expandRequestLdapPutLdapidERSLdap(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdap {
	request := isegosdk.RequestLdapPutLdapidERSLdap{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".general_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".general_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".general_settings")))) {
		request.GeneralSettings = expandRequestLdapPutLdapidERSLdapGeneralSettings(ctx, key+".general_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_settings")))) {
		request.ConnectionSettings = expandRequestLdapPutLdapidERSLdapConnectionSettings(ctx, key+".connection_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".directory_organization")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".directory_organization")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".directory_organization")))) {
		request.DirectoryOrganization = expandRequestLdapPutLdapidERSLdapDirectoryOrganization(ctx, key+".directory_organization.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".groups")))) {
		request.Groups = expandRequestLdapPutLdapidERSLdapGroups(ctx, key+".groups.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = expandRequestLdapPutLdapidERSLdapAttributes(ctx, key+".attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password_change_ldap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password_change_ldap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password_change_ldap")))) {
		request.EnablePasswordChangeLDAP = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapGeneralSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapGeneralSettings {
	request := isegosdk.RequestLdapPutLdapidERSLdapGeneralSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_object_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_object_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_object_class")))) {
		request.UserObjectClass = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name_attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name_attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name_attribute")))) {
		request.UserNameAttribute = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_name_attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_name_attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_name_attribute")))) {
		request.GroupNameAttribute = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_object_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_object_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_object_class")))) {
		request.GroupObjectClass = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_map_attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_map_attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_map_attribute_name")))) {
		request.GroupMapAttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".schema")))) {
		request.Schema = expandRequestLdapPutLdapidERSLdapGeneralSettingsSchema(ctx, key+".schema.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_member_reference")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_member_reference")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_member_reference")))) {
		request.GroupMemberReference = expandRequestLdapPutLdapidERSLdapGeneralSettingsGroupMemberReference(ctx, key+".group_member_reference.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_info_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_info_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_info_attributes")))) {
		request.UserInfoAttributes = expandRequestLdapPutLdapidERSLdapGeneralSettingsUserInfoAttributes(ctx, key+".user_info_attributes.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapGeneralSettingsSchema(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapGeneralSettingsSchema {
	var request isegosdk.RequestLdapPutLdapidERSLdapGeneralSettingsSchema
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapGeneralSettingsGroupMemberReference(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapGeneralSettingsGroupMemberReference {
	var request isegosdk.RequestLdapPutLdapidERSLdapGeneralSettingsGroupMemberReference
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapGeneralSettingsUserInfoAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapGeneralSettingsUserInfoAttributes {
	request := isegosdk.RequestLdapPutLdapidERSLdapGeneralSettingsUserInfoAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_name")))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".department")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".department")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".department")))) {
		request.Department = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_name")))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".organizational_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".organizational_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".organizational_unit")))) {
		request.OrganizationalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".job_title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".job_title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".job_title")))) {
		request.JobTitle = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".locality")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".locality")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".locality")))) {
		request.Locality = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email")))) {
		request.Email = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state_or_province")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state_or_province")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state_or_province")))) {
		request.StateOrProvince = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".telephone")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".telephone")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".telephone")))) {
		request.Telephone = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".country")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".country")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".country")))) {
		request.Country = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".street_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".street_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".street_address")))) {
		request.StreetAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_attribute")))) {
		request.AdditionalAttribute = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapConnectionSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapConnectionSettings {
	request := isegosdk.RequestLdapPutLdapidERSLdapConnectionSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_server")))) {
		request.PrimaryServer = expandRequestLdapPutLdapidERSLdapConnectionSettingsPrimaryServer(ctx, key+".primary_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_server")))) {
		request.SecondaryServer = expandRequestLdapPutLdapidERSLdapConnectionSettingsSecondaryServer(ctx, key+".secondary_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ldap_node_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ldap_node_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ldap_node_data")))) {
		request.LdapNodeData = expandRequestLdapPutLdapidERSLdapConnectionSettingsLdapNodeDataArray(ctx, key+".ldap_node_data", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failover_to_secondary")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failover_to_secondary")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failover_to_secondary")))) {
		request.FailoverToSecondary = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failback_retry_delay")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failback_retry_delay")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failback_retry_delay")))) {
		request.FailbackRetryDelay = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".specify_server_for_each_ise_node")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".specify_server_for_each_ise_node")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".specify_server_for_each_ise_node")))) {
		request.SpecifyServerForEachIseNode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".always_access_primary_first")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".always_access_primary_first")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".always_access_primary_first")))) {
		request.AlwaysAccessPrimaryFirst = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapConnectionSettingsPrimaryServer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapConnectionSettingsPrimaryServer {
	request := isegosdk.RequestLdapPutLdapidERSLdapConnectionSettingsPrimaryServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_connections")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_connections")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_connections")))) {
		request.MaxConnections = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_timeout")))) {
		request.ServerTimeout = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_admin_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_admin_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_admin_access")))) {
		request.UseAdminAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_dn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_dn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_dn")))) {
		request.AdminDN = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_password")))) {
		request.AdminPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_secure_connection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_secure_connection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_secure_connection")))) {
		request.EnableSecureConnection = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_server_identity_check")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_server_identity_check")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_server_identity_check")))) {
		request.EnableServerIDentityCheck = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trust_certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trust_certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trust_certificate")))) {
		request.TrustCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".issuer_cacertificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".issuer_cacertificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".issuer_cacertificate")))) {
		request.IssuerCaCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_force_reconnect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_force_reconnect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_force_reconnect")))) {
		request.EnableForceReconnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".force_reconnect_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".force_reconnect_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".force_reconnect_time")))) {
		request.ForceReconnectTime = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapConnectionSettingsSecondaryServer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapConnectionSettingsSecondaryServer {
	request := isegosdk.RequestLdapPutLdapidERSLdapConnectionSettingsSecondaryServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_connections")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_connections")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_connections")))) {
		request.MaxConnections = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_timeout")))) {
		request.ServerTimeout = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_admin_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_admin_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_admin_access")))) {
		request.UseAdminAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_dn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_dn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_dn")))) {
		request.AdminDN = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_password")))) {
		request.AdminPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_secure_connection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_secure_connection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_secure_connection")))) {
		request.EnableSecureConnection = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_server_identity_check")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_server_identity_check")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_server_identity_check")))) {
		request.EnableServerIDentityCheck = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trust_certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trust_certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trust_certificate")))) {
		request.TrustCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".issuer_cacertificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".issuer_cacertificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".issuer_cacertificate")))) {
		request.IssuerCaCertificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_force_reconnect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_force_reconnect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_force_reconnect")))) {
		request.EnableForceReconnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".force_reconnect_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".force_reconnect_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".force_reconnect_time")))) {
		request.ForceReconnectTime = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapConnectionSettingsLdapNodeDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestLdapPutLdapidERSLdapConnectionSettingsLdapNodeData {
	request := []isegosdk.RequestLdapPutLdapidERSLdapConnectionSettingsLdapNodeData{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestLdapPutLdapidERSLdapConnectionSettingsLdapNodeData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapConnectionSettingsLdapNodeData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapConnectionSettingsLdapNodeData {
	request := isegosdk.RequestLdapPutLdapidERSLdapConnectionSettingsLdapNodeData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_hostname")))) {
		request.PrimaryHostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_hostname")))) {
		request.SecondaryHostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_port")))) {
		request.PrimaryPort = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_port")))) {
		request.SecondaryPort = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapDirectoryOrganization(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapDirectoryOrganization {
	request := isegosdk.RequestLdapPutLdapidERSLdapDirectoryOrganization{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_directory_subtree")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_directory_subtree")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_directory_subtree")))) {
		request.UserDirectorySubtree = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_directory_subtree")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_directory_subtree")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_directory_subtree")))) {
		request.GroupDirectorySubtree = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_format")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_format")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_format")))) {
		request.MacFormat = expandRequestLdapPutLdapidERSLdapDirectoryOrganizationMacFormat(ctx, key+".mac_format.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".strip_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".strip_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".strip_prefix")))) {
		request.StripPrefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".strip_suffix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".strip_suffix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".strip_suffix")))) {
		request.StripSuffix = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prefix_separator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prefix_separator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prefix_separator")))) {
		request.PrefixSeparator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".suffix_separator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".suffix_separator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".suffix_separator")))) {
		request.SuffixSeparator = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapDirectoryOrganizationMacFormat(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapDirectoryOrganizationMacFormat {
	var request isegosdk.RequestLdapPutLdapidERSLdapDirectoryOrganizationMacFormat
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapGroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapGroups {
	request := isegosdk.RequestLdapPutLdapidERSLdapGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".groups_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".groups_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".groups_names")))) {
		request.GroupsNames = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLdapPutLdapidERSLdapAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLdapPutLdapidERSLdapAttributes {
	request := isegosdk.RequestLdapPutLdapidERSLdapAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsLdapGetLdap(m interface{}, response *isegosdk.ResponseLdapGetLdap, queryParams *isegosdk.GetLdapQueryParams) []isegosdk.ResponseLdapGetLdapSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseLdapGetLdapSearchResultResources
	for response.SearchResult != nil && response.SearchResult.Resources != nil && len(*response.SearchResult.Resources) > 0 {
		respItems = append(respItems, *response.SearchResult.Resources...)
		if response.SearchResult.NextPage != nil && response.SearchResult.NextPage.Rel == "next" {
			href := response.SearchResult.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			if queryParams != nil {
				queryParams.Page = page
				queryParams.Size = size
			}
			response, _, err = client.Ldap.GetLdap(queryParams)
			if err != nil {
				break
			}
			// All is good, continue to the next page
			continue
		}
		// Does not have next page finish iteration
		break
	}
	return respItems
}

func searchLdapGetLdap(m interface{}, items []isegosdk.ResponseLdapGetLdapSearchResultResources, name string, id string) (*isegosdk.ResponseLdapGetLdapidERSLdap, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseLdapGetLdapidERSLdap
	for _, item := range items {
		if id != "" && item.ID == id {
			var getItem *isegosdk.ResponseLdapGetLdapid
			getItem, _, err = client.Ldap.GetLdapid(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetLdapid")
			}
			foundItem = getItem.ERSLdap
			return foundItem, err
		}
		if name != "" && item.Name == name {
			var getItem *isegosdk.ResponseLdapGetLdapid
			getItem, _, err = client.Ldap.GetLdapid(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetLdapid")
			}
			foundItem = getItem.ERSLdap
			return foundItem, err
		}
	}
	return foundItem, err
}
