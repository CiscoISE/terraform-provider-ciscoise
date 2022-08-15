package ciscoise

import (
	"context"
	"reflect"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePersonasUpdateRolesServices() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Access - Authentication Rules.
- Network Access Reset HitCount for Authentication Rules
`,

		CreateContext: resourcePersonasUpdateRolesServicesCreate,
		ReadContext:   resourcePersonasUpdateRolesServicesRead,
		DeleteContext: resourcePersonasUpdateRolesServicesDelete,

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
						},
						"roles": &schema.Schema{
							Description: `roles`,
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"services": &schema.Schema{
							Description: `services`,
							Type:        schema.TypeList,
							Required:    true,
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

func resourcePersonasUpdateRolesServicesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PersonasUpdateRolesServices")
	var diags diag.Diagnostics
	node := expandRequestPersonasUpdateRolesServices(ctx, "parameters.0", d)
	err := node.UpdateRolesServices()
	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing UpdateRolesServices function", err,
			"Failure at UpdateRolesServices, unexpected response", ""))
		return diags
	}
	if err := d.Set("item", "Node updated successfully"); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateRolesServices response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourcePersonasUpdateRolesServicesRead(ctx, d, m)
}

func resourcePersonasUpdateRolesServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourcePersonasUpdateRolesServicesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PersonasUpdateRolesServices delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing PersonasUpdateRolesServices delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestPersonasUpdateRolesServices(ctx context.Context, key string, d *schema.ResourceData) Node {
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".roles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".roles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".roles")))) {
		request.Roles = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services")))) {
		request.Services = interfaceToSliceString(v)
	}

	return request
}
