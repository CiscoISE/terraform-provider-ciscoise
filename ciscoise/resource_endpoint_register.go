package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEndpointRegister() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on endpoint.
- This data source action allows the client to register an endpoint.
`,

		CreateContext: resourceEndpointRegisterCreate,
		ReadContext:   resourceEndpointRegisterRead,
		DeleteContext: resourceEndpointRegisterDelete,

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
						"custom_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_attributes": &schema.Schema{
										Description: `Key value map`,
										Type:        schema.TypeMap,
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"group_id": &schema.Schema{
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
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"mdm_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"mdm_compliance_status": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
									},
									"mdm_encrypted": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
									},
									"mdm_enrolled": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
									},
									"mdm_ime_i": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"mdm_jail_broken": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
									},
									"mdm_manufacturer": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"mdm_model": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"mdm_os": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"mdm_phone_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"mdm_pinlock": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
									},
									"mdm_reachable": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
									},
									"mdm_serial": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"mdm_server_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
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
						"profile_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"static_group_assignment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"static_profile_assignment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
					},
				},
			},
		},
	}
}

func resourceEndpointRegisterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning RegisterEndpoint create")
	log.Printf("[DEBUG] Missing RegisterEndpoint create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	request1 := expandRequestEndpointRegisterRegisterEndpoint(ctx, "parameters", d)

	response1, err := client.Endpoint.RegisterEndpoint(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RegisterEndpoint", err,
			"Failure at RegisterEndpoint, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RegisterEndpoint response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceEndpointRegisterRead(ctx, d, m)
}

func resourceEndpointRegisterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceEndpointRegisterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning EndpointRegister delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing EndpointRegister delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestEndpointRegisterRegisterEndpoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointRegisterEndpoint {
	request := isegosdk.RequestEndpointRegisterEndpoint{}
	request.ERSEndPoint = expandRequestEndpointRegisterRegisterEndpointERSEndPoint(ctx, key, d)
	return &request
}

func expandRequestEndpointRegisterRegisterEndpointERSEndPoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointRegisterEndpointERSEndPoint {
	request := isegosdk.RequestEndpointRegisterEndpointERSEndPoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac")))) {
		request.Mac = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_id")))) {
		request.ProfileID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".static_profile_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".static_profile_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".static_profile_assignment")))) {
		request.StaticProfileAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_id")))) {
		request.GroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".static_group_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".static_group_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".static_group_assignment")))) {
		request.StaticGroupAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_user")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_user")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_user")))) {
		request.PortalUser = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_store")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_store")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_store")))) {
		request.IDentityStore = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_store_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_store_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_store_id")))) {
		request.IDentityStoreID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_attributes")))) {
		request.MdmAttributes = expandRequestEndpointRegisterRegisterEndpointERSEndPointMdmAttributes(ctx, key+".mdm_attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_attributes")))) {
		request.CustomAttributes = expandRequestEndpointRegisterRegisterEndpointERSEndPointCustomAttributes(ctx, key+".custom_attributes.0", d)
	}
	return &request
}

func expandRequestEndpointRegisterRegisterEndpointERSEndPointMdmAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointRegisterEndpointERSEndPointMdmAttributes {
	request := isegosdk.RequestEndpointRegisterEndpointERSEndPointMdmAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_server_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_server_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_server_name")))) {
		request.MdmServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_reachable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_reachable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_reachable")))) {
		request.MdmReachable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_enrolled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_enrolled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_enrolled")))) {
		request.MdmEnrolled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_compliance_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_compliance_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_compliance_status")))) {
		request.MdmComplianceStatus = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_os")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_os")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_os")))) {
		request.MdmOS = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_manufacturer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_manufacturer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_manufacturer")))) {
		request.MdmManufacturer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_model")))) {
		request.MdmModel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_serial")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_serial")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_serial")))) {
		request.MdmSerial = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_encrypted")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_encrypted")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_encrypted")))) {
		request.MdmEncrypted = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_pinlock")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_pinlock")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_pinlock")))) {
		request.MdmPinlock = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_jail_broken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_jail_broken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_jail_broken")))) {
		request.MdmJailBroken = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_ime_i")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_ime_i")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_ime_i")))) {
		request.MdmIMEI = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_phone_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_phone_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_phone_number")))) {
		request.MdmPhoneNumber = interfaceToString(v)
	}
	return &request
}

func expandRequestEndpointRegisterRegisterEndpointERSEndPointCustomAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointRegisterEndpointERSEndPointCustomAttributes {
	request := isegosdk.RequestEndpointRegisterEndpointERSEndPointCustomAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_attributes")))) {
		request.CustomAttributes = expandRequestEndpointRegisterRegisterEndpointERSEndPointCustomAttributesCustomAttributes(ctx, key+".custom_attributes.0", d)
	}
	return &request
}

func expandRequestEndpointRegisterRegisterEndpointERSEndPointCustomAttributesCustomAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointRegisterEndpointERSEndPointCustomAttributesCustomAttributes {
	var request isegosdk.RequestEndpointRegisterEndpointERSEndPointCustomAttributesCustomAttributes
	v := d.Get(fixKeyAccess(key))
	request = v.(map[string]interface{})
	return &request
}
