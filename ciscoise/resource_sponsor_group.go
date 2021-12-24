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

func resourceSponsorGroup() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SponsorGroup.

- This resource allows the client to update a sponsor group by ID.

- This resource deletes a sponsor group by ID.

- This resource creates a sponsor group.
`,

		CreateContext: resourceSponsorGroupCreate,
		ReadContext:   resourceSponsorGroupRead,
		UpdateContext: resourceSponsorGroupUpdate,
		DeleteContext: resourceSponsorGroupDelete,
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

						"auto_notification": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_permissions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"can_create_random_accounts": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_import_multiple_accounts": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_set_future_start_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_specify_username_prefix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"default_username_prefix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"import_batch_size_limit": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"random_batch_size_limit": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"start_date_future_limit_days": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"guest_types": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_enabled": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"locations": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"manage_permission": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"member_groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"other_permissions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"can_access_via_rest": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_approve_selfreg_guests": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_delete_guest_accounts": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_extend_guest_accounts": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_reinstate_suspended_accounts": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_reset_guest_passwords": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_send_sms_notifications": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_suspend_guest_accounts": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_update_guest_contact_info": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"can_view_guest_passwords": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"limit_approval_to_sponsors_guests": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"require_suspension_reason": &schema.Schema{
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auto_notification": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"create_permissions": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"can_create_random_accounts": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_import_multiple_accounts": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_set_future_start_date": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_specify_username_prefix": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"default_username_prefix": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"import_batch_size_limit": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"random_batch_size_limit": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_date_future_limit_days": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"guest_types": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"is_default_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"is_enabled": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"locations": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"manage_permission": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"member_groups": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"other_permissions": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"can_access_via_rest": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_approve_selfreg_guests": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_delete_guest_accounts": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_extend_guest_accounts": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_reinstate_suspended_accounts": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_reset_guest_passwords": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_send_sms_notifications": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_suspend_guest_accounts": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_update_guest_contact_info": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"can_view_guest_passwords": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"limit_approval_to_sponsors_guests": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"require_suspension_reason": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
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

func resourceSponsorGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SponsorGroup Create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSponsorGroupCreateSponsorGroup(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.SponsorGroup.GetSponsorGroupByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceSponsorGroupRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetSponsorGroupQueryParams{}

		response2, _, err := client.SponsorGroup.GetSponsorGroup(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsSponsorGroupGetSponsorGroup(m, response2, &queryParams2)
			item2, err := searchSponsorGroupGetSponsorGroup(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceSponsorGroupRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.SponsorGroup.CreateSponsorGroup(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSponsorGroup", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSponsorGroup", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceSponsorGroupRead(ctx, d, m)
}

func resourceSponsorGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SponsorGroup Read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		vvName := vName
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetSponsorGroup")
		queryParams1 := isegosdk.GetSponsorGroupQueryParams{}

		response1, restyResp1, err := client.SponsorGroup.GetSponsorGroup(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsSponsorGroupGetSponsorGroup(m, response1, &queryParams1)
		item1, err := searchSponsorGroupGetSponsorGroup(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenSponsorGroupGetSponsorGroupByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsorGroup search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSponsorGroupByID")
		vvID := vID

		response2, restyResp2, err := client.SponsorGroup.GetSponsorGroupByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSponsorGroupGetSponsorGroupByIDItem(response2.SponsorGroup)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsorGroupByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSponsorGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SponsorGroup Update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetSponsorGroupQueryParams{}

		getResp1, _, err := client.SponsorGroup.GetSponsorGroup(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsSponsorGroupGetSponsorGroup(m, getResp1, &queryParams1)
			item1, err := searchSponsorGroupGetSponsorGroup(m, items1, vName, vID)
			if err == nil && item1 != nil {
				if vID != item1.ID {
					vvID = item1.ID
				} else {
					vvID = vID
				}
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestSponsorGroupUpdateSponsorGroupByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.SponsorGroup.UpdateSponsorGroupByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSponsorGroupByID", err, restyResp1.String(),
					"Failure at UpdateSponsorGroupByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSponsorGroupByID", err,
				"Failure at UpdateSponsorGroupByID, unexpected response", ""))
			return diags
		}
	}

	return resourceSponsorGroupRead(ctx, d, m)
}

func resourceSponsorGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SponsorGroup Delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {
		queryParams1 := isegosdk.GetSponsorGroupQueryParams{}

		getResp1, _, err := client.SponsorGroup.GetSponsorGroup(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSponsorGroupGetSponsorGroup(m, getResp1, &queryParams1)
		item1, err := searchSponsorGroupGetSponsorGroup(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 2 {
		vvID = vID
		getResp, _, err := client.SponsorGroup.GetSponsorGroupByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.SponsorGroup.DeleteSponsorGroupByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSponsorGroupByID", err, restyResp1.String(),
				"Failure at DeleteSponsorGroupByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSponsorGroupByID", err,
			"Failure at DeleteSponsorGroupByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSponsorGroupCreateSponsorGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorGroupCreateSponsorGroup {
	request := isegosdk.RequestSponsorGroupCreateSponsorGroup{}
	request.SponsorGroup = expandRequestSponsorGroupCreateSponsorGroupSponsorGroup(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorGroupCreateSponsorGroupSponsorGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorGroupCreateSponsorGroupSponsorGroup {
	request := isegosdk.RequestSponsorGroupCreateSponsorGroupSponsorGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_enabled")))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_group")))) {
		request.IsDefaultGroup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_groups")))) {
		request.MemberGroups = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_types")))) {
		request.GuestTypes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".locations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".locations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".locations")))) {
		request.Locations = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_notification")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_notification")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_notification")))) {
		request.AutoNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_permissions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_permissions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_permissions")))) {
		request.CreatePermissions = expandRequestSponsorGroupCreateSponsorGroupSponsorGroupCreatePermissions(ctx, key+".create_permissions.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".manage_permission")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".manage_permission")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".manage_permission")))) {
		request.ManagePermission = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".other_permissions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".other_permissions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".other_permissions")))) {
		request.OtherPermissions = expandRequestSponsorGroupCreateSponsorGroupSponsorGroupOtherPermissions(ctx, key+".other_permissions.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorGroupCreateSponsorGroupSponsorGroupCreatePermissions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorGroupCreateSponsorGroupSponsorGroupCreatePermissions {
	request := isegosdk.RequestSponsorGroupCreateSponsorGroupSponsorGroupCreatePermissions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_import_multiple_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_import_multiple_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_import_multiple_accounts")))) {
		request.CanImportMultipleAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".import_batch_size_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".import_batch_size_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".import_batch_size_limit")))) {
		request.ImportBatchSizeLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_create_random_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_create_random_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_create_random_accounts")))) {
		request.CanCreateRandomAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".random_batch_size_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".random_batch_size_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".random_batch_size_limit")))) {
		request.RandomBatchSizeLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_username_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_username_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_username_prefix")))) {
		request.DefaultUsernamePrefix = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_specify_username_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_specify_username_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_specify_username_prefix")))) {
		request.CanSpecifyUsernamePrefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_set_future_start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_set_future_start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_set_future_start_date")))) {
		request.CanSetFutureStartDate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date_future_limit_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date_future_limit_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date_future_limit_days")))) {
		request.StartDateFutureLimitDays = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorGroupCreateSponsorGroupSponsorGroupOtherPermissions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorGroupCreateSponsorGroupSponsorGroupOtherPermissions {
	request := isegosdk.RequestSponsorGroupCreateSponsorGroupSponsorGroupOtherPermissions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_update_guest_contact_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_update_guest_contact_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_update_guest_contact_info")))) {
		request.CanUpdateGuestContactInfo = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_view_guest_passwords")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_view_guest_passwords")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_view_guest_passwords")))) {
		request.CanViewGuestPasswords = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_send_sms_notifications")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_send_sms_notifications")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_send_sms_notifications")))) {
		request.CanSendSmsNotifications = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_reset_guest_passwords")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_reset_guest_passwords")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_reset_guest_passwords")))) {
		request.CanResetGuestPasswords = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_extend_guest_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_extend_guest_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_extend_guest_accounts")))) {
		request.CanExtendGuestAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_delete_guest_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_delete_guest_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_delete_guest_accounts")))) {
		request.CanDeleteGuestAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_suspend_guest_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_suspend_guest_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_suspend_guest_accounts")))) {
		request.CanSuspendGuestAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_suspension_reason")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_suspension_reason")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_suspension_reason")))) {
		request.RequireSuspensionReason = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_reinstate_suspended_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_reinstate_suspended_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_reinstate_suspended_accounts")))) {
		request.CanReinstateSuspendedAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_approve_selfreg_guests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_approve_selfreg_guests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_approve_selfreg_guests")))) {
		request.CanApproveSelfregGuests = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit_approval_to_sponsors_guests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit_approval_to_sponsors_guests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit_approval_to_sponsors_guests")))) {
		request.LimitApprovalToSponsorsGuests = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_access_via_rest")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_access_via_rest")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_access_via_rest")))) {
		request.CanAccessViaRest = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorGroupUpdateSponsorGroupByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorGroupUpdateSponsorGroupByID {
	request := isegosdk.RequestSponsorGroupUpdateSponsorGroupByID{}
	request.SponsorGroup = expandRequestSponsorGroupUpdateSponsorGroupByIDSponsorGroup(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorGroupUpdateSponsorGroupByIDSponsorGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroup {
	request := isegosdk.RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_enabled")))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_group")))) {
		request.IsDefaultGroup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_groups")))) {
		request.MemberGroups = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_types")))) {
		request.GuestTypes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".locations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".locations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".locations")))) {
		request.Locations = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_notification")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_notification")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_notification")))) {
		request.AutoNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_permissions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_permissions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_permissions")))) {
		request.CreatePermissions = expandRequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupCreatePermissions(ctx, key+".create_permissions.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".manage_permission")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".manage_permission")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".manage_permission")))) {
		request.ManagePermission = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".other_permissions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".other_permissions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".other_permissions")))) {
		request.OtherPermissions = expandRequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupOtherPermissions(ctx, key+".other_permissions.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupCreatePermissions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupCreatePermissions {
	request := isegosdk.RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupCreatePermissions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_import_multiple_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_import_multiple_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_import_multiple_accounts")))) {
		request.CanImportMultipleAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".import_batch_size_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".import_batch_size_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".import_batch_size_limit")))) {
		request.ImportBatchSizeLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_create_random_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_create_random_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_create_random_accounts")))) {
		request.CanCreateRandomAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".random_batch_size_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".random_batch_size_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".random_batch_size_limit")))) {
		request.RandomBatchSizeLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_username_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_username_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_username_prefix")))) {
		request.DefaultUsernamePrefix = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_specify_username_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_specify_username_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_specify_username_prefix")))) {
		request.CanSpecifyUsernamePrefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_set_future_start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_set_future_start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_set_future_start_date")))) {
		request.CanSetFutureStartDate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date_future_limit_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date_future_limit_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date_future_limit_days")))) {
		request.StartDateFutureLimitDays = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupOtherPermissions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupOtherPermissions {
	request := isegosdk.RequestSponsorGroupUpdateSponsorGroupByIDSponsorGroupOtherPermissions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_update_guest_contact_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_update_guest_contact_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_update_guest_contact_info")))) {
		request.CanUpdateGuestContactInfo = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_view_guest_passwords")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_view_guest_passwords")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_view_guest_passwords")))) {
		request.CanViewGuestPasswords = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_send_sms_notifications")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_send_sms_notifications")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_send_sms_notifications")))) {
		request.CanSendSmsNotifications = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_reset_guest_passwords")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_reset_guest_passwords")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_reset_guest_passwords")))) {
		request.CanResetGuestPasswords = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_extend_guest_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_extend_guest_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_extend_guest_accounts")))) {
		request.CanExtendGuestAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_delete_guest_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_delete_guest_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_delete_guest_accounts")))) {
		request.CanDeleteGuestAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_suspend_guest_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_suspend_guest_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_suspend_guest_accounts")))) {
		request.CanSuspendGuestAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_suspension_reason")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_suspension_reason")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_suspension_reason")))) {
		request.RequireSuspensionReason = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_reinstate_suspended_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_reinstate_suspended_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_reinstate_suspended_accounts")))) {
		request.CanReinstateSuspendedAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_approve_selfreg_guests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_approve_selfreg_guests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_approve_selfreg_guests")))) {
		request.CanApproveSelfregGuests = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit_approval_to_sponsors_guests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit_approval_to_sponsors_guests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit_approval_to_sponsors_guests")))) {
		request.LimitApprovalToSponsorsGuests = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".can_access_via_rest")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".can_access_via_rest")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".can_access_via_rest")))) {
		request.CanAccessViaRest = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsSponsorGroupGetSponsorGroup(m interface{}, response *isegosdk.ResponseSponsorGroupGetSponsorGroup, queryParams *isegosdk.GetSponsorGroupQueryParams) []isegosdk.ResponseSponsorGroupGetSponsorGroupSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseSponsorGroupGetSponsorGroupSearchResultResources
	for response.SearchResult != nil && response.SearchResult.Resources != nil && len(*response.SearchResult.Resources) > 0 {
		respItems = append(respItems, *response.SearchResult.Resources...)
		if response.SearchResult.NextPage != nil && response.SearchResult.NextPage.Rel == "next" {
			href := response.SearchResult.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			if queryParams != nil {
				queryParams.Page = page
				queryParams.Size = size
			}
			response, _, err = client.SponsorGroup.GetSponsorGroup(queryParams)
			if err != nil {
				break
			}
			// All is good, continue to the next page
			continue
		}
		// Does not have next page finish iteration
		break
	}
	return respItems
}

func searchSponsorGroupGetSponsorGroup(m interface{}, items []isegosdk.ResponseSponsorGroupGetSponsorGroupSearchResultResources, name string, id string) (*isegosdk.ResponseSponsorGroupGetSponsorGroupByIDSponsorGroup, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSponsorGroupGetSponsorGroupByIDSponsorGroup
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSponsorGroupGetSponsorGroupByID
			getItem, _, err = client.SponsorGroup.GetSponsorGroupByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSponsorGroupByID")
			}
			foundItem = getItem.SponsorGroup
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSponsorGroupGetSponsorGroupByID
			getItem, _, err = client.SponsorGroup.GetSponsorGroupByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSponsorGroupByID")
			}
			foundItem = getItem.SponsorGroup
			return foundItem, err
		}
	}
	return foundItem, err
}
