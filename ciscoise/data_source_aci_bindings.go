package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciBindings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAciBindingsRead,
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
			"filter_by": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter_value": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sgt_value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"psn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"learned_from": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"learned_by": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAciBindingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSort, okSort := d.GetOk("sort")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vFilterBy, okFilterBy := d.GetOk("filter_by")
	vFilterValue, okFilterValue := d.GetOk("filter_value")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAciBindings")
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

		response1, _, err := client.AciBindings.GetAciBindings(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAciBindings", err,
				"Failure at GetAciBindings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

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
