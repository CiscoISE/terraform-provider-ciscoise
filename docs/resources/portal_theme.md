---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_portal_theme Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on PortalTheme.
  This resource allows the client to update a portal theme by ID.This resource deletes a portal theme by ID.This resource creates a portal theme.
---

# ciscoise_portal_theme (Resource)

It manages create, read, update and delete operations on PortalTheme.

- This resource allows the client to update a portal theme by ID.

- This resource deletes a portal theme by ID.

- This resource creates a portal theme.

## Example Usage

```terraform
resource "ciscoise_portal_theme" "example" {
  provider = ciscoise
  parameters {

    description = "string"
    id          = "string"
    name        = "string"
    theme_data  = "string"
  }
}

output "ciscoise_portal_theme_example" {
  value = ciscoise_portal_theme.example
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
- `theme_data` (String) Portal Theme for all portals

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
- `theme_data` (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_portal_theme.example "id:=string"
```
