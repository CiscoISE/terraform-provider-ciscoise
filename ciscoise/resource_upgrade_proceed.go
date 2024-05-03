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
func resourceUpgradeProceed() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on full upgrade.

- API's purpose would be to orchestrate upgrade execution on PPAN
`,

		CreateContext: resourceUpgradeProceedCreate,
		ReadContext:   resourceUpgradeProceedRead,
		DeleteContext: resourceUpgradeProceedDelete,
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

func resourceUpgradeProceedCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestUpgradeProceedInitiateUpgradeOnPPAN(ctx, "parameters.0", d)

	response1, restyResp1, err := client.FullUpgrade.InitiateUpgradeOnPPAN(request1)

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

	vItem1 := flattenFullUpgradeInitiateUpgradeOnPPANItemResponse(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting InitiateUpgradeOnPPAN response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}

func expandRequestUpgradeProceedInitiateUpgradeOnPPAN(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestFullUpgradeInitiateUpgradeOnPPAN {
	request := isegosdk.RequestFullUpgradeInitiateUpgradeOnPPAN{}
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

func flattenFullUpgradeInitiateUpgradeOnPPANItemResponse(item *isegosdk.ResponseFullUpgradeInitiateUpgradeOnPPANResponse) []map[string]interface{} {
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

func resourceUpgradeProceedRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceUpgradeProceedUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceUpgradeProceedRead(ctx, d, m)
}

func resourceUpgradeProceedDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
