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
func dataSourceTrustsecVnVLANMappingBulkUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on vnVlanMapping.

- Update VN-Vlan Mappings in bulk
`,

		ReadContext: dataSourceTrustsecVnVLANMappingBulkUpdateRead,
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
				Description: `Array of RequestVnVLANMappingBulkUpdateVnVlanMappings`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Identifier of the VN-Vlan Mapping`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"is_data": &schema.Schema{
							Description:  `Flag which indicates whether the Vlan is data or voice type`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"is_default_vlan": &schema.Schema{
							Description:  `Flag which indicates if the Vlan is default`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"last_update": &schema.Schema{
							Description: `Timestamp for the last update of the VN-Vlan Mapping`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"max_value": &schema.Schema{
							Description: `Max value`,
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"name": &schema.Schema{
							Description: `Name of the Vlan`,
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

func dataSourceTrustsecVnVLANMappingBulkUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BulkUpdateVnVLANMappings")
		request1 := expandRequestTrustsecVnVLANMappingBulkUpdateBulkUpdateVnVLANMappings(ctx, "", d)

		response1, restyResp1, err := client.VnVLANMapping.BulkUpdateVnVLANMappings(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkUpdateVnVLANMappings", err,
				"Failure at BulkUpdateVnVLANMappings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenVnVLANMappingBulkUpdateVnVLANMappingsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkUpdateVnVLANMappings response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestTrustsecVnVLANMappingBulkUpdateBulkUpdateVnVLANMappings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestVnVLANMappingBulkUpdateVnVLANMappings {
	request := isegosdk.RequestVnVLANMappingBulkUpdateVnVLANMappings{}
	if v := expandRequestTrustsecVnVLANMappingBulkUpdateBulkUpdateVnVLANMappingsItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestTrustsecVnVLANMappingBulkUpdateBulkUpdateVnVLANMappingsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestItemVnVLANMappingBulkUpdateVnVLANMappings {
	request := []isegosdk.RequestItemVnVLANMappingBulkUpdateVnVLANMappings{}
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
		i := expandRequestTrustsecVnVLANMappingBulkUpdateBulkUpdateVnVLANMappingsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestTrustsecVnVLANMappingBulkUpdateBulkUpdateVnVLANMappingsItem(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemVnVLANMappingBulkUpdateVnVLANMappings {
	request := isegosdk.RequestItemVnVLANMappingBulkUpdateVnVLANMappings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_data")))) {
		request.IsData = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_vlan")))) {
		request.IsDefaultVLAN = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_id")))) {
		request.VnID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_name")))) {
		request.VnName = interfaceToString(v)
	}
	return &request
}

func flattenVnVLANMappingBulkUpdateVnVLANMappingsItem(item *isegosdk.ResponseVnVLANMappingBulkUpdateVnVLANMappings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
