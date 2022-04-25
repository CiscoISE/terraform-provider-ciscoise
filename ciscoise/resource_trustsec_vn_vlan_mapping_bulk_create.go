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

func resourceTrustsecVnVLANMappingBulkCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on vnVlanMapping.
- Create VN-Vlan Mappings in bulk
`,

		CreateContext: resourceTrustsecVnVLANMappingBulkCreateCreate,
		ReadContext:   resourceTrustsecVnVLANMappingBulkCreateRead,
		DeleteContext: resourceTrustsecVnVLANMappingBulkCreateDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestVnVLANMappingBulkCreateVnVlanMappings`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Identifier of the VN-Vlan Mapping`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"is_data": &schema.Schema{
										Description:  `Flag which indicates whether the Vlan is data or voice type`,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
									},
									"is_default_vlan": &schema.Schema{
										Description:  `Flag which indicates if the Vlan is default`,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
									},
									"last_update": &schema.Schema{
										Description: `Timestamp for the last update of the VN-Vlan Mapping`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"max_value": &schema.Schema{
										Description: `Max value`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
									},
									"name": &schema.Schema{
										Description: `Name of the Vlan`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"vn_id": &schema.Schema{
										Description: `Identifier for the associated Virtual Network which is required unless its name is provided`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"vn_name": &schema.Schema{
										Description: `Name of the associated Virtual Network to be used for identity if id is not provided`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
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

func resourceTrustsecVnVLANMappingBulkCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkCreateVnVLANMappings create")
	log.Printf("[DEBUG] Missing BulkCreateVnVLANMappings create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	request1 := expandRequestTrustsecVnVLANMappingBulkCreateBulkCreateVnVLANMappings(ctx, "parameters.0", d)
	response1, restyResp1, err := client.VnVLANMapping.BulkCreateVnVLANMappings(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkCreateVnVLANMappings", err,
			"Failure at BulkCreateVnVLANMappings, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenVnVLANMappingBulkCreateVnVLANMappingsItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkCreateVnVLANMappings response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return resourceTrustsecVnVLANMappingBulkCreateRead(ctx, d, m)
}

func resourceTrustsecVnVLANMappingBulkCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceTrustsecVnVLANMappingBulkCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVnVLANMappingBulkCreate delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing TrustsecVnVLANMappingBulkCreate delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
func expandRequestTrustsecVnVLANMappingBulkCreateBulkCreateVnVLANMappings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestVnVLANMappingBulkCreateVnVLANMappings {
	request := isegosdk.RequestVnVLANMappingBulkCreateVnVLANMappings{}
	if v := expandRequestTrustsecVnVLANMappingBulkCreateBulkCreateVnVLANMappingsItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestTrustsecVnVLANMappingBulkCreateBulkCreateVnVLANMappingsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestItemVnVLANMappingBulkCreateVnVLANMappings {
	request := []isegosdk.RequestItemVnVLANMappingBulkCreateVnVLANMappings{}
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
		i := expandRequestTrustsecVnVLANMappingBulkCreateBulkCreateVnVLANMappingsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestTrustsecVnVLANMappingBulkCreateBulkCreateVnVLANMappingsItem(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemVnVLANMappingBulkCreateVnVLANMappings {
	request := isegosdk.RequestItemVnVLANMappingBulkCreateVnVLANMappings{}
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

func flattenVnVLANMappingBulkCreateVnVLANMappingsItem(item *isegosdk.ResponseVnVLANMappingBulkCreateVnVLANMappings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
