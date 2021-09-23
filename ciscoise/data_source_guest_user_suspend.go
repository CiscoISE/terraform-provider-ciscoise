package ciscoise

import (
	"context"

	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceGuestUserSuspend() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on GuestUser.

- This data source action allows the client to suspend a guest user by name.

- This data source action allows the client to suspend a guest user by ID.
`,

		ReadContext: dataSourceGuestUserSuspendRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
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

func dataSourceGuestUserSuspendRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SuspendGuestUserByName")
		vvName := vName.(string)

		response1, err := client.GuestUser.SuspendGuestUserByName(vvName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SuspendGuestUserByName", err,
				"Failure at SuspendGuestUserByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SuspendGuestUserByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: SuspendGuestUserByID")
		vvID := vID.(string)
		request2 := expandRequestGuestUserSuspendSuspendGuestUserByID(ctx, "", d)

		response2, err := client.GuestUser.SuspendGuestUserByID(vvID, request2)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SuspendGuestUserByID", err,
				"Failure at SuspendGuestUserByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		if err := d.Set("item", response2.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SuspendGuestUserByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestGuestUserSuspendSuspendGuestUserByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserSuspendGuestUserByID {
	request := isegosdk.RequestGuestUserSuspendGuestUserByID{}
	request.OperationAdditionalData = expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalData {
	request := isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalData{}
	if v, ok := d.GetOkExists(key + ".additional_data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".additional_data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".additional_data"))) {
		request.AdditionalData = expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalDataAdditionalDataArray(ctx, key+".additional_data", d)
	}
	return &request
}

func expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalDataAdditionalData{}
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	return &request
}

func expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
