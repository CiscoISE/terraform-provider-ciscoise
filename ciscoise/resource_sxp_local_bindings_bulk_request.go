package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSxpLocalBindingsBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on SXPLocalBindings.
- This resource allows the client to submit the bulk request.
`,

		CreateContext: resourceSxpLocalBindingsBulkRequestCreate,
		ReadContext:   resourceSxpLocalBindingsBulkRequestRead,
		DeleteContext: resourceSxpLocalBindingsBulkRequestDelete,

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

func resourceSxpLocalBindingsBulkRequestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkRequestForSxpLocalBindings create")
	log.Printf("[DEBUG] Missing BulkRequestForSxpLocalBindings create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	request1 := expandRequestSxpLocalBindingsBulkRequestBulkRequestForSxpLocalBindings(ctx, "parameters.0", d)

	response1, err := client.SxpLocalBindings.BulkRequestForSxpLocalBindings(request1)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkRequestForSxpLocalBindings", err,
			"Failure at BulkRequestForSxpLocalBindings, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkRequestForSxpLocalBindings response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return resourceSxpLocalBindingsBulkRequestRead(ctx, d, m)
}

func resourceSxpLocalBindingsBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSxpLocalBindingsBulkRequestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SxpLocalBindingsBulkRequest delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing SxpLocalBindingsBulkRequest delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestSxpLocalBindingsBulkRequestBulkRequestForSxpLocalBindings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpLocalBindingsBulkRequestForSxpLocalBindings {
	request := isegosdk.RequestSxpLocalBindingsBulkRequestForSxpLocalBindings{}
	request.LocalBindingBulkRequest = expandRequestSxpLocalBindingsBulkRequestBulkRequestForSxpLocalBindingsLocalBindingBulkRequest(ctx, key, d)
	return &request
}

func expandRequestSxpLocalBindingsBulkRequestBulkRequestForSxpLocalBindingsLocalBindingBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSxpLocalBindingsBulkRequestForSxpLocalBindingsLocalBindingBulkRequest {
	request := isegosdk.RequestSxpLocalBindingsBulkRequestForSxpLocalBindingsLocalBindingBulkRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation_type")))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_media_type")))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
