package ciscoise

import (
	"context"
	"log"

	"fmt"
	"reflect"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceDuoIDentitySyncStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Duo-IdentitySync.

- duo-identitysync update sync status.
`,

		CreateContext: resourceDuoIDentitySyncStatusCreate,
		ReadContext:   resourceDuoIDentitySyncStatusRead,
		DeleteContext: resourceDuoIDentitySyncStatusDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sync_name": &schema.Schema{
							Description: `syncName path parameter. Sync connection to be updated`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"reason": &schema.Schema{
							Description: `Reason user failed sync`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"user": &schema.Schema{
							Description: `User to be synced to Duo`,
							Type:        schema.TypeMap,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceDuoIDentitySyncStatusCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))

	vSyncName := resourceItem["sync_name"]

	vvSyncName := vSyncName.(string)
	request1 := expandRequestDuoIDentitySyncStatusUpdateStatus(ctx, "parameters.0", d)

	response1, err := client.DuoIDentitySync.UpdateStatus(vvSyncName, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing UpdateStatus", err,
			"Failure at UpdateStatus, unexpected response", ""))
		return diags
	}

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateStatus response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

	return diags
}

func expandRequestDuoIDentitySyncStatusUpdateStatus(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncUpdateStatus {
	request := isegosdk.RequestDuoIDentitySyncUpdateStatus{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".error_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".error_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".error_list")))) {
		request.ErrorList = expandRequestDuoIDentitySyncStatusUpdateStatusErrorListArray(ctx, key+".error_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	return &request
}

func expandRequestDuoIDentitySyncStatusUpdateStatusErrorListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDuoIDentitySyncUpdateStatusErrorList {
	request := []isegosdk.RequestDuoIDentitySyncUpdateStatusErrorList{}
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
		i := expandRequestDuoIDentitySyncStatusUpdateStatusErrorList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDuoIDentitySyncStatusUpdateStatusErrorList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncUpdateStatusErrorList {
	request := isegosdk.RequestDuoIDentitySyncUpdateStatusErrorList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".reason")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".reason")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".reason")))) {
		request.Reason = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user")))) {
		request.User = expandRequestDuoIDentitySyncStatusUpdateStatusErrorListUser(ctx, key+".user.0", d)
	}
	return &request
}

func expandRequestDuoIDentitySyncStatusUpdateStatusErrorListUser(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoIDentitySyncUpdateStatusErrorListUser {
	request := isegosdk.RequestDuoIDentitySyncUpdateStatusErrorListUser{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".directoryname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".directoryname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".directoryname")))) {
		request.Directoryname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email")))) {
		request.Email = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".firstname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".firstname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".firstname")))) {
		request.Firstname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".groupname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".groupname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".groupname")))) {
		request.Groupname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lastname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lastname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lastname")))) {
		request.Lastname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".notes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".notes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".notes")))) {
		request.Notes = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".realname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".realname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".realname")))) {
		request.Realname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	return &request
}

func resourceDuoIDentitySyncStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceDuoIDentitySyncStatusUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceDuoIDentitySyncStatusRead(ctx, d, m)
}

func resourceDuoIDentitySyncStatusDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
