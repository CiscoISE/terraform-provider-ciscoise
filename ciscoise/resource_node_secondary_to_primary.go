package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNodeSecondaryToPrimary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Node Deployment.
- Execute this API in the secondary PAN in the cluster to promote the node to primary PAN.  Ensure that the API Gateway
setting is enabled in the secondary PAN.
 Approximate execution time 300 seconds.
`,

		CreateContext: resourceNodeSecondaryToPrimaryCreate,
		ReadContext:   resourceNodeSecondaryToPrimaryRead,
		DeleteContext: resourceNodeSecondaryToPrimaryDelete,

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

						"success": &schema.Schema{
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
						"version": &schema.Schema{
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

func resourceNodeSecondaryToPrimaryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PromoteNode create")
	log.Printf("[DEBUG] Missing PromoteNode create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	response1, restyResp1, err := client.NodeDeployment.PromoteNode()

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing PromoteNode", err,
			"Failure at PromoteNode, unexpected response", ""))
		return diags
	}

	vItem1 := flattenNodeDeploymentPromoteNodeItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting PromoteNode response",
			err))
		return diags
	}

	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceNodeSecondaryToPrimaryRead(ctx, d, m)
}

func resourceNodeSecondaryToPrimaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceNodeSecondaryToPrimaryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeSecondaryToPrimary delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing NodeSecondaryToPrimary delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenNodeDeploymentPromoteNodeItem(item *isegosdk.ResponseNodeDeploymentPromoteNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = flattenNodeDeploymentPromoteNodeItemSuccess(item.Success)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeDeploymentPromoteNodeItemSuccess(item *isegosdk.ResponseNodeDeploymentPromoteNodeSuccess) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message

	return []map[string]interface{}{
		respItem,
	}

}
