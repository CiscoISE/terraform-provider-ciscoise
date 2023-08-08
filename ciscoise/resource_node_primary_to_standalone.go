package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNodePrimaryToStandalone() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Node Deployment.
- This resource changes the primary PAN in a single node cluster on which the API is invoked, to a standalone
node.
`,

		CreateContext: resourceNodePrimaryToStandaloneCreate,
		ReadContext:   resourceNodePrimaryToStandaloneRead,
		DeleteContext: resourceNodePrimaryToStandaloneDelete,

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

func resourceNodePrimaryToStandaloneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning MakeStandalone create")
	log.Printf("[DEBUG] Missing MakeStandalone create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	d.Set("parameters", nil)
	var diags diag.Diagnostics
	response1, restyResp1, err := client.NodeDeployment.MakeStandalone()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing MakeStandalone", err,
			"Failure at MakeStandalone, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenNodeDeploymentMakeStandaloneItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting MakeStandalone response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceNodePrimaryToStandaloneRead(ctx, d, m)
}

func resourceNodePrimaryToStandaloneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceNodePrimaryToStandaloneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodePrimaryToStandalone delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing NodePrimaryToStandalone delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenNodeDeploymentMakeStandaloneItem(item *isegosdk.ResponseNodeDeploymentMakeStandalone) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = flattenNodeDeploymentMakeStandaloneItemSuccess(item.Success)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeDeploymentMakeStandaloneItemSuccess(item *isegosdk.ResponseNodeDeploymentMakeStandaloneSuccess) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message

	return []map[string]interface{}{
		respItem,
	}

}
