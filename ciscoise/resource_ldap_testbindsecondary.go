package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceLdapTestbindsecondary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on ldap.

- test-bind-secondary
`,

		CreateContext: resourceLdapTestbindsecondaryCreate,
		ReadContext:   resourceLdapTestbindsecondaryRead,
		DeleteContext: resourceLdapTestbindsecondaryDelete,
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

						"messages": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"code": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"operation": &schema.Schema{
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
						"id": &schema.Schema{
							Description: `id path parameter.`,
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

func resourceLdapTestbindsecondaryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]
	vvID := vID.(string)

	response1, restyResp1, err := client.Ldap.PutLdapidTestbindsecondary(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenLdapPutLdapidTestbindsecondaryItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting PutLdapidTestbindsecondary response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags
}

func flattenLdapPutLdapidTestbindsecondaryItem(item *isegosdk.ResponseLdapPutLdapidTestbindsecondary) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["operation"] = item.Operation
	respItem["messages"] = flattenLdapPutLdapidTestbindsecondaryItemMessages(item.Messages)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenLdapPutLdapidTestbindsecondaryItemMessages(items *[]isegosdk.ResponseLdapPutLdapidTestbindsecondaryMessages) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["title"] = item.Title
		respItem["type"] = item.Type
		respItem["code"] = item.Code
		respItems = append(respItems, respItem)
	}
	return respItems
}

func resourceLdapTestbindsecondaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceLdapTestbindsecondaryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceLdapTestbindsecondaryRead(ctx, d, m)
}

func resourceLdapTestbindsecondaryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
