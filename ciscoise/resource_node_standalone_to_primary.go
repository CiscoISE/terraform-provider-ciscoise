package ciscoise

import (
	"context"
	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNodeStandaloneToPrimary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Node Deployment.
- This resource promotes the standalone node on which the API is invoked to the primary Policy Administration
node (PAN).
`,

		CreateContext: resourceNodeStandaloneToPrimaryCreate,
		ReadContext:   resourceNodeStandaloneToPrimaryRead,
		DeleteContext: resourceNodeStandaloneToPrimaryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(HOTPATCH_INSTALL_TIMEOUT_SLEEP),
			Delete: schema.DefaultTimeout(HOTPATCH_ROLLBACK_TIMEOUT_SLEEP),
		},

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

func resourceNodeStandaloneToPrimaryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	log.Printf("[DEBUG] Selected method: MakePrimary")

	response1, restyResp1, err := client.NodeDeployment.MakePrimary()

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing MakePrimary", err,
			"Failure at MakePrimary, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenNodeDeploymentMakePrimaryItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting MakePrimary response",
			err))
		return diags
	}

	d.Set("last_updated", getUnixTimeString())
	d.SetId(getUnixTimeString())
	return diags
}

func resourceNodeStandaloneToPrimaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNodeStandaloneToPrimaryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNodeStandaloneToPrimaryRead(ctx, d, m)
}

func resourceNodeStandaloneToPrimaryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func flattenNodeDeploymentMakePrimaryItem(item *isegosdk.ResponseNodeDeploymentMakePrimary) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = flattenNodeDeploymentMakePrimaryItemSuccess(item.Success)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeDeploymentMakePrimaryItemSuccess(item *isegosdk.ResponseNodeDeploymentMakePrimarySuccess) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message

	return []map[string]interface{}{
		respItem,
	}

}
