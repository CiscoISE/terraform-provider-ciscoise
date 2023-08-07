package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkAccessGlobalExceptionRulesResetHitcount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Access - Authorization Exception Rules.
- Network Access Reset HitCount for local exceptions
`,

		CreateContext: resourceNetworkAccessGlobalExceptionRulesResetHitcountCreate,
		ReadContext:   resourceNetworkAccessGlobalExceptionRulesResetHitcountRead,
		DeleteContext: resourceNetworkAccessGlobalExceptionRulesResetHitcountDelete,

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

func resourceNetworkAccessGlobalExceptionRulesResetHitcountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ResetHitCountsNetworkAccessGlobalExceptions create")
	log.Printf("[DEBUG] Missing ResetHitCountsNetworkAccessGlobalExceptions create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	d.Set("parameters", nil)
	var diags diag.Diagnostics
	response1, restyResp1, err := client.NetworkAccessAuthorizationGlobalExceptionRules.ResetHitCountsNetworkAccessGlobalExceptions()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ResetHitCountsNetworkAccessGlobalExceptions", err,
			"Failure at ResetHitCountsNetworkAccessGlobalExceptions, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenNetworkAccessAuthorizationGlobalExceptionRulesResetHitCountsNetworkAccessGlobalExceptionsItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ResetHitCountsNetworkAccessGlobalExceptions response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceNetworkAccessGlobalExceptionRulesResetHitcountRead(ctx, d, m)
}

func resourceNetworkAccessGlobalExceptionRulesResetHitcountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkAccessGlobalExceptionRulesResetHitcountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessGlobalExceptionRulesResetHitcount delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing NetworkAccessGlobalExceptionRulesResetHitcount delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenNetworkAccessAuthorizationGlobalExceptionRulesResetHitCountsNetworkAccessGlobalExceptionsItem(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesResetHitCountsNetworkAccessGlobalExceptions) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
