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

func resourcePatch() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Patch.

- Triggers patch installation on the Cisco ISE node. A task ID is returned which can be used to monitor the progress of
the patch installation process. As the patch   installation triggers the Cisco ISE to restart, the task API becomes
unavailable for  a certain period of time.

- Triggers patch rollback on the Cisco ISE node. A task ID is returned which can be used to monitor the progress of the
patch rollback process. As the patch   rollback triggers the Cisco ISE to restart, the task API becomes unavailable for
a certain period of time.
`,

		CreateContext: resourcePatchCreate,
		ReadContext:   resourcePatchRead,
		UpdateContext: resourcePatchUpdate,
		DeleteContext: resourcePatchDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(PATCH_INSTALL_TIMEOUT_SLEEP),
			Delete: schema.DefaultTimeout(PATCH_ROLLBACK_TIMEOUT_SLEEP),
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
						"patch_number": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"patch_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"patch_number": &schema.Schema{
							Type:     schema.TypeInt,
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

func resourcePatchCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Patch create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	vPatchNumber, okPatchNumber := resourceItem["patch_number"]
	vPatchName, okPatchName := resourceItem["patch_name"]
	var vvPatchNumber int
	var vvPatchName string
	if okPatchNumber {
		vvPatchNumber = vPatchNumber.(int)
	}
	if okPatchName {
		vvPatchName = vPatchName.(string)
	}

	getResponse1, _, err := client.Patching.ListInstalledPatches()
	if err == nil && getResponse1 != nil {
		item1, err := searchPatch(m, getResponse1, &vvPatchNumber)
		if err == nil && item1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["patch_number"] = interfaceToString(vvPatchNumber)
			resourceMap["patch_name"] = vvPatchName
			d.SetId(joinResourceID(resourceMap))
			return resourcePatchRead(ctx, d, m)
		}
	}

	request1 := expandRequestPatchInstallInstallPatch(ctx, "parameters.0", d)
	response1, restyResp1, err := client.Patching.InstallPatch(request1)
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
				time.Sleep(PATCH_INSTALL_TIMEOUT_SLEEP)
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
	resourceMap["patch_number"] = interfaceToString(vvPatchNumber)
	resourceMap["patch_name"] = vvPatchName
	d.SetId(joinResourceID(resourceMap))
	return resourcePatchRead(ctx, d, m)
}

func resourcePatchRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Patch read for id=[%s]", d.Id())

	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPatchNumber, _ := resourceMap["patch_number"]
	vvPatchNumber := interfaceToIntPtr(vPatchNumber)

	response1, restyResp1, err := client.Patching.ListInstalledPatches()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	item1, err := searchPatch(m, response1, vvPatchNumber)
	if err != nil || item1 == nil {
		d.SetId("")
		return diags
	}

	vItem1 := flattenPatchingListInstalledPatchesItemPatchVersion(item1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ListInstalledPatches response to item",
			err))
		return diags
	}
	if err := d.Set("parameters", remove_parameters(vItem1, "install_date")); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ListInstalledPatches response to parameters",
			err))
		return diags
	}
	return diags
}

func resourcePatchUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Patch update for id=[%s]", d.Id())
	log.Printf("[DEBUG] Missing Patch update on Cisco ISE. It will only be update it on Terraform")
	// _ = d.Set("last_updated", getUnixTimeString())
	return resourcePatchRead(ctx, d, m)
}

func resourcePatchDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Patch delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPatchNumber, _ := resourceMap["patch_number"]
	vvPatchNumber := interfaceToIntPtr(vPatchNumber)

	getResp1, restyResp1, err := client.Patching.ListInstalledPatches()
	if err != nil || getResp1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*getResp1))
	item1, err := searchPatch(m, getResp1, vvPatchNumber)
	if err != nil || item1 == nil {
		d.SetId("")
		return diags
	}

	request1 := expandRequestPatchRollbackRollbackPatch(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Patching.RollbackPatch(request1)

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
				time.Sleep(PATCH_ROLLBACK_TIMEOUT_SLEEP)
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

func searchPatch(m interface{}, item *isegosdk.ResponsePatchingListInstalledPatches, patchNumber *int) (*[]isegosdk.ResponsePatchingListInstalledPatchesPatchVersion, error) {
	var err error
	if item == nil {
		return nil, err
	}
	if item.PatchVersion == nil {
		return nil, err
	}
	if patchNumber == nil {
		return nil, err
	}
	items := *item.PatchVersion
	var foundItems []isegosdk.ResponsePatchingListInstalledPatchesPatchVersion
	for _, item := range items {
		if item.PatchNumber != nil && *item.PatchNumber == *patchNumber {
			// Call get by _ method and set value to foundItem and return
			foundItems = append(foundItems, item)
			return &foundItems, err
		}
	}
	return nil, err
}

func expandRequestPatchInstallInstallPatch(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPatchingInstallPatch {
	request := isegosdk.RequestPatchingInstallPatch{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".patch_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".patch_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".patch_name")))) {
		request.PatchName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".repository_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".repository_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".repository_name")))) {
		request.RepositoryName = interfaceToString(v)
	}
	return &request
}

func expandRequestPatchRollbackRollbackPatch(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPatchingRollbackPatch {
	request := isegosdk.RequestPatchingRollbackPatch{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".patch_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".patch_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".patch_number")))) {
		if d.HasChange(fixKeyAccess(key + ".patch_number")) {
			if old, new := d.GetChange(fixKeyAccess(key + ".patch_number")); !reflect.DeepEqual(old, new) {
				request.PatchNumber = interfaceToIntPtr(old)
			}
		} else {
			request.PatchNumber = interfaceToIntPtr(v)
		}
	}
	return &request
}
