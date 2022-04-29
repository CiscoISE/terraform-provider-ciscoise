package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceAdministrationLocalExceptionRulesResetHitcount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Administration - Authorization Exception Rules.
- Device Admin Reset HitCount for local exceptions
`,

		CreateContext: resourceDeviceAdministrationLocalExceptionRulesResetHitcountCreate,
		ReadContext:   resourceDeviceAdministrationLocalExceptionRulesResetHitcountRead,
		DeleteContext: resourceDeviceAdministrationLocalExceptionRulesResetHitcountDelete,

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

func resourceDeviceAdministrationLocalExceptionRulesResetHitcountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ResetHitCountsDeviceAdminLocalExceptions create")
	log.Printf("[DEBUG] Missing ResetHitCountsDeviceAdminLocalExceptions create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vPolicyID := resourceItem["policy_id"]
	vvPolicyID := vPolicyID.(string)
	response1, restyResp1, err := client.DeviceAdministrationAuthorizationExceptionRules.ResetHitCountsDeviceAdminLocalExceptions(vvPolicyID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ResetHitCountsDeviceAdminLocalExceptions", err,
			"Failure at ResetHitCountsDeviceAdminLocalExceptions, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptionsItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ResetHitCountsDeviceAdminLocalExceptions response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceDeviceAdministrationLocalExceptionRulesResetHitcountRead(ctx, d, m)
}

func resourceDeviceAdministrationLocalExceptionRulesResetHitcountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceDeviceAdministrationLocalExceptionRulesResetHitcountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeviceAdministrationLocalExceptionRulesResetHitcount delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing DeviceAdministrationLocalExceptionRulesResetHitcount delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptionsItem(item *isegosdk.ResponseDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptions) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
