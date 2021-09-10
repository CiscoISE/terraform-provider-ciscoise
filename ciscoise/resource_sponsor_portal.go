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

func resourceSponsorPortal() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSponsorPortalCreate,
		ReadContext:   resourceSponsorPortalRead,
		UpdateContext: resourceSponsorPortalUpdate,
		DeleteContext: resourceSponsorPortalDelete,
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
												"require_aup_scrolling": &schema.Schema{
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
												"authentication_method": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"available_ssids": &schema.Schema{
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
												"fallback_language": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"fqdn": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"https_port": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"idle_timeout": &schema.Schema{
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
									"sponsor_change_password_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_sponsor_to_change_pwd": &schema.Schema{
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

func resourceSponsorPortalCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestSponsorPortalCreateSponsorPortal(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.SponsorPortal.GetSponsorPortalByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		queryParams2 := isegosdk.GetSponsorPortalQueryParams{}

		response2, _, err := client.SponsorPortal.GetSponsorPortal(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsSponsorPortalGetSponsorPortal(m, response2, &queryParams2)
			item2, err := searchSponsorPortalGetSponsorPortal(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	restyResp1, err := client.SponsorPortal.CreateSponsorPortal(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSponsorPortal", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSponsorPortal", err))
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

func resourceSponsorPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetSponsorPortal")
		queryParams1 := isegosdk.GetSponsorPortalQueryParams{}

		response1, _, err := client.SponsorPortal.GetSponsorPortal(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSponsorPortal", err,
				"Failure at GetSponsorPortal, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsSponsorPortalGetSponsorPortal(m, response1, &queryParams1)
		item1, err := searchSponsorPortalGetSponsorPortal(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetSponsorPortal response", err,
				"Failure when searching item from GetSponsorPortal, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsorPortal search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSponsorPortalByID")
		vvID := vID

		response2, _, err := client.SponsorPortal.GetSponsorPortalByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSponsorPortalByID", err,
				"Failure at GetSponsorPortalByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenSponsorPortalGetSponsorPortalByIDItem(response2.SponsorPortal)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsorPortalByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSponsorPortalUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		queryParams1 := isegosdk.GetSponsorPortalQueryParams{}

		getResp1, _, err := client.SponsorPortal.GetSponsorPortal(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsSponsorPortalGetSponsorPortal(m, getResp1, &queryParams1)
			item1, err := searchSponsorPortalGetSponsorPortal(m, items1, vName, vID)
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
		request1 := expandRequestSponsorPortalUpdateSponsorPortalByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SponsorPortal.UpdateSponsorPortalByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSponsorPortalByID", err, restyResp1.String(),
					"Failure at UpdateSponsorPortalByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSponsorPortalByID", err,
				"Failure at UpdateSponsorPortalByID, unexpected response", ""))
			return diags
		}
	}

	return resourceSponsorPortalRead(ctx, d, m)
}

func resourceSponsorPortalDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		queryParams1 := isegosdk.GetSponsorPortalQueryParams{}

		getResp1, _, err := client.SponsorPortal.GetSponsorPortal(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSponsorPortalGetSponsorPortal(m, getResp1, &queryParams1)
		item1, err := searchSponsorPortalGetSponsorPortal(m, items1, vName, vID)
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
		getResp, _, err := client.SponsorPortal.GetSponsorPortalByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.SponsorPortal.DeleteSponsorPortalByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSponsorPortalByID", err, restyResp1.String(),
				"Failure at DeleteSponsorPortalByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSponsorPortalByID", err,
			"Failure at DeleteSponsorPortalByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSponsorPortalCreateSponsorPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortal {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortal{}
	request.SponsorPortal = expandRequestSponsorPortalCreateSponsorPortalSponsorPortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortal {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortal{}
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
		request.Settings = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".customizations"))) {
		request.Customizations = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettings {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettings{}
	if v, ok := d.GetOkExists(key + ".portal_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_settings"))) {
		request.PortalSettings = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".login_page_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".login_page_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".login_page_settings"))) {
		request.LoginPageSettings = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsLoginPageSettings(ctx, key+".login_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".aup_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aup_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aup_settings"))) {
		request.AupSettings = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".sponsor_change_password_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sponsor_change_password_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sponsor_change_password_settings"))) {
		request.SponsorChangePasswordSettings = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSponsorChangePasswordSettings(ctx, key+".sponsor_change_password_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_login_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_login_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_login_banner_settings"))) {
		request.PostLoginBannerSettings = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_access_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_access_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_access_banner_settings"))) {
		request.PostAccessBannerSettings = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".support_info_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".support_info_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".support_info_settings"))) {
		request.SupportInfoSettings = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPortalSettings {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPortalSettings{}
	if v, ok := d.GetOkExists(key + ".https_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".https_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".https_port"))) {
		request.HTTPSPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allowed_interfaces"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allowed_interfaces"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allowed_interfaces"))) {
		request.AllowedInterfaces = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".certificate_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".certificate_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".certificate_group_tag"))) {
		request.CertificateGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".fqdn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".fqdn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".fqdn"))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".authentication_method"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authentication_method"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authentication_method"))) {
		request.AuthenticationMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".idle_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".idle_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".idle_timeout"))) {
		request.IDleTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".display_lang"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_lang"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_lang"))) {
		request.DisplayLang = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".fallback_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".fallback_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".fallback_language"))) {
		request.FallbackLanguage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".available_ssids"); !isEmptyValue(reflect.ValueOf(d.Get(key+".available_ssids"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".available_ssids"))) {
		request.AvailableSSIDs = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsLoginPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsLoginPageSettings {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsLoginPageSettings{}
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
	if v, ok := d.GetOkExists(key + ".require_aup_scrolling"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_aup_scrolling"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_aup_scrolling"))) {
		request.RequireAupScrolling = interfaceToBoolPtr(v)
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

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsAupSettings {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(key + ".include_aup"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_aup"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_aup"))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".require_scrolling"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_scrolling"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_scrolling"))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".display_frequency"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_frequency"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_frequency"))) {
		request.DisplayFrequency = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".display_frequency_interval_days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_frequency_interval_days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_frequency_interval_days"))) {
		request.DisplayFrequencyIntervalDays = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSponsorChangePasswordSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSponsorChangePasswordSettings {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSponsorChangePasswordSettings{}
	if v, ok := d.GetOkExists(key + ".allow_sponsor_to_change_pwd"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_sponsor_to_change_pwd"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_sponsor_to_change_pwd"))) {
		request.AllowSponsorToChangePwd = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSupportInfoSettings {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalSettingsSupportInfoSettings{}
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

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizations {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizations{}
	if v, ok := d.GetOkExists(key + ".portal_theme"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_theme"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_theme"))) {
		request.PortalTheme = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(key + ".portal_tweak_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_tweak_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_tweak_settings"))) {
		request.PortalTweakSettings = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".language"))) {
		request.Language = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(key + ".global_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".global_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".global_customizations"))) {
		request.GlobalCustomizations = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(key + ".page_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".page_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".page_customizations"))) {
		request.PageCustomizations = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTheme {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTheme{}
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

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPortalTweakSettings{}
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

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsLanguage {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(key + ".view_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".view_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".view_language"))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(key + ".mobile_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mobile_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mobile_logo_image"))) {
		request.MobileLogoImage = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".desktop_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".desktop_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".desktop_logo_image"))) {
		request.DesktopLogoImage = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".banner_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_image"))) {
		request.BannerImage = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".background_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".background_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".background_image"))) {
		request.BackgroundImage = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
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

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizations {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizationsDataArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizationsData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestSponsorPortalCreateSponsorPortalSponsorPortalCustomizationsPageCustomizationsData{}
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

func expandRequestSponsorPortalUpdateSponsorPortalByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByID {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByID{}
	request.SponsorPortal = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortal {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortal{}
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
		request.Settings = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".customizations"))) {
		request.Customizations = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettings {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettings{}
	if v, ok := d.GetOkExists(key + ".portal_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_settings"))) {
		request.PortalSettings = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".login_page_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".login_page_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".login_page_settings"))) {
		request.LoginPageSettings = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsLoginPageSettings(ctx, key+".login_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".aup_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aup_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aup_settings"))) {
		request.AupSettings = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".sponsor_change_password_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sponsor_change_password_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sponsor_change_password_settings"))) {
		request.SponsorChangePasswordSettings = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSponsorChangePasswordSettings(ctx, key+".sponsor_change_password_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_login_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_login_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_login_banner_settings"))) {
		request.PostLoginBannerSettings = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_access_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_access_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_access_banner_settings"))) {
		request.PostAccessBannerSettings = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".support_info_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".support_info_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".support_info_settings"))) {
		request.SupportInfoSettings = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPortalSettings {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPortalSettings{}
	if v, ok := d.GetOkExists(key + ".https_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".https_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".https_port"))) {
		request.HTTPSPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allowed_interfaces"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allowed_interfaces"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allowed_interfaces"))) {
		request.AllowedInterfaces = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".certificate_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".certificate_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".certificate_group_tag"))) {
		request.CertificateGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".fqdn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".fqdn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".fqdn"))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".authentication_method"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authentication_method"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authentication_method"))) {
		request.AuthenticationMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".idle_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".idle_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".idle_timeout"))) {
		request.IDleTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".display_lang"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_lang"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_lang"))) {
		request.DisplayLang = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".fallback_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".fallback_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".fallback_language"))) {
		request.FallbackLanguage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".available_ssids"); !isEmptyValue(reflect.ValueOf(d.Get(key+".available_ssids"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".available_ssids"))) {
		request.AvailableSSIDs = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsLoginPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsLoginPageSettings {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsLoginPageSettings{}
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
	if v, ok := d.GetOkExists(key + ".require_aup_scrolling"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_aup_scrolling"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_aup_scrolling"))) {
		request.RequireAupScrolling = interfaceToBoolPtr(v)
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

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsAupSettings {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(key + ".include_aup"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_aup"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_aup"))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".require_scrolling"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_scrolling"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_scrolling"))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".display_frequency"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_frequency"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_frequency"))) {
		request.DisplayFrequency = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".display_frequency_interval_days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_frequency_interval_days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_frequency_interval_days"))) {
		request.DisplayFrequencyIntervalDays = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSponsorChangePasswordSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSponsorChangePasswordSettings {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSponsorChangePasswordSettings{}
	if v, ok := d.GetOkExists(key + ".allow_sponsor_to_change_pwd"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_sponsor_to_change_pwd"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_sponsor_to_change_pwd"))) {
		request.AllowSponsorToChangePwd = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSupportInfoSettings {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalSettingsSupportInfoSettings{}
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

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizations {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizations{}
	if v, ok := d.GetOkExists(key + ".portal_theme"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_theme"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_theme"))) {
		request.PortalTheme = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(key + ".portal_tweak_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_tweak_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_tweak_settings"))) {
		request.PortalTweakSettings = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".language"))) {
		request.Language = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(key + ".global_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".global_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".global_customizations"))) {
		request.GlobalCustomizations = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(key + ".page_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".page_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".page_customizations"))) {
		request.PageCustomizations = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTheme {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTheme{}
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

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPortalTweakSettings{}
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

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsLanguage {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(key + ".view_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".view_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".view_language"))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(key + ".mobile_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mobile_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mobile_logo_image"))) {
		request.MobileLogoImage = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".desktop_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".desktop_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".desktop_logo_image"))) {
		request.DesktopLogoImage = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".banner_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_image"))) {
		request.BannerImage = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".background_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".background_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".background_image"))) {
		request.BackgroundImage = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
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

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizations {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsDataArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestSponsorPortalUpdateSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData{}
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

func getAllItemsSponsorPortalGetSponsorPortal(m interface{}, response *isegosdk.ResponseSponsorPortalGetSponsorPortal, queryParams *isegosdk.GetSponsorPortalQueryParams) []isegosdk.ResponseSponsorPortalGetSponsorPortalSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseSponsorPortalGetSponsorPortalSearchResultResources
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
			response, _, err = client.SponsorPortal.GetSponsorPortal(queryParams)
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

func searchSponsorPortalGetSponsorPortal(m interface{}, items []isegosdk.ResponseSponsorPortalGetSponsorPortalSearchResultResources, name string, id string) (*isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortal, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortal
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSponsorPortalGetSponsorPortalByID
			getItem, _, err = client.SponsorPortal.GetSponsorPortalByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSponsorPortalByID")
			}
			foundItem = getItem.SponsorPortal
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSponsorPortalGetSponsorPortalByID
			getItem, _, err = client.SponsorPortal.GetSponsorPortalByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSponsorPortalByID")
			}
			foundItem = getItem.SponsorPortal
			return foundItem, err
		}
	}
	return foundItem, err
}
