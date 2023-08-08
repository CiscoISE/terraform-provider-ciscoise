package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEndpointBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on endpoint.
- This resource allows the client to submit the bulk request.
`,

		CreateContext: resourceEndpointBulkRequestCreate,
		ReadContext:   resourceEndpointBulkRequestRead,
		DeleteContext: resourceEndpointBulkRequestDelete,

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

func resourceEndpointBulkRequestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkRequestForEndpoint create")
	log.Printf("[DEBUG] Missing BulkRequestForEndpoint create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	request1 := expandRequestEndpointBulkRequestBulkRequestForEndpoint(ctx, "parameters.0", d)

	response1, err := client.Endpoint.BulkRequestForEndpoint(request1)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkRequestForEndpoint", err,
			"Failure at BulkRequestForEndpoint, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkRequestForEndpoint response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return resourceEndpointBulkRequestRead(ctx, d, m)
}

func resourceEndpointBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceEndpointBulkRequestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning EndpointBulkRequest delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing EndpointBulkRequest delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestEndpointBulkRequestBulkRequestForEndpoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointBulkRequestForEndpoint {
	request := isegosdk.RequestEndpointBulkRequestForEndpoint{}
	request.EndpointBulkRequest = expandRequestEndpointBulkRequestBulkRequestForEndpointEndpointBulkRequest(ctx, key, d)
	return &request
}

func expandRequestEndpointBulkRequestBulkRequestForEndpointEndpointBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointBulkRequestForEndpointEndpointBulkRequest {
	request := isegosdk.RequestEndpointBulkRequestForEndpointEndpointBulkRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation_type")))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_media_type")))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
