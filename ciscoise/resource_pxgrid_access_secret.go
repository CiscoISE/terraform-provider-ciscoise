package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePxgridAccessSecret() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Consumer.
- 🚧 AccessSecret
`,

		CreateContext: resourcePxgridAccessSecretCreate,
		ReadContext:   resourcePxgridAccessSecretRead,
		DeleteContext: resourcePxgridAccessSecretDelete,

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
						"peer_node_name": &schema.Schema{
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

func resourcePxgridAccessSecretCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AccessSecret create")
	log.Printf("[DEBUG] Missing AccessSecret create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	request1 := expandRequestPxgridAccessSecretAccessSecret(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, err := client.Consumer.AccessSecret(request1)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AccessSecret", err,
			"Failure at AccessSecret, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AccessSecret response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourcePxgridAccessSecretRead(ctx, d, m)
}

func resourcePxgridAccessSecretRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourcePxgridAccessSecretDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PxgridAccessSecret delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing PxgridAccessSecret delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestPxgridAccessSecretAccessSecret(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestClearThreatsAndVulnerabilitiesAccessSecret {
	request := isegosdk.RequestClearThreatsAndVulnerabilitiesAccessSecret{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".peer_node_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".peer_node_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".peer_node_name")))) {
		request.PeerNodeName = interfaceToString(v)
	}
	return &request
}
