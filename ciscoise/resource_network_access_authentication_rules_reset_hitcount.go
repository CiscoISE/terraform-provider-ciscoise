package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkAccessAuthenticationRulesResetHitcount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Access - Authentication Rules.
- Network Access Reset HitCount for Authentication Rules
`,

		CreateContext: resourceNetworkAccessAuthenticationRulesResetHitcountCreate,
		ReadContext:   resourceNetworkAccessAuthenticationRulesResetHitcountRead,
		DeleteContext: resourceNetworkAccessAuthenticationRulesResetHitcountDelete,

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
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkAccessAuthenticationRulesResetHitcountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ResetHitCountsNetworkAccessAuthenticationRules create")
	log.Printf("[DEBUG] Missing ResetHitCountsNetworkAccessAuthenticationRules create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPolicyID := d.Get("parameters.0.policy_id")

	vvPolicyID := vPolicyID.(string)

	response1, restyResp1, err := client.NetworkAccessAuthenticationRules.ResetHitCountsNetworkAccessAuthenticationRules(vvPolicyID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ResetHitCountsNetworkAccessAuthenticationRules", err,
			"Failure at ResetHitCountsNetworkAccessAuthenticationRules, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenNetworkAccessAuthenticationRulesResetHitCountsNetworkAccessAuthenticationRulesItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ResetHitCountsNetworkAccessAuthenticationRules response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return resourceNetworkAccessAuthenticationRulesResetHitcountRead(ctx, d, m)
}

func resourceNetworkAccessAuthenticationRulesResetHitcountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkAccessAuthenticationRulesResetHitcountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessAuthenticationRulesResetHitcount delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing NetworkAccessAuthenticationRulesResetHitcount delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenNetworkAccessAuthenticationRulesResetHitCountsNetworkAccessAuthenticationRulesItem(item *isegosdk.ResponseNetworkAccessAuthenticationRulesResetHitCountsNetworkAccessAuthenticationRules) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
