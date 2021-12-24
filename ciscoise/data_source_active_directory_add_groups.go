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

// dataSourceAction
func dataSourceActiveDirectoryAddGroups() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ActiveDirectory.

- This data source action loads domain groups configuration from Active Directory into Cisco ISE.
`,

		ReadContext: dataSourceActiveDirectoryAddGroupsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"ad_attributes": &schema.Schema{
				Description: `Holds list of AD Attributes`,
				Type:        schema.TypeList,
				Optional:    true,
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
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Description: `Resource Name. Maximum 32 characters allowed. Allowed characters are alphanumeric and .-_/\\ characters`,
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceActiveDirectoryAddGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: LoadGroupsFromDomain")
		vvID := vID.(string)
		request1 := expandRequestActiveDirectoryAddGroupsLoadGroupsFromDomain(ctx, "", d)

		response1, err := client.ActiveDirectory.LoadGroupsFromDomain(vvID, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

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

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LoadGroupsFromDomain response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
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
		request.FailedAuthThreshold = interfaceToIntPtr(v)
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
