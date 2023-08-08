package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodeServicesProfilerProbeConfig() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Node Services.

- This data source retrieves the profiler probe configuration of a PSN.
`,

		ReadContext: dataSourceNodeServicesProfilerProbeConfigRead,
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Description: `hostname path parameter. Hostname of the node.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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
		},
	}
}

func dataSourceNodeServicesProfilerProbeConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vHostname := d.Get("hostname")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetProfilerProbeConfig")
		vvHostname := vHostname.(string)

		response1, restyResp1, err := client.NodeServices.GetProfilerProbeConfig(vvHostname)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetProfilerProbeConfig", err,
				"Failure at GetProfilerProbeConfig, unexpected response", ""))
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
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNodeServicesGetProfilerProbeConfigItem(item *isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["active_directory"] = flattenNodeServicesGetProfilerProbeConfigItemActiveDirectory(item.ActiveDirectory)
	respItem["dhcp"] = flattenNodeServicesGetProfilerProbeConfigItemDhcp(item.Dhcp)
	respItem["dhcp_span"] = flattenNodeServicesGetProfilerProbeConfigItemDhcpSpan(item.DhcpSpan)
	respItem["dns"] = flattenNodeServicesGetProfilerProbeConfigItemDNS(item.DNS)
	respItem["http"] = flattenNodeServicesGetProfilerProbeConfigItemHTTP(item.HTTP)
	respItem["netflow"] = flattenNodeServicesGetProfilerProbeConfigItemNetflow(item.Netflow)
	respItem["nmap"] = flattenNodeServicesGetProfilerProbeConfigItemNmap(item.Nmap)
	respItem["pxgrid"] = flattenNodeServicesGetProfilerProbeConfigItemPxgrid(item.Pxgrid)
	respItem["radius"] = flattenNodeServicesGetProfilerProbeConfigItemRadius(item.Radius)
	respItem["snmp_query"] = flattenNodeServicesGetProfilerProbeConfigItemSNMPQuery(item.SNMPQuery)
	respItem["snmp_trap"] = flattenNodeServicesGetProfilerProbeConfigItemSNMPTrap(item.SNMPTrap)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeServicesGetProfilerProbeConfigItemActiveDirectory(item *isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseActiveDirectory) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["days_before_rescan"] = item.DaysBeforeRescan

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeServicesGetProfilerProbeConfigItemDhcp(item *isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseDhcp) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interfaces"] = flattenNodeServicesGetProfilerProbeConfigItemDhcpInterfaces(item.Interfaces)
	respItem["port"] = item.Port

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeServicesGetProfilerProbeConfigItemDhcpInterfaces(items *[]isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseDhcpInterfaces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNodeServicesGetProfilerProbeConfigItemDhcpSpan(item *isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseDhcpSpan) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interfaces"] = flattenNodeServicesGetProfilerProbeConfigItemDhcpSpanInterfaces(item.Interfaces)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeServicesGetProfilerProbeConfigItemDhcpSpanInterfaces(items *[]isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseDhcpSpanInterfaces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNodeServicesGetProfilerProbeConfigItemDNS(item *isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseDNS) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["timeout"] = item.Timeout

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeServicesGetProfilerProbeConfigItemHTTP(item *isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseHTTP) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interfaces"] = flattenNodeServicesGetProfilerProbeConfigItemHTTPInterfaces(item.Interfaces)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeServicesGetProfilerProbeConfigItemHTTPInterfaces(items *[]isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseHTTPInterfaces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNodeServicesGetProfilerProbeConfigItemNetflow(item *isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseNetflow) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interfaces"] = flattenNodeServicesGetProfilerProbeConfigItemNetflowInterfaces(item.Interfaces)
	respItem["port"] = item.Port

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeServicesGetProfilerProbeConfigItemNetflowInterfaces(items *[]isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseNetflowInterfaces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNodeServicesGetProfilerProbeConfigItemNmap(items *[]isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseNmap) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenNodeServicesGetProfilerProbeConfigItemPxgrid(items *[]isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponsePxgrid) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenNodeServicesGetProfilerProbeConfigItemRadius(items *[]isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseRadius) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenNodeServicesGetProfilerProbeConfigItemSNMPQuery(item *isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseSNMPQuery) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["event_timeout"] = item.EventTimeout
	respItem["retries"] = item.Retries
	respItem["timeout"] = item.Timeout

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeServicesGetProfilerProbeConfigItemSNMPTrap(item *isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseSNMPTrap) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interfaces"] = flattenNodeServicesGetProfilerProbeConfigItemSNMPTrapInterfaces(item.Interfaces)
	respItem["link_trap_query"] = boolPtrToString(item.LinkTrapQuery)
	respItem["mac_trap_query"] = boolPtrToString(item.MacTrapQuery)
	respItem["port"] = item.Port

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNodeServicesGetProfilerProbeConfigItemSNMPTrapInterfaces(items *[]isegosdk.ResponseNodeServicesGetProfilerProbeConfigResponseSNMPTrapInterfaces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItems = append(respItems, respItem)
	}
	return respItems
}
