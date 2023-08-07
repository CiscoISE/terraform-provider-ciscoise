package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSgMappingBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMapping.
- This resource allows the client to submit the bulk request.
`,

		CreateContext: resourceSgMappingBulkRequestCreate,
		ReadContext:   resourceSgMappingBulkRequestRead,
		DeleteContext: resourceSgMappingBulkRequestDelete,

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
							ForceNew: true,
							Optional: true,
						},
						"resource_media_type": &schema.Schema{
							Type:     schema.TypeString,
							ForceNew: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSgMappingBulkRequestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkRequestForIPToSgtMapping create")
	log.Printf("[DEBUG] Missing BulkRequestForIPToSgtMapping create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	request1 := expandRequestSgMappingBulkRequestBulkRequestForIPToSgtMapping(ctx, "parameters.0", d)

	response1, err := client.IPToSgtMapping.BulkRequestForIPToSgtMapping(request1)
	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing BulkRequestForIPToSgtMapping", err, response1.String(),
				"Failure at BulkRequestForIPToSgtMapping, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkRequestForIPToSgtMapping", err,
			"Failure at BulkRequestForIPToSgtMapping, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %s", response1.String())
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkRequestForIPToSgtMapping response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceSgMappingBulkRequestRead(ctx, d, m)
}

func resourceSgMappingBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSgMappingBulkRequestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMappingBulkRequest delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing SgMappingBulkRequest delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestSgMappingBulkRequestBulkRequestForIPToSgtMapping(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingBulkRequestForIPToSgtMapping {
	request := isegosdk.RequestIPToSgtMappingBulkRequestForIPToSgtMapping{}
	request.SgMappingBulkRequest = expandRequestSgMappingBulkRequestBulkRequestForIPToSgtMappingSgMappingBulkRequest(ctx, key, d)
	return &request
}

func expandRequestSgMappingBulkRequestBulkRequestForIPToSgtMappingSgMappingBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingBulkRequestForIPToSgtMappingSgMappingBulkRequest {
	request := isegosdk.RequestIPToSgtMappingBulkRequestForIPToSgtMappingSgMappingBulkRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation_type")))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_media_type")))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
