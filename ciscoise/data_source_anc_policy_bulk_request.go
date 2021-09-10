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
func dataSourceAncPolicyBulkRequest() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAncPolicyBulkRequestRead,
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

func dataSourceAncPolicyBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BulkRequestForAncPolicy")
		request1 := expandRequestAncPolicyBulkRequestBulkRequestForAncPolicy(ctx, "", d)

		response1, err := client.AncPolicy.BulkRequestForAncPolicy(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkRequestForAncPolicy", err,
				"Failure at BulkRequestForAncPolicy, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkRequestForAncPolicy response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestAncPolicyBulkRequestBulkRequestForAncPolicy(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncPolicyBulkRequestForAncPolicy {
	request := isegosdk.RequestAncPolicyBulkRequestForAncPolicy{}
	request.ErsAncPolicyBulkRequest = expandRequestAncPolicyBulkRequestBulkRequestForAncPolicyErsAncPolicyBulkRequest(ctx, key, d)
	return &request
}

func expandRequestAncPolicyBulkRequestBulkRequestForAncPolicyErsAncPolicyBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncPolicyBulkRequestForAncPolicyErsAncPolicyBulkRequest {
	request := isegosdk.RequestAncPolicyBulkRequestForAncPolicyErsAncPolicyBulkRequest{}
	if v, ok := d.GetOkExists("operation_type"); !isEmptyValue(reflect.ValueOf(d.Get("operation_type"))) && (ok || !reflect.DeepEqual(v, d.Get("operation_type"))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("resource_media_type"); !isEmptyValue(reflect.ValueOf(d.Get("resource_media_type"))) && (ok || !reflect.DeepEqual(v, d.Get("resource_media_type"))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}