package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePersonasCheckStandalone() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Access - Authentication Rules.
- Network Access Reset HitCount for Authentication Rules
`,

		CreateContext: resourcePersonasCheckStandaloneCreate,
		ReadContext:   resourcePersonasCheckStandaloneRead,
		DeleteContext: resourcePersonasCheckStandaloneDelete,

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
						"ip": &schema.Schema{
							Description: `Node Ip`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"hostname": &schema.Schema{
							Description: `Node hostname`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"username": &schema.Schema{
							Description: `username`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"password": &schema.Schema{
							Description: `password`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						}},
				},
			},
		},
	}
}

func resourcePersonasCheckStandaloneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PersonasCheckStandalone")
	var diags diag.Diagnostics
	node := expandRequestPersonasCheckStandalone(ctx, "parameters.0", d)
	isStandAlone, err := node.IsStandAlone()
	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing IsStandAlone function", err,
			"Failure at IsStandAlone, unexpected response", ""))
		return diags
	}
	serverIsRunning, err := node.AppServerIsRunning()
	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AppServerIsRunning function", err,
			"Failure at AppServerIsRunning, unexpected response", ""))
		return diags
	}
	if !(isStandAlone && serverIsRunning) {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AppServerIsRunning function", fmt.Errorf(fmt.Sprintf("Node %s is not in STANDALONE state or application server is not running.", node.HostName)),
			"Failure at AppServerIsRunning", ""))
		return diags
	}
	if err := d.Set("item", fmt.Sprintf("Node %s is in STANDALONE mode", node.HostName)); err != nil {
		diags = append(diags, diagError(
			"Failure when setting STANDALONE response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourcePersonasCheckStandaloneRead(ctx, d, m)
}

func resourcePersonasCheckStandaloneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourcePersonasCheckStandaloneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PersonasCheckStandalone delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing PersonasCheckStandalone delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestPersonasCheckStandalone(ctx context.Context, key string, d *schema.ResourceData) Node {
	request := Node{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip")))) {
		request.Ip = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.HostName = interfaceToString(v)
	}

	return request
}
