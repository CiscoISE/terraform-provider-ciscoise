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

func resourceDuoMfa() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Duo-Mfa.

- Duo-MFA Create a new Duo-MFA configuration

- Duo-MFA Update the Duo-MFA configuration specified in the connectionName.

- Duo-MFA Delete the Duo-MFA configuration specified in the connectionName.
`,

		CreateContext: resourceDuoMfaCreate,
		ReadContext:   resourceDuoMfaRead,
		UpdateContext: resourceDuoMfaUpdate,
		DeleteContext: resourceDuoMfaDelete,
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
						"account_configurations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"admin_api": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ikey": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"s_key": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"api_host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"authentication_api": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ikey": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"connection_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_sync": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"account_configurations": &schema.Schema{
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"connection_name": &schema.Schema{
							Description:      `Name of the Duo-MFA configuration`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"description": &schema.Schema{
							Description:      `Description of the Duo-MFA configuration`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"identity_sync": &schema.Schema{
							Description:      `Name of the Identity Sync configuration`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"response": &schema.Schema{
							Description: `Duo-MFA configuration information`,
							Type:        schema.TypeMap,
							Computed:    true,
						},
						"type": &schema.Schema{
							Description:      `Protocol type for which this Duo-MFA can be used`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceDuoMfaCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := m.(ClientConfig).EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDuoMfaCreateMfa(ctx, "parameters.0", d)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vConnectionName, okConnectionName := resourceItem["connection_name"]
	vvConnectionName := interfaceToString(vConnectionName)
	if isEnableAutoImport {
		if okConnectionName && vvConnectionName != "" {
			getResponse2, _, err := client.DuoMfa.GetMfaByconnectionName(vvConnectionName)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["connection_name"] = vvConnectionName
				d.SetId(joinResourceID(resourceMap))
				return resourceDuoMfaRead(ctx, d, m)
			}
		} else {
			response2, _, err := client.DuoMfa.GetMfa()
			if response2 != nil && err == nil {
				items2 := getAllItemsDuoMfaGetMfa(m, response2)
				item2, err := searchDuoMfaGetMfa(m, items2, vvConnectionName)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["connection_name"] = vvConnectionName
					d.SetId(joinResourceID(resourceMap))
					return resourceDuoMfaRead(ctx, d, m)
				}
			}
		}
	}
	resp1, err := client.DuoMfa.CreateMfa(request1)
	if err != nil || resp1 == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateMfa", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["connection_name"] = vvConnectionName
	d.SetId(joinResourceID(resourceMap))
	return resourceDuoMfaRead(ctx, d, m)
}

func resourceDuoMfaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vConnectionName, okConnectionName := resourceMap["connection_name"]
	vvConnectionName := vConnectionName
	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okConnectionName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetMfa")

		response1, restyResp1, err := client.DuoMfa.GetMfa()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsDuoMfaGetMfa(m, response1)
		item1, err := searchDuoMfaGetMfa(m, items1, vvConnectionName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenDuoMfaGetMfaByconnectionNameItemResponse(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMfa search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMfa search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetMfaByconnectionName")

		response2, restyResp2, err := client.DuoMfa.GetMfaByconnectionName(vvConnectionName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDuoMfaGetMfaByconnectionNameItemResponse(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMfaByconnectionName response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMfaByconnectionName response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceDuoMfaUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vConnectionName, _ := resourceMap["connection_name"]
	vvConnectionName := vConnectionName

	if d.HasChange("parameters") {

		log.Printf("[DEBUG] Name used for update operation %s", vvConnectionName)

		request1 := expandRequestDuoMfaUpdateMFaByConnectionName(ctx, "parameters.0", d)

		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

		response1, err := client.DuoMfa.UpdateMFaByConnectionName(vvConnectionName, request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(

				"Failure when executing UpdateMFaByConnectionName", err,

				"Failure at UpdateMFaByConnectionName, unexpected response", ""))

			return diags

		}

	}

	return resourceDuoMfaRead(ctx, d, m)
}

func resourceDuoMfaDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vConnectionName, okConnectionName := resourceMap["connection_name"]
	vvConnectionName := vConnectionName

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okConnectionName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.DuoMfa.GetMfa()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsDuoMfaGetMfa(m, getResp1)
		item1, err := searchDuoMfaGetMfa(m, items1, vvConnectionName)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vConnectionName != item1.Mfa.ConnectionName {
			vvConnectionName = item1.Mfa.ConnectionName
		} else {
			vvConnectionName = vConnectionName
		}
	}
	if selectedMethod == 2 {
		getResp, _, err := client.DuoMfa.GetMfaByconnectionName(vvConnectionName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, err := client.DuoMfa.DeleteMfaByConnectionName(vvConnectionName)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteMfaByConnectionName", err,
			"Failure at DeleteMfaByConnectionName, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDuoMfaCreateMfa(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaCreateMfa {
	request := isegosdk.RequestDuoMfaCreateMfa{}
	request.Mfa = expandRequestDuoMfaCreateMfaMfa(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoMfaCreateMfaMfa(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaCreateMfaMfa {
	request := isegosdk.RequestDuoMfaCreateMfaMfa{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".account_configurations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".account_configurations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".account_configurations")))) {
		request.AccountConfigurations = expandRequestDuoMfaCreateMfaMfaAccountConfigurations(ctx, key+".account_configurations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_name")))) {
		request.ConnectionName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_sync")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_sync")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_sync")))) {
		request.IDentitySync = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoMfaCreateMfaMfaAccountConfigurations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaCreateMfaMfaAccountConfigurations {
	request := isegosdk.RequestDuoMfaCreateMfaMfaAccountConfigurations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_api")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_api")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_api")))) {
		request.AdminAPI = expandRequestDuoMfaCreateMfaMfaAccountConfigurationsAdminAPI(ctx, key+".admin_api.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".api_host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".api_host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".api_host_name")))) {
		request.APIHostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_api")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_api")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_api")))) {
		request.AuthenticationAPI = expandRequestDuoMfaCreateMfaMfaAccountConfigurationsAuthenticationAPI(ctx, key+".authentication_api.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoMfaCreateMfaMfaAccountConfigurationsAdminAPI(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaCreateMfaMfaAccountConfigurationsAdminAPI {
	request := isegosdk.RequestDuoMfaCreateMfaMfaAccountConfigurationsAdminAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ikey")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ikey")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ikey")))) {
		request.Ikey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".s_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".s_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".s_key")))) {
		request.SKey = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoMfaCreateMfaMfaAccountConfigurationsAuthenticationAPI(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaCreateMfaMfaAccountConfigurationsAuthenticationAPI {
	request := isegosdk.RequestDuoMfaCreateMfaMfaAccountConfigurationsAuthenticationAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ikey")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ikey")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ikey")))) {
		request.Ikey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".s_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".s_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".s_key")))) {
		request.SKey = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoMfaUpdateMFaByConnectionName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaUpdateMFaByConnectionName {
	request := isegosdk.RequestDuoMfaUpdateMFaByConnectionName{}
	request.Mfa = expandRequestDuoMfaUpdateMFaByConnectionNameMfa(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoMfaUpdateMFaByConnectionNameMfa(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaUpdateMFaByConnectionNameMfa {
	request := isegosdk.RequestDuoMfaUpdateMFaByConnectionNameMfa{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".account_configurations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".account_configurations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".account_configurations")))) {
		request.AccountConfigurations = expandRequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurations(ctx, key+".account_configurations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_name")))) {
		request.ConnectionName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_sync")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_sync")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_sync")))) {
		request.IDentitySync = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurations {
	request := isegosdk.RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_api")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_api")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_api")))) {
		request.AdminAPI = expandRequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAdminAPI(ctx, key+".admin_api.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".api_host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".api_host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".api_host_name")))) {
		request.APIHostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_api")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_api")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_api")))) {
		request.AuthenticationAPI = expandRequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAuthenticationAPI(ctx, key+".authentication_api.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAdminAPI(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAdminAPI {
	request := isegosdk.RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAdminAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ikey")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ikey")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ikey")))) {
		request.Ikey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".s_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".s_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".s_key")))) {
		request.SKey = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAuthenticationAPI(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAuthenticationAPI {
	request := isegosdk.RequestDuoMfaUpdateMFaByConnectionNameMfaAccountConfigurationsAuthenticationAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ikey")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ikey")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ikey")))) {
		request.Ikey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".s_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".s_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".s_key")))) {
		request.SKey = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsDuoMfaGetMfa(m interface{}, response *isegosdk.ResponseDuoMfaGetMfa) []isegosdk.ResponseDuoMfaGetMfaResponse {
	var respItems []isegosdk.ResponseDuoMfaGetMfaResponse
	for response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchDuoMfaGetMfa(m interface{}, items []isegosdk.ResponseDuoMfaGetMfaResponse, connectionName string) (*isegosdk.ResponseDuoMfaGetMfaByconnectionNameResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseDuoMfaGetMfaByconnectionNameResponse
	for _, item := range items {
		if connectionName != "" && item.Name == connectionName {
			var getItem *isegosdk.ResponseDuoMfaGetMfaByconnectionName
			getItem, _, err = client.DuoMfa.GetMfaByconnectionName(connectionName)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetMfaByconnectionName")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
