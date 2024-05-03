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
func resourceCustomAttributesRename() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on customattributes.

- rename custom attribute
`,

		CreateContext: resourceCustomAttributesRenameCreate,
		ReadContext:   resourceCustomAttributesRenameRead,
		DeleteContext: resourceCustomAttributesRenameDelete,
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
						"current_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"new_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceCustomAttributesRenameCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestCustomAttributesRenameRename(ctx, "parameters.0", d)

	response1, err := client.CustomAttributes.Rename(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing Rename", err,
			"Failure at Rename, unexpected response", ""))
		return diags
	}

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting Rename response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags
}

func expandRequestCustomAttributesRenameRename(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCustomAttributesRename {
	request := isegosdk.RequestCustomAttributesRename{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".current_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".current_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".current_name")))) {
		request.CurrentName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_name")))) {
		request.NewName = interfaceToString(v)
	}
	return &request
}

func resourceCustomAttributesRenameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceCustomAttributesRenameUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCustomAttributesRenameRead(ctx, d, m)
}

func resourceCustomAttributesRenameDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
