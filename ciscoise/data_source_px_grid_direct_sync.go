package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePxGridDirectSync() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on pxGrid Direct.

- This data source is used to get the status for SyncNow Status

It returns the sync status as


syncstatus


QUEUED ,means its in  QUEUED state

SUBMITTED ,means its in  Submited to fetch the data

INPROGRESS ,means its inprogress of fetching and saving in ISE

ERRORED ,means some internal error while fetching and saving in ISE further debugging logs will help

COMPLETED ,means its COMPLETED of fetching and saving in ISE

SCH_INPROGRESS ,means its inprogress for schedule time fetch and saving in ISE

SCH_SUBMITTED ,means its submitted for schedule time fetch and will start to saving in ISE

CANCELED ,means its cancelled if any of ISE Service start when its in middle of QUEUED/SUBMITTED/INPROGRESS


connectorName


`,

		ReadContext: dataSourcePxGridDirectSyncRead,
		Schema: map[string]*schema.Schema{
			"connector_name": &schema.Schema{
				Description: `connectorName path parameter. retrieve the connector syncnow status.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connector": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connector_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sync_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourcePxGridDirectSyncRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vConnectorName := d.Get("connector_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetConnectorConfigSyncNowStatus")
		vvConnectorName := vConnectorName.(string)

		response1, restyResp1, err := client.PxGridDirect.GetConnectorConfigSyncNowStatus(vvConnectorName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetConnectorConfigSyncNowStatus", err,
				"Failure at GetConnectorConfigSyncNowStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPxGridDirectGetConnectorConfigSyncNowStatusItemResponse(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConnectorConfigSyncNowStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPxGridDirectGetConnectorConfigSyncNowStatusItemResponse(item *isegosdk.ResponsePxGridDirectGetConnectorConfigSyncNowStatusResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connector"] = flattenPxGridDirectGetConnectorConfigSyncNowStatusItemResponseConnector(item.Connector)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigSyncNowStatusItemResponseConnector(item *isegosdk.ResponsePxGridDirectGetConnectorConfigSyncNowStatusResponseConnector) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connector_name"] = item.ConnectorName
	respItem["sync_status"] = item.SyncStatus

	return []map[string]interface{}{
		respItem,
	}

}
