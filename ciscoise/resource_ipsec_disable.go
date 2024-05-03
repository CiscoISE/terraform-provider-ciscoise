package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceIPsecDisable() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Native IPsec.

- Disables an enabled IPsec node connection.
`,

		CreateContext: resourceIPsecDisableCreate,
		ReadContext:   resourceIPsecDisableRead,
		DeleteContext: resourceIPsecDisableDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_name": &schema.Schema{
							Description: `hostName path parameter. Hostname of the deployed node.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"nad_ip": &schema.Schema{
							Description: `nadIp path parameter. IP address of the NAD.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func resourceIPsecDisableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))

	vHostName := resourceItem["host_name"]

	vNadIP := resourceItem["nad_ip"]

	vvHostName := vHostName.(string)
	vvNadIP := vNadIP.(string)

	response1, restyResp1, err := client.NativeIPsec.DisableIPsecConnection(vvHostName, vvNadIP)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenNativeIPsecDisableIPsecConnectionItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DisableIPsecConnection response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}

func flattenNativeIPsecDisableIPsecConnectionItem(item *isegosdk.ResponseNativeIPsecDisableIPsecConnection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}

func resourceIPsecDisableRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIPsecDisableUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceIPsecDisableRead(ctx, d, m)
}

func resourceIPsecDisableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
