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

func resourceActiveDirectory() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update* and delete operations on ActiveDirectory.

- This resource deletes an AD join point from Cisco ISE.

- This resource creates an AD join point in Cisco ISE.

- *This resource action loads domain groups configuration from Active Directory into Cisco ISE.
`,

		CreateContext: resourceActiveDirectoryCreate,
		ReadContext:   resourceActiveDirectoryRead,
		UpdateContext: resourceActiveDirectoryUpdate,
		DeleteContext: resourceActiveDirectoryDelete,
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

						"ad_attributes": &schema.Schema{
							Description: `Holds list of AD Attributes`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Description: `List of Attributes`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_value": &schema.Schema{
													Description: `Required for each attribute in the attribute list. Can contain an empty string. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"internal_name": &schema.Schema{
													Description: `Required for each attribute in the attribute list. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Required for each attribute in the attribute list with no duplication between attributes. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Required for each group in the group list. Allowed values: STRING, IP, BOOLEAN, INT, OCTET_STRING`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"ad_scopes_names": &schema.Schema{
							Description: `String that contains the names of the scopes that the active directory belongs to. Names are separated by comma. Alphanumeric, underscore (_) characters are allowed`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"adgroups": &schema.Schema{
							Description: `Holds list of AD Groups`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"groups": &schema.Schema{
										Description: `List of Groups`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Required for each group in the group list with no duplication between groups. All characters are allowed except %`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"sid": &schema.Schema{
													Description: `Cisco ISE uses security identifiers (SIDs) for optimization of group membership evaluation. SIDs are useful for efficiency (speed) when the groups are evaluated. All characters are allowed except %`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `No character restriction`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"advanced_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aging_time": &schema.Schema{
										Description: `Range 1-8760 hours`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"auth_protection_type": &schema.Schema{
										Description: `Enable prevent AD account lockout. Allowed values:
- WIRELESS,
- WIRED,
- BOTH`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"country": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"department": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"email": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"enable_callback_for_dialin_client": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_dialin_permission_check": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_failed_auth_protection": &schema.Schema{
										Description: `Enable prevent AD account lockout due to too many bad password attempts`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"enable_machine_access": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_machine_auth": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_pass_change": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_rewrites": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"failed_auth_threshold": &schema.Schema{
										Description: `Number of bad password attempts`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"first_name": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"identity_not_in_ad_behaviour": &schema.Schema{
										Description: `Allowed values: REJECT, SEARCH_JOINED_FOREST, SEARCH_ALL`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"job_title": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"last_name": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"locality": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"organizational_unit": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"plaintext_auth": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rewrite_rules": &schema.Schema{
										Description: `Identity rewrite is an advanced feature that directs Cisco ISE to manipulate the identity
before it is passed to the external Active Directory system. You can create rules to change
the identity to a desired format that includes or excludes a domain prefix and/or suffix or
other additional markup of your choice`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"rewrite_match": &schema.Schema{
													Description: `Required for each rule in the list with no duplication between rules. All characters are allowed except %"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"rewrite_result": &schema.Schema{
													Description: `Required for each rule in the list. All characters are allowed except %"`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"row_id": &schema.Schema{
													Description: `Required for each rule in the list in serial order`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
											},
										},
									},
									"schema": &schema.Schema{
										Description: `Allowed values: ACTIVE_DIRECTORY, CUSTOM.
Choose ACTIVE_DIRECTORY schema when the AD attributes defined in AD can be copied to relevant attributes
in Cisco ISE. If customization is needed, choose CUSTOM schema. All User info attributes are always set to
default value if schema is ACTIVE_DIRECTORY. Values can be changed only for CUSTOM schema`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"state_or_province": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"street_address": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"telephone": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"unreachable_domains_behaviour": &schema.Schema{
										Description: `Allowed values: PROCEED, DROP`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Description: `No character restriction`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"domain": &schema.Schema{
							Description: `The AD domain. Alphanumeric, hyphen (-) and dot (.) characters are allowed`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"enable_domain_allowed_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_domain_white_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Resource UUID value`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `Resource Name. Maximum 32 characters allowed. Allowed characters are alphanumeric and .-_/\\ characters`,
							Type:        schema.TypeString,
							Computed:    true,
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

						"ad_attributes": &schema.Schema{
							Description: `Holds list of AD Attributes`,
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Description: `List of Attributes`,
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_value": &schema.Schema{
													Description: `Required for each attribute in the attribute list. Can contain an empty string. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"internal_name": &schema.Schema{
													Description: `Required for each attribute in the attribute list. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"name": &schema.Schema{
													Description: `Required for each attribute in the attribute list with no duplication between attributes. All characters are allowed except <%"`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"type": &schema.Schema{
													Description: `Required for each group in the group list. Allowed values: STRING, IP, BOOLEAN, INT, OCTET_STRING`,
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
						"ad_scopes_names": &schema.Schema{
							Description: `String that contains the names of the scopes that the active directory belongs to. Names are separated by comma. Alphanumeric, underscore (_) characters are allowed`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"adgroups": &schema.Schema{
							Description: `Holds list of AD Groups`,
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"groups": &schema.Schema{
										Description: `List of Groups`,
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Required for each group in the group list with no duplication between groups. All characters are allowed except %`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"sid": &schema.Schema{
													Description: `Cisco ISE uses security identifiers (SIDs) for optimization of group membership evaluation. SIDs are useful for efficiency (speed) when the groups are evaluated. All characters are allowed except %`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"type": &schema.Schema{
													Description: `No character restriction`,
													Type:        schema.TypeString,
													Optional:    true,
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
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aging_time": &schema.Schema{
										Description: `Range 1-8760 hours`,
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"auth_protection_type": &schema.Schema{
										Description: `Enable prevent AD account lockout. Allowed values:
- WIRELESS,
- WIRED,
- BOTH`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"country": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"department": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"email": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"enable_callback_for_dialin_client": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"enable_dialin_permission_check": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"enable_failed_auth_protection": &schema.Schema{
										Description:  `Enable prevent AD account lockout due to too many bad password attempts`,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"enable_machine_access": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"enable_machine_auth": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"enable_pass_change": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"enable_rewrites": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"failed_auth_threshold": &schema.Schema{
										Description: `Number of bad password attempts`,
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"first_name": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"identity_not_in_ad_behaviour": &schema.Schema{
										Description: `Allowed values: REJECT, SEARCH_JOINED_FOREST, SEARCH_ALL`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"job_title": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"last_name": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"locality": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"organizational_unit": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"plaintext_auth": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"rewrite_rules": &schema.Schema{
										Description: `Identity rewrite is an advanced feature that directs Cisco ISE to manipulate the identity
before it is passed to the external Active Directory system. You can create rules to change
the identity to a desired format that includes or excludes a domain prefix and/or suffix or
other additional markup of your choice`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"rewrite_match": &schema.Schema{
													Description: `Required for each rule in the list with no duplication between rules. All characters are allowed except %"`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"rewrite_result": &schema.Schema{
													Description: `Required for each rule in the list. All characters are allowed except %"`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"row_id": &schema.Schema{
													Description: `Required for each rule in the list in serial order`,
													Type:        schema.TypeInt,
													Optional:    true,
												},
											},
										},
									},
									"schema": &schema.Schema{
										Description: `Allowed values: ACTIVE_DIRECTORY, CUSTOM.
Choose ACTIVE_DIRECTORY schema when the AD attributes defined in AD can be copied to relevant attributes
in Cisco ISE. If customization is needed, choose CUSTOM schema. All User info attributes are always set to
default value if schema is ACTIVE_DIRECTORY. Values can be changed only for CUSTOM schema`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"state_or_province": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"street_address": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"telephone": &schema.Schema{
										Description: `User info attribute. All characters are allowed except %`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"unreachable_domains_behaviour": &schema.Schema{
										Description: `Allowed values: PROCEED, DROP`,
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Description: `No character restriction`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"domain": &schema.Schema{
							Description: `The AD domain. Alphanumeric, hyphen (-) and dot (.) characters are allowed`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"enable_domain_white_list": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_domain_allowed_list": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"id": &schema.Schema{
							Description: `Resource UUID value`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": &schema.Schema{
							Description: `Resource Name. Maximum 32 characters allowed. Allowed characters are alphanumeric and .-_/\\ characters`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceActiveDirectoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ActiveDirectory create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestActiveDirectoryCreateActiveDirectory(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.ActiveDirectory.GetActiveDirectoryByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceActiveDirectoryRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.ActiveDirectory.GetActiveDirectoryByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceActiveDirectoryRead(ctx, d, m)
		}
	}
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
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceActiveDirectoryRead(ctx, d, m)
}

func resourceActiveDirectoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ActiveDirectory read for id=[%s]", d.Id())
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
		log.Printf("[DEBUG] Selected method: GetActiveDirectoryByName")
		vvName := vName

		response1, restyResp1, err := client.ActiveDirectory.GetActiveDirectoryByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenActiveDirectoryGetActiveDirectoryByNameItemName(response1.ERSActiveDirectory)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetActiveDirectoryByName response to item",
				err))
			return diags
		}
		vItemName2 := flattenActiveDirectoryGetActiveDirectoryByNameItemNameForParams(response1.ERSActiveDirectory)
		if err := d.Set("parameters", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetActiveDirectoryByName response to parameters",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetActiveDirectoryByID")
		vvID := vID

		response2, restyResp2, err := client.ActiveDirectory.GetActiveDirectoryByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

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
	log.Printf("[DEBUG] Beginning ActiveDirectory update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		if _, ok := d.GetOk("parameters.0"); ok {
			if _, ok := d.GetOk("parameters.0.adgroups"); ok {
				if d.HasChange("parameters.0.adgroups") {
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
					if selectedMethod == 1 {
						getResp, _, err := client.ActiveDirectory.GetActiveDirectoryByID(vID)
						if err == nil && getResp != nil {
							vvID = vID
						}
					}
					if selectedMethod == 2 {
						getResp, _, err := client.ActiveDirectory.GetActiveDirectoryByName(vName)
						if err == nil && getResp != nil {
							if getResp.ERSActiveDirectory != nil {
								vvID = getResp.ERSActiveDirectory.ID
							}
						}
					}

					if vvID != "" {
						log.Printf("[DEBUG] Selected method: LoadGroupsFromDomain")
						request1 := expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomain(ctx, "parameters.0", d)

						if request1 != nil {
							log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
						}
						if request1 != nil && request1.ERSActiveDirectory != nil {
							request1.ERSActiveDirectory.ID = vvID
						}
						response1, err := client.ActiveDirectory.LoadGroupsFromDomain(vvID, request1)

						if err != nil || response1 == nil {
							if response1 != nil {
								log.Printf("[DEBUG] Retrieved error response %s", response1.String())
								diags = append(diags, diagErrorWithAltAndResponse(
									"Failure when executing LoadGroupsFromDomain", err, response1.String(),
									"Failure at LoadGroupsFromDomain, unexpected response", ""))
								return diags
							}
							diags = append(diags, diagErrorWithAlt(
								"Failure when executing LoadGroupsFromDomain", err,
								"Failure at LoadGroupsFromDomain, unexpected response", ""))
							return diags
						}

						log.Printf("[DEBUG] Retrieved response %s", response1.String())
						d.Set("last_updated", getUnixTimeString())
					}
				}
			}
		}
	} else {
		log.Printf("[DEBUG] Missing ActiveDirectory update on Cisco ISE. It will only be update it on Terraform. Resource only performs update if it has changes in adgroups.")
	}
	return resourceActiveDirectoryRead(ctx, d, m)
}

func resourceActiveDirectoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ActiveDirectory delete for id=[%s]", d.Id())
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
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain")))) {
		request.Domain = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_domain_white_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_domain_white_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_domain_white_list")))) {
		request.EnableDomainWhiteList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".adgroups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".adgroups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".adgroups")))) {
		request.Adgroups = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroups(ctx, key+".adgroups.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_settings")))) {
		request.AdvancedSettings = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettings(ctx, key+".advanced_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ad_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ad_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ad_attributes")))) {
		request.AdAttributes = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributes(ctx, key+".ad_attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ad_scopes_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ad_scopes_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ad_scopes_names")))) {
		request.AdScopesNames = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroups {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".groups")))) {
		request.Groups = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroupsArray(ctx, key+".groups", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroupsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups {
	request := []isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups{}
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
		i := expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdgroupsGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sid")))) {
		request.Sid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettings {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_pass_change")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_pass_change")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_pass_change")))) {
		request.EnablePassChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_machine_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_machine_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_machine_auth")))) {
		request.EnableMachineAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_machine_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_machine_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_machine_access")))) {
		request.EnableMachineAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aging_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aging_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aging_time")))) {
		request.AgingTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_dialin_permission_check")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_dialin_permission_check")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_dialin_permission_check")))) {
		request.EnableDialinPermissionCheck = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_callback_for_dialin_client")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_callback_for_dialin_client")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_callback_for_dialin_client")))) {
		request.EnableCallbackForDialinClient = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".plaintext_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".plaintext_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".plaintext_auth")))) {
		request.PlaintextAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_failed_auth_protection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_failed_auth_protection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_failed_auth_protection")))) {
		request.EnableFailedAuthProtection = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_protection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_protection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_protection_type")))) {
		request.AuthProtectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failed_auth_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failed_auth_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failed_auth_threshold")))) {
		failedAuthThreshold := interfaceToIntPtr(v)
		if failedAuthThreshold != nil && *failedAuthThreshold > 0 {
			request.FailedAuthThreshold = failedAuthThreshold
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_not_in_ad_behaviour")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_not_in_ad_behaviour")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_not_in_ad_behaviour")))) {
		request.IDentityNotInAdBehaviour = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".unreachable_domains_behaviour")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".unreachable_domains_behaviour")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".unreachable_domains_behaviour")))) {
		request.UnreachableDomainsBehaviour = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_rewrites")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_rewrites")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_rewrites")))) {
		request.EnableRewrites = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rewrite_rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rewrite_rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rewrite_rules")))) {
		request.RewriteRules = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRulesArray(ctx, key+".rewrite_rules", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_name")))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".department")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".department")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".department")))) {
		request.Department = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_name")))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".organizational_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".organizational_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".organizational_unit")))) {
		request.OrganizationalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".job_title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".job_title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".job_title")))) {
		request.JobTitle = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".locality")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".locality")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".locality")))) {
		request.Locality = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email")))) {
		request.Email = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state_or_province")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state_or_province")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state_or_province")))) {
		request.StateOrProvince = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".telephone")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".telephone")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".telephone")))) {
		request.Telephone = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".country")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".country")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".country")))) {
		request.Country = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".street_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".street_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".street_address")))) {
		request.StreetAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".schema")))) {
		request.Schema = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules {
	request := []isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules{}
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
		i := expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdvancedSettingsRewriteRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".row_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".row_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".row_id")))) {
		request.RowID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rewrite_match")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rewrite_match")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rewrite_match")))) {
		request.RewriteMatch = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rewrite_result")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rewrite_result")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rewrite_result")))) {
		request.RewriteResult = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributes {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributesArray(ctx, key+".attributes", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes {
	request := []isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes{}
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
		i := expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes {
	request := isegosdk.RequestActiveDirectoryCreateActiveDirectoryERSActiveDirectoryAdAttributesAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".internal_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".internal_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".internal_name")))) {
		request.InternalName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomain(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLoadGroupsFromDomain {
	request := isegosdk.RequestActiveDirectoryLoadGroupsFromDomain{}
	request.ERSActiveDirectory = expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectory(ctx, key, d)
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectory(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectory {
	request := isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectory{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain")))) {
		request.Domain = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_domain_white_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_domain_white_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_domain_white_list")))) {
		request.EnableDomainWhiteList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".adgroups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".adgroups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".adgroups")))) {
		request.Adgroups = expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdgroups(ctx, key+".adgroups.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_settings")))) {
		request.AdvancedSettings = expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdvancedSettings(ctx, key+".advanced_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ad_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ad_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ad_attributes")))) {
		request.AdAttributes = expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdAttributes(ctx, key+".ad_attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ad_scopes_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ad_scopes_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ad_scopes_names")))) {
		request.AdScopesNames = interfaceToString(v)
	}
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdgroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdgroups {
	request := isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdgroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".groups")))) {
		request.Groups = expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdgroupsGroupsArray(ctx, key+".groups", d)
	}
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdgroupsGroupsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdgroupsGroups {
	request := []isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdgroupsGroups{}
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
		i := expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdgroupsGroups(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdgroupsGroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdgroupsGroups {
	request := isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdgroupsGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sid")))) {
		request.Sid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdvancedSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdvancedSettings {
	request := isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdvancedSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_pass_change")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_pass_change")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_pass_change")))) {
		request.EnablePassChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_machine_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_machine_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_machine_auth")))) {
		request.EnableMachineAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_machine_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_machine_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_machine_access")))) {
		request.EnableMachineAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aging_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aging_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aging_time")))) {
		request.AgingTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_dialin_permission_check")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_dialin_permission_check")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_dialin_permission_check")))) {
		request.EnableDialinPermissionCheck = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_callback_for_dialin_client")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_callback_for_dialin_client")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_callback_for_dialin_client")))) {
		request.EnableCallbackForDialinClient = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".plaintext_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".plaintext_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".plaintext_auth")))) {
		request.PlaintextAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_failed_auth_protection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_failed_auth_protection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_failed_auth_protection")))) {
		request.EnableFailedAuthProtection = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_protection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_protection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_protection_type")))) {
		request.AuthProtectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failed_auth_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failed_auth_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failed_auth_threshold")))) {
		failedAuthThreshold := interfaceToIntPtr(v)
		if failedAuthThreshold != nil && *failedAuthThreshold > 0 {
			request.FailedAuthThreshold = failedAuthThreshold
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_not_in_ad_behaviour")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_not_in_ad_behaviour")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_not_in_ad_behaviour")))) {
		request.IDentityNotInAdBehaviour = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".unreachable_domains_behaviour")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".unreachable_domains_behaviour")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".unreachable_domains_behaviour")))) {
		request.UnreachableDomainsBehaviour = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_rewrites")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_rewrites")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_rewrites")))) {
		request.EnableRewrites = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rewrite_rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rewrite_rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rewrite_rules")))) {
		request.RewriteRules = expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdvancedSettingsRewriteRulesArray(ctx, key+".rewrite_rules", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_name")))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".department")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".department")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".department")))) {
		request.Department = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_name")))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".organizational_unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".organizational_unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".organizational_unit")))) {
		request.OrganizationalUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".job_title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".job_title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".job_title")))) {
		request.JobTitle = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".locality")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".locality")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".locality")))) {
		request.Locality = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email")))) {
		request.Email = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state_or_province")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state_or_province")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state_or_province")))) {
		request.StateOrProvince = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".telephone")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".telephone")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".telephone")))) {
		request.Telephone = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".country")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".country")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".country")))) {
		request.Country = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".street_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".street_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".street_address")))) {
		request.StreetAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".schema")))) {
		request.Schema = interfaceToString(v)
	}
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdvancedSettingsRewriteRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdvancedSettingsRewriteRules {
	request := []isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdvancedSettingsRewriteRules{}
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
		i := expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdvancedSettingsRewriteRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdvancedSettingsRewriteRules(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdvancedSettingsRewriteRules {
	request := isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdvancedSettingsRewriteRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".row_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".row_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".row_id")))) {
		request.RowID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rewrite_match")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rewrite_match")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rewrite_match")))) {
		request.RewriteMatch = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rewrite_result")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rewrite_result")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rewrite_result")))) {
		request.RewriteResult = interfaceToString(v)
	}
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdAttributes {
	request := isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdAttributesAttributesArray(ctx, key+".attributes", d)
	}
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdAttributesAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdAttributesAttributes {
	request := []isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdAttributesAttributes{}
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
		i := expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdAttributesAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomainERSActiveDirectoryAdAttributesAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdAttributesAttributes {
	request := isegosdk.RequestActiveDirectoryLoadGroupsFromDomainERSActiveDirectoryAdAttributesAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".internal_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".internal_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".internal_name")))) {
		request.InternalName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	return &request
}
