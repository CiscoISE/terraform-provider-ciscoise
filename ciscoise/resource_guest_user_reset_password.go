package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGuestUserResetPassword() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on GuestUser.
- This data source action allows the client to reset the guest user password.
`,

		CreateContext: resourceGuestUserResetPasswordCreate,
		ReadContext:   resourceGuestUserResetPasswordRead,
		DeleteContext: resourceGuestUserResetPasswordDelete,

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

func resourceGuestUserResetPasswordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ResetGuestUserPasswordByID create")
	log.Printf("[DEBUG] Missing ResetGuestUserPasswordByID create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vID := resourceItem["id"]
	vvID := vID.(string)
	response1, restyResp1, err := client.GuestUser.ResetGuestUserPasswordByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ResetGuestUserPasswordByID", err,
			"Failure at ResetGuestUserPasswordByID, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenGuestUserResetGuestUserPasswordByIDItem(response1.OperationResult)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ResetGuestUserPasswordByID response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceGuestUserResetPasswordRead(ctx, d, m)
}

func resourceGuestUserResetPasswordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceGuestUserResetPasswordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestUserResetPassword delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing GuestUserResetPassword delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func flattenGuestUserResetGuestUserPasswordByIDItem(item *isegosdk.ResponseGuestUserResetGuestUserPasswordByIDOperationResult) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["result_value"] = flattenGuestUserResetGuestUserPasswordByIDItemResultValue(item.ResultValue)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenGuestUserResetGuestUserPasswordByIDItemResultValue(items *[]isegosdk.ResponseGuestUserResetGuestUserPasswordByIDOperationResultResultValue) []map[string]interface{} {
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
