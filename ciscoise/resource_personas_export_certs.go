package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePersonasExportCerts() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Access - Authentication Rules.
- Network Access Reset HitCount for Authentication Rules
`,

		CreateContext: resourcePersonasExportCertsCreate,
		ReadContext:   resourcePersonasExportCertsRead,
		DeleteContext: resourcePersonasExportCertsDelete,

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
						"hostname": &schema.Schema{
							Description: `Node hostname`,
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
						"username": &schema.Schema{
							Description: `username`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"name": &schema.Schema{
							Description: `name`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"ip": &schema.Schema{
							Description: `ip`,
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

func resourcePersonasExportCertsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PersonasExportCerts")
	var diags diag.Diagnostics
	node := expandRequestPersonasExportCerts(ctx, "parameters.0", d)
	primaryNode := expandRequestPersonasExportCertsPrimary(ctx, "parameters.0", d)

	err := node.ImportCertificateIntoPrimary(primaryNode)

	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ImportCertificateIntoPrimary function", err,
			"Failure at ImportCertificateIntoPrimary, unexpected response", ""))
		return diags
	}

	if err := d.Set("item", fmt.Sprintf("The certificate for %s was exported successfully to the primary node", node.HostName)); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ImportCertificateIntoPrimary response",
			err))
		return diags
	}

	_ = d.Set("last_updated", getUnixTimeString())

	d.SetId(getUnixTimeString())
	return resourcePersonasExportCertsRead(ctx, d, m)
}

func resourcePersonasExportCertsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourcePersonasExportCertsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PersonasExportCerts delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing PersonasExportCerts delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestPersonasExportCerts(ctx context.Context, key string, d *schema.ResourceData) Node {
	request := Node{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip")))) {
		request.Ip = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.HostName = interfaceToString(v)
	}
	return request
}
func expandRequestPersonasExportCertsPrimary(ctx context.Context, key string, d *schema.ResourceData) Node {
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
