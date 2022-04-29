package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceAdministrationAuthenticationResetHitcount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Administration - Authentication Rules.
- Device Admin Reset HitCount for Authentication Rules
`,

		CreateContext: resourceDeviceAdministrationAuthenticationResetHitcountCreate,
		ReadContext:   resourceDeviceAdministrationAuthenticationResetHitcountRead,
		DeleteContext: resourceDeviceAdministrationAuthenticationResetHitcountDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_id": &schema.Schema{
							Description: `policyId path parameter. Policy id`,
							Type:        schema.TypeString,
							ForceNew:    true,
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func resourceDeviceAdministrationAuthenticationResetHitcountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ResetHitCountsDeviceAdminAuthenticationRules create")
	log.Printf("[DEBUG] Missing ResetHitCountsDeviceAdminAuthenticationRules create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vPolicyID := resourceItem["policy_id"]
	vvPolicyID := vPolicyID.(string)
	response1, restyResp1, err := client.DeviceAdministrationAuthenticationRules.ResetHitCountsDeviceAdminAuthenticationRules(vvPolicyID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ResetHitCountsDeviceAdminAuthenticationRules", err,
			"Failure at ResetHitCountsDeviceAdminAuthenticationRules, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenDeviceAdministrationAuthenticationRulesResetHitCountsDeviceAdminAuthenticationRulesItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ResetHitCountsDeviceAdminAuthenticationRules response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceDeviceAdministrationAuthenticationResetHitcountRead(ctx, d, m)
}

func resourceDeviceAdministrationAuthenticationResetHitcountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceDeviceAdministrationAuthenticationResetHitcountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeviceAdministrationAuthenticationResetHitcount delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing DeviceAdministrationAuthenticationResetHitcount delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenDeviceAdministrationAuthenticationRulesResetHitCountsDeviceAdminAuthenticationRulesItem(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesResetHitCountsDeviceAdminAuthenticationRules) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
