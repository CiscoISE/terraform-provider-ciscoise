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
func dataSourceEndpointBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on endpoint.

- This data source action allows the client to submit the bulk request.
`,

		ReadContext: dataSourceEndpointBulkRequestRead,
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

func dataSourceEndpointBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BulkRequestForEndpoint")
		request1 := expandRequestEndpointBulkRequestBulkRequestForEndpoint(ctx, "", d)

		response1, err := client.Endpoint.BulkRequestForEndpoint(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkRequestForEndpoint", err,
				"Failure at BulkRequestForEndpoint, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkRequestForEndpoint response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestEndpointBulkRequestBulkRequestForEndpoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointBulkRequestForEndpoint {
	request := isegosdk.RequestEndpointBulkRequestForEndpoint{}
	request.EndpointBulkRequest = expandRequestEndpointBulkRequestBulkRequestForEndpointEndpointBulkRequest(ctx, key, d)
	return &request
}

func expandRequestEndpointBulkRequestBulkRequestForEndpointEndpointBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointBulkRequestForEndpointEndpointBulkRequest {
	request := isegosdk.RequestEndpointBulkRequestForEndpointEndpointBulkRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation_type")))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_media_type")))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
