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
func dataSourceLicensingTierStateCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Licensing.

- Applicable values for
name
 &
status
 parameters:


name:
 ESSENTIAL, ADVANTAGE, PREMIER, DEVICEADMIN

status:
 ENABLED, DISABLED

`,

		ReadContext: dataSourceLicensingTierStateCreateRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"message": &schema.Schema{
							Description: `Response message on successful change of license tier state.`,
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
			"payload": &schema.Schema{
				Description: `Array of RequestLicensingUpdateTierStateInfo`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLicensingTierStateCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateTierStateInfo")
		request1 := expandRequestLicensingTierStateCreateUpdateTierStateInfo(ctx, "", d)

		response1, restyResp1, err := client.Licensing.UpdateTierStateInfo(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTierStateInfo", err,
				"Failure at UpdateTierStateInfo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenLicensingUpdateTierStateInfoItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateTierStateInfo response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestLicensingTierStateCreateUpdateTierStateInfo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestLicensingUpdateTierStateInfo {
	request := isegosdk.RequestLicensingUpdateTierStateInfo{}
	if v := expandRequestLicensingTierStateCreateUpdateTierStateInfoItemArray(ctx, key+".payload", d); v != nil {
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
	for item_no, _ := range objs {
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

func flattenLicensingUpdateTierStateInfoItems(items *[]isegosdk.ResponseLicensingUpdateTierStateInfoResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItem["name"] = item.Name
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
