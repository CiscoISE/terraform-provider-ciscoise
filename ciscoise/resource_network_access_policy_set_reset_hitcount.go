package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkAccessPolicySetResetHitcount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Access - Authorization Exception Rules.
- Network Access Reset HitCount for local exceptions
`,

		CreateContext: resourceNetworkAccessPolicySetResetHitcountCreate,
		ReadContext:   resourceNetworkAccessPolicySetResetHitcountRead,
		DeleteContext: resourceNetworkAccessPolicySetResetHitcountDelete,

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

func resourceNetworkAccessPolicySetResetHitcountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ResetHitCountsNetworkAccessPolicySets create")
	log.Printf("[DEBUG] Missing ResetHitCountsNetworkAccessPolicySets create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)
	d.Set("parameters", nil)
	var diags diag.Diagnostics
	response1, restyResp1, err := client.NetworkAccessPolicySet.ResetHitCountsNetworkAccessPolicySets()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ResetHitCountsNetworkAccessPolicySets", err,
			"Failure at ResetHitCountsNetworkAccessPolicySets, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenNetworkAccessPolicySetResetHitCountsNetworkAccessPolicySetsItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ResetHitCountsNetworkAccessPolicySets response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceNetworkAccessPolicySetResetHitcountRead(ctx, d, m)
}

func resourceNetworkAccessPolicySetResetHitcountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkAccessPolicySetResetHitcountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessPolicySetResetHitcount delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing NetworkAccessPolicySetResetHitcount delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenNetworkAccessPolicySetResetHitCountsNetworkAccessPolicySetsItem(item *isegosdk.ResponseNetworkAccessPolicySetResetHitCountsNetworkAccessPolicySets) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
