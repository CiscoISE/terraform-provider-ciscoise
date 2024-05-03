package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUserEquipmentImeiInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on User Equipment.

- Get a user equipment based on the IMEI
`,

		ReadContext: dataSourceUserEquipmentImeiInfoRead,
		Schema: map[string]*schema.Schema{
			"imei": &schema.Schema{
				Description: `imei path parameter. IMEI for the user equipment object`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"imei": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rel": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"href": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
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
	}
}

func dataSourceUserEquipmentImeiInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vImei := d.Get("imei")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetUserEquipmentByIMEI")
		vvImei := vImei.(string)

		response1, restyResp1, err := client.UserEquipment.GetUserEquipmentByIMEI(vvImei)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetUserEquipmentByIMEI", err,
				"Failure at GetUserEquipmentByIMEI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenUserEquipmentGetUserEquipmentByIMEIItemResponse(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetUserEquipmentByIMEI response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUserEquipmentGetUserEquipmentByIMEIItemResponse(item *isegosdk.ResponseUserEquipmentGetUserEquipmentByIMEIResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["description"] = item.Description
	respItem["device_group"] = item.DeviceGroup
	respItem["imei"] = item.Imei
	respItem["create_time"] = item.CreateTime
	respItem["update_time"] = item.UpdateTime
	respItem["id"] = item.ID
	respItem["link"] = flattenUserEquipmentGetUserEquipmentByIMEIItemResponseLink(item.Link)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenUserEquipmentGetUserEquipmentByIMEIItemResponseLink(item *isegosdk.ResponseUserEquipmentGetUserEquipmentByIMEIResponseLink) []map[string]interface{} {
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
