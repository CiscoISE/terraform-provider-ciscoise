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
func dataSourcePxGridSettingsAutoApprove() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePxGridSettingsAutoApproveRead,
		Schema: map[string]*schema.Schema{
			"allow_password_based_accounts": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_approve_cert_based_accounts": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePxGridSettingsAutoApproveRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: AutoapprovePxGridSettings")
		request1 := expandRequestPxGridSettingsAutoApproveAutoapprovePxGridSettings(ctx, "", d)

		response1, err := client.PxGridSettings.AutoapprovePxGridSettings(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AutoapprovePxGridSettings", err,
				"Failure at AutoapprovePxGridSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting AutoapprovePxGridSettings response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPxGridSettingsAutoApproveAutoapprovePxGridSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridSettingsAutoapprovePxGridSettings {
	request := isegosdk.RequestPxGridSettingsAutoapprovePxGridSettings{}
	request.PxgridSettings = expandRequestPxGridSettingsAutoApproveAutoapprovePxGridSettingsPxgridSettings(ctx, key, d)
	return &request
}

func expandRequestPxGridSettingsAutoApproveAutoapprovePxGridSettingsPxgridSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridSettingsAutoapprovePxGridSettingsPxgridSettings {
	request := isegosdk.RequestPxGridSettingsAutoapprovePxGridSettingsPxgridSettings{}
	if v, ok := d.GetOkExists("auto_approve_cert_based_accounts"); !isEmptyValue(reflect.ValueOf(d.Get("auto_approve_cert_based_accounts"))) && (ok || !reflect.DeepEqual(v, d.Get("auto_approve_cert_based_accounts"))) {
		request.AutoApproveCertBasedAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("allow_password_based_accounts"); !isEmptyValue(reflect.ValueOf(d.Get("allow_password_based_accounts"))) && (ok || !reflect.DeepEqual(v, d.Get("allow_password_based_accounts"))) {
		request.AllowPasswordBasedAccounts = interfaceToBoolPtr(v)
	}
	return &request
}