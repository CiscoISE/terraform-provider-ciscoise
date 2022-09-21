package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEgressMatrixCellClearAll() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on EgressMatrixCell.
- This resource allows the client to clear all the egress matrix cells.
`,

		CreateContext: resourceEgressMatrixCellClearAllCreate,
		ReadContext:   resourceEgressMatrixCellClearAllRead,
		DeleteContext: resourceEgressMatrixCellClearAllDelete,

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

func resourceEgressMatrixCellClearAllCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ClearAllMatrixCells create")
	log.Printf("[DEBUG] Missing ClearAllMatrixCells create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	d.Set("parameters", nil)

	var diags diag.Diagnostics

	response1, err := client.EgressMatrixCell.ClearAllMatrixCells()
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ClearAllMatrixCells", err,
			"Failure at ClearAllMatrixCells, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ClearAllMatrixCells response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceEgressMatrixCellClearAllRead(ctx, d, m)
}

func resourceEgressMatrixCellClearAllRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceEgressMatrixCellClearAllDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning EgressMatrixCellClearAll delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing EgressMatrixCellClearAll delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
