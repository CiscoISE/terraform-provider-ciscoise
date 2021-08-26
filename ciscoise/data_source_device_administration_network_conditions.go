package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationNetworkConditions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceAdministrationNetworkConditionsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

						"condition_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceAdministrationNetworkConditionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminNetworkConditions")

		response1, _, err := client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditions()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminNetworkConditions", err,
				"Failure at GetDeviceAdminNetworkConditions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminNetworkConditions response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceAdminNetworkConditionByID")
		vvID := vID.(string)

		response2, _, err := client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditionByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminNetworkConditionByID", err,
				"Failure at GetDeviceAdminNetworkConditionByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminNetworkConditionByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsItems(items *[]isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["link"] = flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsItemsLink(item.Link)
		respItem["name"] = item.Name
		respItem["conditions"] = flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsItemsConditions(item.Conditions)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsItemsLink(item *isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsItemsConditions(items *[]isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponseConditions) []map[string]interface{} {
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

func flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDItem(item *isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["link"] = flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDItemLink(item.Link)
	respItem["name"] = item.Name
	respItem["conditions"] = flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDItemConditions(item.Conditions)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDItemLink(item *isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDItemConditions(items *[]isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponseConditions) []map[string]interface{} {
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
