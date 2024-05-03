package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourcePxGridDirectTestConnector() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on pxGrid Direct.

- pxGrid Direct test the Connector.
`,

		CreateContext: resourcePxGridDirectTestConnectorCreate,
		ReadContext:   resourcePxGridDirectTestConnectorRead,
		DeleteContext: resourcePxGridDirectTestConnectorDelete,
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

						"connector_name": &schema.Schema{
							Description: `connectorName`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"data": &schema.Schema{
							Description: `Response data`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"error": &schema.Schema{
							Description: `error`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"skip_certificate_validations": &schema.Schema{
							Description: `skipCertificateValidations`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Description: `status of the request`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"unique_id": &schema.Schema{
							Description: `uniqueness to identify`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"x_request_id": &schema.Schema{
							Description: `X-Request-ID header parameter. request Id, will return in the response headers, and appear in logs`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"password": &schema.Schema{
							Description: `password`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Sensitive:   true,
							Computed:    true,
						},
						"user_name": &schema.Schema{
							Description: `userName`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourcePxGridDirectTestConnectorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestPxGridDirectTestConnectorTestConnector(ctx, "parameters.0", d)

	response1, restyResp1, err := client.PxGridDirect.TestConnector(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenPxGridDirectTestConnectorItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting TestConnector response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}

func expandRequestPxGridDirectTestConnectorTestConnector(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectTestConnector {
	request := isegosdk.RequestPxGridDirectTestConnector{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_values")))) {
		request.AuthValues = expandRequestPxGridDirectTestConnectorTestConnectorAuthValues(ctx, key+".auth_values.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_name")))) {
		request.ConnectorName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".response_parsing")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".response_parsing")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".response_parsing")))) {
		request.ResponseParsing = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".skip_certificate_validations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".skip_certificate_validations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".skip_certificate_validations")))) {
		request.SkipCertificateValidations = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".unique_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".unique_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".unique_id")))) {
		request.UniqueID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	return &request
}

func expandRequestPxGridDirectTestConnectorTestConnectorAuthValues(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectTestConnectorAuthValues {
	request := isegosdk.RequestPxGridDirectTestConnectorAuthValues{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	return &request
}

func flattenPxGridDirectTestConnectorItem(item *isegosdk.ResponsePxGridDirectTestConnector) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connector_name"] = item.ConnectorName
	respItem["data"] = item.Data
	respItem["error"] = item.Error
	respItem["skip_certificate_validations"] = boolPtrToString(item.SkipCertificateValidations)
	respItem["status"] = item.Status
	respItem["unique_id"] = item.UniqueID
	return []map[string]interface{}{
		respItem,
	}
}

func resourcePxGridDirectTestConnectorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePxGridDirectTestConnectorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePxGridDirectTestConnectorRead(ctx, d, m)
}

func resourcePxGridDirectTestConnectorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
