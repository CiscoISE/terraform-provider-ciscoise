package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePxgridAccountCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Consumer.
- 🚧 AccountCreate
`,

		CreateContext: resourcePxgridAccountCreateCreate,
		ReadContext:   resourcePxgridAccountCreateRead,
		DeleteContext: resourcePxgridAccountCreateDelete,

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
						"node_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourcePxgridAccountCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning CreateAccount create")
	log.Printf("[DEBUG] Missing CreateAccount create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestPxgridAccountCreateCreateAccount(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, err := client.Consumer.CreateAccount(request1)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateAccount", err,
			"Failure at CreateAccount, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CreateAccount response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourcePxgridAccountCreateRead(ctx, d, m)
}

func resourcePxgridAccountCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourcePxgridAccountCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PxgridAccountCreate delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing PxgridAccountCreate delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
func expandRequestPxgridAccountCreateCreateAccount(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestClearThreatsAndVulnerabilitiesCreateAccount {
	request := isegosdk.RequestClearThreatsAndVulnerabilitiesCreateAccount{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".node_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".node_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".node_name")))) {
		request.NodeName = interfaceToString(v)
	}
	return &request
}
