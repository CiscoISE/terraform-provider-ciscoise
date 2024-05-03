package ciscoise

import (
	"context"
	"log"

	"reflect"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourcePxGridDirectSync() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on pxGrid Direct.

- This syncNow is used on demand on a URLFetch Type connector only
 Following parameters are present in the POST request body




PARAMETER

DESCRIPTION

EXAMPLE





SyncType
*required

its for FULL or INCREMENTAL

"SyncType": "FULL or INCREMENTAL"



connectorName
*required

Name of the Connector for only URLFetcher type

name of Connector



description

Decription of the Connector

"length": "256 character"




NOTE:
For
Use syncNowStatus api to get the status of the connector


`,

		CreateContext: resourcePxGridDirectSyncCreate,
		ReadContext:   resourcePxGridDirectSyncRead,
		DeleteContext: resourcePxGridDirectSyncDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sync_type": &schema.Schema{
							Description: `connector Type list`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"connector_name": &schema.Schema{
							Description: `connectorName`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `description`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourcePxGridDirectSyncCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestPxGridDirectSyncSyncNowConnector(ctx, "parameters.0", d)

	response1, err := client.PxGridDirect.SyncNowConnector(request1)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing SyncNowConnector", err,
			"Failure at SyncNowConnector, unexpected response", ""))
		return diags
	}

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting SyncNowConnector response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags
}

func expandRequestPxGridDirectSyncSyncNowConnector(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectSyncNowConnector {
	request := isegosdk.RequestPxGridDirectSyncNowConnector{}
	request.Connector = expandRequestPxGridDirectSyncSyncNowConnectorConnector(ctx, key, d)
	return &request
}

func expandRequestPxGridDirectSyncSyncNowConnectorConnector(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestPxGridDirectSyncNowConnectorConnector {
	request := isegosdk.RequestPxGridDirectSyncNowConnectorConnector{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_type")))) {
		request.SyncType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_name")))) {
		request.ConnectorName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	return &request
}

func resourcePxGridDirectSyncRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePxGridDirectSyncUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePxGridDirectSyncRead(ctx, d, m)
}

func resourcePxGridDirectSyncDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
