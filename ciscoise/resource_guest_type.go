package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

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
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_time": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_access_on_specific_days_times": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"day_time_limits": &schema.Schema{
										Description: `List of Time Ranges for account access`,
										Type:        schema.TypeList,
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
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"end_time": &schema.Schema{
													Description: `End time in HH:mm format`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"start_time": &schema.Schema{
													Description: `Start time in HH:mm format`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"default_duration": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"duration_time_unit": &schema.Schema{
										Description: `Allowed values are:
- DAYS,
- HOURS,
- MINUTES`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"from_first_login": &schema.Schema{
										Description: `When Account Duration starts from first login or specified date`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"max_account_duration": &schema.Schema{
										Description: `Maximum value of Account Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"expiration_notification": &schema.Schema{
							Description: `Expiration Notification Settings`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advance_notification_duration": &schema.Schema{
										Description: `Send Account Expiration Notification Duration before ( Days, Hours, Minutes )`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"advance_notification_units": &schema.Schema{
										Description: `Allowed values are:
- DAYS,
- HOURS,
- MINUTES`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"email_text": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_notification": &schema.Schema{
										Description: `Enable Notification settings`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"send_email_notification": &schema.Schema{
										Description: `Enable Email Notification`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"send_sms_notification": &schema.Schema{
										Description: `Maximum devices guests can register`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"sms_text": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default_type": &schema.Schema{
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
						"login_options": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_guest_portal_bypass": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"failure_action": &schema.Schema{
										Description: `When Guest Exceeds limit this action will be invoked.
Allowed values are:
- Disconnect_Oldest_Connection,
- Disconnect_Newest_Connection`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"identity_group_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"limit_simultaneous_logins": &schema.Schema{
										Description: `Enable Simultaneous Logins`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"max_registered_devices": &schema.Schema{
										Description: `Maximum devices guests can register`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"max_simultaneous_logins": &schema.Schema{
										Description: `Number of Simultaneous Logins`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sponsor_groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_time": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_access_on_specific_days_times": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"day_time_limits": &schema.Schema{
										Description:      `List of Time Ranges for account access`,
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
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
													Type:             schema.TypeList,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"end_time": &schema.Schema{
													Description:      `End time in HH:mm format`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"start_time": &schema.Schema{
													Description:      `Start time in HH:mm format`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
											},
										},
									},
									"default_duration": &schema.Schema{
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"duration_time_unit": &schema.Schema{
										Description: `Allowed values are:
		- DAYS,
		- HOURS,
		- MINUTES`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"from_first_login": &schema.Schema{
										Description:      `When Account Duration starts from first login or specified date`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"max_account_duration": &schema.Schema{
										Description:      `Maximum value of Account Duration`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"expiration_notification": &schema.Schema{
							Description:      `Expiration Notification Settings`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advance_notification_duration": &schema.Schema{
										Description:      `Send Account Expiration Notification Duration before ( Days, Hours, Minutes )`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"advance_notification_units": &schema.Schema{
										Description: `Allowed values are:
		- DAYS,
		- HOURS,
		- MINUTES`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"email_text": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"enable_notification": &schema.Schema{
										Description:      `Enable Notification settings`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"send_email_notification": &schema.Schema{
										Description:      `Enable Email Notification`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"send_sms_notification": &schema.Schema{
										Description:      `Maximum devices guests can register`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"sms_text": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"is_default_type": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
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
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_guest_portal_bypass": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"failure_action": &schema.Schema{
										Description: `When Guest Exceeds limit this action will be invoked.
		Allowed values are:
		- Disconnect_Oldest_Connection,
		- Disconnect_Newest_Connection`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"identity_group_id": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"limit_simultaneous_logins": &schema.Schema{
										Description:      `Enable Simultaneous Logins`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"max_registered_devices": &schema.Schema{
										Description:      `Maximum devices guests can register`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"max_simultaneous_logins": &schema.Schema{
										Description:      `Number of Simultaneous Logins`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"sponsor_groups": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
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
	log.Printf("[DEBUG] Beginning GuestType create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	isEnableAutoImport := clientConfig.EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGuestTypeCreateGuestType(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse2, _, err := client.GuestType.GetGuestTypeByID(vvID)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceGuestTypeRead(ctx, d, m)
			}
		} else {
			queryParams2 := isegosdk.GetGuestTypeQueryParams{}

			response2, _, err := client.GuestType.GetGuestType(&queryParams2)
			if response2 != nil && err == nil {
				items2 := getAllItemsGuestTypeGetGuestType(m, response2, &queryParams2)
				item2, err := searchGuestTypeGetGuestType(m, items2, vvName, vvID)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["id"] = item2.ID
					resourceMap["name"] = vvName
					d.SetId(joinResourceID(resourceMap))
					return resourceGuestTypeRead(ctx, d, m)
				}
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
	return resourceGuestTypeRead(ctx, d, m)
}

func resourceGuestTypeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestType read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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

		response1, restyResp1, err := client.GuestType.GetGuestType(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsGuestTypeGetGuestType(m, response1, &queryParams1)
		item1, err := searchGuestTypeGetGuestType(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenGuestTypeGetGuestTypeByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestType search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestType search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetGuestTypeByID")
		vvID := vID

		response2, restyResp2, err := client.GuestType.GetGuestTypeByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
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
		if err := d.Set("parameters", vItem2); err != nil {
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
	log.Printf("[DEBUG] Beginning GuestType update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
	// NOTE: Added getAllItems and search function to get missing params
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
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestGuestTypeUpdateGuestTypeByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
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
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceGuestTypeRead(ctx, d, m)
}

func resourceGuestTypeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestType delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_type")))) {
		request.IsDefaultType = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".access_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".access_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".access_time")))) {
		request.AccessTime = expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTime(ctx, key+".access_time.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".login_options")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".login_options")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".login_options")))) {
		request.LoginOptions = expandRequestGuestTypeCreateGuestTypeGuestTypeLoginOptions(ctx, key+".login_options.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".expiration_notification")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".expiration_notification")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".expiration_notification")))) {
		request.ExpirationNotification = expandRequestGuestTypeCreateGuestTypeGuestTypeExpirationNotification(ctx, key+".expiration_notification.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sponsor_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sponsor_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sponsor_groups")))) {
		request.SponsorGroups = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTime(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTime {
	request := isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTime{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_first_login")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_first_login")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_first_login")))) {
		request.FromFirstLogin = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_account_duration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_account_duration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_account_duration")))) {
		request.MaxAccountDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".duration_time_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".duration_time_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".duration_time_unit")))) {
		request.DurationTimeUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_duration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_duration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_duration")))) {
		request.DefaultDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_access_on_specific_days_times")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_access_on_specific_days_times")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_access_on_specific_days_times")))) {
		request.AllowAccessOnSpecificDaysTimes = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".day_time_limits")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".day_time_limits")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".day_time_limits")))) {
		request.DayTimeLimits = expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimitsArray(ctx, key+".day_time_limits", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimitsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits {
	request := []isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits{}
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
		i := expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits {
	request := isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeAccessTimeDayTimeLimits{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".days")))) {
		request.Days = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestTypeLoginOptions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeLoginOptions {
	request := isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeLoginOptions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit_simultaneous_logins")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit_simultaneous_logins")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit_simultaneous_logins")))) {
		request.LimitSimultaneousLogins = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_simultaneous_logins")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_simultaneous_logins")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_simultaneous_logins")))) {
		request.MaxSimultaneousLogins = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failure_action")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failure_action")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failure_action")))) {
		request.FailureAction = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_registered_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_registered_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_registered_devices")))) {
		request.MaxRegisteredDevices = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_group_id")))) {
		request.IDentityGroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_portal_bypass")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_portal_bypass")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_portal_bypass")))) {
		request.AllowGuestPortalBypass = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeCreateGuestTypeGuestTypeExpirationNotification(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeExpirationNotification {
	request := isegosdk.RequestGuestTypeCreateGuestTypeGuestTypeExpirationNotification{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_notification")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_notification")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_notification")))) {
		request.EnableNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advance_notification_duration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advance_notification_duration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advance_notification_duration")))) {
		request.AdvanceNotificationDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advance_notification_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advance_notification_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advance_notification_units")))) {
		request.AdvanceNotificationUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".send_email_notification")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".send_email_notification")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".send_email_notification")))) {
		request.SendEmailNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email_text")))) {
		request.EmailText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".send_sms_notification")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".send_sms_notification")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".send_sms_notification")))) {
		request.SendSmsNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sms_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sms_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sms_text")))) {
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_type")))) {
		request.IsDefaultType = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".access_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".access_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".access_time")))) {
		request.AccessTime = expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTime(ctx, key+".access_time.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".login_options")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".login_options")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".login_options")))) {
		request.LoginOptions = expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeLoginOptions(ctx, key+".login_options.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".expiration_notification")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".expiration_notification")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".expiration_notification")))) {
		request.ExpirationNotification = expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeExpirationNotification(ctx, key+".expiration_notification.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sponsor_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sponsor_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sponsor_groups")))) {
		request.SponsorGroups = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTime(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTime {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTime{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_first_login")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_first_login")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_first_login")))) {
		request.FromFirstLogin = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_account_duration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_account_duration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_account_duration")))) {
		request.MaxAccountDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".duration_time_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".duration_time_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".duration_time_unit")))) {
		request.DurationTimeUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_duration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_duration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_duration")))) {
		request.DefaultDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_access_on_specific_days_times")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_access_on_specific_days_times")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_access_on_specific_days_times")))) {
		request.AllowAccessOnSpecificDaysTimes = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".day_time_limits")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".day_time_limits")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".day_time_limits")))) {
		request.DayTimeLimits = expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimitsArray(ctx, key+".day_time_limits", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimitsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits {
	request := []isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits{}
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
		i := expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeAccessTimeDayTimeLimits{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".days")))) {
		request.Days = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeLoginOptions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeLoginOptions {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeLoginOptions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit_simultaneous_logins")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit_simultaneous_logins")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit_simultaneous_logins")))) {
		request.LimitSimultaneousLogins = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_simultaneous_logins")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_simultaneous_logins")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_simultaneous_logins")))) {
		request.MaxSimultaneousLogins = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failure_action")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failure_action")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failure_action")))) {
		request.FailureAction = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_registered_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_registered_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_registered_devices")))) {
		request.MaxRegisteredDevices = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_group_id")))) {
		request.IDentityGroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_portal_bypass")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_portal_bypass")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_portal_bypass")))) {
		request.AllowGuestPortalBypass = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestTypeUpdateGuestTypeByIDGuestTypeExpirationNotification(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeExpirationNotification {
	request := isegosdk.RequestGuestTypeUpdateGuestTypeByIDGuestTypeExpirationNotification{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_notification")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_notification")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_notification")))) {
		request.EnableNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advance_notification_duration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advance_notification_duration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advance_notification_duration")))) {
		request.AdvanceNotificationDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advance_notification_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advance_notification_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advance_notification_units")))) {
		request.AdvanceNotificationUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".send_email_notification")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".send_email_notification")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".send_email_notification")))) {
		request.SendEmailNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email_text")))) {
		request.EmailText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".send_sms_notification")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".send_sms_notification")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".send_sms_notification")))) {
		request.SendSmsNotification = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sms_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sms_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sms_text")))) {
		request.SmsText = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsGuestTypeGetGuestType(m interface{}, response *isegosdk.ResponseGuestTypeGetGuestType, queryParams *isegosdk.GetGuestTypeQueryParams) []isegosdk.ResponseGuestTypeGetGuestTypeSearchResultResources {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
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
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
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
