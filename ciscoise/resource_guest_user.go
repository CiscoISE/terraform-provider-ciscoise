package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGuestUser() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on GuestUser.

- This resource allows the client to update a guest user by name.

- This resource deletes a guest user.

- This resource allows the client to update a guest user by ID.

- This resource deletes a guest user by ID.

- This resource creates a guest user.

- This resource allows the client to change the sponsor password.

- This resource allows the client to update a guest user email by ID.

- This resource allows the client to update a guest user sms by ID.
`,

		CreateContext: resourceGuestUserCreate,
		ReadContext:   resourceGuestUserRead,
		UpdateContext: resourceGuestUserUpdate,
		DeleteContext: resourceGuestUserDelete,
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

						"custom_fields": &schema.Schema{
							Description: `Key value map`,
							Type:        schema.TypeMap,
							Computed:    true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"guest_access_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"from_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"group_tag": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"location": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"to_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"valid_days": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"guest_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"company": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"creation_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"email_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": &schema.Schema{
										Description: `This field is only for Get operation not applicable for Create, Update operations`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"notification_language": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"phone_number": &schema.Schema{
										Description: `Phone number should be E.164 format`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"sms_service_provider": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_name": &schema.Schema{
										Description: `If account needs be created with mobile number, please provide mobile number here`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"guest_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"portal_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reason_for_visit": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sponsor_user_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sponsor_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_reason": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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

						"change_email_address": &schema.Schema{
							Description:  `Flag to allow call to change update a guest user email by ID.`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"change_sms": &schema.Schema{
							Description:  `Flag to allow call to change update a guest user sms by ID.`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"change_password": &schema.Schema{
							Description:  `Flag to allow call to change the sponsor password.`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"custom_fields": &schema.Schema{
							Description: `Key value map`,
							Type:        schema.TypeMap,
							Optional:    true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"guest_access_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"from_date": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"group_tag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"location": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"to_date": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"valid_days": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"guest_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"company": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"creation_time": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"email_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"enabled": &schema.Schema{
										Description:  `This field is only for Get operation not applicable for Create, Update operations`,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"notification_language": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"phone_number": &schema.Schema{
										Description: `Phone number should be E.164 format`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"sms_service_provider": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"user_name": &schema.Schema{
										Description: `If account needs be created with mobile number, please provide mobile number here`,
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"guest_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"portal_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"reason_for_visit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sponsor_user_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sponsor_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status_reason": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceGuestUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestUser create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGuestUserCreateGuestUser(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.GuestUser.GetGuestUserByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceGuestUserRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.GuestUser.GetGuestUserByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceGuestUserRead(ctx, d, m)
		}
	}
	restyResp1, err := client.GuestUser.CreateGuestUser(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateGuestUser", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateGuestUser", err))
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
	return resourceGuestUserRead(ctx, d, m)
}

func resourceGuestUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestUser read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetGuestUserByName")
		vvName := vName

		response1, restyResp1, err := client.GuestUser.GetGuestUserByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenGuestUserGetGuestUserByNameItemName(response1.GuestUser)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestUserByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetGuestUserByID")
		vvID := vID

		response2, restyResp2, err := client.GuestUser.GetGuestUserByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenGuestUserGetGuestUserByIDItemID(response2.GuestUser)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestUserByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceGuestUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestUser update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.GuestUser.GetGuestUserByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestUserByName", err,
				"Failure at GetGuestUserByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.GuestUser != nil {
			vvID = getResp.GuestUser.ID
		}
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestGuestUserUpdateGuestUserByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.GuestUser.UpdateGuestUserByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateGuestUserByID", err, restyResp1.String(),
					"Failure at UpdateGuestUserByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGuestUserByID", err,
				"Failure at UpdateGuestUserByID, unexpected response", ""))
			return diags
		}
		if _, ok := d.GetOk("parameters"); ok {
			if _, ok := d.GetOk("parameters.0"); ok {
				if _, ok := d.GetOk("parameters.0.guest_info"); ok {
					if d.HasChange("parameters.0.guest_info") {
						if _, ok := d.GetOk("parameters.0.guest_info.0"); ok {
							if vPortalID, okPortalID := d.GetOk("parameters.0.portal_id"); okPortalID {
								vvPortalID := vPortalID.(string)

								change_email_address_bool := false
								change_sms_bool := false
								change_password_bool := false
								if change_email_address, ok_change_email_address := d.GetOk("parameters.0.change_email_address"); ok_change_email_address {
									if vchange_email_address_bool := interfaceToBoolPtr(change_email_address); vchange_email_address_bool != nil {
										change_email_address_bool = *vchange_email_address_bool
									}
								}
								if change_sms, ok_change_sms := d.GetOk("parameters.0.change_sms"); ok_change_sms {
									if vchange_sms_bool := interfaceToBoolPtr(change_sms); vchange_sms_bool != nil {
										change_sms_bool = *vchange_sms_bool
									}
								}
								if change_password, ok_change_password := d.GetOk("parameters.0.change_password"); ok_change_password {
									if vchange_password_bool := interfaceToBoolPtr(change_password); vchange_password_bool != nil {
										change_password_bool = *vchange_password_bool
									}
								}
								if change_email_address_bool && d.HasChange("parameters.0.guest_info.0.password") {
									additional_data := []isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData{}
									old_password, new_password := d.GetChange("parameters.0.guest_info.0.password")
									old_password_additional_data := isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData{
										Name:  "currentPassword",
										Value: old_password.(string),
									}
									new_password_additional_data := isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalDataAdditionalData{
										Name:  "newPassword",
										Value: new_password.(string),
									}
									additional_data = append(additional_data, old_password_additional_data)
									additional_data = append(additional_data, new_password_additional_data)
									operational_data := isegosdk.RequestGuestUserChangeSponsorPasswordOperationAdditionalData{AdditionalData: &additional_data}
									request2 := &isegosdk.RequestGuestUserChangeSponsorPassword{}
									request2.OperationAdditionalData = &operational_data
									response2, err := client.GuestUser.ChangeSponsorPassword(vvPortalID, request2)
									if err != nil || response2 == nil {
										log.Printf("[ERROR] response for ChangeSponsorPassword operation => %v", response2.String())
									}
								}
								if change_password_bool && d.HasChange("parameters.0.guest_info.0.email_address") {
									additional_data := []isegosdk.RequestGuestUserUpdateGuestUserEmailOperationAdditionalDataAdditionalData{}
									sender_email := d.Get("parameters.0.guest_info.0.email_address")
									sender_email_additional_data := isegosdk.RequestGuestUserUpdateGuestUserEmailOperationAdditionalDataAdditionalData{
										Name:  "senderEmail",
										Value: sender_email.(string),
									}
									additional_data = append(additional_data, sender_email_additional_data)
									operational_data := isegosdk.RequestGuestUserUpdateGuestUserEmailOperationAdditionalData{AdditionalData: &additional_data}
									request3 := &isegosdk.RequestGuestUserUpdateGuestUserEmail{}
									request3.OperationAdditionalData = &operational_data
									response3, err := client.GuestUser.UpdateGuestUserEmail(vvID, vvPortalID, request3)
									if err != nil || response3 == nil {
										log.Printf("[ERROR] response for ChangeSponsorPassword operation => %v", response3.String())
									}
								}
								if change_sms_bool && d.HasChange("parameters.0.guest_info.0.phone_number") {
									// additional_data := []isegosdk.RequestGuestUserUpdateGuestUserSmsOperationAdditionalDataAdditionalData{}
									// phone_number := d.Get("parameters.0.guest_info.0.phone_number")
									// phone_number_additional_data := isegosdk.RequestGuestUserUpdateGuestUserSmsOperationAdditionalDataAdditionalData{
									// 	Name:  "phoneNumber",
									// 	Value: phone_number.(string),
									// }
									// additional_data = append(additional_data, phone_number_additional_data)
									// operational_data := isegosdk.RequestGuestUserUpdateGuestUserSmsOperationAdditionalData{AdditionalData: &additional_data}
									// request4 := &isegosdk.RequestGuestUserUpdateGuestUserSms{}
									// request4.OperationAdditionalData = &operational_data
									response4, err := client.GuestUser.UpdateGuestUserSms(vvID, vvPortalID)
									if err != nil || response4 == nil {
										log.Printf("[ERROR] response for ChangeSponsorPassword operation => %v", response4.String())
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return resourceGuestUserRead(ctx, d, m)
}

func resourceGuestUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestUser delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.GuestUser.GetGuestUserByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.GuestUser.GetGuestUserByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.GuestUser != nil {
			vvID = getResp.GuestUser.ID
		}
	}
	restyResp1, err := client.GuestUser.DeleteGuestUserByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteGuestUserByID", err, restyResp1.String(),
				"Failure at DeleteGuestUserByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteGuestUserByID", err,
			"Failure at DeleteGuestUserByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestGuestUserCreateGuestUser(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserCreateGuestUser {
	request := isegosdk.RequestGuestUserCreateGuestUser{}
	request.GuestUser = expandRequestGuestUserCreateGuestUserGuestUser(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserCreateGuestUserGuestUser(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserCreateGuestUserGuestUser {
	request := isegosdk.RequestGuestUserCreateGuestUserGuestUser{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_type")))) {
		request.GuestType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status_reason")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status_reason")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status_reason")))) {
		request.StatusReason = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".reason_for_visit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".reason_for_visit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".reason_for_visit")))) {
		request.ReasonForVisit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sponsor_user_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sponsor_user_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sponsor_user_id")))) {
		request.SponsorUserID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sponsor_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sponsor_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sponsor_user_name")))) {
		request.SponsorUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_info")))) {
		request.GuestInfo = expandRequestGuestUserCreateGuestUserGuestUserGuestInfo(ctx, key+".guest_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_access_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_access_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_access_info")))) {
		request.GuestAccessInfo = expandRequestGuestUserCreateGuestUserGuestUserGuestAccessInfo(ctx, key+".guest_access_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_id")))) {
		request.PortalID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_fields")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_fields")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_fields")))) {
		request.CustomFields = expandRequestGuestUserCreateGuestUserGuestUserCustomFields(ctx, key+".custom_fields", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserCreateGuestUserGuestUserGuestInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserCreateGuestUserGuestUserGuestInfo {
	request := isegosdk.RequestGuestUserCreateGuestUserGuestUserGuestInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_name")))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_name")))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".company")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".company")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".company")))) {
		request.Company = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".notification_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".notification_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".notification_language")))) {
		request.NotificationLanguage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email_address")))) {
		request.EmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phone_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phone_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phone_number")))) {
		request.PhoneNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enabled")))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sms_service_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sms_service_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sms_service_provider")))) {
		request.SmsServiceProvider = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserCreateGuestUserGuestUserGuestAccessInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserCreateGuestUserGuestUserGuestAccessInfo {
	request := isegosdk.RequestGuestUserCreateGuestUserGuestUserGuestAccessInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".valid_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".valid_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".valid_days")))) {
		request.ValidDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_date")))) {
		request.FromDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_date")))) {
		request.ToDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location")))) {
		request.Location = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_tag")))) {
		request.GroupTag = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserCreateGuestUserGuestUserCustomFields(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserCreateGuestUserGuestUserCustomFields {
	var request isegosdk.RequestGuestUserCreateGuestUserGuestUserCustomFields
	v := d.Get(fixKeyAccess(key))
	request = v.(map[string]interface{})
	return &request
}

func expandRequestGuestUserUpdateGuestUserByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserUpdateGuestUserByID {
	request := isegosdk.RequestGuestUserUpdateGuestUserByID{}
	request.GuestUser = expandRequestGuestUserUpdateGuestUserByIDGuestUser(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserUpdateGuestUserByIDGuestUser(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUser {
	request := isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUser{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_type")))) {
		request.GuestType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status_reason")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status_reason")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status_reason")))) {
		request.StatusReason = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".reason_for_visit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".reason_for_visit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".reason_for_visit")))) {
		request.ReasonForVisit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sponsor_user_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sponsor_user_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sponsor_user_id")))) {
		request.SponsorUserID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sponsor_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sponsor_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sponsor_user_name")))) {
		request.SponsorUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_info")))) {
		request.GuestInfo = expandRequestGuestUserUpdateGuestUserByIDGuestUserGuestInfo(ctx, key+".guest_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_access_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_access_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_access_info")))) {
		request.GuestAccessInfo = expandRequestGuestUserUpdateGuestUserByIDGuestUserGuestAccessInfo(ctx, key+".guest_access_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_id")))) {
		request.PortalID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_fields")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_fields")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_fields")))) {
		request.CustomFields = expandRequestGuestUserUpdateGuestUserByIDGuestUserCustomFields(ctx, key+".custom_fields", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserUpdateGuestUserByIDGuestUserGuestInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUserGuestInfo {
	request := isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUserGuestInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_name")))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_name")))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".company")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".company")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".company")))) {
		request.Company = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".notification_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".notification_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".notification_language")))) {
		request.NotificationLanguage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email_address")))) {
		request.EmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phone_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phone_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phone_number")))) {
		request.PhoneNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enabled")))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sms_service_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sms_service_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sms_service_provider")))) {
		request.SmsServiceProvider = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserUpdateGuestUserByIDGuestUserGuestAccessInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUserGuestAccessInfo {
	request := isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUserGuestAccessInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".valid_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".valid_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".valid_days")))) {
		request.ValidDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_date")))) {
		request.FromDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_date")))) {
		request.ToDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location")))) {
		request.Location = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_tag")))) {
		request.GroupTag = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserUpdateGuestUserByIDGuestUserCustomFields(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUserCustomFields {
	var request isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUserCustomFields
	v := d.Get(fixKeyAccess(key))
	request = v.(map[string]interface{})
	return &request
}
