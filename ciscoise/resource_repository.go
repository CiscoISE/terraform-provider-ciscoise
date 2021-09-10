package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceRepositoryCreate,
		ReadContext:   resourceRepositoryRead,
		UpdateContext: resourceRepositoryUpdate,
		DeleteContext: resourceRepositoryDelete,
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

						"enable_pki": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"path": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user_name": &schema.Schema{
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

func resourceRepositoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestRepositoryCreateRepository(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okName && vvName != "" {
		getResponse2, _, err := client.Repository.GetRepository(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.Repository.GetRepositories()
		if response2 != nil && err == nil {
			items2 := getAllItemsRepositoryGetRepositories(m, response2)
			item2, err := searchRepositoryGetRepositories(m, items2, vvName, "")
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	resp1, restyResp1, err := client.Repository.CreateRepository(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateRepository", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateRepository", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		vvName := vName
		log.Printf("[DEBUG] Selected method: GetRepositories")

		response1, _, err := client.Repository.GetRepositories()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRepositories", err,
				"Failure at GetRepositories, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsRepositoryGetRepositories(m, response1)
		item1, err := searchRepositoryGetRepositories(m, items1, vvName, "")
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetRepositories response", err,
				"Failure when searching item from GetRepositories, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRepositories search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetRepository")
		vvName := vName

		response2, _, err := client.Repository.GetRepository(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRepository", err,
				"Failure at GetRepository, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenRepositoryGetRepositoryItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRepository response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceRepositoryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if selectedMethod == 2 {
		vvName = vName
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvName %s", vvName)
		request1 := expandRequestRepositoryUpdateRepository(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Repository.UpdateRepository(vvName, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateRepository", err, restyResp1.String(),
					"Failure at UpdateRepository, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateRepository", err,
				"Failure at UpdateRepository, unexpected response", ""))
			return diags
		}
	}

	return resourceRepositoryRead(ctx, d, m)
}

func resourceRepositoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.Repository.GetRepositories()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsRepositoryGetRepositories(m, getResp1)
		item1, err := searchRepositoryGetRepositories(m, items1, vName, "")
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vName != item1.Name {
			vvName = item1.Name
		} else {
			vvName = vName
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.Repository.GetRepository(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.Repository.DeleteRepository(vvName)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteRepository", err, restyResp1.String(),
				"Failure at DeleteRepository, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteRepository", err,
			"Failure at DeleteRepository, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestRepositoryCreateRepository(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRepositoryCreateRepository {
	request := isegosdk.RequestRepositoryCreateRepository{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".protocol"); !isEmptyValue(reflect.ValueOf(d.Get(key+".protocol"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".protocol"))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".path"); !isEmptyValue(reflect.ValueOf(d.Get(key+".path"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".path"))) {
		request.Path = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".server_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".server_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".server_name"))) {
		request.ServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".user_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".user_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".user_name"))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_pki"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_pki"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_pki"))) {
		request.EnablePki = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRepositoryUpdateRepository(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRepositoryUpdateRepository {
	request := isegosdk.RequestRepositoryUpdateRepository{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".protocol"); !isEmptyValue(reflect.ValueOf(d.Get(key+".protocol"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".protocol"))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".path"); !isEmptyValue(reflect.ValueOf(d.Get(key+".path"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".path"))) {
		request.Path = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".server_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".server_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".server_name"))) {
		request.ServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".user_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".user_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".user_name"))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_pki"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_pki"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_pki"))) {
		request.EnablePki = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsRepositoryGetRepositories(m interface{}, response *isegosdk.ResponseRepositoryGetRepositories) []isegosdk.ResponseRepositoryGetRepositoriesResponse {
	var respItems []isegosdk.ResponseRepositoryGetRepositoriesResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchRepositoryGetRepositories(m interface{}, items []isegosdk.ResponseRepositoryGetRepositoriesResponse, name string, id string) (*isegosdk.ResponseRepositoryGetRepositoryResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseRepositoryGetRepositoryResponse
	for _, item := range items {
		if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseRepositoryGetRepository
			getItem, _, err = client.Repository.GetRepository(name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetRepository")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
