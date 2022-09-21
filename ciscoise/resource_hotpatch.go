package ciscoise

import (
	"context"
	"reflect"
	"time"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHotpatch() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Hotpatch.

- Triggers hot patch installation on the Cisco ISE node. A task ID is returned which  can be used to monitor the
progress of the hot patch installation process.  As hot patch installation triggers the Cisco ISE to restart, the task
API becomes  unavailable for a certain period of time.

- Triggers hot patch rollback on the Cisco ISE node. A task ID is returned which  can be used to monitor the progress of
the hot patch rollback process.  As hot patch rollback triggers the Cisco ISE to restart, the task API becomes
unavailable for a certain period of time.
`,

		CreateContext: resourceHotpatchCreate,
		ReadContext:   resourceHotpatchRead,
		UpdateContext: resourceHotpatchUpdate,
		DeleteContext: resourceHotpatchDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(HOTPATCH_INSTALL_TIMEOUT_SLEEP),
			Delete: schema.DefaultTimeout(HOTPATCH_ROLLBACK_TIMEOUT_SLEEP),
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"hotpatch_name": &schema.Schema{
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: diffSupressHotpatchName(),
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hotpatch_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"install_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceHotpatchCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Hotpatch create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	vRepositoryName, okRepositoryName := resourceItem["repository_name"]
	vHotpatchName, okPatchName := resourceItem["hotpatch_name"]
	var vvRepositoryName string
	var vvHotpatchName string
	if okRepositoryName {
		vvRepositoryName = vRepositoryName.(string)
	}
	if okPatchName {
		vvHotpatchName = vHotpatchName.(string)
	}

	getResponse1, _, err := client.Patching.ListInstalledHotpatches()
	if err == nil && getResponse1 != nil {
		item1, err := searchHotPatch(m, getResponse1, vvHotpatchName)
		if err == nil && item1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["repository_name"] = vvRepositoryName
			resourceMap["hotpatch_name"] = vvHotpatchName
			d.SetId(joinResourceID(resourceMap))
			return resourceHotpatchRead(ctx, d, m)
		}
	}

	request1 := expandRequestHotpatchInstallInstallHotpatch(ctx, "parameters.0", d)
	response1, restyResp1, err := client.Patching.InstallHotpatch(request1)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if response1 != nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing InstallPatch", err,
				"Failure at InstallPatch, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing InstallPatch", err,
			"Failure at InstallPatch, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %s", restyResp1.String())

	if response1.Response != nil {
		taskID := response1.Response.ID
		if taskID != "" {
			taskResponse, restyResp2, err := client.Tasks.GetTaskStatusByID(taskID)
			if err != nil || taskResponse == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskStatusByID", err,
					"Failure at GetTaskStatusByID, unexpected response", ""))
				return diags
			}
			for err != nil || taskResponse == nil || taskResponse.ExecutionStatus == "IN_PROGRESS" {
				time.Sleep(HOTPATCH_INSTALL_TIMEOUT_SLEEP)
				taskResponse, restyResp2, err = client.Tasks.GetTaskStatusByID(taskID)
				if err != nil || taskResponse == nil {
					if restyResp2 != nil {
						log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
					}
				}
			}
		}
	}

	resourceMap := make(map[string]string)
	resourceMap["repository_name"] = vvRepositoryName
	resourceMap["patch_name"] = vvHotpatchName
	d.SetId(joinResourceID(resourceMap))
	return resourceHotpatchRead(ctx, d, m)
}

func resourceHotpatchRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Hotpatch read for id=[%s]", d.Id())

	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvHotpatchName, _ := resourceMap["hotpatch_name"]

	response1, restyResp1, err := client.Patching.ListInstalledHotpatches()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	item1, err := searchHotPatch(m, response1, vvHotpatchName)
	if err != nil || item1 == nil {
		d.SetId("")
		return diags
	}

	vItem1 := flattenPatchingListInstalledHotpatchesItems(item1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ListInstalledHotpatches response to item",
			err))
		return diags
	}
	// if err := d.Set("parameters", remove_parameters(vItem1, "repository_name")); err != nil {
	// 	diags = append(diags, diagError(
	// 		"Failure when setting ListInstalledHotpatches response to parameters",
	// 		err))
	// 	return diags
	// }
	return diags
}

func resourceHotpatchUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Hotpatch update for id=[%s]", d.Id())
	log.Printf("[DEBUG] Missing Hotpatch update on Cisco ISE. It will only be update it on Terraform")
	// _ = d.Set("last_updated", getUnixTimeString())
	return resourceHotpatchRead(ctx, d, m)
}

func resourceHotpatchDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Hotpatch delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvHotpatchName, _ := resourceMap["hotpatch_name"]

	getResp1, restyResp1, err := client.Patching.ListInstalledHotpatches()
	if err != nil || getResp1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*getResp1))
	item1, err := searchHotPatch(m, getResp1, vvHotpatchName)
	if err != nil || item1 == nil {
		d.SetId("")
		return diags
	}

	request1 := expandRequestHotpatchRollbackRollbackHotpatch(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Patching.RollbackHotpatch(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RollbackPatch", err,
			"Failure at RollbackPatch, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response != nil {
		taskID := response1.Response.ID
		if taskID != "" {
			taskResponse, restyResp2, err := client.Tasks.GetTaskStatusByID(taskID)
			if err != nil || taskResponse == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskStatusByID", err,
					"Failure at GetTaskStatusByID, unexpected response", ""))
				return diags
			}
			for err != nil || taskResponse == nil || taskResponse.ExecutionStatus == "IN_PROGRESS" {
				time.Sleep(HOTPATCH_ROLLBACK_TIMEOUT_SLEEP)
				taskResponse, restyResp2, err = client.Tasks.GetTaskStatusByID(taskID)
				if err != nil || taskResponse == nil {
					if restyResp2 != nil {
						log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
					}
				}
			}
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func searchHotPatch(m interface{}, item *isegosdk.ResponsePatchingListInstalledHotpatches, hotpatchName string) (*[]isegosdk.ResponsePatchingListInstalledHotpatchesResponse, error) {
	var err error
	if item == nil {
		return nil, err
	}
	if item.Response == nil {
		return nil, err
	}
	items := *item.Response
	var foundItems []isegosdk.ResponsePatchingListInstalledHotpatchesResponse
	for _, item := range items {
		if item.HotpatchName == hotpatchName {
			// Call get by _ method and set value to foundItem and return
			foundItems = append(foundItems, item)
			return &foundItems, err
		}
	}
	return nil, err
}

func expandRequestHotpatchInstallInstallHotpatch(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPatchingInstallHotpatch {
	request := isegosdk.RequestPatchingInstallHotpatch{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hotpatch_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hotpatch_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hotpatch_name")))) {
		request.HotpatchName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".repository_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".repository_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".repository_name")))) {
		request.RepositoryName = interfaceToString(v)
	}
	return &request
}

func expandRequestHotpatchRollbackRollbackHotpatch(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPatchingRollbackHotpatch {
	request := isegosdk.RequestPatchingRollbackHotpatch{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hotpatch_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hotpatch_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hotpatch_name")))) {
		if d.HasChange(fixKeyAccess(key + ".hotpatch_name")) {
			if old, new := d.GetChange(fixKeyAccess(key + ".hotpatch_name")); !reflect.DeepEqual(old, new) {
				request.HotpatchName = interfaceToString(old)
			}
		} else {
			request.HotpatchName = interfaceToString(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".repository_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".repository_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".repository_name")))) {
		if d.HasChange(fixKeyAccess(key + ".repository_name")) {
			if old, new := d.GetChange(fixKeyAccess(key + ".repository_name")); !reflect.DeepEqual(old, new) {
				request.RepositoryName = interfaceToString(old)
			}
		} else {
			request.RepositoryName = interfaceToString(v)
		}
	}
	return &request
}
