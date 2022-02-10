package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePanHa() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on PanHa.

- This resource allows the client to update ACI settings.
`,

		CreateContext: resourcePanHaCreate,
		ReadContext:   resourcePanHaRead,
		UpdateContext: resourcePanHaUpdate,
		DeleteContext: resourcePanHaDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"failed_attempts": &schema.Schema{
							Description: `Failover occurs if the primary PAN is down for the specified number of failure polls. Count (2 - 60).<br> The default value is 5. `,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"is_enabled": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"polling_interval": &schema.Schema{
							Description: `Administration nodes are checked after each interval. Seconds (30 - 300) <br> The default value is 120. `,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"primary_health_check_node": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"secondary_health_check_node": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							MaxItems: 1,
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
							MaxItems: 1,
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
				},
			},
		},
	}
}

/*
primary_health_check_node_hostname
secondary_health_check_node_hostname
*/

func resourcePanHaCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PanHa create")
	log.Printf("[DEBUG] Missing PanHa create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceMap := make(map[string]string)

	var primary_health_check_node_hostname string
	var secondary_health_check_node_hostname string

	if _, ok := d.GetOk("parameters.0"); ok {
		if _, ok := d.GetOk("parameters.0.primary_health_check_node"); ok {
			if _, ok := d.GetOk("parameters.0.primary_health_check_node.0"); ok {
				if v, ok := d.GetOk("parameters.0.primary_health_check_node.0.hostname"); ok {
					primary_health_check_node_hostname = v.(string)
				}
			}
		}
		if _, ok := d.GetOk("parameters.0.secondary_health_check_node"); ok {
			if _, ok := d.GetOk("parameters.0.secondary_health_check_node.0"); ok {
				if v, ok := d.GetOk("parameters.0.secondary_health_check_node.0.hostname"); ok {
					secondary_health_check_node_hostname = v.(string)
				}
			}
		}
	}

	log.Printf("[DEBUG] Selected method: UpdatePanHa")
	request1 := expandRequestPanHaUpdateUpdatePanHa(ctx, "parameters.0", d)

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

	resourceMap["id"] = getUnixTimeString()
	resourceMap["primary_health_check_node_hostname"] = primary_health_check_node_hostname
	resourceMap["secondary_health_check_node_hostname"] = secondary_health_check_node_hostname
	d.SetId(joinResourceID(resourceMap))
	return resourcePanHaRead(ctx, d, m)
}

func resourcePanHaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PanHa read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPanHaStatus")

		response1, restyResp1, err := client.PanHa.GetPanHaStatus()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPanHaGetPanHaStatusItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPanHaStatus response",
				err))
			return diags
		}
		if err := d.Set("parameters", remove_parameters(vItem1)); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPanHaStatus response to parameters",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourcePanHaUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PanHa update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Selected method: UpdatePanHa")
		request1 := expandRequestPanHaUpdateUpdatePanHa(ctx, "parameters.0", d)

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
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourcePanHaRead(ctx, d, m)
}

func resourcePanHaDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PanHa delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing PanHa delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
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
