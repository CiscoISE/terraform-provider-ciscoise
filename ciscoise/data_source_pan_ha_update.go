package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourcePanHaUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on PAN HA.

- To deploy the auto-failover feature, you must have at least three nodes, where two of the nodes assume the
Administration persona, and one node acts as the health check node. A health check node is a non-administration node and
can be a Policy Service, Monitoring, or pxGrid node, or any combination of these. If the PANs are in different data
centers, you must have a health check node for each PAN.
 All the fields are mandatory to enable PanHA.
 Values of failedAttempts, pollingInterval, primaryHealthCheckNode, and secondaryHealthCheckNode are not considered when
the isEnable value is "false" in the request body.
`,

		ReadContext: dataSourcePanHaUpdateRead,
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"success": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"message": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_enabled": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Required:     true,
			},
			"failed_attempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"polling_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"primary_health_check_node": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"secondary_health_check_node": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePanHaUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: UpdatePanHa")
		request1 := expandRequestPanHaUpdateUpdatePanHa(ctx, "", d)

		response1, restyResp1, err := client.PanHa.UpdatePanHa(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatePanHa", err,
				"Failure at UpdatePanHa, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPanHaUpdatePanHaItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdatePanHa response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPanHaUpdateUpdatePanHa(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPanHaUpdatePanHa {
	request := isegosdk.RequestPanHaUpdatePanHa{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failed_attempts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failed_attempts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failed_attempts")))) {
		request.FailedAttempts = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_enabled")))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".polling_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".polling_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".polling_interval")))) {
		request.PollingInterval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_health_check_node")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_health_check_node")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_health_check_node")))) {
		request.PrimaryHealthCheckNode = expandRequestPanHaUpdateUpdatePanHaPrimaryHealthCheckNode(ctx, key+".primary_health_check_node.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_health_check_node")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_health_check_node")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_health_check_node")))) {
		request.SecondaryHealthCheckNode = expandRequestPanHaUpdateUpdatePanHaSecondaryHealthCheckNode(ctx, key+".secondary_health_check_node.0", d)
	}
	return &request
}

func expandRequestPanHaUpdateUpdatePanHaPrimaryHealthCheckNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPanHaUpdatePanHaPrimaryHealthCheckNode {
	request := isegosdk.RequestPanHaUpdatePanHaPrimaryHealthCheckNode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	return &request
}

func expandRequestPanHaUpdateUpdatePanHaSecondaryHealthCheckNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPanHaUpdatePanHaSecondaryHealthCheckNode {
	request := isegosdk.RequestPanHaUpdatePanHaSecondaryHealthCheckNode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	return &request
}

func flattenPanHaUpdatePanHaItem(item *isegosdk.ResponsePanHaUpdatePanHa) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = flattenPanHaUpdatePanHaItemSuccess(item.Success)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenPanHaUpdatePanHaItemSuccess(item *isegosdk.ResponsePanHaUpdatePanHaSuccess) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message

	return []map[string]interface{}{
		respItem,
	}

}
