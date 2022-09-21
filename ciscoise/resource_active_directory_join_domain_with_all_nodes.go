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

func resourceActiveDirectoryJoinDomainWithAllNodes() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ActiveDirectory.
- This resource joins all Cisco ISE Nodes to an Active Directory domain.
`,

		CreateContext: resourceActiveDirectoryJoinDomainWithAllNodesCreate,
		ReadContext:   resourceActiveDirectoryJoinDomainWithAllNodesRead,
		DeleteContext: resourceActiveDirectoryJoinDomainWithAllNodesDelete,

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
						"id": &schema.Schema{
							Description: `id path parameter.`,
							Type:        schema.TypeString,
							ForceNew:    true,
							Required:    true,
						},
						"additional_data": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										ForceNew: true,
										Optional: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										ForceNew: true,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceActiveDirectoryJoinDomainWithAllNodesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning JoinDomainWithAllNodes create")
	log.Printf("[DEBUG] Missing JoinDomainWithAllNodes create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vID := resourceItem["id"]
	vvID := vID.(string)
	request1 := expandRequestActiveDirectoryJoinDomainWithAllNodesJoinDomainWithAllNodes(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, err := client.ActiveDirectory.JoinDomainWithAllNodes(vvID, request1)
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
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceActiveDirectoryJoinDomainWithAllNodesRead(ctx, d, m)
}

func resourceActiveDirectoryJoinDomainWithAllNodesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceActiveDirectoryJoinDomainWithAllNodesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ActiveDirectoryJoinDomainWithAllNodes delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing ActiveDirectoryJoinDomainWithAllNodes delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
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
