package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciBindings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ACIBindings.

- This data source allows clients to retrieve all the bindings that were sent to Cisco ISE by ACI or received on ACI
from Cisco ISE.The binding information will be identical to the information on ACI bindings page in the Cisco ISE UI.
Filtering will be based on one attribute only, such as ip/sgt/vn/psn/learnedFrom/learnedBy with CONTAINS mode of search.
`,

		ReadContext: dataSourceAciBindingsRead,
		Schema: map[string]*schema.Schema{
			"filter_by": &schema.Schema{
				Description: `filterBy query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter_value": &schema.Schema{
				Description: `filterValue query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Resource UUID value`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ip": &schema.Schema{
							Description: `Binding IPv4 address. Each binding will be exclusively identified by its IP address and virtual network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"learned_by": &schema.Schema{
							Description: `Binding Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"learned_from": &schema.Schema{
							Description: `Binding Source`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Resource Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"psn": &schema.Schema{
							Description: `Cisco ISE Policy Service node (PSN) IP address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sgt_value": &schema.Schema{
							Description: `Security Group Tag (SGT) value. The valid range for SGT values is 0-65534`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"vn": &schema.Schema{
							Description: `Virtual network. Each binding will be exclusively identified by its IP address and virtual network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAciBindingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSort, okSort := d.GetOk("sort")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vFilterBy, okFilterBy := d.GetOk("filter_by")
	vFilterValue, okFilterValue := d.GetOk("filter_value")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAciBindings")
		queryParams1 := isegosdk.GetAciBindingsQueryParams{}

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
		if okFilterBy {
			queryParams1.FilterBy = interfaceToSliceString(vFilterBy)
		}
		if okFilterValue {
			queryParams1.FilterValue = interfaceToSliceString(vFilterValue)
		}

		response1, restyResp1, err := client.AciBindings.GetAciBindings(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAciBindings", err,
				"Failure at GetAciBindings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenAciBindingsGetAciBindingsItem(response1.AciBindings)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAciBindings response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAciBindingsGetAciBindingsItem(item *isegosdk.ResponseAciBindingsGetAciBindingsAciBindings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["ip"] = item.IP
	respItem["sgt_value"] = item.SgtValue
	respItem["vn"] = item.Vn
	respItem["psn"] = item.Psn
	respItem["learned_from"] = item.LearnedFrom
	respItem["learned_by"] = item.LearnedBy
	return []map[string]interface{}{
		respItem,
	}
}
