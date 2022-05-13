package ciscoise

import (
	"context"
	"fmt"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAncEndpoint() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on ANCEndpoint.

- This resource action allows the client to apply the required configuration.

- This resource action allows the client to clear the required configuration.
`,

		CreateContext: resourceAncEndpointCreate,
		ReadContext:   resourceAncEndpointRead,
		UpdateContext: resourceAncEndpointUpdate,
		DeleteContext: resourceAncEndpointDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"mac_address": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: diffSupressMacAddress(),
						},
						"policy_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
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
						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"policy_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceAncEndpointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AncEndpoint create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	vID, okID := resourceItem["id"]
	vIpAddress, okIpAddress := resourceItem["ip_address"]
	vMacAddress, okMacAddress := resourceItem["mac_address"]
	vPolicyName, okPolicyName := resourceItem["policy_name"]
	var vvID string
	var vvIpAddress string
	var vvMacAddress string
	var vvPolicyName string
	if okID {
		vvID = vID.(string)
	}
	if okIpAddress {
		vvIpAddress = vIpAddress.(string)
	}
	if okMacAddress {
		vvMacAddress = vMacAddress.(string)
	}
	if okPolicyName {
		vvPolicyName = vPolicyName.(string)
	}
	if okID && vvID != "" {
		response2, _, err := client.AncEndpoint.GetAncEndpointByID(vvID)
		if err == nil && response2 != nil && response2.ErsAncEndpoint != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			if vvMacAddress == "" {
				vvMacAddress = response2.ErsAncEndpoint.MacAddress
			}
			if vvPolicyName == "" {
				vvPolicyName = response2.ErsAncEndpoint.PolicyName
			}
			resourceMap["ip_address"] = vvIpAddress
			resourceMap["mac_address"] = vvMacAddress
			resourceMap["policy_name"] = vvPolicyName
			d.SetId(joinResourceID(resourceMap))
			return resourceAncEndpointRead(ctx, d, m)
		}
	} else {
		queryParams1 := isegosdk.GetAncEndpointQueryParams{}
		queryParams1.Filter = []string{fmt.Sprintf("name.EQ.%s", vvPolicyName)}
		response1, _, err := client.AncEndpoint.GetAncEndpoint(&queryParams1)
		if err == nil && response1 != nil {
			items1 := getAllItemsAncEndpointGetAncEndpoint(m, response1, &queryParams1)
			item1, err := searchAncEndpointGetAncEndpoint(m, items1, vvPolicyName, vvMacAddress, vvID)
			if err == nil && item1 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["ip_address"] = vvIpAddress
				resourceMap["mac_address"] = vvMacAddress
				resourceMap["policy_name"] = vvPolicyName
				d.SetId(joinResourceID(resourceMap))
				return resourceAncEndpointRead(ctx, d, m)
			}
		}
	}
	additional_data := []isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalDataAdditionalData{}
	if vvIpAddress != "" {
		ip_address_additional_data := isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalDataAdditionalData{
			Name:  "ipAddress",
			Value: vvIpAddress,
		}
		additional_data = append(additional_data, ip_address_additional_data)
	}
	if vvMacAddress != "" {
		mac_address_additional_data := isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalDataAdditionalData{
			Name:  "macAddress",
			Value: vvMacAddress,
		}
		additional_data = append(additional_data, mac_address_additional_data)
	}
	if vvPolicyName != "" {
		policy_name_additional_data := isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalDataAdditionalData{
			Name:  "policyName",
			Value: vvPolicyName,
		}
		additional_data = append(additional_data, policy_name_additional_data)
	}
	operational_data := isegosdk.RequestAncEndpointApplyAncEndpointOperationAdditionalData{AdditionalData: &additional_data}
	request1 := &isegosdk.RequestAncEndpointApplyAncEndpoint{}
	request1.OperationAdditionalData = &operational_data
	response1, err := client.AncEndpoint.ApplyAncEndpoint(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing ApplyAncEndpoint", err, response1.String(),
				"Failure at ApplyAncEndpoint, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ApplyAncEndpoint", err,
			"Failure at ApplyAncEndpoint, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %s", response1.String())

	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["ip_address"] = vvIpAddress
	resourceMap["mac_address"] = vvMacAddress
	resourceMap["policy_name"] = vvPolicyName
	d.SetId(joinResourceID(resourceMap))
	return resourceAncEndpointRead(ctx, d, m)
}

func resourceAncEndpointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AncEndpoint read for id=[%s]", d.Id())

	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID, okID := resourceMap["id"]
	vvMacAddress, _ := resourceMap["mac_address"]
	vvPolicyName, _ := resourceMap["policy_name"]
	if !okID {
		log.Printf("[DEBUG] Selected method: GetAncEndpoint")
		queryParams1 := isegosdk.GetAncEndpointQueryParams{}
		queryParams1.Filter = []string{fmt.Sprintf("name.EQ.%s", vvPolicyName)}
		response1, restyResp1, err := client.AncEndpoint.GetAncEndpoint(&queryParams1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsAncEndpointGetAncEndpoint(m, response1, &queryParams1)
		item1, err := searchAncEndpointGetAncEndpoint(m, items1, vvPolicyName, vvMacAddress, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}

		vItem1 := flattenAncEndpointGetAncEndpointByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAncEndpoint response",
				err))
			return diags
		}
		return diags
	}
	if okID && vvID != "" {
		log.Printf("[DEBUG] Selected method: GetAncEndpointByID")

		response2, restyResp2, err := client.AncEndpoint.GetAncEndpointByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenAncEndpointGetAncEndpointByIDItem(response2.ErsAncEndpoint)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAncEndpointByID response",
				err))
			return diags
		}
		return diags
	}
	return diags
}

func resourceAncEndpointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AncEndpoint update for id=[%s]", d.Id())
	log.Printf("[DEBUG] Missing AncEndpoint update on Cisco ISE. It will only be update it on Terraform")
	// d.Set("last_updated", getUnixTimeString())
	return resourceAncEndpointRead(ctx, d, m)
}

func resourceAncEndpointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AncEndpoint delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID, okID := resourceMap["id"]
	vvIpAddress, _ := resourceMap["ip_address"]
	vvMacAddress, _ := resourceMap["mac_address"]
	vvPolicyName, _ := resourceMap["policy_name"]

	if !okID {
		queryParams1 := isegosdk.GetAncEndpointQueryParams{}
		queryParams1.Filter = []string{fmt.Sprintf("name.EQ.%s", vvPolicyName)}
		response1, _, err := client.AncEndpoint.GetAncEndpoint(&queryParams1)
		if err != nil || response1 == nil {
			// Assume that element it is already gone
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsAncEndpointGetAncEndpoint(m, response1, &queryParams1)
		item1, err := searchAncEndpointGetAncEndpoint(m, items1, vvPolicyName, vvMacAddress, vvID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if okID && vvID != "" {
		response2, _, err := client.AncEndpoint.GetAncEndpointByID(vvID)
		if err != nil || response2 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	additional_data := []isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalDataAdditionalData{}
	if vvIpAddress != "" {
		ip_address_additional_data := isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalDataAdditionalData{
			Name:  "ipAddress",
			Value: vvIpAddress,
		}
		additional_data = append(additional_data, ip_address_additional_data)
	}
	if vvMacAddress != "" {
		mac_address_additional_data := isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalDataAdditionalData{
			Name:  "macAddress",
			Value: vvMacAddress,
		}
		additional_data = append(additional_data, mac_address_additional_data)
	}
	if vvPolicyName != "" {
		policy_name_additional_data := isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalDataAdditionalData{
			Name:  "policyName",
			Value: vvPolicyName,
		}
		additional_data = append(additional_data, policy_name_additional_data)
	}
	operational_data := isegosdk.RequestAncEndpointClearAncEndpointOperationAdditionalData{AdditionalData: &additional_data}
	request1 := &isegosdk.RequestAncEndpointClearAncEndpoint{}
	request1.OperationAdditionalData = &operational_data
	response1, err := client.AncEndpoint.ClearAncEndpoint(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing ClearAncEndpoint", err, response1.String(),
				"Failure at ClearAncEndpoint, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ClearAncEndpoint", err,
			"Failure at ClearAncEndpoint, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %s", response1.String())

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func getAllItemsAncEndpointGetAncEndpoint(m interface{}, response *isegosdk.ResponseAncEndpointGetAncEndpoint, queryParams *isegosdk.GetAncEndpointQueryParams) []isegosdk.ResponseAncEndpointGetAncEndpointSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseAncEndpointGetAncEndpointSearchResultResources
	for response.SearchResult != nil && response.SearchResult.Resources != nil && len(*response.SearchResult.Resources) > 0 {
		respItems = append(respItems, *response.SearchResult.Resources...)
		if response.SearchResult.NextPage != nil && response.SearchResult.NextPage.Rel == "next" {
			href := response.SearchResult.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			queryParams.Page = page
			queryParams.Size = size
			response, _, err = client.AncEndpoint.GetAncEndpoint(queryParams)
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

func searchAncEndpointGetAncEndpoint(m interface{}, items []isegosdk.ResponseAncEndpointGetAncEndpointSearchResultResources, name string, mac_address string, id string) (*isegosdk.ResponseAncEndpointGetAncEndpointByIDErsAncEndpoint, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseAncEndpointGetAncEndpointByIDErsAncEndpoint
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseAncEndpointGetAncEndpointByID
			getItem, _, err = client.AncEndpoint.GetAncEndpointByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSxpVpnByID")
			}
			foundItem = getItem.ErsAncEndpoint
			return foundItem, err
		} else {
			var getItem *isegosdk.ResponseAncEndpointGetAncEndpointByID
			getItem, _, err = client.AncEndpoint.GetAncEndpointByID(item.ID)
			if err != nil {
				continue
			}
			if getItem == nil {
				continue
			}
			is_anc_endpoint := getItem.ErsAncEndpoint.PolicyName == name
			is_anc_endpoint = is_anc_endpoint && compareMacAddress(getItem.ErsAncEndpoint.MacAddress, mac_address)
			if getItem.ErsAncEndpoint != nil && is_anc_endpoint {
				foundItem = getItem.ErsAncEndpoint
				return foundItem, err
			}
		}
	}
	return foundItem, err
}
