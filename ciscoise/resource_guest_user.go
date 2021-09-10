package ciscoise

import (
	"context"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGuestUser() *schema.Resource {
	return &schema.Resource{

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
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_fields": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"guest_access_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"from_date": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"group_tag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"location": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"to_date": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"valid_days": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"guest_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"company": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"creation_time": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"email_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"notification_language": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"phone_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sms_service_provider": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"user_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"guest_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
							Optional: true,
							Computed: true,
						},
						"portal_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"reason_for_visit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sponsor_user_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sponsor_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status_reason": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceGuestUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestGuestUserCreateGuestUser(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

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
			return diags
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.GuestUser.GetGuestUserByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
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
	return diags
}

func resourceGuestUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

		response1, _, err := client.GuestUser.GetGuestUserByName(vvName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestUserByName", err,
				"Failure at GetGuestUserByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

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

		response2, _, err := client.GuestUser.GetGuestUserByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestUserByID", err,
				"Failure at GetGuestUserByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

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
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestGuestUserUpdateGuestUserByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.GuestUser.UpdateGuestUserByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
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
	}

	return resourceGuestUserRead(ctx, d, m)
}

func resourceGuestUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
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
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".guest_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".guest_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".guest_type"))) {
		request.GuestType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".status"); !isEmptyValue(reflect.ValueOf(d.Get(key+".status"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".status"))) {
		request.Status = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".status_reason"); !isEmptyValue(reflect.ValueOf(d.Get(key+".status_reason"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".status_reason"))) {
		request.StatusReason = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".reason_for_visit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".reason_for_visit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".reason_for_visit"))) {
		request.ReasonForVisit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sponsor_user_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sponsor_user_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sponsor_user_id"))) {
		request.SponsorUserID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sponsor_user_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sponsor_user_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sponsor_user_name"))) {
		request.SponsorUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".guest_info"); !isEmptyValue(reflect.ValueOf(d.Get(key+".guest_info"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".guest_info"))) {
		request.GuestInfo = expandRequestGuestUserCreateGuestUserGuestUserGuestInfo(ctx, key+".guest_info.0", d)
	}
	if v, ok := d.GetOkExists(key + ".guest_access_info"); !isEmptyValue(reflect.ValueOf(d.Get(key+".guest_access_info"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".guest_access_info"))) {
		request.GuestAccessInfo = expandRequestGuestUserCreateGuestUserGuestUserGuestAccessInfo(ctx, key+".guest_access_info.0", d)
	}
	if v, ok := d.GetOkExists(key + ".portal_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_id"))) {
		request.PortalID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".custom_fields"); !isEmptyValue(reflect.ValueOf(d.Get(key+".custom_fields"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".custom_fields"))) {
		customFields := v.([]interface{})[0].(map[string]interface{})
		request.CustomFields = &customFields
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserCreateGuestUserGuestUserGuestInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserCreateGuestUserGuestUserGuestInfo {
	request := isegosdk.RequestGuestUserCreateGuestUserGuestUserGuestInfo{}
	if v, ok := d.GetOkExists(key + ".first_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".first_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".first_name"))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".last_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".last_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".last_name"))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".company"); !isEmptyValue(reflect.ValueOf(d.Get(key+".company"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".company"))) {
		request.Company = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".creation_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".creation_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".creation_time"))) {
		request.CreationTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".notification_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".notification_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".notification_language"))) {
		request.NotificationLanguage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".user_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".user_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".user_name"))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".email_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".email_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".email_address"))) {
		request.EmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".phone_number"); !isEmptyValue(reflect.ValueOf(d.Get(key+".phone_number"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".phone_number"))) {
		request.PhoneNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".sms_service_provider"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sms_service_provider"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sms_service_provider"))) {
		request.SmsServiceProvider = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserCreateGuestUserGuestUserGuestAccessInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserCreateGuestUserGuestUserGuestAccessInfo {
	request := isegosdk.RequestGuestUserCreateGuestUserGuestUserGuestAccessInfo{}
	if v, ok := d.GetOkExists(key + ".valid_days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".valid_days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".valid_days"))) {
		request.ValidDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".from_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".from_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".from_date"))) {
		request.FromDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".to_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".to_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".to_date"))) {
		request.ToDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".location"); !isEmptyValue(reflect.ValueOf(d.Get(key+".location"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".location"))) {
		request.Location = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ssid"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ssid"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ssid"))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".group_tag"))) {
		request.GroupTag = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
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
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".guest_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".guest_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".guest_type"))) {
		request.GuestType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".status"); !isEmptyValue(reflect.ValueOf(d.Get(key+".status"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".status"))) {
		request.Status = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".status_reason"); !isEmptyValue(reflect.ValueOf(d.Get(key+".status_reason"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".status_reason"))) {
		request.StatusReason = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".reason_for_visit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".reason_for_visit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".reason_for_visit"))) {
		request.ReasonForVisit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sponsor_user_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sponsor_user_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sponsor_user_id"))) {
		request.SponsorUserID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sponsor_user_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sponsor_user_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sponsor_user_name"))) {
		request.SponsorUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".guest_info"); !isEmptyValue(reflect.ValueOf(d.Get(key+".guest_info"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".guest_info"))) {
		request.GuestInfo = expandRequestGuestUserUpdateGuestUserByIDGuestUserGuestInfo(ctx, key+".guest_info.0", d)
	}
	if v, ok := d.GetOkExists(key + ".guest_access_info"); !isEmptyValue(reflect.ValueOf(d.Get(key+".guest_access_info"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".guest_access_info"))) {
		request.GuestAccessInfo = expandRequestGuestUserUpdateGuestUserByIDGuestUserGuestAccessInfo(ctx, key+".guest_access_info.0", d)
	}
	if v, ok := d.GetOkExists(key + ".portal_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_id"))) {
		request.PortalID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".custom_fields"); !isEmptyValue(reflect.ValueOf(d.Get(key+".custom_fields"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".custom_fields"))) {
		customFields := v.([]interface{})[0].(map[string]interface{})
		request.CustomFields = &customFields
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserUpdateGuestUserByIDGuestUserGuestInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUserGuestInfo {
	request := isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUserGuestInfo{}
	if v, ok := d.GetOkExists(key + ".first_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".first_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".first_name"))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".last_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".last_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".last_name"))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".company"); !isEmptyValue(reflect.ValueOf(d.Get(key+".company"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".company"))) {
		request.Company = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".creation_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".creation_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".creation_time"))) {
		request.CreationTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".notification_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".notification_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".notification_language"))) {
		request.NotificationLanguage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".user_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".user_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".user_name"))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".email_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".email_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".email_address"))) {
		request.EmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".phone_number"); !isEmptyValue(reflect.ValueOf(d.Get(key+".phone_number"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".phone_number"))) {
		request.PhoneNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".sms_service_provider"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sms_service_provider"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sms_service_provider"))) {
		request.SmsServiceProvider = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestUserUpdateGuestUserByIDGuestUserGuestAccessInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUserGuestAccessInfo {
	request := isegosdk.RequestGuestUserUpdateGuestUserByIDGuestUserGuestAccessInfo{}
	if v, ok := d.GetOkExists(key + ".valid_days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".valid_days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".valid_days"))) {
		request.ValidDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".from_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".from_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".from_date"))) {
		request.FromDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".to_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".to_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".to_date"))) {
		request.ToDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".location"); !isEmptyValue(reflect.ValueOf(d.Get(key+".location"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".location"))) {
		request.Location = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ssid"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ssid"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ssid"))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".group_tag"))) {
		request.GroupTag = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
