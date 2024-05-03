package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUpgradeProceed() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on full upgrade.

- get the status of upgrade stage process for the requested nodes
`,

		ReadContext: dataSourceUpgradeProceedRead,
		Schema: map[string]*schema.Schema{
			"pre_check_report_id": &schema.Schema{
				Description: `preCheckReportID query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"db_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"node": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"percentage": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"progress_msg": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceUpgradeProceedRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPreCheckReportID, okPreCheckReportID := d.GetOk("pre_check_report_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: StageStatus")
		queryParams1 := isegosdk.StageStatusQueryParams{}

		if okPreCheckReportID {
			queryParams1.PreCheckReportID = vPreCheckReportID.(string)
		}

		response1, restyResp1, err := client.FullUpgrade.StageStatus(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 StageStatus", err,
				"Failure at StageStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenFullUpgradeStageStatusItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting StageStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFullUpgradeStageStatusItems(items *[]isegosdk.ResponseFullUpgradestageStatusResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["db_status"] = item.DbStatus
		respItem["message"] = item.Message
		respItem["node"] = item.Node
		respItem["percentage"] = item.Percentage
		respItem["progress_msg"] = item.ProgressMsg
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
