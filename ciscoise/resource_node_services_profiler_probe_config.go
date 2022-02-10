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

func resourceNodeServicesProfilerProbeConfig() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Node Services.

- This resource updates the profiler probe configuration of a PSN.
Set probe value as false to disable probe.
Ex: Below payload will disable NMAP, PxGrid and SNMPTRAP probes
{
  "activeDirectory": { "daysBeforeRescan": 1 },
  "dhcp": { "interfaces": "[{"interface":"GigabitEthernet 0"}]", "port": 0 },
  "dhcpSpan": { "interfaces": "[{"interface":"GigabitEthernet 0"}]" },
  "dns": { "timeout": 2 },
  "http": { "interfaces": "[{"interface":"GigabitEthernet 0"}]" },
  "netflow": { "interfaces": "[{"interface":"GigabitEthernet 0"}]", "port": 0 },
  "nmap": false,
  "pxgrid": false,
  "radius": "true",
  "snmpQuery": { "eventTimeout": 30, "retries": 2, "timeout": 1000 },
  "snmpTrap": null
}

`,

		CreateContext: resourceNodeServicesProfilerProbeConfigCreate,
		ReadContext:   resourceNodeServicesProfilerProbeConfigRead,
		UpdateContext: resourceNodeServicesProfilerProbeConfigUpdate,
		DeleteContext: resourceNodeServicesProfilerProbeConfigDelete,
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
						"hostname": &schema.Schema{
							Description: `hostname path parameter.`,
							Type:        schema.TypeString,
							Required:    true,
						},
						"active_directory": &schema.Schema{
							Description: `The Active Directory probe queries the Active Directory for Windows information.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"days_before_rescan": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"dhcp": &schema.Schema{
							Description: `The DHCP probe listens for DHCP packets from IP helpers.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"dhcp_span": &schema.Schema{
							Description: `The DHCP SPAN probe collects DHCP packets.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"dns": &schema.Schema{
							Description: `The DNS probe performs a DNS lookup for the FQDN.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"timeout": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"http": &schema.Schema{
							Description: `The HTTP probe receives and parses HTTP packets.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"netflow": &schema.Schema{
							Description: `The NetFlow probe collects the NetFlow packets that are sent to it from routers.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"nmap": &schema.Schema{
							Description: `The NMAP probe scans endpoints for open ports and OS.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"pxgrid": &schema.Schema{
							Description: `The pxGrid probe fetches attributes of MAC address or IP address as a subscriber from the pxGrid queue.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"radius": &schema.Schema{
							Description: `The RADIUS probe collects RADIUS session attributes as well as CDP, LLDP, DHCP, HTTP, and MDM attributes from IOS Sensors.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"snmp_query": &schema.Schema{
							Description: `The SNMP query probe collects details from network devices such as interface, CDP, LLDP, and ARP.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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
							Description: `The SNMP trap probe receives linkup, linkdown, and MAC notification traps from network devices.`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": &schema.Schema{
							Description: `hostname path parameter. Hostname of the node.`,
							Type:        schema.TypeString,
							Required:    true,
						},
						"active_directory": &schema.Schema{
							Description: `The Active Directory probe queries the Active Directory for Windows information.`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"days_before_rescan": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"dhcp": &schema.Schema{
							Description: `The DHCP probe listens for DHCP packets from IP helpers.`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"dhcp_span": &schema.Schema{
							Description: `The DHCP SPAN probe collects DHCP packets.`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"dns": &schema.Schema{
							Description: `The DNS probe performs a DNS lookup for the FQDN.`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"timeout": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"http": &schema.Schema{
							Description: `The HTTP probe receives and parses HTTP packets.`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"netflow": &schema.Schema{
							Description: `The NetFlow probe collects the NetFlow packets that are sent to it from routers.`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"nmap": &schema.Schema{
							Description: `The NMAP probe scans endpoints for open ports and OS.
							If set to true, it will activate the NMAP probe.
							If set to false, it will deactivate the NMAP probe.
							Finally, if set to empty string or no-set (default), it will maintain the NMAP probe state.
							`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"pxgrid": &schema.Schema{
							Description: `The pxGrid probe fetches attributes of MAC address or IP address as a subscriber from the pxGrid queue.
							If set to true, it will activate the pxGrid probe.
							If set to false, it will deactivate the pxGrid probe.
							Finally, if set to empty string or no-set (default), it will maintain the pxGrid probe state.
							`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"radius": &schema.Schema{
							Description: `The RADIUS probe collects RADIUS session attributes as well as CDP, LLDP, DHCP, HTTP, and MDM attributes from IOS Sensors.
							If set to true, it will activate the RADIUS probe.
							If set to false, it will deactivate the RADIUS probe.
							Finally, if set to empty string or no-set (default), it will maintain the RADIUS probe state.
							`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"snmp_query": &schema.Schema{
							Description: `The SNMP query probe collects details from network devices such as interface, CDP, LLDP, and ARP.`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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
							Description: `The SNMP trap probe receives linkup, linkdown, and MAC notification traps from network devices.`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
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
	}
}

func resourceNodeServicesProfilerProbeConfigCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeServicesProfilerProbeConfig create")
	log.Printf("[DEBUG] Missing NodeServicesProfilerProbeConfig create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	vHostname := interfaceToString(resourceItem["hostname"])
	vvHostname := vHostname
	log.Printf("[DEBUG] Name used for update operation %s", vHostname)
	request1 := expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfig(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, restyResp1, err := client.NodeServices.SetProfilerProbeConfig(vvHostname, request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing SetProfilerProbeConfig", err, restyResp1.String(),
				"Failure at SetProfilerProbeConfig, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing SetProfilerProbeConfig", err,
			"Failure at SetProfilerProbeConfig, unexpected response", ""))
		return diags
	}
	resourceMap["hostname"] = interfaceToString(resourceItem["hostname"])
	d.SetId(joinResourceID(resourceMap))
	return resourceNodeServicesProfilerProbeConfigRead(ctx, d, m)
}

func resourceNodeServicesProfilerProbeConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeServicesProfilerProbeConfig read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostname, _ := resourceMap["hostname"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetProfilerProbeConfig")
		vvHostname := vHostname

		response1, restyResp1, err := client.NodeServices.GetProfilerProbeConfig(vvHostname)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNodeServicesGetProfilerProbeConfigItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetProfilerProbeConfig response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNodeServicesProfilerProbeConfigUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeServicesProfilerProbeConfig update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostname, _ := resourceMap["hostname"]

	var vvHostname string = vHostname
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vHostname)
		request1 := expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfig(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.NodeServices.SetProfilerProbeConfig(vvHostname, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SetProfilerProbeConfig", err, restyResp1.String(),
					"Failure at SetProfilerProbeConfig, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SetProfilerProbeConfig", err,
				"Failure at SetProfilerProbeConfig, unexpected response", ""))
			return diags
		}
		d.Set("last_updated", getUnixTimeString())
	}

	return resourceNodeServicesProfilerProbeConfigRead(ctx, d, m)
}

func resourceNodeServicesProfilerProbeConfigDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeServicesProfilerProbeConfig delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing NodeServicesProfilerProbeConfig delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfig(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfig {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfig{
		ActiveDirectory: nil,
		Dhcp:            nil,
		DhcpSpan:        nil,
		DNS:             nil,
		HTTP:            nil,
		Netflow:         nil,
		Nmap:            nil,
		Pxgrid:          nil,
		Radius:          nil,
		SNMPQuery:       nil,
		SNMPTrap:        nil,
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".active_directory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".active_directory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".active_directory")))) {
		request.ActiveDirectory = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigActiveDirectory(ctx, key+".active_directory.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp")))) {
		request.Dhcp = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcp(ctx, key+".dhcp.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_span")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_span")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_span")))) {
		request.DhcpSpan = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcpSpan(ctx, key+".dhcp_span.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns")))) {
		request.DNS = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDNS(ctx, key+".dns.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http")))) {
		request.HTTP = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigHTTP(ctx, key+".http.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netflow")))) {
		request.Netflow = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigNetflow(ctx, key+".netflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".nmap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".nmap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".nmap")))) {
		request.Nmap = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigNmapArray(ctx, key+".nmap", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pxgrid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pxgrid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pxgrid")))) {
		request.Pxgrid = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigPxgridArray(ctx, key+".pxgrid", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radius")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radius")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radius")))) {
		request.Radius = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigRadiusArray(ctx, key+".radius", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_query")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_query")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_query")))) {
		request.SNMPQuery = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigSNMPQuery(ctx, key+".snmp_query.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_trap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_trap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_trap")))) {
		request.SNMPTrap = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigSNMPTrap(ctx, key+".snmp_trap.0", d)
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigActiveDirectory(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigActiveDirectory {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigActiveDirectory{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".days_before_rescan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".days_before_rescan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".days_before_rescan")))) {
		request.DaysBeforeRescan = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcp(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcp {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcp{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interfaces")))) {
		request.Interfaces = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcpInterfacesArray(ctx, key+".interfaces", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcpInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcpInterfaces {
	request := []isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcpInterfaces{}
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
		i := expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcpInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcpInterfaces(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcpInterfaces {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcpInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface")))) {
		request.Interface = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcpSpan(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcpSpan {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcpSpan{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interfaces")))) {
		request.Interfaces = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcpSpanInterfacesArray(ctx, key+".interfaces", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcpSpanInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcpSpanInterfaces {
	request := []isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcpSpanInterfaces{}
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
		i := expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcpSpanInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDhcpSpanInterfaces(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcpSpanInterfaces {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigDhcpSpanInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface")))) {
		request.Interface = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigDNS(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigDNS {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigDNS{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timeout")))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigHTTP(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigHTTP {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigHTTP{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interfaces")))) {
		request.Interfaces = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigHTTPInterfacesArray(ctx, key+".interfaces", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigHTTPInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNodeServicesSetProfilerProbeConfigHTTPInterfaces {
	request := []isegosdk.RequestNodeServicesSetProfilerProbeConfigHTTPInterfaces{}
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
		i := expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigHTTPInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigHTTPInterfaces(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigHTTPInterfaces {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigHTTPInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface")))) {
		request.Interface = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigNetflow(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigNetflow {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigNetflow{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interfaces")))) {
		request.Interfaces = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigNetflowInterfacesArray(ctx, key+".interfaces", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigNetflowInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNodeServicesSetProfilerProbeConfigNetflowInterfaces {
	request := []isegosdk.RequestNodeServicesSetProfilerProbeConfigNetflowInterfaces{}
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
		i := expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigNetflowInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigNetflowInterfaces(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigNetflowInterfaces {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigNetflowInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface")))) {
		request.Interface = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigNmapArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNodeServicesSetProfilerProbeConfigNmap {
	request := []isegosdk.RequestNodeServicesSetProfilerProbeConfigNmap{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	if v := interfaceToBoolPtr(o); v != nil && *v {
		return &request
	}
	return nil
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigPxgridArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNodeServicesSetProfilerProbeConfigPxgrid {
	request := []isegosdk.RequestNodeServicesSetProfilerProbeConfigPxgrid{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	if v := interfaceToBoolPtr(o); v != nil && *v {
		return &request
	}
	return nil
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigPxgrid(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigPxgrid {
	var request isegosdk.RequestNodeServicesSetProfilerProbeConfigPxgrid
	keyValue := d.Get(fixKeyAccess(key))
	request = requestStringToInterface(interfaceToString(keyValue))
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigRadiusArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNodeServicesSetProfilerProbeConfigRadius {
	request := []isegosdk.RequestNodeServicesSetProfilerProbeConfigRadius{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	if v := interfaceToBoolPtr(o); v != nil && *v {
		return &request
	}
	return nil
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigSNMPQuery(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigSNMPQuery {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigSNMPQuery{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_timeout")))) {
		request.EventTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".retries")))) {
		request.Retries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timeout")))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigSNMPTrap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigSNMPTrap {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigSNMPTrap{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interfaces")))) {
		request.Interfaces = expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigSNMPTrapInterfacesArray(ctx, key+".interfaces", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".link_trap_query")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".link_trap_query")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".link_trap_query")))) {
		request.LinkTrapQuery = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_trap_query")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_trap_query")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_trap_query")))) {
		request.MacTrapQuery = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigSNMPTrapInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNodeServicesSetProfilerProbeConfigSNMPTrapInterfaces {
	request := []isegosdk.RequestNodeServicesSetProfilerProbeConfigSNMPTrapInterfaces{}
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
		i := expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigSNMPTrapInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNodeServicesProfilerProbeConfigSetProfilerProbeConfigSNMPTrapInterfaces(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetProfilerProbeConfigSNMPTrapInterfaces {
	request := isegosdk.RequestNodeServicesSetProfilerProbeConfigSNMPTrapInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface")))) {
		request.Interface = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
