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

func resourcePortalGlobalSetting() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on PortalGlobalSetting.

- This resource allows the client to update the portal global settings by id.
`,

		CreateContext: resourcePortalGlobalSettingCreate,
		ReadContext:   resourcePortalGlobalSettingRead,
		UpdateContext: resourcePortalGlobalSettingUpdate,
		DeleteContext: resourcePortalGlobalSettingDelete,
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

						"customization": &schema.Schema{
							Description: `Allowed values:
- HTML,
- HTMLANDJAVASCRIPT`,
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
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"customization": &schema.Schema{
							Description: `Allowed values:
- HTML,
- HTMLANDJAVASCRIPT`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourcePortalGlobalSettingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return resourcePortalGlobalSettingRead(ctx, d, m)
}

func resourcePortalGlobalSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetPortalGlobalSettings")
		queryParams1 := isegosdk.GetPortalGlobalSettingsQueryParams{}

		response1, restyResp1, err := client.PortalGlobalSetting.GetPortalGlobalSettings(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPortalGlobalSettings", err,
				"Failure at GetPortalGlobalSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsPortalGlobalSettingGetPortalGlobalSettings(m, response1, &queryParams1)
		item1, err := searchPortalGlobalSettingGetPortalGlobalSettings(m, items1, "", vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetPortalGlobalSettings response", err,
				"Failure when searching item from GetPortalGlobalSettings, unexpected response", ""))
			return diags
		}
		vItem1 := flattenPortalGlobalSettingGetPortalGlobalSettingByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortalGlobalSettings search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetPortalGlobalSettingByID")
		vvID := vID

		response2, restyResp2, err := client.PortalGlobalSetting.GetPortalGlobalSettingByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPortalGlobalSettingByID", err,
				"Failure at GetPortalGlobalSettingByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenPortalGlobalSettingGetPortalGlobalSettingByIDItem(response2.PortalCustomizationGlobalSetting)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortalGlobalSettingByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourcePortalGlobalSettingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	if selectedMethod == 2 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestPortalGlobalSettingUpdatePortalGlobalSettingByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.PortalGlobalSetting.UpdatePortalGlobalSettingByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatePortalGlobalSettingByID", err, restyResp1.String(),
					"Failure at UpdatePortalGlobalSettingByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatePortalGlobalSettingByID", err,
				"Failure at UpdatePortalGlobalSettingByID, unexpected response", ""))
			return diags
		}
	}

	return resourcePortalGlobalSettingRead(ctx, d, m)
}

func resourcePortalGlobalSettingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete PortalGlobalSetting on Cisco ISE
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestPortalGlobalSettingUpdatePortalGlobalSettingByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPortalGlobalSettingUpdatePortalGlobalSettingByID {
	request := isegosdk.RequestPortalGlobalSettingUpdatePortalGlobalSettingByID{}
	request.PortalCustomizationGlobalSetting = expandRequestPortalGlobalSettingUpdatePortalGlobalSettingByIDPortalCustomizationGlobalSetting(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPortalGlobalSettingUpdatePortalGlobalSettingByIDPortalCustomizationGlobalSetting(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPortalGlobalSettingUpdatePortalGlobalSettingByIDPortalCustomizationGlobalSetting {
	request := isegosdk.RequestPortalGlobalSettingUpdatePortalGlobalSettingByIDPortalCustomizationGlobalSetting{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".customization")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".customization")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".customization")))) {
		request.Customization = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsPortalGlobalSettingGetPortalGlobalSettings(m interface{}, response *isegosdk.ResponsePortalGlobalSettingGetPortalGlobalSettings, queryParams *isegosdk.GetPortalGlobalSettingsQueryParams) []isegosdk.ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultResources
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
			response, _, err = client.PortalGlobalSetting.GetPortalGlobalSettings(queryParams)
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

func searchPortalGlobalSettingGetPortalGlobalSettings(m interface{}, items []isegosdk.ResponsePortalGlobalSettingGetPortalGlobalSettingsSearchResultResources, name string, id string) (*isegosdk.ResponsePortalGlobalSettingGetPortalGlobalSettingByIDPortalCustomizationGlobalSetting, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponsePortalGlobalSettingGetPortalGlobalSettingByIDPortalCustomizationGlobalSetting
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponsePortalGlobalSettingGetPortalGlobalSettingByID
			getItem, _, err = client.PortalGlobalSetting.GetPortalGlobalSettingByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetPortalGlobalSettingByID")
			}
			foundItem = getItem.PortalCustomizationGlobalSetting
			return foundItem, err
		}
	}
	return foundItem, err
}
