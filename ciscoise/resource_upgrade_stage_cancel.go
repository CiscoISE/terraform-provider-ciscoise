package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceUpgradestageCancel() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on full upgrade.

- API to cancel staging process of specified nodes from PPAN
`,

		CreateContext: resourceUpgradestageCancelCreate,
		ReadContext:   resourceUpgradestageCancelRead,
		DeleteContext: resourceUpgradestageCancelDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pre_check_report_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
						"x_request_id": &schema.Schema{
							Description: `X-Request-ID header parameter. request Id, will return in the response headers, and appear in logs`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"hostnames": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"pre_check_report_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"upgrade_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceUpgradestageCancelCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestUpgradestageCancelCancelStagingOnPPAN(ctx, "parameters.0", d)

	response1, restyResp1, err := client.FullUpgrade.CancelStagingOnPPAN(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenFullUpgradeCancelStagingOnPPANItemResponse(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CancelStagingOnPPAN response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}

func expandRequestUpgradestageCancelCancelStagingOnPPAN(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestFullUpgradeCancelStagingOnPPAN {
	request := isegosdk.RequestFullUpgradeCancelStagingOnPPAN{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostnames")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostnames")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostnames")))) {
		request.Hostnames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pre_check_report_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pre_check_report_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pre_check_report_id")))) {
		request.PreCheckReportID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upgrade_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upgrade_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upgrade_type")))) {
		request.UpgradeType = interfaceToString(v)
	}
	return &request
}

func flattenFullUpgradeCancelStagingOnPPANItemResponse(item *isegosdk.ResponseFullUpgradeCancelStagingOnPPANResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	respItem["pre_check_report_id"] = item.PreCheckReportID

	return []map[string]interface{}{
		respItem,
	}

}

func resourceUpgradestageCancelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceUpgradestageCancelUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceUpgradestageCancelRead(ctx, d, m)
}

func resourceUpgradestageCancelDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
