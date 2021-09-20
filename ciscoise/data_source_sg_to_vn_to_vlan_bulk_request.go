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
func dataSourceSgToVnToVLANBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on SecurityGroupToVirtualNetwork.

- This data source action allows the client to submit the bulk request.
`,

		ReadContext: dataSourceSgToVnToVLANBulkRequestRead,
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

func dataSourceSgToVnToVLANBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BulkRequestForSecurityGroupsToVnToVLAN")
		request1 := expandRequestSgToVnToVLANBulkRequestBulkRequestForSecurityGroupsToVnToVLAN(ctx, "", d)

		response1, err := client.SecurityGroupToVirtualNetwork.BulkRequestForSecurityGroupsToVnToVLAN(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkRequestForSecurityGroupsToVnToVLAN", err,
				"Failure at BulkRequestForSecurityGroupsToVnToVLAN, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkRequestForSecurityGroupsToVnToVLAN response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSgToVnToVLANBulkRequestBulkRequestForSecurityGroupsToVnToVLAN(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupToVirtualNetworkBulkRequestForSecurityGroupsToVnToVLAN {
	request := isegosdk.RequestSecurityGroupToVirtualNetworkBulkRequestForSecurityGroupsToVnToVLAN{}
	request.SgtVnVLANContainerBulkRequest = expandRequestSgToVnToVLANBulkRequestBulkRequestForSecurityGroupsToVnToVLANSgtVnVLANContainerBulkRequest(ctx, key, d)
	return &request
}

func expandRequestSgToVnToVLANBulkRequestBulkRequestForSecurityGroupsToVnToVLANSgtVnVLANContainerBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupToVirtualNetworkBulkRequestForSecurityGroupsToVnToVLANSgtVnVLANContainerBulkRequest {
	request := isegosdk.RequestSecurityGroupToVirtualNetworkBulkRequestForSecurityGroupsToVnToVLANSgtVnVLANContainerBulkRequest{}
	if v, ok := d.GetOkExists("operation_type"); !isEmptyValue(reflect.ValueOf(d.Get("operation_type"))) && (ok || !reflect.DeepEqual(v, d.Get("operation_type"))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("resource_media_type"); !isEmptyValue(reflect.ValueOf(d.Get("resource_media_type"))) && (ok || !reflect.DeepEqual(v, d.Get("resource_media_type"))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
