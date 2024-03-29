---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_identity_group Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read and update operations on IdentityGroups.
  This resource allows the client to update an identity group.This resource creates an identity group.
---

# ciscoise_identity_group (Resource)

It manages create, read and update operations on IdentityGroups.

- This resource allows the client to update an identity group.

- This resource creates an identity group.

## Example Usage

```terraform
resource "ciscoise_identity_group" "example" {
  provider = ciscoise
  parameters {

    description = "string"
    id          = "string"
    name        = "string"
    parent      = "string"
  }
}

output "ciscoise_identity_group_example" {
  value = ciscoise_identity_group.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String) Unix timestamp records the last time that the resource was updated.

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `description` (String)
- `name` (String)
- `parent` (String)

Read-Only:

- `id` (String) The ID of this resource.
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

- `description` (String)
- `id` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `name` (String)
- `parent` (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_identity_group.example "id:=string\name:=string"
```
