package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUserEquipmentSubscriberInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on User Equipment.

- Get user equipments associated with a subscriber GUID
`,

		ReadContext: dataSourceUserEquipmentSubscriberInfoRead,
		Schema: map[string]*schema.Schema{
			"subscriber_id": &schema.Schema{
				Description: `subscriberId path parameter. Unique ID for a subscriber object`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"create_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `Description for User Equipment`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"device_group": &schema.Schema{
							Description: `Device or Endpoint Group`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"imei": &schema.Schema{
							Description: `IMEI for User Equipment`,
							Type:        schema.TypeString,
							Computed:    true,
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
						"update_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceUserEquipmentSubscriberInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vSubscriberID := d.Get("subscriber_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetUserEquipmentsBySubscriberID")
		vvSubscriberID := vSubscriberID.(string)

		response1, restyResp1, err := client.UserEquipment.GetUserEquipmentsBySubscriberID(vvSubscriberID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetUserEquipmentsBySubscriberID", err,
				"Failure at GetUserEquipmentsBySubscriberID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenUserEquipmentGetUserEquipmentsBySubscriberIDItemsResponse(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetUserEquipmentsBySubscriberID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUserEquipmentGetUserEquipmentsBySubscriberIDItemsResponse(items *[]isegosdk.ResponseUserEquipmentGetUserEquipmentsBySubscriberIDResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["description"] = item.Description
		respItem["device_group"] = item.DeviceGroup
		respItem["imei"] = item.Imei
		respItem["create_time"] = item.CreateTime
		respItem["update_time"] = item.UpdateTime
		respItem["id"] = item.ID
		respItem["link"] = flattenUserEquipmentGetUserEquipmentsBySubscriberIDItemsResponseLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenUserEquipmentGetUserEquipmentsBySubscriberIDItemsResponseLink(item *isegosdk.ResponseUserEquipmentGetUserEquipmentsBySubscriberIDResponseLink) []map[string]interface{} {
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
