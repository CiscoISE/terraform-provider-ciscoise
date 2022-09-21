package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSgMappingGroupDeploy() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMappingGroup.
- This resource allows the client to deploy an IP to SGT mapping group by ID.
Only one Deploy process can run at any given time
`,

		CreateContext: resourceSgMappingGroupDeployCreate,
		ReadContext:   resourceSgMappingGroupDeployRead,
		DeleteContext: resourceSgMappingGroupDeployDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Description: `id path parameter.`,
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

func resourceSgMappingGroupDeployCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeployIPToSgtMappingGroupByID create")
	log.Printf("[DEBUG] Missing DeployIPToSgtMappingGroupByID create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vID := resourceItem["id"]
	vvID := vID.(string)
	response1, err := client.IPToSgtMappingGroup.DeployIPToSgtMappingGroupByID(vvID)
	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeployIPToSgtMappingGroupByID", err, response1.String(),
				"Failure at DeployIPToSgtMappingGroupByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeployIPToSgtMappingGroupByID", err,
			"Failure at DeployIPToSgtMappingGroupByID, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DeployIPToSgtMappingGroupByID response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceSgMappingGroupDeployRead(ctx, d, m)
}

func resourceSgMappingGroupDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSgMappingGroupDeployDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMappingGroupDeploy delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing SgMappingGroupDeploy delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
