package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePxGridDirectDictionaryInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on pxGrid Direct.

- pxGrid Direct Get a map of references to pxgrid-direct dictionaries
`,

		ReadContext: dataSourcePxGridDirectDictionaryInfoRead,
		Schema: map[string]*schema.Schema{
			"x_request_id": &schema.Schema{
				Description: `X-Request-ID header parameter. request Id, will return in the response headers, and appear in logs`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePxGridDirectDictionaryInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetpxgridDirectDictionaryReferences")

		response1, restyResp1, err := client.PxGridDirect.GetpxgridDirectDictionaryReferences()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetpxgridDirectDictionaryReferences", err,
				"Failure at GetpxgridDirectDictionaryReferences, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPxGridDirectGetpxgridDirectDictionaryReferencesItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetpxgridDirectDictionaryReferences response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPxGridDirectGetpxgridDirectDictionaryReferencesItem(item *isegosdk.ResponsePxGridDirectGetpxgridDirectDictionaryReferences) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = flattenPxGridDirectGetpxgridDirectDictionaryReferencesItemResponse(item.Response)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenPxGridDirectGetpxgridDirectDictionaryReferencesItemResponse(item *isegosdk.ResponsePxGridDirectGetpxgridDirectDictionaryReferencesResponse) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
