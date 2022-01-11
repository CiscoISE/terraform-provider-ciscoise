package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTransportGatewaySettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on telemetry.

- Transport Gateway acts a proxy for the communication between the ISE servers in your network and the Telemetry servers
in case of air-gapped network.
`,

		CreateContext: resourceTransportGatewaySettingsCreate,
		ReadContext:   resourceTransportGatewaySettingsRead,
		UpdateContext: resourceTransportGatewaySettingsUpdate,
		DeleteContext: resourceTransportGatewaySettingsDelete,
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

						"enable_transport_gateway": &schema.Schema{
							Description: `Indicates whether transport gateway is enabled or not.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"url": &schema.Schema{
							Description: `URL of transport gateway`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enable_transport_gateway": &schema.Schema{
							Description:  `Indicates whether transport gateway is enabled or not.`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"url": &schema.Schema{
							Description: `URL of transport gateway`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceTransportGatewaySettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TransportGatewaySettings Create")
	// var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	resourceMap["url"] = interfaceToString(resourceItem["url"])
	d.SetId(joinResourceID(resourceMap))
	return resourceTransportGatewaySettingsRead(ctx, d, m)
}

func resourceTransportGatewaySettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TransportGatewaySettings Read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTransportGateway")

		response1, restyResp1, err := client.Telemetry.GetTransportGateway()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTelemetryGetTransportGatewayItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTransportGateway response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTransportGatewaySettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TransportGatewaySettings Update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation")
		request1 := expandRequestTransportGatewaySettingsUpdateTransportGateway(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.Telemetry.UpdateTransportGateway(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTransportGateway", err, restyResp1.String(),
					"Failure at UpdateTransportGateway, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTransportGateway", err,
				"Failure at UpdateTransportGateway, unexpected response", ""))
			return diags
		}
	}

	return resourceTransportGatewaySettingsRead(ctx, d, m)
}

func resourceTransportGatewaySettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TransportGatewaySettings Delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	// NOTE: Unable to delete TransportGatewaySettings on Cisco ISE
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestTransportGatewaySettingsUpdateTransportGateway(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTelemetryUpdateTransportGateway {
	request := isegosdk.RequestTelemetryUpdateTransportGateway{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_transport_gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_transport_gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_transport_gateway")))) {
		request.EnableTransportGateway = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
