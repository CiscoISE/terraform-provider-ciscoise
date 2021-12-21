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
func dataSourceTrustsecSgVnMappingBulkDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on sgVnMapping.

- Delete SG-VN Mappings in bulk
`,

		ReadContext: dataSourceTrustsecSgVnMappingBulkDeleteRead,
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
				Description: `Array of RequestSgVnMappingBulkDeleteSgVnMappings`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceTrustsecSgVnMappingBulkDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BulkDeleteSgVnMappings")
		request1 := expandRequestTrustsecSgVnMappingBulkDeleteBulkDeleteSgVnMappings(ctx, "", d)

		response1, restyResp1, err := client.SgVnMapping.BulkDeleteSgVnMappings(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkDeleteSgVnMappings", err,
				"Failure at BulkDeleteSgVnMappings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSgVnMappingBulkDeleteSgVnMappingsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkDeleteSgVnMappings response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestTrustsecSgVnMappingBulkDeleteBulkDeleteSgVnMappings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSgVnMappingBulkDeleteSgVnMappings {
	request := isegosdk.RequestSgVnMappingBulkDeleteSgVnMappings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".payload")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".payload")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".payload")))) {
		request = interfaceToSliceString(v)
	}
	return &request
}

func flattenSgVnMappingBulkDeleteSgVnMappingsItem(item *isegosdk.ResponseSgVnMappingBulkDeleteSgVnMappings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
