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
func dataSourceTrustsecSgVnMappingBulkUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on sgVnMapping.

- Update SG-VN Mappings in bulk
`,

		ReadContext: dataSourceTrustsecSgVnMappingBulkUpdateRead,
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
				Description: `Array of RequestSgVnMappingBulkUpdateSgVnMappings`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Identifier of the SG-VN mapping`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"last_update": &schema.Schema{
							Description: `Timestamp for the last update of the SG-VN mapping`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"sg_name": &schema.Schema{
							Description: `Name of the associated Security Group to be used for identity if id is not provided`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"sgt_id": &schema.Schema{
							Description: `Identifier of the associated Security Group which is required unless its name is provided`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"vn_id": &schema.Schema{
							Description: `Identifier for the associated Virtual Network which is required unless its name is provided`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"vn_name": &schema.Schema{
							Description: `Name of the associated Virtual Network to be used for identity if id is not provided`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTrustsecSgVnMappingBulkUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BulkUpdateSgVnMappings")
		request1 := expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappings(ctx, "", d)

		response1, restyResp1, err := client.SgVnMapping.BulkUpdateSgVnMappings(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkUpdateSgVnMappings", err,
				"Failure at BulkUpdateSgVnMappings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSgVnMappingBulkUpdateSgVnMappingsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkUpdateSgVnMappings response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSgVnMappingBulkUpdateSgVnMappings {
	request := isegosdk.RequestSgVnMappingBulkUpdateSgVnMappings{}
	if v := expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappingsItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappingsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestItemSgVnMappingBulkUpdateSgVnMappings {
	request := []isegosdk.RequestItemSgVnMappingBulkUpdateSgVnMappings{}
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
		i := expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappingsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappingsItem(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemSgVnMappingBulkUpdateSgVnMappings {
	request := isegosdk.RequestItemSgVnMappingBulkUpdateSgVnMappings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sg_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sg_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sg_name")))) {
		request.SgName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt_id")))) {
		request.SgtID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_id")))) {
		request.VnID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_name")))) {
		request.VnName = interfaceToString(v)
	}
	return &request
}

func flattenSgVnMappingBulkUpdateSgVnMappingsItem(item *isegosdk.ResponseSgVnMappingBulkUpdateSgVnMappings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
