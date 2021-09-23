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

func resourceGuestType() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on GuestType.

- This resource allows the client to update a guest type.

- This resource deletes a guest type.

- This resource creates a guest type.
`,

		CreateContext: resourceGuestTypeCreate,
		ReadContext:   resourceGuestTypeRead,
		UpdateContext: resourceGuestTypeUpdate,
		DeleteContext: resourceGuestTypeDelete,
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
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_time": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_access_on_specific_days_times": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"day_time_limits": &schema.Schema{
										Description: `List of Time Ranges for account access`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"days": &schema.Schema{
													Description: `List of Days
Values should be one of Week day.
Allowed values are:
- Sunday,
- Monday,
- Tuesday,
- Wednesday,
- Thursday,
- Friday,
- Saturday`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"end_time": &schema.Schema{
													Description: `End time in HH:mm format`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"start_time": &schema.Schema{
													Description: `Start time in HH:mm format`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"default_duration": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"duration_time_unit": &schema.Schema{
										Description: `Allowed values are:
- DAYS,
- HOURS,
- MINUTES`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"from_first_login": &schema.Schema{
										Description: `When Account Duration starts from first login or specified date`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"max_account_duration": &schema.Schema{
										Description: `Maximum value of Account Duration`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expiration_notification": &schema.Schema{
							Description: `Expiration Notification Settings`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advance_notification_duration": &schema.Schema{
										Description: `Send Account Expiration Notification Duration before ( Days, Hours, Minutes )`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"advance_notification_units": &schema.Schema{
										Description: `Allowed values are:
- DAYS,
- HOURS,
- MINUTES`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"email_text": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"enable_notification": &schema.Schema{
										Description: `Enable Notification settings`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"send_email_notification": &schema.Schema{
										Description: `Enable Email Notification`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"send_sms_notification": &schema.Schema{
										Description: `Maximum devices guests can register`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"sms_text": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_default_type": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
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
						"login_options": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_guest_portal_bypass": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"failure_action": &schema.Schema{
										Description: `When Guest Exceeds limit this action will be invoked.
