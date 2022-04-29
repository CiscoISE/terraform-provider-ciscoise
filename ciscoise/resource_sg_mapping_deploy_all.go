package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSgMappingDeployAll() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMapping.
- This resource allows the client to deploy all the IP to SGT mappings.
Only one Deploy process can run at any given time
`,

		CreateContext: resourceSgMappingDeployAllCreate,
		ReadContext:   resourceSgMappingDeployAllRead,
		DeleteContext: resourceSgMappingDeployAllDelete,

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
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}
}

func resourceSgMappingDeployAllCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeployAllIPToSgtMapping create")
	log.Printf("[DEBUG] Missing DeployAllIPToSgtMapping create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)
	d.Set("parameters", nil)
	var diags diag.Diagnostics
	response1, err := client.IPToSgtMapping.DeployAllIPToSgtMapping()
	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeployAllIPToSgtMapping", err, response1.String(),
				"Failure at DeployAllIPToSgtMapping, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeployAllIPToSgtMapping", err,
			"Failure at DeployAllIPToSgtMapping, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %s", response1.String())
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DeployAllIPToSgtMapping response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceSgMappingDeployAllRead(ctx, d, m)
}

func resourceSgMappingDeployAllRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSgMappingDeployAllDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMappingDeployAll delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing SgMappingDeployAll delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
