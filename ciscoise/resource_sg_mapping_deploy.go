package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSgMappingDeploy() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMapping.
- This resource allows the client to deploy an IP to SGT mapping by ID.
Only one Deploy process can run at any given time
`,

		CreateContext: resourceSgMappingDeployCreate,
		ReadContext:   resourceSgMappingDeployRead,
		DeleteContext: resourceSgMappingDeployDelete,

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
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSgMappingDeployCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeployIPToSgtMappingByID create")
	log.Printf("[DEBUG] Missing DeployIPToSgtMappingByID create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("parameters.0.id")

	log.Printf("[DEBUG] Selected method: DeployIPToSgtMappingByID")
	vvID := vID.(string)

	response1, err := client.IPToSgtMapping.DeployIPToSgtMappingByID(vvID)

	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeployIPToSgtMappingByID", err,
			"Failure at DeployIPToSgtMappingByID, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %s", response1.String())

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DeployIPToSgtMappingByID response",
			err))
		return diags
	}

	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceSgMappingDeployRead(ctx, d, m)
}

func resourceSgMappingDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSgMappingDeployDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMappingDeploy delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing SgMappingDeploy delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
