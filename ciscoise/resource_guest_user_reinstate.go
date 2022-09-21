package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGuestUserReinstate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on GuestUser.
- This resource allows the client to reinstate a guest user by name.
- This resource allows the client to reinstate a guest user by ID.
`,

		CreateContext: resourceGuestUserReinstateCreate,
		ReadContext:   resourceGuestUserReinstateRead,
		DeleteContext: resourceGuestUserReinstateDelete,

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
							Optional:    true,
							ForceNew:    true,
						},
						"name": &schema.Schema{
							Description: `name path parameter.`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func resourceGuestUserReinstateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ReinstateGuestUserByName create")
	log.Printf("[DEBUG] Missing ReinstateGuestUserByName create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vName, okName := d.GetOk("parameters.0.name")
	vID, okID := d.GetOk("parameters.0.id")

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReinstateGuestUserByName")
		vvName := vName.(string)

		response1, err := client.GuestUser.ReinstateGuestUserByName(vvName)

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing ReinstateGuestUserByName", err, response1.String(),
					"Failure at ReinstateGuestUserByName, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReinstateGuestUserByName", err,
				"Failure at ReinstateGuestUserByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReinstateGuestUserByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		_ = d.Set("last_updated", getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: ReinstateGuestUserByID")
		vvID := vID.(string)

		response2, err := client.GuestUser.ReinstateGuestUserByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReinstateGuestUserByID", err,
				"Failure at ReinstateGuestUserByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response2.String())

		if err := d.Set("item", response2.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReinstateGuestUserByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		_ = d.Set("last_updated", getUnixTimeString())
		return diags

	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceGuestUserReinstateRead(ctx, d, m)
}

func resourceGuestUserReinstateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceGuestUserReinstateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestUserReinstate delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing GuestUserReinstate delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
