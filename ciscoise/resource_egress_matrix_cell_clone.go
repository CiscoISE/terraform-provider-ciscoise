package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEgressMatrixCellClone() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on EgressMatrixCell.
- This resource allows the client to clone an egress matrix cell.
`,

		CreateContext: resourceEgressMatrixCellCloneCreate,
		ReadContext:   resourceEgressMatrixCellCloneRead,
		DeleteContext: resourceEgressMatrixCellCloneDelete,

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

						"result_value": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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
						"dst_sgt_id": &schema.Schema{
							Description: `dstSgtId path parameter.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"id": &schema.Schema{
							Description: `id path parameter.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"src_sgt_id": &schema.Schema{
							Description: `srcSgtId path parameter.`,
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

func resourceEgressMatrixCellCloneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning CloneMatrixCell create")
	log.Printf("[DEBUG] Missing CloneMatrixCell create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	vID := d.Get("parameters.0.id")
	vSrcSgtID := d.Get("parameters.0.src_sgt_id")
	vDstSgtID := d.Get("parameters.0.dst_sgt_id")

	var diags diag.Diagnostics

	vvID := vID.(string)
	vvSrcSgtID := vSrcSgtID.(string)
	vvDstSgtID := vDstSgtID.(string)

	response1, restyResp1, err := client.EgressMatrixCell.CloneMatrixCell(vvID, vvSrcSgtID, vvDstSgtID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing CloneMatrixCell", err, restyResp1.String(),
				"Failure at CloneMatrixCell, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CloneMatrixCell", err,
			"Failure at CloneMatrixCell, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenEgressMatrixCellCloneMatrixCellItem(response1.OperationResult)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CloneMatrixCell response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceEgressMatrixCellCloneRead(ctx, d, m)
}

func resourceEgressMatrixCellCloneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceEgressMatrixCellCloneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning EgressMatrixCellClone delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing EgressMatrixCellClone delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenEgressMatrixCellCloneMatrixCellItem(item *isegosdk.ResponseEgressMatrixCellCloneMatrixCellOperationResult) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["result_value"] = flattenEgressMatrixCellCloneMatrixCellItemResultValue(item.ResultValue)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEgressMatrixCellCloneMatrixCellItemResultValue(items *[]isegosdk.ResponseEgressMatrixCellCloneMatrixCellOperationResultResultValue) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["value"] = item.Value
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
