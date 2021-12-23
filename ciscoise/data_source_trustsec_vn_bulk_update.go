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

// dataSourceAction
func dataSourceTrustsecVnBulkUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on virtualNetwork.

- Update Virtual Network in bulk
`,

		ReadContext: dataSourceTrustsecVnBulkUpdateRead,
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
				Description: `Array of RequestVirtualNetworkBulkUpdateVirtualNetworks`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_attributes": &schema.Schema{
							Description: `JSON String of additional attributes for the Virtual Network`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier of the Virtual Network`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"last_update": &schema.Schema{
							Description: `Timestamp for the last update of the Virtual Network`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": &schema.Schema{
							Description: `Name of the Virtual Network`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTrustsecVnBulkUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: BulkUpdateVirtualNetworks")
		request1 := expandRequestTrustsecVnBulkUpdateBulkUpdateVirtualNetworks(ctx, "", d)

		response1, restyResp1, err := client.VirtualNetwork.BulkUpdateVirtualNetworks(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkUpdateVirtualNetworks", err,
				"Failure at BulkUpdateVirtualNetworks, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenVirtualNetworkBulkUpdateVirtualNetworksItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkUpdateVirtualNetworks response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestTrustsecVnBulkUpdateBulkUpdateVirtualNetworks(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestVirtualNetworkBulkUpdateVirtualNetworks {
	request := isegosdk.RequestVirtualNetworkBulkUpdateVirtualNetworks{}
	if v := expandRequestTrustsecVnBulkUpdateBulkUpdateVirtualNetworksItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestTrustsecVnBulkUpdateBulkUpdateVirtualNetworksItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestItemVirtualNetworkBulkUpdateVirtualNetworks {
	request := []isegosdk.RequestItemVirtualNetworkBulkUpdateVirtualNetworks{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestTrustsecVnBulkUpdateBulkUpdateVirtualNetworksItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestTrustsecVnBulkUpdateBulkUpdateVirtualNetworksItem(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemVirtualNetworkBulkUpdateVirtualNetworks {
	request := isegosdk.RequestItemVirtualNetworkBulkUpdateVirtualNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_attributes")))) {
		request.AdditionalAttributes = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}

func flattenVirtualNetworkBulkUpdateVirtualNetworksItem(item *isegosdk.ResponseVirtualNetworkBulkUpdateVirtualNetworks) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
