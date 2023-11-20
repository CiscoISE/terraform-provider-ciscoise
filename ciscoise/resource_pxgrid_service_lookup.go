package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePxgridServiceLookup() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Consumer.
- 🚧 ServiceLookup
`,

		CreateContext: resourcePxgridServiceLookupCreate,
		ReadContext:   resourcePxgridServiceLookupRead,
		DeleteContext: resourcePxgridServiceLookupDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
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
						"name": &schema.Schema{
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

func resourcePxgridServiceLookupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning LookupService create")
	log.Printf("[DEBUG] Missing LookupService create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	request1 := expandRequestPxgridServiceLookupLookupService(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, err := client.Consumer.LookupService(request1)
	if err != nil || response1 == nil {
		if response1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", response1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing LookupService", err, response1.String(),
				"Failure at LookupService, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing LookupService", err,
			"Failure at LookupService, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %s", response1.String())

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting LookupService response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourcePxgridServiceLookupRead(ctx, d, m)
}

func resourcePxgridServiceLookupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePxgridServiceLookupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PxgridServiceLookup delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing PxgridServiceLookup delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestPxgridServiceLookupLookupService(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestConfigurationLookupService {
	request := isegosdk.RequestConfigurationLookupService{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	return &request
}
