package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFilterPolicy() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceFilterPolicyCreate,
		ReadContext:   resourceFilterPolicyRead,
		UpdateContext: resourceFilterPolicyUpdate,
		DeleteContext: resourceFilterPolicyDelete,
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

						"domains": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sgt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceFilterPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceFilterPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceFilterPolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceFilterPolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func expandRequestFilterPolicyCreateFilterPolicy(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestFilterPolicyCreateFilterPolicy {
	request := isegosdk.RequestFilterPolicyCreateFilterPolicy{}
	request.ERSFilterPolicy = expandRequestFilterPolicyCreateFilterPolicyERSFilterPolicy(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestFilterPolicyCreateFilterPolicyERSFilterPolicy(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestFilterPolicyCreateFilterPolicyERSFilterPolicy {
	request := isegosdk.RequestFilterPolicyCreateFilterPolicyERSFilterPolicy{}
	if v, ok := d.GetOkExists(key + ".subnet"); !isEmptyValue(reflect.ValueOf(d.Get(key+".subnet"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".subnet"))) {
		request.Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".domains"); !isEmptyValue(reflect.ValueOf(d.Get(key+".domains"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".domains"))) {
		request.Domains = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sgt"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sgt"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sgt"))) {
		request.Sgt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".vn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".vn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".vn"))) {
		request.Vn = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestFilterPolicyUpdateFilterPolicyByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestFilterPolicyUpdateFilterPolicyByID {
	request := isegosdk.RequestFilterPolicyUpdateFilterPolicyByID{}
	request.ERSFilterPolicy = expandRequestFilterPolicyUpdateFilterPolicyByIDERSFilterPolicy(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestFilterPolicyUpdateFilterPolicyByIDERSFilterPolicy(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestFilterPolicyUpdateFilterPolicyByIDERSFilterPolicy {
	request := isegosdk.RequestFilterPolicyUpdateFilterPolicyByIDERSFilterPolicy{}
	if v, ok := d.GetOkExists(key + ".subnet"); !isEmptyValue(reflect.ValueOf(d.Get(key+".subnet"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".subnet"))) {
		request.Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".domains"); !isEmptyValue(reflect.ValueOf(d.Get(key+".domains"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".domains"))) {
		request.Domains = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sgt"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sgt"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sgt"))) {
		request.Sgt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".vn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".vn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".vn"))) {
		request.Vn = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsFilterPolicyGetFilterPolicy(m interface{}, response *isegosdk.ResponseFilterPolicyGetFilterPolicy, queryParams *isegosdk.GetFilterPolicyQueryParams) []isegosdk.ResponseFilterPolicyGetFilterPolicySearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseFilterPolicyGetFilterPolicySearchResultResources
	for response.SearchResult != nil && response.SearchResult.Resources != nil && len(*response.SearchResult.Resources) > 0 {
		respItems = append(respItems, *response.SearchResult.Resources...)
		if response.SearchResult.NextPage != nil && response.SearchResult.NextPage.Rel == "next" {
			href := response.SearchResult.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			if queryParams != nil {
				queryParams.Page = page
				queryParams.Size = size
			}
			response, _, err = client.FilterPolicy.GetFilterPolicy(queryParams)
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

func searchFilterPolicyGetFilterPolicy(m interface{}, items []isegosdk.ResponseFilterPolicyGetFilterPolicySearchResultResources, name string, id string) (*isegosdk.ResponseFilterPolicyGetFilterPolicyByIDERSFilterPolicy, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseFilterPolicyGetFilterPolicyByIDERSFilterPolicy
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseFilterPolicyGetFilterPolicyByID
			getItem, _, err = client.FilterPolicy.GetFilterPolicyByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetFilterPolicyByID")
			}
			foundItem = getItem.ERSFilterPolicy
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseFilterPolicyGetFilterPolicyByID
			getItem, _, err = client.FilterPolicy.GetFilterPolicyByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetFilterPolicyByID")
			}
			foundItem = getItem.ERSFilterPolicy
			return foundItem, err
		}
	}
	return foundItem, err
}
