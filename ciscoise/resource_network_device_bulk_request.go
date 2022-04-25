package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDeviceBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on NetworkDevice.
- This data source action allows the client to submit the bulk request.
`,

		CreateContext: resourceNetworkDeviceBulkRequestCreate,
		ReadContext:   resourceNetworkDeviceBulkRequestRead,
		DeleteContext: resourceNetworkDeviceBulkRequestDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operation_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_media_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDeviceBulkRequestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkRequestForNetworkDevice create")
	log.Printf("[DEBUG] Missing BulkRequestForNetworkDevice create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	request1 := expandRequestNetworkDeviceBulkRequestBulkRequestForNetworkDevice(ctx, "parameters.0", d)

	response1, err := client.NetworkDevice.BulkRequestForNetworkDevice(request1)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkRequestForNetworkDevice", err,
			"Failure at BulkRequestForNetworkDevice, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkRequestForNetworkDevice response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return resourceNetworkDeviceBulkRequestRead(ctx, d, m)
}

func resourceNetworkDeviceBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDeviceBulkRequestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkDeviceBulkRequest delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing NetworkDeviceBulkRequest delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestNetworkDeviceBulkRequestBulkRequestForNetworkDevice(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceBulkRequestForNetworkDevice {
	request := isegosdk.RequestNetworkDeviceBulkRequestForNetworkDevice{}
	request.NetworkDeviceBulkRequest = expandRequestNetworkDeviceBulkRequestBulkRequestForNetworkDeviceNetworkDeviceBulkRequest(ctx, key, d)
	return &request
}

func expandRequestNetworkDeviceBulkRequestBulkRequestForNetworkDeviceNetworkDeviceBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceBulkRequestForNetworkDeviceNetworkDeviceBulkRequest {
	request := isegosdk.RequestNetworkDeviceBulkRequestForNetworkDeviceNetworkDeviceBulkRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation_type")))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_media_type")))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
