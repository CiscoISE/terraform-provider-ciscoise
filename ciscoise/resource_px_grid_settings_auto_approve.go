package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePxGridSettingsAutoApprove() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on PxGridSettings.
- This data source action allows the client to auto approve the pxGrid settings.
`,

		CreateContext: resourcePxGridSettingsAutoApproveCreate,
		ReadContext:   resourcePxGridSettingsAutoApproveRead,
		DeleteContext: resourcePxGridSettingsAutoApproveDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_password_based_accounts": &schema.Schema{
							Description:  `Allow password based accounts when true`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							ForceNew:     true,
							Optional:     true,
						},
						"auto_approve_cert_based_accounts": &schema.Schema{
							Description:  `Auto approve certificate based accounts when true`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							ForceNew:     true,
							Optional:     true,
						},
					},
				},
			},
		},
	}
}

func resourcePxGridSettingsAutoApproveCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AutoapprovePxGridSettings create")
	log.Printf("[DEBUG] Missing AutoapprovePxGridSettings create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	request1 := expandRequestPxGridSettingsAutoApproveAutoapprovePxGridSettings(ctx, "parameters.0", d)
	response1, err := client.PxGridSettings.AutoapprovePxGridSettings(request1)
	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing AutoapprovePxGridSettings", err, response1.String(),
				"Failure at AutoapprovePxGridSettings, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AutoapprovePxGridSettings", err,
			"Failure at AutoapprovePxGridSettings, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %s", response1.String())
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AutoapprovePxGridSettings response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourcePxGridSettingsAutoApproveRead(ctx, d, m)
}

func resourcePxGridSettingsAutoApproveRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourcePxGridSettingsAutoApproveDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PxGridSettingsAutoApprove delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing PxGridSettingsAutoApprove delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestPxGridSettingsAutoApproveAutoapprovePxGridSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridSettingsAutoapprovePxGridSettings {
	request := isegosdk.RequestPxGridSettingsAutoapprovePxGridSettings{}
	request.PxgridSettings = expandRequestPxGridSettingsAutoApproveAutoapprovePxGridSettingsPxgridSettings(ctx, key, d)
	return &request
}

func expandRequestPxGridSettingsAutoApproveAutoapprovePxGridSettingsPxgridSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridSettingsAutoapprovePxGridSettingsPxgridSettings {
	request := isegosdk.RequestPxGridSettingsAutoapprovePxGridSettingsPxgridSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_approve_cert_based_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_approve_cert_based_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_approve_cert_based_accounts")))) {
		request.AutoApproveCertBasedAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_password_based_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_password_based_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_password_based_accounts")))) {
		request.AllowPasswordBasedAccounts = interfaceToBoolPtr(v)
	}
	return &request
}
