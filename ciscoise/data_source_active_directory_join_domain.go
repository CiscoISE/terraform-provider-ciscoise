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
func dataSourceActiveDirectoryJoinDomain() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceActiveDirectoryJoinDomainRead,
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
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceActiveDirectoryJoinDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: JoinDomain")
		vvID := vID.(string)
		request1 := expandRequestActiveDirectoryJoinDomainJoinDomain(ctx, "", d)

		response1, err := client.ActiveDirectory.JoinDomain(vvID, request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing JoinDomain", err,
				"Failure at JoinDomain, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting JoinDomain response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestActiveDirectoryJoinDomainJoinDomain(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryJoinDomain {
	request := isegosdk.RequestActiveDirectoryJoinDomain{}
	request.OperationAdditionalData = expandRequestActiveDirectoryJoinDomainJoinDomainOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestActiveDirectoryJoinDomainJoinDomainOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryJoinDomainOperationAdditionalData {
	request := isegosdk.RequestActiveDirectoryJoinDomainOperationAdditionalData{}
	if v, ok := d.GetOkExists("additional_data"); !isEmptyValue(reflect.ValueOf(d.Get("additional_data"))) && (ok || !reflect.DeepEqual(v, d.Get("additional_data"))) {
		request.AdditionalData = expandRequestActiveDirectoryJoinDomainJoinDomainOperationAdditionalDataAdditionalDataArray(ctx, key, d)
	}
	return &request
}

func expandRequestActiveDirectoryJoinDomainJoinDomainOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryJoinDomainOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestActiveDirectoryJoinDomainOperationAdditionalDataAdditionalData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestActiveDirectoryJoinDomainJoinDomainOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	return &request
}

func expandRequestActiveDirectoryJoinDomainJoinDomainOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryJoinDomainOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestActiveDirectoryJoinDomainOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists("value"); !isEmptyValue(reflect.ValueOf(d.Get("value"))) && (ok || !reflect.DeepEqual(v, d.Get("value"))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(d.Get("name"))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
