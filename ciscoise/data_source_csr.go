package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCsr() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Certificates.

- This data source supports Filtering, Sorting and Pagination.

Filtering and Sorting supported on below mentioned attributes:


friendlyName

subject

timeStamp


Supported Date Format: yyyy-MM-dd HH:mm:ss.SSS

Supported Operators: EQ, NEQ, GT and LT




- This data source displays details of a Certificate Signing Request of a particular node for given HostName and ID.
`,

		ReadContext: dataSourceCsrRead,
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
			"host_name": &schema.Schema{
				Description: `hostName path parameter. Name of the host of which CSR's should be returned`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. ID of the Certificate Signing Request returned`,
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

						"csr_contents": &schema.Schema{
							Description: `Contents of the certificate file.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"friendly_name": &schema.Schema{
							Description: `Friendly name of the certificate.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"group_tag": &schema.Schema{
							Description: `GroupTag of the certificate.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"host_name": &schema.Schema{
							Description: `Hostname or IP address of the Cisco ISE node.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `ID of the certificate.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"key_size": &schema.Schema{
							Description: `Size of the cryptographic key used.`,
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
						"signature_algorithm": &schema.Schema{
							Description: `Algorithm used for encrypting CSR`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"subject": &schema.Schema{
							Description: `Subject of the certificate. Includes Common Name (CN), Organizational Unit (OU), etc.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"time_stamp": &schema.Schema{
							Description: `Timestamp of the certificate generation.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"used_for": &schema.Schema{
							Description: `Services for which the certificate is used for (for eg- MGMT, GENERIC).`,
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

						"friendly_name": &schema.Schema{
							Description: `Friendly name of the certificate.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"group_tag": &schema.Schema{
							Description: `GroupTag of the certificate.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"host_name": &schema.Schema{
							Description: `Hostname or IP address of the Cisco ISE node.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `ID of the certificate.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"key_size": &schema.Schema{
							Description: `Size of the cryptographic key used.`,
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
						"signature_algorithm": &schema.Schema{
							Description: `Algorithm used for encrypting CSR`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"subject": &schema.Schema{
							Description: `Subject of the certificate. Includes Common Name (CN), Organizational Unit (OU), etc.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"time_stamp": &schema.Schema{
							Description: `Timestamp of the certificate generation.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"used_for": &schema.Schema{
							Description: `Services for which the certificate is used for (for eg- MGMT, GENERIC).`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCsrRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSort, okSort := d.GetOk("sort")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vHostName, okHostName := d.GetOk("host_name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okSort, okSortBy, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostName, okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCsrs")
		queryParams1 := isegosdk.GetCsrsQueryParams{}

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

		response1, restyResp1, err := client.Certificates.GetCsrs(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetCsrs", err,
				"Failure at GetCsrs, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseCertificatesGetCsrsResponse
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
				response1, _, err = client.Certificates.GetCsrs(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenCertificatesGetCsrsItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCsrs response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetCsrByID")
		vvHostName := vHostName.(string)
		vvID := vID.(string)

		response2, restyResp2, err := client.Certificates.GetCsrByID(vvHostName, vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetCsrByID", err,
				"Failure at GetCsrByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenCertificatesGetCsrByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCsrByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenCertificatesGetCsrsItems(items *[]isegosdk.ResponseCertificatesGetCsrsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["friendly_name"] = item.FriendlyName
		respItem["group_tag"] = item.GroupTag
		respItem["host_name"] = item.HostName
		respItem["id"] = item.ID
		respItem["key_size"] = item.KeySize
		respItem["link"] = flattenCertificatesGetCsrsItemsLink(item.Link)
		respItem["signature_algorithm"] = item.SignatureAlgorithm
		respItem["subject"] = item.Subject
		respItem["time_stamp"] = item.TimeStamp
		respItem["used_for"] = item.UsedFor
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenCertificatesGetCsrsItemsLink(item *isegosdk.ResponseCertificatesGetCsrsResponseLink) []map[string]interface{} {
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

func flattenCertificatesGetCsrByIDItem(item *isegosdk.ResponseCertificatesGetCsrByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["csr_contents"] = item.CsrContents
	respItem["friendly_name"] = item.FriendlyName
	respItem["group_tag"] = item.GroupTag
	respItem["host_name"] = item.HostName
	respItem["id"] = item.ID
	respItem["key_size"] = item.KeySize
	respItem["link"] = flattenCertificatesGetCsrByIDItemLink(item.Link)
	respItem["signature_algorithm"] = item.SignatureAlgorithm
	respItem["subject"] = item.Subject
	respItem["time_stamp"] = item.TimeStamp
	respItem["used_for"] = item.UsedFor
	return []map[string]interface{}{
		respItem,
	}
}

func flattenCertificatesGetCsrByIDItemLink(item *isegosdk.ResponseCertificatesGetCsrByIDResponseLink) []map[string]interface{} {
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
