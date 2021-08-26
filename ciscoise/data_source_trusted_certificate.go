package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTrustedCertificate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTrustedCertificateRead,
		Schema: map[string]*schema.Schema{
			"page": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sort": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authenticate_before_crl_received": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"automatic_crl_update": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"automatic_crl_update_period": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"automatic_crl_update_units": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"crl_distribution_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"crl_download_failure_retries": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"crl_download_failure_retries_units": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"download_crl": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_ocsp_validation": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_server_identity_check": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"expiration_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ignore_crl_expiration": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"issued_to": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"non_automatic_crl_update_units": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reject_if_no_status_from_ocs_p": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reject_if_unreachable_from_ocs_p": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"selected_ocsp_service": &schema.Schema{
							Type:     schema.TypeString,
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
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"subject": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"trusted_for": &schema.Schema{
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
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authenticate_before_crl_received": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"automatic_crl_update": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"automatic_crl_update_period": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"automatic_crl_update_units": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"crl_distribution_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"crl_download_failure_retries": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"crl_download_failure_retries_units": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"download_crl": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_ocsp_validation": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_server_identity_check": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"expiration_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ignore_crl_expiration": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"issued_to": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"non_automatic_crl_update_units": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reject_if_no_status_from_ocs_p": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reject_if_unreachable_from_ocs_p": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"selected_ocsp_service": &schema.Schema{
							Type:     schema.TypeString,
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
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"subject": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"trusted_for": &schema.Schema{
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
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

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
			if response1.NextPage.Rel == "next" {
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
