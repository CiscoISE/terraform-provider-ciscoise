package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTrustedCertificate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Certificates.

 This data source supports Filtering, Sorting and Pagination.


Filtering and Sorting supported on below mentioned attributes:




friendlyName


subject


issuedTo


issuedBy


validFrom




Supported Date Format: yyyy-MM-dd HH:mm:ss


Supported Operators: EQ, NEQ, GT and LT




expirationDate




Supported Date Format: yyyy-MM-dd HH:mm:ss


Supported Operators: EQ, NEQ, GT and LT




status




Allowed values: enabled, disabled


Supported Operators: EQ, NEQ






This data source can displays details of a Trust Certificate based on a given ID.`,

		ReadContext: dataSourceTrustedCertificateRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Description: `filter query parameter. 
 
 
 
Simple filtering
 should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the 
"filterType=or"
 query string parameter. Each resource Data model description should specify if an attribute is a filtered field. 
 
 
 
 
 
OPERATOR
 
DESCRIPTION
 
 
 
 
 
EQ
 
Equals
 
 
 
NEQ
 
Not Equals
 
 
 
GT
 
Greater Than
 
 
 
LT
 
Less Then
 
 
 
STARTSW
 
Starts With
 
 
 
NSTARTSW
 
Not Starts With
 
 
 
ENDSW
 
Ends With
 
 
 
NENDSW
 
Not Ends With
 
 
 
CONTAINS
 
Contains
 
 
 
NCONTAINS
 
