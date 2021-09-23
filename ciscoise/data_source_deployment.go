package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeployment() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on PullDeploymentInfo.

- This data source allows the client to pull the deployment information.
`,

		ReadContext: dataSourceDeploymentRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"deployment_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"deployment_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"fipsstatus": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"node_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node_and_node_count_and_count_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"declared_type": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"global_scope": &schema.Schema{
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"nil": &schema.Schema{
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"scope": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"type_substituted": &schema.Schema{
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"value": &schema.Schema{
																Description: `It varies type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"version_history_info": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"epoch_time": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"main_version": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"op_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"kong_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"deployment_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"node_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"service": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"route": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"http_count": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"latency_count": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"latency_sum": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"route_name": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"service_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"sn": &schema.Schema{
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
						"licenses_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"deployment_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"node_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"mdm_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"active_desktop_mdm_servers_count": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"active_mdm_servers_count": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"active_mobile_mdm_servers_count": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"deployment_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"node_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node_and_scope": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"nad_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"nadcount_info": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"total_active_nad_count": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"node_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node_and_scope": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"network_access_info": &schema.Schema{
							Description: `networkAccessInfo Details`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"deployment_id": &schema.Schema{
										Description: `Deployment ID`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"is_csn_enabled": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"node_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node_and_scope": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"radius3_rd_party": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"sda_vns": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"trust_sec_control": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"posture_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"content": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"declared_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"global_scope": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"nil": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"scope": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"type_substituted": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Description: `It varies type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"profiler_info": &schema.Schema{
							Description: `profilerInfo Details`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"deployment_id": &schema.Schema{
										Description: `Deployment ID`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"node_list": &schema.Schema{
										Description: `Deployment ID`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"last_applied_feed_date_time": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"online_subscription_enabled": &schema.Schema{
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"profiles": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"scope": &schema.Schema{
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

func dataSourceDeploymentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeploymentInfo")

		response1, _, err := client.PullDeploymentInfo.GetDeploymentInfo()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeploymentInfo", err,
				"Failure at GetDeploymentInfo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPullDeploymentInfoGetDeploymentInfoItem(response1.ERSDeploymentInfo)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeploymentInfo response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPullDeploymentInfoGetDeploymentInfoItem(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["network_access_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemNetworkAccessInfo(item.NetworkAccessInfo)
	respItem["profiler_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemProfilerInfo(item.ProfilerInfo)
	respItem["deployment_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemDeploymentInfo(item.DeploymentInfo)
	respItem["nad_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemNadInfo(item.NadInfo)
	respItem["mdm_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemMdmInfo(item.MdmInfo)
	respItem["licenses_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemLicensesInfo(item.LicensesInfo)
	respItem["posture_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemPostureInfo(item.PostureInfo)
	respItem["kong_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemKongInfo(item.KongInfo)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenPullDeploymentInfoGetDeploymentInfoItemNetworkAccessInfo(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNetworkAccessInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["deployment_id"] = item.DeploymentID
	respItem["is_csn_enabled"] = boolPtrToString(item.IsCsnEnabled)
	respItem["node_list"] = flattenPullDeploymentInfoGetDeploymentInfoItemNetworkAccessInfoNodeList(item.NodeList)
	respItem["sda_vns"] = responseInterfaceToSliceString(item.SdaVns)
	respItem["trust_sec_control"] = item.TrustSecControl
	respItem["radius3_rd_party"] = responseInterfaceToSliceString(item.Radius3RdParty)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemNetworkAccessInfoNodeList(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNetworkAccessInfoNodeList) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["node_and_scope"] = responseInterfaceToSliceString(item.NodeAndScope)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemProfilerInfo(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoProfilerInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["deployment_id"] = item.DeploymentID
	respItem["node_list"] = flattenPullDeploymentInfoGetDeploymentInfoItemProfilerInfoNodeList(item.NodeList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemProfilerInfoNodeList(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoProfilerInfoNodeList) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["node"] = flattenPullDeploymentInfoGetDeploymentInfoItemProfilerInfoNodeListNode(item.Node)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemProfilerInfoNodeListNode(items *[]isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoProfilerInfoNodeListNode) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["online_subscription_enabled"] = boolPtrToString(item.OnlineSubscriptionEnabled)
		respItem["last_applied_feed_date_time"] = item.LastAppliedFeedDateTime
		respItem["scope"] = item.Scope
		respItem["profiles"] = responseInterfaceToString(item.Profiles)
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenPullDeploymentInfoGetDeploymentInfoItemDeploymentInfo(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["deployment_id"] = item.DeploymentID
	respItem["version_history_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemDeploymentInfoVersionHistoryInfo(item.VersionHistoryInfo)
	respItem["node_list"] = flattenPullDeploymentInfoGetDeploymentInfoItemDeploymentInfoNodeList(item.NodeList)
	respItem["fipsstatus"] = item.Fipsstatus

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemDeploymentInfoVersionHistoryInfo(items *[]isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfoVersionHistoryInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["op_type"] = item.OpType
		respItem["main_version"] = item.MainVersion
		respItem["epoch_time"] = item.EpochTime
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenPullDeploymentInfoGetDeploymentInfoItemDeploymentInfoNodeList(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfoNodeList) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["node_and_node_count_and_count_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemDeploymentInfoNodeListNodeAndNodeCountAndCountInfo(item.NodeAndNodeCountAndCountInfo)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemDeploymentInfoNodeListNodeAndNodeCountAndCountInfo(items *[]isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoDeploymentInfoNodeListNodeAndNodeCountAndCountInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["value"] = responseInterfaceToString(item.Value)
		respItem["declared_type"] = item.DeclaredType
		respItem["scope"] = item.Scope
		respItem["nil"] = boolPtrToString(item.Nil)
		respItem["global_scope"] = boolPtrToString(item.GlobalScope)
		respItem["type_substituted"] = boolPtrToString(item.TypeSubstituted)
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenPullDeploymentInfoGetDeploymentInfoItemNadInfo(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNadInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["node_list"] = flattenPullDeploymentInfoGetDeploymentInfoItemNadInfoNodeList(item.NodeList)
	respItem["nadcount_info"] = flattenPullDeploymentInfoGetDeploymentInfoItemNadInfoNadcountInfo(item.NadcountInfo)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemNadInfoNodeList(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNadInfoNodeList) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["node_and_scope"] = responseInterfaceToSliceString(item.NodeAndScope)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemNadInfoNadcountInfo(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoNadInfoNadcountInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_active_nad_count"] = item.TotalActiveNADCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemMdmInfo(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoMdmInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["active_mdm_servers_count"] = item.ActiveMdmServersCount
	respItem["active_desktop_mdm_servers_count"] = item.ActiveDesktopMdmServersCount
	respItem["active_mobile_mdm_servers_count"] = item.ActiveMobileMdmServersCount
	respItem["deployment_id"] = item.DeploymentID
	respItem["node_list"] = flattenPullDeploymentInfoGetDeploymentInfoItemMdmInfoNodeList(item.NodeList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemMdmInfoNodeList(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoMdmInfoNodeList) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["node_and_scope"] = responseInterfaceToSliceString(item.NodeAndScope)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemLicensesInfo(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoLicensesInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["deployment_id"] = item.DeploymentID
	respItem["node_list"] = flattenPullDeploymentInfoGetDeploymentInfoItemLicensesInfoNodeList(item.NodeList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemLicensesInfoNodeList(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoLicensesInfoNodeList) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["node"] = responseInterfaceToSliceString(item.Node)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemPostureInfo(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoPostureInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["content"] = flattenPullDeploymentInfoGetDeploymentInfoItemPostureInfoContent(item.Content)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemPostureInfoContent(items *[]isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoPostureInfoContent) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["value"] = responseInterfaceToString(item.Value)
		respItem["declared_type"] = item.DeclaredType
		respItem["scope"] = item.Scope
		respItem["nil"] = boolPtrToString(item.Nil)
		respItem["global_scope"] = boolPtrToString(item.GlobalScope)
		respItem["type_substituted"] = boolPtrToString(item.TypeSubstituted)
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenPullDeploymentInfoGetDeploymentInfoItemKongInfo(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["deployment_id"] = item.DeploymentID
	respItem["node_list"] = flattenPullDeploymentInfoGetDeploymentInfoItemKongInfoNodeList(item.NodeList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemKongInfoNodeList(item *isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeList) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["node"] = flattenPullDeploymentInfoGetDeploymentInfoItemKongInfoNodeListNode(item.Node)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPullDeploymentInfoGetDeploymentInfoItemKongInfoNodeListNode(items *[]isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeListNode) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["sn"] = item.Sn
		respItem["service"] = flattenPullDeploymentInfoGetDeploymentInfoItemKongInfoNodeListNodeService(item.Service)
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenPullDeploymentInfoGetDeploymentInfoItemKongInfoNodeListNodeService(items *[]isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeListNodeService) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["service_name"] = item.ServiceName
		respItem["route"] = flattenPullDeploymentInfoGetDeploymentInfoItemKongInfoNodeListNodeServiceRoute(item.Route)
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenPullDeploymentInfoGetDeploymentInfoItemKongInfoNodeListNodeServiceRoute(items *[]isegosdk.ResponsePullDeploymentInfoGetDeploymentInfoERSDeploymentInfoKongInfoNodeListNodeServiceRoute) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["route_name"] = item.RouteName
		respItem["http_count"] = responseInterfaceToString(item.HTTPCount)
		respItem["latency_count"] = responseInterfaceToString(item.LatencyCount)
		respItem["latency_sum"] = responseInterfaceToString(item.LatencySum)
		respItems = append(respItems, respItem)
	}
	return respItems

}
