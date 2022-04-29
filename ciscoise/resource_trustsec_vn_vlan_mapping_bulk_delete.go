package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrustsecVnVLANMappingBulkDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on vnVlanMapping.
- Delete VN-Vlan Mappings in bulk
`,

		CreateContext: resourceTrustsecVnVLANMappingBulkDeleteCreate,
		ReadContext:   resourceTrustsecVnVLANMappingBulkDeleteRead,
		DeleteContext: resourceTrustsecVnVLANMappingBulkDeleteDelete,

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
							Description: `Array of RequestVnVLANMappingBulkDeleteVnVlanMappings`,
							Type:        schema.TypeList,
							ForceNew:    true,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceTrustsecVnVLANMappingBulkDeleteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkDeleteVnVLANMappings create")
	log.Printf("[DEBUG] Missing BulkDeleteVnVLANMappings create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	request1 := expandRequestTrustsecVnVLANMappingBulkDeleteBulkDeleteVnVLANMappings(ctx, "parameters.0", d)
	response1, restyResp1, err := client.VnVLANMapping.BulkDeleteVnVLANMappings(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkDeleteVnVLANMappings", err,
			"Failure at BulkDeleteVnVLANMappings, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenVnVLANMappingBulkDeleteVnVLANMappingsItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkDeleteVnVLANMappings response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceTrustsecVnVLANMappingBulkDeleteRead(ctx, d, m)
}

func resourceTrustsecVnVLANMappingBulkDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceTrustsecVnVLANMappingBulkDeleteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVnVLANMappingBulkDelete delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing TrustsecVnVLANMappingBulkDelete delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestTrustsecVnVLANMappingBulkDeleteBulkDeleteVnVLANMappings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestVnVLANMappingBulkDeleteVnVLANMappings {
	request := isegosdk.RequestVnVLANMappingBulkDeleteVnVLANMappings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".payload")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".payload")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".payload")))) {
		request = interfaceToSliceString(v)
	}
	return &request
}

func flattenVnVLANMappingBulkDeleteVnVLANMappingsItem(item *isegosdk.ResponseVnVLANMappingBulkDeleteVnVLANMappings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
