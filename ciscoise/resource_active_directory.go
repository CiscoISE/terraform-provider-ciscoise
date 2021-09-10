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

func resourceActiveDirectory() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceActiveDirectoryCreate,
		ReadContext:   resourceActiveDirectoryRead,
		UpdateContext: resourceActiveDirectoryUpdate,
		DeleteContext: resourceActiveDirectoryDelete,
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

						"ad_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"internal_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"type": &schema.Schema{
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
						"ad_scopes_names": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"adgroups": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"groups": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"sid": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"type": &schema.Schema{
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
						"advanced_settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aging_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"auth_protection_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"country": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"department": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"email": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"enable_callback_for_dialin_client": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"enable_dialin_permission_check": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"enable_failed_auth_protection": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"enable_machine_access": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"enable_machine_auth": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"enable_pass_change": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"enable_rewrites": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"failed_auth_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"identity_not_in_ad_behaviour": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"job_title": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"locality": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"organizational_unit": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"plaintext_auth": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"rewrite_rules": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"rewrite_match": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"rewrite_result": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"row_id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"schema": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"state_or_province": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"street_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"telephone": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"unreachable_domains_behaviour": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enable_domain_allowed_list": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enable_domain_white_list": &schema.Schema{
							Type:     schema.TypeBool,
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
					},
				},
			},
		},
	}
}

func resourceActiveDirectoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestActiveDirectoryCreateActiveDirectory(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	// vDomain, okDomain := resourceItem["domain"]
	// vvDomain := interfaceToString(vDomain)

	if okID && vvID != "" {
		getResponse1, _, err := client.ActiveDirectory.GetActiveDirectoryByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			// resourceMap["domain"] = vvDomain
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.ActiveDirectory.GetActiveDirectoryByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			// resourceMap["domain"] = vvDomain
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	// if okDomain && vvDomain != "" {
	// 	getResponse2, _, err := client.ActiveDirectory.GetActiveDirectoryByName(vvDomain)
	// 	if err == nil && getResponse2 != nil {
	// 		resourceMap := make(map[string]string)
	// 		resourceMap["id"] = vvID
	// 		resourceMap["name"] = vvName
	// 		resourceMap["domain"] = vvDomain
	// 		d.SetId(joinResourceID(resourceMap))
	// 		return diags
	// 	}
	// }
	restyResp1, err := client.ActiveDirectory.CreateActiveDirectory(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateActiveDirectory", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateActiveDirectory", err))
		return diags
	}
	// REVIEW: if location header contains id
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	// NOTE: attr name is part of creation, now is part of the Tf resource identifier
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	// resourceMap["domain"] = vvDomain
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceActiveDirectoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	// Changed order of selection to give priority to id
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetActiveDirectoryByName")
		vvName := vName

		response1, _, err := client.ActiveDirectory.GetActiveDirectoryByName(vvName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetActiveDirectoryByName", err,
				"Failure at GetActiveDirectoryByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItemName1 := flattenActiveDirectoryGetActiveDirectoryByNameItemName(response1.ERSActiveDirectory)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetActiveDirectoryByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetActiveDirectoryByID")
		vvID := vID

		response2, _, err := client.ActiveDirectory.GetActiveDirectoryByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetActiveDirectoryByID", err,
				"Failure at GetActiveDirectoryByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemID2 := flattenActiveDirectoryGetActiveDirectoryByIDItemID(response2.ERSActiveDirectory)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetActiveDirectoryByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceActiveDirectoryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceActiveDirectoryRead(ctx, d, m)
}

func resourceActiveDirectoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	// Changed order of selection to give priority to id
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.ActiveDirectory.GetActiveDirectoryByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.ActiveDirectory.GetActiveDirectoryByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.ERSActiveDirectory != nil {
			vvID = getResp.ERSActiveDirectory.ID
		}
	}
	restyResp1, err := client.ActiveDirectory.DeleteActiveDirectoryByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteActiveDirectoryByID", err, restyResp1.String(),
				"Failure at DeleteActiveDirectoryByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteActiveDirectoryByID", err,
			"Failure at DeleteActiveDirectoryByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestActiveDirectoryCreateActiveDirectory(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectory {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectory{}
	request.ERSActiveDirectory = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectory(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectory(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectory {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectory{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".domain"); !isEmptyValue(reflect.ValueOf(d.Get(key+".domain"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".domain"))) {
		request.Domain = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_domain_white_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_domain_white_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_domain_white_list"))) {
		request.EnableDomainWhiteList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".adgroups"); !isEmptyValue(reflect.ValueOf(d.Get(key+".adgroups"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".adgroups"))) {
		request.Adgroups = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroups(ctx, key+".adgroups.0", d)
	}
	if v, ok := d.GetOkExists(key + ".advanced_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advanced_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advanced_settings"))) {
		request.AdvancedSettings = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettings(ctx, key+".advanced_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".ad_attributes"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ad_attributes"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ad_attributes"))) {
		request.AdAttributes = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributes(ctx, key+".ad_attributes.0", d)
	}
	if v, ok := d.GetOkExists(key + ".ad_scopes_names"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ad_scopes_names"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ad_scopes_names"))) {
		request.AdScopesNames = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroups {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroups{}
	if v, ok := d.GetOkExists(key + ".groups"); !isEmptyValue(reflect.ValueOf(d.Get(key+".groups"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".groups"))) {
		request.Groups = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroupsArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroupsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups {
	request := []isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sid"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sid"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sid"))) {
		request.Sid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettings {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettings{}
	if v, ok := d.GetOkExists(key + ".enable_pass_change"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_pass_change"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_pass_change"))) {
		request.EnablePassChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_machine_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_machine_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_machine_auth"))) {
		request.EnableMachineAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_machine_access"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_machine_access"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_machine_access"))) {
		request.EnableMachineAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".aging_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aging_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aging_time"))) {
		request.AgingTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_dialin_permission_check"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_dialin_permission_check"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_dialin_permission_check"))) {
		request.EnableDialinPermissionCheck = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_callback_for_dialin_client"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_callback_for_dialin_client"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_callback_for_dialin_client"))) {
		request.EnableCallbackForDialinClient = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".plaintext_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".plaintext_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".plaintext_auth"))) {
		request.PlaintextAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_failed_auth_protection"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_failed_auth_protection"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_failed_auth_protection"))) {
		request.EnableFailedAuthProtection = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".auth_protection_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".auth_protection_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".auth_protection_type"))) {
		request.AuthProtectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".failed_auth_threshold"); !isEmptyValue(reflect.ValueOf(d.Get(key+".failed_auth_threshold"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".failed_auth_threshold"))) {
		request.FailedAuthThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".identity_not_in_ad_behaviour"); !isEmptyValue(reflect.ValueOf(d.Get(key+".identity_not_in_ad_behaviour"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".identity_not_in_ad_behaviour"))) {
		request.IDentityNotInAdBehaviour = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".unreachable_domains_behaviour"); !isEmptyValue(reflect.ValueOf(d.Get(key+".unreachable_domains_behaviour"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".unreachable_domains_behaviour"))) {
		request.UnreachableDomainsBehaviour = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_rewrites"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_rewrites"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_rewrites"))) {
		request.EnableRewrites = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".rewrite_rules"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rewrite_rules"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rewrite_rules"))) {
		request.RewriteRules = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRulesArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".first_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".first_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".first_name"))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".department"); !isEmptyValue(reflect.ValueOf(d.Get(key+".department"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".department"))) {
		request.Department = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".last_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".last_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".last_name"))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".organizational_unit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".organizational_unit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".organizational_unit"))) {
		request.OrganizationalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".job_title"); !isEmptyValue(reflect.ValueOf(d.Get(key+".job_title"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".job_title"))) {
		request.JobTitle = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".locality"); !isEmptyValue(reflect.ValueOf(d.Get(key+".locality"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".locality"))) {
		request.Locality = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".email"); !isEmptyValue(reflect.ValueOf(d.Get(key+".email"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".email"))) {
		request.Email = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".state_or_province"); !isEmptyValue(reflect.ValueOf(d.Get(key+".state_or_province"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".state_or_province"))) {
		request.StateOrProvince = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".telephone"); !isEmptyValue(reflect.ValueOf(d.Get(key+".telephone"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".telephone"))) {
		request.Telephone = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".country"); !isEmptyValue(reflect.ValueOf(d.Get(key+".country"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".country"))) {
		request.Country = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".street_address"); !isEmptyValue(reflect.ValueOf(d.Get(key+".street_address"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".street_address"))) {
		request.StreetAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".schema"); !isEmptyValue(reflect.ValueOf(d.Get(key+".schema"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".schema"))) {
		request.Schema = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules {
	request := []isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules{}
	if v, ok := d.GetOkExists(key + ".row_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".row_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".row_id"))) {
		request.RowID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".rewrite_match"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rewrite_match"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rewrite_match"))) {
		request.RewriteMatch = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rewrite_result"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rewrite_result"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rewrite_result"))) {
		request.RewriteResult = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributes {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributes{}
	if v, ok := d.GetOkExists(key + ".attributes"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attributes"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attributes"))) {
		request.Attributes = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributesArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes {
	request := []isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".internal_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".internal_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".internal_name"))) {
		request.InternalName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_value"))) {
		request.DefaultValue = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
