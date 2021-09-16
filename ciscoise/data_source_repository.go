package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Repository.

This will get the full list of repository definitions on the system.

Get a specific repository identified by the name passed in the URL.
`,

		ReadContext: dataSourceRepositoryRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: `name path parameter. Unique name for a repository`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enable_pki": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password": &schema.Schema{
							Description: `Password can contain alphanumeric and/or special characters.`,
							Type:        schema.TypeString,
							Sensitive:   true,
							Computed:    true,
						},
						"path": &schema.Schema{
							Description: `Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name": &schema.Schema{
							Description: `Username can contain alphanumeric characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enable_pki": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password": &schema.Schema{
							Description: `Password can contain alphanumeric and/or special characters.`,
							Type:        schema.TypeString,
							Sensitive:   true,
							Computed:    true,
						},
						"path": &schema.Schema{
							Description: `Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name": &schema.Schema{
							Description: `Username can contain alphanumeric characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetRepositories")

		response1, _, err := client.Repository.GetRepositories()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRepositories", err,
				"Failure at GetRepositories, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenRepositoryGetRepositoriesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRepositories response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetRepository")
		vvName := vName.(string)

		response2, _, err := client.Repository.GetRepository(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRepository", err,
				"Failure at GetRepository, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenRepositoryGetRepositoryItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRepository response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenRepositoryGetRepositoriesItems(items *[]isegosdk.ResponseRepositoryGetRepositoriesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["protocol"] = item.Protocol
		respItem["path"] = item.Path
		respItem["password"] = item.Password
		respItem["server_name"] = item.ServerName
		respItem["user_name"] = item.UserName
		respItem["enable_pki"] = item.EnablePki
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenRepositoryGetRepositoryItem(item *isegosdk.ResponseRepositoryGetRepositoryResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["protocol"] = item.Protocol
	respItem["path"] = item.Path
	respItem["password"] = item.Password
	respItem["server_name"] = item.ServerName
	respItem["user_name"] = item.UserName
	respItem["enable_pki"] = item.EnablePki
	return []map[string]interface{}{
		respItem,
	}
}
