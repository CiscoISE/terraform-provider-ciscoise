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
func dataSourceGuestUserChangeSponsorPassword() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on GuestUser.

- This data source action allows the client to change the sponsor password.
`,

		ReadContext: dataSourceGuestUserChangeSponsorPasswordRead,
		Schema: map[string]*schema.Schema{
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

func dataSourceGuestUserChangeSponsorPasswordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPortalID := d.Get("portal_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ChangeSponsorPassword")
		vvPortalID := vPortalID.(string)
		request1 := expandRequestGuestUserChangeSponsorPasswordChangeSponsorPassword(ctx, "", d)

		response1, err := client.GuestUser.ChangeSponsorPassword(vvPortalID, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing ChangeSponsorPassword", err, response1.String(),
					"Failure at ChangeSponsorPassword, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ChangeSponsorPassword", err,
				"Failure at ChangeSponsorPassword, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_data")))) {
		request.AdditionalData = expandRequestGuestUserChangeSponsorPasswordChangeSponsorPasswordOperationAdditionalDataAdditionalDataArray(ctx, key+".additional_data", d)
	}
	return &request
}

func expandRequestGuestUserChangeSponsorPasswordChangeSponsorPasswordOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestGuestUserChangeSponsorPasswordChangeSponsorPasswordOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestGuestUserChangeSponsorPasswordChangeSponsorPasswordOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
