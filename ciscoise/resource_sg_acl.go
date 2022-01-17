package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSgACL() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SecurityGroupsACLs.

- This resource allows the client to update a security group ACL.

- This resource deletes a security group ACL.

- This resource creates a security group ACL.
`,

		CreateContext: resourceSgACLCreate,
		ReadContext:   resourceSgACLRead,
		UpdateContext: resourceSgACLUpdate,
		DeleteContext: resourceSgACLDelete,
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
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aclcontent": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"generation_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_version": &schema.Schema{
							Description: `Allowed values:
- IPV4,
- IPV6,
- IP_AGNOSTIC`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_read_only": &schema.Schema{
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
						"modelled_content": &schema.Schema{
							Description: `Modelled content of contract`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aclcontent": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"generation_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_version": &schema.Schema{
							Description: `Allowed values:
- IPV4,
- IPV6,
- IP_AGNOSTIC`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"is_read_only": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"modelled_content": &schema.Schema{
							Description: `Modelled content of contract`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSgACLCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgACL create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSgACLCreateSecurityGroupsACL(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.SecurityGroupsACLs.GetSecurityGroupsACLByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceSgACLRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetSecurityGroupsACLQueryParams{}

		response2, _, err := client.SecurityGroupsACLs.GetSecurityGroupsACL(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsSecurityGroupsACLsGetSecurityGroupsACL(m, response2, &queryParams2)
			item2, err := searchSecurityGroupsACLsGetSecurityGroupsACL(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceSgACLRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.SecurityGroupsACLs.CreateSecurityGroupsACL(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSecurityGroupsACL", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSecurityGroupsACL", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceSgACLRead(ctx, d, m)
}

func resourceSgACLRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgACL read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		vvName := vName
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetSecurityGroupsACL")
		queryParams1 := isegosdk.GetSecurityGroupsACLQueryParams{}

		response1, restyResp1, err := client.SecurityGroupsACLs.GetSecurityGroupsACL(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsSecurityGroupsACLsGetSecurityGroupsACL(m, response1, &queryParams1)
		item1, err := searchSecurityGroupsACLsGetSecurityGroupsACL(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenSecurityGroupsACLsGetSecurityGroupsACLByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityGroupsACL search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSecurityGroupsACLByID")
		vvID := vID

		response2, restyResp2, err := client.SecurityGroupsACLs.GetSecurityGroupsACLByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSecurityGroupsACLsGetSecurityGroupsACLByIDItem(response2.Sgacl)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityGroupsACLByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSgACLUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgACL update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetSecurityGroupsACLQueryParams{}

		getResp1, _, err := client.SecurityGroupsACLs.GetSecurityGroupsACL(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsSecurityGroupsACLsGetSecurityGroupsACL(m, getResp1, &queryParams1)
			item1, err := searchSecurityGroupsACLsGetSecurityGroupsACL(m, items1, vName, vID)
			if err == nil && item1 != nil {
				if vID != item1.ID {
					vvID = item1.ID
				} else {
					vvID = vID
				}
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestSgACLUpdateSecurityGroupsACLByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.SecurityGroupsACLs.UpdateSecurityGroupsACLByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSecurityGroupsACLByID", err, restyResp1.String(),
					"Failure at UpdateSecurityGroupsACLByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSecurityGroupsACLByID", err,
				"Failure at UpdateSecurityGroupsACLByID, unexpected response", ""))
			return diags
		}
	}

	return resourceSgACLRead(ctx, d, m)
}

func resourceSgACLDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgACL delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetSecurityGroupsACLQueryParams{}

		getResp1, _, err := client.SecurityGroupsACLs.GetSecurityGroupsACL(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSecurityGroupsACLsGetSecurityGroupsACL(m, getResp1, &queryParams1)
		item1, err := searchSecurityGroupsACLsGetSecurityGroupsACL(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.SecurityGroupsACLs.GetSecurityGroupsACLByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.SecurityGroupsACLs.DeleteSecurityGroupsACLByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSecurityGroupsACLByID", err, restyResp1.String(),
				"Failure at DeleteSecurityGroupsACLByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSecurityGroupsACLByID", err,
			"Failure at DeleteSecurityGroupsACLByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSgACLCreateSecurityGroupsACL(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsACLsCreateSecurityGroupsACL {
	request := isegosdk.RequestSecurityGroupsACLsCreateSecurityGroupsACL{}
	request.Sgacl = expandRequestSgACLCreateSecurityGroupsACLSgacl(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgACLCreateSecurityGroupsACLSgacl(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsACLsCreateSecurityGroupsACLSgacl {
	request := isegosdk.RequestSecurityGroupsACLsCreateSecurityGroupsACLSgacl{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".generation_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".generation_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".generation_id")))) {
		request.GenerationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aclcontent")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aclcontent")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aclcontent")))) {
		request.ACLcontent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_read_only")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_read_only")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_read_only")))) {
		request.IsReadOnly = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".modelled_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".modelled_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".modelled_content")))) {
		request.ModelledContent = expandRequestSgACLCreateSecurityGroupsACLSgaclModelledContent(ctx, key+".modelled_content.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_version")))) {
		request.IPVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgACLCreateSecurityGroupsACLSgaclModelledContent(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsACLsCreateSecurityGroupsACLSgaclModelledContent {
	var request isegosdk.RequestSecurityGroupsACLsCreateSecurityGroupsACLSgaclModelledContent
	keyValue := d.Get(fixKeyAccess(key))
	request = requestStringToInterface(interfaceToString(keyValue))
	return &request
}

func expandRequestSgACLUpdateSecurityGroupsACLByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsACLsUpdateSecurityGroupsACLByID {
	request := isegosdk.RequestSecurityGroupsACLsUpdateSecurityGroupsACLByID{}
	request.Sgacl = expandRequestSgACLUpdateSecurityGroupsACLByIDSgacl(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgACLUpdateSecurityGroupsACLByIDSgacl(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsACLsUpdateSecurityGroupsACLByIDSgacl {
	request := isegosdk.RequestSecurityGroupsACLsUpdateSecurityGroupsACLByIDSgacl{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".generation_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".generation_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".generation_id")))) {
		request.GenerationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aclcontent")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aclcontent")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aclcontent")))) {
		request.ACLcontent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_read_only")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_read_only")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_read_only")))) {
		request.IsReadOnly = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".modelled_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".modelled_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".modelled_content")))) {
		request.ModelledContent = expandRequestSgACLUpdateSecurityGroupsACLByIDSgaclModelledContent(ctx, key+".modelled_content.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_version")))) {
		request.IPVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgACLUpdateSecurityGroupsACLByIDSgaclModelledContent(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupsACLsUpdateSecurityGroupsACLByIDSgaclModelledContent {
	var request isegosdk.RequestSecurityGroupsACLsUpdateSecurityGroupsACLByIDSgaclModelledContent
	keyValue := d.Get(fixKeyAccess(key))
	request = requestStringToInterface(interfaceToString(keyValue))
	return &request
}

func getAllItemsSecurityGroupsACLsGetSecurityGroupsACL(m interface{}, response *isegosdk.ResponseSecurityGroupsACLsGetSecurityGroupsACL, queryParams *isegosdk.GetSecurityGroupsACLQueryParams) []isegosdk.ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultResources
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
			response, _, err = client.SecurityGroupsACLs.GetSecurityGroupsACL(queryParams)
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

func searchSecurityGroupsACLsGetSecurityGroupsACL(m interface{}, items []isegosdk.ResponseSecurityGroupsACLsGetSecurityGroupsACLSearchResultResources, name string, id string) (*isegosdk.ResponseSecurityGroupsACLsGetSecurityGroupsACLByIDSgacl, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSecurityGroupsACLsGetSecurityGroupsACLByIDSgacl
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSecurityGroupsACLsGetSecurityGroupsACLByID
			getItem, _, err = client.SecurityGroupsACLs.GetSecurityGroupsACLByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSecurityGroupsACLByID")
			}
			foundItem = getItem.Sgacl
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSecurityGroupsACLsGetSecurityGroupsACLByID
			getItem, _, err = client.SecurityGroupsACLs.GetSecurityGroupsACLByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSecurityGroupsACLByID")
			}
			foundItem = getItem.Sgacl
			return foundItem, err
		}
	}
	return foundItem, err
}
