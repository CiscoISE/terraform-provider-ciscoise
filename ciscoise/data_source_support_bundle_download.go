package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSupportBundleDownload() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on SupportBundleDownload.

- This data source action allows the client to upload a support bundle.
`,

		ReadContext: dataSourceSupportBundleDownloadRead,
		Schema: map[string]*schema.Schema{
			"dirpath": &schema.Schema{
				Description: `Directory absolute path in which to save the file.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"file_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceSupportBundleDownloadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: DownloadSupportBundle")
		request1 := expandRequestSupportBundleDownloadDownloadSupportBundle(ctx, "", d)

		response1, _, err := client.SupportBundleDownload.DownloadSupportBundle(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing DownloadSupportBundle", err))
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

func expandRequestSupportBundleDownloadDownloadSupportBundle(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSupportBundleDownloadDownloadSupportBundle {
	request := isegosdk.RequestSupportBundleDownloadDownloadSupportBundle{}
	request.ErsSupportBundleDownload = expandRequestSupportBundleDownloadDownloadSupportBundleErsSupportBundleDownload(ctx, key, d)
	return &request
}

func expandRequestSupportBundleDownloadDownloadSupportBundleErsSupportBundleDownload(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSupportBundleDownloadDownloadSupportBundleErsSupportBundleDownload {
	request := isegosdk.RequestSupportBundleDownloadDownloadSupportBundleErsSupportBundleDownload{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".file_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".file_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".file_name")))) {
		request.FileName = interfaceToString(v)
	}
	return &request
}
