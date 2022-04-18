package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEgressMatrixCellSetAllStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on EgressMatrixCell.
- his API allows the client to set status of all the egress matrix cells.
`,

		CreateContext: resourceEgressMatrixCellSetAllStatusCreate,
		ReadContext:   resourceEgressMatrixCellSetAllStatusRead,
		DeleteContext: resourceEgressMatrixCellSetAllStatusDelete,

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
						"status": &schema.Schema{
							Description: `status path parameter.`,
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

func resourceEgressMatrixCellSetAllStatusCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SetAllCellsStatus create")
	log.Printf("[DEBUG] Missing SetAllCellsStatus create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vStatus := d.Get("parameters.0.status")
	vvStatus := vStatus.(string)
	response1, err := client.EgressMatrixCell.SetAllCellsStatus(vvStatus)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing SetAllCellsStatus", err,
			"Failure at SetAllCellsStatus, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting SetAllCellsStatus response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceEgressMatrixCellSetAllStatusRead(ctx, d, m)
}

func resourceEgressMatrixCellSetAllStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceEgressMatrixCellSetAllStatusDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning EgressMatrixCellSetAllStatus delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing EgressMatrixCellSetAllStatus delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
