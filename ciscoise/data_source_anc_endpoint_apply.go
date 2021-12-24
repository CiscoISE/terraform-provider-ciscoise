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
func dataSourceAncEndpointApply() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ANCEndpoint.

- This data source action allows the client to apply the required configuration.
`,

		ReadContext: dataSourceAncEndpointApplyRead,
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

func dataSourceAncEndpointApplyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ApplyAncEndpoint")
		request1 := expandRequestAncEndpointApplyApplyAncEndpoint(ctx, "", d)

		response1, err := client.AncEndpoint.ApplyAncEndpoint(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing ApplyAncEndpoint", err, response1.String(),
					"Failure at ApplyAncEndpoint, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ApplyAncEndpoint", err,
				"Failure at ApplyAncEndpoint, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ApplyAncEndpoint response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestAncEndpointApplyApplyAncEndpoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncEndpointApplyAncEndpoint {
	request := isegosdk.RequestAncEndpointApplyAncEndpoint{}
	request.OperationAdditionalData = expandRequestAncEndpointApplyApplyAncEndpointOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestAncEndpointApplyApplyAncEndpointOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalData {
	request := isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_data")))) {
		request.AdditionalData = expandRequestAncEndpointApplyApplyAncEndpointOperationAdditionalDataAdditionalDataArray(ctx, key+".additional_data", d)
	}
	return &request
}

func expandRequestAncEndpointApplyApplyAncEndpointOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalDataAdditionalData{}
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
		i := expandRequestAncEndpointApplyApplyAncEndpointOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAncEndpointApplyApplyAncEndpointOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
