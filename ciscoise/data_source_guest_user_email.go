package ciscoise

import (
	"context"

	"fmt"
	"reflect"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceGuestUserEmail() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on GuestUser.

- This data source action allows the client to update a guest user email by ID.
`,

		ReadContext: dataSourceGuestUserEmailRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"portal_id": &schema.Schema{
				Description: `portalId path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
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

func dataSourceGuestUserEmailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vPortalID := d.Get("portal_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateGuestUserEmail")
		vvID := vID.(string)
		vvPortalID := vPortalID.(string)
		request1 := expandRequestGuestUserEmailUpdateGuestUserEmail(ctx, "", d)

		response1, err := client.GuestUser.UpdateGuestUserEmail(vvID, vvPortalID, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGuestUserEmail", err,
				"Failure at UpdateGuestUserEmail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateGuestUserEmail response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestGuestUserEmailUpdateGuestUserEmail(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserUpdateGuestUserEmail {
	request := isegosdk.RequestGuestUserUpdateGuestUserEmail{}
	request.OperationAdditionalData = expandRequestGuestUserEmailUpdateGuestUserEmailOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestGuestUserEmailUpdateGuestUserEmailOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserUpdateGuestUserEmailOperationAdditionalData {
	request := isegosdk.RequestGuestUserUpdateGuestUserEmailOperationAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_data")))) {
		request.AdditionalData = expandRequestGuestUserEmailUpdateGuestUserEmailOperationAdditionalDataAdditionalDataArray(ctx, key+".additional_data", d)
	}
	return &request
}

func expandRequestGuestUserEmailUpdateGuestUserEmailOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestUserUpdateGuestUserEmailOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestGuestUserUpdateGuestUserEmailOperationAdditionalDataAdditionalData{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestGuestUserEmailUpdateGuestUserEmailOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestGuestUserEmailUpdateGuestUserEmailOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserUpdateGuestUserEmailOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestGuestUserUpdateGuestUserEmailOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
