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

func resourceLicensingTierState() *schema.Resource {
	return &schema.Resource{
		Description: `It manages ENABLED and DISABLED operations on Tier - state information.

- Applicable values for name & status parameters:

* name: ESSENTIAL, ADVANTAGE, PREMIER, DEVICEADMIN

* status: ENABLED, DISABLED
`,

		CreateContext: resourceLicensingTierStateCreate,
		ReadContext:   resourceLicensingTierStateRead,
		UpdateContext: resourceLicensingTierStateUpdate,
		DeleteContext: resourceLicensingTierStateDelete,
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

						"compliance": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"consumption_counter": &schema.Schema{
							Description: `Compliance counter for tier`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"days_out_of_compliance": &schema.Schema{
							Description: `Number of days tier is out of compliance`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_authorization": &schema.Schema{
							Description: `Last date of authorization`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
						"name": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validateStringHasValueFunc([]string{"", "ESSENTIAL", "ADVANTAGE", "PREMIER", "DEVICEADMIN"}),
						},
						"status": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringHasValueFunc([]string{"", "ENABLED", "DISABLED"}),
						},
					},
				},
			},
		},
	}
}

func resourceLicensingTierStateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TierState create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestLicensingTierStateCreateUpdateTierStateInfo(ctx, "", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	response1, restyResp1, err := client.Licensing.UpdateTierStateInfo(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing UpdateTierStateInfo", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing UpdateTierStateInfo", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return resourceLicensingTierStateRead(ctx, d, m)
}

func resourceLicensingTierStateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TierState read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]

	response1, restyResp1, err := client.Licensing.GetTierStateInfo()

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	response_nodes, err := searchLicensingGetTierStateInfo(m, response1, vName)
	if err != nil || response_nodes == nil || len(*response_nodes) == 0 {
		d.SetId("")
		return diags
	}

	vItem1 := flattenLicensingGetTierStateInfoItems(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetTierStateInfo response",
			err))
		return diags
	}
	return diags
}

func resourceLicensingTierStateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TierState update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestLicensingTierStateCreateUpdateTierStateInfo(ctx, "", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		response1, restyResp1, err := client.Licensing.UpdateTierStateInfo(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				diags = append(diags, diagErrorWithResponse(
					"Failure when executing UpdateTierStateInfo", err, restyResp1.String()))
				return diags
			}
			diags = append(diags, diagError(
				"Failure when executing UpdateTierStateInfo", err))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}
	return resourceLicensingTierStateRead(ctx, d, m)
}

func resourceLicensingTierStateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TierState delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing TierState delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestLicensingTierStateCreateUpdateTierStateInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLicensingUpdateTierStateInfo {
	request := isegosdk.RequestLicensingUpdateTierStateInfo{}
	if v := expandRequestLicensingTierStateCreateUpdateTierStateInfoItemArray(ctx, key+".parameters", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestLicensingTierStateCreateUpdateTierStateInfoItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestItemLicensingUpdateTierStateInfo {
	request := []isegosdk.RequestItemLicensingUpdateTierStateInfo{}
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
		i := expandRequestLicensingTierStateCreateUpdateTierStateInfoItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestLicensingTierStateCreateUpdateTierStateInfoItem(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemLicensingUpdateTierStateInfo {
	request := isegosdk.RequestItemLicensingUpdateTierStateInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	return &request
}

func searchLicensingGetTierStateInfo(m interface{}, items *isegosdk.ResponseLicensingGetTierStateInfo, name string) (*isegosdk.ResponseLicensingGetTierStateInfo, error) {
	var err error
	var foundItem isegosdk.ResponseLicensingGetTierStateInfo
	var foundItems []isegosdk.ResponseItemLicensingGetTierStateInfo
	if items == nil {
		return nil, err
	}
	for _, item := range *items {
		if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			foundItems = append(foundItems, item)
			foundItem = foundItems
			return &foundItem, err
		}
	}
	return nil, err
}
