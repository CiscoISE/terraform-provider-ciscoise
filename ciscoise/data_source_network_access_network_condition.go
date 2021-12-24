package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessNetworkCondition() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Access - Network Conditions.

- Network Access Returns a list of network conditions.

- Network Access Returns a network condition.
`,

		ReadContext: dataSourceNetworkAccessNetworkConditionRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Condition id`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"condition_type": &schema.Schema{
							Description: `This field determines the content of the conditions field`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"conditions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cli_dnis_list": &schema.Schema{
										Description: `<p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"condition_type": &schema.Schema{
										Description: `This field determines the content of the conditions field`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_group_list": &schema.Schema{
										Description: `<p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_list": &schema.Schema{
										Description: `<p>This field should contain Device-Name,port-number. The device name must be the same as the name field in a Network Device object.<br> Line format - Device Name,Port</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_addr_list": &schema.Schema{
										Description: `<p>This field should contain IP-address-or-subnet,port number<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet,Port</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
									"mac_addr_list": &schema.Schema{
										Description: `<p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": &schema.Schema{
										Description: `Network Condition name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
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
							Description: `Network Condition name`,
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

						"condition_type": &schema.Schema{
							Description: `This field determines the content of the conditions field`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"conditions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cli_dnis_list": &schema.Schema{
										Description: `<p>This field should contain a Caller ID (CLI), comma, and Called ID (DNIS).<br> Line format -  Caller ID (CLI), Called ID (DNIS)</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"condition_type": &schema.Schema{
										Description: `This field determines the content of the conditions field`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_group_list": &schema.Schema{
										Description: `<p>This field should contain a tuple with NDG Root, comma, and an NDG (that it under the root).<br> Line format - NDG Root Name, NDG, Port</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_list": &schema.Schema{
										Description: `<p>This field should contain Device-Name,port-number. The device name must be the same as the name field in a Network Device object.<br> Line format - Device Name,Port</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_addr_list": &schema.Schema{
										Description: `<p>This field should contain IP-address-or-subnet,port number<br> IP address can be IPV4 format (n.n.n.n) or IPV6 format (n:n:n:n:n:n:n:n).<br> IP subnet can be IPV4 format (n.n.n.n/m) or IPV6 format (n:n:n:n:n:n:n:n/m).<br> Line format - IP Address or subnet,Port</p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
									"mac_addr_list": &schema.Schema{
										Description: `<p>This field should contain Endstation MAC address, comma, and Destination MAC addresses.<br> Each Max address must include twelve hexadecimal digits using formats nn:nn:nn:nn:nn:nn or nn-nn-nn-nn-nn-nn or nnnn.nnnn.nnnn or nnnnnnnnnnnn.<br> Line format - Endstation MAC,Destination MAC </p>`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": &schema.Schema{
										Description: `Network Condition name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
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
							Description: `Network Condition name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkAccessNetworkConditionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessNetworkConditions")

		response1, restyResp1, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditions()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessNetworkConditions", err,
				"Failure at GetNetworkAccessNetworkConditions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessNetworkConditions response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessNetworkConditionByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessNetworkConditionByID", err,
				"Failure at GetNetworkAccessNetworkConditionByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessNetworkConditionByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsItems(items *[]isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["link"] = flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsItemsLink(item.Link)
		respItem["name"] = item.Name
		respItem["conditions"] = flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsItemsConditions(item.Conditions)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsItemsLink(item *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsItemsConditions(items *[]isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponseConditions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["cli_dnis_list"] = item.CliDnisList
		respItem["condition_type"] = item.ConditionType
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["ip_addr_list"] = item.IPAddrList
		respItem["link"] = flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsItemsConditionsLink(item.Link)
		respItem["mac_addr_list"] = item.MacAddrList
		respItem["name"] = item.Name
		respItem["device_group_list"] = item.DeviceGroupList
		respItem["device_list"] = item.DeviceList
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsItemsConditionsLink(item *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponseConditionsLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItem(item *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["link"] = flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItemLink(item.Link)
	respItem["name"] = item.Name
	respItem["conditions"] = flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItemConditions(item.Conditions)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItemLink(item *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItemConditions(items *[]isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponseConditions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["cli_dnis_list"] = item.CliDnisList
		respItem["condition_type"] = item.ConditionType
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["ip_addr_list"] = item.IPAddrList
		respItem["link"] = flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItemConditionsLink(item.Link)
		respItem["mac_addr_list"] = item.MacAddrList
		respItem["name"] = item.Name
		respItem["device_group_list"] = item.DeviceGroupList
		respItem["device_list"] = item.DeviceList
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItemConditionsLink(item *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponseConditionsLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
