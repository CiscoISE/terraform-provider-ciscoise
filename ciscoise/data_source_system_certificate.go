package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemCertificate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSystemCertificateRead,
		Schema: map[string]*schema.Schema{
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
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"expiration_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
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
							Type:     schema.TypeInt,
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
						"portals_using_the_tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"expiration_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
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
							Type:     schema.TypeInt,
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
						"portals_using_the_tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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

func dataSourceSystemCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vHostName, okHostName := d.GetOk("host_name")
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSort, okSort := d.GetOk("sort")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vID, okID := d.GetOk("id")

	method1 := []bool{okHostName, okPage, okSize, okSort, okSortBy, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostName, okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSystemCertificates")
		vvHostName := vHostName.(string)
		queryParams1 := isegosdk.GetSystemCertificatesQueryParams{}

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

		response1, _, err := client.Certificates.GetSystemCertificates(vvHostName, &queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSystemCertificates", err,
				"Failure at GetSystemCertificates, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseCertificatesGetSystemCertificatesResponse
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
				response1, _, err = client.Certificates.GetSystemCertificates(vvHostName, &queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenCertificatesGetSystemCertificatesItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSystemCertificates response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetSystemCertificateByID")
		vvHostName := vHostName.(string)
		vvID := vID.(string)

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
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenCertificatesGetSystemCertificatesItems(items *[]isegosdk.ResponseCertificatesGetSystemCertificatesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["expiration_date"] = item.ExpirationDate
		respItem["friendly_name"] = item.FriendlyName
		respItem["group_tag"] = item.GroupTag
		respItem["id"] = item.ID
		respItem["issued_by"] = item.IssuedBy
		respItem["issued_to"] = item.IssuedTo
		respItem["key_size"] = item.KeySize
		respItem["link"] = flattenCertificatesGetSystemCertificatesItemsLink(item.Link)
		respItem["portals_using_the_tag"] = item.PortalsUsingTheTag
		respItem["self_signed"] = item.SelfSigned
		respItem["serial_number_decimal_format"] = item.SerialNumberDecimalFormat
		respItem["sha256_fingerprint"] = item.Sha256Fingerprint
		respItem["signature_algorithm"] = item.SignatureAlgorithm
		respItem["used_by"] = item.UsedBy
		respItem["valid_from"] = item.ValidFrom
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenCertificatesGetSystemCertificatesItemsLink(item *isegosdk.ResponseCertificatesGetSystemCertificatesResponseLink) []map[string]interface{} {
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

func flattenCertificatesGetSystemCertificateByIDItem(item *isegosdk.ResponseCertificatesGetSystemCertificateByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["expiration_date"] = item.ExpirationDate
	respItem["friendly_name"] = item.FriendlyName
	respItem["group_tag"] = item.GroupTag
	respItem["id"] = item.ID
	respItem["issued_by"] = item.IssuedBy
	respItem["issued_to"] = item.IssuedTo
	respItem["key_size"] = item.KeySize
	respItem["link"] = flattenCertificatesGetSystemCertificateByIDItemLink(item.Link)
	respItem["portals_using_the_tag"] = item.PortalsUsingTheTag
	respItem["self_signed"] = item.SelfSigned
	respItem["serial_number_decimal_format"] = item.SerialNumberDecimalFormat
	respItem["sha256_fingerprint"] = item.Sha256Fingerprint
	respItem["signature_algorithm"] = item.SignatureAlgorithm
	respItem["used_by"] = item.UsedBy
	respItem["valid_from"] = item.ValidFrom
	return []map[string]interface{}{
		respItem,
	}
}

func flattenCertificatesGetSystemCertificateByIDItemLink(item *isegosdk.ResponseCertificatesGetSystemCertificateByIDResponseLink) []map[string]interface{} {
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
