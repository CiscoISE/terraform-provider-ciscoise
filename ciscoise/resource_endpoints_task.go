package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceEndpointsTask() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on endpoints.

- Create Endpoint task
`,

		CreateContext: resourceEndpointsTaskCreate,
		ReadContext:   resourceEndpointsTaskRead,
		DeleteContext: resourceEndpointsTaskDelete,
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
						"connected_links": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							ForceNew: true,
						},
						"custom_attributes": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							ForceNew: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"device_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"group_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"hardware_revision": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"identity_store": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"identity_store_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"mdm_attributes": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							ForceNew: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"portal_user": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"product_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"software_revision": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"static_group_assignment": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"static_profile_assignment": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"vendor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceEndpointsTaskCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestEndpointsTaskCreateEndPointTask(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Endpoints.CreateEndPointTask(request1)

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

	vItem1 := flattenEndpointsCreateEndPointTaskItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CreateEndPointTask response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	return diags
}

func expandRequestEndpointsTaskCreateEndPointTask(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointsCreateEndPointTask {
	request := isegosdk.RequestEndpointsCreateEndPointTask{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_links")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_links")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_links")))) {
		request.ConnectedLinks = expandRequestEndpointsTaskCreateEndPointTaskConnectedLinks(ctx, key+".connected_links.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_attributes")))) {
		request.CustomAttributes = expandRequestEndpointsTaskCreateEndPointTaskCustomAttributes(ctx, key+".custom_attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_id")))) {
		request.GroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_revision")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_revision")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_revision")))) {
		request.HardwareRevision = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_store")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_store")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_store")))) {
		request.IDentityStore = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_store_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_store_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_store_id")))) {
		request.IDentityStoreID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac")))) {
		request.Mac = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_attributes")))) {
		request.MdmAttributes = expandRequestEndpointsTaskCreateEndPointTaskMdmAttributes(ctx, key+".mdm_attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_user")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_user")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_user")))) {
		request.PortalUser = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_id")))) {
		request.ProductID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_id")))) {
		request.ProfileID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_revision")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_revision")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_revision")))) {
		request.SoftwareRevision = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".static_group_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".static_group_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".static_group_assignment")))) {
		request.StaticGroupAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".static_profile_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".static_profile_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".static_profile_assignment")))) {
		request.StaticProfileAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vendor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vendor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vendor")))) {
		request.Vendor = interfaceToString(v)
	}
	return &request
}

func expandRequestEndpointsTaskCreateEndPointTaskConnectedLinks(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointsCreateEndPointTaskConnectedLinks {
	var request isegosdk.RequestEndpointsCreateEndPointTaskConnectedLinks
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestEndpointsTaskCreateEndPointTaskCustomAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointsCreateEndPointTaskCustomAttributes {
	var request isegosdk.RequestEndpointsCreateEndPointTaskCustomAttributes
	request = d.Get(fixKeyAccess(key)).(isegosdk.RequestEndpointsCreateEndPointTaskCustomAttributes)
	return &request
}

func expandRequestEndpointsTaskCreateEndPointTaskMdmAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointsCreateEndPointTaskMdmAttributes {
	var request isegosdk.RequestEndpointsCreateEndPointTaskMdmAttributes
	request = d.Get(fixKeyAccess(key)).(isegosdk.RequestEndpointsCreateEndPointTaskMdmAttributes)
	return &request
}

func flattenEndpointsCreateEndPointTaskItem(item *isegosdk.ResponseEndpointsCreateEndPointTask) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}

func resourceEndpointsTaskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceEndpointsTaskUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceEndpointsTaskRead(ctx, d, m)
}

func resourceEndpointsTaskDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
