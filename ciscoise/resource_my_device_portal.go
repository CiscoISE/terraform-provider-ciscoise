package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMyDevicePortal() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceMyDevicePortalCreate,
		ReadContext:   resourceMyDevicePortalRead,
		UpdateContext: resourceMyDevicePortalUpdate,
		DeleteContext: resourceMyDevicePortalDelete,
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

						"customizations": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_customizations": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"background_image": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"banner_image": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"banner_title": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"contact_text": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"desktop_logo_image": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"footer_element": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"mobile_logo_image": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"language": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"view_language": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"page_customizations": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"data": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"value": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"portal_theme": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"theme_data": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"portal_tweak_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"banner_color": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"banner_text_color": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"page_background_color": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"page_label_and_text_color": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"description": &schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"portal_test_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"portal_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aup_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_frequency": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"display_frequency_interval_days": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"require_scrolling": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"employee_change_password_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_employee_to_change_pwd": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"login_page_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"aup_display": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"max_failed_attempts_before_rate_limit": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"require_scrolling": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"social_configs": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"time_between_logins_during_rate_limit": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"portal_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allowed_interfaces": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"always_used_language": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"certificate_group_tag": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"display_lang": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"endpoint_identity_group": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"fallback_language": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"https_port": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"post_access_banner_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"include_post_access_banner": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"post_login_banner_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"include_post_access_banner": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"support_info_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_empty_field_value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"empty_field_display": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"include_browser_user_agent": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"include_failure_code": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"include_ip_address": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"include_mac_addr": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"include_policy_server": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"include_support_info_page": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
											},
										},
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

func resourceMyDevicePortalCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestMyDevicePortalCreateMyDevicePortal(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.MyDevicePortal.GetMyDevicePortalByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		queryParams2 := isegosdk.GetMyDevicePortalQueryParams{}

		response2, _, err := client.MyDevicePortal.GetMyDevicePortal(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsMyDevicePortalGetMyDevicePortal(m, response2, &queryParams2)
			item2, err := searchMyDevicePortalGetMyDevicePortal(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	restyResp1, err := client.MyDevicePortal.CreateMyDevicePortal(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateMyDevicePortal", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateMyDevicePortal", err))
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

func resourceMyDevicePortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		vvName := vName
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetMyDevicePortal")
		queryParams1 := isegosdk.GetMyDevicePortalQueryParams{}

		response1, _, err := client.MyDevicePortal.GetMyDevicePortal(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetMyDevicePortal", err,
				"Failure at GetMyDevicePortal, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsMyDevicePortalGetMyDevicePortal(m, response1, &queryParams1)
		item1, err := searchMyDevicePortalGetMyDevicePortal(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetMyDevicePortal response", err,
				"Failure when searching item from GetMyDevicePortal, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMyDevicePortal search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetMyDevicePortalByID")
		vvID := vID

		response2, _, err := client.MyDevicePortal.GetMyDevicePortalByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetMyDevicePortalByID", err,
				"Failure at GetMyDevicePortalByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenMyDevicePortalGetMyDevicePortalByIDItem(response2.MyDevicePortal)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMyDevicePortalByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceMyDevicePortalUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetMyDevicePortalQueryParams{}

		getResp1, _, err := client.MyDevicePortal.GetMyDevicePortal(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsMyDevicePortalGetMyDevicePortal(m, getResp1, &queryParams1)
			item1, err := searchMyDevicePortalGetMyDevicePortal(m, items1, vName, vID)
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
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestMyDevicePortalUpdateMyDevicePortalByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.MyDevicePortal.UpdateMyDevicePortalByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateMyDevicePortalByID", err, restyResp1.String(),
					"Failure at UpdateMyDevicePortalByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateMyDevicePortalByID", err,
				"Failure at UpdateMyDevicePortalByID, unexpected response", ""))
			return diags
		}
	}

	return resourceMyDevicePortalRead(ctx, d, m)
}

func resourceMyDevicePortalDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetMyDevicePortalQueryParams{}

		getResp1, _, err := client.MyDevicePortal.GetMyDevicePortal(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsMyDevicePortalGetMyDevicePortal(m, getResp1, &queryParams1)
		item1, err := searchMyDevicePortalGetMyDevicePortal(m, items1, vName, vID)
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
		getResp, _, err := client.MyDevicePortal.GetMyDevicePortalByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.MyDevicePortal.DeleteMyDevicePortalByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteMyDevicePortalByID", err, restyResp1.String(),
				"Failure at DeleteMyDevicePortalByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteMyDevicePortalByID", err,
			"Failure at DeleteMyDevicePortalByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestMyDevicePortalCreateMyDevicePortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortal {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortal{}
	request.MyDevicePortal = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortal {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortal{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_type"))) {
		request.PortalType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_test_url"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_test_url"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_test_url"))) {
		request.PortalTestURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".settings"))) {
		request.Settings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".customizations"))) {
		request.Customizations = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettings{}
	if v, ok := d.GetOkExists(key + ".portal_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_settings"))) {
		request.PortalSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".login_page_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".login_page_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".login_page_settings"))) {
		request.LoginPageSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettings(ctx, key+".login_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".aup_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aup_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aup_settings"))) {
		request.AupSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".employee_change_password_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".employee_change_password_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".employee_change_password_settings"))) {
		request.EmployeeChangePasswordSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsEmployeeChangePasswordSettings(ctx, key+".employee_change_password_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_login_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_login_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_login_banner_settings"))) {
		request.PostLoginBannerSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_access_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_access_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_access_banner_settings"))) {
		request.PostAccessBannerSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".support_info_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".support_info_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".support_info_settings"))) {
		request.SupportInfoSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPortalSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPortalSettings{}
	if v, ok := d.GetOkExists(key + ".https_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".https_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".https_port"))) {
		request.HTTPSPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allowed_interfaces"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allowed_interfaces"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allowed_interfaces"))) {
		request.AllowedInterfaces = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".certificate_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".certificate_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".certificate_group_tag"))) {
		request.CertificateGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".endpoint_identity_group"); !isEmptyValue(reflect.ValueOf(d.Get(key+".endpoint_identity_group"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".endpoint_identity_group"))) {
		request.EndpointIDentityGroup = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".display_lang"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_lang"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_lang"))) {
		request.DisplayLang = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".fallback_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".fallback_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".fallback_language"))) {
		request.FallbackLanguage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".always_used_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".always_used_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".always_used_language"))) {
		request.AlwaysUsedLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettings{}
	if v, ok := d.GetOkExists(key + ".max_failed_attempts_before_rate_limit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_failed_attempts_before_rate_limit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_failed_attempts_before_rate_limit"))) {
		request.MaxFailedAttemptsBeforeRateLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".time_between_logins_during_rate_limit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".time_between_logins_during_rate_limit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".time_between_logins_during_rate_limit"))) {
		request.TimeBetweenLoginsDuringRateLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_aup"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_aup"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_aup"))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".aup_display"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aup_display"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aup_display"))) {
		request.AupDisplay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".require_aup_acceptance"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_aup_acceptance"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_aup_acceptance"))) {
		request.RequireAupAcceptance = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".require_scrolling"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_scrolling"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_scrolling"))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".social_configs"); !isEmptyValue(reflect.ValueOf(d.Get(key+".social_configs"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".social_configs"))) {
		socialConfigs := v.([]interface{})
		request.SocialConfigs = &socialConfigs
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsAupSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(key + ".display_frequency_interval_days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_frequency_interval_days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_frequency_interval_days"))) {
		request.DisplayFrequencyIntervalDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".display_frequency"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_frequency"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_frequency"))) {
		request.DisplayFrequency = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".include_aup"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_aup"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_aup"))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".require_scrolling"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_scrolling"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_scrolling"))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsEmployeeChangePasswordSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsEmployeeChangePasswordSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsEmployeeChangePasswordSettings{}
	if v, ok := d.GetOkExists(key + ".allow_employee_to_change_pwd"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_employee_to_change_pwd"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_employee_to_change_pwd"))) {
		request.AllowEmployeeToChangePwd = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsSupportInfoSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsSupportInfoSettings{}
	if v, ok := d.GetOkExists(key + ".include_support_info_page"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_support_info_page"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_support_info_page"))) {
		request.IncludeSupportInfoPage = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_mac_addr"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_mac_addr"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_mac_addr"))) {
		request.IncludeMacAddr = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_ip_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_ip_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_ip_address"))) {
		request.IncludeIPAddress = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_browser_user_agent"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_browser_user_agent"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_browser_user_agent"))) {
		request.IncludeBrowserUserAgent = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_policy_server"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_policy_server"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_policy_server"))) {
		request.IncludePolicyServer = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_failure_code"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_failure_code"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_failure_code"))) {
		request.IncludeFailureCode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".empty_field_display"); !isEmptyValue(reflect.ValueOf(d.Get(key+".empty_field_display"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".empty_field_display"))) {
		request.EmptyFieldDisplay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_empty_field_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_empty_field_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_empty_field_value"))) {
		request.DefaultEmptyFieldValue = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizations {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizations{}
	if v, ok := d.GetOkExists(key + ".portal_theme"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_theme"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_theme"))) {
		request.PortalTheme = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(key + ".portal_tweak_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_tweak_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_tweak_settings"))) {
		request.PortalTweakSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".language"))) {
		request.Language = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(key + ".global_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".global_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".global_customizations"))) {
		request.GlobalCustomizations = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(key + ".page_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".page_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".page_customizations"))) {
		request.PageCustomizations = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTheme {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTheme{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".theme_data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".theme_data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".theme_data"))) {
		request.ThemeData = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTweakSettings{}
	if v, ok := d.GetOkExists(key + ".banner_color"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_color"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_color"))) {
		request.BannerColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".banner_text_color"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_text_color"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_text_color"))) {
		request.BannerTextColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".page_background_color"); !isEmptyValue(reflect.ValueOf(d.Get(key+".page_background_color"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".page_background_color"))) {
		request.PageBackgroundColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".page_label_and_text_color"); !isEmptyValue(reflect.ValueOf(d.Get(key+".page_label_and_text_color"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".page_label_and_text_color"))) {
		request.PageLabelAndTextColor = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsLanguage {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(key + ".view_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".view_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".view_language"))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(key + ".mobile_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mobile_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mobile_logo_image"))) {
		request.MobileLogoImage = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".desktop_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".desktop_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".desktop_logo_image"))) {
		request.DesktopLogoImage = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".banner_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_image"))) {
		request.BannerImage = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".background_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".background_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".background_image"))) {
		request.BackgroundImage = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".banner_title"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_title"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_title"))) {
		request.BannerTitle = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".contact_text"); !isEmptyValue(reflect.ValueOf(d.Get(key+".contact_text"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".contact_text"))) {
		request.ContactText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".footer_element"); !isEmptyValue(reflect.ValueOf(d.Get(key+".footer_element"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".footer_element"))) {
		request.FooterElement = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizations {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsDataArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData{}
	if v, ok := d.GetOkExists(key + ".key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key"))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByID {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByID{}
	request.MyDevicePortal = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortal {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortal{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_type"))) {
		request.PortalType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_test_url"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_test_url"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_test_url"))) {
		request.PortalTestURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".settings"))) {
		request.Settings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".customizations"))) {
		request.Customizations = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettings{}
	if v, ok := d.GetOkExists(key + ".portal_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_settings"))) {
		request.PortalSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".login_page_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".login_page_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".login_page_settings"))) {
		request.LoginPageSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings(ctx, key+".login_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".aup_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aup_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aup_settings"))) {
		request.AupSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".employee_change_password_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".employee_change_password_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".employee_change_password_settings"))) {
		request.EmployeeChangePasswordSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings(ctx, key+".employee_change_password_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_login_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_login_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_login_banner_settings"))) {
		request.PostLoginBannerSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_access_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_access_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_access_banner_settings"))) {
		request.PostAccessBannerSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".support_info_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".support_info_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".support_info_settings"))) {
		request.SupportInfoSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPortalSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPortalSettings{}
	if v, ok := d.GetOkExists(key + ".https_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".https_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".https_port"))) {
		request.HTTPSPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allowed_interfaces"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allowed_interfaces"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allowed_interfaces"))) {
		request.AllowedInterfaces = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".certificate_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".certificate_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".certificate_group_tag"))) {
		request.CertificateGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".endpoint_identity_group"); !isEmptyValue(reflect.ValueOf(d.Get(key+".endpoint_identity_group"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".endpoint_identity_group"))) {
		request.EndpointIDentityGroup = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".display_lang"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_lang"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_lang"))) {
		request.DisplayLang = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".fallback_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".fallback_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".fallback_language"))) {
		request.FallbackLanguage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".always_used_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".always_used_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".always_used_language"))) {
		request.AlwaysUsedLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings{}
	if v, ok := d.GetOkExists(key + ".max_failed_attempts_before_rate_limit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_failed_attempts_before_rate_limit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_failed_attempts_before_rate_limit"))) {
		request.MaxFailedAttemptsBeforeRateLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".time_between_logins_during_rate_limit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".time_between_logins_during_rate_limit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".time_between_logins_during_rate_limit"))) {
		request.TimeBetweenLoginsDuringRateLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_aup"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_aup"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_aup"))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".aup_display"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aup_display"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aup_display"))) {
		request.AupDisplay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".require_aup_acceptance"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_aup_acceptance"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_aup_acceptance"))) {
		request.RequireAupAcceptance = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".require_scrolling"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_scrolling"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_scrolling"))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".social_configs"); !isEmptyValue(reflect.ValueOf(d.Get(key+".social_configs"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".social_configs"))) {
		socialConfigs := v.([]interface{})
		request.SocialConfigs = &socialConfigs
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsAupSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(key + ".display_frequency_interval_days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_frequency_interval_days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_frequency_interval_days"))) {
		request.DisplayFrequencyIntervalDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".display_frequency"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_frequency"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_frequency"))) {
		request.DisplayFrequency = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".include_aup"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_aup"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_aup"))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".require_scrolling"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_scrolling"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_scrolling"))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings{}
	if v, ok := d.GetOkExists(key + ".allow_employee_to_change_pwd"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_employee_to_change_pwd"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_employee_to_change_pwd"))) {
		request.AllowEmployeeToChangePwd = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings{}
	if v, ok := d.GetOkExists(key + ".include_support_info_page"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_support_info_page"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_support_info_page"))) {
		request.IncludeSupportInfoPage = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_mac_addr"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_mac_addr"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_mac_addr"))) {
		request.IncludeMacAddr = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_ip_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_ip_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_ip_address"))) {
		request.IncludeIPAddress = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_browser_user_agent"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_browser_user_agent"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_browser_user_agent"))) {
		request.IncludeBrowserUserAgent = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_policy_server"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_policy_server"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_policy_server"))) {
		request.IncludePolicyServer = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".include_failure_code"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_failure_code"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_failure_code"))) {
		request.IncludeFailureCode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".empty_field_display"); !isEmptyValue(reflect.ValueOf(d.Get(key+".empty_field_display"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".empty_field_display"))) {
		request.EmptyFieldDisplay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_empty_field_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_empty_field_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_empty_field_value"))) {
		request.DefaultEmptyFieldValue = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizations {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizations{}
	if v, ok := d.GetOkExists(key + ".portal_theme"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_theme"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_theme"))) {
		request.PortalTheme = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(key + ".portal_tweak_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_tweak_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_tweak_settings"))) {
		request.PortalTweakSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".language"))) {
		request.Language = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(key + ".global_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".global_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".global_customizations"))) {
		request.GlobalCustomizations = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(key + ".page_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".page_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".page_customizations"))) {
		request.PageCustomizations = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".theme_data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".theme_data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".theme_data"))) {
		request.ThemeData = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings{}
	if v, ok := d.GetOkExists(key + ".banner_color"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_color"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_color"))) {
		request.BannerColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".banner_text_color"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_text_color"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_text_color"))) {
		request.BannerTextColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".page_background_color"); !isEmptyValue(reflect.ValueOf(d.Get(key+".page_background_color"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".page_background_color"))) {
		request.PageBackgroundColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".page_label_and_text_color"); !isEmptyValue(reflect.ValueOf(d.Get(key+".page_label_and_text_color"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".page_label_and_text_color"))) {
		request.PageLabelAndTextColor = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsLanguage {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(key + ".view_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".view_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".view_language"))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(key + ".mobile_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mobile_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mobile_logo_image"))) {
		request.MobileLogoImage = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".desktop_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".desktop_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".desktop_logo_image"))) {
		request.DesktopLogoImage = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".banner_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_image"))) {
		request.BannerImage = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".background_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".background_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".background_image"))) {
		request.BackgroundImage = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".banner_title"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_title"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_title"))) {
		request.BannerTitle = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".contact_text"); !isEmptyValue(reflect.ValueOf(d.Get(key+".contact_text"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".contact_text"))) {
		request.ContactText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".footer_element"); !isEmptyValue(reflect.ValueOf(d.Get(key+".footer_element"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".footer_element"))) {
		request.FooterElement = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsDataArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData{}
	if v, ok := d.GetOkExists(key + ".key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key"))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsMyDevicePortalGetMyDevicePortal(m interface{}, response *isegosdk.ResponseMyDevicePortalGetMyDevicePortal, queryParams *isegosdk.GetMyDevicePortalQueryParams) []isegosdk.ResponseMyDevicePortalGetMyDevicePortalSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseMyDevicePortalGetMyDevicePortalSearchResultResources
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
			response, _, err = client.MyDevicePortal.GetMyDevicePortal(queryParams)
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

func searchMyDevicePortalGetMyDevicePortal(m interface{}, items []isegosdk.ResponseMyDevicePortalGetMyDevicePortalSearchResultResources, name string, id string) (*isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortal, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortal
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseMyDevicePortalGetMyDevicePortalByID
			getItem, _, err = client.MyDevicePortal.GetMyDevicePortalByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetMyDevicePortalByID")
			}
			foundItem = getItem.MyDevicePortal
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseMyDevicePortalGetMyDevicePortalByID
			getItem, _, err = client.MyDevicePortal.GetMyDevicePortalByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetMyDevicePortalByID")
			}
			foundItem = getItem.MyDevicePortal
			return foundItem, err
		}
	}
	return foundItem, err
}
