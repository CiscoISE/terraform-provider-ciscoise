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

// dataSourceAction
func dataSourceActiveDirectoryGetGroupsByDomainInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceActiveDirectoryGetGroupsByDomainInfoRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func dataSourceActiveDirectoryGetGroupsByDomainInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGroupsByDomain")
		vvID := vID.(string)
		request1 := expandRequestActiveDirectoryGetGroupsByDomainInfoGetGroupsByDomain(ctx, "", d)

		response1, _, err := client.ActiveDirectory.GetGroupsByDomain(vvID, request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGroupsByDomain", err,
				"Failure at GetGroupsByDomain, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenActiveDirectoryGetGroupsByDomainItem(response1.ERSActiveDirectoryGroups)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGroupsByDomain response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestActiveDirectoryGetGroupsByDomainInfoGetGroupsByDomain(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryGetGroupsByDomain {
	request := isegosdk.RequestActiveDirectoryGetGroupsByDomain{}
	request.OperationAdditionalData = expandRequestActiveDirectoryGetGroupsByDomainInfoGetGroupsByDomainOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestActiveDirectoryGetGroupsByDomainInfoGetGroupsByDomainOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryGetGroupsByDomainOperationAdditionalData {
	request := isegosdk.RequestActiveDirectoryGetGroupsByDomainOperationAdditionalData{}
	if v, ok := d.GetOkExists("additional_data"); !isEmptyValue(reflect.ValueOf(d.Get("additional_data"))) && (ok || !reflect.DeepEqual(v, d.Get("additional_data"))) {
		request.AdditionalData = expandRequestActiveDirectoryGetGroupsByDomainInfoGetGroupsByDomainOperationAdditionalDataAdditionalDataArray(ctx, key, d)
	}
	return &request
}

func expandRequestActiveDirectoryGetGroupsByDomainInfoGetGroupsByDomainOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryGetGroupsByDomainOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestActiveDirectoryGetGroupsByDomainOperationAdditionalDataAdditionalData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestActiveDirectoryGetGroupsByDomainInfoGetGroupsByDomainOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	return &request
}

func expandRequestActiveDirectoryGetGroupsByDomainInfoGetGroupsByDomainOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryGetGroupsByDomainOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestActiveDirectoryGetGroupsByDomainOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists("value"); !isEmptyValue(reflect.ValueOf(d.Get("value"))) && (ok || !reflect.DeepEqual(v, d.Get("value"))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(d.Get("name"))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	return &request
}

func flattenActiveDirectoryGetGroupsByDomainItem(item *isegosdk.ResponseActiveDirectoryGetGroupsByDomainERSActiveDirectoryGroups) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["groups"] = flattenActiveDirectoryGetGroupsByDomainItemGroups(item.Groups)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenActiveDirectoryGetGroupsByDomainItemGroups(items *[]isegosdk.ResponseActiveDirectoryGetGroupsByDomainERSActiveDirectoryGroupsGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["group_name"] = item.GroupName
		respItem["sid"] = item.Sid
		respItem["type"] = item.Type
	}
	return respItems

}
