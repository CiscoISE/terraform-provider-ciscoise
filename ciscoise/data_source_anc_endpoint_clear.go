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
func dataSourceAncEndpointClear() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ANCEndpoint.

- This data source action allows the client to clear the required configuration.
`,

		ReadContext: dataSourceAncEndpointClearRead,
		Schema: map[string]*schema.Schema{
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

func dataSourceAncEndpointClearRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ClearAncEndpoint")
		request1 := expandRequestAncEndpointClearClearAncEndpoint(ctx, "", d)

		response1, err := client.AncEndpoint.ClearAncEndpoint(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ClearAncEndpoint", err,
				"Failure at ClearAncEndpoint, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ClearAncEndpoint response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestAncEndpointClearClearAncEndpoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncEndpointClearAncEndpoint {
	request := isegosdk.RequestAncEndpointClearAncEndpoint{}
	request.OperationAdditionalData = expandRequestAncEndpointClearClearAncEndpointOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestAncEndpointClearClearAncEndpointOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalData {
	request := isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalData{}
	if v, ok := d.GetOkExists("additional_data"); !isEmptyValue(reflect.ValueOf(d.Get("additional_data"))) && (ok || !reflect.DeepEqual(v, d.Get("additional_data"))) {
		request.AdditionalData = expandRequestAncEndpointClearClearAncEndpointOperationAdditionalDataAdditionalDataArray(ctx, key, d)
	}
	return &request
}

func expandRequestAncEndpointClearClearAncEndpointOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalDataAdditionalData{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAncEndpointClearClearAncEndpointOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	return &request
}

func expandRequestAncEndpointClearClearAncEndpointOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists("value"); !isEmptyValue(reflect.ValueOf(d.Get("value"))) && (ok || !reflect.DeepEqual(v, d.Get("value"))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(d.Get("name"))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
