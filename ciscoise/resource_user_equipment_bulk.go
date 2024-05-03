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

// resourceAction
func resourceUserEquipmentBulk() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on User Equipment.

- Create, update and delete multiple user equipments
`,

		CreateContext: resourceUserEquipmentBulkCreate,
		ReadContext:   resourceUserEquipmentBulkRead,
		DeleteContext: resourceUserEquipmentBulkDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
						"x_request_id": &schema.Schema{
							Description: `X-Request-ID header parameter. The Request ID is returned in the response headers and appear in logs`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"device_group": &schema.Schema{
							Description: `Device or Endpoint Group`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"imei": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceUserEquipmentBulkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestUserEquipmentBulkBulkUserEquipmentOperation(ctx, "parameters.0", d)

	response1, restyResp1, err := client.UserEquipment.BulkUserEquipmentOperation(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenUserEquipmentBulkUserEquipmentOperationItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkUserEquipmentOperation response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}

func expandRequestUserEquipmentBulkBulkUserEquipmentOperation(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestUserEquipmentBulkUserEquipmentOperation {
	request := isegosdk.RequestUserEquipmentBulkUserEquipmentOperation{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation")))) {
		request.Operation = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".item_list")))) {
		request.ItemList = expandRequestUserEquipmentBulkBulkUserEquipmentOperationItemListArray(ctx, key+".item_list", d)
	}
	return &request
}

func expandRequestUserEquipmentBulkBulkUserEquipmentOperationItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestUserEquipmentBulkUserEquipmentOperationItemList {
	request := []isegosdk.RequestUserEquipmentBulkUserEquipmentOperationItemList{}
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
		i := expandRequestUserEquipmentBulkBulkUserEquipmentOperationItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestUserEquipmentBulkBulkUserEquipmentOperationItemList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestUserEquipmentBulkUserEquipmentOperationItemList {
	request := isegosdk.RequestUserEquipmentBulkUserEquipmentOperationItemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".imei")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".imei")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".imei")))) {
		request.Imei = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_group")))) {
		request.DeviceGroup = interfaceToString(v)
	}
	return &request
}

func flattenUserEquipmentBulkUserEquipmentOperationItem(item *isegosdk.ResponseUserEquipmentBulkUserEquipmentOperation) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}

func resourceUserEquipmentBulkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceUserEquipmentBulkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceUserEquipmentBulkRead(ctx, d, m)
}

func resourceUserEquipmentBulkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
