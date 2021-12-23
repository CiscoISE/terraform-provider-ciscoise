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
func dataSourceSgMappingBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMapping.

- This data source action allows the client to submit the bulk request.
`,

		ReadContext: dataSourceSgMappingBulkRequestRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"operation_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_media_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceSgMappingBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: BulkRequestForIPToSgtMapping")
		request1 := expandRequestSgMappingBulkRequestBulkRequestForIPToSgtMapping(ctx, "", d)

		response1, err := client.IPToSgtMapping.BulkRequestForIPToSgtMapping(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkRequestForIPToSgtMapping", err,
				"Failure at BulkRequestForIPToSgtMapping, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkRequestForIPToSgtMapping response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSgMappingBulkRequestBulkRequestForIPToSgtMapping(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingBulkRequestForIPToSgtMapping {
	request := isegosdk.RequestIPToSgtMappingBulkRequestForIPToSgtMapping{}
	request.SgMappingBulkRequest = expandRequestSgMappingBulkRequestBulkRequestForIPToSgtMappingSgMappingBulkRequest(ctx, key, d)
	return &request
}

func expandRequestSgMappingBulkRequestBulkRequestForIPToSgtMappingSgMappingBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingBulkRequestForIPToSgtMappingSgMappingBulkRequest {
	request := isegosdk.RequestIPToSgtMappingBulkRequestForIPToSgtMappingSgMappingBulkRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation_type")))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_media_type")))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
