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
func dataSourceActiveDirectoryGetUserGroupsInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ActiveDirectory.

- This data source action allows the client to get groups of which a given user is a member.
`,

		ReadContext: dataSourceActiveDirectoryGetUserGroupsInfoRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"additional_data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"group_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"sid": &schema.Schema{
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
		},
	}
}

func dataSourceActiveDirectoryGetUserGroupsInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetUserGroups")
		vvID := vID.(string)
		request1 := expandRequestActiveDirectoryGetUserGroupsInfoGetUserGroups(ctx, "", d)

		response1, restyResp1, err := client.ActiveDirectory.GetUserGroups(vvID, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetUserGroups", err,
				"Failure at GetUserGroups, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenActiveDirectoryGetUserGroupsItem(response1.ERSActiveDirectoryGroups)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetUserGroups response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestActiveDirectoryGetUserGroupsInfoGetUserGroups(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryGetUserGroups {
	request := isegosdk.RequestActiveDirectoryGetUserGroups{}
	request.OperationAdditionalData = expandRequestActiveDirectoryGetUserGroupsInfoGetUserGroupsOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestActiveDirectoryGetUserGroupsInfoGetUserGroupsOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryGetUserGroupsOperationAdditionalData {
	request := isegosdk.RequestActiveDirectoryGetUserGroupsOperationAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_data")))) {
		request.AdditionalData = expandRequestActiveDirectoryGetUserGroupsInfoGetUserGroupsOperationAdditionalDataAdditionalDataArray(ctx, key+".additional_data", d)
	}
	return &request
}

func expandRequestActiveDirectoryGetUserGroupsInfoGetUserGroupsOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryGetUserGroupsOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestActiveDirectoryGetUserGroupsOperationAdditionalDataAdditionalData{}
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
		i := expandRequestActiveDirectoryGetUserGroupsInfoGetUserGroupsOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestActiveDirectoryGetUserGroupsInfoGetUserGroupsOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryGetUserGroupsOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestActiveDirectoryGetUserGroupsOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}

func flattenActiveDirectoryGetUserGroupsItem(item *isegosdk.ResponseActiveDirectoryGetUserGroupsERSActiveDirectoryGroups) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["groups"] = flattenActiveDirectoryGetUserGroupsItemGroups(item.Groups)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenActiveDirectoryGetUserGroupsItemGroups(items *[]isegosdk.ResponseActiveDirectoryGetUserGroupsERSActiveDirectoryGroupsGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["group_name"] = item.GroupName
		respItem["sid"] = item.Sid
		respItem["type"] = item.Type
		respItems = append(respItems, respItem)
	}
	return respItems
}