Not Contains
 
 
 
 `,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter_type": &schema.Schema{
				Description: `filterType query parameter. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. The id of the trust certificate`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"page": &schema.Schema{
				Description: `page query parameter. Page number`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. Number of objects returned per page`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"sort": &schema.Schema{
				Description: `sort query parameter. sort type asc or desc`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. sort column by which objects needs to be sorted`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authenticate_before_crl_received": &schema.Schema{
							Description: `Switch to enable/disable authentication before receiving CRL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"automatic_crl_update": &schema.Schema{
							Description: `Switch to enable/disable automatic CRL update`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"automatic_crl_update_period": &schema.Schema{
							Description: `Automatic CRL update period`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"automatic_crl_update_units": &schema.Schema{
							Description: `Unit of time of automatic CRL update`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"crl_distribution_url": &schema.Schema{
							Description: `CRL Distribution URL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"crl_download_failure_retries": &schema.Schema{
							Description: `If CRL download fails, wait time before retry`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"crl_download_failure_retries_units": &schema.Schema{
							Description: `Unit of time before retry if CRL download fails`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `Description of trust certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"download_crl": &schema.Schema{
							Description: `Switch to enable/disable download of CRL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"enable_ocsp_validation": &schema.Schema{
							Description: `Switch to enable/disable OCSP Validation`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"enable_server_identity_check": &schema.Schema{
							Description: `Switch to enable/disable Server Identity Check`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Computed:    true,
						},
						"ignore_crl_expiration": &schema.Schema{
							Description: `Switch to enable/disable ignore CRL Expiration`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"internal_ca": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_referred_in_policy": &schema.Schema{
							Type:     schema.TypeBool,
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
						"non_automatic_crl_update_period": &schema.Schema{
							Description: `Non automatic CRL update period`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"non_automatic_crl_update_units": &schema.Schema{
							Description: `Unit of time of non automatic CRL update`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"reject_if_no_status_from_ocs_p": &schema.Schema{
							Description: `Switch to reject certificate if there is no status from OCSP`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"reject_if_unreachable_from_ocs_p": &schema.Schema{
							Description: `Switch to reject certificate if unreachable from OCSP`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"selected_ocsp_service": &schema.Schema{
							Description: `Name of selected OCSP Service`,
							Type:        schema.TypeString,
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
							Computed: true,
						},
						"subject": &schema.Schema{
							Description: `The Subject or entity with which public key of trust certificate is associated`,
							Type:        schema.TypeString,
							Computed:    true,
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
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authenticate_before_crl_received": &schema.Schema{
							Description: `Switch to enable/disable authentication before receiving CRL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"automatic_crl_update": &schema.Schema{
							Description: `Switch to enable/disable automatic CRL update`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"automatic_crl_update_period": &schema.Schema{
							Description: `Automatic CRL update period`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"automatic_crl_update_units": &schema.Schema{
							Description: `Unit of time of automatic CRL update`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"crl_distribution_url": &schema.Schema{
							Description: `CRL Distribution URL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"crl_download_failure_retries": &schema.Schema{
							Description: `If CRL download fails, wait time before retry`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"crl_download_failure_retries_units": &schema.Schema{
							Description: `Unit of time before retry if CRL download fails`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `Description of trust certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"download_crl": &schema.Schema{
							Description: `Switch to enable/disable download of CRL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"enable_ocsp_validation": &schema.Schema{
							Description: `Switch to enable/disable OCSP Validation`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"enable_server_identity_check": &schema.Schema{
							Description: `Switch to enable/disable Server Identity Check`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Computed:    true,
						},
						"ignore_crl_expiration": &schema.Schema{
							Description: `Switch to enable/disable ignore CRL Expiration`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"internal_ca": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_referred_in_policy": &schema.Schema{
							Type:     schema.TypeBool,
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
						"non_automatic_crl_update_period": &schema.Schema{
							Description: `Non automatic CRL update period`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"non_automatic_crl_update_units": &schema.Schema{
							Description: `Unit of time of non automatic CRL update`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"reject_if_no_status_from_ocs_p": &schema.Schema{
							Description: `Switch to reject certificate if there is no status from OCSP`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"reject_if_unreachable_from_ocs_p": &schema.Schema{
							Description: `Switch to reject certificate if unreachable from OCSP`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"selected_ocsp_service": &schema.Schema{
							Description: `Name of selected OCSP Service`,
							Type:        schema.TypeString,
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
							Computed: true,
						},
						"subject": &schema.Schema{
							Description: `The Subject or entity with which public key of trust certificate is associated`,
							Type:        schema.TypeString,
							Computed:    true,
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

func dataSourceTrustedCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSort, okSort := d.GetOk("sort")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okSort, okSortBy, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTrustedCertificates")
		queryParams1 := isegosdk.GetTrustedCertificatesQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}
		if okSort {
			queryParams1.Sort = vSort.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okFilter {
			queryParams1.Filter = interfaceToSliceString(vFilter)
		}
		if okFilterType {
			queryParams1.FilterType = vFilterType.(string)
		}

		response1, _, err := client.Certificates.GetTrustedCertificates(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTrustedCertificates", err,
				"Failure at GetTrustedCertificates, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseCertificatesGetTrustedCertificatesResponse
		for response1.Response != nil && len(*response1.Response) > 0 {
			items1 = append(items1, *response1.Response...)
			if response1.NextPage != nil && response1.NextPage.Rel == "next" {
				href := response1.NextPage.Href
				page, size, err := getNextPageAndSizeParams(href)
				if err != nil {
					break
				}
				queryParams1.Page = page
				queryParams1.Size = size
				response1, _, err = client.Certificates.GetTrustedCertificates(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenCertificatesGetTrustedCertificatesItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTrustedCertificates response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetTrustedCertificateByID")
		vvID := vID.(string)

		response2, _, err := client.Certificates.GetTrustedCertificateByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTrustedCertificateByID", err,
				"Failure at GetTrustedCertificateByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenCertificatesGetTrustedCertificateByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTrustedCertificateByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenCertificatesGetTrustedCertificatesItems(items *[]isegosdk.ResponseCertificatesGetTrustedCertificatesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["authenticate_before_crl_received"] = item.AuthenticateBeforeCRLReceived
		respItem["automatic_crl_update"] = item.AutomaticCRLUpdate
		respItem["automatic_crl_update_period"] = item.AutomaticCRLUpdatePeriod
		respItem["automatic_crl_update_units"] = item.AutomaticCRLUpdateUnits
		respItem["crl_distribution_url"] = item.CrlDistributionURL
		respItem["crl_download_failure_retries"] = item.CrlDownloadFailureRetries
		respItem["crl_download_failure_retries_units"] = item.CrlDownloadFailureRetriesUnits
		respItem["description"] = item.Description
		respItem["download_crl"] = item.DownloadCRL
		respItem["enable_ocsp_validation"] = item.EnableOCSpValidation
		respItem["enable_server_identity_check"] = item.EnableServerIDentityCheck
		respItem["expiration_date"] = item.ExpirationDate
		respItem["friendly_name"] = item.FriendlyName
		respItem["id"] = item.ID
		respItem["ignore_crl_expiration"] = item.IgnoreCRLExpiration
		respItem["internal_ca"] = item.InternalCa
		respItem["is_referred_in_policy"] = item.IsReferredInPolicy
		respItem["issued_by"] = item.IssuedBy
		respItem["issued_to"] = item.IssuedTo
		respItem["key_size"] = item.KeySize
		respItem["link"] = flattenCertificatesGetTrustedCertificatesItemsLink(item.Link)
		respItem["non_automatic_crl_update_period"] = item.NonAutomaticCRLUpdatePeriod
		respItem["non_automatic_crl_update_units"] = item.NonAutomaticCRLUpdateUnits
		respItem["reject_if_no_status_from_ocs_p"] = item.RejectIfNoStatusFromOCSP
		respItem["reject_if_unreachable_from_ocs_p"] = item.RejectIfUnreachableFromOCSP
		respItem["selected_ocsp_service"] = item.SelectedOCSpService
		respItem["serial_number_decimal_format"] = item.SerialNumberDecimalFormat
		respItem["sha256_fingerprint"] = item.Sha256Fingerprint
		respItem["signature_algorithm"] = item.SignatureAlgorithm
		respItem["status"] = item.Status
		respItem["subject"] = item.Subject
		respItem["trusted_for"] = item.TrustedFor
		respItem["valid_from"] = item.ValidFrom
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenCertificatesGetTrustedCertificatesItemsLink(item *isegosdk.ResponseCertificatesGetTrustedCertificatesResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenCertificatesGetTrustedCertificateByIDItem(item *isegosdk.ResponseCertificatesGetTrustedCertificateByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["authenticate_before_crl_received"] = item.AuthenticateBeforeCRLReceived
	respItem["automatic_crl_update"] = item.AutomaticCRLUpdate
	respItem["automatic_crl_update_period"] = item.AutomaticCRLUpdatePeriod
	respItem["automatic_crl_update_units"] = item.AutomaticCRLUpdateUnits
	respItem["crl_distribution_url"] = item.CrlDistributionURL
	respItem["crl_download_failure_retries"] = item.CrlDownloadFailureRetries
	respItem["crl_download_failure_retries_units"] = item.CrlDownloadFailureRetriesUnits
	respItem["description"] = item.Description
	respItem["download_crl"] = item.DownloadCRL
	respItem["enable_ocsp_validation"] = item.EnableOCSpValidation
	respItem["enable_server_identity_check"] = item.EnableServerIDentityCheck
	respItem["expiration_date"] = item.ExpirationDate
	respItem["friendly_name"] = item.FriendlyName
	respItem["id"] = item.ID
	respItem["ignore_crl_expiration"] = item.IgnoreCRLExpiration
	respItem["internal_ca"] = item.InternalCa
	respItem["is_referred_in_policy"] = item.IsReferredInPolicy
	respItem["issued_by"] = item.IssuedBy
	respItem["issued_to"] = item.IssuedTo
	respItem["key_size"] = item.KeySize
	respItem["link"] = flattenCertificatesGetTrustedCertificateByIDItemLink(item.Link)
	respItem["non_automatic_crl_update_period"] = item.NonAutomaticCRLUpdatePeriod
	respItem["non_automatic_crl_update_units"] = item.NonAutomaticCRLUpdateUnits
	respItem["reject_if_no_status_from_ocs_p"] = item.RejectIfNoStatusFromOCSP
	respItem["reject_if_unreachable_from_ocs_p"] = item.RejectIfUnreachableFromOCSP
	respItem["selected_ocsp_service"] = item.SelectedOCSpService
	respItem["serial_number_decimal_format"] = item.SerialNumberDecimalFormat
	respItem["sha256_fingerprint"] = item.Sha256Fingerprint
	respItem["signature_algorithm"] = item.SignatureAlgorithm
	respItem["status"] = item.Status
	respItem["subject"] = item.Subject
	respItem["trusted_for"] = item.TrustedFor
	respItem["valid_from"] = item.ValidFrom
	return []map[string]interface{}{
		respItem,
	}
}

func flattenCertificatesGetTrustedCertificateByIDItemLink(item *isegosdk.ResponseCertificatesGetTrustedCertificateByIDResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
