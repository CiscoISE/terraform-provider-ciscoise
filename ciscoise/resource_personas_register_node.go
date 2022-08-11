package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePersonasRegisterNode() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Access - Authentication Rules.
- Network Access Reset HitCount for Authentication Rules
`,

		CreateContext: resourcePersonasRegisterNodeCreate,
		ReadContext:   resourcePersonasRegisterNodeRead,
		DeleteContext: resourcePersonasRegisterNodeDelete,

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
						"primary_ip": &schema.Schema{
							Description: `Primary Node Ip`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"primary_username": &schema.Schema{
							Description: `Primary username`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"primary_password": &schema.Schema{
							Description: `Primary password`,
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
						},
						"fqdn": &schema.Schema{
							Description: `fqdn`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"roles": &schema.Schema{
							Description: `roles`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"services": &schema.Schema{
							Description: `services`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourcePersonasRegisterNodeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PersonasRegisterNode")
	var diags diag.Diagnostics
	node := expandRequestPersonasRegisterNode(ctx, "parameters.0", d)
	primaryNode := expandRequestPersonasRegisterNodePrimary(ctx, "parameters.0", d)

	primaryAppServerIsRunning, err := primaryNode.AppServerIsRunning()
	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AppServerIsRunning function", err,
			"Failure at AppServerIsRunning, unexpected response", ""))
		return diags
	}
	if !primaryAppServerIsRunning {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AppServerIsRunning function", fmt.Errorf("Application server is not running."),
			"Failure at AppServerIsRunning", ""))
		return diags
	}

	err = node.RegisterToPrimary(primaryNode)

	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RegisterToPrimary function", err,
			"Failure at RegisterToPrimary, unexpected response", ""))
		return diags
	}

	if err := d.Set("item", fmt.Sprintf("Node %s updated successfully", node.Fqdn)); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RegisterToPrimary response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourcePersonasRegisterNodeRead(ctx, d, m)
}

func resourcePersonasRegisterNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourcePersonasRegisterNodeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PersonasRegisterNode delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing PersonasRegisterNode delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestPersonasRegisterNode(ctx context.Context, key string, d *schema.ResourceData) Node {
	request := Node{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".roles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".roles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".roles")))) {
		request.Roles = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services")))) {
		request.Services = interfaceToSliceString(v)
	}
	return request
}
func expandRequestPersonasRegisterNodePrimary(ctx context.Context, key string, d *schema.ResourceData) Node {
	request := Node{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_ip")))) {
		request.Ip = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_username")))) {
		request.UserName = interfaceToString(v)
	}

	return request
}
