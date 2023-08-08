package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceThreatVulnerabilitiesClear() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ClearThreatsAndVulnerabilities.
- This resource allows the client to delete the ThreatContext and Threat events that are associated with the
given MAC Address.
`,

		CreateContext: resourceThreatVulnerabilitiesClearCreate,
		ReadContext:   resourceThreatVulnerabilitiesClearRead,
		DeleteContext: resourceThreatVulnerabilitiesClearDelete,

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
						"mac_addresses": &schema.Schema{
							Type:     schema.TypeString,
							ForceNew: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceThreatVulnerabilitiesClearCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ClearThreatsAndVulnerabilities create")
	log.Printf("[DEBUG] Missing ClearThreatsAndVulnerabilities create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	request1 := expandRequestThreatVulnerabilitiesClearClearThreatsAndVulnerabilities(ctx, "parameters.0", d)
	response1, err := client.ClearThreatsAndVulnerabilities.ClearThreatsAndVulnerabilities(request1)
	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing ClearThreatsAndVulnerabilities", err, response1.String(),
				"Failure at ClearThreatsAndVulnerabilities, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ClearThreatsAndVulnerabilities", err,
			"Failure at ClearThreatsAndVulnerabilities, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %s", response1.String())
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ClearThreatsAndVulnerabilities response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourceThreatVulnerabilitiesClearRead(ctx, d, m)
}

func resourceThreatVulnerabilitiesClearRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceThreatVulnerabilitiesClearDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ThreatVulnerabilitiesClear delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing ThreatVulnerabilitiesClear delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestThreatVulnerabilitiesClearClearThreatsAndVulnerabilities(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilities {
	request := isegosdk.RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilities{}
	request.ERSIrfThreatContext = expandRequestThreatVulnerabilitiesClearClearThreatsAndVulnerabilitiesERSIrfThreatContext(ctx, key, d)
	return &request
}

func expandRequestThreatVulnerabilitiesClearClearThreatsAndVulnerabilitiesERSIrfThreatContext(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilitiesERSIrfThreatContext {
	request := isegosdk.RequestClearThreatsAndVulnerabilitiesClearThreatsAndVulnerabilitiesERSIrfThreatContext{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_addresses")))) {
		request.MacAddresses = interfaceToString(v)
	}
	return &request
}
