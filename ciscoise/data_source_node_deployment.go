package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodeDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNodeDeploymentRead,
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"persona_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"roles": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"services": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"node_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"fqdn": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"node_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"administration": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"is_enabled": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"role": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"general_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"monitoring": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"is_enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"role": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"other_monitoring_node": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"is_mnt_dedicated": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"policyservice": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"enabled": &schema.Schema{
																			Type:     schema.TypeBool,
																			Computed: true,
																		},
																		"session_service": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"is_enabled": &schema.Schema{
																						Type:     schema.TypeBool,
																						Computed: true,
																					},
																					"nodegroup": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"enable_profiling_service": &schema.Schema{
																			Type:     schema.TypeBool,
																			Computed: true,
																		},
																		"enable_nac_service": &schema.Schema{
																			Type:     schema.TypeBool,
																			Computed: true,
																		},
																		"sxpservice": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"is_enabled": &schema.Schema{
																						Type:     schema.TypeBool,
																						Computed: true,
																					},
																					"user_interface": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"enable_device_admin_service": &schema.Schema{
																			Type:     schema.TypeBool,
																			Computed: true,
																		},
																		"enable_passive_identity_service": &schema.Schema{
																			Type:     schema.TypeBool,
																			Computed: true,
																		},
																	},
																},
															},
															"enable_pxgrid": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"profiling_configuration": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"netflow": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"interface": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"dhcp": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"interface": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"dhcp_span": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"interface": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"http": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"interface": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"radius": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"nmap": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"dns": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"snmp_query": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"retries": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"timeout": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"event_timeout": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"snmp_trap": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"link_trap_query": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"mac_trap_query": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"interface": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"active_directory": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"days_before_rescan": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"pxgrid": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enabled": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
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

func dataSourceNodeDeploymentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vHostname, okHostname := d.GetOk("hostname")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okHostname}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNodes")

		response1, _, err := client.NodeDeployment.GetNodes()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodes", err,
				"Failure at GetNodes, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenNodeDeploymentGetNodesItems(&response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodes response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNodeDetails")
		vvHostname := vHostname.(string)

		response2, _, err := client.NodeDeployment.GetNodeDetails(vvHostname)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeDetails", err,
				"Failure at GetNodeDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNodeDeploymentGetNodeDetailsItem(&response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNodeDeploymentGetNodesItems(items *[]isegosdk.ResponseNodeDeploymentGetNodesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["hostname"] = item.Hostname
		respItem["persona_type"] = item.PersonaType
		respItem["roles"] = item.Roles
		respItem["services"] = item.Services
		respItem["node_status"] = item.NodeStatus
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNodeDeploymentGetNodeDetailsItem(item *isegosdk.ResponseNodeDeploymentGetNodeDetailsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["hostname"] = item.Hostname
	respItem["fqdn"] = item.Fqdn
	respItem["ip_address"] = item.IPAddress
	respItem["node_type"] = item.NodeType
	respItem["administration"] = flattenNodeDeploymentGetNodeDetailsItemAdministration(item.Administration)
	respItem["general_settings"] = flattenNodeDeploymentGetNodeDetailsItemGeneralSettings(item.GeneralSettings)
	respItem["profiling_configuration"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfiguration(item.ProfilingConfiguration)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeDeploymentGetNodeDetailsItemAdministration(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseAdministration) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["is_enabled"] = item.IsEnabled
	respItem["role"] = item.Role

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemGeneralSettings(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["monitoring"] = flattenNodeDeploymentGetNodeDetailsItemGeneralSettingsMonitoring(item.Monitoring)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemGeneralSettingsMonitoring(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoring) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["is_enabled"] = item.IsEnabled
	respItem["role"] = item.Role
	respItem["other_monitoring_node"] = item.OtherMonitoringNode
	respItem["is_mnt_dedicated"] = item.IsMntDedicated
	respItem["policyservice"] = flattenNodeDeploymentGetNodeDetailsItemGeneralSettingsMonitoringPolicyservice(item.Policyservice)
	respItem["enable_pxgrid"] = item.EnablePXGrid

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemGeneralSettingsMonitoringPolicyservice(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoringPolicyservice) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["session_service"] = flattenNodeDeploymentGetNodeDetailsItemGeneralSettingsMonitoringPolicyserviceSessionService(item.SessionService)
	respItem["enable_profiling_service"] = item.EnableProfilingService
	respItem["enable_nac_service"] = item.EnableNACService
	respItem["sxpservice"] = flattenNodeDeploymentGetNodeDetailsItemGeneralSettingsMonitoringPolicyserviceSxpservice(item.Sxpservice)
	respItem["enable_device_admin_service"] = item.EnableDeviceAdminService
	respItem["enable_passive_identity_service"] = item.EnablePassiveIDentityService

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemGeneralSettingsMonitoringPolicyserviceSessionService(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoringPolicyserviceSessionService) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["is_enabled"] = item.IsEnabled
	respItem["nodegroup"] = item.Nodegroup

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemGeneralSettingsMonitoringPolicyserviceSxpservice(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseGeneralSettingsMonitoringPolicyserviceSxpservice) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["is_enabled"] = item.IsEnabled
	respItem["user_interface"] = item.UserInterface

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfiguration(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfiguration) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["netflow"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationNetflow(item.Netflow)
	respItem["dhcp"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationDhcp(item.Dhcp)
	respItem["dhcp_span"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationDhcpSpan(item.DhcpSpan)
	respItem["http"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationHTTP(item.HTTP)
	respItem["radius"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationRadius(item.Radius)
	respItem["nmap"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationNmap(item.Nmap)
	respItem["dns"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationDNS(item.DNS)
	respItem["snmp_query"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationSNMPQuery(item.SNMPQuery)
	respItem["snmp_trap"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationSNMPTrap(item.SNMPTrap)
	respItem["active_directory"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationActiveDirectory(item.ActiveDirectory)
	respItem["pxgrid"] = flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationPxgrid(item.Pxgrid)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationNetflow(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationNetflow) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["interface"] = item.Interface
	respItem["port"] = item.Port
	respItem["description"] = item.Description

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationDhcp(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationDhcp) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["interface"] = item.Interface
	respItem["port"] = item.Port
	respItem["description"] = item.Description

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationDhcpSpan(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationDhcpSpan) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["interface"] = item.Interface
	respItem["description"] = item.Description

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationHTTP(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationHTTP) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["interface"] = item.Interface
	respItem["description"] = item.Description

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationRadius(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationRadius) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["description"] = item.Description

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationNmap(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationNmap) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["description"] = item.Description

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationDNS(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationDNS) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["description"] = item.Description

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationSNMPQuery(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationSNMPQuery) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["description"] = item.Description
	respItem["retries"] = item.Retries
	respItem["timeout"] = item.Timeout
	respItem["event_timeout"] = item.EventTimeout

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationSNMPTrap(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationSNMPTrap) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["link_trap_query"] = item.LinkTrapQuery
	respItem["mac_trap_query"] = item.MacTrapQuery
	respItem["interface"] = item.Interface
	respItem["port"] = item.Port
	respItem["description"] = item.Description

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationActiveDirectory(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationActiveDirectory) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["days_before_rescan"] = item.DaysBeforeRescan
	respItem["description"] = item.Description

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeDeploymentGetNodeDetailsItemProfilingConfigurationPxgrid(item isegosdk.ResponseNodeDeploymentGetNodeDetailsResponseProfilingConfigurationPxgrid) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["enabled"] = item.Enabled
	respItem["description"] = item.Description

	return []map[string]interface{}{
		respItem,
	}

}
