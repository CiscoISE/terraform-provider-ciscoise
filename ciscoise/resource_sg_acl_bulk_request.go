package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSgACLBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on SecurityGroupsACLs.
- This resource allows the client to submit the bulk request.
`,

		CreateContext: resourceSgACLBulkRequestCreate,
		ReadContext:   resourceSgACLBulkRequestRead,
		DeleteContext: resourceSgACLBulkRequestDelete,

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

func resourceSgACLBulkRequestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkRequestForSecurityGroupsACL create")
	log.Printf("[DEBUG] Missing BulkRequestForSecurityGroupsACL create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	request1 := expandRequestSgACLBulkRequestBulkRequestForSecurityGroupsACL(ctx, "parameters.0", d)

	response1, err := client.SecurityGroupsACLs.BulkRequestForSecurityGroupsACL(request1)
	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing BulkRequestForSecurityGroupsACL", err, response1.String(),
				"Failure at BulkRequestForSecurityGroupsACL, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkRequestForSecurityGroupsACL", err,
			"Failure at BulkRequestForSecurityGroupsACL, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %s", response1.String())
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkRequestForSecurityGroupsACL response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceSgACLBulkRequestRead(ctx, d, m)
}

func resourceSgACLBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSgACLBulkRequestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgACLBulkRequest delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing SgACLBulkRequest delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestSgACLBulkRequestBulkRequestForSecurityGroupsACL(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsACLsBulkRequestForSecurityGroupsACL {
	request := isegosdk.RequestSecurityGroupsACLsBulkRequestForSecurityGroupsACL{}
	request.SgaclBulkRequest = expandRequestSgACLBulkRequestBulkRequestForSecurityGroupsACLSgaclBulkRequest(ctx, key, d)
	return &request
}

func expandRequestSgACLBulkRequestBulkRequestForSecurityGroupsACLSgaclBulkRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsACLsBulkRequestForSecurityGroupsACLSgaclBulkRequest {
	request := isegosdk.RequestSecurityGroupsACLsBulkRequestForSecurityGroupsACLSgaclBulkRequest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation_type")))) {
		request.OperationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_media_type")))) {
		request.ResourceMediaType = interfaceToString(v)
	}
	return &request
}
