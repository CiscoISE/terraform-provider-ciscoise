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

func resourceTrustedCertificate() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read, update and delete operations on Certificates.
  
  - Update a trusted certificate present in ISE trust store.
  
  
  - This resource deletes a Trust Certificate from Trusted Certificate Store based on a given ID.
  `,

		CreateContext: resourceTrustedCertificateCreate,
		ReadContext:   resourceTrustedCertificateRead,
		UpdateContext: resourceTrustedCertificateUpdate,
		DeleteContext: resourceTrustedCertificateDelete,
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

						"authenticate_before_crl_received": &schema.Schema{
							Description: `Switch to enable/disable CRL Verification if CRL is not Received
  ERROR: Different types for param authenticateBeforeCRLReceived schema.TypeBool schema.TypeString`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false", "on", "off"}),
							Optional:     true,
							Computed:     true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								log.Printf("[DEBUG] Performing comparison to see if key %s requires diff suppression", k)
								if old == "off" {
									return old == new || "false" == new
								}
								if old == "false" {
									return old == new || "off" == new
								}
								if old == "on" {
									return old == new || "true" == new
								}
								if old == "true" {
									return old == new || "on" == new
								}
								return true
							},
						},
						"automatic_crl_update": &schema.Schema{
							Description: `Switch to enable/disable automatic CRL update
  ERROR: Different types for param automaticCRLUpdate schema.TypeBool schema.TypeString`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false", "on", "off"}),
							Optional:     true,
							Computed:     true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								log.Printf("[DEBUG] Performing comparison to see if key %s requires diff suppression", k)
								if old == "off" {
									return old == new || "false" == new
								}
								if old == "false" {
									return old == new || "off" == new
								}
								if old == "on" {
									return old == new || "true" == new
								}
								if old == "true" {
									return old == new || "on" == new
								}
								return true
							},
						},
						"automatic_crl_update_period": &schema.Schema{
							Description: `Automatic CRL update period
  ERROR: Different types for param automaticCRLUpdatePeriod schema.TypeInt schema.TypeString`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"automatic_crl_update_units": &schema.Schema{
							Description: `Unit of time for automatic CRL update`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"crl_distribution_url": &schema.Schema{
							Description: `CRL Distribution URL`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"crl_download_failure_retries": &schema.Schema{
							Description: `If CRL download fails, wait time before retry
  ERROR: Different types for param crlDownloadFailureRetries schema.TypeInt schema.TypeString`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"crl_download_failure_retries_units": &schema.Schema{
							Description: `Unit of time before retry if CRL download fails`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `Description for trust certificate`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"download_crl": &schema.Schema{
							Description: `Switch to enable/disable download of CRL
  ERROR: Different types for param downloadCRL schema.TypeBool schema.TypeString`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false", "on", "off"}),
							Optional:     true,
							Computed:     true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								log.Printf("[DEBUG] Performing comparison to see if key %s requires diff suppression", k)
								if old == "off" {
									return old == new || "false" == new
								}
								if old == "false" {
									return old == new || "off" == new
								}
								if old == "on" {
									return old == new || "true" == new
								}
								if old == "true" {
									return old == new || "on" == new
								}
								return true
							},
						},
						"enable_ocsp_validation": &schema.Schema{
							Description: `Switch to enable/disable OCSP Validation
  ERROR: Different types for param enableOCSPValidation schema.TypeBool schema.TypeString`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false", "on", "off"}),
							Optional:     true,
							Computed:     true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								log.Printf("[DEBUG] Performing comparison to see if key %s requires diff suppression", k)
								if old == "off" {
									return old == new || "false" == new
								}
								if old == "false" {
									return old == new || "off" == new
								}
								if old == "on" {
									return old == new || "true" == new
								}
								if old == "true" {
									return old == new || "on" == new
								}
								return true
							},
						},
						"enable_server_identity_check": &schema.Schema{
							Description: `Switch to enable/disable verification if HTTPS or LDAP server certificate name fits the configured server URL
  ERROR: Different types for param enableServerIdentityCheck schema.TypeBool schema.TypeString`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false", "on", "off"}),
							Optional:     true,
							Computed:     true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								log.Printf("[DEBUG] Performing comparison to see if key %s requires diff suppression", k)
								if old == "off" {
									return old == new || "false" == new
								}
								if old == "false" {
									return old == new || "off" == new
								}
								if old == "on" {
									return old == new || "true" == new
								}
								if old == "true" {
									return old == new || "on" == new
								}
								return true
							},
						},
						"expiration_date": &schema.Schema{
							Description: `The time and date past which the certificate is no longer valid`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"friendly_name": &schema.Schema{
							Description: `Friendly name of trust certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `ID of trust certificate`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ignore_crl_expiration": &schema.Schema{
							Description: `Switch to enable/disable ignore CRL Expiration
  ERROR: Different types for param ignoreCRLExpiration schema.TypeBool schema.TypeString`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false", "on", "off"}),
							Optional:     true,
							Computed:     true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								log.Printf("[DEBUG] Performing comparison to see if key %s requires diff suppression", k)
								if old == "off" {
									return old == new || "false" == new
								}
								if old == "false" {
									return old == new || "off" == new
								}
								if old == "on" {
									return old == new || "true" == new
								}
								if old == "true" {
									return old == new || "on" == new
								}
								return true
							},
						},
						"internal_ca": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_referred_in_policy": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"issued_by": &schema.Schema{
							Description: `The entity that verified the information and signed the certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"issued_to": &schema.Schema{
							Description: `Entity to which trust certificate is issued`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"key_size": &schema.Schema{
							Description: `The length of key used for encrypting trust certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `Friendly name of the certificate`,
							Type:        schema.TypeString,
							Required:    true,
						},
						"non_automatic_crl_update_period": &schema.Schema{
							Description: `Non automatic CRL update period
  ERROR: Different types for param nonAutomaticCRLUpdatePeriod schema.TypeInt schema.TypeString`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"non_automatic_crl_update_units": &schema.Schema{
							Description: `Unit of time of non automatic CRL update`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"reject_if_no_status_from_ocs_p": &schema.Schema{
							Description: `Switch to reject certificate if there is no status from OCSP
  ERROR: Different types for param rejectIfNoStatusFromOCSP schema.TypeBool schema.TypeString`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false", "on", "off"}),
							Optional:     true,
							Computed:     true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								log.Printf("[DEBUG] Performing comparison to see if key %s requires diff suppression", k)
								if old == "off" {
									return old == new || "false" == new
								}
								if old == "false" {
									return old == new || "off" == new
								}
								if old == "on" {
									return old == new || "true" == new
								}
								if old == "true" {
									return old == new || "on" == new
								}
								return true
							},
						},
						"reject_if_unreachable_from_ocs_p": &schema.Schema{
							Description: `Switch to reject certificate if unreachable from OCSP
  ERROR: Different types for param rejectIfUnreachableFromOCSP schema.TypeBool schema.TypeString`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false", "on", "off"}),
							Optional:     true,
							Computed:     true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								log.Printf("[DEBUG] Performing comparison to see if key %s requires diff suppression", k)
								if old == "off" {
									return old == new || "false" == new
								}
								if old == "false" {
									return old == new || "off" == new
								}
								if old == "on" {
									return old == new || "true" == new
								}
								if old == "true" {
									return old == new || "on" == new
								}
								return true
							},
						},
						"selected_ocsp_service": &schema.Schema{
							Description: `Name of selected OCSP Service`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"serial_number_decimal_format": &schema.Schema{
							Description: `Used to uniquely identify the certificate within a CA's systems`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sha256_fingerprint": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"signature_algorithm": &schema.Schema{
							Description: `Algorithm used for encrypting trust certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subject": &schema.Schema{
							Description: `The Subject or entity with which public key of trust certificate is associated`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"trust_for_certificate_based_admin_auth": &schema.Schema{
							Description:  `Trust for Certificate based Admin authentication`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"trust_for_cisco_services_auth": &schema.Schema{
							Description:  `Trust for authentication of Cisco Services`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"trust_for_client_auth": &schema.Schema{
							Description:  `Trust for client authentication and Syslog`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"trust_for_ise_auth": &schema.Schema{
							Description:  `Trust for authentication within ISE`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"trusted_for": &schema.Schema{
							Description: `Different services for which the certificated is trusted`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"valid_from": &schema.Schema{
							Description: `The earliest time and date on which the certificate is valid`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceTrustedCertificateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("item"))
	resourceMap := make(map[string]string)
	// NOTE: Function does not perform create on ISE
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceTrustedCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	vvName := vName
	vvID := vID

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetTrustedCertificates")
		queryParams1 := isegosdk.GetTrustedCertificatesQueryParams{}
		response1, _, err := client.Certificates.GetTrustedCertificates(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTrustedCertificates", err,
				"Failure at GetTrustedCertificates, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsCertificatesGetTrustedCertificates(m, response1, &queryParams1)
		item1, err := searchCertificatesGetTrustedCertificates(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetTrustedCertificates response", err,
				"Failure when searching item from GetTrustedCertificates, unexpected response", ""))
			return diags
		}
		vItem1 := flattenCertificatesGetTrustedCertificateByIDItem2(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTrustedCertificates search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTrustedCertificateByID")

		response2, _, err := client.Certificates.GetTrustedCertificateByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTrustedCertificateByID", err,
				"Failure at GetTrustedCertificateByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenCertificatesGetTrustedCertificateByIDItem2(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTrustedCertificateByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTrustedCertificateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetTrustedCertificatesQueryParams{}
		getResp1, _, err := client.Certificates.GetTrustedCertificates(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsCertificatesGetTrustedCertificates(m, getResp1, &queryParams1)
			item1, err := searchCertificatesGetTrustedCertificates(m, items1, vName, vID)
			if err == nil && item1 != nil {
				if vID != item1.ID {
					vvID = item1.ID
				} else {
					vvID = vID
				}
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestTrustedCertificateUpdateTrustedCertificate(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Certificates.UpdateTrustedCertificate(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTrustedCertificate", err, restyResp1.String(),
					"Failure at UpdateTrustedCertificate, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTrustedCertificate", err,
				"Failure at UpdateTrustedCertificate, unexpected response", ""))
			return diags
		}
	}

	return resourceTrustedCertificateRead(ctx, d, m)
}

func resourceTrustedCertificateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetTrustedCertificatesQueryParams{}

		getResp1, _, err := client.Certificates.GetTrustedCertificates(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsCertificatesGetTrustedCertificates(m, getResp1, &queryParams1)
		item1, err := searchCertificatesGetTrustedCertificates(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.Certificates.GetTrustedCertificateByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.Certificates.DeleteTrustedCertificateByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteTrustedCertificateByID", err, restyResp1.String(),
				"Failure at DeleteTrustedCertificateByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteTrustedCertificateByID", err,
			"Failure at DeleteTrustedCertificateByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTrustedCertificateUpdateTrustedCertificate(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesUpdateTrustedCertificate {
	request := isegosdk.RequestCertificatesUpdateTrustedCertificate{}
	vDownloadCRL, okDownloadCRL := d.GetOk(key + ".download_crl")
	vvDownloadCRL := interfaceToBoolPtr(vDownloadCRL)
	if okDownloadCRL && vvDownloadCRL != nil && *vvDownloadCRL {
		if v, ok := d.GetOkExists(key + ".automatic_crl_update"); !isEmptyValue(reflect.ValueOf(d.Get(key+".automatic_crl_update"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".automatic_crl_update"))) {
			request.AutomaticCRLUpdate = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(key + ".enable_server_identity_check"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_server_identity_check"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_server_identity_check"))) {
			request.EnableServerIDentityCheck = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(key + ".authenticate_before_crl_received"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authenticate_before_crl_received"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authenticate_before_crl_received"))) {
			request.AuthenticateBeforeCRLReceived = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(key + ".ignore_crl_expiration"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ignore_crl_expiration"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ignore_crl_expiration"))) {
			request.IgnoreCRLExpiration = interfaceToBoolPtr(v)
		}
	}
	if v, ok := d.GetOkExists(key + ".crl_distribution_url"); !isEmptyValue(reflect.ValueOf(d.Get(key+".crl_distribution_url"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".crl_distribution_url"))) {
		request.CrlDistributionURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".automatic_crl_update_period"); !isEmptyValue(reflect.ValueOf(d.Get(key+".automatic_crl_update_period"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".automatic_crl_update_period"))) {
		request.AutomaticCRLUpdatePeriod = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".automatic_crl_update_units"); !isEmptyValue(reflect.ValueOf(d.Get(key+".automatic_crl_update_units"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".automatic_crl_update_units"))) {
		request.AutomaticCRLUpdateUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".non_automatic_crl_update_period"); !isEmptyValue(reflect.ValueOf(d.Get(key+".non_automatic_crl_update_period"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".non_automatic_crl_update_period"))) {
		request.NonAutomaticCRLUpdatePeriod = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".non_automatic_crl_update_units"); !isEmptyValue(reflect.ValueOf(d.Get(key+".non_automatic_crl_update_units"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".non_automatic_crl_update_units"))) {
		request.NonAutomaticCRLUpdateUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".crl_download_failure_retries"); !isEmptyValue(reflect.ValueOf(d.Get(key+".crl_download_failure_retries"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".crl_download_failure_retries"))) {
		request.CrlDownloadFailureRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".crl_download_failure_retries_units"); !isEmptyValue(reflect.ValueOf(d.Get(key+".crl_download_failure_retries_units"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".crl_download_failure_retries_units"))) {
		request.CrlDownloadFailureRetriesUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".download_crl"); !isEmptyValue(reflect.ValueOf(d.Get(key+".download_crl"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".download_crl"))) {
		request.DownloadCRL = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_ocsp_validation"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_ocsp_validation"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_ocsp_validation"))) {
		request.EnableOCSpValidation = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}

	vEnableOCSpValidation, okEnableOCSpValidation := d.GetOk(key + ".enable_ocsp_validation")
	vvEnableOCSpValidation := interfaceToBoolPtr(vEnableOCSpValidation)
	if okEnableOCSpValidation && vvEnableOCSpValidation != nil && *vvEnableOCSpValidation {
		if v, ok := d.GetOkExists(key + ".reject_if_no_status_from_ocs_p"); !isEmptyValue(reflect.ValueOf(d.Get(key+".reject_if_no_status_from_ocs_p"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".reject_if_no_status_from_ocs_p"))) {
			request.RejectIfNoStatusFromOCSP = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(key + ".reject_if_unreachable_from_ocs_p"); !isEmptyValue(reflect.ValueOf(d.Get(key+".reject_if_unreachable_from_ocs_p"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".reject_if_unreachable_from_ocs_p"))) {
			request.RejectIfUnreachableFromOCSP = interfaceToBoolPtr(v)
		}
	}
	if v, ok := d.GetOkExists(key + ".selected_ocsp_service"); !isEmptyValue(reflect.ValueOf(d.Get(key+".selected_ocsp_service"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".selected_ocsp_service"))) {
		request.SelectedOCSpService = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".status"); !isEmptyValue(reflect.ValueOf(d.Get(key+".status"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".status"))) {
		request.Status = interfaceToString(v)
	}
	vTrustForIseAuth, okTrustForIseAuth := d.GetOk(key + ".trust_for_ise_auth")
	vvTrustForIseAuth := interfaceToBoolPtr(vTrustForIseAuth)
	if okTrustForIseAuth && vvTrustForIseAuth != nil && *vvTrustForIseAuth {
		vTrustForClientAuth, okTrustForClientAuth := d.GetOk(key + ".trust_for_client_auth")
		vvTrustForClientAuth := interfaceToBoolPtr(vTrustForClientAuth)
		if okTrustForClientAuth && vvTrustForClientAuth != nil && *vvTrustForClientAuth {
			if v, ok := d.GetOkExists(key + ".trust_for_certificate_based_admin_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".trust_for_certificate_based_admin_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".trust_for_certificate_based_admin_auth"))) {
				request.TrustForCertificateBasedAdminAuth = interfaceToBoolPtr(v)
			}
		}
	}
	if v, ok := d.GetOkExists(key + ".trust_for_cisco_services_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".trust_for_cisco_services_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".trust_for_cisco_services_auth"))) {
		request.TrustForCiscoServicesAuth = interfaceToBoolPtr(v)
	}
	if okTrustForIseAuth && vvTrustForIseAuth != nil && *vvTrustForIseAuth {
		if v, ok := d.GetOkExists(key + ".trust_for_client_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".trust_for_client_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".trust_for_client_auth"))) {
			request.TrustForClientAuth = interfaceToBoolPtr(v)
		}
	}
	if v, ok := d.GetOkExists(key + ".trust_for_ise_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".trust_for_ise_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".trust_for_ise_auth"))) {
		request.TrustForIseAuth = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsCertificatesGetTrustedCertificates(m interface{}, response *isegosdk.ResponseCertificatesGetTrustedCertificates, queryParams *isegosdk.GetTrustedCertificatesQueryParams) []isegosdk.ResponseCertificatesGetTrustedCertificatesResponse {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseCertificatesGetTrustedCertificatesResponse
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
			response, _, err = client.Certificates.GetTrustedCertificates(queryParams)
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

func searchCertificatesGetTrustedCertificates(m interface{}, items []isegosdk.ResponseCertificatesGetTrustedCertificatesResponse, name string, id string) (*isegosdk.ResponseCertificatesGetTrustedCertificateByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseCertificatesGetTrustedCertificateByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseCertificatesGetTrustedCertificateByID
			getItem, _, err = client.Certificates.GetTrustedCertificateByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetTrustedCertificateByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.FriendlyName == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseCertificatesGetTrustedCertificateByID
			getItem, _, err = client.Certificates.GetTrustedCertificateByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetTrustedCertificateByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
