package ciscoise

import (
	"context"

	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceEndpointRegister() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEndpointRegisterRead,
		Schema: map[string]*schema.Schema{
			"custom_attributes": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"identity_store": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"identity_store_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mdm_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"mdm_compliance_status": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"mdm_encrypted": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"mdm_enrolled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"mdm_ime_i": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mdm_jail_broken": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"mdm_manufacturer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mdm_model": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mdm_os": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mdm_phone_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mdm_pinlock": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"mdm_reachable": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"mdm_serial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mdm_server_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"portal_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"static_group_assignment": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"static_profile_assignment": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func dataSourceEndpointRegisterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RegisterEndpoint")
		request1 := expandRequestEndpointRegisterRegisterEndpoint(ctx, "", d)

		response1, err := client.Endpoint.RegisterEndpoint(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RegisterEndpoint", err,
				"Failure at RegisterEndpoint, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RegisterEndpoint response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestEndpointRegisterRegisterEndpoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointRegisterEndpoint {
	request := isegosdk.RequestEndpointRegisterEndpoint{}
	request.ERSEndPoint = expandRequestEndpointRegisterRegisterEndpointERSEndPoint(ctx, key, d)
	return &request
}

func expandRequestEndpointRegisterRegisterEndpointERSEndPoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointRegisterEndpointERSEndPoint {
	request := isegosdk.RequestEndpointRegisterEndpointERSEndPoint{}
	if v, ok := d.GetOkExists("id"); !isEmptyValue(reflect.ValueOf(d.Get("id"))) && (ok || !reflect.DeepEqual(v, d.Get("id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(d.Get("name"))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(d.Get("description"))) && (ok || !reflect.DeepEqual(v, d.Get("description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("mac"); !isEmptyValue(reflect.ValueOf(d.Get("mac"))) && (ok || !reflect.DeepEqual(v, d.Get("mac"))) {
		request.Mac = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("profile_id"); !isEmptyValue(reflect.ValueOf(d.Get("profile_id"))) && (ok || !reflect.DeepEqual(v, d.Get("profile_id"))) {
		request.ProfileID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("static_profile_assignment"); !isEmptyValue(reflect.ValueOf(d.Get("static_profile_assignment"))) && (ok || !reflect.DeepEqual(v, d.Get("static_profile_assignment"))) {
		request.StaticProfileAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("group_id"); !isEmptyValue(reflect.ValueOf(d.Get("group_id"))) && (ok || !reflect.DeepEqual(v, d.Get("group_id"))) {
		request.GroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("static_group_assignment"); !isEmptyValue(reflect.ValueOf(d.Get("static_group_assignment"))) && (ok || !reflect.DeepEqual(v, d.Get("static_group_assignment"))) {
		request.StaticGroupAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("portal_user"); !isEmptyValue(reflect.ValueOf(d.Get("portal_user"))) && (ok || !reflect.DeepEqual(v, d.Get("portal_user"))) {
		request.PortalUser = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("identity_store"); !isEmptyValue(reflect.ValueOf(d.Get("identity_store"))) && (ok || !reflect.DeepEqual(v, d.Get("identity_store"))) {
		request.IDentityStore = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("identity_store_id"); !isEmptyValue(reflect.ValueOf(d.Get("identity_store_id"))) && (ok || !reflect.DeepEqual(v, d.Get("identity_store_id"))) {
		request.IDentityStoreID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("mdm_attributes"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_attributes"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_attributes"))) {
		request.MdmAttributes = expandRequestEndpointRegisterRegisterEndpointERSEndPointMdmAttributes(ctx, key+".mdm_attributes.0", d)
	}
	if v, ok := d.GetOkExists("custom_attributes"); !isEmptyValue(reflect.ValueOf(d.Get("custom_attributes"))) && (ok || !reflect.DeepEqual(v, d.Get("custom_attributes"))) {
		customAttributes := v.([]interface{})[0].(map[string]interface{})
		request.CustomAttributes = &customAttributes
	}
	return &request
}

func expandRequestEndpointRegisterRegisterEndpointERSEndPointMdmAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointRegisterEndpointERSEndPointMdmAttributes {
	request := isegosdk.RequestEndpointRegisterEndpointERSEndPointMdmAttributes{}
	if v, ok := d.GetOkExists("mdm_server_name"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_server_name"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_server_name"))) {
		request.MdmServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("mdm_reachable"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_reachable"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_reachable"))) {
		request.MdmReachable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("mdm_enrolled"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_enrolled"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_enrolled"))) {
		request.MdmEnrolled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("mdm_compliance_status"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_compliance_status"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_compliance_status"))) {
		request.MdmComplianceStatus = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("mdm_os"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_os"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_os"))) {
		request.MdmOS = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("mdm_manufacturer"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_manufacturer"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_manufacturer"))) {
		request.MdmManufacturer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("mdm_model"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_model"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_model"))) {
		request.MdmModel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("mdm_serial"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_serial"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_serial"))) {
		request.MdmSerial = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("mdm_encrypted"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_encrypted"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_encrypted"))) {
		request.MdmEncrypted = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("mdm_pinlock"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_pinlock"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_pinlock"))) {
		request.MdmPinlock = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("mdm_jail_broken"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_jail_broken"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_jail_broken"))) {
		request.MdmJailBroken = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("mdm_ime_i"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_ime_i"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_ime_i"))) {
		request.MdmIMEI = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("mdm_phone_number"); !isEmptyValue(reflect.ValueOf(d.Get("mdm_phone_number"))) && (ok || !reflect.DeepEqual(v, d.Get("mdm_phone_number"))) {
		request.MdmPhoneNumber = interfaceToString(v)
	}
	return &request
}
