package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSystemCertificateExportInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Certificates.

- Export System Certificate.

`,

		ReadContext: dataSourceSystemCertificateExportInfoRead,
		Schema: map[string]*schema.Schema{
			"dirpath": &schema.Schema{
				Description: `Directory absolute path in which to save the file.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"export": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceSystemCertificateExportInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ExportSystemCertificate")
		request1 := expandRequestSystemCertificateExportInfoExportSystemCertificate(ctx, "", d)

		response1, _, err := client.Certificates.ExportSystemCertificate(request1)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing ExportSystemCertificate", err))
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

func expandRequestSystemCertificateExportInfoExportSystemCertificate(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesExportSystemCertificate {
	request := isegosdk.RequestCertificatesExportSystemCertificate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".export")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".export")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".export")))) {
		request.Export = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	return &request
}
