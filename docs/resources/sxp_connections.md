---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_sxp_connections Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on SXPConnections.
  This resource allows the client to update a SXP connection.This resource deletes a SXP connection.This resource creates a SXP connection.
---

# ciscoise_sxp_connections (Resource)

It manages create, read, update and delete operations on SXPConnections.

- This resource allows the client to update a SXP connection.

- This resource deletes a SXP connection.

- This resource creates a SXP connection.

## Example Usage

```terraform
resource "ciscoise_sxp_connections" "example" {
  provider = ciscoise
  parameters {

    description = "string"
    enabled     = "false"
    id          = "string"
    ip_address  = "string"
    sxp_mode    = "string"
    sxp_node    = "string"
    sxp_peer    = "string"
    sxp_version = "string"
    sxp_vpn     = "string"
  }
}

output "ciscoise_sxp_connections_example" {
  value = ciscoise_sxp_connections.example
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
- `enabled` (String)
- `ip_address` (String)
- `sxp_mode` (String)
- `sxp_node` (String)
- `sxp_peer` (String)
- `sxp_version` (String)
- `sxp_vpn` (String)

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
- `enabled` (String)
- `id` (String)
- `ip_address` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `sxp_mode` (String)
- `sxp_node` (String)
- `sxp_peer` (String)
- `sxp_version` (String)
- `sxp_vpn` (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_sxp_connections.example "id:=string"
```
