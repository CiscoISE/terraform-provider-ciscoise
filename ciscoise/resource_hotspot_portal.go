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

func resourceHotspotPortal() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on HotspotPortal.

- This resource allows the client to update a hotspot portal by ID.

- This resource deletes a hotspot portal by ID.

- This resource creates a hotspot portal.
`,

		CreateContext: resourceHotspotPortalCreate,
		ReadContext:   resourceHotspotPortalRead,
		UpdateContext: resourceHotspotPortalUpdate,
		DeleteContext: resourceHotspotPortalDelete,
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
							Description: `Defines all of the Portal Customizations available`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
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
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
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
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
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
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
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
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"language": &schema.Schema{
										Description: `This property is supported only for Read operation and it allows to show the customizations in English.
Other languages are not supported`,
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
										Description: `Defines the configuration for portal theme`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `The unique internal identifier of the portal theme`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `The system- or user-assigned name of the portal theme`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"theme_data": &schema.Schema{
													Description: `A CSS file, represented as a Base64-encoded byte array`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"portal_tweak_settings": &schema.Schema{
										Description: `The Tweak Settings are a customization of the Portal Theme that has been selected for the portal.
When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme.
The Tweak Settings can subsequently be changed by the user`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"banner_color": &schema.Schema{
													Description: `Hex value of color`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
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
						"portal_test_url": &schema.Schema{
							Description: `URL to bring up a test page for this portal`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"portal_type": &schema.Schema{
							Description: `Allowed values:
- BYOD,
- HOTSPOTGUEST,
- MYDEVICE,
- SELFREGGUEST,
- SPONSOR,
- SPONSOREDGUEST`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"settings": &schema.Schema{
							Description: `Defines all of the settings groups available for a BYOD`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aup_settings": &schema.Schema{
										Description: `Configuration of the Acceptable Use Policy (AUP) for a portal`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"access_code": &schema.Schema{
													Description: `Access code that must be entered by the portal user (only valid if requireAccessCode = true)`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"include_aup": &schema.Schema{
													Description: `Require the portal user to read and accept an AUP`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"require_access_code": &schema.Schema{
													Description: `Require the portal user to enter an access code.
