package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSgACLBulkRequest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on SecurityGroupsACLs.

- This data source action allows the client to submit the bulk request.
`,

		ReadContext: dataSourceSgACLBulkRequestRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"operation_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_media_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceSgACLBulkRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BulkRequestForSecurityGroupsACL")
		request1 := expandRequestSgACLBulkRequestBulkRequestForSecurityGroupsACL(ctx, "", d)

		response1, err := client.SecurityGroupsACLs.BulkRequestForSecurityGroupsACL(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if err != nil || response1 == nil {
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
		d.SetId(getUnixTimeString())
		return diags

	}
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
