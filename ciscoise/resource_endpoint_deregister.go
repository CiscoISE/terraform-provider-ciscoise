package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEndpointDeregister() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on endpoint.
		- This resource allows the client to de-register an endpoint.
		`,

		CreateContext: resourceEndpointDeregisterCreate,
		ReadContext:   resourceEndpointDeregisterRead,
		DeleteContext: resourceEndpointDeregisterDelete,

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

func resourceEndpointDeregisterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeregisterEndpoint create")
	log.Printf("[DEBUG] Missing DeregisterEndpoint create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vID := d.Get("parameters.0.id")
	vvID := vID.(string)
	response1, err := client.Endpoint.DeregisterEndpoint(vvID)

	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeregisterEndpoint", err,
			"Failure at DeregisterEndpoint, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DeregisterEndpoint response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceEndpointDeregisterRead(ctx, d, m)
}

func resourceEndpointDeregisterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceEndpointDeregisterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning EndpointDeregister delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing EndpointDeregister delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
