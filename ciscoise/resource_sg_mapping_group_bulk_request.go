package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSgMappingGroupBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMappingGroup.
- This resource allows the client to submit the bulk request.
`,

		CreateContext: resourceSgMappingGroupBulkRequestCreate,
		ReadContext:   resourceSgMappingGroupBulkRequestRead,
		DeleteContext: resourceSgMappingGroupBulkRequestDelete,

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

func resourceSgMappingGroupBulkRequestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkRequestForIPToSgtMappingGroup create")
	log.Printf("[DEBUG] Missing BulkRequestForIPToSgtMappingGroup create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	request1 := expandRequestSgMappingGroupBulkRequestBulkRequestForIPToSgtMappingGroup(ctx, "parameters.0", d)
	response1, err := client.IPToSgtMappingGroup.BulkRequestForIPToSgtMappingGroup(request1)
	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing BulkRequestForIPToSgtMappingGroup", err, response1.String(),
				"Failure at BulkRequestForIPToSgtMappingGroup, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkRequestForIPToSgtMappingGroup", err,
			"Failure at BulkRequestForIPToSgtMappingGroup, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %s", response1.String())
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkRequestForIPToSgtMappingGroup response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceSgMappingGroupBulkRequestRead(ctx, d, m)
}

func resourceSgMappingGroupBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSgMappingGroupBulkRequestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMappingGroupBulkRequest delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing SgMappingGroupBulkRequest delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestSgMappingGroupBulkRequestBulkRequestForIPToSgtMappingGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingGroupBulkRequestForIPToSgtMappingGroup {
	request := isegosdk.RequestIPToSgtMappingGroupBulkRequestForIPToSgtMappingGroup{}
	request.SgMappingGroupBulkRequest = expandRequestSgMappingGroupBulkRequestBulkRequestForIPToSgtMappingGroupSgMappingGroupBulkRequest(ctx, key, d)
	return &request
}

func expandRequestSgMappingGroupBulkRequestBulkRequestForIPToSgtMappingGroupSgMappingGroupBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingGroupBulkRequestForIPToSgtMappingGroupSgMappingGroupBulkRequest {
	request := isegosdk.RequestIPToSgtMappingGroupBulkRequestForIPToSgtMappingGroupSgMappingGroupBulkRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation_type")))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_media_type")))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
