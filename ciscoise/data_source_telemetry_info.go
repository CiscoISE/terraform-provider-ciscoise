package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTelemetryInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTelemetryInfoRead,
		Schema: map[string]*schema.Schema{
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
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"href": &schema.Schema{
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
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"udi_sn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"href": &schema.Schema{
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
					},
				},
			},
		},
	}
}

func dataSourceTelemetryInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	vID, okID := d.GetOk("id")
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")

	method1 := []bool{okID}
	method2 := []bool{okPage, okSize, okFilter, okFilterType}

	if pickMethod(method1, method2) == 1 {
		response1, _, err := client.TelemetryInformation.GetTelemetryInfoByID(vID.(string))
		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTelemetryInfoByID", err,
				"Failure at GetTelemetryInfoByID, unexpected response", ""))
			return diags
		}

		vItem1 := flattenTelemetryInformationGetTelemetryInfoByIDItem(&response1.TelemetryInfo)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTelemetryInfoByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())

		return diags
	}

	queryParams2 := isegosdk.GetTelemetryInformationQueryParams{}
	if okPage {
		queryParams2.Page = vPage.(int)
	}
	if okSize {
		queryParams2.Size = vSize.(int)
	}
	if okFilter {
		queryParams2.Filter = interfaceToSliceString(vFilter)
	}
	if okFilterType {
		queryParams2.FilterType = vFilterType.(string)
	}

	response2, _, err := client.TelemetryInformation.GetTelemetryInformation(&queryParams2)
	if err != nil || response2 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetTelemetryInformation", err,
			"Failure at GetTelemetryInformation, unexpected response", ""))
		return diags
	}

	var items2 []isegosdk.ResponseTelemetryInformationGetTelemetryInformationSearchResultResources
	for len(response2.SearchResult.Resources) > 0 {
		items2 = append(items2, response2.SearchResult.Resources...)
		if response2.SearchResult.NextPage.Rel == "next" {
			href := response2.SearchResult.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			queryParams2.Page = page
			queryParams2.Size = size
			response2, _, err = client.TelemetryInformation.GetTelemetryInformation(&queryParams2)
			if err != nil {
				break
			}
			// All is good, continue to the next page
			continue
		}
		// Does not have next page finish iteration
		break
	}

	vItems2 := flattenTelemetryInformationGetTelemetryInformationItems(&items2)
	if err := d.Set("items", vItems2); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetTelemetryInformation response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())

	return diags
}

func flattenTelemetryInformationGetTelemetryInfoByIDItem(item *isegosdk.ResponseTelemetryInformationGetTelemetryInfoByIDTelemetryInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["status"] = item.Status
	respItem["deployment_id"] = item.DeploymentID
	respItem["udi_sn"] = item.UdiSN
	respItem["link"] = flattenTelemetryInformationGetTelemetryInfoByIDItemLink(item.Link)

	return []map[string]interface{}{
		respItem,
	}
}

func flattenTelemetryInformationGetTelemetryInfoByIDItemLink(item isegosdk.ResponseTelemetryInformationGetTelemetryInfoByIDTelemetryInfoLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}
}

func flattenTelemetryInformationGetTelemetryInformationItems(items *[]isegosdk.ResponseTelemetryInformationGetTelemetryInformationSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["link"] = flattenTelemetryInformationGetTelemetryInformationItemsLink(item.Link)

		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTelemetryInformationGetTelemetryInformationItemsLink(item isegosdk.ResponseTelemetryInformationGetTelemetryInformationSearchResultResourcesLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}
}
