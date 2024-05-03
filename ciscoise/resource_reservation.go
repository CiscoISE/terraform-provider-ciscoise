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

func resourceReservation() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SgtRangeReservation.

- Reserve given number of SGTs in a continuous range for the given Client.

- Update the reserved ranges of a specific Client by giving the custom start and end index

- Delete the reserved range of SGT for the given Client
`,

		CreateContext: resourceReservationCreate,
		ReadContext:   resourceReservationRead,
		UpdateContext: resourceReservationUpdate,
		DeleteContext: resourceReservationDelete,
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

						"client_id": &schema.Schema{
							Description: `Unique ID of the given client`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"client_name": &schema.Schema{
							Description: `Name of the given client`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"end_index": &schema.Schema{
							Description: `End index of the reserved range`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"start_index": &schema.Schema{
							Description: `Start index of the reserved range`,
							Type:        schema.TypeInt,
							Computed:    true,
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

						"client_id": &schema.Schema{
							Description:      `Unique ID of the given client`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"client_name": &schema.Schema{
							Description:      `Name of the given client`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"end_index": &schema.Schema{
							Description:      `End index of the reserved range`,
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"number_of_tags": &schema.Schema{
							Description:      `Nummber of tags required to be reserved in ISE.`,
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"response": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"start_index": &schema.Schema{
							Description:      `Start index of the reserved range`,
							Type:             schema.TypeInt,
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

func resourceReservationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := m.(ClientConfig).EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestReservationReserveSgtRange(ctx, "parameters.0", d)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vClientID, okClientID := resourceItem["client_id"]
	vvClientID := interfaceToString(vClientID)
	if isEnableAutoImport {
		if okClientID && vvClientID != "" {
			getResponse2, _, err := client.SgtRangeReservation.GetSgtReservedRange(vvClientID)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["client_id"] = vvClientID
				d.SetId(joinResourceID(resourceMap))
				return resourceReservationRead(ctx, d, m)
			}
		} else {
			queryParams2 := isegosdk.GetSgtReservedRangesQueryParams{}

			response2, _, err := client.SgtRangeReservation.GetSgtReservedRanges(&queryParams2)
			if response2 != nil && err == nil {
				items2 := getAllItemsSgtRangeReservationGetSgtReservedRanges(m, response2, &queryParams2)
				item2, err := searchSgtRangeReservationGetSgtReservedRanges(m, items2, vvClientID)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["client_id"] = vvClientID
					d.SetId(joinResourceID(resourceMap))
					return resourceReservationRead(ctx, d, m)
				}
			}
		}
	}
	resp1, restyResp1, err := client.SgtRangeReservation.ReserveSgtRange(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing ReserveSgtRange", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing ReserveSgtRange", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["client_id"] = vvClientID
	d.SetId(joinResourceID(resourceMap))
	return resourceReservationRead(ctx, d, m)
}

func resourceReservationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vClientID, okClientID := resourceMap["client_id"]
	vvClientID := vClientID

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okClientID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSgtReservedRanges")
		queryParams1 := isegosdk.GetSgtReservedRangesQueryParams{}

		response1, restyResp1, err := client.SgtRangeReservation.GetSgtReservedRanges(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsSgtRangeReservationGetSgtReservedRanges(m, response1, nil)
		item1, err := searchSgtRangeReservationGetSgtReservedRanges(m, items1, vvClientID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenSgtRangeReservationGetSgtReservedRangeItemResponse(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSgtReservedRanges search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSgtReservedRanges search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSgtReservedRange")

		response2, restyResp2, err := client.SgtRangeReservation.GetSgtReservedRange(vvClientID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSgtRangeReservationGetSgtReservedRangeItemResponse(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSgtReservedRange response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSgtReservedRange response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceReservationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vClientID, _ := resourceMap["client_id"]
	vvClientID := vClientID

	if d.HasChange("parameters") {

		log.Printf("[DEBUG] ID used for update operation %s", vvClientID)

		request1 := expandRequestReservationUpdateReservedRange(ctx, "parameters.0", d)

		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

		response1, restyResp1, err := client.SgtRangeReservation.UpdateReservedRange(vvClientID, request1)

		if err != nil || response1 == nil {

			if restyResp1 != nil {

				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())

				diags = append(diags, diagErrorWithAltAndResponse(

					"Failure when executing UpdateReservedRange", err, restyResp1.String(),

					"Failure at UpdateReservedRange, unexpected response", ""))

				return diags

			}

			diags = append(diags, diagErrorWithAlt(

				"Failure when executing UpdateReservedRange", err,

				"Failure at UpdateReservedRange, unexpected response", ""))

			return diags

		}

	}

	return resourceReservationRead(ctx, d, m)
}

func resourceReservationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vClientID, okClientID := resourceMap["client_id"]
	vvClientID := vClientID

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okClientID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {
		queryParams1 := isegosdk.GetSgtReservedRangesQueryParams{}

		getResp1, _, err := client.SgtRangeReservation.GetSgtReservedRanges(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSgtRangeReservationGetSgtReservedRanges(m, getResp1, &queryParams1)
		item1, err := searchSgtRangeReservationGetSgtReservedRanges(m, items1, vvClientID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vClientID != item1.ClientID {
			vvClientID = item1.ClientID
		} else {
			vvClientID = vClientID
		}
	}
	if selectedMethod == 2 {
		getResp, _, err := client.SgtRangeReservation.GetSgtReservedRange(vvClientID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.SgtRangeReservation.DeleteSgtReserveRange(vvClientID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSgtReserveRange", err, restyResp1.String(),
				"Failure at DeleteSgtReserveRange, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSgtReserveRange", err,
			"Failure at DeleteSgtReserveRange, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestReservationReserveSgtRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSgtRangeReservationReserveSgtRange {
	request := isegosdk.RequestSgtRangeReservationReserveSgtRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_name")))) {
		request.ClientName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".number_of_tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".number_of_tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".number_of_tags")))) {
		request.NumberOfTags = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestReservationUpdateReservedRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSgtRangeReservationUpdateReservedRange {
	request := isegosdk.RequestSgtRangeReservationUpdateReservedRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_id")))) {
		request.ClientID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_index")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_index")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_index")))) {
		request.EndIndex = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_index")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_index")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_index")))) {
		request.StartIndex = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsSgtRangeReservationGetSgtReservedRanges(m interface{}, response *isegosdk.ResponseSgtRangeReservationGetSgtReservedRanges, queryParams *isegosdk.GetSgtReservedRangesQueryParams) []isegosdk.ResponseSgtRangeReservationGetSgtReservedRangesResponse {
	var respItems []isegosdk.ResponseSgtRangeReservationGetSgtReservedRangesResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchSgtRangeReservationGetSgtReservedRanges(m interface{}, items []isegosdk.ResponseSgtRangeReservationGetSgtReservedRangesResponse, clientID string) (*isegosdk.ResponseSgtRangeReservationGetSgtReservedRangeResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSgtRangeReservationGetSgtReservedRangeResponse
	for _, item := range items {
		if clientID != "" && item.ClientID == clientID {
			var getItem *isegosdk.ResponseSgtRangeReservationGetSgtReservedRange
			getItem, _, err = client.SgtRangeReservation.GetSgtReservedRange(clientID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSgtReservedRange")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
