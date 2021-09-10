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

func resourcePortalGlobalSetting() *schema.Resource {
	return &schema.Resource{

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
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"customization": &schema.Schema{
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
					},
				},
			},
		},
	}
}

func resourcePortalGlobalSettingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("item"))
	resourceMap := make(map[string]string)
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourcePortalGlobalSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetPortalGlobalSettings")
		queryParams1 := isegosdk.GetPortalGlobalSettingsQueryParams{}

		response1, _, err := client.PortalGlobalSetting.GetPortalGlobalSettings(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPortalGlobalSettings", err,
				"Failure at GetPortalGlobalSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsPortalGlobalSettingGetPortalGlobalSettings(m, response1, &queryParams1)
		item1, err := searchPortalGlobalSettingGetPortalGlobalSettings(m, items1, "", vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetPortalGlobalSettings response", err,
				"Failure when searching item from GetPortalGlobalSettings, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortalGlobalSettings search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetPortalGlobalSettingByID")
		vvID := vID

		response2, _, err := client.PortalGlobalSetting.GetPortalGlobalSettingByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPortalGlobalSettingByID", err,
				"Failure at GetPortalGlobalSettingByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

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
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if selectedMethod == 2 {
		vvID = vID
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestPortalGlobalSettingUpdatePortalGlobalSettingByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.PortalGlobalSetting.UpdatePortalGlobalSettingByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
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
	// NOTE: Function does not perform delete on ISE
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
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".customization"); !isEmptyValue(reflect.ValueOf(d.Get(key+".customization"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".customization"))) {
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
