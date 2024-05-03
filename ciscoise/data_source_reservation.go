package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReservation() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SgtRangeReservation.

- Get all the reserved Security Group tag ranges in ISE.

- Get the reserved range of SGT for the specific client which is passed in the URL.
`,

		ReadContext: dataSourceReservationRead,
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Description: `clientID path parameter. Unique name for a Client`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"page": &schema.Schema{
				Description: `page query parameter. Page number`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. Number of objects returned per page`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"client_id": &schema.Schema{
							Description: `Unique ID of the given client`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"client_name": &schema.Schema{
							Description: `Name of the given client`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"end_index": &schema.Schema{
							Description: `End index of the reserved range`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"start_index": &schema.Schema{
							Description: `Start index of the reserved range`,
							Type:        schema.TypeInt,
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

						"client_id": &schema.Schema{
							Description: `Unique ID of the given client`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"client_name": &schema.Schema{
							Description: `Name of the given client`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"end_index": &schema.Schema{
							Description: `End index of the reserved range`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"start_index": &schema.Schema{
							Description: `Start index of the reserved range`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceReservationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vClientID, okClientID := d.GetOk("client_id")

	method1 := []bool{okPage, okSize}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okClientID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSgtReservedRanges")
		queryParams1 := isegosdk.GetSgtReservedRangesQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, restyResp1, err := client.SgtRangeReservation.GetSgtReservedRanges(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSgtReservedRanges", err,
				"Failure at GetSgtReservedRanges, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSgtRangeReservationGetSgtReservedRangesItemsResponse(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSgtReservedRanges response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSgtReservedRange")
		vvClientID := vClientID.(string)

		response2, restyResp2, err := client.SgtRangeReservation.GetSgtReservedRange(vvClientID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSgtReservedRange", err,
				"Failure at GetSgtReservedRange, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSgtRangeReservationGetSgtReservedRangeItemResponse(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSgtReservedRange response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSgtRangeReservationGetSgtReservedRangesItemsResponse(items *[]isegosdk.ResponseSgtRangeReservationGetSgtReservedRangesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["client_id"] = item.ClientID
		respItem["client_name"] = item.ClientName
		respItem["end_index"] = item.EndIndex
		respItem["start_index"] = item.StartIndex
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSgtRangeReservationGetSgtReservedRangeItemResponse(item *isegosdk.ResponseSgtRangeReservationGetSgtReservedRangeResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["client_id"] = item.ClientID
	respItem["client_name"] = item.ClientName
	respItem["end_index"] = item.EndIndex
	respItem["start_index"] = item.StartIndex

	return []map[string]interface{}{
		respItem,
	}

}
