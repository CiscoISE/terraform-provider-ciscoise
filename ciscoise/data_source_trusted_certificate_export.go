package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTrustedCertificateExport() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Certificates.

- The response of this API carries a trusted certificate file mapped to the requested ID
`,

		ReadContext: dataSourceTrustedCertificateExportRead,
		Schema: map[string]*schema.Schema{
			"dirpath": &schema.Schema{
				Description: `Directory absolute path in which to save the file.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. ID of the Trusted Certificate to be exported.`,
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceTrustedCertificateExportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ExportTrustedCert")
		vvID := vID.(string)

		response1, _, err := client.Certificates.ExportTrustedCert(vvID)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing ExportTrustedCert", err))
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