Only used in Hotspot portal`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"require_scrolling": &schema.Schema{
													Description: `Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
									"auth_success_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"redirect_url": &schema.Schema{
													Description: `Target URL for redirection, used when successRedirect = URL`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"success_redirect": &schema.Schema{
													Description: `After an Authentication Success where should device be redirected. Allowed values:
- AUTHSUCCESSPAGE,
- ORIGINATINGURL,
- URL`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"portal_settings": &schema.Schema{
										Description: `The port, interface, certificate, and other basic settings of a portal`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allowed_interfaces": &schema.Schema{
													Description: `Interfaces that the portal will be reachable on.
Allowed values:
- eth0
- eth1
- eth2
- eth3
- eth4
- eth5
- bond0
- bond1
- bond2`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"always_used_language": &schema.Schema{
													Description: `Used when displayLang = ALWAYSUSE`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"certificate_group_tag": &schema.Schema{
													Description: `Logical name of the x.509 server certificate that will be used for the portal`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"coa_type": &schema.Schema{
													Description: `Allowed Values:
- COAREAUTHENTICATE,
- COATERMINATE`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"display_lang": &schema.Schema{
													Description: `Allowed values:
- USEBROWSERLOCALE,
- ALWAYSUSE`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"endpoint_identity_group": &schema.Schema{
													Description: `Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"fallback_language": &schema.Schema{
													Description: `Used when displayLang = USEBROWSERLOCALE`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"https_port": &schema.Schema{
													Description: `The port number that the allowed interfaces will listen on.
Range from 8000 to 8999`,
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
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
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
													Description: `Include a Post-Login Banner page`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
									"support_info_settings": &schema.Schema{
										Description: `Portal Support Information Settings`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_empty_field_value": &schema.Schema{
													Description: `The default value displayed for an empty field.
Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"empty_field_display": &schema.Schema{
													Description: `Specifies how empty fields are handled on the Support Information Page.
Allowed values:
- HIDE,
- DISPLAYWITHNOVALUE,
- DISPLAYWITHDEFAULTVALUE`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"include_browser_user_agent": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"include_failure_code": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"include_ip_address": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"include_mac_addr": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"include_policy_server": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"include_support_info_page": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
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

func resourceHotspotPortalCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestHotspotPortalCreateHotspotPortal(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.HotspotPortal.GetHotspotPortalByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		queryParams2 := isegosdk.GetHotspotPortalQueryParams{}

		response2, _, err := client.HotspotPortal.GetHotspotPortal(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsHotspotPortalGetHotspotPortal(m, response2, &queryParams2)
			item2, err := searchHotspotPortalGetHotspotPortal(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	restyResp1, err := client.HotspotPortal.CreateHotspotPortal(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateHotspotPortal", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateHotspotPortal", err))
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

func resourceHotspotPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetHotspotPortal")
		queryParams1 := isegosdk.GetHotspotPortalQueryParams{}

		response1, _, err := client.HotspotPortal.GetHotspotPortal(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetHotspotPortal", err,
				"Failure at GetHotspotPortal, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsHotspotPortalGetHotspotPortal(m, response1, &queryParams1)
		item1, err := searchHotspotPortalGetHotspotPortal(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetHotspotPortal response", err,
				"Failure when searching item from GetHotspotPortal, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetHotspotPortal search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetHotspotPortalByID")
		vvID := vID

		response2, _, err := client.HotspotPortal.GetHotspotPortalByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetHotspotPortalByID", err,
				"Failure at GetHotspotPortalByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenHotspotPortalGetHotspotPortalByIDItem(response2.HotspotPortal)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetHotspotPortalByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceHotspotPortalUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		queryParams1 := isegosdk.GetHotspotPortalQueryParams{}
		getResp1, _, err := client.HotspotPortal.GetHotspotPortal(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsHotspotPortalGetHotspotPortal(m, getResp1, &queryParams1)
			item1, err := searchHotspotPortalGetHotspotPortal(m, items1, vName, vID)
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
		request1 := expandRequestHotspotPortalUpdateHotspotPortalByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.HotspotPortal.UpdateHotspotPortalByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateHotspotPortalByID", err, restyResp1.String(),
					"Failure at UpdateHotspotPortalByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateHotspotPortalByID", err,
				"Failure at UpdateHotspotPortalByID, unexpected response", ""))
			return diags
		}
	}

	return resourceHotspotPortalRead(ctx, d, m)
}

func resourceHotspotPortalDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		queryParams1 := isegosdk.GetHotspotPortalQueryParams{}

		getResp1, _, err := client.HotspotPortal.GetHotspotPortal(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsHotspotPortalGetHotspotPortal(m, getResp1, &queryParams1)
		item1, err := searchHotspotPortalGetHotspotPortal(m, items1, vName, vID)
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
		getResp, _, err := client.HotspotPortal.GetHotspotPortalByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.HotspotPortal.DeleteHotspotPortalByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteHotspotPortalByID", err, restyResp1.String(),
				"Failure at DeleteHotspotPortalByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteHotspotPortalByID", err,
			"Failure at DeleteHotspotPortalByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestHotspotPortalCreateHotspotPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortal {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortal{}
	request.HotspotPortal = expandRequestHotspotPortalCreateHotspotPortalHotspotPortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortal {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortal{}
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
		request.Settings = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".customizations"))) {
		request.Customizations = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettings {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettings{}
	if v, ok := d.GetOkExists(key + ".portal_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_settings"))) {
		request.PortalSettings = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".aup_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aup_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aup_settings"))) {
		request.AupSettings = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_access_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_access_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_access_banner_settings"))) {
		request.PostAccessBannerSettings = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".auth_success_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".auth_success_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".auth_success_settings"))) {
		request.AuthSuccessSettings = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAuthSuccessSettings(ctx, key+".auth_success_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_login_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_login_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_login_banner_settings"))) {
		request.PostLoginBannerSettings = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".support_info_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".support_info_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".support_info_settings"))) {
		request.SupportInfoSettings = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPortalSettings {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPortalSettings{}
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
	if v, ok := d.GetOkExists(key + ".coa_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".coa_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".coa_type"))) {
		request.CoaType = interfaceToString(v)
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

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAupSettings {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(key + ".require_access_code"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_access_code"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_access_code"))) {
		request.RequireAccessCode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".access_code"); !isEmptyValue(reflect.ValueOf(d.Get(key+".access_code"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".access_code"))) {
		request.AccessCode = interfaceToString(v)
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

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAuthSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAuthSuccessSettings {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsAuthSuccessSettings{}
	if v, ok := d.GetOkExists(key + ".success_redirect"); !isEmptyValue(reflect.ValueOf(d.Get(key+".success_redirect"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".success_redirect"))) {
		request.SuccessRedirect = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".redirect_url"); !isEmptyValue(reflect.ValueOf(d.Get(key+".redirect_url"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".redirect_url"))) {
		request.RedirectURL = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsSupportInfoSettings {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalSettingsSupportInfoSettings{}
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

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizations {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizations{}
	if v, ok := d.GetOkExists(key + ".portal_theme"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_theme"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_theme"))) {
		request.PortalTheme = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(key + ".portal_tweak_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_tweak_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_tweak_settings"))) {
		request.PortalTweakSettings = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".language"))) {
		request.Language = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(key + ".global_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".global_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".global_customizations"))) {
		request.GlobalCustomizations = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(key + ".page_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".page_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".page_customizations"))) {
		request.PageCustomizations = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTheme {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTheme{}
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

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPortalTweakSettings{}
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

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsLanguage {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(key + ".view_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".view_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".view_language"))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(key + ".mobile_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mobile_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mobile_logo_image"))) {
		request.MobileLogoImage = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".desktop_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".desktop_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".desktop_logo_image"))) {
		request.DesktopLogoImage = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".background_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".background_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".background_image"))) {
		request.BackgroundImage = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".banner_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_image"))) {
		request.BannerImage = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
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

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizations {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizationsDataArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizationsData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestHotspotPortalCreateHotspotPortalHotspotPortalCustomizationsPageCustomizationsData{}
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

func expandRequestHotspotPortalUpdateHotspotPortalByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByID {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByID{}
	request.HotspotPortal = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortal {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortal{}
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
		request.Settings = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".customizations"))) {
		request.Customizations = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettings {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettings{}
	if v, ok := d.GetOkExists(key + ".portal_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_settings"))) {
		request.PortalSettings = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".aup_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aup_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aup_settings"))) {
		request.AupSettings = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_access_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_access_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_access_banner_settings"))) {
		request.PostAccessBannerSettings = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".auth_success_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".auth_success_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".auth_success_settings"))) {
		request.AuthSuccessSettings = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAuthSuccessSettings(ctx, key+".auth_success_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".post_login_banner_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".post_login_banner_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".post_login_banner_settings"))) {
		request.PostLoginBannerSettings = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".support_info_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".support_info_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".support_info_settings"))) {
		request.SupportInfoSettings = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPortalSettings {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPortalSettings{}
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
	if v, ok := d.GetOkExists(key + ".coa_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".coa_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".coa_type"))) {
		request.CoaType = interfaceToString(v)
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

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAupSettings {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(key + ".require_access_code"); !isEmptyValue(reflect.ValueOf(d.Get(key+".require_access_code"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".require_access_code"))) {
		request.RequireAccessCode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".access_code"); !isEmptyValue(reflect.ValueOf(d.Get(key+".access_code"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".access_code"))) {
		request.AccessCode = interfaceToString(v)
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

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAuthSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAuthSuccessSettings {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsAuthSuccessSettings{}
	if v, ok := d.GetOkExists(key + ".success_redirect"); !isEmptyValue(reflect.ValueOf(d.Get(key+".success_redirect"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".success_redirect"))) {
		request.SuccessRedirect = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".redirect_url"); !isEmptyValue(reflect.ValueOf(d.Get(key+".redirect_url"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".redirect_url"))) {
		request.RedirectURL = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(key + ".include_post_access_banner"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_post_access_banner"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_post_access_banner"))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsSupportInfoSettings {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalSettingsSupportInfoSettings{}
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

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizations {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizations{}
	if v, ok := d.GetOkExists(key + ".portal_theme"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_theme"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_theme"))) {
		request.PortalTheme = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(key + ".portal_tweak_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_tweak_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_tweak_settings"))) {
		request.PortalTweakSettings = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".language"))) {
		request.Language = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(key + ".global_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".global_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".global_customizations"))) {
		request.GlobalCustomizations = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(key + ".page_customizations"); !isEmptyValue(reflect.ValueOf(d.Get(key+".page_customizations"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".page_customizations"))) {
		request.PageCustomizations = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTheme {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTheme{}
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

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPortalTweakSettings{}
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

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsLanguage {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(key + ".view_language"); !isEmptyValue(reflect.ValueOf(d.Get(key+".view_language"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".view_language"))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(key + ".mobile_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mobile_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mobile_logo_image"))) {
		request.MobileLogoImage = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".desktop_logo_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".desktop_logo_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".desktop_logo_image"))) {
		request.DesktopLogoImage = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".background_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".background_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".background_image"))) {
		request.BackgroundImage = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
	}
	if v, ok := d.GetOkExists(key + ".banner_image"); !isEmptyValue(reflect.ValueOf(d.Get(key+".banner_image"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".banner_image"))) {
		request.BannerImage = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
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

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizations {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsDataArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestHotspotPortalUpdateHotspotPortalByIDHotspotPortalCustomizationsPageCustomizationsData{}
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

func getAllItemsHotspotPortalGetHotspotPortal(m interface{}, response *isegosdk.ResponseHotspotPortalGetHotspotPortal, queryParams *isegosdk.GetHotspotPortalQueryParams) []isegosdk.ResponseHotspotPortalGetHotspotPortalSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseHotspotPortalGetHotspotPortalSearchResultResources
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
			response, _, err = client.HotspotPortal.GetHotspotPortal(queryParams)
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

func searchHotspotPortalGetHotspotPortal(m interface{}, items []isegosdk.ResponseHotspotPortalGetHotspotPortalSearchResultResources, name string, id string) (*isegosdk.ResponseHotspotPortalGetHotspotPortalByIDHotspotPortal, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseHotspotPortalGetHotspotPortalByIDHotspotPortal
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseHotspotPortalGetHotspotPortalByID
			getItem, _, err = client.HotspotPortal.GetHotspotPortalByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetHotspotPortalByID")
			}
			foundItem = getItem.HotspotPortal
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseHotspotPortalGetHotspotPortalByID
			getItem, _, err = client.HotspotPortal.GetHotspotPortalByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetHotspotPortalByID")
			}
			foundItem = getItem.HotspotPortal
			return foundItem, err
		}
	}
	return foundItem, err
}
