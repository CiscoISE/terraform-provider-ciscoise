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
func dataSourceGuestTypeSms() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGuestTypeSmsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
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

func dataSourceGuestTypeSmsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateGuestTypeSms")
		vvID := vID.(string)
		request1 := expandRequestGuestTypeSmsUpdateGuestTypeSms(ctx, "", d)

		response1, err := client.GuestType.UpdateGuestTypeSms(vvID, request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGuestTypeSms", err,
				"Failure at UpdateGuestTypeSms, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateGuestTypeSms response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestGuestTypeSmsUpdateGuestTypeSms(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeSms {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeSms{}
	request.OperationAdditionalData = expandRequestGuestTypeSmsUpdateGuestTypeSmsOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestGuestTypeSmsUpdateGuestTypeSmsOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeSmsOperationAdditionalData {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeSmsOperationAdditionalData{}
	if v, ok := d.GetOkExists("additional_data"); !isEmptyValue(reflect.ValueOf(d.Get("additional_data"))) && (ok || !reflect.DeepEqual(v, d.Get("additional_data"))) {
		request.AdditionalData = expandRequestGuestTypeSmsUpdateGuestTypeSmsOperationAdditionalDataAdditionalDataArray(ctx, key, d)
	}
	return &request
}

func expandRequestGuestTypeSmsUpdateGuestTypeSmsOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestTypeUpdateGuestTypeSmsOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestGuestTypeUpdateGuestTypeSmsOperationAdditionalDataAdditionalData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestGuestTypeSmsUpdateGuestTypeSmsOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	return &request
}

func expandRequestGuestTypeSmsUpdateGuestTypeSmsOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeSmsOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeSmsOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists("value"); !isEmptyValue(reflect.ValueOf(d.Get("value"))) && (ok || !reflect.DeepEqual(v, d.Get("value"))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(d.Get("name"))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	return &request
}