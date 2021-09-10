package ciscoise

import (
	"context"

	"fmt"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceGuestUserChangeSponsorPassword() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGuestUserChangeSponsorPasswordRead,
		Schema: map[string]*schema.Schema{
			"portal_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"additional_data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGuestUserChangeSponsorPasswordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPortalID := d.Get("portal_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ChangeSponsorPassword")
		vvPortalID := vPortalID.(string)
		request1 := expandRequestGuestUserChangeSponsorPasswordChangeSponsorPassword(ctx, "", d)

		response1, err := client.GuestUser.ChangeSponsorPassword(vvPortalID, request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ChangeSponsorPassword", err,
				"Failure at ChangeSponsorPassword, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ChangeSponsorPassword response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestGuestUserChangeSponsorPasswordChangeSponsorPassword(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserChangeSponsorPassword {
	request := isegosdk.RequestGuestUserChangeSponsorPassword{}
	request.OperationAdditionalData = expandRequestGuestUserChangeSponsorPasswordChangeSponsorPasswordOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestGuestUserChangeSponsorPasswordChangeSponsorPasswordOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalData {
	request := isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalData{}
	if v, ok := d.GetOkExists("additional_data"); !isEmptyValue(reflect.ValueOf(d.Get("additional_data"))) && (ok || !reflect.DeepEqual(v, d.Get("additional_data"))) {
		request.AdditionalData = expandRequestGuestUserChangeSponsorPasswordChangeSponsorPasswordOperationAdditionalDataAdditionalDataArray(ctx, key, d)
	}
	return &request
}

func expandRequestGuestUserChangeSponsorPasswordChangeSponsorPasswordOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestGuestUserChangeSponsorPasswordChangeSponsorPasswordOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	return &request
}

func expandRequestGuestUserChangeSponsorPasswordChangeSponsorPasswordOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists("value"); !isEmptyValue(reflect.ValueOf(d.Get("value"))) && (ok || !reflect.DeepEqual(v, d.Get("value"))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(d.Get("name"))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	return &request
}