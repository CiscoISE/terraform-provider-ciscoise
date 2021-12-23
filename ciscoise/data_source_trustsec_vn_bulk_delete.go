package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceTrustsecVnBulkDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on virtualNetwork.

- Delete Virtual Network in bulk
`,

		ReadContext: dataSourceTrustsecVnBulkDeleteRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"payload": &schema.Schema{
				Description: `Array of RequestVirtualNetworkBulkDeleteVirtualNetworks`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceTrustsecVnBulkDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: BulkDeleteVirtualNetworks")
		request1 := expandRequestTrustsecVnBulkDeleteBulkDeleteVirtualNetworks(ctx, "", d)

		response1, restyResp1, err := client.VirtualNetwork.BulkDeleteVirtualNetworks(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkDeleteVirtualNetworks", err,
				"Failure at BulkDeleteVirtualNetworks, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenVirtualNetworkBulkDeleteVirtualNetworksItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkDeleteVirtualNetworks response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestTrustsecVnBulkDeleteBulkDeleteVirtualNetworks(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestVirtualNetworkBulkDeleteVirtualNetworks {
	request := isegosdk.RequestVirtualNetworkBulkDeleteVirtualNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".payload")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".payload")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".payload")))) {
		request = interfaceToSliceString(v)
	}
	return &request
}

func flattenVirtualNetworkBulkDeleteVirtualNetworksItem(item *isegosdk.ResponseVirtualNetworkBulkDeleteVirtualNetworks) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
