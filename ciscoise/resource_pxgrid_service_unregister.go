package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePxgridServiceUnregister() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Provider.
- 🚧 ServiceUnregister
`,

		CreateContext: resourcePxgridServiceUnregisterCreate,
		ReadContext:   resourcePxgridServiceUnregisterRead,
		DeleteContext: resourcePxgridServiceUnregisterDelete,

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

func resourcePxgridServiceUnregisterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning UnregisterService create")
	log.Printf("[DEBUG] Missing UnregisterService create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	d.Set("parameters", nil)
	var diags diag.Diagnostics
	response1, err := client.Provider.UnregisterService()
	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing UnregisterService", err, response1.String(),
				"Failure at UnregisterService, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing UnregisterService", err,
			"Failure at UnregisterService, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UnregisterService response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourcePxgridServiceUnregisterRead(ctx, d, m)
}

func resourcePxgridServiceUnregisterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourcePxgridServiceUnregisterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PxgridServiceUnregister delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing PxgridServiceUnregister delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
