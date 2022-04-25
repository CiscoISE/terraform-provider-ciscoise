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

func resourceTrustsecVnBulkCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on virtualNetwork.
- Create Virtual Network in bulk
`,

		CreateContext: resourceTrustsecVnBulkCreateCreate,
		ReadContext:   resourceTrustsecVnBulkCreateRead,
		DeleteContext: resourceTrustsecVnBulkCreateDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestVirtualNetworkBulkCreateVirtualNetworks`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"additional_attributes": &schema.Schema{
										Description: `JSON String of additional attributes for the Virtual Network`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"id": &schema.Schema{
										Description: `Identifier of the Virtual Network`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"last_update": &schema.Schema{
										Description: `Timestamp for the last update of the Virtual Network`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"name": &schema.Schema{
										Description: `Name of the Virtual Network`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
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

func resourceTrustsecVnBulkCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkCreateVirtualNetworks create")
	log.Printf("[DEBUG] Missing BulkCreateVirtualNetworks create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	request1 := expandRequestTrustsecVnBulkCreateBulkCreateVirtualNetworks(ctx, "parameters.0", d)

	response1, restyResp1, err := client.VirtualNetwork.BulkCreateVirtualNetworks(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkCreateVirtualNetworks", err,
			"Failure at BulkCreateVirtualNetworks, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenVirtualNetworkBulkCreateVirtualNetworksItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkCreateVirtualNetworks response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return resourceTrustsecVnBulkCreateRead(ctx, d, m)
}

func resourceTrustsecVnBulkCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceTrustsecVnBulkCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVnBulkCreate delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing TrustsecVnBulkCreate delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestTrustsecVnBulkCreateBulkCreateVirtualNetworks(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestVirtualNetworkBulkCreateVirtualNetworks {
	request := isegosdk.RequestVirtualNetworkBulkCreateVirtualNetworks{}
	if v := expandRequestTrustsecVnBulkCreateBulkCreateVirtualNetworksItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestTrustsecVnBulkCreateBulkCreateVirtualNetworksItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestItemVirtualNetworkBulkCreateVirtualNetworks {
	request := []isegosdk.RequestItemVirtualNetworkBulkCreateVirtualNetworks{}
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
		i := expandRequestTrustsecVnBulkCreateBulkCreateVirtualNetworksItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestTrustsecVnBulkCreateBulkCreateVirtualNetworksItem(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemVirtualNetworkBulkCreateVirtualNetworks {
	request := isegosdk.RequestItemVirtualNetworkBulkCreateVirtualNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_attributes")))) {
		request.AdditionalAttributes = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}

func flattenVirtualNetworkBulkCreateVirtualNetworksItem(item *isegosdk.ResponseVirtualNetworkBulkCreateVirtualNetworks) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
