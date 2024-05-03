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

func resourceUserEquipment() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on User Equipment.

- Create a user equipment

- Update the user equipment for a given ID and request payload

- Delete the user equipment for a given ID
`,

		CreateContext: resourceUserEquipmentCreate,
		ReadContext:   resourceUserEquipmentRead,
		UpdateContext: resourceUserEquipmentUpdate,
		DeleteContext: resourceUserEquipmentDelete,
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

						"create_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `Description for User Equipment`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"device_group": &schema.Schema{
							Description: `Device or Endpoint Group`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"imei": &schema.Schema{
							Description: `IMEI for User Equipment`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"link": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rel": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"href": {
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
						"update_time": &schema.Schema{
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

						"description": &schema.Schema{
							Description:      `Description for User Equipment`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"device_group": &schema.Schema{
							Description:      `Device or Endpoint Group`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"imei": &schema.Schema{
							Description:      `IMEI for User Equipment`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"response": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"user_equipment_id": &schema.Schema{
							Description:      `userEquipmentId path parameter. Unique ID for a user equipment object`,
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: diffSupressOptional(),
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

func resourceUserEquipmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := m.(ClientConfig).EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestUserEquipmentCreateUserEquipment(ctx, "parameters.0", d)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vUserEquipmentID, okUserEquipmentID := resourceItem["user_equipment_id"]
	vvUserEquipmentID := interfaceToString(vUserEquipmentID)
	if isEnableAutoImport {
		if okUserEquipmentID && vvUserEquipmentID != "" {
			getResponse2, _, err := client.UserEquipment.GetUserEquipmentByID(vvUserEquipmentID)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["user_equipment_id"] = vvUserEquipmentID
				d.SetId(joinResourceID(resourceMap))
				return resourceUserEquipmentRead(ctx, d, m)
			}
		} else {
			queryParams2 := isegosdk.GetUserEquipmentsQueryParams{}

			response2, _, err := client.UserEquipment.GetUserEquipments(&queryParams2)
			if response2 != nil && err == nil {
				items2 := getAllItemsUserEquipmentGetUserEquipments(m, response2, &queryParams2)
				item2, err := searchUserEquipmentGetUserEquipments(m, items2, vvUserEquipmentID)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["user_equipment_id"] = vvUserEquipmentID
					d.SetId(joinResourceID(resourceMap))
					return resourceUserEquipmentRead(ctx, d, m)
				}
			}
		}
	}
	resp1, restyResp1, err := client.UserEquipment.CreateUserEquipment(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateUserEquipment", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateUserEquipment", err))
		return diags
	}
	if vvUserEquipmentID != resp1.Response.ID {
		vvUserEquipmentID = resp1.Response.ID
	}
	resourceMap := make(map[string]string)
	resourceMap["user_equipment_id"] = vvUserEquipmentID
	d.SetId(joinResourceID(resourceMap))
	return resourceUserEquipmentRead(ctx, d, m)
}

func resourceUserEquipmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vUserEquipmentID, okUserEquipmentID := resourceMap["user_equipment_id"]
	vvUserEquipmentID := vUserEquipmentID

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okUserEquipmentID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetUserEquipments")
		queryParams1 := isegosdk.GetUserEquipmentsQueryParams{}

		response1, restyResp1, err := client.UserEquipment.GetUserEquipments(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsUserEquipmentGetUserEquipments(m, response1, &queryParams1)
		item1, err := searchUserEquipmentGetUserEquipments(m, items1, vvUserEquipmentID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenUserEquipmentGetUserEquipmentByIDItemResponse(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetUserEquipments search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetUserEquipments search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetUserEquipmentByID")

		response2, restyResp2, err := client.UserEquipment.GetUserEquipmentByID(vvUserEquipmentID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenUserEquipmentGetUserEquipmentByIDItemResponse(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetUserEquipmentByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetUserEquipmentByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceUserEquipmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vUserEquipmentID, _ := resourceMap["user_equipment_id"]
	vvUserEquipmentID := vUserEquipmentID

	if d.HasChange("parameters") {

		log.Printf("[DEBUG] ID used for update operation %s", vvUserEquipmentID)

		request1 := expandRequestUserEquipmentUpdateUserEquipment(ctx, "parameters.0", d)

		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

		response1, restyResp1, err := client.UserEquipment.UpdateUserEquipment(vvUserEquipmentID, request1)

		if err != nil || response1 == nil {

			if restyResp1 != nil {

				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())

				diags = append(diags, diagErrorWithAltAndResponse(

					"Failure when executing UpdateUserEquipment", err, restyResp1.String(),

					"Failure at UpdateUserEquipment, unexpected response", ""))

				return diags

			}

			diags = append(diags, diagErrorWithAlt(

				"Failure when executing UpdateUserEquipment", err,

				"Failure at UpdateUserEquipment, unexpected response", ""))

			return diags

		}

	}

	return resourceUserEquipmentRead(ctx, d, m)
}

func resourceUserEquipmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vUserEquipmentID, okUserEquipmentID := resourceMap["user_equipment_id"]
	vvUserEquipmentID := vUserEquipmentID

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okUserEquipmentID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {
		queryParams1 := isegosdk.GetUserEquipmentsQueryParams{}

		getResp1, _, err := client.UserEquipment.GetUserEquipments(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsUserEquipmentGetUserEquipments(m, getResp1, &queryParams1)
		item1, err := searchUserEquipmentGetUserEquipments(m, items1, vvUserEquipmentID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vUserEquipmentID != item1.ID {
			vvUserEquipmentID = item1.ID
		} else {
			vvUserEquipmentID = vUserEquipmentID
		}
	}
	if selectedMethod == 2 {
		getResp, _, err := client.UserEquipment.GetUserEquipmentByID(vvUserEquipmentID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.UserEquipment.DeleteUserEquipment(vvUserEquipmentID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteUserEquipment", err, restyResp1.String(),
				"Failure at DeleteUserEquipment, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteUserEquipment", err,
			"Failure at DeleteUserEquipment, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestUserEquipmentCreateUserEquipment(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestUserEquipmentCreateUserEquipment {
	request := isegosdk.RequestUserEquipmentCreateUserEquipment{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_group")))) {
		request.DeviceGroup = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".imei")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".imei")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".imei")))) {
		request.Imei = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestUserEquipmentUpdateUserEquipment(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestUserEquipmentUpdateUserEquipment {
	request := isegosdk.RequestUserEquipmentUpdateUserEquipment{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_group")))) {
		request.DeviceGroup = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsUserEquipmentGetUserEquipments(m interface{}, response *isegosdk.ResponseUserEquipmentGetUserEquipments, queryParams *isegosdk.GetUserEquipmentsQueryParams) []isegosdk.ResponseUserEquipmentGetUserEquipmentsResponse {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseUserEquipmentGetUserEquipmentsResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
		if response.NextPage != nil && response.NextPage.Rel == "next" {
			href := response.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			if queryParams != nil {
				queryParams.Page = page
				queryParams.Size = size
			}
			response, _, err = client.UserEquipment.GetUserEquipments(queryParams)
			if err != nil {
				break
			}
			// All is good, continue to the next page
			continue
		}
		// Does not have next page finish iteration
		break
	}
	return respItems
}

func searchUserEquipmentGetUserEquipments(m interface{}, items []isegosdk.ResponseUserEquipmentGetUserEquipmentsResponse, userEquipmentID string) (*isegosdk.ResponseUserEquipmentGetUserEquipmentByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseUserEquipmentGetUserEquipmentByIDResponse
	for _, item := range items {
		if userEquipmentID != "" && item.ID == userEquipmentID {
			var getItem *isegosdk.ResponseUserEquipmentGetUserEquipmentByID
			getItem, _, err = client.UserEquipment.GetUserEquipmentByID(userEquipmentID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetUserEquipmentByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}

	return foundItem, err
}
