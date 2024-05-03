package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDuoIDentitysyncSyncInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Duo-IdentitySync.

- Initiate the sync between the Active Directory and the corresponding Mfa provider associated with this Identitysync
config.
`,

		ReadContext: dataSourceDuoIDentitysyncSyncInfoRead,
		Schema: map[string]*schema.Schema{
			"sync_name": &schema.Schema{
				Description: `syncName path parameter. Name of the Identitysync configuration used to initiate sync.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDuoIDentitysyncSyncInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vSyncName := d.Get("sync_name")

	log.Printf("[DEBUG] Selected method: Sync")
	vvSyncName := vSyncName.(string)

	response1, err := client.DuoIDentitySync.Sync(vvSyncName)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing Sync", err,
			"Failure at Sync, unexpected response", ""))
		return diags
	}
	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting Sync response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags
}
