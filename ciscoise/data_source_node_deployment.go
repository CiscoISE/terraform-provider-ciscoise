package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodeDeployment() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Node Deployment.

- The API lists all the nodes that are deployed in the cluster.
 Returns basic information about each of the deployed nodes in the cluster like hostname, status, roles, and services.
 Supports filtering on FQDN, hostname, IP address, roles, services and node status.

- This data source retrieves detailed information of the deployed node.
`,

		ReadContext: dataSourceNodeDeploymentRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Description: `filter query parameter. 
 
 
 
Simple filtering
 is available through the filter query string parameter. The structure of a filter is a triplet of field operator and value, separated by dots. More than one filter can be sent. The logical operator common to all filter criteria is AND by default, and can be changed by using the 
"filterType=or"
 query string parameter. Each resource Data model description should specify if an attribute is a filtered field. 
 
 
 
 
 
OPERATOR
 
DESCRIPTION
 
 
 
 
 
EQ
 
Equals
 
 
 
NEQ
 
Not Equals
 
 
 
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
				Description: `filterType query parameter. The logical operator common to all filter criteria is AND by default, and can be changed by using this parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hostname": &schema.Schema{
				Description: `hostname path parameter. Hostname of the deployed node.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": &schema.Schema{
							Description: `Roles can be empty or have many values for a node. `,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"services": &schema.Schema{
							Description: `Services can be empty or have many values for a node. `,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": &schema.Schema{
							Description: `Roles can be empty or have many values for a node. `,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"services": &schema.Schema{
							Description: `Services can be empty or have many values for a node. `,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceNodeDeploymentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vHostname, okHostname := d.GetOk("hostname")

	method1 := []bool{okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostname}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeploymentNodes")
		queryParams1 := isegosdk.GetDeploymentNodesQueryParams{}

		if okFilter {
			queryParams1.Filter = interfaceToSliceString(vFilter)
		}
		if okFilterType {
			queryParams1.FilterType = vFilterType.(string)
		}

		response1, restyResp1, err := client.NodeDeployment.GetDeploymentNodes(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeploymentNodes", err,
				"Failure at GetDeploymentNodes, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNodeDeploymentGetDeploymentNodesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeploymentNodes response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNodeDetails")
		vvHostname := vHostname.(string)

		response2, restyResp2, err := client.NodeDeployment.GetNodeDetails(vvHostname)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNodeDetails", err,
				"Failure at GetNodeDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNodeDeploymentGetNodeDetailsItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNodeDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNodeDeploymentGetDeploymentNodesItems(items *[]isegosdk.ResponseNodeDeploymentGetDeploymentNodesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["fqdn"] = item.Fqdn
		respItem["hostname"] = item.Hostname
		respItem["ip_address"] = item.IPAddress
		respItem["node_status"] = item.NodeStatus
		respItem["roles"] = item.Roles
		respItem["services"] = item.Services
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNodeDeploymentGetNodeDetailsItem(item *isegosdk.ResponseNodeDeploymentGetNodeDetailsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["fqdn"] = item.Fqdn
	respItem["hostname"] = item.Hostname
	respItem["ip_address"] = item.IPAddress
	respItem["node_status"] = item.NodeStatus
	respItem["roles"] = item.Roles
	respItem["services"] = item.Services
	return []map[string]interface{}{
		respItem,
	}
}
