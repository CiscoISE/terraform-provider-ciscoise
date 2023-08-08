package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAncEndpointBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ANCEndpoint.
- This resource allows the client to submit the bulk request.
`,

		CreateContext: resourceAncEndpointBulkRequestCreate,
		ReadContext:   resourceAncEndpointBulkRequestRead,
		DeleteContext: resourceAncEndpointBulkRequestDelete,

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

func resourceAncEndpointBulkRequestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkRequestForAncEndpoint create")
	log.Printf("[DEBUG] Missing BulkRequestForAncEndpoint create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	request1 := expandRequestAncEndpointBulkRequestBulkRequestForAncEndpoint(ctx, "parameters.0", d)
	response1, err := client.AncEndpoint.BulkRequestForAncEndpoint(request1)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkRequestForAncEndpoint", err,
			"Failure at BulkRequestForAncEndpoint, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkRequestForAncEndpoint response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return resourceAncEndpointBulkRequestRead(ctx, d, m)
}

func resourceAncEndpointBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceAncEndpointBulkRequestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AncEndpointBulkRequest delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing AncEndpointBulkRequest delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestAncEndpointBulkRequestBulkRequestForAncEndpoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncEndpointBulkRequestForAncEndpoint {
	request := isegosdk.RequestAncEndpointBulkRequestForAncEndpoint{}
	request.ErsAncEndpointBulkRequest = expandRequestAncEndpointBulkRequestBulkRequestForAncEndpointErsAncEndpointBulkRequest(ctx, key, d)
	return &request
}

func expandRequestAncEndpointBulkRequestBulkRequestForAncEndpointErsAncEndpointBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAncEndpointBulkRequestForAncEndpointErsAncEndpointBulkRequest {
	request := isegosdk.RequestAncEndpointBulkRequestForAncEndpointErsAncEndpointBulkRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation_type")))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_media_type")))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
