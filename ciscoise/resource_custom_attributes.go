package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCustomAttributes() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on customattributes.

- Create Custom Attribute

- Delete custom attribute by name
`,

		CreateContext: resourceCustomAttributesCreate,
		ReadContext:   resourceCustomAttributesRead,
		UpdateContext: resourceCustomAttributesUpdate,
		DeleteContext: resourceCustomAttributesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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

						"attribute_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_name": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"attribute_type": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"name": &schema.Schema{
							Description:      `name path parameter. The name of the custom attribute`,
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: diffSupressOptional(),
						},
					},
				},
			},
		},
	}
}

func resourceCustomAttributesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := m.(ClientConfig).EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestCustomAttributesCreateCustomAttribute(ctx, "parameters.0", d)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okName && vvName != "" {
		vName = resourceItem["attribute_name"]
		vvName = interfaceToString(vName)
	}
	if isEnableAutoImport {
		if okName && vvName != "" {
			getResponse2, _, err := client.CustomAttributes.Get(vvName)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceCustomAttributesRead(ctx, d, m)
			}
		}
	}

	resourceItem = *getResourceItem(d.Get("parameters"))
	request1 = expandRequestCustomAttributesCreateCustomAttribute(ctx, "parameters.0", d)
	resp1, err := client.CustomAttributes.CreateCustomAttribute(request1)
	if err != nil || resp1 == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateCustomAttribute", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceCustomAttributesRead(ctx, d, m)
}

func resourceCustomAttributesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvName := resourceMap["name"]

	response2, restyResp2, err := client.CustomAttributes.Get(vvName)

	if err != nil || response2 == nil {
		if restyResp2 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

	vItem2 := flattenCustomAttributesGetItem(response2)
	if err := d.Set("item", vItem2); err != nil {
		diags = append(diags, diagError(
			"Failure when setting Get response",
			err))
		return diags
	}
	if err := d.Set("parameters", vItem2); err != nil {
		diags = append(diags, diagError(
			"Failure when setting Get response",
			err))
		return diags
	}
	return diags

}

func resourceCustomAttributesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCustomAttributesRead(ctx, d, m)
}

func resourceCustomAttributesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvName := resourceMap["name"]

	getResp, _, err := client.CustomAttributes.Get(vvName)
	if err != nil || getResp == nil {
		// Assume that element it is already gone
		return diags
	}
	response1, err := client.CustomAttributes.Delete(vvName)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing Delete", err,
			"Failure at Delete, unexpected response", ""))
		return diags
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestCustomAttributesCreateCustomAttribute(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCustomAttributesCreateCustomAttribute {
	request := isegosdk.RequestCustomAttributesCreateCustomAttribute{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_name")))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_type")))) {
		request.AttributeType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