Allowed values are:
- Disconnect_Oldest_Connection,
- Disconnect_Newest_Connection`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"identity_group_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"limit_simultaneous_logins": &schema.Schema{
										Description: `Enable Simultaneous Logins`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"max_registered_devices": &schema.Schema{
										Description: `Maximum devices guests can register`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"max_simultaneous_logins": &schema.Schema{
										Description: `Number of Simultaneous Logins`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sponsor_groups": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceGuestTypeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestGuestTypeCreateGuestType(ctx, "item.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.GuestType.GetGuestTypeByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		queryParams2 := isegosdk.GetGuestTypeQueryParams{}

		response2, _, err := client.GuestType.GetGuestType(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsGuestTypeGetGuestType(m, response2, &queryParams2)
			item2, err := searchGuestTypeGetGuestType(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	restyResp1, err := client.GuestType.CreateGuestType(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateGuestType", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateGuestType", err))
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
	return diags
}

func resourceGuestTypeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetGuestType")
		queryParams1 := isegosdk.GetGuestTypeQueryParams{}

		response1, _, err := client.GuestType.GetGuestType(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestType", err,
				"Failure at GetGuestType, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsGuestTypeGetGuestType(m, response1, &queryParams1)
		item1, err := searchGuestTypeGetGuestType(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetGuestType response", err,
				"Failure when searching item from GetGuestType, unexpected response", ""))
			return diags
		}
		vItem1 := flattenGuestTypeGetGuestTypeByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestType search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetGuestTypeByID")
		vvID := vID

		response2, _, err := client.GuestType.GetGuestTypeByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestTypeByID", err,
				"Failure at GetGuestTypeByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenGuestTypeGetGuestTypeByIDItem(response2.GuestType)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestTypeByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceGuestTypeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		queryParams1 := isegosdk.GetGuestTypeQueryParams{}
		getResp1, _, err := client.GuestType.GetGuestType(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsGuestTypeGetGuestType(m, getResp1, &queryParams1)
			item1, err := searchGuestTypeGetGuestType(m, items1, vName, vID)
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
	if d.HasChange("item") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestGuestTypeUpdateGuestTypeByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.GuestType.UpdateGuestTypeByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateGuestTypeByID", err, restyResp1.String(),
					"Failure at UpdateGuestTypeByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGuestTypeByID", err,
				"Failure at UpdateGuestTypeByID, unexpected response", ""))
			return diags
		}
	}

	return resourceGuestTypeRead(ctx, d, m)
}

func resourceGuestTypeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetGuestTypeQueryParams{}

		getResp1, _, err := client.GuestType.GetGuestType(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsGuestTypeGetGuestType(m, getResp1, &queryParams1)
		item1, err := searchGuestTypeGetGuestType(m, items1, vName, vID)
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
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.GuestType.GetGuestTypeByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.GuestType.DeleteGuestTypeByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteGuestTypeByID", err, restyResp1.String(),
				"Failure at DeleteGuestTypeByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteGuestTypeByID", err,
			"Failure at DeleteGuestTypeByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestGuestTypeCreateGuestType(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeCreateGuestType {
	request := isegosdk.RequestGuestTypeCreateGuestType{}
	request.GuestType = expandRequestGuestTypeCreateGuestTypeGuestType(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestType(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeCreateGuestTypeGuestType {
	request := isegosdk.RequestGuestTypeCreateGuestTypeGuestType{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_default_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_default_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_default_type"))) {
		request.IsDefaultType = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".access_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".access_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".access_time"))) {
		request.AccessTime = expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTime(ctx, key+".access_time.0", d)
	}
	if v, ok := d.GetOkExists(key + ".login_options"); !isEmptyValue(reflect.ValueOf(d.Get(key+".login_options"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".login_options"))) {
		request.LoginOptions = expandRequestGuestTypeCreateGuestTypeGuestTypeLoginOptions(ctx, key+".login_options.0", d)
	}
	if v, ok := d.GetOkExists(key + ".expiration_notification"); !isEmptyValue(reflect.ValueOf(d.Get(key+".expiration_notification"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".expiration_notification"))) {
		request.ExpirationNotification = expandRequestGuestTypeCreateGuestTypeGuestTypeExpirationNotification(ctx, key+".expiration_notification.0", d)
	}
	if v, ok := d.GetOkExists(key + ".sponsor_groups"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sponsor_groups"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sponsor_groups"))) {
		request.SponsorGroups = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTime(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTime {
	request := isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTime{}
	if v, ok := d.GetOkExists(key + ".from_first_login"); !isEmptyValue(reflect.ValueOf(d.Get(key+".from_first_login"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".from_first_login"))) {
		request.FromFirstLogin = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".max_account_duration"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_account_duration"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_account_duration"))) {
		request.MaxAccountDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".duration_time_unit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".duration_time_unit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".duration_time_unit"))) {
		request.DurationTimeUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_duration"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_duration"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_duration"))) {
		request.DefaultDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_access_on_specific_days_times"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_access_on_specific_days_times"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_access_on_specific_days_times"))) {
		request.AllowAccessOnSpecificDaysTimes = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".day_time_limits"); !isEmptyValue(reflect.ValueOf(d.Get(key+".day_time_limits"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".day_time_limits"))) {
		request.DayTimeLimits = expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimitsArray(ctx, key+".day_time_limits", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimitsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits {
	request := []isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits{}
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits {
	request := isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits{}
	if v, ok := d.GetOkExists(key + ".start_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_time"))) {
		request.StartTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".end_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_time"))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".days"))) {
		request.Days = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestTypeLoginOptions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeLoginOptions {
	request := isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeLoginOptions{}
	if v, ok := d.GetOkExists(key + ".limit_simultaneous_logins"); !isEmptyValue(reflect.ValueOf(d.Get(key+".limit_simultaneous_logins"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".limit_simultaneous_logins"))) {
		request.LimitSimultaneousLogins = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".max_simultaneous_logins"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_simultaneous_logins"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_simultaneous_logins"))) {
		request.MaxSimultaneousLogins = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".failure_action"); !isEmptyValue(reflect.ValueOf(d.Get(key+".failure_action"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".failure_action"))) {
		request.FailureAction = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".max_registered_devices"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_registered_devices"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_registered_devices"))) {
		request.MaxRegisteredDevices = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".identity_group_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".identity_group_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".identity_group_id"))) {
		request.IDentityGroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_guest_portal_bypass"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_guest_portal_bypass"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_guest_portal_bypass"))) {
		request.AllowGuestPortalBypass = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestTypeExpirationNotification(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeExpirationNotification {
	request := isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeExpirationNotification{}
	if v, ok := d.GetOkExists(key + ".enable_notification"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_notification"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_notification"))) {
		request.EnableNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".advance_notification_duration"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advance_notification_duration"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advance_notification_duration"))) {
		request.AdvanceNotificationDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".advance_notification_units"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advance_notification_units"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advance_notification_units"))) {
		request.AdvanceNotificationUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".send_email_notification"); !isEmptyValue(reflect.ValueOf(d.Get(key+".send_email_notification"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".send_email_notification"))) {
		request.SendEmailNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".email_text"); !isEmptyValue(reflect.ValueOf(d.Get(key+".email_text"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".email_text"))) {
		request.EmailText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".send_sms_notification"); !isEmptyValue(reflect.ValueOf(d.Get(key+".send_sms_notification"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".send_sms_notification"))) {
		request.SendSmsNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".sms_text"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sms_text"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sms_text"))) {
		request.SmsText = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeByID {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeByID{}
	request.GuestType = expandRequestGuestTypeUpdateGuestTypeByIDGuestType(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestType(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestType {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestType{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_default_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_default_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_default_type"))) {
		request.IsDefaultType = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".access_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".access_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".access_time"))) {
		request.AccessTime = expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTime(ctx, key+".access_time.0", d)
	}
	if v, ok := d.GetOkExists(key + ".login_options"); !isEmptyValue(reflect.ValueOf(d.Get(key+".login_options"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".login_options"))) {
		request.LoginOptions = expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeLoginOptions(ctx, key+".login_options.0", d)
	}
	if v, ok := d.GetOkExists(key + ".expiration_notification"); !isEmptyValue(reflect.ValueOf(d.Get(key+".expiration_notification"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".expiration_notification"))) {
		request.ExpirationNotification = expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeExpirationNotification(ctx, key+".expiration_notification.0", d)
	}
	if v, ok := d.GetOkExists(key + ".sponsor_groups"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sponsor_groups"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sponsor_groups"))) {
		request.SponsorGroups = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTime(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTime {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTime{}
	if v, ok := d.GetOkExists(key + ".from_first_login"); !isEmptyValue(reflect.ValueOf(d.Get(key+".from_first_login"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".from_first_login"))) {
		request.FromFirstLogin = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".max_account_duration"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_account_duration"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_account_duration"))) {
		request.MaxAccountDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".duration_time_unit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".duration_time_unit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".duration_time_unit"))) {
		request.DurationTimeUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_duration"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_duration"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_duration"))) {
		request.DefaultDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_access_on_specific_days_times"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_access_on_specific_days_times"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_access_on_specific_days_times"))) {
		request.AllowAccessOnSpecificDaysTimes = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".day_time_limits"); !isEmptyValue(reflect.ValueOf(d.Get(key+".day_time_limits"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".day_time_limits"))) {
		request.DayTimeLimits = expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimitsArray(ctx, key+".day_time_limits", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimitsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits {
	request := []isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits{}
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits{}
	if v, ok := d.GetOkExists(key + ".start_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_time"))) {
		request.StartTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".end_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_time"))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".days"))) {
		request.Days = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeLoginOptions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeLoginOptions {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeLoginOptions{}
	if v, ok := d.GetOkExists(key + ".limit_simultaneous_logins"); !isEmptyValue(reflect.ValueOf(d.Get(key+".limit_simultaneous_logins"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".limit_simultaneous_logins"))) {
		request.LimitSimultaneousLogins = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".max_simultaneous_logins"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_simultaneous_logins"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_simultaneous_logins"))) {
		request.MaxSimultaneousLogins = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".failure_action"); !isEmptyValue(reflect.ValueOf(d.Get(key+".failure_action"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".failure_action"))) {
		request.FailureAction = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".max_registered_devices"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_registered_devices"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_registered_devices"))) {
		request.MaxRegisteredDevices = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".identity_group_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".identity_group_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".identity_group_id"))) {
		request.IDentityGroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_guest_portal_bypass"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_guest_portal_bypass"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_guest_portal_bypass"))) {
		request.AllowGuestPortalBypass = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeExpirationNotification(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeExpirationNotification {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeExpirationNotification{}
	if v, ok := d.GetOkExists(key + ".enable_notification"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_notification"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_notification"))) {
		request.EnableNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".advance_notification_duration"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advance_notification_duration"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advance_notification_duration"))) {
		request.AdvanceNotificationDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".advance_notification_units"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advance_notification_units"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advance_notification_units"))) {
		request.AdvanceNotificationUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".send_email_notification"); !isEmptyValue(reflect.ValueOf(d.Get(key+".send_email_notification"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".send_email_notification"))) {
		request.SendEmailNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".email_text"); !isEmptyValue(reflect.ValueOf(d.Get(key+".email_text"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".email_text"))) {
		request.EmailText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".send_sms_notification"); !isEmptyValue(reflect.ValueOf(d.Get(key+".send_sms_notification"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".send_sms_notification"))) {
		request.SendSmsNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".sms_text"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sms_text"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sms_text"))) {
		request.SmsText = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsGuestTypeGetGuestType(m interface{}, response *isegosdk.ResponseGuestTypeGetGuestType, queryParams *isegosdk.GetGuestTypeQueryParams) []isegosdk.ResponseGuestTypeGetGuestTypeSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseGuestTypeGetGuestTypeSearchResultResources
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
			response, _, err = client.GuestType.GetGuestType(queryParams)
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

func searchGuestTypeGetGuestType(m interface{}, items []isegosdk.ResponseGuestTypeGetGuestTypeSearchResultResources, name string, id string) (*isegosdk.ResponseGuestTypeGetGuestTypeByIDGuestType, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseGuestTypeGetGuestTypeByIDGuestType
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseGuestTypeGetGuestTypeByID
			getItem, _, err = client.GuestType.GetGuestTypeByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetGuestTypeByID")
			}
			foundItem = getItem.GuestType
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseGuestTypeGetGuestTypeByID
			getItem, _, err = client.GuestType.GetGuestTypeByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetGuestTypeByID")
			}
			foundItem = getItem.GuestType
			return foundItem, err
		}
	}
	return foundItem, err
}
