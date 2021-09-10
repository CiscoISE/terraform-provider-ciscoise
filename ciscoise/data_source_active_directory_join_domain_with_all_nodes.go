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
func dataSourceActiveDirectoryJoinDomainWithAllNodes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceActiveDirectoryJoinDomainWithAllNodesRead,
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

func dataSourceActiveDirectoryJoinDomainWithAllNodesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: JoinDomainWithAllNodes")
		vvID := vID.(string)
		request1 := expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodes(ctx, "", d)

		response1, err := client.ActiveDirectory.JoinDomainWithAllNodes(vvID, request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing JoinDomainWithAllNodes", err,
				"Failure at JoinDomainWithAllNodes, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting JoinDomainWithAllNodes response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryJoinDomainWithAllNodes {
	request := isegosdk.RequestActiveDirectoryJoinDomainWithAllNodes{}
	request.OperationAdditionalData = expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalData {
	request := isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalData{}
	if v, ok := d.GetOkExists("additional_data"); !isEmptyValue(reflect.ValueOf(d.Get("additional_data"))) && (ok || !reflect.DeepEqual(v, d.Get("additional_data"))) {
		request.AdditionalData = expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalDataAdditionalDataArray(ctx, key, d)
	}
	return &request
}

func expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalDataAdditionalData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	return &request
}

func expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists("value"); !isEmptyValue(reflect.ValueOf(d.Get("value"))) && (ok || !reflect.DeepEqual(v, d.Get("value"))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(d.Get("name"))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
