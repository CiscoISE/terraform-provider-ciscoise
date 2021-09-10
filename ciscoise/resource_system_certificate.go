package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificate() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSystemCertificateCreate,
		ReadContext:   resourceSystemCertificateRead,
		UpdateContext: resourceSystemCertificateUpdate,
		DeleteContext: resourceSystemCertificateDelete,
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
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"admin": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"allow_replacement_of_portal_group_tag": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"eap": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"expiration_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"expiration_ttl_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"expiration_ttl_units": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ims": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"issued_by": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"issued_to": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_size": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:             schema.TypeList,
							DiffSuppressFunc: diffSuppressAlways(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
									"rel": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
									"type": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"portal": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"portal_group_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"portals_using_the_tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pxgrid": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"radius": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"renew_self_signed_certificate": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"saml": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"self_signed": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"serial_number_decimal_format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sha256_fingerprint": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"signature_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"used_by": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid_from": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemCertificateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("item"))
	resourceMap := make(map[string]string)
	// NOTE: Function does not perform create on ISE
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	resourceMap["host_name"] = interfaceToString(resourceItem["host_name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSystemCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostName, okHostName := resourceMap["host_name"]
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okHostName, okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostName, okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	vvHostName := vHostName
	vvName := vName
	vvID := vID
	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSystemCertificates")
		queryParams1 := isegosdk.GetSystemCertificatesQueryParams{}

		response1, _, err := client.Certificates.GetSystemCertificates(vvHostName, &queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSystemCertificates", err,
				"Failure at GetSystemCertificates, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsCertificatesGetSystemCertificates(m, response1, vvHostName, &queryParams1)
		item1, err := searchCertificatesGetSystemCertificates(m, items1, vvName, vvID, vvHostName)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetSystemCertificates response", err,
				"Failure when searching item from GetSystemCertificates, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSystemCertificates search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSystemCertificateByID")

		response2, _, err := client.Certificates.GetSystemCertificateByID(vvHostName, vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSystemCertificateByID", err,
				"Failure at GetSystemCertificateByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenCertificatesGetSystemCertificateByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSystemCertificateByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSystemCertificateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostName, okHostName := resourceMap["host_name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okHostName, okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvHostName string
	if selectedMethod == 1 {
		vvID = vID
		vvHostName = vHostName
	}
	if selectedMethod == 2 {
		vvHostName = vHostName
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestSystemCertificateUpdateSystemCertificate(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Certificates.UpdateSystemCertificate(vvID, vvHostName, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSystemCertificate", err, restyResp1.String(),
					"Failure at UpdateSystemCertificate, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSystemCertificate", err,
				"Failure at UpdateSystemCertificate, unexpected response", ""))
			return diags
		}
	}

	return resourceSystemCertificateRead(ctx, d, m)
}

func resourceSystemCertificateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostName, okHostName := resourceMap["host_name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okHostName, okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	vvHostName := vHostName
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.Certificates.GetSystemCertificateByID(vvHostName, vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.Certificates.DeleteSystemCertificateByID(vvHostName, vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSystemCertificateByID", err, restyResp1.String(),
				"Failure at DeleteSystemCertificateByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSystemCertificateByID", err,
			"Failure at DeleteSystemCertificateByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSystemCertificateUpdateSystemCertificate(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesUpdateSystemCertificate {
	request := isegosdk.RequestCertificatesUpdateSystemCertificate{}
	if v, ok := d.GetOkExists(key + ".admin"); !isEmptyValue(reflect.ValueOf(d.Get(key+".admin"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".admin"))) {
		request.Admin = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_replacement_of_portal_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_replacement_of_portal_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_replacement_of_portal_group_tag"))) {
		request.AllowReplacementOfPortalGroupTag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".eap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".eap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".eap"))) {
		request.Eap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".expiration_ttl_period"); !isEmptyValue(reflect.ValueOf(d.Get(key+".expiration_ttl_period"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".expiration_ttl_period"))) {
		request.ExpirationTTLPeriod = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".expiration_ttl_units"); !isEmptyValue(reflect.ValueOf(d.Get(key+".expiration_ttl_units"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".expiration_ttl_units"))) {
		request.ExpirationTTLUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ims"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ims"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ims"))) {
		request.Ims = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal"))) {
		request.Portal = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_group_tag"))) {
		request.PortalGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".pxgrid"); !isEmptyValue(reflect.ValueOf(d.Get(key+".pxgrid"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".pxgrid"))) {
		request.Pxgrid = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".radius"); !isEmptyValue(reflect.ValueOf(d.Get(key+".radius"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".radius"))) {
		request.Radius = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".renew_self_signed_certificate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".renew_self_signed_certificate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".renew_self_signed_certificate"))) {
		request.RenewSelfSignedCertificate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".saml"); !isEmptyValue(reflect.ValueOf(d.Get(key+".saml"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".saml"))) {
		request.Saml = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsCertificatesGetSystemCertificates(m interface{}, response *isegosdk.ResponseCertificatesGetSystemCertificates, hostname string, queryParams *isegosdk.GetSystemCertificatesQueryParams) []isegosdk.ResponseCertificatesGetSystemCertificatesResponse {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseCertificatesGetSystemCertificatesResponse
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
			response, _, err = client.Certificates.GetSystemCertificates(hostname, queryParams)
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

func searchCertificatesGetSystemCertificates(m interface{}, items []isegosdk.ResponseCertificatesGetSystemCertificatesResponse, name string, id string, hostname string) (*isegosdk.ResponseCertificatesGetSystemCertificateByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseCertificatesGetSystemCertificateByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseCertificatesGetSystemCertificateByID
			getItem, _, err = client.Certificates.GetSystemCertificateByID(hostname, id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSystemCertificateByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.FriendlyName == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseCertificatesGetSystemCertificateByID
			getItem, _, err = client.Certificates.GetSystemCertificateByID(hostname, item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSystemCertificateByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}