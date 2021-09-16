package ciscoise

import (
	"context"

	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSxpLocalBindingsBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on SXPLocalBindings.

This data source action allows the client to submit the bulk request.`,

		ReadContext: dataSourceSxpLocalBindingsBulkRequestRead,
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

func dataSourceSxpLocalBindingsBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BulkRequestForSxpLocalBindings")
		request1 := expandRequestSxpLocalBindingsBulkRequestBulkRequestForSxpLocalBindings(ctx, "", d)

		response1, err := client.SxpLocalBindings.BulkRequestForSxpLocalBindings(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkRequestForSxpLocalBindings", err,
				"Failure at BulkRequestForSxpLocalBindings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkRequestForSxpLocalBindings response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSxpLocalBindingsBulkRequestBulkRequestForSxpLocalBindings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpLocalBindingsBulkRequestForSxpLocalBindings {
	request := isegosdk.RequestSxpLocalBindingsBulkRequestForSxpLocalBindings{}
	request.LocalBindingBulkRequest = expandRequestSxpLocalBindingsBulkRequestBulkRequestForSxpLocalBindingsLocalBindingBulkRequest(ctx, key, d)
	return &request
}

func expandRequestSxpLocalBindingsBulkRequestBulkRequestForSxpLocalBindingsLocalBindingBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpLocalBindingsBulkRequestForSxpLocalBindingsLocalBindingBulkRequest {
	request := isegosdk.RequestSxpLocalBindingsBulkRequestForSxpLocalBindingsLocalBindingBulkRequest{}
	if v, ok := d.GetOkExists("operation_type"); !isEmptyValue(reflect.ValueOf(d.Get("operation_type"))) && (ok || !reflect.DeepEqual(v, d.Get("operation_type"))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("resource_media_type"); !isEmptyValue(reflect.ValueOf(d.Get("resource_media_type"))) && (ok || !reflect.DeepEqual(v, d.Get("resource_media_type"))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
