package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCsrExport() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCsrExportRead,
		Schema: map[string]*schema.Schema{
			"dirpath": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceCsrExportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vHostname := d.Get("hostname")
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ExportCsr")
		vvHostname := vHostname.(string)
		vvID := vID.(string)

		response1, _, err := client.Certificates.ExportCsr(vvHostname, vvID)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing ExportCsr", err))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response")

		vvDirpath := d.Get("dirpath").(string)
		err = response1.SaveDownload(vvDirpath)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when downloading file", err))
			return diags
		}
		log.Printf("[DEBUG] Downloaded file %s", vvDirpath)

	}
	return diags
}
