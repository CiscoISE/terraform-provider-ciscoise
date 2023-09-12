package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceActiveDirectoryLeaveDomain() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ActiveDirectory.
- This resource makes a Cisco ISE node to leave an Active Directory domain.
`,

		CreateContext: resourceActiveDirectoryLeaveDomainCreate,
		ReadContext:   resourceActiveDirectoryLeaveDomainRead,
		DeleteContext: resourceActiveDirectoryLeaveDomainDelete,

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
							Required:    true,
							ForceNew:    true,
						},
						"additional_data": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
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

func resourceActiveDirectoryLeaveDomainCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning LeaveDomain create")
	log.Printf("[DEBUG] Missing LeaveDomain create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	vID := d.Get("parameters.0.id")
	var diags diag.Diagnostics

	request1 := expandRequestActiveDirectoryLeaveDomainLeaveDomain(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vvID := vID.(string)
	response1, err := client.ActiveDirectory.LeaveDomain(vvID, request1)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing LeaveDomain", err,
			"Failure at LeaveDomain, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting LeaveDomain response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceActiveDirectoryLeaveDomainRead(ctx, d, m)
}

func resourceActiveDirectoryLeaveDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceActiveDirectoryLeaveDomainDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ActiveDirectoryLeaveDomain delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing ActiveDirectoryLeaveDomain delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
func expandRequestActiveDirectoryLeaveDomainLeaveDomain(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLeaveDomain {
	request := isegosdk.RequestActiveDirectoryLeaveDomain{}
	request.OperationAdditionalData = expandRequestActiveDirectoryLeaveDomainLeaveDomainOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestActiveDirectoryLeaveDomainLeaveDomainOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLeaveDomainOperationAdditionalData {
	request := isegosdk.RequestActiveDirectoryLeaveDomainOperationAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_data")))) {
		request.AdditionalData = expandRequestActiveDirectoryLeaveDomainLeaveDomainOperationAdditionalDataAdditionalDataArray(ctx, key+".additional_data", d)
	}
	return &request
}

func expandRequestActiveDirectoryLeaveDomainLeaveDomainOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestActiveDirectoryLeaveDomainOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestActiveDirectoryLeaveDomainOperationAdditionalDataAdditionalData{}
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
		i := expandRequestActiveDirectoryLeaveDomainLeaveDomainOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestActiveDirectoryLeaveDomainLeaveDomainOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestActiveDirectoryLeaveDomainOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestActiveDirectoryLeaveDomainOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
