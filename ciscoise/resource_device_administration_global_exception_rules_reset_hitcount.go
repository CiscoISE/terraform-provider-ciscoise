package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceAdministrationGlobalExceptionRulesResetHitcount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Administration - Authorization Global Exception Rules.
- Device Admin Reset HitCount for Global Exceptions
`,

		CreateContext: resourceDeviceAdministrationGlobalExceptionRulesResetHitcountCreate,
		ReadContext:   resourceDeviceAdministrationGlobalExceptionRulesResetHitcountRead,
		DeleteContext: resourceDeviceAdministrationGlobalExceptionRulesResetHitcountDelete,

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
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}
}

func resourceDeviceAdministrationGlobalExceptionRulesResetHitcountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ResetHitCountsDeviceAdminGlobalExceptions create")
	log.Printf("[DEBUG] Missing ResetHitCountsDeviceAdminGlobalExceptions create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	d.Set("parameters", nil)
	response1, restyResp1, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.ResetHitCountsDeviceAdminGlobalExceptions()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ResetHitCountsDeviceAdminGlobalExceptions", err,
			"Failure at ResetHitCountsDeviceAdminGlobalExceptions, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenDeviceAdministrationAuthorizationGlobalExceptionRulesResetHitCountsDeviceAdminGlobalExceptionsItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ResetHitCountsDeviceAdminGlobalExceptions response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceDeviceAdministrationGlobalExceptionRulesResetHitcountRead(ctx, d, m)
}

func resourceDeviceAdministrationGlobalExceptionRulesResetHitcountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceDeviceAdministrationGlobalExceptionRulesResetHitcountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeviceAdministrationGlobalExceptionRulesResetHitcount delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing DeviceAdministrationGlobalExceptionRulesResetHitcount delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesResetHitCountsDeviceAdminGlobalExceptionsItem(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesResetHitCountsDeviceAdminGlobalExceptions) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
