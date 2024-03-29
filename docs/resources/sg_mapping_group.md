---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_sg_mapping_group Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on IPToSGTMappingGroup.
  This resource allows the client to update an IP to SGT mapping group by ID.This resource deletes an IP to SGT mapping group.This resource creates an IP to SGT mapping group.
---

# ciscoise_sg_mapping_group (Resource)

It manages create, read, update and delete operations on IPToSGTMappingGroup.

- This resource allows the client to update an IP to SGT mapping group by ID.

- This resource deletes an IP to SGT mapping group.

- This resource creates an IP to SGT mapping group.

## Example Usage

```terraform
resource "ciscoise_sg_mapping_group" "example" {
  provider = ciscoise
  parameters {

    deploy_to   = "string"
    deploy_type = "string"
    id          = "string"
    name        = "string"
    sgt         = "string"
  }
}

output "ciscoise_sg_mapping_group_example" {
  value = ciscoise_sg_mapping_group.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String) Unix timestamp records the last time that the resource was updated.

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `id` (String) id path parameter.

Optional:

- `deploy_to` (String) Mandatory unless mappingGroup is set or unless deployType=ALL
- `deploy_type` (String) Allowed values:
		- ALL,
		- ND,
		- NDG
- `name` (String)
- `sgt` (String) Mandatory unless mappingGroup is set

Read-Only:

- `link` (List of Object) (see [below for nested schema](#nestedatt--parameters--link))

<a id="nestedatt--parameters--link"></a>
### Nested Schema for `parameters.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `deploy_to` (String)
- `deploy_type` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `name` (String)
- `sgt` (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_sg_mapping_group.example "id:=string"
terraform import ciscoise_sg_mapping_group.example "name:=string"
```
