package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessNetworkCondition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkAccessNetworkConditionRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"condition_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"conditions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cli_dnis_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_group_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ip_addr_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"mac_addr_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
							Type:     schema.TypeString,
							Computed: true,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"conditions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cli_dnis_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_group_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ip_addr_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"mac_addr_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
							Type:     schema.TypeString,
							Computed: true,
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
		log.Printf("[DEBUG] Selected method 1: GetNetworkAccessNetworkConditions")

		response1, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditions()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessNetworkConditions", err,
				"Failure at GetNetworkAccessNetworkConditions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

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
		log.Printf("[DEBUG] Selected method 2: GetNetworkAccessNetworkConditionByID")
		vvID := vID.(string)

		response2, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessNetworkConditionByID", err,
				"Failure at GetNetworkAccessNetworkConditionByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

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
		respItem["ip_addr_list"] = item.IPAddrList
		respItem["mac_addr_list"] = item.MacAddrList
		respItem["device_group_list"] = item.DeviceGroupList
		respItem["device_list"] = item.DeviceList
	}
	return respItems

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
		respItem["ip_addr_list"] = item.IPAddrList
		respItem["mac_addr_list"] = item.MacAddrList
		respItem["device_group_list"] = item.DeviceGroupList
		respItem["device_list"] = item.DeviceList
	}
	return respItems

}
