package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNodeDeployment() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Node Deployment.

- Register ISE node to form a multi-node deployment

- Updates the deployed ISE node with the information provided

- The de-register ednode becomes a standalone Cisco ISE node.
It retains the last configuration that it received rom the PrimaryPAN and assumes the default personas of a standalone
node
that are Administration, PolicyService, and Monitoring.
`,

		CreateContext: resourceNodeDeploymentCreate,
		ReadContext:   resourceNodeDeploymentRead,
		UpdateContext: resourceNodeDeploymentUpdate,
		DeleteContext: resourceNodeDeploymentDelete,
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
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"administration": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"is_enabled": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"role": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"fdqn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"general_settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"monitoring": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_pxgrid": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"is_enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"is_mnt_dedicated": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"other_monitoring_node": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"policyservice": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enable_device_admin_service": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"enable_nac_service": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"enable_passive_identity_service": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"enable_profiling_service": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"session_service": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"is_enabled": &schema.Schema{
																			Type:         schema.TypeString,
																			ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																			Optional:     true,
																			Computed:     true,
																		},
																		"nodegroup": &schema.Schema{
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																	},
																},
															},
															"sxpservice": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"is_enabled": &schema.Schema{
																			Type:         schema.TypeString,
																			ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																			Optional:     true,
																			Computed:     true,
																		},
																		"user_interface": &schema.Schema{
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"role": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"hostname": &schema.Schema{
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
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"profile_configuration": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"active_directory": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"days_before_rescan": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"dhcp": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"port": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"dhcp_span": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"dns": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"http": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"netflow": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"port": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"nmap": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"pxgrid": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"radius": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"snmp_query": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"enabled": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"event_timeout": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"retries": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"timeout": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"snmp_trap": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"link_trap_query": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"mac_trap_query": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"port": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
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

									"active_directory": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"days_before_rescan": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"enabled": &schema.Schema{
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

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"enabled": &schema.Schema{
													Type:     schema.TypeString,
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
											},
										},
									},
									"dhcp_span": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"enabled": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface": &schema.Schema{
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

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"enabled": &schema.Schema{
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

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"enabled": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"netflow": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"enabled": &schema.Schema{
													Type:     schema.TypeString,
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
											},
										},
									},
									"nmap": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"enabled": &schema.Schema{
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

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"enabled": &schema.Schema{
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

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"enabled": &schema.Schema{
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

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"enabled": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"event_timeout": &schema.Schema{
													Type:     schema.TypeInt,
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
											},
										},
									},
									"snmp_trap": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"link_trap_query": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"mac_trap_query": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"port": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"response": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"general_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"monitoring": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enable_pxgrid": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"is_enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"is_mnt_dedicated": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"other_monitoring_node": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"policyservice": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"enable_device_admin_service": &schema.Schema{
																			Type:         schema.TypeString,
																			ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																			Optional:     true,
																		},
																		"enable_nac_service": &schema.Schema{
																			Type:         schema.TypeString,
																			ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																			Optional:     true,
																		},
																		"enable_passive_identity_service": &schema.Schema{
																			Type:         schema.TypeString,
																			ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																			Optional:     true,
																		},
																		"enable_profiling_service": &schema.Schema{
																			Type:         schema.TypeString,
																			ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																			Optional:     true,
																		},
																		"enabled": &schema.Schema{
																			Type:         schema.TypeString,
																			ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																			Optional:     true,
																		},
																		"session_service": &schema.Schema{
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"is_enabled": &schema.Schema{
																						Type:         schema.TypeString,
																						ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																						Optional:     true,
																					},
																					"nodegroup": &schema.Schema{
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},
																		"sxpservice": &schema.Schema{
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"is_enabled": &schema.Schema{
																						Type:         schema.TypeString,
																						ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																						Optional:     true,
																					},
																					"user_interface": &schema.Schema{
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"role": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"profile_configuration": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"active_directory": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"days_before_rescan": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"dhcp": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"interface": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"dhcp_span": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"interface": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"dns": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"http": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"interface": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"netflow": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"interface": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"nmap": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"pxgrid": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"radius": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"snmp_query": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"enabled": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"event_timeout": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"retries": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"timeout": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"snmp_trap": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"description": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"interface": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"link_trap_query": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"mac_trap_query": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
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
						"user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceNodeDeploymentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestNodeDeploymentRegisterNode(ctx, "item.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vHostname, okHostname := resourceItem["hostname"]
	vvHostname := interfaceToString(vHostname)
	if okHostname && vvHostname != "" {
		getResponse2, _, err := client.NodeDeployment.GetNodeDetails(vvHostname)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["hostname"] = vvHostname
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.NodeDeployment.GetNodes()
		if response2 != nil && err == nil {
			items2 := getAllItemsNodeDeploymentGetNodes(m, response2)
			item2, err := searchNodeDeploymentGetNodes(m, items2, vvHostname, "")
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["hostname"] = vvHostname
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	resp1, restyResp1, err := client.NodeDeployment.RegisterNode(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing RegisterNode", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing RegisterNode", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["hostname"] = vvHostname
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNodeDeploymentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostname, okHostname := resourceMap["hostname"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostname}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNodes")

		response1, _, err := client.NodeDeployment.GetNodes()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodes", err,
				"Failure at GetNodes, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsNodeDeploymentGetNodes(m, response1)
		item1, err := searchNodeDeploymentGetNodes(m, items1, vHostname, "")
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetNodes response", err,
				"Failure when searching item from GetNodes, unexpected response", ""))
			return diags
		}
		vItem1 := flattenNodeDeploymentGetNodeDetailsItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodes search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNodeDetails")
		vvHostname := vHostname

		response2, _, err := client.NodeDeployment.GetNodeDetails(vvHostname)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeDetails", err,
				"Failure at GetNodeDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNodeDeploymentGetNodeDetailsItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeDetails response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNodeDeploymentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostname, okHostname := resourceMap["hostname"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostname}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvHostname string
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		vvHostname = vHostname
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] Hostname used for update operation %s", vvHostname)
		request1 := expandRequestNodeDeploymentUpdateNode(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NodeDeployment.UpdateNode(vvHostname, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNode", err, restyResp1.String(),
					"Failure at UpdateNode, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNode", err,
				"Failure at UpdateNode, unexpected response", ""))
			return diags
		}
	}

	return resourceNodeDeploymentRead(ctx, d, m)
}

func resourceNodeDeploymentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostname, okHostname := resourceMap["hostname"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostname}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvHostname string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.NodeDeployment.GetNodes()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNodeDeploymentGetNodes(m, getResp1)
		item1, err := searchNodeDeploymentGetNodes(m, items1, vHostname, "")
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vHostname != item1.Hostname {
			vvHostname = item1.Hostname
		} else {
			vvHostname = vHostname
		}
	}
	if selectedMethod == 2 {
		vvHostname = vHostname
		getResp, _, err := client.NodeDeployment.GetNodeDetails(vvHostname)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NodeDeployment.DeleteNode(vvHostname)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNode", err, restyResp1.String(),
				"Failure at DeleteNode, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNode", err,
			"Failure at DeleteNode, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNodeDeploymentRegisterNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNode {
	request := isegosdk.RequestNodeDeploymentRegisterNode{}
	if v, ok := d.GetOkExists(key + ".fdqn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".fdqn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".fdqn"))) {
		request.Fdqn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".user_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".user_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".user_name"))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".administration"); !isEmptyValue(reflect.ValueOf(d.Get(key+".administration"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".administration"))) {
		request.Administration = expandRequestNodeDeploymentRegisterNodeAdministration(ctx, key+".administration.0", d)
	}
	if v, ok := d.GetOkExists(key + ".general_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".general_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".general_settings"))) {
		request.GeneralSettings = expandRequestNodeDeploymentRegisterNodeGeneralSettings(ctx, key+".general_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".profile_configuration"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile_configuration"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile_configuration"))) {
		request.ProfileConfiguration = expandRequestNodeDeploymentRegisterNodeProfileConfiguration(ctx, key+".profile_configuration.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeAdministration(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeAdministration {
	request := isegosdk.RequestNodeDeploymentRegisterNodeAdministration{}
	if v, ok := d.GetOkExists(key + ".is_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_enabled"))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".role"); !isEmptyValue(reflect.ValueOf(d.Get(key+".role"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".role"))) {
		request.Role = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeGeneralSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeGeneralSettings {
	request := isegosdk.RequestNodeDeploymentRegisterNodeGeneralSettings{}
	if v, ok := d.GetOkExists(key + ".monitoring"); !isEmptyValue(reflect.ValueOf(d.Get(key+".monitoring"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".monitoring"))) {
		request.Monitoring = expandRequestNodeDeploymentRegisterNodeGeneralSettingsMonitoring(ctx, key+".monitoring.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeGeneralSettingsMonitoring(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoring {
	request := isegosdk.RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoring{}
	if v, ok := d.GetOkExists(key + ".is_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_enabled"))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".role"); !isEmptyValue(reflect.ValueOf(d.Get(key+".role"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".role"))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".other_monitoring_node"); !isEmptyValue(reflect.ValueOf(d.Get(key+".other_monitoring_node"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".other_monitoring_node"))) {
		request.OtherMonitoringNode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_mnt_dedicated"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_mnt_dedicated"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_mnt_dedicated"))) {
		request.IsMntDedicated = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".policyservice"); !isEmptyValue(reflect.ValueOf(d.Get(key+".policyservice"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".policyservice"))) {
		request.Policyservice = expandRequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyservice(ctx, key+".policyservice.0", d)
	}
	if v, ok := d.GetOkExists(key + ".enable_pxgrid"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_pxgrid"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_pxgrid"))) {
		request.EnablePXGrid = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyservice(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyservice {
	request := isegosdk.RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyservice{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".session_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".session_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".session_service"))) {
		request.SessionService = expandRequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSessionService(ctx, key+".session_service.0", d)
	}
	if v, ok := d.GetOkExists(key + ".enable_profiling_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_profiling_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_profiling_service"))) {
		request.EnableProfilingService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_nac_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_nac_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_nac_service"))) {
		request.EnableNACService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".sxpservice"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxpservice"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxpservice"))) {
		request.Sxpservice = expandRequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSxpservice(ctx, key+".sxpservice.0", d)
	}
	if v, ok := d.GetOkExists(key + ".enable_device_admin_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_device_admin_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_device_admin_service"))) {
		request.EnableDeviceAdminService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_passive_identity_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_passive_identity_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_passive_identity_service"))) {
		request.EnablePassiveIDentityService = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSessionService(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSessionService {
	request := isegosdk.RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSessionService{}
	if v, ok := d.GetOkExists(key + ".is_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_enabled"))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".nodegroup"); !isEmptyValue(reflect.ValueOf(d.Get(key+".nodegroup"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".nodegroup"))) {
		request.Nodegroup = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSxpservice(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSxpservice {
	request := isegosdk.RequestNodeDeploymentRegisterNodeGeneralSettingsMonitoringPolicyserviceSxpservice{}
	if v, ok := d.GetOkExists(key + ".is_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_enabled"))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".user_interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".user_interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".user_interface"))) {
		request.UserInterface = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfiguration(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfiguration {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfiguration{}
	if v, ok := d.GetOkExists(key + ".netflow"); !isEmptyValue(reflect.ValueOf(d.Get(key+".netflow"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".netflow"))) {
		request.Netflow = expandRequestNodeDeploymentRegisterNodeProfileConfigurationNetflow(ctx, key+".netflow.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dhcp"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dhcp"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dhcp"))) {
		request.Dhcp = expandRequestNodeDeploymentRegisterNodeProfileConfigurationDhcp(ctx, key+".dhcp.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dhcp_span"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dhcp_span"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dhcp_span"))) {
		request.DhcpSpan = expandRequestNodeDeploymentRegisterNodeProfileConfigurationDhcpSpan(ctx, key+".dhcp_span.0", d)
	}
	if v, ok := d.GetOkExists(key + ".http"); !isEmptyValue(reflect.ValueOf(d.Get(key+".http"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".http"))) {
		request.HTTP = expandRequestNodeDeploymentRegisterNodeProfileConfigurationHTTP(ctx, key+".http.0", d)
	}
	if v, ok := d.GetOkExists(key + ".radius"); !isEmptyValue(reflect.ValueOf(d.Get(key+".radius"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".radius"))) {
		request.Radius = expandRequestNodeDeploymentRegisterNodeProfileConfigurationRadius(ctx, key+".radius.0", d)
	}
	if v, ok := d.GetOkExists(key + ".nmap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".nmap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".nmap"))) {
		request.Nmap = expandRequestNodeDeploymentRegisterNodeProfileConfigurationNmap(ctx, key+".nmap.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dns"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dns"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dns"))) {
		request.DNS = expandRequestNodeDeploymentRegisterNodeProfileConfigurationDNS(ctx, key+".dns.0", d)
	}
	if v, ok := d.GetOkExists(key + ".snmp_query"); !isEmptyValue(reflect.ValueOf(d.Get(key+".snmp_query"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".snmp_query"))) {
		request.SNMPQuery = expandRequestNodeDeploymentRegisterNodeProfileConfigurationSNMPQuery(ctx, key+".snmp_query.0", d)
	}
	if v, ok := d.GetOkExists(key + ".snmp_trap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".snmp_trap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".snmp_trap"))) {
		request.SNMPTrap = expandRequestNodeDeploymentRegisterNodeProfileConfigurationSNMPTrap(ctx, key+".snmp_trap.0", d)
	}
	if v, ok := d.GetOkExists(key + ".active_directory"); !isEmptyValue(reflect.ValueOf(d.Get(key+".active_directory"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".active_directory"))) {
		request.ActiveDirectory = expandRequestNodeDeploymentRegisterNodeProfileConfigurationActiveDirectory(ctx, key+".active_directory.0", d)
	}
	if v, ok := d.GetOkExists(key + ".pxgrid"); !isEmptyValue(reflect.ValueOf(d.Get(key+".pxgrid"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".pxgrid"))) {
		request.Pxgrid = expandRequestNodeDeploymentRegisterNodeProfileConfigurationPxgrid(ctx, key+".pxgrid.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationNetflow(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationNetflow {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationNetflow{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface"))) {
		request.Interface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".port"))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationDhcp(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationDhcp {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationDhcp{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface"))) {
		request.Interface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".port"))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationDhcpSpan(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationDhcpSpan {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationDhcpSpan{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface"))) {
		request.Interface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationHTTP(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationHTTP {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationHTTP{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface"))) {
		request.Interface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationRadius(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationRadius {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationRadius{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationNmap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationNmap {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationNmap{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationDNS(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationDNS {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationDNS{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationSNMPQuery(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationSNMPQuery {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationSNMPQuery{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".retries"); !isEmptyValue(reflect.ValueOf(d.Get(key+".retries"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".retries"))) {
		request.Retries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".timeout"))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".event_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".event_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".event_timeout"))) {
		request.EventTimeout = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationSNMPTrap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationSNMPTrap {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationSNMPTrap{}
	if v, ok := d.GetOkExists(key + ".link_trap_query"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link_trap_query"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link_trap_query"))) {
		request.LinkTrapQuery = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".mac_trap_query"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mac_trap_query"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mac_trap_query"))) {
		request.MacTrapQuery = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface"))) {
		request.Interface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".port"))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationActiveDirectory(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationActiveDirectory {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationActiveDirectory{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".days_before_rescan"); !isEmptyValue(reflect.ValueOf(d.Get(key+".days_before_rescan"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".days_before_rescan"))) {
		request.DaysBeforeRescan = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentRegisterNodeProfileConfigurationPxgrid(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationPxgrid {
	request := isegosdk.RequestNodeDeploymentRegisterNodeProfileConfigurationPxgrid{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNode {
	request := isegosdk.RequestNodeDeploymentUpdateNode{}
	request.Response = expandRequestNodeDeploymentUpdateNodeResponse(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponse(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponse {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponse{}
	if v, ok := d.GetOkExists(key + ".general_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".general_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".general_settings"))) {
		request.GeneralSettings = expandRequestNodeDeploymentUpdateNodeResponseGeneralSettings(ctx, key+".general_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".profile_configuration"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile_configuration"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile_configuration"))) {
		request.ProfileConfiguration = expandRequestNodeDeploymentUpdateNodeResponseProfileConfiguration(ctx, key+".profile_configuration.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseGeneralSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseGeneralSettings {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseGeneralSettings{}
	if v, ok := d.GetOkExists(key + ".monitoring"); !isEmptyValue(reflect.ValueOf(d.Get(key+".monitoring"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".monitoring"))) {
		request.Monitoring = expandRequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoring(ctx, key+".monitoring.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoring(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoring {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoring{}
	if v, ok := d.GetOkExists(key + ".is_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_enabled"))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".role"); !isEmptyValue(reflect.ValueOf(d.Get(key+".role"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".role"))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".other_monitoring_node"); !isEmptyValue(reflect.ValueOf(d.Get(key+".other_monitoring_node"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".other_monitoring_node"))) {
		request.OtherMonitoringNode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_mnt_dedicated"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_mnt_dedicated"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_mnt_dedicated"))) {
		request.IsMntDedicated = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".policyservice"); !isEmptyValue(reflect.ValueOf(d.Get(key+".policyservice"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".policyservice"))) {
		request.Policyservice = expandRequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyservice(ctx, key+".policyservice.0", d)
	}
	if v, ok := d.GetOkExists(key + ".enable_pxgrid"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_pxgrid"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_pxgrid"))) {
		request.EnablePXGrid = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyservice(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyservice {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyservice{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".session_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".session_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".session_service"))) {
		request.SessionService = expandRequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSessionService(ctx, key+".session_service.0", d)
	}
	if v, ok := d.GetOkExists(key + ".enable_profiling_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_profiling_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_profiling_service"))) {
		request.EnableProfilingService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_nac_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_nac_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_nac_service"))) {
		request.EnableNACService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".sxpservice"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sxpservice"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sxpservice"))) {
		request.Sxpservice = expandRequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSxpservice(ctx, key+".sxpservice.0", d)
	}
	if v, ok := d.GetOkExists(key + ".enable_device_admin_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_device_admin_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_device_admin_service"))) {
		request.EnableDeviceAdminService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_passive_identity_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_passive_identity_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_passive_identity_service"))) {
		request.EnablePassiveIDentityService = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSessionService(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSessionService {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSessionService{}
	if v, ok := d.GetOkExists(key + ".is_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_enabled"))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".nodegroup"); !isEmptyValue(reflect.ValueOf(d.Get(key+".nodegroup"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".nodegroup"))) {
		request.Nodegroup = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSxpservice(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSxpservice {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseGeneralSettingsMonitoringPolicyserviceSxpservice{}
	if v, ok := d.GetOkExists(key + ".is_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_enabled"))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".user_interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".user_interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".user_interface"))) {
		request.UserInterface = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfiguration(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfiguration {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfiguration{}
	if v, ok := d.GetOkExists(key + ".netflow"); !isEmptyValue(reflect.ValueOf(d.Get(key+".netflow"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".netflow"))) {
		request.Netflow = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationNetflow(ctx, key+".netflow.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dhcp"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dhcp"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dhcp"))) {
		request.Dhcp = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcp(ctx, key+".dhcp.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dhcp_span"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dhcp_span"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dhcp_span"))) {
		request.DhcpSpan = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcpSpan(ctx, key+".dhcp_span.0", d)
	}
	if v, ok := d.GetOkExists(key + ".http"); !isEmptyValue(reflect.ValueOf(d.Get(key+".http"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".http"))) {
		request.HTTP = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationHTTP(ctx, key+".http.0", d)
	}
	if v, ok := d.GetOkExists(key + ".radius"); !isEmptyValue(reflect.ValueOf(d.Get(key+".radius"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".radius"))) {
		request.Radius = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationRadius(ctx, key+".radius.0", d)
	}
	if v, ok := d.GetOkExists(key + ".nmap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".nmap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".nmap"))) {
		request.Nmap = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationNmap(ctx, key+".nmap.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dns"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dns"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dns"))) {
		request.DNS = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationDNS(ctx, key+".dns.0", d)
	}
	if v, ok := d.GetOkExists(key + ".snmp_query"); !isEmptyValue(reflect.ValueOf(d.Get(key+".snmp_query"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".snmp_query"))) {
		request.SNMPQuery = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPQuery(ctx, key+".snmp_query.0", d)
	}
	if v, ok := d.GetOkExists(key + ".snmp_trap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".snmp_trap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".snmp_trap"))) {
		request.SNMPTrap = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPTrap(ctx, key+".snmp_trap.0", d)
	}
	if v, ok := d.GetOkExists(key + ".active_directory"); !isEmptyValue(reflect.ValueOf(d.Get(key+".active_directory"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".active_directory"))) {
		request.ActiveDirectory = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationActiveDirectory(ctx, key+".active_directory.0", d)
	}
	if v, ok := d.GetOkExists(key + ".pxgrid"); !isEmptyValue(reflect.ValueOf(d.Get(key+".pxgrid"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".pxgrid"))) {
		request.Pxgrid = expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationPxgrid(ctx, key+".pxgrid.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationNetflow(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationNetflow {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationNetflow{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface"))) {
		request.Interface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".port"))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcp(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcp {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcp{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface"))) {
		request.Interface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".port"))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcpSpan(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcpSpan {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDhcpSpan{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface"))) {
		request.Interface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationHTTP(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationHTTP {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationHTTP{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface"))) {
		request.Interface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationRadius(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationRadius {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationRadius{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationNmap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationNmap {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationNmap{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationDNS(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDNS {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationDNS{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPQuery(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPQuery {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPQuery{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".retries"); !isEmptyValue(reflect.ValueOf(d.Get(key+".retries"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".retries"))) {
		request.Retries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".timeout"))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".event_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".event_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".event_timeout"))) {
		request.EventTimeout = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPTrap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPTrap {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationSNMPTrap{}
	if v, ok := d.GetOkExists(key + ".link_trap_query"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link_trap_query"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link_trap_query"))) {
		request.LinkTrapQuery = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".mac_trap_query"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mac_trap_query"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mac_trap_query"))) {
		request.MacTrapQuery = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".interface"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface"))) {
		request.Interface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".port"))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationActiveDirectory(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationActiveDirectory {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationActiveDirectory{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".days_before_rescan"); !isEmptyValue(reflect.ValueOf(d.Get(key+".days_before_rescan"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".days_before_rescan"))) {
		request.DaysBeforeRescan = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeDeploymentUpdateNodeResponseProfileConfigurationPxgrid(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationPxgrid {
	request := isegosdk.RequestNodeDeploymentUpdateNodeResponseProfileConfigurationPxgrid{}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNodeDeploymentGetNodes(m interface{}, response *isegosdk.ResponseNodeDeploymentGetNodes) []isegosdk.ResponseNodeDeploymentGetNodesResponse {
	var respItems []isegosdk.ResponseNodeDeploymentGetNodesResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNodeDeploymentGetNodes(m interface{}, items []isegosdk.ResponseNodeDeploymentGetNodesResponse, name string, id string) (*isegosdk.ResponseNodeDeploymentGetNodeDetailsResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNodeDeploymentGetNodeDetailsResponse
	for _, item := range items {
		if name != "" && item.Hostname == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNodeDeploymentGetNodeDetails
			getItem, _, err = client.NodeDeployment.GetNodeDetails(name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNodeDetails")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
