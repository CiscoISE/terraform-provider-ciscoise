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
func dataSourceActiveDirectoryJoinDomainWithAllNodes() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ActiveDirectory.

- This data source action joins all Cisco ISE Nodes to an Active Directory domain.
`,

		ReadContext: dataSourceActiveDirectoryJoinDomainWithAllNodesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
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

func dataSourceActiveDirectoryJoinDomainWithAllNodesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: JoinDomainWithAllNodes")
		vvID := vID.(string)
		request1 := expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodes(ctx, "", d)

		response1, err := client.ActiveDirectory.JoinDomainWithAllNodes(vvID, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing JoinDomainWithAllNodes", err, response1.String(),
					"Failure at JoinDomainWithAllNodes, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing JoinDomainWithAllNodes", err,
				"Failure at JoinDomainWithAllNodes, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting JoinDomainWithAllNodes response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryJoinDomainWithAllNodes {
	request := isegosdk.RequestActiveDirectoryJoinDomainWithAllNodes{}
	request.OperationAdditionalData = expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalData {
	request := isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_data")))) {
		request.AdditionalData = expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalDataAdditionalDataArray(ctx, key+".additional_data", d)
	}
	return &request
}

func expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalDataAdditionalData{}
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
		i := expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodesOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestActiveDirectoryJoinDomainWithAllNodesOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
