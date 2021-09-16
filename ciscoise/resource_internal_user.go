package ciscoise

import (
	"context"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInternalUser() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on InternalUser.
  
  - This resource allows the client to update an internal user by name.
  - This resource deletes an internal user by name.
  - This resource allows the client to update an internal user by ID.
  - This resource deletes an internal user by ID.
  - This resource creates an internal user.`,

		CreateContext: resourceInternalUserCreate,
		ReadContext:   resourceInternalUserRead,
		UpdateContext: resourceInternalUserUpdate,
		DeleteContext: resourceInternalUserDelete,
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
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"change_password": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"custom_attributes": &schema.Schema{
							Description: `Key value map`,
							Type:        schema.TypeMap,
							Optional:    true,
							Computed:    true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enable_password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enabled": &schema.Schema{
							Description: `Whether the user is enabled/disabled. To use it as filter, the values should be 'Enabled' or 'Disabled'.
  The values are case sensitive. For example, '[ERSObjectURL]?filter=enabled.EQ.Enabled'`,
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"expiry_date": &schema.Schema{
							Description: `To store the internal user's expiry date information. It's format is = 'YYYY-MM-DD'`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"expiry_date_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"first_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"identity_groups": &schema.Schema{
							Description: `CSV of identity group IDs`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"last_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"password_idstore": &schema.Schema{
							Description: `The id store where the internal user's password is kept`,
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceInternalUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestInternalUserCreateInternalUser(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.InternalUser.GetInternalUserByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.InternalUser.GetInternalUserByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	restyResp1, err := client.InternalUser.CreateInternalUser(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateInternalUser", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateInternalUser", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceInternalUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetInternalUserByName")
		vvName := vName

		response1, _, err := client.InternalUser.GetInternalUserByName(vvName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetInternalUserByName", err,
				"Failure at GetInternalUserByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItemName1 := flattenInternalUserGetInternalUserByNameItemName(response1.InternalUser)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInternalUserByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetInternalUserByID")
		vvID := vID

		response2, _, err := client.InternalUser.GetInternalUserByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetInternalUserByID", err,
				"Failure at GetInternalUserByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemID2 := flattenInternalUserGetInternalUserByIDItemID(response2.InternalUser)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInternalUserByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceInternalUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.InternalUser.GetInternalUserByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetInternalUserByName", err,
				"Failure at GetInternalUserByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.InternalUser != nil {
			vvID = getResp.InternalUser.ID
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestInternalUserUpdateInternalUserByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.InternalUser.UpdateInternalUserByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateInternalUserByID", err, restyResp1.String(),
					"Failure at UpdateInternalUserByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateInternalUserByID", err,
				"Failure at UpdateInternalUserByID, unexpected response", ""))
			return diags
		}
	}

	return resourceInternalUserRead(ctx, d, m)
}

func resourceInternalUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.InternalUser.GetInternalUserByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.InternalUser.GetInternalUserByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.InternalUser != nil {
			vvID = getResp.InternalUser.ID
		}
	}
	restyResp1, err := client.InternalUser.DeleteInternalUserByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteInternalUserByID", err, restyResp1.String(),
				"Failure at DeleteInternalUserByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteInternalUserByID", err,
			"Failure at DeleteInternalUserByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestInternalUserCreateInternalUser(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestInternalUserCreateInternalUser {
	request := isegosdk.RequestInternalUserCreateInternalUser{}
	request.InternalUser = expandRequestInternalUserCreateInternalUserInternalUser(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestInternalUserCreateInternalUserInternalUser(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestInternalUserCreateInternalUserInternalUser {
	request := isegosdk.RequestInternalUserCreateInternalUserInternalUser{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".email"); !isEmptyValue(reflect.ValueOf(d.Get(key+".email"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".email"))) {
		request.Email = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".first_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".first_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".first_name"))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".last_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".last_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".last_name"))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".change_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".change_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".change_password"))) {
		request.ChangePassword = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".identity_groups"); !isEmptyValue(reflect.ValueOf(d.Get(key+".identity_groups"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".identity_groups"))) {
		request.IDentityGroups = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".expiry_date_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".expiry_date_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".expiry_date_enabled"))) {
		request.ExpiryDateEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".expiry_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".expiry_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".expiry_date"))) {
		request.ExpiryDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_password"))) {
		request.EnablePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".custom_attributes"); !isEmptyValue(reflect.ValueOf(d.Get(key+".custom_attributes"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".custom_attributes"))) {
		customAttributes := v.([]interface{})[0].(map[string]interface{})
		request.CustomAttributes = &customAttributes
	}
	if v, ok := d.GetOkExists(key + ".password_idstore"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password_idstore"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password_idstore"))) {
		request.PasswordIDStore = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestInternalUserUpdateInternalUserByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestInternalUserUpdateInternalUserByID {
	request := isegosdk.RequestInternalUserUpdateInternalUserByID{}
	request.InternalUser = expandRequestInternalUserUpdateInternalUserByIDInternalUser(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestInternalUserUpdateInternalUserByIDInternalUser(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestInternalUserUpdateInternalUserByIDInternalUser {
	request := isegosdk.RequestInternalUserUpdateInternalUserByIDInternalUser{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".email"); !isEmptyValue(reflect.ValueOf(d.Get(key+".email"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".email"))) {
		request.Email = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".first_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".first_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".first_name"))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".last_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".last_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".last_name"))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".change_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".change_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".change_password"))) {
		request.ChangePassword = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".identity_groups"); !isEmptyValue(reflect.ValueOf(d.Get(key+".identity_groups"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".identity_groups"))) {
		request.IDentityGroups = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".expiry_date_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".expiry_date_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".expiry_date_enabled"))) {
		request.ExpiryDateEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".expiry_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".expiry_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".expiry_date"))) {
		request.ExpiryDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_password"))) {
		request.EnablePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".custom_attributes"); !isEmptyValue(reflect.ValueOf(d.Get(key+".custom_attributes"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".custom_attributes"))) {
		customAttributes := v.([]interface{})[0].(map[string]interface{})
		request.CustomAttributes = &customAttributes
	}
	if v, ok := d.GetOkExists(key + ".password_idstore"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password_idstore"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password_idstore"))) {
		request.PasswordIDStore = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
