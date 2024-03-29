---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_personas_check_standalone Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs create operation on Network Access - Authentication Rules.
  - Network Access Reset HitCount for Authentication Rules
---

# ciscoise_personas_check_standalone (Resource)

It performs create operation on Network Access - Authentication Rules.
- Network Access Reset HitCount for Authentication Rules

## Example Usage

```terraform
resource "ciscoise_personas_check_standalone" "check_standalone" {
  parameters {
    ip       = "string"
    username = "string"
    password = "string"
    hostname = "string"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (String)
- `last_updated` (String) Unix timestamp records the last time that the resource was updated.

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `hostname` (String) Node hostname
- `ip` (String) Node Ip
- `password` (String) password
- `username` (String) username


