package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDuoIDentitySync() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Duo-IdentitySync.

- Duo-IdentitySync Create a new IdentitySync configuration

- Duo-Identitysync Update the Identitysync configuration specified in the syncName.

- Duo-Identitysync Delete the Identitysync configuration specified in the syncName.
`,

		CreateContext: resourceDuoIDentitySyncCreate,
		ReadContext:   resourceDuoIDentitySyncRead,
		UpdateContext: resourceDuoIDentitySyncUpdate,
		DeleteContext: resourceDuoIDentitySyncDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ad_groups": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description:      `Active Directory Group ID`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"source": &schema.Schema{
										Description:      `Source of the Active Directory Group`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"configurations": &schema.Schema{
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"last_sync": &schema.Schema{
							Description:      `Time of the last Sync`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"response": &schema.Schema{
							Description: `Identitysync configuration information`,
							Type:        schema.TypeMap,
							Computed:    true,
						},
						"sync_name": &schema.Schema{
							Description:      `Name of the Identitysync configuration`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"sync_schedule": &schema.Schema{
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"sync_status": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceDuoIDentitySyncCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := m.(ClientConfig).EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDuoIDentitySyncCreateIDentitysync(ctx, "parameters.0", d)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vSyncName, okSyncName := resourceItem["sync_name"]
	vvSyncName := interfaceToString(vSyncName)
	if isEnableAutoImport {
		if okSyncName && vvSyncName != "" {
			getResponse2, _, err := client.DuoIDentitySync.GetIDentitysyncBySyncName(vvSyncName)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["sync_name"] = vvSyncName
				d.SetId(joinResourceID(resourceMap))
				return resourceDuoIDentitySyncRead(ctx, d, m)
			}
		} else {
			response2, _, err := client.DuoIDentitySync.GetIDentitysync()
			if response2 != nil && err == nil {
				items2 := getAllItemsDuoIDentitySyncGetIDentitysync(m, response2)
				item2, err := searchDuoIDentitySyncGetIDentitysync(m, items2, vvSyncName)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["sync_name"] = vvSyncName
					d.SetId(joinResourceID(resourceMap))
					return resourceDuoIDentitySyncRead(ctx, d, m)
				}
			}
		}
	}
	resp1, err := client.DuoIDentitySync.CreateIDentitysync(request1)
	if err != nil || resp1 == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateIDentitysync", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["sync_name"] = vvSyncName
	d.SetId(joinResourceID(resourceMap))
	return resourceDuoIDentitySyncRead(ctx, d, m)
}

func resourceDuoIDentitySyncRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvSyncName, okSyncName := resourceMap["sync_name"]

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
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsDuoIDentitySyncGetIDentitysync(m, response1)
		item1, err := searchDuoIDentitySyncGetIDentitysync(m, items1, vvSyncName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenDuoIDentitySyncGetIDentitysyncBySyncNameItemResponse(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIDentitysync search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIDentitysync search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetIDentitysyncBySyncName")
		response2, restyResp2, err := client.DuoIDentitySync.GetIDentitysyncBySyncName(vvSyncName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
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
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIDentitysyncBySyncName response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceDuoIDentitySyncUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvSyncName := resourceMap["sync_name"]

	if d.HasChange("parameters") {

		log.Printf("[DEBUG] Name used for update operation %s", vvSyncName)

		request1 := expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncName(ctx, "parameters.0", d)

		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

		response1, err := client.DuoIDentitySync.UpdateIDentitysyncBySyncName(vvSyncName, request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateIDentitysyncBySyncName", err,
				"Failure at UpdateIDentitysyncBySyncName, unexpected response", ""))

			return diags

		}

	}

	return resourceDuoIDentitySyncRead(ctx, d, m)
}

func resourceDuoIDentitySyncDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vSyncName, okSyncName := resourceMap["sync_name"]
	vvSyncName, _ := resourceMap["sync_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okSyncName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.DuoIDentitySync.GetIDentitysync()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsDuoIDentitySyncGetIDentitysync(m, getResp1)
		item1, err := searchDuoIDentitySyncGetIDentitysync(m, items1, vvSyncName)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vSyncName != item1.IDentitySync.SyncName {
			vvSyncName = item1.IDentitySync.SyncName
		} else {
			vvSyncName = vSyncName
		}
	}
	if selectedMethod == 2 {
		getResp, _, err := client.DuoIDentitySync.GetIDentitysyncBySyncName(vvSyncName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, err := client.DuoIDentitySync.DeleteIDentitySyncBySyncName(vvSyncName)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteIDentitySyncBySyncName", err,
			"Failure at DeleteIDentitySyncBySyncName, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDuoIDentitySyncCreateIDentitysync(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncCreateIDentitysync {
	request := isegosdk.RequestDuoIDentitySyncCreateIDentitysync{}
	request.IDentitySync = expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySync(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySync(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySync {
	request := isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySync{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ad_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ad_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ad_groups")))) {
		request.AdGroups = expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncAdGroupsArray(ctx, key+".ad_groups", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configurations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configurations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configurations")))) {
		request.Configurations = expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurations(ctx, key+".configurations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_sync")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_sync")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_sync")))) {
		request.LastSync = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_name")))) {
		request.SyncName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_schedule")))) {
		request.SyncSchedule = expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncSyncSchedule(ctx, key+".sync_schedule.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_status")))) {
		request.SyncStatus = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncAdGroupsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncAdGroups {
	request := []isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncAdGroups{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncAdGroups(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncAdGroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncAdGroups {
	request := isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncAdGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source")))) {
		request.Source = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurations {
	request := isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".active_directories")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".active_directories")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".active_directories")))) {
		request.ActiveDirectories = expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurationsActiveDirectoriesArray(ctx, key+".active_directories", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurationsActiveDirectoriesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurationsActiveDirectories {
	request := []isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurationsActiveDirectories{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurationsActiveDirectories(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurationsActiveDirectories(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurationsActiveDirectories {
	request := isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncConfigurationsActiveDirectories{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".directory_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".directory_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".directory_id")))) {
		request.DirectoryID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain")))) {
		request.Domain = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncCreateIDentitysyncIDentitySyncSyncSchedule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncSyncSchedule {
	request := isegosdk.RequestDuoIDentitySyncCreateIDentitysyncIDentitySyncSyncSchedule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval")))) {
		request.Interval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval_unit")))) {
		request.IntervalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scheduler_sync")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scheduler_sync")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scheduler_sync")))) {
		request.SchedulerSync = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncName {
	request := isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncName{}
	request.IDentitySync = expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySync(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySync(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySync {
	request := isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySync{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ad_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ad_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ad_groups")))) {
		request.AdGroups = expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncAdGroupsArray(ctx, key+".ad_groups", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configurations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configurations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configurations")))) {
		request.Configurations = expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurations(ctx, key+".configurations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_sync")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_sync")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_sync")))) {
		request.LastSync = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_name")))) {
		request.SyncName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_schedule")))) {
		request.SyncSchedule = expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncSyncSchedule(ctx, key+".sync_schedule.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_status")))) {
		request.SyncStatus = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncAdGroupsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncAdGroups {
	request := []isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncAdGroups{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncAdGroups(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncAdGroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncAdGroups {
	request := isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncAdGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source")))) {
		request.Source = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurations {
	request := isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".active_directories")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".active_directories")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".active_directories")))) {
		request.ActiveDirectories = expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurationsActiveDirectoriesArray(ctx, key+".active_directories", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurationsActiveDirectoriesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurationsActiveDirectories {
	request := []isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurationsActiveDirectories{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurationsActiveDirectories(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurationsActiveDirectories(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurationsActiveDirectories {
	request := isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncConfigurationsActiveDirectories{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".directory_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".directory_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".directory_id")))) {
		request.DirectoryID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain")))) {
		request.Domain = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncSyncSchedule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncSyncSchedule {
	request := isegosdk.RequestDuoIDentitySyncUpdateIDentitysyncBySyncNameIDentitySyncSyncSchedule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval")))) {
		request.Interval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval_unit")))) {
		request.IntervalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scheduler_sync")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scheduler_sync")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scheduler_sync")))) {
		request.SchedulerSync = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsDuoIDentitySyncGetIDentitysync(m interface{}, response *isegosdk.ResponseDuoIDentitySyncGetIDentitysync) []isegosdk.ResponseDuoIDentitySyncGetIDentitysyncResponse {
	var respItems []isegosdk.ResponseDuoIDentitySyncGetIDentitysyncResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchDuoIDentitySyncGetIDentitysync(m interface{}, items []isegosdk.ResponseDuoIDentitySyncGetIDentitysyncResponse, syncName string) (*isegosdk.ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponse, error) {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var err error
	var foundItem *isegosdk.ResponseDuoIDentitySyncGetIDentitysyncBySyncNameResponse
	for _, item := range items {
		if syncName != "" && item.Name == syncName {
			var getItem *isegosdk.ResponseDuoIDentitySyncGetIDentitysyncBySyncName
			getItem, _, err = client.DuoIDentitySync.GetIDentitysyncBySyncName(syncName)
			if err != nil {
				return foundItem, nil
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetIDentitysyncBySyncName")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
