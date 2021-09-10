package ciscoise

import (
	"context"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePanHa() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourcePanHaCreate,
		ReadContext:   resourcePanHaRead,
		UpdateContext: resourcePanHaUpdate,
		DeleteContext: resourcePanHaDelete,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"failed_attempts": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"is_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"polling_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"primary_health_check_node": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"secondary_health_check_node": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourcePanHaCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	request1 := expandRequestPanHaEnablePanHa(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.PanHa.EnablePanHa(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing EnablePanHa", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing EnablePanHa", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourcePanHaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPanHaStatus")

		response1, _, err := client.PanHa.GetPanHaStatus()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPanHaStatus", err,
				"Failure at GetPanHaStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)
		items1 := getAllItemsPanHaGetPanHaStatus(m, response1)
		if err := d.Set("item", items1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPanHaStatus search response",
				err))
			return diags
		}

	}
	return diags
}

func resourcePanHaUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePanHaRead(ctx, d, m)
}

func resourcePanHaDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.PanHa.GetPanHaStatus()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsPanHaGetPanHaStatus(m, getResp1)
		if err != nil || items1 == nil || len(items1) == 0 {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.PanHa.DisablePanHa()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DisablePanHa", err, restyResp1.String(),
				"Failure at DisablePanHa, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DisablePanHa", err,
			"Failure at DisablePanHa, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestPanHaEnablePanHa(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPanHaEnablePanHa {
	request := isegosdk.RequestPanHaEnablePanHa{}
	request.Request = expandRequestPanHaEnablePanHaRequest(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPanHaEnablePanHaRequest(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPanHaEnablePanHaRequest {
	request := isegosdk.RequestPanHaEnablePanHaRequest{}
	if v, ok := d.GetOkExists(key + ".is_enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_enabled"))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".primary_health_check_node"); !isEmptyValue(reflect.ValueOf(d.Get(key+".primary_health_check_node"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".primary_health_check_node"))) {
		request.PrimaryHealthCheckNode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".secondary_health_check_node"); !isEmptyValue(reflect.ValueOf(d.Get(key+".secondary_health_check_node"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".secondary_health_check_node"))) {
		request.SecondaryHealthCheckNode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".polling_interval"); !isEmptyValue(reflect.ValueOf(d.Get(key+".polling_interval"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".polling_interval"))) {
		request.PollingInterval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".failed_attempts"); !isEmptyValue(reflect.ValueOf(d.Get(key+".failed_attempts"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".failed_attempts"))) {
		request.FailedAttempts = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsPanHaGetPanHaStatus(m interface{}, response *isegosdk.ResponsePanHaGetPanHaStatus) []isegosdk.ResponsePanHaGetPanHaStatusResponse {
	var respItems []isegosdk.ResponsePanHaGetPanHaStatusResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}
