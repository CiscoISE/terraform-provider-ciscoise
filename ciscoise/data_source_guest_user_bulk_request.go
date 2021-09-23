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
func dataSourceGuestUserBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on GuestUser.

- This data source action allows the client to submit the bulk request.
`,

		ReadContext: dataSourceGuestUserBulkRequestRead,
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

func dataSourceGuestUserBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BulkRequestForGuestUser")
		request1 := expandRequestGuestUserBulkRequestBulkRequestForGuestUser(ctx, "", d)

		response1, err := client.GuestUser.BulkRequestForGuestUser(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BulkRequestForGuestUser", err,
				"Failure at BulkRequestForGuestUser, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BulkRequestForGuestUser response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestGuestUserBulkRequestBulkRequestForGuestUser(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserBulkRequestForGuestUser {
	request := isegosdk.RequestGuestUserBulkRequestForGuestUser{}
	request.GuestUserBulkRequest = expandRequestGuestUserBulkRequestBulkRequestForGuestUserGuestUserBulkRequest(ctx, key, d)
	return &request
}

func expandRequestGuestUserBulkRequestBulkRequestForGuestUserGuestUserBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserBulkRequestForGuestUserGuestUserBulkRequest {
	request := isegosdk.RequestGuestUserBulkRequestForGuestUserGuestUserBulkRequest{}
	if v, ok := d.GetOkExists(key + ".operation_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".operation_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".operation_type"))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".resource_media_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".resource_media_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".resource_media_type"))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
