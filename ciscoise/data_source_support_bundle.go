package ciscoise

import (
	"context"

	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSupportBundle() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSupportBundleRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"support_bundle_include_options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"from_date": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"include_config_db": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"include_core_files": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"include_debug_logs": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"include_local_logs": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"include_system_logs": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"mnt_logs": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"policy_xml": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"to_date": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateSupportBundle", err,
				"Failure at CreateSupportBundle, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

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
	if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(d.Get("name"))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(d.Get("description"))) && (ok || !reflect.DeepEqual(v, d.Get("description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("host_name"); !isEmptyValue(reflect.ValueOf(d.Get("host_name"))) && (ok || !reflect.DeepEqual(v, d.Get("host_name"))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("support_bundle_include_options"); !isEmptyValue(reflect.ValueOf(d.Get("support_bundle_include_options"))) && (ok || !reflect.DeepEqual(v, d.Get("support_bundle_include_options"))) {
		request.SupportBundleIncludeOptions = expandRequestSupportBundleCreateSupportBundleSupportBundleSupportBundleIncludeOptions(ctx, key+".support_bundle_include_options.0", d)
	}
	return &request
}

func expandRequestSupportBundleCreateSupportBundleSupportBundleSupportBundleIncludeOptions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSupportBundleTriggerConfigurationCreateSupportBundleSupportBundleSupportBundleIncludeOptions {
	request := isegosdk.RequestSupportBundleTriggerConfigurationCreateSupportBundleSupportBundleSupportBundleIncludeOptions{}
	if v, ok := d.GetOkExists("include_config_db"); !isEmptyValue(reflect.ValueOf(d.Get("include_config_db"))) && (ok || !reflect.DeepEqual(v, d.Get("include_config_db"))) {
		request.IncludeConfigDB = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("include_debug_logs"); !isEmptyValue(reflect.ValueOf(d.Get("include_debug_logs"))) && (ok || !reflect.DeepEqual(v, d.Get("include_debug_logs"))) {
		request.IncludeDebugLogs = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("include_local_logs"); !isEmptyValue(reflect.ValueOf(d.Get("include_local_logs"))) && (ok || !reflect.DeepEqual(v, d.Get("include_local_logs"))) {
		request.IncludeLocalLogs = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("include_core_files"); !isEmptyValue(reflect.ValueOf(d.Get("include_core_files"))) && (ok || !reflect.DeepEqual(v, d.Get("include_core_files"))) {
		request.IncludeCoreFiles = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("mnt_logs"); !isEmptyValue(reflect.ValueOf(d.Get("mnt_logs"))) && (ok || !reflect.DeepEqual(v, d.Get("mnt_logs"))) {
		request.MntLogs = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("include_system_logs"); !isEmptyValue(reflect.ValueOf(d.Get("include_system_logs"))) && (ok || !reflect.DeepEqual(v, d.Get("include_system_logs"))) {
		request.IncludeSystemLogs = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("policy_xml"); !isEmptyValue(reflect.ValueOf(d.Get("policy_xml"))) && (ok || !reflect.DeepEqual(v, d.Get("policy_xml"))) {
		request.PolicyXml = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists("from_date"); !isEmptyValue(reflect.ValueOf(d.Get("from_date"))) && (ok || !reflect.DeepEqual(v, d.Get("from_date"))) {
		request.FromDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists("to_date"); !isEmptyValue(reflect.ValueOf(d.Get("to_date"))) && (ok || !reflect.DeepEqual(v, d.Get("to_date"))) {
		request.ToDate = interfaceToString(v)
	}
	return &request
}
