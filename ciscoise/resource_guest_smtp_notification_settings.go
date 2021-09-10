package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGuestSmtpNotificationSettings() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceGuestSmtpNotificationSettingsCreate,
		ReadContext:   resourceGuestSmtpNotificationSettingsRead,
		UpdateContext: resourceGuestSmtpNotificationSettingsUpdate,
		DeleteContext: resourceGuestSmtpNotificationSettingsDelete,
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

						"connection_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default_from_address": &schema.Schema{
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
							Type:             schema.TypeList,
							DiffSuppressFunc: diffSuppressAlways(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
									"rel": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
									"type": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
								},
							},
						},
						"notification_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"smtp_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"smtp_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"use_default_from_address": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"use_password_authentication": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"use_tlsor_ssl_encryption": &schema.Schema{
							Type:     schema.TypeBool,
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
		},
	}
}

func resourceGuestSmtpNotificationSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestGuestSmtpNotificationSettingsCreateGuestSmtpNotificationSettings(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okID && vvID != "" {
		getResponse2, _, err := client.GuestSmtpNotificationConfiguration.GetGuestSmtpNotificationSettingsByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		queryParams2 := isegosdk.GetGuestSmtpNotificationSettingsQueryParams{}

		response2, _, err := client.GuestSmtpNotificationConfiguration.GetGuestSmtpNotificationSettings(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings(m, response2, &queryParams2)
			item2, err := searchGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings(m, items2, "", vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	restyResp1, err := client.GuestSmtpNotificationConfiguration.CreateGuestSmtpNotificationSettings(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateGuestSmtpNotificationSettings", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateGuestSmtpNotificationSettings", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceGuestSmtpNotificationSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetGuestSmtpNotificationSettings")
		queryParams1 := isegosdk.GetGuestSmtpNotificationSettingsQueryParams{}

		response1, _, err := client.GuestSmtpNotificationConfiguration.GetGuestSmtpNotificationSettings(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestSmtpNotificationSettings", err,
				"Failure at GetGuestSmtpNotificationSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings(m, response1, &queryParams1)
		item1, err := searchGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings(m, items1, "", vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetGuestSmtpNotificationSettings response", err,
				"Failure when searching item from GetGuestSmtpNotificationSettings, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestSmtpNotificationSettings search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetGuestSmtpNotificationSettingsByID")
		vvID := vID

		response2, _, err := client.GuestSmtpNotificationConfiguration.GetGuestSmtpNotificationSettingsByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGuestSmtpNotificationSettingsByID", err,
				"Failure at GetGuestSmtpNotificationSettingsByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDItem(response2.ERSGuestSmtpNotificationSettings)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGuestSmtpNotificationSettingsByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceGuestSmtpNotificationSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if selectedMethod == 2 {
		vvID = vID
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestGuestSmtpNotificationSettingsUpdateGuestSmtpNotificationSettingsByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.GuestSmtpNotificationConfiguration.UpdateGuestSmtpNotificationSettingsByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateGuestSmtpNotificationSettingsByID", err, restyResp1.String(),
					"Failure at UpdateGuestSmtpNotificationSettingsByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGuestSmtpNotificationSettingsByID", err,
				"Failure at UpdateGuestSmtpNotificationSettingsByID, unexpected response", ""))
			return diags
		}
	}

	return resourceGuestSmtpNotificationSettingsRead(ctx, d, m)
}

func resourceGuestSmtpNotificationSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Function does not perform delete on ISE
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestGuestSmtpNotificationSettingsCreateGuestSmtpNotificationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestSmtpNotificationConfigurationCreateGuestSmtpNotificationSettings {
	request := isegosdk.RequestGuestSmtpNotificationConfigurationCreateGuestSmtpNotificationSettings{}
	request.ERSGuestSmtpNotificationSettings = expandRequestGuestSmtpNotificationSettingsCreateGuestSmtpNotificationSettingsERSGuestSmtpNotificationSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestSmtpNotificationSettingsCreateGuestSmtpNotificationSettingsERSGuestSmtpNotificationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestSmtpNotificationConfigurationCreateGuestSmtpNotificationSettingsERSGuestSmtpNotificationSettings {
	request := isegosdk.RequestGuestSmtpNotificationConfigurationCreateGuestSmtpNotificationSettingsERSGuestSmtpNotificationSettings{}
	if v, ok := d.GetOkExists(key + ".smtp_server"); !isEmptyValue(reflect.ValueOf(d.Get(key+".smtp_server"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".smtp_server"))) {
		request.SmtpServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".notification_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".notification_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".notification_enabled"))) {
		request.NotificationEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".use_default_from_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".use_default_from_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".use_default_from_address"))) {
		request.UseDefaultFromAddress = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".default_from_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_from_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_from_address"))) {
		request.DefaultFromAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".smtp_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".smtp_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".smtp_port"))) {
		request.SmtpPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".connection_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".connection_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".connection_timeout"))) {
		request.ConnectionTimeout = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".use_tlsor_ssl_encryption"); !isEmptyValue(reflect.ValueOf(d.Get(key+".use_tlsor_ssl_encryption"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".use_tlsor_ssl_encryption"))) {
		request.UseTLSorSSLEncryption = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".use_password_authentication"); !isEmptyValue(reflect.ValueOf(d.Get(key+".use_password_authentication"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".use_password_authentication"))) {
		request.UsePasswordAuthentication = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".user_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".user_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".user_name"))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password"))) {
		request.Password = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestSmtpNotificationSettingsUpdateGuestSmtpNotificationSettingsByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByID {
	request := isegosdk.RequestGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByID{}
	request.ERSGuestSmtpNotificationSettings = expandRequestGuestSmtpNotificationSettingsUpdateGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGuestSmtpNotificationSettingsUpdateGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings {
	request := isegosdk.RequestGuestSmtpNotificationConfigurationUpdateGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".smtp_server"); !isEmptyValue(reflect.ValueOf(d.Get(key+".smtp_server"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".smtp_server"))) {
		request.SmtpServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".notification_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".notification_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".notification_enabled"))) {
		request.NotificationEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".use_default_from_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".use_default_from_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".use_default_from_address"))) {
		request.UseDefaultFromAddress = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".default_from_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_from_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_from_address"))) {
		request.DefaultFromAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".smtp_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".smtp_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".smtp_port"))) {
		request.SmtpPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".connection_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".connection_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".connection_timeout"))) {
		request.ConnectionTimeout = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".use_tlsor_ssl_encryption"); !isEmptyValue(reflect.ValueOf(d.Get(key+".use_tlsor_ssl_encryption"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".use_tlsor_ssl_encryption"))) {
		request.UseTLSorSSLEncryption = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".use_password_authentication"); !isEmptyValue(reflect.ValueOf(d.Get(key+".use_password_authentication"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".use_password_authentication"))) {
		request.UsePasswordAuthentication = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".user_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".user_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".user_name"))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password"))) {
		request.Password = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings(m interface{}, response *isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings, queryParams *isegosdk.GetGuestSmtpNotificationSettingsQueryParams) []isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultResources
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
			response, _, err = client.GuestSmtpNotificationConfiguration.GetGuestSmtpNotificationSettings(queryParams)
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

func searchGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettings(m interface{}, items []isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsSearchResultResources, name string, id string) (*isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByIDERSGuestSmtpNotificationSettings
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseGuestSmtpNotificationConfigurationGetGuestSmtpNotificationSettingsByID
			getItem, _, err = client.GuestSmtpNotificationConfiguration.GetGuestSmtpNotificationSettingsByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetGuestSmtpNotificationSettingsByID")
			}
			foundItem = getItem.ERSGuestSmtpNotificationSettings
			return foundItem, err
		}
	}
	return foundItem, err
}
