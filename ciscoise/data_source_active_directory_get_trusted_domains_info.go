package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceActiveDirectoryGetTrustedDomainsInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceActiveDirectoryGetTrustedDomainsInfoRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"domains": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"forest": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"unusable_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceActiveDirectoryGetTrustedDomainsInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTrustedDomains")
		vvID := vID.(string)

		response1, _, err := client.ActiveDirectory.GetTrustedDomains(vvID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTrustedDomains", err,
				"Failure at GetTrustedDomains, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenActiveDirectoryGetTrustedDomainsItem(response1.ERSActiveDirectoryDomains)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTrustedDomains response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenActiveDirectoryGetTrustedDomainsItem(item *isegosdk.ResponseActiveDirectoryGetTrustedDomainsERSActiveDirectoryDomains) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["domains"] = flattenActiveDirectoryGetTrustedDomainsItemDomains(item.Domains)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenActiveDirectoryGetTrustedDomainsItemDomains(items *[]isegosdk.ResponseActiveDirectoryGetTrustedDomainsERSActiveDirectoryDomainsDomains) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dns_name"] = item.DNSName
		respItem["forest"] = item.Forest
		respItem["unusable_reason"] = item.UnusableReason
	}
	return respItems

}
