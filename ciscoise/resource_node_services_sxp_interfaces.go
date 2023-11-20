package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNodeServicesSxpInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Node Services.

- This resource configures the SXP interface.
`,

		CreateContext: resourceNodeServicesSxpInterfacesCreate,
		ReadContext:   resourceNodeServicesSxpInterfacesRead,
		UpdateContext: resourceNodeServicesSxpInterfacesUpdate,
		DeleteContext: resourceNodeServicesSxpInterfacesDelete,
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

						"interface": &schema.Schema{
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

						"hostname": &schema.Schema{
							Description:      `hostname path parameter. Hostname of the node.`,
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: diffSupressOptional(),
						},
						"interface": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func resourceNodeServicesSxpInterfacesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeServicesSxpInterfaces create")
	log.Printf("[DEBUG] Missing NodeServicesSxpInterfaces create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)

	vHostname := interfaceToString(resourceItem["hostname"])
	log.Printf("[DEBUG] Name used for update operation %s", vHostname)
	request1 := expandRequestNodeServicesSxpInterfacesSetSxpInterface(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, restyResp1, err := client.NodeServices.SetSxpInterface(vHostname, request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing SetSxpInterface", err, restyResp1.String(),
				"Failure at SetSxpInterface, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing SetSxpInterface", err,
			"Failure at SetSxpInterface, unexpected response", ""))
		return diags
	}

	resourceMap["hostname"] = interfaceToString(resourceItem["hostname"])
	d.SetId(joinResourceID(resourceMap))
	return resourceNodeServicesSxpInterfacesRead(ctx, d, m)
}

func resourceNodeServicesSxpInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeServicesSxpInterfaces read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostname, _ := resourceMap["hostname"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSxpInterface")
		vvHostname := vHostname

		response1, restyResp1, err := client.NodeServices.GetSxpInterface(vvHostname)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNodeServicesGetSxpInterfaceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSxpInterface response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSxpInterface response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNodeServicesSxpInterfacesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeServicesSxpInterfaces update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostname, _ := resourceMap["hostname"]

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vHostname)
		request1 := expandRequestNodeServicesSxpInterfacesSetSxpInterface(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.NodeServices.SetSxpInterface(vHostname, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SetSxpInterface", err, restyResp1.String(),
					"Failure at SetSxpInterface, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SetSxpInterface", err,
				"Failure at SetSxpInterface, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceNodeServicesSxpInterfacesRead(ctx, d, m)
}

func resourceNodeServicesSxpInterfacesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NodeServicesSxpInterfaces delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing NodeServicesSxpInterfaces delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
func expandRequestNodeServicesSxpInterfacesSetSxpInterface(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeServicesSetSxpInterface {
	request := isegosdk.RequestNodeServicesSetSxpInterface{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface")))) {
		request.Interface = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
