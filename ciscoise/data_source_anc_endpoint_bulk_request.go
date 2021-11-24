package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceAncEndpointBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ANCEndpoint.

- This data source action allows the client to submit the bulk request.
`,

		ReadContext: dataSourceAncEndpointBulkRequestRead,
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

func dataSourceAncEndpointBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BulkRequestForAncEndpoint")
		request1 := expandRequestAncEndpointBulkRequestBulkRequestForAncEndpoint(ctx, "", d)

		response1, err := client.AncEndpoint.BulkRequestForAncEndpoint(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkRequestForAncEndpoint", err,
				"Failure at BulkRequestForAncEndpoint, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkRequestForAncEndpoint response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestAncEndpointBulkRequestBulkRequestForAncEndpoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncEndpointBulkRequestForAncEndpoint {
	request := isegosdk.RequestAncEndpointBulkRequestForAncEndpoint{}
	request.ErsAncEndpointBulkRequest = expandRequestAncEndpointBulkRequestBulkRequestForAncEndpointErsAncEndpointBulkRequest(ctx, key, d)
	return &request
}

func expandRequestAncEndpointBulkRequestBulkRequestForAncEndpointErsAncEndpointBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncEndpointBulkRequestForAncEndpointErsAncEndpointBulkRequest {
	request := isegosdk.RequestAncEndpointBulkRequestForAncEndpointErsAncEndpointBulkRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation_type")))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_media_type")))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
