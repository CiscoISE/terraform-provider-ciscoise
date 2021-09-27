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
func dataSourceSupportBundle() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on SupportBundleTriggerConfiguration.

- This data source action allows the client to create a support bundle trigger configuration.
`,

		ReadContext: dataSourceSupportBundleRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Description: `This parameter is hostName only, xxxx of xxxx.yyy.zz`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Description: `Resource Name`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"support_bundle_include_options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"from_date": &schema.Schema{
							Description: `Date from where support bundle should include the logs`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"include_config_db": &schema.Schema{
							Description:  `Set to include Config DB in Support Bundle`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"include_core_files": &schema.Schema{
							Description:  `Set to include Core files in Support Bundle`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"include_debug_logs": &schema.Schema{
							Description:  `Set to include Debug logs in Support Bundle`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"include_local_logs": &schema.Schema{
							Description:  `Set to include Local logs in Support Bundle`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"include_system_logs": &schema.Schema{
							Description:  `Set to include System logs in Support Bundle`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"mnt_logs": &schema.Schema{
							Description:  `Set to include Monitoring and troublshooting logs in Support Bundle`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"policy_xml": &schema.Schema{
							Description:  `Set to include Policy XML in Support Bundle`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"to_date": &schema.Schema{
							Description: `Date upto where support bundle should include the logs`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSupportBundleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateSupportBundle")
		request1 := expandRequestSupportBundleCreateSupportBundle(ctx, "", d)

		response1, err := client.SupportBundleTriggerConfiguration.CreateSupportBundle(request1)

		if err != nil || response1 == nil {
			if request1 != nil {
				log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateSupportBundle", err,
				"Failure at CreateSupportBundle, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateSupportBundle response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSupportBundleCreateSupportBundle(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSupportBundleTriggerConfigurationCreateSupportBundle {
	request := isegosdk.RequestSupportBundleTriggerConfigurationCreateSupportBundle{}
	request.SupportBundle = expandRequestSupportBundleCreateSupportBundleSupportBundle(ctx, key, d)
	return &request
}

func expandRequestSupportBundleCreateSupportBundleSupportBundle(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSupportBundleTriggerConfigurationCreateSupportBundleSupportBundle {
	request := isegosdk.RequestSupportBundleTriggerConfigurationCreateSupportBundleSupportBundle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".support_bundle_include_options")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".support_bundle_include_options")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".support_bundle_include_options")))) {
		request.SupportBundleIncludeOptions = expandRequestSupportBundleCreateSupportBundleSupportBundleSupportBundleIncludeOptions(ctx, key+".support_bundle_include_options.0", d)
	}
	return &request
}

func expandRequestSupportBundleCreateSupportBundleSupportBundleSupportBundleIncludeOptions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSupportBundleTriggerConfigurationCreateSupportBundleSupportBundleSupportBundleIncludeOptions {
	request := isegosdk.RequestSupportBundleTriggerConfigurationCreateSupportBundleSupportBundleSupportBundleIncludeOptions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_config_db")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_config_db")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_config_db")))) {
		request.IncludeConfigDB = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_debug_logs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_debug_logs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_debug_logs")))) {
		request.IncludeDebugLogs = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_local_logs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_local_logs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_local_logs")))) {
		request.IncludeLocalLogs = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_core_files")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_core_files")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_core_files")))) {
		request.IncludeCoreFiles = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mnt_logs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mnt_logs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mnt_logs")))) {
		request.MntLogs = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_system_logs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_system_logs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_system_logs")))) {
		request.IncludeSystemLogs = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".policy_xml")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".policy_xml")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".policy_xml")))) {
		request.PolicyXml = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_date")))) {
		request.FromDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_date")))) {
		request.ToDate = interfaceToString(v)
	}
	return &request
}
