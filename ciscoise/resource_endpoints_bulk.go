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

func resourceEndpointsBulk() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on endpoints.

- Create Endpoint

- Update Endpoint in bulk

- Create Endpoint in bulk

- Delete Endpoint in bulk

- Update Endpoint by id or mac

- Delete endpoint by id or mac
`,

		CreateContext: resourceEndpointsBulkCreate,
		ReadContext:   resourceEndpointsBulkRead,
		UpdateContext: resourceEndpointsBulkUpdate,
		DeleteContext: resourceEndpointsBulkDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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
				Description: `Array of RequestEndpointsCreateBulkEndPoints`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connected_links": &schema.Schema{
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"custom_attributes": &schema.Schema{
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"description": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"device_type": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"group_id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"hardware_revision": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"identity_store": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"identity_store_id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"ip_address": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"mac": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"mdm_attributes": &schema.Schema{
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"name": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"portal_user": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"product_id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"profile_id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"protocol": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"serial_number": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"software_revision": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"static_group_assignment": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"static_profile_assignment": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"value": &schema.Schema{
							Description:      `value path parameter. The id or MAC of the endpoint`,
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: diffSupressOptional(),
						},
						"vendor": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func resourceEndpointsBulkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	// resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestEndpointsBulkCreateBulkEndPoints(ctx, "parameters", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	resp1, restyResp1, err := client.Endpoints.CreateBulkEndPoints(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateBulkEndPoints", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateBulkEndPoints", err))
		return diags
	}

	vItem1 := flattenEndpointsBulkItem(resp1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetAciSettings response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags
}

func resourceEndpointsBulkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceEndpointsBulkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceEndpointsBulkRead(ctx, d, m)
}

func resourceEndpointsBulkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenEndpointsBulkItem(item *isegosdk.ResponseEndpointsCreateBulkEndPoints) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}

func expandRequestEndpointsBulkCreateBulkEndPoints(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointsCreateBulkEndPoints {
	request := isegosdk.RequestEndpointsCreateBulkEndPoints{}
	if v := expandRequestEndpointsBulkCreateBulkEndPointsItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointsBulkCreateBulkEndPointsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestItemEndpointsCreateBulkEndPoints {
	request := []isegosdk.RequestItemEndpointsCreateBulkEndPoints{}
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
		i := expandRequestEndpointsBulkCreateBulkEndPointsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointsBulkCreateBulkEndPointsItem(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemEndpointsCreateBulkEndPoints {
	request := isegosdk.RequestItemEndpointsCreateBulkEndPoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_links")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_links")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_links")))) {
		request.ConnectedLinks = expandRequestEndpointsBulkCreateBulkEndPointsItemConnectedLinks(ctx, key+".connected_links.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_attributes")))) {
		request.CustomAttributes = expandRequestEndpointsBulkCreateBulkEndPointsItemCustomAttributes(ctx, key+".custom_attributes.0", d)
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
		request.MdmAttributes = expandRequestEndpointsBulkCreateBulkEndPointsItemMdmAttributes(ctx, key+".mdm_attributes.0", d)
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
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointsBulkCreateBulkEndPointsItemConnectedLinks(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemEndpointsCreateBulkEndPointsConnectedLinks {
	var request isegosdk.RequestItemEndpointsCreateBulkEndPointsConnectedLinks
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointsBulkCreateBulkEndPointsItemCustomAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemEndpointsCreateBulkEndPointsCustomAttributes {
	var request isegosdk.RequestItemEndpointsCreateBulkEndPointsCustomAttributes
	request, _ = d.Get(fixKeyAccess(key)).(isegosdk.RequestItemEndpointsCreateBulkEndPointsCustomAttributes)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointsBulkCreateBulkEndPointsItemMdmAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemEndpointsCreateBulkEndPointsMdmAttributes {
	var request isegosdk.RequestItemEndpointsCreateBulkEndPointsMdmAttributes
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointsBulkUpdateEndpoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointsUpdateEndpoint {
	request := isegosdk.RequestEndpointsUpdateEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_links")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_links")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_links")))) {
		request.ConnectedLinks = expandRequestEndpointsBulkUpdateEndpointConnectedLinks(ctx, key+".connected_links.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_attributes")))) {
		request.CustomAttributes = expandRequestEndpointsBulkUpdateEndpointCustomAttributes(ctx, key+".custom_attributes.0", d)
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
		request.MdmAttributes = expandRequestEndpointsBulkUpdateEndpointMdmAttributes(ctx, key+".mdm_attributes.0", d)
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
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointsBulkUpdateEndpointConnectedLinks(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointsUpdateEndpointConnectedLinks {
	var request isegosdk.RequestEndpointsUpdateEndpointConnectedLinks
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointsBulkUpdateEndpointCustomAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointsUpdateEndpointCustomAttributes {
	var request isegosdk.RequestEndpointsUpdateEndpointCustomAttributes
	request, _ = d.Get(fixKeyAccess(key)).(isegosdk.RequestEndpointsUpdateEndpointCustomAttributes)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointsBulkUpdateEndpointMdmAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointsUpdateEndpointMdmAttributes {
	var request isegosdk.RequestEndpointsUpdateEndpointMdmAttributes
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
