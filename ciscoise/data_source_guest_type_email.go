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
func dataSourceGuestTypeEmail() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGuestTypeEmailRead,
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

func dataSourceGuestTypeEmailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateGuestTypeEmail")
		vvID := vID.(string)
		request1 := expandRequestGuestTypeEmailUpdateGuestTypeEmail(ctx, "", d)

		response1, err := client.GuestType.UpdateGuestTypeEmail(vvID, request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGuestTypeEmail", err,
				"Failure at UpdateGuestTypeEmail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateGuestTypeEmail response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestGuestTypeEmailUpdateGuestTypeEmail(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeEmail {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeEmail{}
	request.OperationAdditionalData = expandRequestGuestTypeEmailUpdateGuestTypeEmailOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestGuestTypeEmailUpdateGuestTypeEmailOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeEmailOperationAdditionalData {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeEmailOperationAdditionalData{}
	if v, ok := d.GetOkExists("additional_data"); !isEmptyValue(reflect.ValueOf(d.Get("additional_data"))) && (ok || !reflect.DeepEqual(v, d.Get("additional_data"))) {
		request.AdditionalData = expandRequestGuestTypeEmailUpdateGuestTypeEmailOperationAdditionalDataAdditionalDataArray(ctx, key, d)
	}
	return &request
}

func expandRequestGuestTypeEmailUpdateGuestTypeEmailOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestTypeUpdateGuestTypeEmailOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestGuestTypeUpdateGuestTypeEmailOperationAdditionalDataAdditionalData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestGuestTypeEmailUpdateGuestTypeEmailOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	return &request
}

func expandRequestGuestTypeEmailUpdateGuestTypeEmailOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeEmailOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeEmailOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists("value"); !isEmptyValue(reflect.ValueOf(d.Get("value"))) && (ok || !reflect.DeepEqual(v, d.Get("value"))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(d.Get("name"))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
