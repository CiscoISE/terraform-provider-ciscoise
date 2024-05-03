package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPsecInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Native IPsec.

- Returns all the IPsec enabled nodes with configuration details.

 This data source supports filtering, sorting and pagination.

The attributes that are suppported for filtering are:


hostName

nadIp

status

authType


The attribute that is suppported for sorting is:


hostName



- Returns the IPsec configuration details of a given node with the hostname and the NAD IP.
`,

		ReadContext: dataSourceIPsecInfoRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Description: `filter query parameter. 
 
 
 
Simple filtering
 should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the 
"filterType=or"
 query string parameter. Each resource Data model description should specify if an attribute is a filtered field. 
 
 
 
 
 
 
OPERATOR
 
DESCRIPTION
 
APPLICABLE ON FIELDS
 
 
 
 
 
EQ
 
Equals
 
authType
 
 
 
NEQ
 
Not Equals
 
authType
 
 
 
EQ
 
Equals
 
hostName
 
 
 
NEQ
 
Not Equals
 
hostName
 
 
 
EQ
 
Equals
 
nadIp
 
 
 
NEQ
 
Not Equals
 
nadIp
 
 
 
EQ
 
Equals
 
status
 
 
 
NEQ
 
Not Equals
 
status
 
 
 
 `,
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter_type": &schema.Schema{
				Description: `filterType query parameter. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"host_name": &schema.Schema{
				Description: `hostName path parameter. Hostname of the deployed node.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nad_ip": &schema.Schema{
				Description: `nadIp path parameter. IP address of the NAD.`,
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
				Description: `sortBy query parameter. Sort column  The IPsec enabled nodes are sorted based on the columns. This is applicable for the field hostName.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cert_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"configure_vti": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"esp_ah_protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"iface": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ike_re_auth_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ike_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"local_internal_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mode_option": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nad_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_one_dhgroup": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_one_encryption_algo": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_one_hash_algo": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_one_life_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"phase_two_dhgroup": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_two_encryption_algo": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_two_hash_algo": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_two_life_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"psk": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_peer_internal_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_time": {
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

						"auth_type": &schema.Schema{
							Description: `Authentication type for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"cert_id": &schema.Schema{
							Description: `ID of the certificate for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"configure_vti": &schema.Schema{
							Description: `Authentication type for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"create_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"esp_ah_protocol": &schema.Schema{
							Description: `Encryption protocol used for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"host_name": &schema.Schema{
							Description: `Hostname of the node`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"iface": &schema.Schema{
							Description: `Ethernet port of the node`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ike_re_auth_time": &schema.Schema{
							Description: `IKE re-authentication time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"ike_version": &schema.Schema{
							Description: `IKE version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"local_internal_ip": &schema.Schema{
							Description: `Local Tunnel IP address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"mode_option": &schema.Schema{
							Description: `The Mode type used for establishing the connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"nad_ip": &schema.Schema{
							Description: `NAD IP address for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"phase_one_dhgroup": &schema.Schema{
							Description: `Phase-one DH group used for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"phase_one_encryption_algo": &schema.Schema{
							Description: `Phase-one encryption algorithm used for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"phase_one_hash_algo": &schema.Schema{
							Description: `Phase-one hashing algorithm used for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"phase_one_life_time": &schema.Schema{
							Description: `Phase-one connection lifetime`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"phase_two_dhgroup": &schema.Schema{
							Description: `Phase-two DH group used for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"phase_two_encryption_algo": &schema.Schema{
							Description: `Phase-two encryption algorithm used for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"phase_two_hash_algo": &schema.Schema{
							Description: `Phase-two hashing algorithm used for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"phase_two_life_time": &schema.Schema{
							Description: `Phase-two connection lifetime`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"psk": &schema.Schema{
							Description: `Pre-shared key used for establishing connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"remote_peer_internal_ip": &schema.Schema{
							Description: `Remote Tunnel IP address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIPsecInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vSort, okSort := d.GetOk("sort")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vHostName, okHostName := d.GetOk("host_name")
	vNadIP, okNadIP := d.GetOk("nad_ip")

	method1 := []bool{okPage, okSize, okFilter, okFilterType, okSort, okSortBy}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostName, okNadIP}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetIPsecEnabledNodes")
		queryParams1 := isegosdk.GetIPsecEnabledNodesQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}
		if okFilter {
			queryParams1.Filter = vFilter.(string)
		}
		if okFilterType {
			queryParams1.FilterType = vFilterType.(string)
		}
		if okSort {
			queryParams1.Sort = vSort.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}

		response1, restyResp1, err := client.NativeIPsec.GetIPsecEnabledNodes(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetIPsecEnabledNodes", err,
				"Failure at GetIPsecEnabledNodes, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseNativeIPsecGetIPsecEnabledNodesResponse
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
				response1, _, err = client.NativeIPsec.GetIPsecEnabledNodes(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenNativeIPsecGetIPsecEnabledNodesItemsResponse(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPsecEnabledNodes response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetIPsecNode")
		vvHostName := vHostName.(string)
		vvNadIP := vNadIP.(string)

		response2, restyResp2, err := client.NativeIPsec.GetIPsecNode(vvHostName, vvNadIP)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetIPsecNode", err,
				"Failure at GetIPsecNode, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNativeIPsecGetIPsecNodeItemResponse(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPsecNode response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNativeIPsecGetIPsecEnabledNodesItemsResponse(items *[]isegosdk.ResponseNativeIPsecGetIPsecEnabledNodesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["auth_type"] = item.AuthType
		respItem["cert_id"] = item.CertID
		respItem["configure_vti"] = boolPtrToString(item.ConfigureVti)
		respItem["create_time"] = item.CreateTime
		respItem["esp_ah_protocol"] = item.EspAhProtocol
		respItem["host_name"] = item.HostName
		respItem["id"] = item.ID
		respItem["iface"] = item.Iface
		respItem["ike_re_auth_time"] = item.IkeReAuthTime
		respItem["ike_version"] = item.IkeVersion
		respItem["local_internal_ip"] = item.LocalInternalIP
		respItem["mode_option"] = item.ModeOption
		respItem["nad_ip"] = item.NadIP
		respItem["phase_one_dhgroup"] = item.PhaseOneDHGroup
		respItem["phase_one_encryption_algo"] = item.PhaseOneEncryptionAlgo
		respItem["phase_one_hash_algo"] = item.PhaseOneHashAlgo
		respItem["phase_one_life_time"] = item.PhaseOneLifeTime
		respItem["phase_two_dhgroup"] = item.PhaseTwoDHGroup
		respItem["phase_two_encryption_algo"] = item.PhaseTwoEncryptionAlgo
		respItem["phase_two_hash_algo"] = item.PhaseTwoHashAlgo
		respItem["phase_two_life_time"] = item.PhaseTwoLifeTime
		respItem["psk"] = item.Psk
		respItem["remote_peer_internal_ip"] = item.RemotePeerInternalIP
		respItem["status"] = item.Status
		respItem["update_time"] = item.UpdateTime
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNativeIPsecGetIPsecNodeItemResponse(item *isegosdk.ResponseNativeIPsecGetIPsecNodeResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["auth_type"] = item.AuthType
	respItem["cert_id"] = item.CertID
	respItem["configure_vti"] = boolPtrToString(item.ConfigureVti)
	respItem["create_time"] = item.CreateTime
	respItem["esp_ah_protocol"] = item.EspAhProtocol
	respItem["host_name"] = item.HostName
	respItem["id"] = item.ID
	respItem["iface"] = item.Iface
	respItem["ike_re_auth_time"] = item.IkeReAuthTime
	respItem["ike_version"] = item.IkeVersion
	respItem["link"] = flattenNativeIPsecGetIPsecNodeItemResponseLink(item.Link)
	respItem["local_internal_ip"] = item.LocalInternalIP
	respItem["mode_option"] = item.ModeOption
	respItem["nad_ip"] = item.NadIP
	respItem["phase_one_dhgroup"] = item.PhaseOneDHGroup
	respItem["phase_one_encryption_algo"] = item.PhaseOneEncryptionAlgo
	respItem["phase_one_hash_algo"] = item.PhaseOneHashAlgo
	respItem["phase_one_life_time"] = item.PhaseOneLifeTime
	respItem["phase_two_dhgroup"] = item.PhaseTwoDHGroup
	respItem["phase_two_encryption_algo"] = item.PhaseTwoEncryptionAlgo
	respItem["phase_two_hash_algo"] = item.PhaseTwoHashAlgo
	respItem["phase_two_life_time"] = item.PhaseTwoLifeTime
	respItem["psk"] = item.Psk
	respItem["remote_peer_internal_ip"] = item.RemotePeerInternalIP
	respItem["status"] = item.Status
	respItem["update_time"] = item.UpdateTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNativeIPsecGetIPsecNodeItemResponseLink(item *isegosdk.ResponseNativeIPsecGetIPsecNodeResponseLink) []map[string]interface{} {
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
