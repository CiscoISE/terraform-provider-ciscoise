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

func resourceNativeSupplicantProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read, update and delete operations on NativeSupplicantProfile.

- This resource allows the client to update a native supplicant profile

- This resource deletes a native supplicant profile.
`,

		CreateContext: resourceNativeSupplicantProfileCreate,
		ReadContext:   resourceNativeSupplicantProfileRead,
		UpdateContext: resourceNativeSupplicantProfileUpdate,
		DeleteContext: resourceNativeSupplicantProfileDelete,
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

						"description": &schema.Schema{
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
						"wireless_profiles": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action_type": &schema.Schema{
										Description: `Action type for WifiProfile.
Allowed values:
- ADD,
- UPDATE,
- DELETE
(required for updating existing WirelessProfile)`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"allowed_protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"certificate_template_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"previous_ssid": &schema.Schema{
										Description: `Previous ssid for WifiProfile (required for updating existing WirelessProfile)`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
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

						"description": &schema.Schema{
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
						"wireless_profiles": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action_type": &schema.Schema{
										Description: `Action type for WifiProfile.
Allowed values:
- ADD,
- UPDATE,
- DELETE
(required for updating existing WirelessProfile)`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"allowed_protocol": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"certificate_template_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"previous_ssid": &schema.Schema{
										Description: `Previous ssid for WifiProfile (required for updating existing WirelessProfile)`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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

func resourceNativeSupplicantProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NativeSupplicantProfile create")
	log.Printf("[DEBUG] Missing NativeSupplicantProfile create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	vvID := interfaceToString(resourceItem["id"])
	log.Printf("[DEBUG] ID used for update operation %s", vvID)
	request1 := expandRequestNativeSupplicantProfileUpdateNativeSupplicantProfileByID(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, restyResp1, err := client.NativeSupplicantProfile.UpdateNativeSupplicantProfileByID(vvID, request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing UpdateNativeSupplicantProfileByID", err, restyResp1.String(),
				"Failure at UpdateNativeSupplicantProfileByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing UpdateNativeSupplicantProfileByID", err,
			"Failure at UpdateNativeSupplicantProfileByID, unexpected response", ""))
		return diags
	}
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return resourceNativeSupplicantProfileRead(ctx, d, m)
}

func resourceNativeSupplicantProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NativeSupplicantProfile read for id=[%s]", d.Id())
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
		log.Printf("[DEBUG] Selected method: GetNativeSupplicantProfile")
		queryParams1 := isegosdk.GetNativeSupplicantProfileQueryParams{}

		response1, restyResp1, err := client.NativeSupplicantProfile.GetNativeSupplicantProfile(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsNativeSupplicantProfileGetNativeSupplicantProfile(m, response1, &queryParams1)
		item1, err := searchNativeSupplicantProfileGetNativeSupplicantProfile(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenNativeSupplicantProfileGetNativeSupplicantProfileByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNativeSupplicantProfile search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNativeSupplicantProfileByID")
		vvID := vID

		response2, restyResp2, err := client.NativeSupplicantProfile.GetNativeSupplicantProfileByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNativeSupplicantProfileGetNativeSupplicantProfileByIDItem(response2.ERSNSpProfile)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNativeSupplicantProfileByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNativeSupplicantProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NativeSupplicantProfile update for id=[%s]", d.Id())
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
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetNativeSupplicantProfileQueryParams{}

		getResp1, _, err := client.NativeSupplicantProfile.GetNativeSupplicantProfile(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsNativeSupplicantProfileGetNativeSupplicantProfile(m, getResp1, &queryParams1)
			item1, err := searchNativeSupplicantProfileGetNativeSupplicantProfile(m, items1, vName, vID)
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
		request1 := expandRequestNativeSupplicantProfileUpdateNativeSupplicantProfileByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.NativeSupplicantProfile.UpdateNativeSupplicantProfileByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNativeSupplicantProfileByID", err, restyResp1.String(),
					"Failure at UpdateNativeSupplicantProfileByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNativeSupplicantProfileByID", err,
				"Failure at UpdateNativeSupplicantProfileByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceNativeSupplicantProfileRead(ctx, d, m)
}

func resourceNativeSupplicantProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NativeSupplicantProfile delete for id=[%s]", d.Id())
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
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetNativeSupplicantProfileQueryParams{}

		getResp1, _, err := client.NativeSupplicantProfile.GetNativeSupplicantProfile(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNativeSupplicantProfileGetNativeSupplicantProfile(m, getResp1, &queryParams1)
		item1, err := searchNativeSupplicantProfileGetNativeSupplicantProfile(m, items1, vName, vID)
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
		getResp, _, err := client.NativeSupplicantProfile.GetNativeSupplicantProfileByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.NativeSupplicantProfile.DeleteNativeSupplicantProfileByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNativeSupplicantProfileByID", err, restyResp1.String(),
				"Failure at DeleteNativeSupplicantProfileByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNativeSupplicantProfileByID", err,
			"Failure at DeleteNativeSupplicantProfileByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNativeSupplicantProfileUpdateNativeSupplicantProfileByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByID {
	request := isegosdk.RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByID{}
	request.ERSNSpProfile = expandRequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfile(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfile(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfile {
	request := isegosdk.RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wireless_profiles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wireless_profiles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wireless_profiles")))) {
		request.WirelessProfiles = expandRequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfileWirelessProfilesArray(ctx, key+".wireless_profiles", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfileWirelessProfilesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles {
	request := []isegosdk.RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles{}
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
		i := expandRequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles {
	request := isegosdk.RequestNativeSupplicantProfileUpdateNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allowed_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allowed_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allowed_protocol")))) {
		request.AllowedProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_template_id")))) {
		request.CertificateTemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".action_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".action_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".action_type")))) {
		request.ActionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".previous_ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".previous_ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".previous_ssid")))) {
		request.PreviousSSID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNativeSupplicantProfileGetNativeSupplicantProfile(m interface{}, response *isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfile, queryParams *isegosdk.GetNativeSupplicantProfileQueryParams) []isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultResources
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
			response, _, err = client.NativeSupplicantProfile.GetNativeSupplicantProfile(queryParams)
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

func searchNativeSupplicantProfileGetNativeSupplicantProfile(m interface{}, items []isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultResources, name string, id string) (*isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfile, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfile
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileByID
			getItem, _, err = client.NativeSupplicantProfile.GetNativeSupplicantProfileByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNativeSupplicantProfileByID")
			}
			foundItem = getItem.ERSNSpProfile
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileByID
			getItem, _, err = client.NativeSupplicantProfile.GetNativeSupplicantProfileByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNativeSupplicantProfileByID")
			}
			foundItem = getItem.ERSNSpProfile
			return foundItem, err
		}
	}
	return foundItem, err
}
