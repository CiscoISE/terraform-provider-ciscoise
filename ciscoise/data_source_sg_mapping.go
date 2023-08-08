package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSgMapping() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on IPToSGTMapping.

- This data source allows the client to get an IP to SGT mapping by ID.

- This data source allows the client to get all the IP to SGT mappings.

Filter:

[hostName, groupName, ip, sgtName]

Sorting:

[hostName, groupName, ip, sgtName]
`,

		ReadContext: dataSourceSgMappingRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Description: `filter query parameter. 

**Simple filtering** should be available through the filter query string parameter. The structure of a filter is
a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator
common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query
string parameter. Each resource Data model description should specify if an attribute is a filtered field.



              Operator    | Description 

              ------------|----------------

              EQ          | Equals 

              NEQ         | Not Equals 

              GT          | Greater Than 

              LT          | Less Then 

              STARTSW     | Starts With 

              NSTARTSW    | Not Starts With 

              ENDSW       | Ends With 

              NENDSW      | Not Ends With 

              CONTAINS	  | Contains 

              NCONTAINS	  | Not Contains 

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
				Description: `id path parameter.`,
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
			"sortasc": &schema.Schema{
				Description: `sortasc query parameter. sort asc`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sortdsc": &schema.Schema{
				Description: `sortdsc query parameter. sort desc`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"deploy_to": &schema.Schema{
							Description: `Mandatory unless mappingGroup is set or unless deployType=ALL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"deploy_type": &schema.Schema{
							Description: `Allowed values:
- ALL,
- ND,
- NDG`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_ip": &schema.Schema{
							Description: `Mandatory if hostName is empty -- valid IP`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"host_name": &schema.Schema{
							Description: `Mandatory if hostIp is empty`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
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
						"mapping_group": &schema.Schema{
							Description: `Mapping Group Id. Mandatory unless sgt and deployTo and deployType are set`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sgt": &schema.Schema{
							Description: `Mandatory unless mappingGroup is set`,
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

						"id": &schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSgMappingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSortasc, okSortasc := d.GetOk("sortasc")
	vSortdsc, okSortdsc := d.GetOk("sortdsc")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okSortasc, okSortdsc, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetIPToSgtMapping")
		queryParams1 := isegosdk.GetIPToSgtMappingQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}
		if okSortasc {
			queryParams1.Sortasc = vSortasc.(string)
		}
		if okSortdsc {
			queryParams1.Sortdsc = vSortdsc.(string)
		}
		if okFilter {
			queryParams1.Filter = interfaceToSliceString(vFilter)
		}
		if okFilterType {
			queryParams1.FilterType = vFilterType.(string)
		}

		response1, restyResp1, err := client.IPToSgtMapping.GetIPToSgtMapping(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetIPToSgtMapping", err,
				"Failure at GetIPToSgtMapping, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingSearchResultResources
		for response1.SearchResult != nil && response1.SearchResult.Resources != nil && len(*response1.SearchResult.Resources) > 0 {
			items1 = append(items1, *response1.SearchResult.Resources...)
			if response1.SearchResult.NextPage != nil && response1.SearchResult.NextPage.Rel == "next" {
				href := response1.SearchResult.NextPage.Href
				page, size, err := getNextPageAndSizeParams(href)
				if err != nil {
					break
				}
				queryParams1.Page = page
				queryParams1.Size = size
				response1, _, err = client.IPToSgtMapping.GetIPToSgtMapping(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenIPToSgtMappingGetIPToSgtMappingItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPToSgtMapping response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetIPToSgtMappingByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.IPToSgtMapping.GetIPToSgtMappingByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetIPToSgtMappingByID", err,
				"Failure at GetIPToSgtMappingByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenIPToSgtMappingGetIPToSgtMappingByIDItem(response2.SgMapping)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPToSgtMappingByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIPToSgtMappingGetIPToSgtMappingItems(items *[]isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["link"] = flattenIPToSgtMappingGetIPToSgtMappingItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIPToSgtMappingGetIPToSgtMappingItemsLink(item *isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingSearchResultResourcesLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenIPToSgtMappingGetIPToSgtMappingByIDItem(item *isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingByIDSgMapping) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["sgt"] = item.Sgt
	respItem["deploy_to"] = item.DeployTo
	respItem["deploy_type"] = item.DeployType
	respItem["host_name"] = item.HostName
	respItem["host_ip"] = item.HostIP
	respItem["mapping_group"] = item.MappingGroup
	respItem["link"] = flattenIPToSgtMappingGetIPToSgtMappingByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIPToSgtMappingGetIPToSgtMappingByIDItemLink(item *isegosdk.ResponseIPToSgtMappingGetIPToSgtMappingByIDSgMappingLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
