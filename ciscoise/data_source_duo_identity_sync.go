package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDuoIDentitySync() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Duo-IdentitySync.

- Duo-IdentitySync Get the list of all Identitysync configurations

- Duo-IdentitySync Get the Identitysync config specified in the syncName
`,

		ReadContext: dataSourceDuoIDentitySyncRead,
		Schema: map[string]*schema.Schema{
			"sync_name": &schema.Schema{
				Description: `syncName path parameter. This name is used to update, delete or retrieve the specific Identitysync config.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"identity_sync": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ad_groups": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sid": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"source": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"configurations": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"active_directories": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"directory_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"domain": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"last_sync": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sync_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sync_schedule": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"interval": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"interval_unit": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"scheduler_sync": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
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
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ad_groups": &schema.Schema{
							Description: `Number of Active Directory Groups`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"destination": &schema.Schema{
							Description: `Destination of Identitysync (Mfa Provider)`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_sync": &schema.Schema{
							Description: `Time of the last Sync`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name of the Identitysync configuration`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"source": &schema.Schema{
							Description: `Source of the Identitysync (Active Directory)`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sync_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDuoIDentitySyncRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vSyncName, okSyncName := d.GetOk("sync_name")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okSyncName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetIDentitysync")

		response1, restyResp1, err := client.DuoIDentitySync.GetIDentitysync()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetIDentitysync", err,
				"Failure at GetIDentitysync, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDuoIDentitySyncGetIDentitysyncItemsResponse(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIDentitysync response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetIDentitysyncBySyncName")
		vvSyncName := vSyncName.(string)

		response2, restyResp2, err := client.DuoIDentitySync.GetIDentitysyncBySyncName(vvSyncName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetIDentitysyncBySyncName", err,
				"Failure at GetIDentitysyncBySyncName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponse(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIDentitysyncBySyncName response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDuoIDentitySyncGetIDentitysyncItemsResponse(items *[]isegosdk.ResponseDuoIDentitySyncGetIDentitysyncResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ad_groups"] = item.AdGroups
		respItem["destination"] = item.Destination
		respItem["last_sync"] = item.LastSync
		respItem["name"] = item.Name
		respItem["source"] = item.Source
		respItem["sync_status"] = item.SyncStatus
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponse(item *isegosdk.ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["identity_sync"] = flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponseIDentitySync(item.IDentitySync)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponseIDentitySync(item *isegosdk.ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySync) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ad_groups"] = flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponseIDentitySyncAdGroups(item.AdGroups)
	respItem["configurations"] = flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponseIDentitySyncConfigurations(item.Configurations)
	respItem["last_sync"] = item.LastSync
	respItem["sync_name"] = item.SyncName
	respItem["sync_schedule"] = flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponseIDentitySyncSyncSchedule(item.SyncSchedule)
	respItem["sync_status"] = item.SyncStatus

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponseIDentitySyncAdGroups(items *[]isegosdk.ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncAdGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["sid"] = item.Sid
		respItem["source"] = item.Source
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponseIDentitySyncConfigurations(item *isegosdk.ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncConfigurations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["active_directories"] = flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponseIDentitySyncConfigurationsActiveDirectories(item.ActiveDirectories)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponseIDentitySyncConfigurationsActiveDirectories(items *[]isegosdk.ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncConfigurationsActiveDirectories) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["directory_id"] = item.DirectoryID
		respItem["domain"] = item.Domain
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponseIDentitySyncSyncSchedule(item *isegosdk.ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponseIDentitySyncSyncSchedule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interval"] = item.Interval
	respItem["interval_unit"] = item.IntervalUnit
	respItem["scheduler_sync"] = item.SchedulerSync
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}
