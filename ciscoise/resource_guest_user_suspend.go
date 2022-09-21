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

func resourceGuestUserSuspend() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on GuestUser.
- This resource allows the client to suspend a guest user by name.
- This resource allows the client to suspend a guest user by ID.
`,

		CreateContext: resourceGuestUserSuspendCreate,
		ReadContext:   resourceGuestUserSuspendRead,
		DeleteContext: resourceGuestUserSuspendDelete,

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
							Optional:    true,
							ForceNew:    true,
						},
						"name": &schema.Schema{
							Description: `name path parameter.`,
							Type:        schema.TypeString,
							Optional:    true,
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

func resourceGuestUserSuspendCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SuspendGuestUserByName create")
	log.Printf("[DEBUG] Missing SuspendGuestUserByName create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vName, okName := d.GetOk("parameters.0.name")
	vID, okID := d.GetOk("parameters.0.id")

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: SuspendGuestUserByName")
		vvName := vName.(string)

		response1, err := client.GuestUser.SuspendGuestUserByName(vvName)

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SuspendGuestUserByName", err, response1.String(),
					"Failure at SuspendGuestUserByName, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SuspendGuestUserByName", err,
				"Failure at SuspendGuestUserByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SuspendGuestUserByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		_ = d.Set("last_updated", getUnixTimeString())

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: SuspendGuestUserByID")
		vvID := vID.(string)
		request2 := expandRequestGuestUserSuspendSuspendGuestUserByID(ctx, "parameters.0", d)

		response2, err := client.GuestUser.SuspendGuestUserByID(vvID, request2)

		if request2 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request2))
		}

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SuspendGuestUserByID", err,
				"Failure at SuspendGuestUserByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response2.String())

		if err := d.Set("item", response2.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SuspendGuestUserByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		_ = d.Set("last_updated", getUnixTimeString())

	}
	return resourceGuestUserSuspendRead(ctx, d, m)
}

func resourceGuestUserSuspendRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceGuestUserSuspendDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning GuestUserSuspend delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing GuestUserSuspend delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestGuestUserSuspendSuspendGuestUserByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserSuspendGuestUserByID {
	request := isegosdk.RequestGuestUserSuspendGuestUserByID{}
	request.OperationAdditionalData = expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalData(ctx, key, d)
	return &request
}

func expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalData {
	request := isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_data")))) {
		request.AdditionalData = expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalDataAdditionalDataArray(ctx, key+".additional_data", d)
	}
	return &request
}

func expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalDataAdditionalDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalDataAdditionalData {
	request := []isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalDataAdditionalData{}
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
		i := expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalDataAdditionalData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestGuestUserSuspendSuspendGuestUserByIDOperationAdditionalDataAdditionalData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalDataAdditionalData {
	request := isegosdk.RequestGuestUserSuspendGuestUserByIDOperationAdditionalDataAdditionalData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
