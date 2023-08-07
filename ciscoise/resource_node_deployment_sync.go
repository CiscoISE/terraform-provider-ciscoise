package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNodeDeploymentSync() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Node Deployment.
- Performing a manual synchronization involves a reload of the target node, but not the primary PAN.
Approximate execution time 300 seconds.
`,

		CreateContext: resourceNodeDeploymentSyncCreate,
		ReadContext:   resourceNodeDeploymentSyncRead,
		DeleteContext: resourceNodeDeploymentSyncDelete,

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

						"id": &schema.Schema{
							Description: `ID which can be used to track the status of the deployment task using the Task Service API.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
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
						"hostname": &schema.Schema{
							Description: `hostname path parameter. Hostname of the node.`,
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

func resourceNodeDeploymentSyncCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SyncNode create")
	log.Printf("[DEBUG] Missing SyncNode create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vHostname := resourceItem["hostname"]
	vvHostname := vHostname.(string)
	response1, restyResp1, err := client.NodeDeployment.SyncNode(vvHostname)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing SyncNode", err, restyResp1.String(),
				"Failure at SyncNode, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing SyncNode", err,
			"Failure at SyncNode, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenNodeDeploymentSyncNodeItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting SyncNode response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceNodeDeploymentSyncRead(ctx, d, m)
}

func resourceNodeDeploymentSyncRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceNodeDeploymentSyncDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeDeploymenySync delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing NodeDeploymenySync delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenNodeDeploymentSyncNodeItem(item *isegosdk.ResponseNodeDeploymentSyncNodeResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
