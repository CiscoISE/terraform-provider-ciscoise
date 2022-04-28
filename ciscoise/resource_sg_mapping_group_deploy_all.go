package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSgMappingGroupDeployAll() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMappingGroup.
- This data source action allows the client to deploy all the IP to SGT mapping groups.
Only one Deploy process can run at any given time
`,

		CreateContext: resourceSgMappingGroupDeployAllCreate,
		ReadContext:   resourceSgMappingGroupDeployAllRead,
		DeleteContext: resourceSgMappingGroupDeployAllDelete,

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

func resourceSgMappingGroupDeployAllCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeployAllIPToSgtMappingGroup create")
	log.Printf("[DEBUG] Missing DeployAllIPToSgtMappingGroup create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)
	d.Set("parameters", nil)
	var diags diag.Diagnostics
	response1, err := client.IPToSgtMappingGroup.DeployAllIPToSgtMappingGroup()
	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeployAllIPToSgtMappingGroup", err, response1.String(),
				"Failure at DeployAllIPToSgtMappingGroup, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeployAllIPToSgtMappingGroup", err,
			"Failure at DeployAllIPToSgtMappingGroup, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DeployAllIPToSgtMappingGroup response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceSgMappingGroupDeployAllRead(ctx, d, m)
}

func resourceSgMappingGroupDeployAllRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSgMappingGroupDeployAllDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMappingGroupDeployAll delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing SgMappingGroupDeployAll delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
