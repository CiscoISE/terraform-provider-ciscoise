package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePxGridNode() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and delete operations on PxGridNode.

- This resource action deletes a pxGrid node by name.
`,

		CreateContext: resourcePxGridNodeCreate,
		ReadContext:   resourcePxGridNodeRead,
		UpdateContext: resourcePxGridNodeUpdate,
		DeleteContext: resourcePxGridNodeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auth_method": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"groups": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourcePxGridNodeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PxGridNode create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okName && vvName != "" {
		getResponse2, _, err := client.PxGridNode.GetPxGridNodeByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourcePxGridNodeRead(ctx, d, m)
		}
	}
	diags = append(diags, diagErrorWithAlt(
		"Failure when executing GetPxGridNodeByName", nil,
		"Failure at GetPxGridNodeByName, unexpected response", ""))
	return diags
}

func resourcePxGridNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PxGridNode read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vName, okName := resourceMap["name"]

	if okName && vName != "" {
		vvName := vName
		log.Printf("[DEBUG] Selected method: GetPxGridNodeByName")
		response1, restyResp1, err := client.PxGridNode.GetPxGridNodeByName(vvName)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPxGridNodeGetPxGridNodeByNameItemName(response1.PxgridNode)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPxGridNodes search response",
				err))
			return diags
		}
		if err := d.Set("parameters", remove_parameters(vItem1, "link")); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPxGridNodes response to parameters",
				err))
			return diags
		}
	}
	return diags
}

func resourcePxGridNodeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PxGridNode update for id=[%s]", d.Id())
	log.Printf("[DEBUG] Missing PxGridNode update on Cisco ISE. It will only be update it on Terraform")
	// _ = d.Set("last_updated", getUnixTimeString())
	return resourcePxGridNodeRead(ctx, d, m)
}

func resourcePxGridNodeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning PxGridNode delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]

	if okName && vName != "" {
		response1, _, err := client.PxGridNode.GetPxGridNodeByName(vName)
		if err != nil || response1 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.PxGridNode.DeletePxGridNodeByName(vName)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletePxGridNodeByName", err, restyResp1.String(),
				"Failure at DeletePxGridNodeByName, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletePxGridNodeByName", err,
			"Failure at DeletePxGridNodeByName, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
